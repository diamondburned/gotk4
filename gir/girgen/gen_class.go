package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

var classTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .InterfaceName }}
	type {{ .InterfaceName }} interface {
		{{ range .Implements -}}
		{{ . }}
		{{ end }}

		{{ range .Methods }}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Name }}{{ .Tail -}}
		{{ end }}
	}

	// {{ .StructName }} implements the {{ .InterfaceName }} class.
	type {{ .StructName }} struct {
		{{ .ParentInterface }}
	}

	// Wrap{{ .InterfaceName }} wraps a GObject to the right type. It is
	// primarily used internally.
	func Wrap{{ .InterfaceName }}(obj *externglib.Object) {{ .InterfaceName }} {
		return {{ .Tree.WrapClass "obj" }}
	}

	{{ if .GLibGetType }}
	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return Wrap{{ .InterfaceName }}(obj), nil
	}
	{{ end }}

	{{ range .Constructors }}
	{{ GoDoc .Doc 0 .Name }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}

	{{ range .Methods }}
	func ({{ .Recv }} {{ $.StructName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}

	{{ range .InheritedMethods }}
	func ({{ .Recv }} {{ $.StructName }}) {{ .Name }}{{ .Tail }} {
		{{ if .Return }} return {{ end -}}
		{{ .Wrapper }}(gextras.InternObject({{ .Recv }})).{{ .Name }}({{ .CallParams }})
	}
	{{ end }}
`)

func (ng *NamespaceGenerator) generateClasses() {
	clgen := newClassGenerator(ng)

	for _, class := range ng.current.Namespace.Classes {
		if !class.IsIntrospectable() {
			continue
		}
		if ng.mustIgnore(&class.Name, &class.CType) {
			continue
		}
		if !clgen.Use(class) {
			continue
		}

		// Need for the object wrapper.
		ng.needsExternGLib()

		if class.GLibGetType != "" && !ng.mustIgnoreC(class.GLibGetType) {
			ng.addMarshaler(class.GLibGetType, clgen.InterfaceName)
		}

		ng.pen.WriteTmpl(classTmpl, &clgen)
	}
}

type classCallable struct {
	Doc   *gir.Doc
	Recv  string
	Name  string
	Tail  string
	Block string
}

// classInheritedMethod describes a method inherited from an interface (using
// the implements thing). The Use function must add imports for the wrappers if
// needed.
type classInheritedMethod struct {
	classCallable
	Return     bool
	Wrapper    string
	CallParams string
}

func newClassCallable(cgen *callableGenerator) classCallable {
	return classCallable{
		Doc:   cgen.Doc,
		Recv:  cgen.Recv(),
		Name:  cgen.Name,
		Tail:  cgen.Tail,
		Block: cgen.Block,
	}
}

type classGenerator struct {
	gir.Class
	StructName      string
	InterfaceName   string
	ParentInterface string

	Implements       []string // list of PUBLIC INTERFACE NAMES
	Constructors     []classCallable
	Methods          []classCallable
	InheritedMethods []classInheritedMethod

	Tree TypeTree // starts from current resolved type

	ng   *NamespaceGenerator
	cgen *callableGenerator
	igen *ifaceGenerator
}

func newClassGenerator(ng *NamespaceGenerator) classGenerator {
	cgen := newCallableGenerator(ng)
	igen := newIfaceGenerator(ng)
	return classGenerator{
		ng:   ng,
		cgen: &cgen,
		igen: &igen,
	}
}

func (cg *classGenerator) Use(class gir.Class) bool {
	cg.Tree = cg.ng.TypeTree()
	cg.Tree.Level = 2

	if class.Parent == "" {
		// TODO: check what happens if a class has no parent. It should have a
		// GObject parent, usually.
		cg.Logln(LogSkip, "class", class.Name, "because it has no parents")
		return false
	}

	if !cg.Tree.Resolve(class.Name) {
		cg.Logln(LogSkip, "class", class.Name, "because unknown parent type", class.Parent)
		return false
	}

	cg.Class = class
	cg.InterfaceName = PascalToGo(class.Name)
	cg.StructName = UnexportPascal(cg.InterfaceName)

	cg.Implements = cg.Tree.PublicEmbeds()
	cg.ParentInterface = GoPublicType(cg.ng, cg.Tree.Requires[0].ResolvedType)

	// Add imports for the parent class in the struct.
	cg.ng.importResolvedType(cg.Tree.Requires[0].PublImport)
	// Add imports for the embedded interfaces.
	for _, imp := range cg.Tree.Embeds {
		cg.ng.importResolvedType(imp.PublImport)
	}

	cg.Methods = classCallableGrow(cg.Methods, len(class.Methods))
	cg.Constructors = classCallableGrow(cg.Constructors, len(class.Constructors))
	cg.InheritedMethods = cg.InheritedMethods[:0]

	// Initialize the Callable constructor generator.
	cg.cgen.ReturnWrap = "Wrap" + cg.InterfaceName

	for _, ctor := range class.Constructors {
		ctor = bodgeClassCtor(class, ctor)
		if !cg.cgen.Use(ctor.CallableAttrs) {
			continue
		}
		cg.Constructors = append(cg.Constructors, newClassCallable(cg.cgen))
	}

	// Reset the ReturnWrap for methods.
	cg.cgen.ReturnWrap = ""

	// The first requirement type is always the parent class type.
	for _, impl := range cg.Tree.Requires[1:] {
		iface := impl.ResolvedType.Extern.Result.Interface
		if iface == nil {
			cg.Logln(LogUnusuality,
				"implemented type", impl.ResolvedType.PublicType(true), "not interface")
			continue
		}

		if !cg.igen.Use(*iface) {
			continue
		}

		needsNamespace := impl.ResolvedType.NeedsNamespace(cg.ng.current)
		wrapper := impl.ResolvedType.WrapName(needsNamespace)

		for _, method := range cg.igen.Methods {
			// Parse the parameter values out of the function in a pretty hacky
			// way by extracting the types out.
			params := append([]string(nil), method.goArgs.Joints()...)
			for i, word := range params {
				params[i] = strings.SplitN(word, " ", 2)[0]
			}
			cg.InheritedMethods = append(cg.InheritedMethods, classInheritedMethod{
				classCallable: newClassCallable(&method),
				Return:        method.goRets.Len() > 0,
				Wrapper:       wrapper,
				CallParams:    strings.Join(params, ", "),
			})
		}
	}

	if len(cg.InheritedMethods) > 0 {
		cg.ng.addImportInternal("gextras")
	}

	// Only generate the methods after we've generated the inherited methods,
	// because we want the class to satisfy the inherited interfaces as a first
	// priority, and then generate our methods down below. This way, we can also
	// avoid method collisions.
	for _, method := range class.Methods {
		if !cg.cgen.Use(method.CallableAttrs) {
			continue
		}
		cg.Methods = append(cg.Methods, newClassCallable(cg.cgen))
	}

	// Rename all methods to have idiomatic and non-colliding names, if possible.
	for i, method := range cg.Methods {
		newName, isGetter := renameGetter(method.Name)
		isDuplicate := cg.hasField(newName)

		// Avoid duplicating field names with inherited interfaces, including
		// Objector, but only this if we're not renaming a getter (since that's
		// not important).
		if isDuplicate && !isGetter {
			newName += cg.InterfaceName
			isDuplicate = cg.hasField(newName)
		}

		if !isDuplicate {
			cg.Methods[i].Name = newName
		}
	}

	return true
}

func (cg *classGenerator) hasField(goName string) bool {
	if isObjectorMethod(goName) || cg.ParentInterface == goName {
		return true
	}
	for _, impl := range cg.Implements {
		if impl == goName {
			return true
		}
	}
	for _, callable := range cg.Methods {
		if callable.Name == goName {
			return true
		}
	}
	for _, parent := range cg.InheritedMethods {
		if parent.Name == goName {
			return true
		}
	}
	return false
}

// bodgeClassCtor bodges the given constructor to return exactly the class type
// instead of any other. It returns the original ctor if the conditions don't
// match for bodging.
//
// We have to do this to work around some cases where widget constructors would
// return the widget class instead of the actual class.
func bodgeClassCtor(class gir.Class, ctor gir.Constructor) gir.Constructor {
	if ctor.ReturnValue == nil || ctor.ReturnValue.Type == nil {
		return ctor
	}

	retVal := *ctor.ReturnValue
	retTyp := *retVal.AnyType.Type

	// Note: this has caused me quite a lot of trouble. It's probably wrong as
	// well. The whole point is to work around the C API's weird class typing.
	retTyp.CType = moveCPtr(class.CType, retTyp.CType)

	retTyp.Name = class.Name
	retTyp.Introspectable = class.Introspectable
	retTyp.AnyType = gir.AnyType{}

	retVal.AnyType.Type = &retTyp
	ctor.ReturnValue = &retVal

	ctor.Name = strings.TrimPrefix(ctor.Name, "new")
	ctor.Name = strings.TrimPrefix(ctor.Name, "_")
	if ctor.Name != "" {
		ctor.Name = "_" + ctor.Name
	}

	ctor.Name = "new_" + class.Name + ctor.Name

	return ctor
}

func (cg *classGenerator) Logln(lvl LogLevel, v ...interface{}) {
	v = append(v, nil)
	copy(v[1:], v)
	v[0] = fmt.Sprintf("class %s (C.%s):", cg.InterfaceName, cg.CType)

	cg.ng.Logln(lvl, v...)
}

func classCallableGrow(callables []classCallable, n int) []classCallable {
	if cap(callables) <= n {
		return callables[:0]
	}
	return make([]classCallable, 0, n*2)
}

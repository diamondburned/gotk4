package generators

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var classTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	type {{ .InterfaceName }} interface {
		{{ range .Implements -}}
		{{ . }}
		{{ end }}

		{{ range .Methods }}
		{{ GoDoc . 1 }}
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
	{{ GoDoc . 0 }}
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

// GenerateClass generates the given class into files.
func GenerateClass(gen FileGeneratorWriter, class *gir.Class) bool {
	classGen := NewClassGenerator(gen)
	if !classGen.Use(class) {
		return false
	}

	writer := FileWriterFromType(gen, class)

	if class.GLibGetType != "" && !types.FilterCType(gen, class.GLibGetType) {
		writer.Header().AddMarshaler(class.GLibGetType, classGen.InterfaceName)
	}

	// Need for the object wrapper.
	writer.Header().NeedsExternGLib()

	file.ApplyHeader(writer, &classGen)
	writer.Pen().WriteTmpl(classTmpl, &classGen)
	return true
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

func newClassCallable(cgen *callable.Generator) classCallable {
	return classCallable{
		Doc:   cgen.Doc,
		Recv:  cgen.Recv(),
		Name:  cgen.Name,
		Tail:  cgen.Tail,
		Block: cgen.Block,
	}
}

type ClassGenerator struct {
	*gir.Class
	StructName      string
	InterfaceName   string
	ParentInterface string

	Implements       []string // list of PUBLIC INTERFACE NAMES
	Constructors     []classCallable
	Methods          []classCallable
	InheritedMethods []classInheritedMethod

	Tree types.Tree // starts from current resolved type

	header file.Header
	gen    FileGenerator
	igen   InterfaceGenerator
	cgen   callable.Generator
}

// NewClassGenerator creates a new class generator.
func NewClassGenerator(gen FileGenerator) ClassGenerator {
	classGenerator := ClassGenerator{
		gen:  gen,
		igen: NewInterfaceGenerator(gen),
		Tree: types.NewTree(gen),
	}
	classGenerator.cgen = callable.NewGenerator(headeredFileGenerator{
		FileGenerator: gen,
		Headerer:      &classGenerator,
	})

	return classGenerator
}

func (cg *ClassGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("class %s (C.%s):", cg.InterfaceName, cg.CType)
	cg.gen.Logln(lvl, logger.Prefix(v, p)...)
}

// Reset resets the callback generator.
func (g *ClassGenerator) Reset() {
	g.igen.Reset()
	g.cgen.Reset()
	g.Tree.Reset()

	*g = ClassGenerator{
		gen:  g.gen,
		igen: g.igen,
		cgen: g.cgen,
		Tree: g.Tree,
	}
}

// Header returns the callback generator's current header.
func (g *ClassGenerator) Header() *file.Header {
	return &g.header
}

func (cg *ClassGenerator) Use(class *gir.Class) bool {
	cg.Reset()
	cg.Class = class

	if !class.IsIntrospectable() || types.Filter(cg.gen, class.Name, class.CType) {
		return false
	}

	// Skip generating built-in types as well as orphaned types.
	if !cg.Tree.Resolve(class.Name) || cg.Tree.Builtin != nil {
		cg.Logln(logger.Debug, "class", class.Name, "because unknown parent type", class.Parent)
		return false
	}

	cg.InterfaceName = strcases.PascalToGo(class.Name)
	cg.StructName = strcases.UnexportPascal(cg.InterfaceName)

	cg.Implements = cg.Tree.PublicEmbeds()
	cg.ParentInterface = types.GoPublicType(cg.gen, cg.Tree.Requires[0].Resolved)

	// Add imports for the parent class in the struct.
	cg.header.ImportResolvedType(cg.Tree.Requires[0].PublImport)
	// Add imports for the embedded interfaces.
	for _, imp := range cg.Tree.Embeds {
		cg.header.ImportResolvedType(imp.PublImport)
	}

	cg.Methods = classCallableGrow(cg.Methods, len(class.Methods))
	cg.Constructors = classCallableGrow(cg.Constructors, len(class.Constructors))
	cg.InheritedMethods = cg.InheritedMethods[:0]

	// Initialize the Callable constructor generator.
	cg.cgen.ReturnWrap = "Wrap" + cg.InterfaceName

	for _, ctor := range class.Constructors {
		// Bodge this so the constructors and stuff are named properly. This
		// copies things safely, so class is not modified.
		ctor := bodgeClassCtor(class, ctor)
		if !cg.cgen.Use(&ctor.CallableAttrs) {
			continue
		}
		cg.Constructors = append(cg.Constructors, newClassCallable(&cg.cgen))
	}

	// Reset the ReturnWrap for methods.
	cg.cgen.ReturnWrap = ""

	// The first requirement type is always the parent class type.
	cg.Tree.WalkPublInterfaces(func(typ *types.Resolved) {
		// This fucking sucks. I fucking hate this. Because the GNOME people
		// wants to maximize misery and make your life a fucking pain, the
		// methods inside this interface aren't namespaced properly, so we have
		// no choice but to override the global namespace.
		//
		// Ideally, once the file generator is restored, the type resolver
		// should be decoupled from the namespace generator as a whole, and
		// instead, the namespace should be overrideable by having TypeResolver
		// returning the namespace we wants and ifaceGenerator (and everything
		// else) to use TypeResolver for resolving and generating methods.

		cg.igen.UseMethods(typ.Extern.Type.(*gir.Interface), typ.Extern.NamespaceFindResult)
		file.ApplyHeader(cg, &cg.igen)

		needsNamespace := typ.NeedsNamespace(cg.gen.Namespace())
		wrapper := typ.WrapName(needsNamespace)

		for _, method := range cg.igen.Methods {
			// Parse the parameter values out of the function in a pretty hacky
			// way by extracting the types out.
			params := append([]string(nil), method.GoArgs.Joints()...)
			for i, word := range params {
				params[i] = strings.SplitN(word, " ", 2)[0]
			}

			cg.InheritedMethods = append(cg.InheritedMethods, classInheritedMethod{
				classCallable: newClassCallable(&method),
				Return:        method.GoRets.Len() > 0,
				Wrapper:       wrapper,
				CallParams:    strings.Join(params, ", "),
			})
		}
	})

	if len(cg.InheritedMethods) > 0 {
		cg.header.ImportCore("gextras")
	}

	// Only generate the methods after we've generated the inherited methods,
	// because we want the class to satisfy the inherited interfaces as a first
	// priority, and then generate our methods down below. This way, we can also
	// avoid method collisions.
	for _, method := range class.Methods {
		if types.FilterMethod(cg.gen, cg.Class.Name, &method) {
			continue
		}
		if !cg.cgen.Use(&method.CallableAttrs) {
			continue
		}
		file.ApplyHeader(cg, &cg.cgen)
		cg.Methods = append(cg.Methods, newClassCallable(&cg.cgen))
	}

	// Rename all methods to have idiomatic and non-colliding names, if possible.
	for i, method := range cg.Methods {
		newName, isGetter := callable.RenameGetter(method.Name)
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

func (cg *ClassGenerator) hasField(goName string) bool {
	if types.IsObjectorMethod(goName) || cg.ParentInterface == goName {
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
func bodgeClassCtor(class *gir.Class, ctor gir.Constructor) gir.Constructor {
	if ctor.ReturnValue == nil || ctor.ReturnValue.Type == nil {
		return ctor
	}

	retVal := *ctor.ReturnValue
	retTyp := *retVal.AnyType.Type

	// Note: this has caused me quite a lot of trouble. It's probably wrong as
	// well. The whole point is to work around the C API's weird class typing.
	retTyp.CType = types.MoveCPtr(class.CType, retTyp.CType)

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

func classCallableGrow(callables []classCallable, n int) []classCallable {
	if cap(callables) <= n {
		return callables[:0]
	}
	return make([]classCallable, 0, n*2)
}

package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

var classTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .Name }}
	type {{ .Name }} struct {
		{{ range .TypeTree.Children -}}
		{{ . }}
		{{ end }}
	}

	// {{ .Name }}Class is an interface that the {{ .Name }} class always
	// implements. It is only used for parameters that take in not just this
	// class but any other class that extends it.
	type {{ .Name }}Class interface {
		gextras.Objector
		_{{ UnexportPascal .Name }}()
	}

	func ({{ .Name }}) _{{ UnexportPascal .Name }}() {}

	{{ if .GLibGetType }}
	func marshal{{ .Name }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return Wrap{{ .Name }}(obj), nil
	}
	{{ end }}

	{{ range .Constructors }}
	{{ GoDoc .Doc 0 .Name }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}

	{{ range .Methods }}
	{{ GoDoc .Doc 0 .Name }}
	func ({{ .Recv }} {{ $.Name }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

type classCallable struct {
	Doc   *gir.Doc
	Recv  string
	Name  string
	Tail  string
	Block string
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

func classCallableGrow(callables []classCallable, n int) []classCallable {
	if cap(callables) <= n {
		return callables[:0]
	}
	return make([]classCallable, 0, n*2)
}

type classGenerator struct {
	gir.Class
	Name string

	Methods      []classCallable
	Constructors []classCallable

	TypeTree TypeTree // starts from current resolved type

	ng   *NamespaceGenerator
	cgen *callableGenerator
}

func newClassGenerator(ng *NamespaceGenerator) *classGenerator {
	cgen := newCallableGenerator(ng)
	return &classGenerator{
		ng:   ng,
		cgen: &cgen,
	}
}

func (cg *classGenerator) Use(class gir.Class) bool {
	cg.TypeTree = cg.ng.TypeTree()
	// cg.TypeTree.Level = 2

	if class.Parent == "" {
		// TODO: check what happens if a class has no parent. It should have a
		// GObject parent, usually.
		cg.Logln(LogSkip, "class", class.Name, "because it has no parents")
		return false
	}

	cg.Class = class
	cg.Name = PascalToGo(class.Name)

	if !cg.TypeTree.Resolve(class.Name) {
		cg.Logln(LogSkip, "class", class.Name, "because unknown parent type", class.Parent)
		return false
	}

	cg.TypeTree.ImportChildren(cg.ng)

	cg.Methods = classCallableGrow(cg.Methods, len(class.Methods))
	cg.Constructors = classCallableGrow(cg.Constructors, len(class.Constructors))

	// Initialize the Callable constructor generator.
	cg.cgen.ManualWrap = true

	for _, ctor := range class.Constructors {
		ctor = bodgeClassCtor(class, ctor)
		if !cg.cgen.Use(ctor.CallableAttrs) {
			continue
		}
		cg.Constructors = append(cg.Constructors, newClassCallable(cg.cgen))
	}

	// Reset the ReturnWrap for methods.
	cg.cgen.ManualWrap = false

	for _, method := range class.Methods {
		if !cg.cgen.Use(method.CallableAttrs) {
			continue
		}
		cg.Methods = append(cg.Methods, newClassCallable(cg.cgen))
	}

	// Rename all methods to have idiomatic getter names if possible.
	for i, callable := range cg.Methods {
		newName := renameGetter(callable.Name)

		// Avoid duplicating method names with Objector.
		// TODO: account for other interfaces as well.
		if isObjectorMethod(newName) {
			newName += cg.Name
		}

		if !cg.hasField(newName) {
			cg.Methods[i].Name = newName
		}
	}

	return true
}

func (cg *classGenerator) hasField(goName string) bool {
	for _, callable := range cg.Methods {
		if callable.Name == goName {
			return true
		}
	}

	for _, parent := range cg.TypeTree.Requires {
		if parent.Resolved.GoType(false) == goName {
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
	v[0] = fmt.Sprintf("class %s (C.%s):", cg.Name, cg.CType)

	cg.ng.Logln(lvl, v...)
}

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
		// Need for the Class interface.
		ng.addImportInternal("gextras")

		if class.GLibGetType != "" && !ng.mustIgnoreC(class.GLibGetType) {
			ng.addMarshaler(class.GLibGetType, clgen.Name)
		}

		ng.pen.WriteTmpl(classTmpl, &clgen)
	}
}

func renameGetterMethod(all []gir.Method, method gir.Method) string {
	newName := renameGetter(method.Name)

	for _, m := range all {
		if m.Name == newName {
			return method.Name
		}
	}

	return newName
}

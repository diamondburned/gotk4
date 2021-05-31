package girgen

import (
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

var classTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .InterfaceName }}
	type {{ .InterfaceName }} interface {
		{{ range .TypeTree.PublicChildren -}}
		{{ . }}
		{{ end }}

		{{ range .Methods }}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Name }}{{ .Tail -}}
		{{ end }}
	}

	// {{ .StructName }} implements the {{ .InterfaceName }} interface.
	type {{ .StructName }} struct {
		{{ range .TypeTree.PublicChildren -}}
		{{ . }}
		{{ end }}
	}

	var _ {{ .InterfaceName }} = (*{{ .StructName }})(nil)

	// Wrap{{ .InterfaceName }} wraps a GObject to the right type. It is
	// primarily used internally.
	func Wrap{{ .InterfaceName }}(obj *externglib.Object) {{ .InterfaceName }} {
		return {{ .TypeTree.Wrap "obj" }}
	}

	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return Wrap{{ .InterfaceName }}(obj), nil
	}

	{{ range .Constructors }}
	{{ if $.UseConstructor . }}
	// {{ $.Callable.Name }} constructs a class {{ $.InterfaceName }}.
	func {{ $.Callable.Name }}{{ $.Callable.Tail }} {{ $.Callable.Block }}
	{{ end }}
	{{ end }}

	{{ range .Methods }}
	{{ GoDoc .Doc 1 .Name }}
	func ({{ .Recv }} {{ $.StructName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

type classGenerator struct {
	gir.Class
	StructName    string
	InterfaceName string

	TypeTree TypeTree // starts from current resolved type
	Methods  []callableGenerator

	Callable callableGenerator

	Ng *NamespaceGenerator
}

func newClassGenerator(ng *NamespaceGenerator) *classGenerator {
	return &classGenerator{
		TypeTree: *ng.TypeTree(),
		Callable: callableGenerator{Ng: ng},
		Ng:       ng,
	}
}

func (cg *classGenerator) Use(class gir.Class) bool {
	if class.Parent == "" {
		// TODO: check what happens if a class has no parent. It should have a
		// GObject parent, usually.
		return false
	}

	cg.Class = class
	cg.InterfaceName = PascalToGo(class.Name)
	cg.StructName = UnexportPascal(cg.InterfaceName)

	resolved := TypeFromResult(cg.Ng, gir.TypeFindResult{Class: &class})
	if !cg.TypeTree.ResolveFromType(resolved) {
		return false
	}

	cg.Methods = callableGrow(cg.Methods, len(class.Methods))

	for _, method := range class.Methods {
		cbgen := newCallableGenerator(cg.Ng)
		if !cbgen.Use(method.CallableAttrs) {
			continue
		}

		cg.Methods = append(cg.Methods, cbgen)
	}

	// Use Go-idiomatic getter names, unless there's a duplicate.
	callableRenameGetters(cg.Methods)

	return true
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

	retTyp.Name = class.Name
	retTyp.CType = class.CType
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

func (cg *classGenerator) UseConstructor(ctor gir.Constructor) bool {
	ctor = bodgeClassCtor(cg.Class, ctor)
	return cg.Callable.Use(ctor.CallableAttrs)
}

func (ng *NamespaceGenerator) generateClasses() {
	cg := newClassGenerator(ng)

	for _, class := range ng.current.Namespace.Classes {
		if !cg.Use(class) {
			ng.logln(logInfo, "skipping class", class.Name)
			continue
		}

		ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
		ng.pen.BlockTmpl(classTmpl, &cg)
	}
}

package girgen

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
)

// TODO: unexported type implementation
// TODO: methods for implementation

var interfaceTmpl = newGoTemplate(`
	{{ if .Virtuals }}
	// {{ .InterfaceName }}Overrider contains methods that are overridable. This
	// interface is a subset of the interface {{ .InterfaceName }}.
	type {{ .InterfaceName }}Overrider interface {
		{{ range .Virtuals -}}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Name }}{{ .Tail }}
		{{ end }}
	}
	{{ end }}

	{{ GoDoc .Doc 0 .InterfaceName }}
	type {{ .InterfaceName }} interface {
		{{ range .TypeTree.PublicChildren -}}
		{{ . }}
		{{ end }}

		{{ range .Methods -}}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Name }}{{ .Tail }}
		{{ end }}
	}

	// {{ .StructName }} implements the {{ .InterfaceName }} interface.
	type {{ .StructName }} struct {
		{{ range .TypeTree.PublicChildren -}}
		{{ . }}
		{{ end }}
	}

	var _ {{ .InterfaceName }} = (*{{ .StructName }})(nil)

	// Wrap{{ .InterfaceName }} wraps a GObject to a type that implements interface
	// {{ .InterfaceName }}. It is primarily used internally.
	func Wrap{{ .InterfaceName }}(obj *externglib.Object) {{ .InterfaceName }} {
		return {{ .TypeTree.Wrap "obj" }}
	}

	{{ if .GLibGetType }}
	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return Wrap{{ .InterfaceName }}(obj), nil
	}
	{{ end }}

	{{ range .Methods }}
	func ({{ .Recv }} {{ $.StructName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

type ifaceGenerator struct {
	gir.Interface
	InterfaceName string
	StructName    string

	TypeTree TypeTree
	Virtuals []callableGenerator // for overrider
	Methods  []callableGenerator // for big interface

	ng *NamespaceGenerator
}

type ifaceMethod struct {
	*callableGenerator
	StructName string
}

func newIfaceGenerator(ng *NamespaceGenerator) *ifaceGenerator {
	return &ifaceGenerator{
		ng: ng,
	}
}

func (ig *ifaceGenerator) Use(iface gir.Interface) bool {
	ig.TypeTree = ig.ng.TypeTree()
	ig.TypeTree.Level = 2

	ig.Interface = iface
	ig.InterfaceName = PascalToGo(iface.Name)
	ig.StructName = UnexportPascal(ig.InterfaceName)

	if !ig.TypeTree.Resolve(iface.Name) {
		ig.Logln(LogSkip, "cannot be type-resolved")
		return false
	}

	ig.TypeTree.ImportChildren(ig.ng)
	ig.updateMethods()

	return true
}

func (ig *ifaceGenerator) updateMethods() {
	ig.Methods = callableGrow(ig.Methods, len(ig.Interface.Methods))
	ig.Virtuals = callableGrow(ig.Virtuals, len(ig.Interface.VirtualMethods))

	for _, vmethod := range ig.Interface.VirtualMethods {
		cbgen := newCallableGenerator(ig.ng)
		if !cbgen.Use(vmethod.CallableAttrs) {
			continue
		}

		ig.Virtuals = append(ig.Virtuals, cbgen)
	}

	for _, method := range ig.Interface.Methods {
		cbgen := newCallableGenerator(ig.ng)
		if !cbgen.Use(method.CallableAttrs) {
			continue
		}

		ig.Methods = append(ig.Methods, cbgen)
	}

	callableRenameGetters(ig.InterfaceName, ig.Methods)
	callableRenameGetters(ig.InterfaceName, ig.Virtuals)
}

func (ig *ifaceGenerator) Logln(lvl LogLevel, v ...interface{}) {
	v = append(v, nil)
	copy(v[1:], v)
	v[0] = fmt.Sprintf("interface %s (C.%s):", ig.InterfaceName, ig.CType)

	ig.ng.Logln(lvl, v...)
}

func (ng *NamespaceGenerator) generateIfaces() {
	ig := newIfaceGenerator(ng)

	for _, iface := range ng.current.Namespace.Interfaces {
		if !iface.IsIntrospectable() {
			continue
		}
		if ng.mustIgnore(&iface.Name, &iface.CType) {
			continue
		}
		if !ig.Use(iface) {
			continue
		}

		if iface.GLibGetType != "" && !ng.mustIgnoreC(iface.GLibGetType) {
			ng.addMarshaler(iface.GLibGetType, ig.InterfaceName)
		}

		ng.pen.WriteTmpl(interfaceTmpl, &ig)
	}
}

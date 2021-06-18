package girgen

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
)

// TODO: unexported type implementation
// TODO: methods for implementation

// TODO: find a proper way to add an -er suffix.

var interfaceTmpl = newGoTemplate(`
	{{ if .Virtuals }}
	// {{ .Name }}Interface contains virtual methods for {{ .Name }}, or
	// methods that can be overridden.
	type {{ .Name }}Interface interface {
		gextras.Objector

		{{ range .Virtuals -}}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Name }}{{ .Tail }}
		{{ end }}
	}
	{{ end }}

	{{ GoDoc .Doc 0 .Name }}
	type {{ .Name }} struct {
		{{ range .TypeTree.Children -}}
		{{ . }}
		{{ end }}
	}

	{{ if .GLibGetType }}
	func marshal{{ .Name }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return {{ .TypeTree.Wrap "obj" }}, nil
	}
	{{ end }}

	{{ range .Methods }}
	{{ GoDoc .Doc 0 .Name }}
	func ({{ .Recv }} {{ $.Name }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

type ifaceGenerator struct {
	gir.Interface
	Name string

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
	// ig.TypeTree.Level = 4

	ig.Interface = iface
	ig.Name = PascalToGo(iface.Name)

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

	if len(ig.Virtuals) > 0 {
		ig.ng.needsExternGLib()
	}

	callableRenameGetters(ig.Name, ig.Methods)
	callableRenameGetters(ig.Name, ig.Virtuals)
}

func (ig *ifaceGenerator) Logln(lvl LogLevel, v ...interface{}) {
	v = append(v, nil)
	copy(v[1:], v)
	v[0] = fmt.Sprintf("interface %s (C.%s):", ig.Name, ig.CType)

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
			ng.addMarshaler(iface.GLibGetType, ig.Name)
		}

		ng.pen.WriteTmpl(interfaceTmpl, &ig)
	}
}

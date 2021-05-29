package girgen

import (
	"github.com/diamondburned/gotk4/gir"
)

// TODO: unexported type implementation
// TODO: methods for implementation
// TODO: wrap GObject into implementation
// TODO: Go->C and C->Go conversions for implementation

var interfaceTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .GoName }}
	type {{ .GoName }} interface {
		{{- range .TypeTree.PublicChildren }}
		{{ . }}
		{{- else }}
		gextras.Objector
		{{ end }}

		{{ range .Virtuals -}}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Name }}{{ .Tail }}
		{{ end }}
	}

	// Wrap{{ .GoName }} wraps a GObject to a type that implements interface
	// {{ .GoName }}. It is primarily used internally.
	func Wrap{{ .GoName }}(obj *externglib.Object) {{ .GoName }} {
	}
`)

type ifaceGenerator struct {
	gir.Interface
	GoName   string
	TypeTree TypeTree
	Virtuals []callableGenerator // for interface
	Methods  []callableGenerator // for object implementation

	Ng *NamespaceGenerator
}

func (ig *ifaceGenerator) Use(iface gir.Interface) bool {
	ig.Interface = iface
	ig.GoName = PascalToGo(iface.Name)

	ig.updateMethods()
	if len(ig.Virtuals) == 0 {
		return false
	}

	resolved := TypeFromResult(ig.Ng, gir.TypeFindResult{Interface: &iface})
	if !ig.TypeTree.ResolveFromType(resolved) {
		return false
	}

	return true
}

func (ig *ifaceGenerator) updateMethods() {
	ig.Methods = callableGrow(ig.Methods, len(ig.Interface.Methods))
	ig.Virtuals = callableGrow(ig.Virtuals, len(ig.Interface.VirtualMethods))

	for _, vmethod := range ig.Interface.VirtualMethods {
		cbgen := newCallableGenerator(ig.Ng)
		if !cbgen.Use(vmethod.CallableAttrs) {
			continue
		}

		cbgen.Parent = ig.GoName
		ig.Virtuals = append(ig.Virtuals, cbgen)
	}

	for _, method := range ig.Interface.Methods {
		cbgen := newCallableGenerator(ig.Ng)
		if !cbgen.Use(method.CallableAttrs) {
			continue
		}

		cbgen.Parent = ig.GoName
		ig.Methods = append(ig.Methods, cbgen)
	}

	callableRenameGetters(ig.Methods)
	callableRenameGetters(ig.Virtuals)
}

func (ng *NamespaceGenerator) generateIfaces() {
	ig := ifaceGenerator{Ng: ng}

	for _, iface := range ng.current.Namespace.Interfaces {
		if !ig.Use(iface) {
			ng.logln(logInfo, "skipping interface", iface.Name)
			continue
		}

		ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
		ng.pen.BlockTmpl(interfaceTmpl, &ig)
	}
}

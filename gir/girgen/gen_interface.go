package girgen

import "github.com/diamondburned/gotk4/gir"

var interfaceTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .GoName }}
	type {{ .GoName }} interface {
		{{ range .Calls -}}
		{{ . }}
		{{ end }}
	}
`)

type ifaceGenerator struct {
	gir.Interface
	GoName string
	Calls  []string // functions

	Ng *NamespaceGenerator
}

func (ig *ifaceGenerator) Use(iface gir.Interface) bool {
	ig.Interface = iface
	ig.GoName = PascalToGo(iface.Name)
	ig.updateMethods()
	return len(ig.Calls) > 0
}

func (ig *ifaceGenerator) updateMethods() {
	ig.Calls = ig.Calls[:0]

	for _, method := range ig.Interface.Methods {
		call := ig.Ng.FnCall(method.CallableAttrs)
		if call == "" {
			continue
		}

		ig.Calls = append(ig.Calls, SnakeToGo(true, method.Name)+call)
	}
}

func (ng *NamespaceGenerator) generateIfaces() {
	ig := ifaceGenerator{Ng: ng}

	for _, iface := range ng.current.Namespace.Interfaces {
		if !ig.Use(iface) {
			ng.logln(logInfo, "skipping interface", iface.Name)
			continue
		}

		ng.pen.BlockTmpl(interfaceTmpl, &ig)
	}
}
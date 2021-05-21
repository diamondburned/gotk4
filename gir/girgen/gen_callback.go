package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

var callbackTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .Name }}
	type {{ .GoName }} func{{ .Tail }}

	//export c{{ .GoName }}
	func c{{ .GoName }}{{ .CTail }}
`)

type callbackGenerator struct {
	gir.Callback
	GoName string
	CTail  string
	Tail   string

	Ng *NamespaceGenerator
}

func newCallbackGenerator(ng *NamespaceGenerator) callbackGenerator {
	return callbackGenerator{Ng: ng}
}

// Use sets the callback generator to the given GIR callback.
func (fg *callbackGenerator) Use(cb gir.Callback) bool {
	call := fg.Ng.FnCall(cb.CallableAttrs)
	if call == "" {
		return false
	}

	fg.GoName = PascalToGo(cb.Name)
	fg.Tail = call
	fg.CTail = "()"

	if cb.Parameters != nil {
		ctail := make([]string, 0, len(cb.Parameters.Parameters))

		for i, callback := range cb.Parameters.Parameters {
			ctype := anyTypeCGo(callback.AnyType)
			if ctype == "" {
				return false // probably var_args
			}

			ctail = append(ctail, fmt.Sprintf("arg%d %s", i, ctype))
		}

		fg.CTail = "(" + strings.Join(ctail, ", ") + ")"
	}

	if !returnIsVoid(cb.ReturnValue) {
		ctype := anyTypeCGo(cb.ReturnValue.AnyType)
		if ctype == "" {
			return false
		}

		fg.CTail += " " + ctype
	}

	return true
}

func (ng *NamespaceGenerator) generateCallbacks() {
	cg := newCallbackGenerator(ng)

	for _, callback := range ng.current.Namespace.Callbacks {
		if !cg.Use(callback) {
			ng.logln(logInfo, "skipping callback", callback.Name)
			continue
		}

		ng.pen.BlockTmpl(callbackTmpl, &cg)
	}
}

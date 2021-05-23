package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

var callbackTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .Name }}
	type {{ .GoName }} func{{ .Tail }}

	//export c{{ .GoName }}
	func c{{ .GoName }}{{ .CTail }} {{ .CBlock }}
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

func (fg *callbackGenerator) CBlock() string {
	// We need a data parameter to access our callbacks; we can't do this if
	// there are no params.
	if fg.Parameters == nil || len(fg.Parameters.Parameters) == 0 {
		return ""
	}

	b := pen.NewBlock()

	// indices in fg.Parameters.Parameters
	var closureParam *int

	for _, param := range fg.Parameters.Parameters {
		if param.Closure != nil {
			closureParam = param.Closure
			break
		}
	}

	// No data parameter to use; exit.
	if closureParam == nil {
		return ""
	}

	fg.Ng.addImport("github.com/diamondburned/gotk4/internal/box")

	b.Linef("v := box.Get(box.Callback, uintptr(arg%d))", *closureParam)
	b.Linef("if v == nil {")
	b.Linef("  panic(`callback not found`)")
	b.Linef("}")
	b.EmptyLine()

	argAt := func(i int) string { return fmt.Sprintf("arg%d", i) }
	goArgs := pen.NewJoints(", ", len(fg.Parameters.Parameters))
	goRets := pen.NewJoints(", ", len(fg.Parameters.Parameters)+1)

	iterateParams(fg.CallableAttrs, func(i int, param gir.Parameter) bool {
		goName := SnakeToGo(false, param.Name)
		goType, _ := fg.Ng.ResolveAnyType(param.AnyType, false)
		b.Linef("var %s %s", goName, goType)

		conv := fg.Ng.CGoConverter(argAt(i), goName, param.AnyType, argAt)
		b.Line(conv)
		b.EmptyLine()

		goArgs.Add(goName)
		return true
	})

	iterateReturns(fg.CallableAttrs, func(goName string, i int, typ gir.AnyType) bool {
		goRets.Add(goName)
		return true
	})

	b.Linef("%s := v.(%s)(%s)", goRets.Join(), fg.GoName, goArgs.Join())

	return b.String()
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

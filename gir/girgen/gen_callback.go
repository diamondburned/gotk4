package girgen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

// callbackPrefix is the prefix to prepend to a C callback that bridges Cgo.
const callbackPrefix = "gotk4_"

var callbackTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .Name }}
	type {{ .GoName }} func{{ .GoTail }}

	//export gotk4_{{ .GoName }}
	func gotk4_{{ .GoName }}{{ .CGoTail }} {{ .CBlock }}
`)

type callbackGenerator struct {
	gir.Callback
	GoName  string
	GoTail  string
	CGoTail string

	Closure *int
	Destroy *int

	Ng *NamespaceGenerator
}

func newCallbackGenerator(ng *NamespaceGenerator) callbackGenerator {
	return callbackGenerator{Ng: ng}
}

// Use sets the callback generator to the given GIR callback.
func (fg *callbackGenerator) Use(cb gir.Callback) bool {
	// We can't use the callback if it has no closure parameters.
	if cb.Parameters == nil || len(cb.Parameters.Parameters) == 0 {
		return false
	}
	// Don't generate destroy notifiers. It's an edge case that we handle
	// separately and mostly manually. There are also no good ways to detect
	// this.
	if strings.HasSuffix(cb.Name, "DestroyNotify") {
		return false
	}

	fg.Closure = nil
	for _, param := range cb.Parameters.Parameters {
		if param.Closure != nil {
			fg.Closure = param.Closure
			break
		}
	}
	if fg.Closure == nil {
		return false
	}

	fg.GoName = PascalToGo(cb.Name)
	fg.Callback = cb

	fg.GoTail = fg.Ng.FnCall(cb.CallableAttrs)
	if fg.GoTail == "" {
		return false
	}

	cgotail := pen.NewJoints(", ", len(cb.Parameters.Parameters))
	ctail := pen.NewJoints(", ", len(cb.Parameters.Parameters))

	for i, param := range cb.Parameters.Parameters {
		ctype := anyTypeC(param.AnyType)
		if ctype == "" {
			return false // probably var_args
		}

		ctail.Add(ctype + " _" + strconv.Itoa(i))
		cgotype := anyTypeCGo(param.AnyType)
		cgotail.Addf("arg%d %s", i, cgotype)
	}

	fg.CGoTail = "(" + cgotail.Join() + ")"
	cReturn := "void"

	if !returnIsVoid(cb.ReturnValue) {
		ctype := anyTypeC(cb.ReturnValue.AnyType)
		if ctype == "" {
			return false
		}

		cReturn = ctype
		fg.CGoTail += " " + anyTypeCGo(cb.ReturnValue.AnyType)
	}

	fg.Ng.cgo.Wordf("extern %s %s(%s);", cReturn, callbackPrefix+fg.GoName, ctail.Join())

	return true
}

func (fg *callbackGenerator) CBlock() string {
	b := pen.NewBlockSections(128, 1024, 4096, 128, 4096)

	fg.Ng.addImport("github.com/diamondburned/gotk4/internal/box")

	b.Linef(0, "v := box.Get(uintptr(arg%d))", *fg.Closure)
	b.Linef(0, "if v == nil {")
	b.Linef(0, "  panic(`callback not found`)")
	b.Linef(0, "}")
	b.EmptyLine(0)

	argAt := func(i int) string { return fmt.Sprintf("arg%d", i) }
	goArgs := pen.NewJoints(", ", len(fg.Parameters.Parameters))
	goRets := pen.NewJoints(", ", len(fg.Parameters.Parameters)+1)

	iterateParams(fg.CallableAttrs, func(i int, param gir.Parameter) bool {
		goName := SnakeToGo(false, param.Name)
		goType, _ := fg.Ng.ResolveAnyType(param.AnyType, true)

		b.Linef(1, "var %s %s", goName, goType)

		converter := fg.Ng.CGoConverter(TypeConversionToGo{
			TypeConversion: TypeConversion{
				Value:  argAt(i),
				Target: goName,
				Type:   param.AnyType,
				Owner:  param.TransferOwnership,
				ArgAt:  argAt,
			},
		})

		b.Line(2, converter)
		b.EmptyLine(2)

		goArgs.Add(goName)
		return true
	})

	iterateReturns(fg.CallableAttrs, func(goName string, i int, typ gir.AnyType) bool {
		goRets.Add(goName)
		return true
	})

	if goRets.Len() == 0 {
		b.Linef(3, "v.(%s)(%s)", fg.GoName, goArgs.Join())
		return b.String()
	}

	b.Linef(3, "%s := v.(%s)(%s)", goRets.Join(), fg.GoName, goArgs.Join())

	return b.String()
}

func (ng *NamespaceGenerator) generateCallbacks() {
	cg := newCallbackGenerator(ng)

	for _, callback := range ng.current.Namespace.Callbacks {
		if ng.mustIgnore(callback.Name, callback.CIdentifier) {
			continue
		}

		if !cg.Use(callback) {
			ng.logln(logInfo, "skipping callback", callback.Name)
			continue
		}

		ng.pen.BlockTmpl(callbackTmpl, &cg)
	}
}

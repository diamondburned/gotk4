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

	fg *FileGenerator
	ng *NamespaceGenerator
}

func newCallbackGenerator(ng *NamespaceGenerator) callbackGenerator {
	return callbackGenerator{ng: ng}
}

// Use sets the callback generator to the given GIR callback.
func (cg *callbackGenerator) Use(cb gir.Callback) bool {
	cg.fg = cg.ng.FileFromSource(cb.SourcePosition)

	// We can't use the callback if it has no closure parameters.
	if cb.Parameters == nil || len(cb.Parameters.Parameters) == 0 {
		cg.fg.Logln(LogSkip, "callback", cb.Name, "no closure parameter")
		return false
	}

	// Don't generate destroy notifiers. It's an edge case that we handle
	// separately and mostly manually. There are also no good ways to detect
	// this.
	if strings.HasSuffix(cb.Name, "DestroyNotify") {
		cg.fg.Logln(LogSkip, "callback", cb.Name, "is DestroyNotify")
		return false
	}

	cg.Closure = nil
	for _, param := range cb.Parameters.Parameters {
		if param.Closure != nil {
			cg.Closure = param.Closure
			break
		}
	}
	if cg.Closure == nil {
		cg.fg.Logln(LogSkip, "callback", cb.Name, "is DestroyNotify")
		return false
	}

	cg.GoName = PascalToGo(cb.Name)
	cg.Callback = cb

	cg.GoTail = cg.fg.fnCall(cb.CallableAttrs)
	if cg.GoTail == "" {
		return false
	}

	cgotail := pen.NewJoints(", ", len(cb.Parameters.Parameters))
	ctail := pen.NewJoints(", ", len(cb.Parameters.Parameters))

	for i, param := range cb.Parameters.Parameters {
		ctype := anyTypeC(param.AnyType)
		if ctype == "" {
			cg.fg.Logln(LogSkip, "callback", cb.Name, "anyTypeC parameter is empty")
			return false // probably var_args
		}

		ctail.Add(ctype + " _" + strconv.Itoa(i))
		cgotype := anyTypeCGo(param.AnyType)
		cgotail.Addf("arg%d %s", i, cgotype)
	}

	cg.CGoTail = "(" + cgotail.Join() + ")"
	cReturn := "void"

	if !returnIsVoid(cb.ReturnValue) {
		ctype := anyTypeC(cb.ReturnValue.AnyType)
		if ctype == "" {
			cg.fg.Logln(LogSkip, "callback", cb.Name, "anyTypeC return is empty")
			return false
		}

		cReturn = ctype
		cg.CGoTail += " " + anyTypeCGo(cb.ReturnValue.AnyType)
	}

	cg.fg.cgo.Wordf("extern %s %s(%s);", cReturn, callbackPrefix+cg.GoName, ctail.Join())

	return true
}

func (cg *callbackGenerator) CBlock() string {
	b := pen.NewBlockSections(128, 1024, 4096, 128, 4096)

	cg.fg.addImport("github.com/diamondburned/gotk4/internal/box")

	b.Linef(0, "v := box.Get(uintptr(arg%d))", *cg.Closure)
	b.Linef(0, "if v == nil {")
	b.Linef(0, "  panic(`callback not found`)")
	b.Linef(0, "}")
	b.EmptyLine(0)

	argAt := func(i int) string { return fmt.Sprintf("arg%d", i) }
	goArgs := pen.NewJoints(", ", len(cg.Parameters.Parameters))
	goRets := pen.NewJoints(", ", len(cg.Parameters.Parameters)+1)

	iterateParams(cg.CallableAttrs, func(i int, param gir.Parameter) bool {
		goName := SnakeToGo(false, param.Name)
		goType, _ := GoAnyType(cg.fg, param.AnyType, true)

		b.Linef(1, "var %s %s", goName, goType)

		converter := cg.fg.CGoConverter(TypeConversionToGo{
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

	iterateReturns(cg.CallableAttrs, func(goName string, i int, typ gir.AnyType) bool {
		goRets.Add(goName)
		return true
	})

	if goRets.Len() == 0 {
		b.Linef(3, "v.(%s)(%s)", cg.GoName, goArgs.Join())
		return b.String()
	}

	b.Linef(3, "%s := v.(%s)(%s)", goRets.Join(), cg.GoName, goArgs.Join())

	return b.String()
}

func (ng *NamespaceGenerator) generateCallbacks() {
	cg := newCallbackGenerator(ng)

	for _, callback := range ng.current.Namespace.Callbacks {
		if ng.mustIgnore(callback.Name, callback.CIdentifier) {
			continue
		}

		if !cg.Use(callback) {
			continue
		}

		cg.fg.pen.BlockTmpl(callbackTmpl, &cg)
	}
}

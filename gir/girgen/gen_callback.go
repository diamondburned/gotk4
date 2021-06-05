package girgen

import (
	"fmt"
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
	func gotk4_{{ .GoName }}{{ .CGoTail }} {{ .Block }}
`)

// CallbackCHeader renders the C function signature.
func CallbackCHeader(cb *gir.Callback) string {
	var ctail pen.Joints
	if cb.Parameters != nil {
		ctail = pen.NewJoints(", ", len(cb.Parameters.Parameters))

		for _, param := range cb.Parameters.Parameters {
			ctail.Add(anyTypeC(param.AnyType))
		}
	}

	cReturn := "void"
	if cb.ReturnValue != nil {
		cReturn = anyTypeC(cb.ReturnValue.AnyType)
	}

	goName := PascalToGo(cb.Name)
	return fmt.Sprintf("%s %s(%s);", cReturn, callbackPrefix+goName, ctail.Join())
}

type callbackGenerator struct {
	gir.Callback
	GoName  string
	GoTail  string
	CGoTail string
	Block   string

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

	cg.GoName = PascalToGo(cb.Name)
	cg.Callback = cb

	cg.Closure = findClosure(cb.Parameters.Parameters)
	if cg.Closure == nil {
		cg.fg.Logln(LogSkip, "callback", cb.Name, "is missing closure arg")
		return false
	}

	cg.CGoTail = cg.cgoTail()
	if cg.CGoTail == "" {
		return false
	}

	cg.GoTail = cg.fg.fnCall(cb.CallableAttrs)
	if cg.GoTail == "" {
		return false
	}

	cg.Block = cg.block()
	if cg.Block == "" {
		return false
	}

	return true
}

// findClosure returns the closure number or nil.
func findClosure(params []gir.Parameter) *int {
	for _, param := range params {
		if param.Closure != nil {
			return param.Closure
		}
	}
	return nil
}

func (cg *callbackGenerator) cgoTail() string {
	cgotail := pen.NewJoints(", ", len(cg.Parameters.Parameters))

	for i, param := range cg.Parameters.Parameters {
		ctype := anyTypeC(param.AnyType)
		if ctype == "" {
			cg.fg.Logln(LogSkip, "callback", cg.Name, "anyTypeC parameter is empty")
			return "" // probably var_args
		}

		cgotype := anyTypeCGo(param.AnyType)
		cgotail.Addf("arg%d %s", i, cgotype)
	}

	callTail := "(" + cgotail.Join() + ")"

	if !returnIsVoid(cg.ReturnValue) {
		ctype := anyTypeC(cg.ReturnValue.AnyType)
		if ctype == "" {
			cg.fg.Logln(LogSkip, "callback", cg.Name, "anyTypeC return is empty")
			return ""
		}

		callTail += " " + anyTypeCGo(cg.ReturnValue.AnyType)
	}

	return callTail
}

func (cg *callbackGenerator) block() string {
	const (
		secPrefix = iota
		secInputPre
		secInputConv
		secFnCall
		secOutputPre
		secOutputConv
		secReturn
	)

	// b := pen.NewBlockSections(256, 1024, 4096, 128, 1024, 4096, 128)

	b := pen.NewBlockSections() // TODO
	cap := len(cg.Parameters.Parameters) + 2

	b.Linef(secPrefix, "v := box.Get(uintptr(arg%d))", *cg.Closure)
	b.Linef(secPrefix, "if v == nil {")
	b.Linef(secPrefix, "  panic(`callback not found`)")
	b.Linef(secPrefix, "}")
	b.EmptyLine(secPrefix)

	b.Linef(secFnCall, "fn := v.(%s)", cg.GoName)

	inputAt := func(i int) string { return fmt.Sprintf("arg%d", i) }
	goArgs := pen.NewJoints(", ", cap)
	goRets := pen.NewJoints(", ", cap)

	inputValues := make([]CValueProp, 0, cap)
	outputValues := make([]GoValueProp, 0, cap)

	for i, param := range cg.Parameters.Parameters {
		if param.Direction != "out" {
			out := SnakeToGo(false, param.Name)
			goArgs.Add(out)

			inputValues = append(inputValues, CValueProp{
				ValueProp: NewValuePropParam(inputAt(i), out, &i, param.ParameterAttrs),
			})
		} else {
			in := SnakeToGo(false, param.Name)
			goRets.Add(in)

			// No need to have this declare a variable, since we're using the
			// walrus operator in the function call.
			outputValues = append(outputValues, GoValueProp{
				ValueProp: NewValuePropParam(in, inputAt(i), &i, param.ParameterAttrs),
			})
		}
	}

	// TODO: add GError support.

	var fnReturn string

	if !returnIsVoid(cg.ReturnValue) {
		in := "ret"
		goRets.Add(in)

		fnReturn = "cret"

		outputValues = append(outputValues, GoValueProp{
			ValueProp: NewValuePropReturn(in, fnReturn, *cg.ReturnValue),
		})
	}

	convI := cg.fg.CGoConverter(cg.Name, inputValues).WriteAll(
		nil, b.Section(secInputPre), b.Section(secInputConv),
	)

	if !convI {
		cg.fg.Logln(LogSkip, "callback (no C->Go conversion)", cFunctionSig(cg.CallableAttrs))
		return ""
	}

	convO := cg.fg.GoCConverter(cg.Name, outputValues).WriteAll(
		nil, nil, b.Section(secOutputConv),
	)

	if !convO {
		cg.fg.Logln(LogSkip, "callback (no Go->C conversion)", cFunctionSig(cg.CallableAttrs))
		return ""
	}

	cg.fg.addImport(importInternal("box"))

	if goRets.Len() == 0 {
		b.Linef(secFnCall, "fn(%s)", goArgs.Join())
	} else {
		b.Linef(secFnCall, "%s := fn(%s)", goRets.Join(), goArgs.Join())
	}

	if fnReturn != "" {
		b.Linef(secReturn, "return "+fnReturn)
	}

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

		cg.fg.pen.WriteTmpl(callbackTmpl, &cg)
	}
}

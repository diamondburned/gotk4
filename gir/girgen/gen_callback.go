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

	for i, param := range cb.Parameters.Parameters {
		ctype := anyTypeC(param.AnyType)
		if ctype == "" {
			cg.fg.Logln(LogSkip, "callback", cb.Name, "anyTypeC parameter is empty")
			return false // probably var_args
		}

		cgotype := anyTypeCGo(param.AnyType)
		cgotail.Addf("arg%d %s", i, cgotype)
	}

	cg.CGoTail = "(" + cgotail.Join() + ")"

	if !returnIsVoid(cb.ReturnValue) {
		ctype := anyTypeC(cb.ReturnValue.AnyType)
		if ctype == "" {
			cg.fg.Logln(LogSkip, "callback", cb.Name, "anyTypeC return is empty")
			return false
		}

		cg.CGoTail += " " + anyTypeCGo(cb.ReturnValue.AnyType)
	}

	return true
}

func (cg *callbackGenerator) Block() string {
	b := pen.NewBlockSections(256, 4096, 128, 4096, 128)
	cap := len(cg.Parameters.Parameters) + 2

	b.Linef(0, "v := box.Get(uintptr(arg%d))", *cg.Closure)
	b.Linef(0, "if v == nil {")
	b.Linef(0, "  panic(`callback not found`)")
	b.Linef(0, "}")
	b.EmptyLine(0)

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
				ValueProp: NewValuePropParam(inputAt(i), out, param.ParameterAttrs),
			})
		} else {
			in := SnakeToGo(false, param.Name)
			goRets.Add(in)

			// Set the output to *v, which is already declared in the function
			// parameter.
			out := "*" + inputAt(i)

			// No need to have this declare a variable, since we're using the
			// walrus operator in the function call.
			outputValues = append(outputValues, GoValueProp{
				ValueProp: NewValuePropParam(in, out, param.ParameterAttrs),
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

	var (
		paramsConv = cg.fg.CGoConverter(TypeConversionToGo{
			Values: inputValues,
			Parent: cg.Name,
		})
		// Ignore preamble since we didn't ask to create any decls.
		returnsConv = cg.fg.GoCConverter(TypeConversionToC{
			Values: outputValues,
			Parent: cg.Name,
		})
	)

	if paramsConv == nil || returnsConv == nil {
		return ""
	}

	cg.fg.addImport("github.com/diamondburned/gotk4/internal/box")
	paramsConv.Apply(cg.fg)
	returnsConv.Apply(cg.fg)

	b.Linef(1, paramsConv.Conversion)
	b.Linef(3, returnsConv.Conversion)

	b.Linef(2, "fn := v.(%s)", cg.GoName)
	if goRets.Len() == 0 {
		b.Linef(2, "fn(%s)", goArgs.Join())
	} else {
		b.Linef(2, "%s := fn(%s)", goRets.Join(), goArgs.Join())
	}

	if fnReturn != "" {
		b.Linef(4, "return "+fnReturn)
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

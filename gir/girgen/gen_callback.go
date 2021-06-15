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

	pen *pen.BlockSections
	fg  *FileGenerator
	ng  *NamespaceGenerator
}

func newCallbackGenerator(ng *NamespaceGenerator) callbackGenerator {
	return callbackGenerator{
		ng:  ng,
		pen: pen.NewBlockSections(256, 1024, 4096, 128, 1024, 4096, 128),
	}
}

// Use sets the callback generator to the given GIR callback.
func (cg *callbackGenerator) Use(cb gir.Callback) bool {
	cg.fg = cg.ng.FileFromSource(cb.DocElements)

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

	if !cg.renderBlock() {
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

func callbackArg(i int) string {
	return "arg" + strconv.Itoa(i)
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
		cgotail.Addf("%s %s", callbackArg(i), cgotype)
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

func (cg *callbackGenerator) renderBlock() bool {
	defer cg.pen.Reset()

	const (
		secPrefix = iota
		secInputPre
		secInputConv
		secFnCall
		secOutputPre
		secOutputConv
		secReturn
	)

	cg.pen.Linef(secPrefix, "v := box.Get(uintptr(%s))", callbackArg(*cg.Closure))
	cg.pen.Linef(secPrefix, "if v == nil {")
	cg.pen.Linef(secPrefix, "  panic(`callback not found`)")
	cg.pen.Linef(secPrefix, "}")
	cg.pen.EmptyLine(secPrefix)

	cg.pen.Linef(secFnCall, "fn := v.(%s)", cg.GoName)

	callbackValues := make([]ConversionValue, 0, len(cg.Parameters.Parameters)+2)

	for i, param := range cg.Parameters.Parameters {
		if param.Skip {
			continue
		}

		if param.Direction == "" {
			param.Direction = "in" // default
		}

		var in string
		var out string
		var dir ConversionDirection

		switch param.Direction {
		case "in":
			in = callbackArg(i)
			out = SnakeToGo(false, param.Name)
			dir = ConvertCToGo
		case "out":
			in = SnakeToGo(false, param.Name)
			out = callbackArg(i)
			dir = ConvertGoToC
		default:
			// TODO: inout
			return false
		}

		value := NewConversionValue(in, out, i, dir, param.ParameterAttrs)
		callbackValues = append(callbackValues, value)
	}

	var hasReturn bool
	if !returnIsVoid(cg.ReturnValue) {
		hasReturn = true
		returnName := returnName(cg.CallableAttrs)

		value := NewConversionValueReturn(returnName, "cret", ConvertGoToC, *cg.ReturnValue)
		callbackValues = append(callbackValues, value)
	}

	convert := NewTypeConverter(cg.fg, cg.Name, callbackValues)
	results := convert.ConvertAll()
	if results == nil {
		cg.fg.Logln(LogSkip, "callback has no conversion", cFunctionSig(cg.CallableAttrs))
		return false
	}

	cg.fg.applyConvertedFxs(results)

	goCallArgs := pen.NewJoints(", ", len(results))
	goCallRets := pen.NewJoints(", ", len(results))

	goTypeArgs := pen.NewJoints(", ", len(results))
	goTypeRets := pen.NewJoints(", ", len(results))

	for _, result := range results {
		if result.Skip {
			continue
		}

		switch result.Direction {
		case ConvertCToGo:
			goCallArgs.Add(result.OutCall)
			goTypeArgs.Addf("%s %s", result.OutName, result.OutType)

			cg.pen.Line(secInputPre, result.OutDeclare)
			cg.pen.Line(secInputConv, result.Conversion)

		case ConvertGoToC:
			goCallRets.Add(result.InCall)
			goTypeArgs.Addf("%s %s", result.InName, result.InType)

			cg.pen.Line(secOutputPre, result.OutDeclare)
			cg.pen.Line(secOutputConv, result.Conversion)
		}
	}

	if goCallRets.Len() == 0 {
		cg.pen.Linef(secFnCall, "fn(%s)", goCallArgs.Join())
	} else {
		cg.pen.Linef(secFnCall, "%s := fn(%s)", goCallRets.Join(), goCallArgs.Join())
	}

	if hasReturn {
		cg.pen.Linef(secReturn, "return cret")
	}

	cg.Block = cg.pen.String()

	cg.GoTail = "(" + goTypeArgs.Join() + ")"
	if goTypeRets.Len() > 0 {
		cg.GoTail += " (" + goTypeRets.Join() + ")"
	}

	// Only add the import now, since we know that the callback will be
	// generated.
	cg.fg.addImportInternal("box")

	return true
}

func (ng *NamespaceGenerator) generateCallbacks() {
	cg := newCallbackGenerator(ng)

	for _, callback := range ng.current.Namespace.Callbacks {
		if !callback.IsIntrospectable() {
			continue
		}
		if ng.mustIgnore(&callback.Name, &callback.CIdentifier) {
			continue
		}
		if !cg.Use(callback) {
			continue
		}

		cg.fg.pen.WriteTmpl(callbackTmpl, &cg)
	}
}

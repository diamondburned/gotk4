package callback

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

// CallbackPreamble writes the preamble for the Go callback function. It uses
// the Closure parameter to find the Go function on the heap and casts it to the
// right type. The returned name is always "fn".
func CallbackPreamble(g *Generator, sect *pen.BlockSection) (string, bool) {
	if g.Closure == nil {
		g.Logln(logger.Debug, "skipping since no closure argument")
		return "", false
	}

	g.header.ImportCore("gbox")

	sect.Linef("var fn %s", g.GoName)
	sect.Linef("{")
	sect.Linef("  v := gbox.Get(uintptr(%s))", CallbackArg(*g.Closure))
	sect.Linef("  if v == nil {")
	sect.Linef("    panic(`callback not found`)")
	sect.Linef("  }")
	sect.Linef("  fn = v.(%s)", g.GoName)
	sect.Linef("}")

	return "fn", true
}

// Generator generates a callback in Go.
type Generator struct {
	*gir.CallableAttrs
	Parent interface{}

	GoName  string
	GoTail  string
	CGoName string
	CGoTail string
	Block   string

	Preamble func(g *Generator, sect *pen.BlockSection) (call string, ok bool)
	Closure  *int
	Destroy  *int // TODO: why is this unused?

	pen    *pen.BlockSections
	gen    types.FileGenerator
	header file.Header
}

// NewGenerator creates a new Generator instance.
func NewGenerator(gen types.FileGenerator) Generator {
	return Generator{
		pen:      pen.NewBlockSections(256, 1024, 4096, 128, 1024, 4096, 128),
		gen:      gen,
		Preamble: CallbackPreamble,
	}
}

func (g *Generator) Logln(lvl logger.Level, v ...interface{}) {
	g.gen.Logln(lvl, logger.Prefix(v, fmt.Sprintf("callback %s (C.%s):", g.Name, g.CIdentifier))...)
}

// Reset resets the callback generator.
func (g *Generator) Reset() {
	g.pen.Reset()

	*g = Generator{
		Parent:   g.Parent,
		Preamble: g.Preamble,
		pen:      g.pen,
		gen:      g.gen,
	}
}

// Header returns the callback generator's current header.
func (g *Generator) Header() *file.Header {
	return &g.header
}

// Use sets the callback generator to the given GIR callback.
func (g *Generator) Use(cb *gir.CallableAttrs) bool {
	g.Reset()
	g.CallableAttrs = cb

	// This check doesn't apply for virtual methods.

	// if cb.Parameters == nil || len(cb.Parameters.Parameters) == 0 {
	// 	g.Logln(logger.Debug, "has no parameters at all")
	// 	return false
	// }

	if !cb.IsIntrospectable() || types.Filter(g.gen, cb.Name, cb.CIdentifier) {
		return false
	}

	// Don't generate destroy notifiers. It's an edge case that we handle
	// separately and mostly manually. There are also no good ways to detect
	// this.
	if strings.HasSuffix(cb.Name, "DestroyNotify") {
		g.Logln(logger.Debug, "skipping DestroyNotify-ish callback")
		return false
	}

	g.GoName = strcases.PascalToGo(cb.Name)
	g.CGoName = file.CallableExportedName(g.gen.Namespace(), cb)

	if cb.Parameters != nil && len(cb.Parameters.Parameters) > 0 {
		g.Closure = findClosure(cb.Parameters.Parameters)
	}

	if g.CGoTail = g.cgoTail(); g.CGoTail == "" {
		return false
	}

	if !g.renderBlock() {
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

// CallbackArg formats the callback variable name for the given variable index.
func CallbackArg(i int) string {
	return "arg" + strconv.Itoa(i+1)
}

// CGoTail generates the CGo function tail for the given callable.
func CGoTail(gen types.FileGenerator, g *gir.CallableAttrs) (string, error) {
	callTail := "()"

	if g.Parameters != nil {
		cgotail := pen.NewJoints(", ", len(g.Parameters.Parameters)+1)

		var addParam func(ix int, param *gir.ParameterAttrs) bool

		switch gen.LinkMode() {
		case types.DynamicLinkMode:
			addParam = func(ix int, param *gir.ParameterAttrs) bool {
				anyType := types.ResolveAnyType(gen, param.AnyType)

				ctype := types.AnyTypeC(anyType)
				if ctype == "" {
					// probably var_args
					return false
				}

				cgoType := types.AnyTypeCGo(anyType)
				cgotail.Addf("%s %s", CallbackArg(ix), cgoType)
				return true
			}
		case types.RuntimeLinkMode:
			addParam = func(ix int, param *gir.ParameterAttrs) bool {
				anyType := types.ResolveAnyType(gen, param.AnyType)

				// Keep in sync with types.CallbackCHeader.
				cType := types.AnyTypeCPrimitive(gen, anyType)
				if cType == "" {
					return false
				}

				cgoType := types.CGoTypeFromC(cType)
				cgotail.Addf("%s %s", CallbackArg(ix), cgoType)
				return true
			}
		}

		if g.Parameters.InstanceParameter != nil {
			if !addParam(-1, &g.Parameters.InstanceParameter.ParameterAttrs) {
				return "", errors.New("instance parameter is empty")
			}
		}

		for i, param := range g.Parameters.Parameters {
			if !addParam(i, &param.ParameterAttrs) {
				return "", fmt.Errorf("parameter %q is empty", param.Name)
			}
		}

		if g.Throws {
			cgotail.Add("_cerr **C.GError")
		}

		callTail = "(" + cgotail.Join() + ")"
	}

	if !types.ReturnIsVoid(g.ReturnValue) {
		anyTypeCGo := types.ResolveAnyTypeCGo(gen, g.ReturnValue.AnyType)
		if anyTypeCGo == "" {
			return "", fmt.Errorf("anyTypeCGo return is empty")
		}

		callTail += fmt.Sprintf(" (cret %s)", anyTypeCGo)
	}

	return callTail, nil
}

func (g *Generator) cgoTail() string {
	v, err := CGoTail(g.gen, g.CallableAttrs)
	if err != nil {
		g.Logln(logger.Debug, err.Error())
	}
	return v
}

func (g *Generator) renderBlock() bool {
	defer g.pen.Reset()

	const (
		secPrefix = iota
		secInputPre
		secInputConv
		secFnCall
		secOutputPre
		secOutputConv
		secReturn
	)

	var callbackValues []typeconv.ConversionValue

	if g.Parameters != nil && len(g.Parameters.Parameters) > 0 {
		callbackValues = make([]typeconv.ConversionValue, 0, len(g.Parameters.Parameters)+2)

		for i, param := range g.Parameters.Parameters {
			// Skip generating the closure parameter.
			if param.Skip || (g.Closure != nil && i == *g.Closure) {
				continue
			}

			// Reguess the parameter.
			param.Direction = types.GuessParameterOutput(&param)

			var in string
			var out string
			var dir typeconv.ConversionDirection

			switch param.Direction {
			case "in":
				in = CallbackArg(i)
				out = "_" + strcases.SnakeToGo(false, param.Name)
				dir = typeconv.ConvertCToGo
			case "out":
				in = strcases.SnakeToGo(false, param.Name)
				out = CallbackArg(i)
				dir = typeconv.ConvertGoToC
			default:
				// TODO: inout
				return false
			}

			value := typeconv.NewValue(in, out, i, dir, param)
			callbackValues = append(callbackValues, value)
		}
	}

	var hasReturn bool
	if !types.ReturnIsVoid(g.ReturnValue) {
		hasReturn = true
		returnName := callable.ReturnName(g.CallableAttrs)

		// Ignore ok returns, since the Go API will only have an error return.
		if !g.Throws || returnName != "ok" {
			value := typeconv.NewReturnValue(returnName, "cret", typeconv.ConvertGoToC, *g.ReturnValue)
			callbackValues = append(callbackValues, value)
		}
	}

	if g.Throws {
		err := typeconv.NewThrowValue("_goerr", "_cerr")
		err.Direction = typeconv.ConvertGoToC
		callbackValues = append(callbackValues, err)
	}

	var typ *gir.TypeFindResult
	if g.Parent == nil {
		typ = types.Find(g.gen, g.CallableAttrs.Name)
	} else {
		typ = &gir.TypeFindResult{
			NamespaceFindResult: g.gen.Namespace(),
			Type:                g.Parent,
		}
	}

	convert := typeconv.NewConverter(g.gen, typ, callbackValues)
	if convert == nil {
		return false
	}

	convert.Callback = true
	convert.UseLogger(g)

	results := convert.ConvertAll()
	if results == nil {
		g.Logln(logger.Debug, "has no conversion")
		return false
	}

	file.ApplyHeader(g, convert)

	// Do a bit of trickery: if we have a GCancellable in the function, then it
	// should be the first parameter. The GCancellable will then be resolved to
	// a context.Context during conversion.
	callable.MoveContextResult(g.gen, results)

	goCallArgs := pen.NewJoints(", ", len(results))
	goCallRets := pen.NewJoints(", ", len(results))

	goTypeArgs := pen.NewJoints(", ", len(results))
	goTypeRets := pen.NewJoints(", ", len(results))

	for _, result := range results {
		if result.Skip {
			continue
		}

		switch result.Direction {
		case typeconv.ConvertCToGo:
			// Undo the underscore.
			decoOut := strings.TrimPrefix(result.OutName, "_")

			goCallArgs.Add(result.Out.Call)
			goTypeArgs.Addf("%s %s", decoOut, result.Out.Type)

			g.pen.Line(secInputPre, result.Out.Declare)
			g.pen.Line(secInputConv, result.Conversion)

		case typeconv.ConvertGoToC:
			goCallRets.Add(result.In.Call)
			goTypeRets.Addf("%s %s", result.InName, result.In.Type)

			// Out.Declare is declared in the function signature.
			// g.pen.Line(secOutputPre, result.Out.Declare)

			g.pen.Line(secOutputConv, result.Conversion)
		}
	}

	g.GoTail = "(" + goTypeArgs.Join() + ")"
	if goTypeRets.Len() > 0 {
		g.GoTail += " (" + goTypeRets.Join() + ")"
	}

	g.GoTail = callable.CoalesceTail(g.GoTail)

	// Call this after the type conversion so Preamble can use the above
	// outputs.
	fn, ok := g.Preamble(g, g.pen.Section(secPrefix))
	if !ok {
		return false
	}
	g.pen.EmptyLine(secPrefix)

	if goCallRets.Len() == 0 {
		g.pen.Linef(secFnCall, "%s(%s)", fn, goCallArgs.Join())
	} else {
		g.pen.Linef(secFnCall, "%s := %s(%s)", goCallRets.Join(), fn, goCallArgs.Join())
	}

	if hasReturn {
		g.pen.Linef(secReturn, "return cret")
	}

	g.Block = g.pen.String()

	return true
}

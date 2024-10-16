// Package callable provides a generic callable generator.
package callable

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

// FileGenerator describes the interface of a file generator.
type FileGenerator interface {
	types.FileGenerator
}

// CallablePreamble returns CIdentifier or fails.
func CallablePreamble(g *Generator, sect *pen.BlockSection) (string, bool) {
	if g.CIdentifier == "" {
		return "", false
	}
	return "C." + g.CIdentifier, true
}

// Generator is a generator instance that generates a GIR callable.
type Generator struct {
	*gir.CallableAttrs
	Name  string
	Tail  string
	Block string

	constructor bool
	Converts    []string

	Preamble  func(g *Generator, sect *pen.BlockSection) (call string, ok bool)
	ExtraArgs [2][]string // {front, back}
	Conv      *typeconv.Converter
	Results   []typeconv.ValueConverted
	GoArgs    pen.Joints
	GoRets    pen.Joints

	ParamDocs  []cmt.ParamDoc
	ReturnDocs []cmt.ParamDoc

	typ *gir.TypeFindResult
	pen *pen.BlockSections
	hdr file.Header
	gen FileGenerator
}

// IgnoredNames is a list of method names that would be ignored. For more
// information, see typeconv/c-go.go.
var IgnoredNames = []string{
	"ref",
	"unref",
	"free",
}

// NewGenerator creates a new callable generator from the given generator.
func NewGenerator(gen FileGenerator) Generator {
	// Arbitrary sizes, whatever.
	pen := pen.NewBlockSections(1024, 4096, 256, 1024, 4096, 128)

	return Generator{
		Preamble: CallablePreamble,
		pen:      pen,
		gen:      gen,
	}
}

// Reset resets the state of the generator while reusing the backing pen.
func (g *Generator) Reset() {
	g.hdr.Reset()
	g.pen.Reset()
	g.GoArgs.Reset(", ")
	g.GoRets.Reset(", ")

	*g = Generator{
		pen:      g.pen,
		gen:      g.gen,
		hdr:      g.hdr,
		Preamble: g.Preamble,
		GoArgs:   g.GoArgs,
		GoRets:   g.GoRets,

		constructor: g.constructor,
	}
}

var _ file.Headerer = (*Generator)(nil)

// Header returns the generator's current headers.
func (g *Generator) Header() *file.Header {
	return &g.hdr
}

// FileGenerator returns the generator's internal file generator.
func (g *Generator) FileGenerator() FileGenerator {
	return g.gen
}

// UseConstructor calls Use with the constructor flag.
func (g *Generator) UseConstructor(typ *gir.TypeFindResult, call *gir.CallableAttrs) bool {
	g.constructor = true
	ok := g.Use(typ, call)
	g.constructor = false
	return ok
}

// Use uses the given CallableAttrs for the generator.
func (g *Generator) Use(typ *gir.TypeFindResult, call *gir.CallableAttrs) bool {
	g.Reset()
	g.CallableAttrs = call
	g.typ = typ

	if call.ShadowedBy != "" || call.MovedTo != "" {
		// Skip this one. Hope the caller reaches the Shadows method,
		// eventually.
		return false
	}
	if !call.IsIntrospectable() {
		return false
	}
	// This is broken; unsure why.
	if types.FilterSub(g.gen, typ.NamespacedType(), call.Name, call.CIdentifier) {
		return false
	}
	// Double-check that the C identifier is allowed.
	if call.CIdentifier != "" && types.FilterCType(g.gen, call.CIdentifier) {
		return false
	}

	for _, name := range IgnoredNames {
		if name == call.Name || strings.HasSuffix(call.Name, "_"+name) {
			g.Logln(logger.Debug, "not generating ignored name", call.Name)
			return false
		}
	}

	g.Name = strcases.SnakeToGo(true, call.Name)
	if call.Shadows != "" {
		g.Name = strcases.SnakeToGo(true, call.Shadows)
	}

	if !g.renderBlock() {
		return false
	}

	g.ParamDocs = g.ParamDocs[:0]
	g.EachParamResult(gatherParamDoc(&g.ParamDocs))

	g.ReturnDocs = g.ReturnDocs[:0]
	g.EachReturnResult(gatherParamDoc(&g.ReturnDocs))

	return true
}

func gatherParamDoc(docs *[]cmt.ParamDoc) func(*typeconv.ValueConverted) {
	return func(value *typeconv.ValueConverted) {
		var name string // GoName

		switch value.Direction {
		case typeconv.ConvertCToGo:
			name = value.OutName
		case typeconv.ConvertGoToC:
			name = value.InName
		}

		// Returns have trailing underscores.
		name = strings.TrimPrefix(name, "_")

		*docs = append(*docs, cmt.ParamDoc{
			Name:     name,
			Optional: value.Optional || value.Nullable,
			InfoElements: gir.InfoElements{
				DocElements: gir.DocElements{Doc: value.Doc},
			},
		})
	}
}

// Recv returns the receiver variable name. This method should only be called
// for methods.
func (g *Generator) Recv() string {
	if g.Parameters != nil && g.Parameters.InstanceParameter != nil {
		return strcases.SnakeToGo(false, g.Parameters.InstanceParameter.Name)
	}

	return "v"
}

func (g *Generator) renderBlock() bool {
	switch g.gen.LinkMode() {
	case types.DynamicLinkMode:
		return g.renderDynamicLinkedBlock()
	case types.RuntimeLinkMode:
		return g.renderRuntimeLinkedBlock()
	default:
		panic("unknown LinkMode")
	}
}

func (g *Generator) renderDynamicLinkedBlock() bool {
	const (
		secPreamble = iota
		secInputDecl
		secInputConv
		secFnCall
		secOutputDecl
		secOutputConv
		secReturn
	)

	var (
		instanceParam  bool
		callableValues []typeconv.ConversionValue
	)

	if g.Parameters != nil {
		callableValues = make([]typeconv.ConversionValue, 0, len(g.Parameters.Parameters)+2)

		if g.Parameters.InstanceParameter != nil {
			instanceParam = true

			callableValues = append(callableValues, typeconv.NewReceiverValue(
				strcases.SnakeToGo(false, g.Parameters.InstanceParameter.Name),
				"_arg0",
				typeconv.ConvertGoToC,
				g.Parameters.InstanceParameter,
			))
		}

		contextParamIx := findGCancellableParam(g.gen, g.Parameters.Parameters)

		// Copy the parameters list so we can freely mutate it.
		parameters := types.ResolveParameters(g.gen, g.Parameters.Parameters)

		// Preprocess the values to normalize an edge case; see comment below.
		for i, value := range parameters {
			if value.AnyType.Array == nil || value.AnyType.Array.Length == nil {
				continue
			}

			length := &g.Parameters.Parameters[*value.Array.Length]

			// Special case: check if the length is both an output parameter AND
			// has a type that isn't a pointer. In such a case, the function
			// expects the user to supply a buffer of constant length, so we
			// should treat them as input parameters. We can only do this if the
			// type is directly castable, however. Changing this will also allow
			// the function to pass the above output-pointer check.
			if value.ParameterAttrs.Direction == "out" && value.CallerAllocates {
				if !types.AnyTypeIsPtr(length.AnyType) {
					length.Direction = "in"
					parameters[i].Direction = "in"
				}
			}
		}

		for i, param := range parameters {
			var in string
			var out string
			var dir typeconv.ConversionDirection

			switch types.GuessParameterOutput(&param) {
			case "in":
				in = strcases.SnakeToGo(false, param.Name)
				if contextParamIx == i {
					// Idiomatic context naming.
					in = "ctx"
				}
				out = fmt.Sprintf("_arg%d", i+1)
				dir = typeconv.ConvertGoToC
			case "out":
				in = fmt.Sprintf("_arg%d", i+1)
				out = "_" + strcases.SnakeToGo(false, param.Name)
				dir = typeconv.ConvertCToGo
			default:
				return false
			}

			value := typeconv.NewValue(in, out, i, dir, param)
			callableValues = append(callableValues, value)
		}
	}

	var hasReturn bool
	if !types.ReturnIsVoid(g.ReturnValue) {
		returnName := ReturnName(g.CallableAttrs)

		// If the last return is a bool and the function can throw an error,
		// then the boolean is probably to indicate that things are OK. We can
		// skip generating this boolean.
		if !g.Throws || returnName != "ok" {
			hasReturn = true
			returnName = "_" + returnName

			value := typeconv.NewReturnValue(
				"_cret", returnName, typeconv.ConvertCToGo, *g.ReturnValue,
			)
			// Constructors are bodged, so the returned type is concretely
			// accuratee.
			value.KeepType = g.constructor

			callableValues = append(callableValues, value)
		}
	}

	if g.Throws {
		callableValues = append(callableValues, typeconv.NewThrowValue("_cerr", "_goerr"))
	}

	g.Conv = typeconv.NewConverter(g.gen, g.typ, callableValues)
	g.Conv.UseLogger(g)

	if g.Conv == nil {
		g.Logln(logger.Debug, "converter failed", cFunctionHeader(g.CallableAttrs))
		return false
	}

	g.Results = g.Conv.ConvertAll()
	if g.Results == nil {
		g.Logln(logger.Debug, "no conversion", cFunctionHeader(g.CallableAttrs))
		return false
	}

	// Apply imports and such.
	file.ApplyHeader(g, g.Conv)

	// Do a bit of trickery: if we have a GCancellable in the function, then it
	// should be the first parameter. The GCancellable will then be resolved to
	// a context.Context during conversion.
	resultParams := g.Results
	if instanceParam {
		// Don't count the instance parameter.
		resultParams = resultParams[1:]
	}
	MoveContextResult(g.gen, resultParams)

	// For Go variables after the return statement.
	goReturns := pen.NewJoints(", ", 2)

	for i, converted := range g.Results {
		switch converted.Direction {
		case typeconv.ConvertGoToC: // parameter
			// Skip the instance parameter if any.
			if i != 0 || !instanceParam {
				g.GoArgs.Addf("%s %s", converted.InName, converted.In.Type)
			}

			// Go inputs are declared in the parameters, so no In.Declare.
			// C outputs have to be declared (input means C function input).
			g.pen.Line(secInputDecl, converted.Out.Declare)
			// Conversions follow right after declaring all outputs.
			g.pen.Line(secInputConv, converted.Conversion)

		case typeconv.ConvertCToGo: // return
			// decoOut is the name that's used solely for documentation
			// purposes. It is not used internally at all, and so it doesn't
			// have the underscore.
			decoOut := strings.TrimPrefix(converted.OutName, "_")
			g.GoRets.Addf("%s %s", decoOut, converted.Out.Type)

			goReturns.Add(converted.OutName)

			g.pen.Line(secInputDecl, converted.In.Declare)
			// Go outputs should be redeclared.
			g.pen.Line(secOutputDecl, converted.Out.Declare)
			// Conversions follow right after declaring all outputs.
			g.pen.Line(secOutputConv, converted.Conversion)
		}
	}

	fn, ok := g.Preamble(g, g.pen.Section(secPreamble))
	if !ok {
		return false
	}

	// For C function calling.
	ccallParams := g.Conv.CCallParams()
	if g.ExtraArgs[0] != nil {
		ccallParams = append(g.ExtraArgs[0], ccallParams...)
	}
	if g.ExtraArgs[1] != nil {
		ccallParams = append(ccallParams, g.ExtraArgs[1]...)
	}

	callParams := strings.Join(ccallParams, ", ")

	if !hasReturn {
		g.pen.Linef(secFnCall, "%s(%s)", fn, callParams)
	} else {
		g.pen.Linef(secFnCall, "_cret = %s(%s)", fn, callParams)
	}

	// Generate the right statements to ensure that nothing is freed before or
	// while we're invoking our function.
	for _, converted := range g.Results {
		if converted.Direction == typeconv.ConvertGoToC {
			g.hdr.Import("runtime")
			g.pen.Linef(secFnCall, "runtime.KeepAlive(%s)", converted.InName)
		}
	}

	if goReturns.Len() > 0 {
		g.pen.EmptyLine(secFnCall)
		g.pen.Line(secReturn, "return "+goReturns.Join())
	}

	g.Block = g.pen.String()
	g.Tail = "(" + g.GoArgs.Join() + ") " + formatReturnSig(g.GoRets)

	g.pen.Reset()
	return true
}

func (g *Generator) renderRuntimeLinkedBlock() bool {
	const (
		secInputDecl = iota
		secInputConv
		secFnCall
		secKeepAlive
		secOutputDecl
		secOutputConv
		secReturn
	)

	var (
		instanceParam   bool
		callableValues  []typeconv.ConversionValue
		argumentIndices map[int]int

		nIn  int
		nOut int
	)

	g.hdr.ImportCore("girepository")

	argsName := "nil"
	outsName := "nil"

	if g.Parameters != nil {
		contextParamIx := findGCancellableParam(g.gen, g.Parameters.Parameters)

		// Copy the parameters list so we can freely mutate it.
		parameters := types.ResolveParameters(g.gen, g.Parameters.Parameters)

		// Preprocess the values to normalize an edge case; see comment below.
		for i, value := range parameters {
			if value.AnyType.Array == nil || value.AnyType.Array.Length == nil {
				continue
			}

			length := &g.Parameters.Parameters[*value.Array.Length]

			// Special case: check if the length is both an output parameter AND
			// has a type that isn't a pointer. In such a case, the function
			// expects the user to supply a buffer of constant length, so we
			// should treat them as input parameters. We can only do this if the
			// type is directly castable, however. Changing this will also allow
			// the function to pass the above output-pointer check.
			if value.ParameterAttrs.Direction == "out" && value.CallerAllocates {
				if !types.AnyTypeIsPtr(length.AnyType) {
					length.Direction = "in"
					parameters[i].Direction = "in"
				}
			}
		}

		callableValues = make([]typeconv.ConversionValue, 0, len(g.Parameters.Parameters)+2)
		argumentIndices = make(map[int]int, cap(callableValues))

		if g.Parameters.InstanceParameter != nil {
			instanceParam = true
			argumentIndices[len(callableValues)] = nIn // 0

			callableValues = append(callableValues, typeconv.NewReceiverValue(
				strcases.SnakeToGo(false, g.Parameters.InstanceParameter.Name),
				"_args[0]",
				typeconv.ConvertGoToC,
				g.Parameters.InstanceParameter,
			))
			nIn++
		}

		for i, param := range parameters {
			var in string
			var out string
			var dir typeconv.ConversionDirection

			switch types.GuessParameterOutput(&param) {
			case "in":
				in = strcases.SnakeToGo(false, param.Name)
				if contextParamIx == i {
					// Idiomatic context naming.
					in = "ctx"
				}
				out = fmt.Sprintf("_args[%d]", nIn)
				dir = typeconv.ConvertGoToC

				argumentIndices[len(callableValues)] = nIn
				nIn++

			case "out":
				in = fmt.Sprintf("_outs[%d]", nOut)
				out = "_" + strcases.SnakeToGo(false, param.Name)
				dir = typeconv.ConvertCToGo

				argumentIndices[len(callableValues)] = nOut
				nOut++

			default:
				return false
			}

			value := typeconv.NewValue(in, out, i, dir, param)
			callableValues = append(callableValues, value)
		}

		if nIn > 0 {
			g.pen.Linef(secInputDecl, "var _args [%d]girepository.Argument", nIn)
			argsName = "_args[:]"
		}
		if nOut > 0 {
			g.pen.Linef(secInputDecl, "var _outs [%d]girepository.Argument", nOut)
			outsName = "_outs[:]"
		}
	}

	var hasReturn bool
	if !types.ReturnIsVoid(g.ReturnValue) {
		returnName := ReturnName(g.CallableAttrs)

		// If the last return is a bool and the function can throw an error,
		// then the boolean is probably to indicate that things are OK. We can
		// skip generating this boolean.
		if !g.Throws || returnName != "ok" {
			hasReturn = true
			returnName = "_" + returnName

			value := typeconv.NewReturnValue(
				"_cret", returnName, typeconv.ConvertCToGo, *g.ReturnValue,
			)
			// Constructors are bodged, so the returned type is concretely
			// accurate.
			value.KeepType = g.constructor

			callableValues = append(callableValues, value)
		}
	}

	if g.Throws {
		callableValues = append(callableValues, typeconv.NewThrowValue("_cerr", "_goerr"))
	}

	g.Conv = typeconv.NewConverter(g.gen, g.typ, callableValues)
	if g.Conv == nil {
		g.Logln(logger.Debug, "converter failed", cFunctionHeader(g.CallableAttrs))
		return false
	}

	g.Conv.UseLogger(g)
	// Force unsafe.Pointer casting from girepository.Argument.
	g.Conv.MustCast = true

	g.Results = g.Conv.ConvertAll()
	if g.Results == nil {
		g.Logln(logger.Debug, "no conversion", cFunctionHeader(g.CallableAttrs))
		return false
	}

	// Apply imports and such.
	file.ApplyHeader(g, g.Conv)

	// Do a bit of trickery: if we have a GCancellable in the function, then it
	// should be the first parameter. The GCancellable will then be resolved to
	// a context.Context during conversion.
	resultParams := g.Results
	if instanceParam {
		// Don't count the instance parameter.
		resultParams = resultParams[1:]
	}
	MoveContextResult(g.gen, resultParams)

	// For Go variables after the return statement.
	goReturns := pen.NewJoints(", ", 2)

	for i, converted := range g.Results {
		switch converted.Direction {
		case typeconv.ConvertGoToC: // parameter
			// Skip the instance parameter if any.
			if i != 0 || !instanceParam {
				g.GoArgs.Addf("%s %s", converted.InName, converted.In.Type)
			}

			// Go inputs are declared in the parameters, so no In.Declare.
			// C outputs have to be declared (input means C function input).
			// g.pen.Line(secInputDecl, converted.Out.Declare)
			// Conversions follow right after declaring all outputs.
			g.pen.Line(secInputConv, converted.Conversion)

		case typeconv.ConvertCToGo: // return
			// decoOut is the name that's used solely for documentation
			// purposes. It is not used internally at all, and so it doesn't
			// have the underscore.
			decoOut := strings.TrimPrefix(converted.OutName, "_")
			g.GoRets.Addf("%s %s", decoOut, converted.Out.Type)

			goReturns.Add(converted.OutName)

			// g.pen.Line(secInputDecl, converted.In.Declare)
			// Go outputs should be redeclared.
			g.pen.Line(secOutputDecl, converted.Out.Declare)
			// Conversions follow right after declaring all outputs.
			g.pen.Line(secOutputConv, converted.Conversion)
		}
	}

	g.pen.Linef(secFnCall, "_info := girepository.MustFind(%q, %q)", g.typ.Namespace.Name, g.typ.Name())

	var decl string
	if hasReturn {
		decl = "_gret :="
	}

	switch g.typ.Type.(type) {
	case *gir.Function:
		g.pen.Linef(secFnCall, "%s _info.InvokeFunction(%s, %s)", decl, argsName, outsName)
	case *gir.Class:
		g.pen.Linef(secFnCall, "%s _info.InvokeClassMethod(%q, %s, %s)", decl, g.CallableAttrs.Name, argsName, outsName)
	case *gir.Interface:
		g.pen.Linef(secFnCall, "%s _info.InvokeIfaceMethod(%q, %s, %s)", decl, g.CallableAttrs.Name, argsName, outsName)
	case *gir.Record:
		g.pen.Linef(secFnCall, "%s _info.InvokeRecordMethod(%q, %s, %s)", decl, g.CallableAttrs.Name, argsName, outsName)
	}

	if hasReturn {
		ret := g.Results[len(g.Results)-1]
		g.pen.Linef(secFnCall, "_cret := *(*%s)(unsafe.Pointer(&_gret))", ret.In.Type)
	}

	// Generate the right statements to ensure that nothing is freed before or
	// while we're invoking our function.
	for _, converted := range g.Results {
		if converted.Direction == typeconv.ConvertGoToC {
			g.hdr.Import("runtime")
			g.pen.Linef(secKeepAlive, "runtime.KeepAlive(%s)", converted.InName)
		}
	}

	if goReturns.Len() > 0 {
		g.pen.Line(secReturn, "return "+goReturns.Join())
	}

	g.Block = g.pen.String()
	g.Tail = "(" + g.GoArgs.Join() + ") " + formatReturnSig(g.GoRets)

	g.pen.Reset()
	return true
}

// EachParamResult iterates over the list of Go function parameters.
func (g *Generator) EachParamResult(f func(*typeconv.ValueConverted)) {
	direction := typeconv.ConvertGoToC
	if g.Parameters != nil && g.Parameters.InstanceParameter != nil {
		direction = g.Results[0].Direction
	}

	for i, res := range g.Results {
		if res.ParameterIndex == typeconv.ReceiverValueIndex {
			continue
		}
		if res.Direction == direction {
			f(&g.Results[i])
		}
	}
}

// EachReturnResult iterates over the list of Go function returns. Note that the
// direction is flipped if this is for a callback.
func (g *Generator) EachReturnResult(f func(*typeconv.ValueConverted)) {
	direction := typeconv.ConvertCToGo

	if g.Parameters != nil && g.Parameters.InstanceParameter != nil {
		// Flip.
		switch d := g.Results[0].Direction; d {
		case typeconv.ConvertCToGo:
			direction = typeconv.ConvertGoToC
		case typeconv.ConvertGoToC:
			direction = typeconv.ConvertCToGo
		default:
			log.Panicln("unknown instance parameter conv direction", d)
		}
	}

	for i, res := range g.Results {
		if res.Direction != direction {
			continue
		}
		if res.ParameterIndex == typeconv.ReturnValueIndex || res.ParameterIndex >= 0 {
			f(&g.Results[i])
		}
	}
}

// CoalesceTail calls CoalesceTail on the generator's tail.
func (g *Generator) CoalesceTail() {
	g.Tail = CoalesceTail(g.Tail)
}

var partsRegex = regexp.MustCompile(`\(.*?\)`)

// CoalesceTail coalesces certain parameters with the same type to be shorter.
func CoalesceTail(tail string) string {
	return partsRegex.ReplaceAllStringFunc(tail, func(whole string) string {
		part := whole
		part = strings.TrimPrefix(part, "(")
		part = strings.TrimSuffix(part, ")")

		params := strings.Split(part, ",")
		if !strings.Contains(params[0], " ") {
			// Probably a return; this should skip the whole loop.
			return whole
		}

		newParams := strings.Builder{}
		newParams.Grow(len(part))
		newParams.WriteByte('(')

		for i, param := range params {
			if !strings.Contains(param, " ") {
				// Probably a return; this should skip the whole loop.
				continue
			}

			if i == len(params)-1 {
				newParams.WriteString(param)
				break
			}

			n, t1 := splitParam(param)
			_, t2 := splitParam(params[i+1])

			if t1 == t2 {
				// Same type; write the name only.
				newParams.WriteString(n)
			} else {
				newParams.WriteString(param)
			}

			newParams.WriteString(", ")
		}

		newParams.WriteByte(')')
		return newParams.String()
	})
}

func splitParam(param string) (name, typ string) {
	param = strings.TrimSpace(param)

	parts := strings.SplitN(param, " ", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}

	panic("splitParam: invalid non-two-part param " + strconv.Quote(param))
}

func (g *Generator) Logln(lvl logger.Level, v ...interface{}) {
	g.gen.Logln(lvl, logger.Prefix(v, fmt.Sprintf("callable %s (C.%s)", g.Name, g.CIdentifier))...)
}

// findGCancellableParam finds the GCancellable type from the given list of
// parameters, but only if the list has one instance. -1 is returned otherwise.
func findGCancellableParam(g FileGenerator, params []gir.Parameter) int {
	found := -1

	for i, v := range params {
		if v.Skip || types.GuessParameterOutput(&v) != "in" || v.Type == nil {
			continue
		}

		gType := types.EnsureNamespace(g.Namespace(), v.Type.Name)
		cType := v.Type.CType

		// Ensure that the pointer level is what we expect.
		if gType == "Gio.Cancellable" && cType == "GCancellable*" {
			if found != -1 {
				// More than one instance; bail.
				return -1
			}

			found = i
		}
	}

	return found
}

// FindContextResult finds the context.Context type from the given list of
// callable results, but only if the list has one instance. -1 is returned
// otherwise.
func FindContextResult(g FileGenerator, result []typeconv.ValueConverted) int {
	found := -1

	for i, v := range result {
		if v.ConversionValue == nil || v.Resolved == nil {
			continue
		}

		// Ensure that the pointer level is what we expect.
		if v.Resolved.IsBuiltin("context.Context") {
			if found != -1 {
				// More than one instance; bail.
				return -1
			}

			found = i
		}
	}

	return found
}

// MoveContextResult moves the context variable. True is returned if the
// variable is found and moved.
func MoveContextResult(g FileGenerator, result []typeconv.ValueConverted) bool {
	ix := FindContextResult(g, result)
	if ix <= 0 {
		return false
	}

	cancellable := result[ix]
	// Shift everything up to the cancellable value up 1 value.
	copy(result[1:], result[:ix])
	// Set the first value to the cancellable one.
	result[0] = cancellable
	return true
}

func formatReturnSig(joints pen.Joints) string {
	if joints.Len() == 0 {
		return ""
	}

	parts := joints.Joints()
	types := make([]string, len(parts))

	for i, part := range parts {
		types[i] = extractTypeFromPair(part)
	}

	for i := range parts {
		for j := range parts {
			if i == j {
				continue
			}

			if types[i] == types[j] {
				goto dupeType
			}
		}
	}

	// No duplicate type, so only keep the types.
	joints.SetJoints(types)

dupeType:
	if joints.Len() == 1 {
		return joints.Join()
	}

	return "(" + joints.Join() + ")"
}

// extractTypeFromPair returns the second word (which is the type) from the
// name-type pair.
func extractTypeFromPair(namePair string) string {
	return namePair[strings.IndexByte(namePair, ' ')+1:]
}

func ReturnName(attrs *gir.CallableAttrs) string {
	if attrs.ReturnValue == nil {
		return ""
	}

	name := AnyTypeName(attrs.ReturnValue.AnyType, "ret")
	if attrs.Parameters == nil {
		return name
	}

	if attrs.Parameters.InstanceParameter != nil {
		if attrs.Parameters.InstanceParameter.Name == name {
			return "ret"
		}
	}

	for _, param := range attrs.Parameters.Parameters {
		if param.Name == name {
			return "ret"
		}
	}

	return name
}

// AnyTypeName returns the name from the given AnyType, or the given string if
// the type does not have a name.
func AnyTypeName(typ gir.AnyType, or string) string {
	switch {
	case typ.Type != nil:
		if typ.Type.Name == "gboolean" {
			return "ok"
		}
		parts := strings.Split(typ.Type.Name, ".")
		return strcases.UnexportPascal(parts[len(parts)-1])

	case typ.Array != nil:
		name := AnyTypeName(gir.AnyType{Type: typ.Array.Type}, or)
		if !strings.HasSuffix(name, "s") {
			return name + "s"
		}
		return name

	default:
		return or
	}
}

// Find finds a callable with the given Go name. The index within the slice is
// returned, or if nothing is found, then -1 is returned.
func Find(callables []Generator, goName string) int {
	for i, callable := range callables {
		if callable.Name == goName {
			return i
		}
	}
	return -1
}

// Grow grows or shrinks the callables slice to the given length. The returned
// slice will have a length of 0.
func Grow(callables []Generator, n int) []Generator {
	if cap(callables) <= n {
		return callables[:0]
	}
	return make([]Generator, 0, n*2)
}

// RenameGetters renames the given list of callables to have idiomatic Go getter
// names.
func RenameGetters(parentName string, callables []Generator) {
	for i, callable := range callables {
		newName, _ := RenameGetter(callable.Name)

		// Avoid duplicating method names with Objector.
		// TODO: account for other interfaces as well.
		objectorMethod := parentName != "" && types.IsObjectorMethod(newName)
		if objectorMethod {
			newName += parentName
		}

		if Find(callables, newName) > -1 {
			if !objectorMethod {
				continue
			}

			// We cannot not rename this method if it's an objectorMethod.
			newName += "_"
		}

		callables[i].Name = newName
	}
}

// RenameGetter renames a getter. True is returned if the name is changed.
func RenameGetter(name string) (string, bool) {
	if name == "ToString" {
		return "String", true
	}

	if strings.HasPrefix(name, "Get") && name != "Get" {
		return strings.TrimPrefix(name, "Get"), true
	}

	return name, false
}

// cFunctionHeader renders the given GIR function in its C function signature
// string for debugging or callback generation.
func cFunctionHeader(fn *gir.CallableAttrs) string {
	b := strings.Builder{}
	b.Grow(256)

	if fn.ReturnValue != nil {
		b.WriteString(resolveAnyCType(fn.ReturnValue.AnyType))
		b.WriteByte(' ')
	}

	b.WriteString(fn.CIdentifier)
	b.WriteByte('(')

	if fn.Parameters != nil && len(fn.Parameters.Parameters) > 0 {
		if fn.Parameters.InstanceParameter != nil {
			b.WriteString(resolveAnyCType(fn.Parameters.InstanceParameter.AnyType))
		}

		for i, param := range fn.Parameters.Parameters {
			if i != 0 || fn.Parameters.InstanceParameter != nil {
				b.WriteString(", ")
			}

			b.WriteString(resolveAnyCType(param.AnyType))
		}
	}

	b.WriteByte(')')

	return b.String()
}

// resolveAnyCType resolves an AnyType and returns the C type signature.
func resolveAnyCType(any gir.AnyType) string {
	switch {
	case any.Array != nil:
		return any.Array.CType
	case any.Type != nil:
		return any.Type.CType
	default:
		return "..."
	}
}

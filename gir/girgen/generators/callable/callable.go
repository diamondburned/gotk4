// Package callable provides a generic callable generator.
package callable

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/gir"
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

type ParamDoc struct {
	Name         string
	InfoElements gir.InfoElements
}

// Generator is a generator instance that generates a GIR callable.
type Generator struct {
	*gir.CallableAttrs
	Name  string
	Tail  string
	Block string

	constructor bool
	Converts    []string

	Conv    *typeconv.Converter
	Results []typeconv.ValueConverted
	GoArgs  pen.Joints
	GoRets  pen.Joints

	ParamDocs []ParamDoc

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
		pen: pen,
		gen: gen,
	}
}

// Reset resets the state of the generator while reusing the backing pen.
func (g *Generator) Reset() {
	g.hdr.Reset()
	g.pen.Reset()
	g.GoArgs.Reset(", ")
	g.GoRets.Reset(", ")

	*g = Generator{
		pen:    g.pen,
		gen:    g.gen,
		hdr:    g.hdr,
		GoArgs: g.GoArgs,
		GoRets: g.GoRets,

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
	g.EachParamResult(func(value *typeconv.ValueConverted) {
		g.ParamDocs = append(g.ParamDocs, ParamDoc{
			Name: value.InName, // GoName
			InfoElements: gir.InfoElements{
				DocElements: gir.DocElements{Doc: value.Doc},
			},
		})
	})

	return true
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
	const (
		secInputDecl = iota
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
		parameters := append([]gir.Parameter(nil), g.Parameters.Parameters...)

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
				out = fmt.Sprintf("_arg%d", i+1)
				dir = typeconv.ConvertGoToC
			case "out":
				in = fmt.Sprintf("_arg%d", i+1)
				out = "_" + strcases.SnakeToGo(false, param.Name)
				dir = typeconv.ConvertCToGo
			default:
				return false
			}

			if contextParamIx == i {
				// Idiomatic context naming.
				in = "ctx"
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
	if ix := findContextResult(g.gen, resultParams); ix > 0 {
		cancellable := resultParams[ix]
		// Shift everything up to the cancellable value up 1 value.
		copy(resultParams[1:], resultParams[:ix])
		// Set the first value to the cancellable one.
		resultParams[0] = cancellable
	}

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

	// For C function calling.
	callParams := strings.Join(g.Conv.CCallParams(), ", ")

	if !hasReturn {
		g.pen.Linef(secFnCall, "C.%s(%s)", g.CIdentifier, callParams)
	} else {
		g.pen.Linef(secFnCall, "_cret = C.%s(%s)", g.CIdentifier, callParams)
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

func (g *Generator) EachParamResult(f func(*typeconv.ValueConverted)) {
	results := g.Results
	if g.Parameters != nil && g.Parameters.InstanceParameter != nil {
		results = results[1:]
	}

	for i, res := range results {
		if res.Direction != typeconv.ConvertGoToC {
			continue
		}

		f(&results[i])
	}
}

// CoalesceTail calls CoalesceTail on the generator's tail.
func (g *Generator) CoalesceTail() {
	g.Tail = CoalesceTail(g.Tail)
}

// CoalesceTail coalesces certain parameters with the same type to be shorter.
func CoalesceTail(tail string) string {
	if !strings.HasPrefix(tail, "(") {
		return tail
	}

	paramIx := strings.Index(tail, ")")
	params := strings.Split(tail[1:paramIx], ",")

	newParams := strings.Builder{}
	newParams.Grow(len(tail))
	newParams.WriteByte('(')

	for i, param := range params {
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
	newParams.WriteString(tail[paramIx+1:])

	return newParams.String()
}

func splitParam(param string) (name, typ string) {
	parts := strings.Split(strings.TrimSpace(param), " ")
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

// findContextResult finds the context.Context type from the given list of
// callable results, but only if the list has one instance. -1 is returned
// otherwise.
func findContextResult(g FileGenerator, result []typeconv.ValueConverted) int {
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

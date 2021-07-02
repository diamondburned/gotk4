package callable

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/core/pen"
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

// FileGenerator describes the interface of a file generator.
type FileGenerator interface {
	types.FileGenerator
	file.Headerer
}

// Generator is a generator instance that generates a GIR callable.
type Generator struct {
	*gir.CallableAttrs
	Name  string
	Tail  string
	Block string

	ReturnWrap string // passed to ConversionValue, ctor only
	Converts   []string

	conv   *typeconv.Converter
	goArgs pen.Joints
	goRets pen.Joints

	gen FileGenerator
	pen *pen.BlockSections
}

// NewGenerator creates a new callable generator from the given generator.
func NewGenerator(gen FileGenerator) Generator {
	// Arbitrary sizes, whatever.
	pen := pen.NewBlockSections(1024, 4096, 256, 1024, 4096, 128)

	return Generator{
		gen: gen,
		pen: pen,
	}
}

// Reset resets the state of the generator while reusing the backing pen.
func (g *Generator) Reset() {
	if g.CallableAttrs == nil {
		return
	}

	g.pen.Reset()

	*g = Generator{
		gen:    g.gen,
		pen:    g.pen,
		goArgs: g.goArgs,
		goRets: g.goRets,
	}
}

func (g *Generator) Use(cattrs *gir.CallableAttrs) bool {
	g.Reset()

	if cattrs.ShadowedBy != "" || cattrs.MovedTo != "" {
		// Skip this one. Hope the caller reaches the Shadows method,
		// eventually.
		g.Reset()
		return false
	}
	if cattrs.CIdentifier == "" || !cattrs.IsIntrospectable() {
		g.Reset()
		return false
	}

	g.Name = strcases.SnakeToGo(true, cattrs.Name)
	g.CallableAttrs = cattrs

	if !g.renderBlock() {
		g.Reset()
		return false
	}

	return true
}

// Recv returns the receiver variable name. This method should only be called
// for methods.
func (g *Generator) Recv() string {
	if g.Parameters != nil && g.Parameters.InstanceParameter != nil {
		return strcases.FirstLetter(g.Parameters.InstanceParameter.Name)
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

			callableValues = append(callableValues, typeconv.NewValue(
				strcases.FirstLetter(g.Parameters.InstanceParameter.Name),
				"_arg0",
				-1,
				typeconv.ConvertGoToC,
				g.Parameters.InstanceParameter.ParameterAttrs,
			))
		}

		for i, param := range g.Parameters.Parameters {
			if param.Direction == "" {
				param.Direction = "in"
			}

			var in string
			var out string
			var dir typeconv.ConversionDirection

			switch param.Direction {
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

			value := typeconv.NewValue(in, out, i, dir, param.ParameterAttrs)
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
			if g.ReturnWrap != "" {
				value.WrapObject = g.ReturnWrap
			}

			callableValues = append(callableValues, value)
		}
	}

	if g.Throws {
		callableValues = append(callableValues, typeconv.NewThrowValue("_cerr", "_goerr"))
	}

	g.conv = typeconv.NewConverter(g.gen, callableValues)
	g.conv.UseLogger(cg)

	results := g.conv.ConvertAll()
	if results == nil {
		g.Logln(logger.Debug, "no conversion", CFunctionHeader(g.CallableAttrs))
		return false
	}

	// Apply imports and such.
	file.ApplyHeader(g.gen, g.conv)

	// For Go variables after the return statement.
	goReturns := pen.NewJoints(", ", 2)

	g.goArgs.Reset(", ")
	g.goRets.Reset(", ")

	for i, converted := range results {
		if converted.Skip {
			continue
		}

		switch converted.Direction {
		case typeconv.ConvertGoToC: // parameter
			// Skip the instance parameter if any.
			if i != 0 || !instanceParam {
				g.goArgs.Addf("%s %s", converted.InName, converted.InType)
			}

			// Go inputs are declared in the parameters, so no InDeclare.
			// C outputs have to be declared (input means C function input).
			g.pen.Line(secInputDecl, converted.OutDeclare)
			// Conversions follow right after declaring all outputs.
			g.pen.Line(secInputConv, converted.Conversion)

		case typeconv.ConvertCToGo: // return
			// decoOut is the name that's used solely for documentation
			// purposes. It is not used internally at all, and so it doesn't
			// have the underscore.
			decoOut := strings.TrimPrefix(converted.OutName, "_")
			g.goRets.Addf("%s %s", decoOut, converted.OutType)

			goReturns.Add(converted.OutName)

			g.pen.Line(secInputDecl, converted.InDeclare)
			// Go outputs should be redeclared.
			g.pen.Line(secOutputDecl, converted.OutDeclare)
			// Conversions follow right after declaring all outputs.
			g.pen.Line(secOutputConv, converted.Conversion)
		}
	}

	// For C function calling.
	callParams := strings.Join(typeconv.AddCCallParam(g.conv), ", ")

	if !hasReturn {
		g.pen.Linef(secFnCall, "C.%s(%s)", g.CIdentifier, callParams)
	} else {
		g.pen.Linef(secFnCall, "_cret = C.%s(%s)", g.CIdentifier, callParams)
		g.pen.EmptyLine(secFnCall)
	}

	if goReturns.Len() > 0 {
		g.pen.Line(secReturn, "return "+goReturns.Join())
	}

	g.Block = g.pen.String()
	g.Tail = "(" + g.goArgs.Join() + ") " + formatReturnSig(g.goRets)

	g.pen.Reset()
	return true
}

func (g *Generator) Logln(lvl logger.Level, v ...interface{}) {
	g.gen.Logln(lvl, logger.Prefix(v, fmt.Sprintf("callable %s (C.%s)", g.Name, g.CIdentifier)))
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
		name := AnyTypeName(typ.Array.AnyType, or)
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

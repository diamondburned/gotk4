package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

type callableGenerator struct {
	gir.CallableAttrs
	Name  string
	Tail  string
	Block string

	Converts []string

	pen *pen.BlockSections
	ng  *NamespaceGenerator
	fg  *FileGenerator
}

func newCallableGenerator(ng *NamespaceGenerator) callableGenerator {
	// Arbitrary sizes, whatever.
	pen := pen.NewBlockSections(1024, 4096, 256, 1024, 4096, 128)

	return callableGenerator{
		ng:  ng,
		pen: pen,
	}
}

func (cg *callableGenerator) reset() {
	cg.pen.Reset()

	*cg = callableGenerator{
		ng:  cg.ng,
		pen: cg.pen,
	}
}

func (cg *callableGenerator) Use(cattrs gir.CallableAttrs) bool {
	// Skip this one. Hope the caller reaches the Shadows method, eventually.
	if cattrs.ShadowedBy != "" || cattrs.MovedTo != "" || !cattrs.IsIntrospectable() {
		cg.reset()
		return false
	}

	cg.fg = cg.ng.FileFromSource(cattrs.DocElements)

	cg.Name = SnakeToGo(true, cattrs.Name)
	cg.CallableAttrs = cattrs

	if !cg.renderBlock() {
		cg.reset()
		return false
	}

	return true
}

// cFunctionSig renders the given GIR function in its C function signature
// string for debugging.
func cFunctionSig(fn gir.CallableAttrs) string {
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

// Recv returns the receiver variable name. This method should only be called
// for methods.
func (cg *callableGenerator) Recv() string {
	if cg.Parameters != nil && cg.Parameters.InstanceParameter != nil {
		return FirstLetter(cg.Parameters.InstanceParameter.Name)
	}

	return "v"
}

func (cg *callableGenerator) renderBlock() bool {
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
		callableValues []ConversionValue
	)

	if cg.Parameters != nil {
		callableValues = make([]ConversionValue, 0, len(cg.Parameters.Parameters)+2)

		if cg.Parameters.InstanceParameter != nil {
			instanceParam = true

			callableValues = append(callableValues, NewConversionValue(
				FirstLetter(cg.Parameters.InstanceParameter.Name), "_arg0", -1, ConvertGoToC,
				cg.Parameters.InstanceParameter.ParameterAttrs,
			))
		}

		for i, param := range cg.Parameters.Parameters {
			if param.Direction == "" {
				param.Direction = "in"
			}

			var in string
			var out string
			var dir ConversionDirection

			switch param.Direction {
			case "in":
				in = SnakeToGo(false, param.Name)
				out = fmt.Sprintf("_arg%d", i+1)
				dir = ConvertGoToC
			case "out":
				in = fmt.Sprintf("_arg%d", i+1)
				out = "_" + SnakeToGo(false, param.Name)
				dir = ConvertCToGo
			default:
				return false
			}

			value := NewConversionValue(in, out, i, dir, param.ParameterAttrs)
			callableValues = append(callableValues, value)
		}
	}

	var hasReturn bool
	if !returnIsVoid(cg.ReturnValue) {
		returnName := returnName(cg.CallableAttrs)

		// If the last return is a bool and the function can throw an error,
		// then the boolean is probably to indicate that things are OK. We can
		// skip generating this boolean.
		if !cg.Throws || returnName != "ok" {
			hasReturn = true
			returnName = "_" + returnName

			value := NewConversionValueReturn("_cret", returnName, ConvertCToGo, *cg.ReturnValue)
			callableValues = append(callableValues, value)
		}
	}

	if cg.Throws {
		callableValues = append(callableValues, NewThrowValue("_cerr", "_goerr"))
	}

	convert := NewTypeConverter(cg.fg, cg.Name, callableValues)
	results := convert.ConvertAll()
	if results == nil {
		cg.fg.Logln(LogSkip, "callable has no conversion", cFunctionSig(cg.CallableAttrs))
		return false
	}

	cg.fg.applyConvertedFxs(results)

	// For Go variables after the return statement.
	goReturns := pen.NewJoints(", ", 2)

	goFnArgs := pen.NewJoints(", ", len(results))
	goFnRets := pen.NewJoints(", ", len(results))

	for i, converted := range results {
		if converted.Skip {
			continue
		}

		switch converted.Direction {
		case ConvertGoToC: // parameter
			// Skip the instance parameter if any.
			if i != 0 || !instanceParam {
				goFnArgs.Addf("%s %s", converted.InName, converted.InType)
			}

			// Go inputs are declared in the parameters, so no InDeclare.
			// C outputs have to be declared (input means C function input).
			cg.pen.Line(secInputDecl, converted.OutDeclare)
			// Conversions follow right after declaring all outputs.
			cg.pen.Line(secInputConv, converted.Conversion)

		case ConvertCToGo: // return
			// decoOut is the name that's used solely for documentation
			// purposes. It is not used internally at all, and so it doesn't
			// have the underscore.
			decoOut := strings.TrimPrefix(converted.OutName, "_")
			goFnRets.Addf("%s %s", decoOut, converted.OutType)

			goReturns.Add(converted.OutName)

			cg.pen.Line(secInputDecl, converted.InDeclare)
			// Go outputs should be redeclared.
			cg.pen.Line(secOutputDecl, converted.OutDeclare)
			// Conversions follow right after declaring all outputs.
			cg.pen.Line(secOutputConv, converted.Conversion)
		}
	}

	// For C function calling.
	callParams := strings.Join(AddCCallParam(convert), ", ")

	if !hasReturn {
		cg.pen.Linef(secFnCall, "C.%s(%s)", cg.CIdentifier, callParams)
	} else {
		cg.pen.Linef(secFnCall, "_cret = C.%s(%s)", cg.CIdentifier, callParams)
		cg.pen.EmptyLine(secFnCall)
	}

	if goReturns.Len() > 0 {
		cg.pen.Line(secReturn, "return "+goReturns.Join())
	}

	cg.Block = cg.pen.String()
	cg.Tail = "(" + goFnArgs.Join() + ") " + formatReturnSig(goFnRets)

	cg.pen.Reset()
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

func returnName(attrs gir.CallableAttrs) string {
	if attrs.ReturnValue == nil {
		return ""
	}

	name := anyTypeName(attrs.ReturnValue.AnyType, "ret")

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

func anyTypeName(typ gir.AnyType, or string) string {
	switch {
	case typ.Type != nil:
		if typ.Type.Name == "gboolean" {
			return "ok"
		}
		parts := strings.Split(typ.Type.Name, ".")
		return UnexportPascal(parts[len(parts)-1])

	case typ.Array != nil:
		name := anyTypeName(typ.Array.AnyType, or)
		if !strings.HasSuffix(name, "s") {
			return name + "s"
		}
		return name

	default:
		return or
	}
}

// callableRenameGetters renames the given list of callables to have idiomatic
// Go getter names.
func callableRenameGetters(callables []callableGenerator) {
	methodNames := make(map[string]struct{}, len(callables))
	for _, callable := range callables {
		methodNames[callable.Name] = struct{}{}
	}

	for i, callable := range callables {
		var newName string

		switch callable.Name {
		case "ToString":
			newName = "String"

		default:
			if !strings.HasPrefix(callable.Name, "Get") || callable.Name == "Get" {
				continue
			}

			newName = strings.TrimPrefix(callable.Name, "Get")
		}

		_, dup := methodNames[newName]
		if dup {
			continue // skip
		}

		delete(methodNames, callable.Name)
		methodNames[newName] = struct{}{}

		callables[i].Name = newName
	}
}

// callableGrow grows or shrinks the callables slice to the given length. The
// returned slice will have a length of 0.
func callableGrow(callables []callableGenerator, n int) []callableGenerator {
	if cap(callables) <= n {
		return callables[:0]
	}
	return make([]callableGenerator, 0, n*2)
}

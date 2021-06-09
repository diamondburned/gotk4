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
	pen := pen.NewBlockSections(1024, 4096, 1024, 1024, 256, 4096, 128)

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
	if cattrs.ShadowedBy != "" || cattrs.MovedTo != "" {
		cg.reset()
		return false
	}

	cg.fg = cg.ng.FileFromSource(cattrs.SourcePosition)

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

	b.WriteString(fn.Name)
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
		secReturnDecl
		secFnCall
		secOutputDecl
		secOutputConv
		secReturn
	)

	var (
		instanceParam bool
		inputValues   []GoValueProp
		outputValues  []CValueProp
	)

	if cg.Parameters != nil {
		inputValues = make([]GoValueProp, 0, len(cg.Parameters.Parameters))
		outputValues = make([]CValueProp, 0, len(cg.Parameters.Parameters)+2)

		if cg.Parameters.InstanceParameter != nil {
			instanceParam = true
			// params.Add("arg0")

			inputValues = append(inputValues, GoValueProp{
				ValueProp: NewValuePropParam(
					FirstLetter(cg.Parameters.InstanceParameter.Name), "arg0", nil,
					cg.Parameters.InstanceParameter.ParameterAttrs,
				),
			})
		}

		for i, param := range cg.Parameters.Parameters {
			if param.Direction != "out" {
				out := fmt.Sprintf("arg%d", i+1)
				// params.Add(out)

				inputValues = append(inputValues, GoValueProp{
					ValueProp: NewValuePropParam(
						SnakeToGo(false, param.Name), out, &i,
						param.ParameterAttrs,
					),
				})
			} else {
				in := fmt.Sprintf("arg%d", i+1)
				// params.Add(in)

				out := SnakeToGo(false, param.Name)
				// returns.Add(out)
				// returnSigs = append(returnSigs, SnakeToGo(false, param.Name))

				outputValues = append(outputValues, CValueProp{
					ValueProp: NewValuePropParam(in, out, &i, param.ParameterAttrs),
				})
			}
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

			outputValues = append(outputValues, CValueProp{
				ValueProp: NewValuePropReturn("cret", returnName, *cg.ReturnValue),
			})
		}
	}

	if cg.Throws {
		outputValues = append(outputValues, CValueProp{
			ValueProp: NewThrowValue("cerr", "goerr"),
		})
	}

	convertedInputs := cg.fg.GoCConverter(cg.Name, inputValues).ConvertAll()
	if convertedInputs == nil {
		cg.fg.Logln(LogSkip, "callable (no Go->C conversion)", cFunctionSig(cg.CallableAttrs))
		return false
	}

	convertedOutputs := cg.fg.CGoConverter(cg.Name, outputValues).ConvertAll()
	if convertedOutputs == nil {
		cg.fg.Logln(LogSkip, "callable (no C->Go conversion)", cFunctionSig(cg.CallableAttrs))
		return false
	}

	// For C function calling.
	callParams := pen.NewJoints(", ", len(convertedInputs)+len(convertedOutputs))
	// For Go variables after the return statement.
	goReturns := pen.NewJoints(", ", 2)

	goFnArgs := pen.NewJoints(", ", len(convertedInputs))
	goFnRets := pen.NewJoints(", ", len(convertedOutputs))

	for i, converted := range convertedInputs {
		converted.Apply(cg.fg)
		callParams.Add(converted.OutCall)

		// Skip the instance parameter if any.
		if i != 0 || !instanceParam {
			goFnArgs.Addf("%s %s", converted.InName, converted.InType)
		}

		// Go inputs are declared in the parameters, so no InDeclare.
		// C outputs have to be declared (input means C function input).
		cg.pen.Line(secInputDecl, converted.OutDeclare)
		// Conversions follow right after declaring all outputs.
		cg.pen.Line(secInputConv, converted.Conversion)
	}

	for _, converted := range convertedOutputs {
		converted.Apply(cg.fg)
		goReturns.Add(converted.OutName)

		if converted.ParameterIsOutput {
			callParams.Add(converted.InCall)
		}

		// The return variable is not declared in the signature if there's only
		// 1 output, so we only declare it then.
		goFnRets.Addf("%s %s", converted.OutName, converted.OutType)

		cg.pen.Line(secReturnDecl, converted.InDeclare)
		// Go outputs should be redeclared.
		cg.pen.Line(secOutputDecl, converted.OutDeclare)
		// Conversions follow right after declaring all outputs.
		cg.pen.Line(secOutputConv, converted.Conversion)
	}

	if !hasReturn {
		cg.pen.Linef(secFnCall, "C.%s(%s)", cg.CIdentifier, callParams.Join())
	} else {
		cg.pen.Linef(secFnCall, "cret = C.%s(%s)", cg.CIdentifier, callParams.Join())
		cg.pen.EmptyLine(secFnCall)
	}

	if len(outputValues) > 0 {
		cg.pen.Line(secReturn, "return "+goReturns.Join())
	}

	cg.Block = cg.pen.String()
	cg.pen.Reset()

	cg.Tail = "(" + goFnArgs.Join() + ")"

	switch goFnRets.Len() {
	case 0:
	// ok
	case 1:
		cg.Tail += " " + strings.SplitN(goFnRets.Join(), " ", 2)[1] // type only
	default:
		cg.Tail += " (" + goFnRets.Join() + ")"
	}

	return true
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

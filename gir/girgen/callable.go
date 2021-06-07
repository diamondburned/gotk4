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
	pen := pen.NewBlockSections(1024, 4096, 1024, 256, 4096, 128)

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
	if cattrs.ShadowedBy != "" {
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
		secOutputConv
		secReturn
	)

	var (
		params  pen.Joints
		returns pen.Joints

		inputValues  []GoValueProp
		outputValues []CValueProp
	)

	if cg.Parameters != nil {
		params = pen.NewJoints(", ", len(cg.Parameters.Parameters))
		returns = pen.NewJoints(", ", len(cg.Parameters.Parameters)+2)

		inputValues = make([]GoValueProp, 0, len(cg.Parameters.Parameters))
		outputValues = make([]CValueProp, 0, len(cg.Parameters.Parameters)+2)

		if cg.Parameters.InstanceParameter != nil {
			params.Add("arg0")

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
				params.Add(out)

				inputValues = append(inputValues, GoValueProp{
					ValueProp: NewValuePropParam(
						SnakeToGo(false, param.Name), out, &i,
						param.ParameterAttrs,
					),
				})
			} else {
				in := fmt.Sprintf("arg%d", i+1)
				params.Add("&" + in)

				out := SnakeToGo(false, param.Name)
				returns.Add(out)

				outputValues = append(outputValues, CValueProp{
					ValueProp: NewValuePropParam(in, out, &i, param.ParameterAttrs),
				})
			}
		}
	}

	if cg.Throws {
		params.Add("&errout")
		returns.Add("err")

		outputValues = append(outputValues, CValueProp{
			ValueProp: newThrowValue("errout", "err"),
		})
	}

	var hasReturn bool
	// If the last return is a bool and the function can throw an error,
	// then the boolean is probably to indicate that things are OK. We can
	// skip generating this boolean.
	if !returnIsVoid(cg.ReturnValue) {
		returnName := anyTypeName(cg.ReturnValue.AnyType, "ret")

		if !cg.Throws && returnName == "ok" {
			hasReturn = true

			returns.Add(returnName)
			outputValues = append(outputValues, CValueProp{
				ValueProp: NewValuePropReturn("cret", returnName, *cg.ReturnValue),
			})
		}
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

	goArgs := pen.NewJoints(", ", len(convertedInputs))
	for _, converted := range convertedInputs {
		goArgs.Addf("%s %s", converted.In, converted.InType)

		// Go inputs are declared in the parameters, so no InDeclare.
		// C outputs have to be declared (input means C function input).
		cg.pen.Line(secInputDecl, converted.OutDeclare)
		// Conversions follow right after declaring all outputs.
		cg.pen.Line(secInputConv, converted.Conversion)
	}

	goRets := pen.NewJoints(", ", len(convertedOutputs))
	for _, converted := range convertedOutputs {
		goRets.Addf("%s %s", converted.Out, converted.OutType)

		cg.pen.Line(secReturnDecl, converted.InDeclare)
		// Go outputs should be redeclared.
		cg.pen.Line(secReturnDecl, converted.OutDeclare)
		// Conversions follow right after declaring all outputs.
		cg.pen.Line(secOutputConv, converted.Conversion)
	}

	if !hasReturn {
		cg.pen.Linef(secFnCall, "C.%s(%s)", cg.CIdentifier, params.Join())
	} else {
		cg.pen.Linef(secFnCall, "cret = C.%s(%s)", cg.CIdentifier, params.Join())
		cg.pen.EmptyLine(secFnCall)
	}

	if len(outputValues) > 0 {
		cg.pen.Line(secReturn, "return "+returns.Join())
	}

	cg.Block = cg.pen.String()
	cg.pen.Reset()

	cg.Tail = "(" + goArgs.Join() + ")"
	switch goRets.Len() {
	case 0:
	// ok
	case 1:
		cg.Tail += " " + strings.SplitN(goRets.Join(), " ", 2)[1] // type only
	default:
		cg.Tail += " (" + goRets.Join() + ")"
	}

	return true
}

// fnCall generates the tail of the function, that is, everything underlined
// below:
//
//    func FunctionName(arguments...) (returns...)
//                     ^^^^^^^^^^^^^^^^^^^^^^^^^^^
// An empty string is returned if the function cannot be generated.
func (fg *FileGenerator) fnCall(attrs gir.CallableAttrs) string {
	args, ok := fg.fnArgs(attrs)
	if !ok {
		fg.Logln(LogDebug, "fnArgs failed for callable", cFunctionSig(attrs))
		return ""
	}

	returns, ok := fg.fnReturns(attrs)
	if !ok {
		fg.Logln(LogDebug, "fnReturns failed for callable", cFunctionSig(attrs))
		return ""
	}

	return "(" + args + ") " + returns
}

// fnArgs returns the function arguments as a Go string and true. It returns
// false if the argument types cannot be fully resolved.
func (fg *FileGenerator) fnArgs(attrs gir.CallableAttrs) (string, bool) {
	if attrs.Parameters == nil || len(attrs.Parameters.Parameters) == 0 {
		return "", true
	}

	goArgs := make([]string, 0, len(attrs.Parameters.Parameters))

	ok := iterateParams(attrs, func(_ int, param gir.Parameter) bool {
		goName := SnakeToGo(false, param.Name)

		resolved, ok := GoAnyType(fg, param.AnyType, true)
		if !ok {
			if goName == "..." {
				fg.Logln(LogSkip, "function", attrs.Name, "is variadic")
			} else {
				fg.Logln(LogUnknown, "function argument", goName, "for", attrs.Name)
			}

			return false
		}

		goArgs = append(goArgs, goName+" "+resolved)
		return true
	})

	if !ok {
		return "", false
	}

	return strings.Join(goArgs, ", "), true
}

// fnReturns returns the function return type and true. It returns false if the
// function's return type cannot be resolved.
func (fg *FileGenerator) fnReturns(attrs gir.CallableAttrs) (string, bool) {
	var returns []string

	ok := iterateReturns(attrs, func(goName string, i int, any gir.AnyType) bool {
		typ, ok := GoAnyType(fg, any, true)
		if !ok {
			fg.Logln(LogUnknown, "function output", goName, "for", attrs.Name)
			return false
		}

		// if parameter
		if i != -1 {
			// Hacky way to "dereference" a pointer once.
			if strings.HasPrefix(typ, "*") {
				typ = typ[1:]
			}
		}

		// if returning bool and we're throwing, then skip
		if i == -1 && attrs.Throws && goName == "ok" {
			return true
		}

		returns = append(returns, goName+" "+typ)
		return true
	})

	if attrs.Throws {
		returns = append(returns, "err error")
	}

	if len(returns) == 0 || !ok {
		return "", ok
	}
	if len(returns) == 1 {
		// Only use the type if we have 1 return.
		return strings.Split(returns[0], " ")[1], true
	}

	return "(" + strings.Join(returns, ", ") + ")", true
}

// iterateParams iterates over parameters.
func iterateParams(attr gir.CallableAttrs, fn func(int, gir.Parameter) bool) bool {
	if attr.Parameters == nil {
		return true
	}

	var ignores ignoreIxs

	for i, param := range attr.Parameters.Parameters {
		ignores.paramIgnore(param)

		ignore := ignores.ignore(i) ||
			// Ignore out params (treat as return).
			(param.Direction == "out") ||
			// Ignore exposing destroy notifiers.
			(param.Name == "destroy_fn") ||
			(param.Type != nil && strings.HasSuffix(param.Type.Name, "DestroyNotify"))

		if ignore {
			continue
		}

		if !fn(i, param) {
			return false
		}
	}

	return true
}

// iterateReturns iterates over returns. The given index integer is -1 if the
// given type is from the return. The given string is the Go name.
func iterateReturns(attr gir.CallableAttrs, fn func(string, int, gir.AnyType) bool) bool {
	if attr.Parameters != nil {
		for i, param := range attr.Parameters.Parameters {
			if param.Direction != "out" || param.AnyType.VarArgs != nil {
				continue
			}

			name := SnakeToGo(false, param.Name)
			if name == "error" {
				name = "err"
			}

			if !fn(name, i, param.AnyType) {
				return false
			}
		}
	}

	if !returnIsVoid(attr.ReturnValue) {
		retName := anyTypeName(attr.ReturnValue.AnyType, "ret")
		if !fn(UnexportPascal(retName), -1, attr.ReturnValue.AnyType) {
			return false
		}
	}

	return true
}

func anyTypeName(typ gir.AnyType, or string) string {
	switch {
	case typ.Type != nil:
		if typ.Type.Name == "gboolean" {
			return "ok"
		}
		parts := strings.Split(typ.Type.Name, ".")
		return parts[len(parts)-1]

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

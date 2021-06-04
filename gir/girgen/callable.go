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

	ng *NamespaceGenerator
	fg *FileGenerator
}

func newCallableGenerator(ng *NamespaceGenerator) callableGenerator {
	return callableGenerator{ng: ng}
}

func (cg *callableGenerator) reset() {
	*cg = callableGenerator{ng: cg.ng}
}

func (cg *callableGenerator) Use(cattrs gir.CallableAttrs) bool {
	// Skip this one. Hope the caller reaches the Shadows method, eventually.
	if cattrs.ShadowedBy != "" {
		cg.reset()
		return false
	}

	cg.fg = cg.ng.FileFromSource(cattrs.SourcePosition)

	call := cg.fg.fnCall(cattrs)
	if call == "" {
		cg.reset()
		return false
	}

	cg.Name = SnakeToGo(true, cattrs.Name)
	cg.Tail = call
	cg.CallableAttrs = cattrs

	if cg.Block = cg.block(); cg.Block == "" {
		cg.fg.Logln(LogSkip, "callable (no block)", cFunctionSig(cattrs))
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

// Block renders the function block. It returns an empty string if a conversion
// cannot be generated.
func (cg *callableGenerator) block() string {
	type retVar struct {
		Name      string
		Type      gir.AnyType
		Owner     gir.TransferOwnership
		Return    bool
		AllowNone bool
	}

	blocks := pen.NewBlockSections(4096, 256, 4096, 512)

	var (
		params  pen.Joints
		returns pen.Joints

		inputValues  []GoValueProp
		outputValues []CValueProp
	)

	if cg.Parameters != nil {
		cap := len(cg.Parameters.Parameters) + 2

		params = pen.NewJoints(", ", cap)
		returns = pen.NewJoints(", ", cap)

		inputValues = make([]GoValueProp, 0, cap)
		outputValues = make([]CValueProp, 0, cap)

		if cg.Parameters.InstanceParameter != nil {
			in := FirstLetter(cg.Parameters.InstanceParameter.Name)
			params.Add(in)

			inputValues = append(inputValues, GoValueProp{
				ValueProp: NewValuePropParam(
					in, "arg0",
					cg.Parameters.InstanceParameter.ParameterAttrs,
				),
			})
		}

		for i, param := range cg.Parameters.Parameters {
			if param.Direction != "out" {
				in := SnakeToGo(false, param.Name)
				params.Add(in)

				inputValues = append(inputValues, GoValueProp{
					ValueProp: NewValuePropParam(
						in, fmt.Sprintf("arg%d", i+1), // declared
						param.ParameterAttrs,
					),
				})
			} else {
				in := fmt.Sprintf("arg%d", i+1)
				params.Add("&" + in)

				out := fmt.Sprintf("ret%d", i+1)
				returns.Add(out)

				outputValues = append(outputValues, CValueProp{
					ValueProp:   NewValuePropParam(in, out, param.ParameterAttrs),
					OutputParam: true,
				})
			}
		}
	}

	// If the last return is a bool and the function can throw an error,
	// then the boolean is probably to indicate that things are OK. We can
	// skip generating this boolean.
	if !returnIsVoid(cg.ReturnValue) {
		if !cg.Throws && anyTypeName(cg.ReturnValue.AnyType, "") != "ok" {
			out := fmt.Sprintf("ret%d", len(outputValues)+1)
			returns.Add(out)

			outputValues = append(outputValues, CValueProp{
				ValueProp: NewValuePropReturn("ret", out, *cg.ReturnValue),
			})
		}
	}

	if cg.Throws {
		blocks.Line(0, "var errout *C.GError")
		params.Add("&errout")

		blocks.Linef(3, "var goerr error")
		returns.Add("goerr")

		blocks.Line(3, "if errout != nil {")
		blocks.Line(3, `  goerr = fmt.Errorf("%d: %s", errout.code, C.GoString(errout.message))`)
		blocks.Line(3, "  C.g_error_free(errout)")
		blocks.Line(3, "}")
		blocks.EmptyLine(3)
	}

	var (
		paramsConv = cg.fg.GoCConverter(TypeConversionToC{
			Values: inputValues,
			Parent: cg.Name,
		})
		returnsConv = cg.fg.CGoConverter(TypeConversionToGo{
			Values: outputValues,
			Parent: cg.Name,
		})
	)

	if paramsConv == nil || returnsConv == nil {
		return ""
	}

	paramsConv.Apply(cg.fg)
	returnsConv.Apply(cg.fg)

	blocks.Linef(0, paramsConv.Conversion)
	blocks.Linef(0, returnsConv.Preamble)
	blocks.Linef(2, returnsConv.Conversion)

	if returnIsVoid(cg.ReturnValue) {
		blocks.Linef(1, "C.%s(%s)", cg.CIdentifier, params.Join())
	} else {
		blocks.Linef(1, "ret := C.%s(%s)", cg.CIdentifier, params.Join())
	}

	if len(outputValues) > 0 {
		blocks.EmptyLine(1)
		blocks.Line(3, "return "+returns.Join())
	}

	return blocks.String()
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
		return ""
	}

	returns, ok := fg.fnReturns(attrs)
	if !ok {
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

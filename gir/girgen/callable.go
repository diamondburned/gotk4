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

	Ng *NamespaceGenerator
}

func newCallableGenerator(ng *NamespaceGenerator) callableGenerator {
	return callableGenerator{Ng: ng}
}

func (cg *callableGenerator) reset() {
	*cg = callableGenerator{Ng: cg.Ng}
}

func (cg *callableGenerator) Use(cattrs gir.CallableAttrs) bool {
	// Skip this one. Hope the caller reaches the Shadows method, eventually.
	if cattrs.ShadowedBy != "" {
		cg.reset()
		return false
	}

	call := cg.Ng.FnCall(cattrs)
	if call == "" {
		cg.reset()
		return false
	}

	cg.Name = SnakeToGo(true, cattrs.Name)
	cg.Tail = call
	cg.CallableAttrs = cattrs

	if cg.Block = cg.block(); cg.Block == "" {
		cg.reset()
		return false
	}

	return true
}

// Recv returns the receiver variable name. This method should only be called
// for methods.
func (cg *callableGenerator) Recv() string {
	if cg.Parameters != nil && cg.Parameters.InstanceParameter != nil {
		return SnakeToGo(false, cg.Parameters.InstanceParameter.Name)
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
		AllowNone bool
	}

	blocks := pen.NewBlockSections(1024, 4096)

	var args pen.Joints
	var rets []retVar

	// argNamer returns arg with a +1 offset, except for when i is negative,
	// then the method receiver is returned.
	argNamer := func(i int) string {
		if i < 0 {
			return cg.Recv()
		}

		return fmt.Sprintf("arg%d", i+1)
	}

	if cg.Parameters != nil {
		var ignores ignoreIxs
		ignores.paramsIgnore(cg.Parameters)

		// Preallocate if possible.
		args = *pen.NewJoints(", ", len(cg.Parameters.Parameters))

		doParam := func(i int, param gir.ParameterAttrs, targ, valn string) bool {
			// Generate a zero-value variable regardless if we have the
			// conversions or not.
			blocks.Linef(0, "var %s %s", targ, anyTypeCGo(param.AnyType))
			args.Add(targ)

			// Ignored arguments may be covered by GoCConverter if the
			// attributes are part of the type and not the parameter attributes.
			if !ignores.ignore(i) {
				converter := cg.Ng.GoCConverter(TypeConversionToC{
					TypeConversion: TypeConversion{
						Value:      valn,
						Target:     targ,
						Type:       param.AnyType,
						Owner:      param.TransferOwnership,
						ArgAt:      argNamer,
						ParentName: cg.CIdentifier,
					},
					Closure: param.Closure,
					Destroy: param.Destroy,
				})

				if converter != "" {
					blocks.Line(1, converter)
					return true
				}

				// If null is fine, then we're fine.
				return param.AllowNone
			}

			return true
		}

		if cg.Parameters.InstanceParameter != nil {
			attrs := cg.Parameters.InstanceParameter.ParameterAttrs
			if !doParam(-1, attrs, "arg0", cg.Recv()) {
				return ""
			}
		}

		for i, param := range cg.Parameters.Parameters {
			// Be careful with using the index `i`, since the name is different
			// from the index within the slice.

			targ := argNamer(i)
			valn := SnakeToGo(false, param.Name)

			if param.Direction == "out" {
				blocks.Linef(0, "var %s %s // out", targ, anyTypeCGo(param.AnyType))

				args.Add("&" + targ)
				rets = append(rets, retVar{
					Name:      targ,
					Type:      param.AnyType,
					Owner:     param.TransferOwnership,
					AllowNone: param.AllowNone,
				})
				continue
			}

			// TODO: nullability.
			// TODO: GoCConverter.

			if !doParam(i, param.ParameterAttrs, targ, valn) {
				return ""
			}
		}
	}

	if cg.Throws {
		blocks.Linef(0, "var gError *C.GError")
		args.Add("&gError")
	}

	funcReturns := !returnIsVoid(cg.ReturnValue)
	if funcReturns {
		rets = append(rets, retVar{
			Name:  "ret",
			Type:  cg.ReturnValue.AnyType,
			Owner: cg.ReturnValue.TransferOwnership,
		})
	}

	if !funcReturns {
		blocks.Linef(2, "C.%s(%s)", cg.CIdentifier, args.Join())
	} else {
		blocks.Linef(2, "ret := C.%s(%s)", cg.CIdentifier, args.Join())
	}

	// If there is no return NOR output parameters, then exit.
	if len(rets) == 0 {
		return blocks.String()
	}

	blocks.EmptyLine(2)

	retvars := pen.NewJoints(", ", len(rets))

	for i, ret := range rets {
		// If the last return is a bool and the function can throw an error,
		// then the boolean is probably to indicate that things are OK. We can
		// skip generating this boolean.
		if cg.Throws && i == len(rets)-1 && anyTypeName(ret.Type, "") == "ok" {
			break
		}

		goType, _ := cg.Ng.ResolveAnyType(ret.Type, true)
		retName := fmt.Sprintf("ret%d", i)
		retvars.Add(retName)

		blocks.Linef(3, "var %s %s", retName, goType)

		converter := cg.Ng.CGoConverter(TypeConversionToGo{
			TypeConversion: TypeConversion{
				Value:      ret.Name,
				Target:     retName,
				Type:       ret.Type,
				Owner:      ret.Owner,
				ArgAt:      argNamer,
				ParentName: cg.CIdentifier,
			},
			BoxCast: goType,
		})

		if converter != "" {
			blocks.Line(4, converter)
			blocks.EmptyLine(4)
			continue
		}

		if !ret.AllowNone {
			return ""
		}
	}

	if cg.Throws {
		blocks.Linef(3, "var goError error")
		retvars.Add("goError")

		blocks.Line(4, "if gError != nil {")
		blocks.Line(4, `  goError = fmt.Errorf("%d: %s", gError.code, C.GoString(gError.message))`)
		blocks.Line(4, "  C.g_error_free(gError)")
		blocks.Line(4, "}")
	}

	blocks.Line(5, "return "+retvars.Join())
	return blocks.String()
}

// FnCall generates the tail of the function, that is, everything underlined
// below:
//
//    func FunctionName(arguments...) (returns...)
//                     ^^^^^^^^^^^^^^^^^^^^^^^^^^^
// An empty string is returned if the function cannot be generated.
func (ng *NamespaceGenerator) FnCall(attrs gir.CallableAttrs) string {
	args, ok := ng.FnArgs(attrs)
	if !ok {
		return ""
	}

	returns, ok := ng.FnReturns(attrs)
	if !ok {
		return ""
	}

	return "(" + args + ") " + returns
}

// FnArgs returns the function arguments as a Go string and true. It returns
// false if the argument types cannot be fully resolved.
func (ng *NamespaceGenerator) FnArgs(attrs gir.CallableAttrs) (string, bool) {
	if attrs.Parameters == nil || len(attrs.Parameters.Parameters) == 0 {
		return "", true
	}

	goArgs := make([]string, 0, len(attrs.Parameters.Parameters))

	ok := iterateParams(attrs, func(_ int, param gir.Parameter) bool {
		resolved, ok := ng.ResolveAnyType(param.AnyType, true)
		if !ok {
			return false
		}

		goName := SnakeToGo(false, param.Name)
		goArgs = append(goArgs, goName+" "+resolved)
		return true
	})

	if !ok {
		return "", false
	}

	return strings.Join(goArgs, ", "), true
}

// FnReturns returns the function return type and true. It returns false if the
// function's return type cannot be resolved.
func (ng *NamespaceGenerator) FnReturns(attrs gir.CallableAttrs) (string, bool) {
	var returns []string

	ok := iterateReturns(attrs, func(goName string, i int, any gir.AnyType) bool {
		typ, ok := ng.ResolveAnyType(any, true)
		if !ok {
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

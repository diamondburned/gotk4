package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

type callableGenerator struct {
	gir.CallableAttrs
	Name string
	Tail string

	Ng *NamespaceGenerator
}

func newCallableGenerator(ng *NamespaceGenerator) callableGenerator {
	return callableGenerator{Ng: ng}
}

func (cg *callableGenerator) Use(cattrs gir.CallableAttrs) bool {
	call := cg.Ng.FnCall(cattrs)
	if call == "" {
		return false
	}

	cg.Name = SnakeToGo(true, cattrs.Name)
	cg.Tail = call
	cg.CallableAttrs = cattrs

	return true
}

// Block renders the function block.
func (cg *callableGenerator) Block() string {
	type retVar struct {
		Name string
		Type gir.AnyType
	}

	var block pen.Block
	var args []string
	var rets []retVar
	var ignores ignoreIxs

	argNamer := func(i int) string { return fmt.Sprintf("arg%d", i) }

	if cg.Parameters != nil {
		args = make([]string, 0, len(cg.Parameters.Parameters))
		var closure *int // user-data arg
		var destroy *int

		for i, param := range cg.Parameters.Parameters {
			if param.Closure != nil {
				closure = param.Closure
			}
			if param.Destroy != nil {
				destroy = param.Destroy
			}

			ignores.paramIgnore(param)
			targ := argNamer(i)
			valn := SnakeToGo(false, param.Name)

			if param.Direction == "out" {
				block.Linef("var %s %s // out", targ, anyTypeCGo(param.AnyType))
				block.EmptyLine()

				args = append(args, "&"+targ)
				rets = append(rets, retVar{
					Name: targ,
					Type: param.AnyType,
				})
				continue
			}

			// Handle non-closure destroy notifiers: since we're always copying
			// Go memory to C using malloc, we can just give it a free.
			if param.Type != nil &&
				param.Type.Name != "GLib.DestroyNotify" &&
				strings.HasSuffix(param.Type.Name, "DestroyNotify") {

				// https://github.com/golang/go/issues/19835
				args = append(args, "(*[0]byte)(C.free)")
				continue
			}

			if closure != nil && *closure == i {
				cg.Ng.addImport("github.com/diamondburned/gotk4/internal/box")
				// TODO: this is probably not always gpointer. Handle with
				// anyCGoType.
				block.Linef("%s := C.gpointer(box.Assign(box.Callback, %s))", targ, valn)
				continue
			}

			if destroy != nil && *destroy == i {
				cg.Ng.needsCallbackDelete()
				// TODO: this is probably not always true.
				// https://github.com/golang/go/issues/19835
				args = append(args, "(*[0]byte)(C.callbackDelete)")
				continue
			}

			// Skip user_data fields because we need them for the callback
			// registry.
			if ignores.ignore(i) {
				continue
			}

			// TODO: nullability.
			// TODO: GoCConverter.

			conv := cg.Ng.CGoConverter(CGoConversion{
				Value:  valn,
				Target: targ,
				Type:   param.AnyType,
				ArgAt:  func(int) string { return "a" },
			})
			if conv == "" {
				continue
			}

			resolved, _ := cg.Ng.ResolveAnyType(param.AnyType, true)
			block.Linef("var %s %s", targ, resolved)
			block.Line(conv)
			block.EmptyLine()
			args = append(args, targ)
		}
	}

	if !returnIsVoid(cg.ReturnValue) {
		rets = append(rets, retVar{
			Name: "ret",
			Type: cg.ReturnValue.AnyType,
		})
	}

	callArgs := strings.Join(args, ", ")

	if len(rets) == 0 {
		block.Linef("C.%s(%s)", cg.CIdentifier, callArgs)
		return block.String()
	}

	block.Linef("ret := C.%s(%s)", cg.CIdentifier, callArgs)
	block.EmptyLine()

	retvars := pen.NewJoints(", ", len(rets))

	for i, ret := range rets {
		resolved, _ := cg.Ng.ResolveAnyType(ret.Type, true)
		retName := fmt.Sprintf("ret%d", i)
		retvars.Add(retName)

		block.Linef("var %s %s", retName, resolved)
		block.Line(cg.Ng.CGoConverter(CGoConversion{
			Value:  ret.Name,
			Target: retName,
			Type:   ret.Type,
			ArgAt:  argNamer,
		}))
		block.EmptyLine()
	}

	block.Line("return " + retvars.Join())
	return block.String()
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

		returns = append(returns, goName+" "+typ)
		return true
	})

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

		if param.Direction == "out" || ignores.ignore(i) {
			continue
		}

		// Ignore exposing destroy notifiers.
		if strings.HasSuffix(param.Name, "DestroyNotify") {
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

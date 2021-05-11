package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

var functionTmpl = newGoTemplate(`
	{{ $name := .GoName }}

	{{ GoDoc .Doc 0 $name }}
	func {{ $name }}({{ .Args }}) {{ .Return }}
`)

type functionGenerator struct {
	gir.Function
	Args   string
	Return string

	ng *NamespaceGenerator
}

func newFunctionGenerator(ng *NamespaceGenerator) functionGenerator {
	return functionGenerator{
		ng: ng,
	}
}

// GoName returns the current function's Go name.
func (fg *functionGenerator) GoName() string {
	return SnakeToGo(true, fg.Name)
}

// Use sets the function generator to the given GIR function.
func (fg *functionGenerator) Use(fn gir.Function) bool {
	fg.Function = fn
	fg.Args = ""
	fg.Return = ""

	var ok bool

	if fn.ReturnValue != nil {
		fg.Return, ok = fg.ng.fnReturns(fn.Parameters, fn.ReturnValue)
		if !ok {
			return false
		}
	}

	if fn.Parameters != nil {
		fg.Args, ok = fg.ng.fnArgs(fn.Parameters)
		if !ok {
			return false
		}
	}

	return true
}

// fnArgs returns the function arguments as a Go string and true. It returns
// false if the argument types cannot be fully resolved.
func (ng *NamespaceGenerator) fnArgs(params *gir.Parameters) (string, bool) {
	if params == nil || len(params.Parameters) == 0 {
		return "", true
	}

	goArgs := make([]string, 0, len(params.Parameters))

	for _, param := range params.Parameters {
		// Skip output parameters.
		if param.Direction == "out" {
			continue
		}

		resolved, ok := ng.ResolveAnyType(param.AnyType)
		if !ok {
			return "", false
		}

		goName := SnakeToGo(false, param.Name)
		goArgs = append(goArgs, goName+" "+resolved)
	}

	return strings.Join(goArgs, ", "), true
}

// fnReturns returns the function return type and true. It returns false if the
// function's return type cannot be resolved.
func (ng *NamespaceGenerator) fnReturns(ps *gir.Parameters, rs *gir.ReturnValue) (string, bool) {
	var returns []string

	if ps != nil {
		for _, param := range ps.Parameters {
			if param.Direction != "out" {
				continue
			}

			typ, ok := ng.ResolveAnyType(param.AnyType)
			if !ok {
				return "", false
			}

			// Hacky way to "dereference" a pointer once.
			if strings.HasPrefix(typ, "*") {
				typ = typ[1:]
			}

			returns = append(returns, typ)
		}
	}

	if rs != nil {
		typ, ok := ng.ResolveAnyType(rs.AnyType)
		if !ok {
			return "", false
		}

		returns = append(returns, typ)
	}

	if len(returns) == 0 {
		return "", true
	}
	if len(returns) == 1 {
		return returns[0], true
	}
	return "(" + strings.Join(returns, ", ") + ")", true
}

func (ng *NamespaceGenerator) generateFuncs() {
	fg := newFunctionGenerator(ng)

	for _, function := range ng.current.Namespace.Functions {
		if !fg.Use(function) {
			ng.logln(logInfo, "skipping function", cFunctionSig(function))
			continue
		}

		ng.pen.BlockTmpl(functionTmpl, &fg)
	}
}

// cFunctionSig renders the given GIR function in its C function signature
// string for debugging.
func cFunctionSig(fn gir.Function) string {
	b := strings.Builder{}
	b.Grow(256)

	if fn.ReturnValue != nil {
		b.WriteString(resolveAnyCType(fn.ReturnValue.AnyType))
		b.WriteByte(' ')
	}

	b.WriteString(fn.Name)
	b.WriteByte('(')

	if fn.Parameters != nil && len(fn.Parameters.Parameters) > 0 {
		for i, param := range fn.Parameters.Parameters {
			if i != 0 {
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
		pre := resolveAnyCType(any.Array.AnyType)

		if any.Array.FixedSize == 0 {
			return pre + "[]"
		}
		return pre + fmt.Sprintf("[%d]", any.Array.FixedSize)

	case any.Type != nil:
		return any.Type.CType
	default:
		return ""
	}
}

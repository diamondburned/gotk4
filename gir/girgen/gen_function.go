package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

var functionTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .Name }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
`)

type functionGenerator struct {
	callableGenerator
}

func newFunctionGenerator(ng *NamespaceGenerator) functionGenerator {
	return functionGenerator{
		callableGenerator: newCallableGenerator(ng),
	}
}

// GoName returns the current function's Go name.
func (fg *functionGenerator) GoName() string {
	return SnakeToGo(true, fg.Name)
}

// Use sets the function generator to the given GIR function.
func (fg *functionGenerator) Use(fn gir.Function) bool {
	return fg.callableGenerator.Use(fn.CallableAttrs)
}

func (ng *NamespaceGenerator) generateFuncs() {
	fg := newFunctionGenerator(ng)

	for _, function := range ng.current.Namespace.Functions {
		if ng.mustIgnore(function.Name, function.CIdentifier) {
			continue
		}

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

	case any.VarArgs != nil:
		return "..."

	default:
		return ""
	}
}

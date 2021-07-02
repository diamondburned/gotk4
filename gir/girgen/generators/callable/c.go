package callable

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

// CFunctionHeader renders the given GIR function in its C function signature
// string for debugging or callback generation.
func CFunctionHeader(fn *gir.CallableAttrs) string {
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

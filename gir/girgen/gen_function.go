package girgen

import (
	"fmt"

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
		if !function.IsIntrospectable() {
			continue
		}
		if ng.mustIgnore(function.Name, function.CIdentifier) {
			continue
		}
		if !fg.Use(function) {
			continue
		}

		fileGen := fg.fg
		fileGen.pen.WriteTmpl(functionTmpl, &fg)
	}
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

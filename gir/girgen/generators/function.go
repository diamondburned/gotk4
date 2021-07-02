package generators

import (
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
		if ng.mustIgnore(&function.Name, &function.CIdentifier) {
			continue
		}
		if !fg.Use(function) {
			continue
		}

		ng.pen.WriteTmpl(functionTmpl, &fg)
	}
}

package girgen

import "github.com/diamondburned/gotk4/gir"

var functionTmpl = newGoTemplate(`
	{{ $name := (SnakeToGo .Name) }}

	func {{ $name }}()
`)

type functionGenerator struct {
	gir.Function
	Ng *NamespaceGenerator
}

func (fg *functionGenerator) args() string {
	return fg.Ng.fnArgs(fg.Parameters)
}

func (ng *NamespaceGenerator) fnArgs(params *gir.Parameters) string {
	if params == nil || len(params.Parameters) == 0 {
		return ""
	}

	// goArgs := make([]string, 0, len(params.Parameters))

	return ""
}

func (ng *NamespaceGenerator) generateFuncs() {
	for _, function := range ng.current.Namespace.Functions {
		ng.pen.BlockTmpl(functionTmpl, functionGenerator{
			Function: function,
			Ng:       ng,
		})
	}
}

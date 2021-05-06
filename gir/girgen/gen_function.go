package girgen

import (
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
	Ng *NamespaceGenerator
}

func (fg functionGenerator) GoName() string {
	return SnakeToGo(true, fg.Name)
}

func (fg functionGenerator) Args() string {
	return fg.Ng.fnArgs(fg.Parameters)
}

func (fg functionGenerator) Return() string {
	return fg.Ng.fnReturns(fg.ReturnValue)
}

func (ng *NamespaceGenerator) fnArgs(params *gir.Parameters) string {
	if params == nil || len(params.Parameters) == 0 {
		return ""
	}

	goArgs := make([]string, 0, len(params.Parameters))

	for _, param := range params.Parameters {
		resolved := ng.resolveType(param.Type)
		if resolved == nil {
			continue
		}

		goName := SnakeToGo(false, param.Name)
		goArgs = append(goArgs, goName+" "+resolved.GoType)
	}

	return strings.Join(goArgs, ", ")
}

func (ng *NamespaceGenerator) fnReturns(rets *gir.ReturnValue) string {
	if rets == nil {
		return ""
	}

	// TODO: arrays
	return ng.resolveAnyType(rets.AnyType)
}

func (ng *NamespaceGenerator) generateFuncs() {
	for _, function := range ng.current.Namespace.Functions {
		ng.pen.BlockTmpl(functionTmpl, functionGenerator{
			Function: function,
			Ng:       ng,
		})
	}
}

package generators

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var functionTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
`)

// GenerateFunction generates the function call for the given GIR function.
func GenerateFunction(gen FileGeneratorWriter, fn *gir.Function) bool {
	if !fn.IsIntrospectable() || types.Filter(gen, fn.Name, fn.CIdentifier) {
		return false
	}

	callableGen := callable.NewGenerator(gen)
	if !callableGen.Use(&fn.CallableAttrs) {
		return false
	}

	writer := FileWriterFromType(gen, fn)
	writer.Pen().WriteTmpl(functionTmpl, &callableGen)
	file.ApplyHeader(writer, &callableGen)
	return true
}

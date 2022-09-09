package generators

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callback"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
)

var callbackTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	type {{ .GoName }} func{{ .GoTail }}
`)

var callbackExportTmpl = gotmpl.NewGoTemplate(`
	//export {{ .CGoName }}
	func {{ .CGoName }}{{ .CGoTail }} {{ .Block }}
`)

// GenerateCallback generates a callback type declaration and handler into the
// given file generator.
func GenerateCallback(gen FileGeneratorWriter, cb *gir.Callback) bool {
	generator := callback.NewGenerator(gen)
	generator.Parent = cb
	if !generator.Use(&cb.CallableAttrs) {
		return false
	}

	{
		writer := FileWriterFromType(gen, cb)
		writer.Pen().WriteTmpl(callbackTmpl, &generator)
		for _, result := range generator.Results {
			result.Resolved.ImportPubl(gen, writer.Header())
		}
	}

	{
		writer := FileWriterExportedFromType(gen, cb)
		writer.Pen().WriteTmpl(callbackExportTmpl, &generator)
		file.ApplyHeader(writer, &generator)
	}

	return true
}

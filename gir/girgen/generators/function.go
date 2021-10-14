package generators

import (
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var functionTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	{{- if .ParamDocs }}
	//
	// The function takes the following parameters:
	//
	{{- range .ParamDocs }}
	{{ GoDoc . 0 (OverrideSelfName .Name) (AdditionalPrefix "- ") (ParagraphIndent 1) }}
	{{- end }}
	//
	{{- end }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
`)

// GenerateFunction generates the function call for the given GIR function.
func GenerateFunction(gen FileGeneratorWriter, fn *gir.Function) bool {
	return GeneratePrefixedFunction(gen, fn, "")
}

// GeneratePrefixedFunction generates the given GIR function with the prefix
// prepended into the name.
func GeneratePrefixedFunction(gen FileGeneratorWriter, fn *gir.Function, prefix string) bool {
	if fn.CIdentifier == "" || types.Filter(gen, fn.Name, fn.CIdentifier) {
		return false
	}

	typ := gir.TypeFindResult{
		NamespaceFindResult: gen.Namespace(),
		Type:                fn,
	}

	callableGen := callable.NewGenerator(gen)
	if !callableGen.Use(&typ, &fn.CallableAttrs) {
		return false
	}

	if prefix != "" {
		prefix = strcases.Go(prefix)

		// Check if this function is actually a constructor.
		if strings.HasPrefix(callableGen.Name, "New") {
			callableGen.Name = strings.TrimPrefix(callableGen.Name, "New")
			callableGen.Name = "New" + prefix + callableGen.Name
		} else {
			callableGen.Name = prefix + callableGen.Name
		}
	}

	callableGen.CoalesceTail()

	writer := FileWriterFromType(gen, fn)
	writer.Pen().WriteTmpl(functionTmpl, &callableGen)
	file.ApplyHeader(writer, &callableGen)
	return true
}

package girgen

import "github.com/diamondburned/gotk4/gir"

var initTmpl = newGoTemplate(`
	func init() {
		externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
			{{ if .Enums -}}
			// Enums
			{{- range .Enums }}
			{{- if .GLibGetType }}
			{T: externglib.Type(C.{{.GLibGetType}}()), F: marshal{{.Name}}},
			{{- else }}
			// Skipped {{.Name}}.
			{{- end -}}
			{{ end }}
			{{ end }}

			// Objects/Classes
		})
	}
`)

type initGenerator struct {
	gir.Namespace
	Ng *NamespaceGenerator
}

func (ng *NamespaceGenerator) generateInit() {
	ng.addImportAlias("github.com/gotk3/gotk3/glib", "externglib")
	ng.pen.BlockTmpl(initTmpl, initGenerator{
		Namespace: *ng.current.Namespace,
		Ng:        ng,
	})
}

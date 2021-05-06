package girgen

import "github.com/diamondburned/gotk4/gir"

var initTmpl = newGoTemplate(`
	func init() {
		glib.RegisterGValueMarshalers([]glib.TypeMarshaler{
			{{ if .Enums -}}
			// Enums
			{{- range .Enums }}
			{{- if .GLibGetType }}
			{T: glib.Type(C.{{.GLibGetType}}()), F: marshal{{.Name}}},
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
	ng.pen.BlockTmpl(initTmpl, initGenerator{
		Namespace: ng.current.Namespace,
		Ng:        ng,
	})
}

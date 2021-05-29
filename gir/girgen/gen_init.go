package girgen

import "github.com/diamondburned/gotk4/gir"

var initTmpl = newGoTemplate(`
	func init() {
		externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
			{{ if .Enums -}}
			// Enums
			{{- range .Enums }}
			{{- if .GLibGetType }}
			{T: externglib.Type(C.{{.GLibGetType}}()), F: marshal{{PascalToGo .Name}}},
			{{- else }}
			// Skipped {{.Name}}.
			{{- end -}}
			{{ end }}
			{{ end }}

			{{ if .Records -}}
			// Records
			{{- range .Records }}
			{{- if (and .GLibGetType ($.R.Use .)) }}
			{T: externglib.Type(C.{{.GLibGetType}}()), F: marshal{{$.R.GoName}}},
			{{- else }}
			// Skipped {{.Name}}.
			{{- end -}}
			{{ end }}
			{{ end }}

			{{ if .Classes -}}
			// Classes
			{{- range .Classes }}
			{{- if (and .GLibGetType ($.C.Use .)) }}
			{T: externglib.Type(C.{{.GLibGetType}}()), F: marshal{{$.C.InterfaceName}}},
			{{- else }}
			// Skipped {{.Name}}.
			{{- end -}}
			{{ end }}
			{{ end }}
		})
	}
`)

type initGenerator struct {
	*gir.Namespace
	C  *classGenerator
	R  *recordGenerator
	Ng *NamespaceGenerator
}

func (ng *NamespaceGenerator) generateInit() {
	ng.addImportAlias("github.com/gotk3/gotk3/glib", "externglib")

	ng.pen.BlockTmpl(initTmpl, initGenerator{
		Namespace: ng.current.Namespace,

		Ng: ng,

		C: newClassGenerator(ng),
		R: newRecordGenerator(ng),
	})
}

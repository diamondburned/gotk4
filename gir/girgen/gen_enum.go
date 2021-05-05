package girgen

import (
	"github.com/diamondburned/gotk4/gir"
)

var enumTmpl = newGoTemplate(`
	type {{ .Name }} int

	const (
		{{ range .Members }}
		{{- $name := ($.FormatMember .Name) -}}
		{{- GoDoc .Doc 1 $name -}}
		{{ $name }} {{ $.Name }} = {{ .Value }}
		{{ end }})
`)

type enumGenerator struct {
	gir.Enum
	Ng *NamespaceGenerator
}

func (eg *enumGenerator) FormatMember(memberName string) string {
	return eg.Name + SnakeToGo(true, memberName)
}

func (ng *NamespaceGenerator) generateEnums() {
	for _, enum := range ng.current.Namespace.Enums {
		ng.pen.BlockTmpl(enumTmpl, &enumGenerator{
			Enum: enum,
			Ng:   ng,
		})
	}
}

package girgen

import (
	"github.com/diamondburned/gotk4/gir"
)

var enumTmpl = newGoTemplate(`
	{{ $type := (PascalToGo .Name) }}

	{{ GoDoc .Doc 0 $type }}
	type {{ $type }} int

	const (
		{{ range .Members -}}
		{{- $name := ($.FormatMember .Name) -}}
		{{ with (GoDoc .Doc 1 $name) }} {{- . -}} {{ end }}
		{{ $name }} {{ $type }} = {{ .Value }}
		{{ end -}}
	)

	{{ if .GLibGetType }}
	func marshal{{ $type }}(p uintptr) (interface{}, error) {
		return {{ $type }}(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
	}
	{{ end }}
`)

type enumGenerator struct {
	gir.Enum
	Ng *NamespaceGenerator
}

func (eg *enumGenerator) FormatMember(memberName string) string {
	return PascalToGo(eg.Name) + SnakeToGo(true, memberName)
}

func (ng *NamespaceGenerator) generateEnums() {
	for _, enum := range ng.current.Namespace.Enums {
		if ng.mustIgnore(enum.Name, enum.CType) {
			continue
		}

		fg := ng.FileFromSource(enum.SourcePosition)
		fg.needsGLibObject()

		if enum.GLibGetType != "" && !ng.mustIgnoreC(enum.GLibGetType) {
			fg.addMarshaler(enum.GLibGetType, PascalToGo(enum.Name))
		}

		fg.pen.BlockTmpl(enumTmpl, &enumGenerator{
			Enum: enum,
			Ng:   ng,
		})
	}
}

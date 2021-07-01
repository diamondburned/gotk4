package girgen

import (
	"strconv"

	"github.com/diamondburned/gotk4/gir"
)

var bitfieldTmpl = newGoTemplate(`
	{{ $type := (PascalToGo .Name) }}

	{{ GoDoc .Doc 0 $type }}
	type {{ $type }} int

	const (
		{{ range .Members -}}
		{{- $name := ($.FormatMember .Name) -}}
		{{- if .Doc -}}
		{{ GoDoc .Doc 1 $name }}
		{{ end -}}
		{{ $name }} {{ $type }} = {{ .Value }}
		{{ end -}}
	)

	{{ if .GLibGetType }}
	func marshal{{ $type }}(p uintptr) (interface{}, error) {
		return {{ $type }}(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
	}
	{{ end }}
`)

type bitfieldGenerator struct {
	gir.Bitfield
	Ng *NamespaceGenerator
}

func (eg *bitfieldGenerator) Bits(v string) string {
	b, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return v
	}

	return "0b" + strconv.FormatUint(b, 2)
}

func (eg *bitfieldGenerator) FormatMember(memberName string) string {
	return PascalToGo(eg.Name) + SnakeToGo(true, memberName)
}

func (ng *NamespaceGenerator) generateBitfields() {
	for _, bitfield := range ng.current.Namespace.Bitfields {
		if !bitfield.IsIntrospectable() {
			continue
		}
		if ng.mustIgnore(&bitfield.Name, &bitfield.CType) {
			continue
		}

		if bitfield.GLibGetType != "" && !ng.mustIgnoreC(bitfield.GLibGetType) {
			ng.addMarshaler(bitfield.GLibGetType, PascalToGo(bitfield.Name))
		}

		ng.pen.WriteTmpl(bitfieldTmpl, &bitfieldGenerator{
			Bitfield: bitfield,
			Ng:       ng,
		})
	}
}

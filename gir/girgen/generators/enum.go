package generators

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var enumTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	type {{ .GoName }} int

	const (
		{{ range .Members -}}
		{{- $name := ($.FormatMember .Name) -}}
		{{- if .Doc -}}
		{{ GoDoc . 1 }}
		{{ end -}}
		{{ $name }} {{ $.GoName }} = {{ .Value }}
		{{ end -}}
	)

	{{ if .GLibGetType }}
	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		return {{ .GoName }}(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
	}
	{{ end }}
`)

type enumData struct {
	*gir.Enum
	GoName string
}

func (eg *enumData) FormatMember(memberName string) string {
	return strcases.PascalToGo(eg.Name) + strcases.SnakeToGo(true, memberName)
}

// GenerateEnum generates an enum type declaration as well as the constants and
// the type marshaler.
func GenerateEnum(gen FileGeneratorWriter, enum *gir.Enum) bool {
	if !enum.IsIntrospectable() || types.Filter(gen, enum.Name, enum.CType) {
		return false
	}

	goName := strcases.PascalToGo(enum.Name)
	writer := FileWriterFromType(gen, enum)

	if enum.GLibGetType != "" && !types.FilterCType(gen, enum.GLibGetType) {
		writer.Header().AddMarshaler(enum.GLibGetType, goName)
	}

	writer.Pen().WriteTmpl(enumTmpl, &enumData{
		Enum:   enum,
		GoName: goName,
	})
	return true
}

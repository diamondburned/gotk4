package generators

import (
	"strconv"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var enumTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	type {{ .GoName }} int

	{{ if .IsIota }}
	const (
		{{ range $ix, $member := .Members -}}
		{{- if .Doc -}}
		{{ GoDoc . 1 }}
		{{ end -}}
		{{- if (eq $ix 0) -}}
		{{ $.FormatMember .Name }} {{ $.GoName }} = iota
		{{ else -}}
		{{ $.FormatMember .Name }}
		{{ end -}}
		{{ end }}
	)
	{{ else }}
	const (
		{{ range .Members -}}
		{{- if .Doc -}}
		{{ GoDoc . 1 }}
		{{ end -}}
		{{ $.FormatMember .Name }} {{ $.GoName }} = {{ .Value }}
		{{ end -}}
	)
	{{ end }}

	{{ if .GLibGetType }}
	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		return {{ .GoName }}(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
	}
	{{ end }}
`)

type enumData struct {
	*gir.Enum
	GoName string
	IsIota bool
}

func (eg *enumData) FormatMember(memberName string) string {
	return strcases.PascalToGo(eg.Name) + strcases.SnakeToGo(true, memberName)
}

// CanGenerateEnum returns false if the given enum cannot be generated.
func CanGenerateEnum(gen FileGenerator, enum *gir.Enum) bool {
	if !enum.IsIntrospectable() || types.Filter(gen, enum.Name, enum.CType) {
		return false
	}
	return true
}

// GenerateEnum generates an enum type declaration as well as the constants and
// the type marshaler.
func GenerateEnum(gen FileGeneratorWriter, enum *gir.Enum) bool {
	if !CanGenerateEnum(gen, enum) {
		return false
	}

	goName := strcases.PascalToGo(enum.Name)
	writer := FileWriterFromType(gen, enum)

	if enum.GLibGetType != "" && !types.FilterCType(gen, enum.GLibGetType) {
		writer.Header().AddMarshaler(enum.GLibGetType, goName)
	}

	isIota := true
	for i := 0; i < len(enum.Members); i++ {
		if enum.Members[i].Value != strconv.Itoa(i) {
			isIota = false
			break
		}
	}

	writer.Pen().WriteTmpl(enumTmpl, &enumData{
		Enum:   enum,
		GoName: goName,
		IsIota: isIota,
	})
	return true
}

package generators

import (
	"strconv"
	"unicode"
	"unicode/utf8"

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
		{{ range $ix, $member := .Members }}
		{{- GoDoc . 1 TrailingNewLine -}}
		{{- if (eq $ix 0) }}
		{{- $.FormatMember .Name }} {{ $.GoName }} = iota
		{{- else }}
		{{- $.FormatMember .Name }}
		{{- end }}
		{{ end }}
	)
	{{ else }}
	const (
		{{ range .Members }}
		{{- GoDoc . 1 TrailingNewLine -}}
		{{- $.FormatMember .Name }} {{ $.GoName }} = {{ .Value }}
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

var numberMap = map[rune]string{
	'0': "Zero",
	'1': "One",
	'2': "Two",
	'3': "Three",
	'4': "Four",
	'5': "Five",
	'6': "Six",
	'7': "Seven",
	'8': "Eight",
	'9': "Nine",
}

func (eg *enumData) FormatMember(memberName string) string {
	memberName = strcases.SnakeToGo(true, memberName)

	r, sz := utf8.DecodeRuneInString(memberName)
	if sz > 0 && unicode.IsNumber(r) {
		memberName = numberMap[r] + memberName[sz:]
	}

	return memberName
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

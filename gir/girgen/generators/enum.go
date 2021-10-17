package generators

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var enumTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 (OverrideSelfName .GoName) }}
	type {{ .GoName }} int

	{{ if .IsIota }}
	const (
		{{ range $ix, $member := .Members }}
		{{- $name := $.FormatMember $member -}}
		{{- GoDoc . 1 TrailingNewLine (OverrideSelfName $name) -}}
		{{- if (eq $ix 0) }}
		{{- $name }} {{ $.GoName }} = iota
		{{- else }}
		{{- $name }}
		{{- end }}
		{{ end }}
	)
	{{ else }}
	const (
		{{ range .Members -}}
		{{- $name := $.FormatMember . -}}
		{{- GoDoc . 1 TrailingNewLine (OverrideSelfName $name) -}}
		{{- $name }} {{ $.GoName }} = {{ .Value }}
		{{ end -}}
	)
	{{ end }}

	{{ if .Marshaler }}
	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		return {{ .GoName }}(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
	}
	{{ end }}

	{{ $recv := FirstLetter .GoName }}
	// String returns the name in string for {{ .GoName }}.
	func ({{ $recv }} {{ .GoName }}) String() string {
		switch {{ $recv }} {
		{{- range .UniqueMembers }} {{ $name := $.FormatMember . }}
		case {{ $name }}: return "{{ SnakeToGo true .Name }}"
		{{- end }}
		default: return fmt.Sprintf("{{ .GoName }}(%d)", {{ $recv }})
		}
	}
`)

type enumData struct {
	*gir.Enum
	GoName    string
	IsIota    bool
	Marshaler bool
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

func (eg *enumData) FormatMember(member gir.Member) string {
	return FormatEnumMember(member)
}

func (eg *enumData) UniqueMembers() []gir.Member {
	return UniqueEnumMembers(eg.Members)
}

// FormatEnumMember returns the enum member's Go name.
func FormatEnumMember(member gir.Member) string {
	// Pop the namespace off. Probably works most of the time.
	if parts := strings.SplitN(member.CIdentifier, "_", 2); len(parts) == 2 {
		member.CIdentifier = parts[1]
	}

	memberName := strcases.SnakeToGo(true, strings.ToLower(member.CIdentifier))

	// TODO: prepend GoName instead.
	r, sz := utf8.DecodeRuneInString(memberName)
	if sz > 0 && unicode.IsNumber(r) {
		memberName = numberMap[r] + memberName[sz:]
	}

	return memberName
}

// UniqueEnumMembers returns the enum members with unique values only.
func UniqueEnumMembers(members []gir.Member) []gir.Member {
	uniques := make([]gir.Member, 0, len(members))
	known := make(map[string]struct{}, len(members))

	for _, member := range members {
		_, isKnown := known[member.Value]
		if isKnown {
			continue
		}

		uniques = append(uniques, member)
		known[member.Value] = struct{}{}
	}

	return uniques
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

	data := enumData{
		Enum:   enum,
		GoName: goName,
		IsIota: true,
	}

	if enum.GLibGetType != "" && !types.FilterCType(gen, enum.GLibGetType) {
		data.Marshaler = true
		writer.Header().NeedsExternGLib()
		writer.Header().AddMarshaler(enum.GLibGetType, goName)
	}

	for i := 0; i < len(enum.Members); i++ {
		if enum.Members[i].Value != strconv.Itoa(i) {
			data.IsIota = false
			break
		}
	}

	writer.Header().Import("fmt")
	writer.Pen().WriteTmpl(enumTmpl, &data)
	return true
}

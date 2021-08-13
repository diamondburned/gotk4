package generators

import (
	"strconv"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var bitfieldTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 (OverrideSelfName .GoName) }}
	type {{ .GoName }} int

	const (
		{{ range .Members -}}
		{{- $name := ($.FormatMember .) -}}
		{{- GoDoc . 1 TrailingNewLine (OverrideSelfName $name) -}}
		{{- $name }} {{ $.GoName }} = {{ $.Bits .Value }}
		{{ end -}}
	)

	{{ if .GLibGetType }}
	func marshal{{.GoName}}(p uintptr) (interface{}, error) {
		return {{.GoName}}(C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))), nil
	}
	{{ end }}

	{{ $recv := FirstLetter .GoName }}
	// String returns the names in string for {{.GoName}}.
	func ({{$recv}} {{.GoName}}) String() string {
		if {{$recv}} == 0 {
			return "{{.GoName}}(0)"
		}

		var builder strings.Builder
		builder.Grow({{.StrLen}})

		for {{$recv}} != 0 {
			next := {{$recv}} & ({{$recv}} - 1)
			bit := {{$recv}} - next

			switch bit {
			{{- range .UniqueMembers }} {{ $name := $.FormatMember . }}
			case {{$name}}: builder.WriteString("{{SnakeToGo true .Name}}|")
			{{- end }}
			default: builder.WriteString(fmt.Sprintf("{{.GoName}}(0b%b)|", bit))
			}

			{{$recv}} = next
		}

		return strings.TrimSuffix(builder.String(), "|")
	}

	// Has returns true if {{$recv}} contains other.
	func ({{$recv}} {{.GoName}}) Has(other {{.GoName}}) bool {
		return ({{$recv}} & other) == other
	}
`)

type bitfieldData struct {
	*gir.Bitfield
	GoName string
	StrLen int // length of all enum strings concatenated

	gen FileGenerator
}

func (*bitfieldData) Bits(v string) string {
	b, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return v
	}

	return "0b" + strconv.FormatUint(b, 2)
}

func (b *bitfieldData) FormatMember(member gir.Member) string {
	return FormatEnumMember(member)
}

func (b *bitfieldData) UniqueMembers() []gir.Member {
	return UniqueEnumMembers(b.Members)
}

// CanGenerateBitfield returns false if the bitfield cannot be generated.
func CanGenerateBitfield(gen FileGenerator, bitfield *gir.Bitfield) bool {
	if !bitfield.IsIntrospectable() || types.Filter(gen, bitfield.Name, bitfield.CType) {
		return false
	}
	return true
}

// GenerateBitfield generates a bitfield type declaration as well as the
// constants and the type marshaler into the given file generator. If the
// generation fails or is ignored, then false is returned.
func GenerateBitfield(gen FileGeneratorWriter, bitfield *gir.Bitfield) bool {
	if !CanGenerateBitfield(gen, bitfield) {
		return false
	}

	goName := strcases.PascalToGo(bitfield.Name)
	writer := FileWriterFromType(gen, bitfield)

	if bitfield.GLibGetType != "" && !types.FilterCType(gen, bitfield.GLibGetType) {
		writer.Header().AddMarshaler(bitfield.GLibGetType, goName)
	}

	// Need GLibObject for g_value_*.
	writer.Header().NeedsGLibObject()
	// Need this for String().
	writer.Header().Import("strings")
	writer.Header().Import("fmt")

	data := &bitfieldData{
		Bitfield: bitfield,
		GoName:   goName,
		gen:      gen,
	}

	for i, member := range bitfield.Members {
		data.StrLen += len(data.FormatMember(member))
		if i > 0 {
			data.StrLen++ // account for '|'
		}
	}

	writer.Pen().WriteTmpl(bitfieldTmpl, data)

	return true
}

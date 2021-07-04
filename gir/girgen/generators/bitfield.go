package generators

import (
	"strconv"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var bitfieldTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	type {{ .GoName }} int

	const (
		{{ range .Members -}}
		{{- $name := ($.FormatMember .Name) -}}
		{{- if .Doc -}}
		{{ GoDoc . 1 (OverrideSelfName $name) }}
		{{ end -}}
		{{ $name }} {{ $.GoName }} = {{ $.Bits .Value }}
		{{ end -}}
	)

	{{ if .GLibGetType }}
	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		return {{ .GoName }}(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
	}
	{{ end }}
`)

type bitfieldData struct {
	*gir.Bitfield
	GoName string

	gen FileGenerator
}

func (*bitfieldData) Bits(v string) string {
	b, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return v
	}

	return "0b" + strconv.FormatUint(b, 2)
}

func (b *bitfieldData) FormatMember(memberName string) string {
	return strcases.PascalToGo(b.Name) + strcases.SnakeToGo(true, memberName)
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

	writer.Pen().WriteTmpl(bitfieldTmpl, &bitfieldData{
		Bitfield: bitfield,
		GoName:   goName,
		gen:      gen,
	})

	return true
}

package generators

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

var unionTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	type {{ .GoName }} struct {
		*{{ .ImplName }}
	}

	// {{ .ImplName }} has the finalizer attached to it.
	type {{ .ImplName }} struct {
		native *C.{{ .CType }}
	}

	{{ $iface := Interfacify .GoName }}

	// {{ $iface }} is used for all functions that accept any kind of {{ .GoName }}.
	type {{ $iface }} interface {
		// Implementing types:
		//
		{{- range .Fields }}
		{{- if .Record }}
		//    {{ .GoType }}
		{{- end }}
		{{- end }}
		//

		underlying{{ .GoName }}() unsafe.Pointer
	}

	// Copy{{ $iface }} copies any type that belongs to a {{ .GoName }} union
	// into a new {{ .GoName }} instance. To see supported types, refer to
	// {{ $iface }}'s documentation.
	func Copy{{ $iface }}({{ .Recv }} {{ $iface }}) *{{ .GoName }} {{ .CastBlock }}

	{{ if .Marshaler }}
	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
		return &{{.GoName}}{&{{.ImplName}}{(*C.{{.CType}})(b)}}, nil
	}
	{{ end }}

	func (v *{{ .GoName }}) underlying{{ .GoName }}() unsafe.Pointer {
		return unsafe.Pointer(v.native)
	}

	{{ range .Fields }}
	{{ if .Record }}
	// underlying{{ $.GoName }} marks the struct for {{ $iface }}.
	func (v {{ .GoType }}) underlying{{ $.GoName }}() unsafe.Pointer {
		return unsafe.Pointer(v.native)
	}
	{{ end }}
	// As{{ .GoName }} returns a copy of {{ $.Recv }} as the struct {{ .GoType }}.
	// It does this without any knowledge on the actual type of the value, so
	// the caller must take care of type-checking beforehand.
	func ({{ $.Recv }} *{{ $.GoName }}) As{{ .GoName }}() {{ .GoType }} {{ .Block }}
	{{ end }}
`)

type UnionGenerator struct {
	*gir.Union
	GoName    string
	ImplName  string
	Marshaler bool

	CastBlock string

	Marshalers []UnionFieldMarshaler
	Fields     []UnionField

	gen FileGenerator
	hdr file.Header
}

type UnionFieldMarshaler struct {
	*gir.Field
	GLibGetType string
}

type UnionField struct {
	*gir.Field
	GoName string
	GoType string
	Block  string
	Record bool
}

// GenerateUnion generates the union.
func GenerateUnion(gen FileGeneratorWriter, union *gir.Union) bool {
	unionGen := NewUnionGenerator(gen)
	if !unionGen.Use(union) {
		return false
	}

	writer := FileWriterFromType(gen, union)

	if len(unionGen.Fields) > 0 {
		writer.Header().Import("unsafe")
	}

	if gtype, ok := GenerateGType(gen, union.Name, union.GLibGetType); ok {
		unionGen.Marshaler = true
		writer.Header().AddMarshaler(gtype.GetType, unionGen.GoName)
	}

	writer.Pen().WriteTmpl(unionTmpl, &unionGen)
	// Write the header after using the template to ensure that UseConstructor
	// registers everything.
	file.ApplyHeader(writer, &unionGen)

	return true
}

func NewUnionGenerator(gen FileGenerator) UnionGenerator {
	return UnionGenerator{
		gen: gen,
	}
}

// Header returns the union generator's headers.
func (ug *UnionGenerator) Header() *file.Header {
	return &ug.hdr
}

// Recv returns the method receiver.
func (ug *UnionGenerator) Recv() string {
	return strcases.FirstLetter(ug.GoName)
}

func (ug *UnionGenerator) Use(union *gir.Union) bool {
	ug.hdr.Reset()
	ug.Marshaler = false
	ug.Fields = ug.Fields[:0]

	// If there's no C type or the name has an underscore, then we probably
	// shouldn't touch it.
	if union.CType == "" || strings.HasPrefix(union.Name, "_") {
		return false
	}

	ug.Union = union
	ug.GoName = strcases.PascalToGo(union.Name)
	ug.ImplName = strcases.UnexportPascal(ug.GoName)

	typ := gir.TypeFindResult{
		NamespaceFindResult: ug.gen.Namespace(),
		Type:                union,
	}

	typRecord := gir.TypeFindResult{
		NamespaceFindResult: typ.NamespaceFindResult,
		Type: &gir.Record{
			Name:    union.Name,
			Methods: union.Methods,
		},
	}

	// Can we copy? Exit if not. We want the finalizer to work properly.
	copyMethod := types.FindMethodName(union.Methods, "copy")
	if copyMethod == nil {
		ug.Logln(logger.Skip, "no copy method")
		return false
	}

	// We can optionally have a freeMethod.
	// freeMethod := types.RecordHasFree(&gir.Record{Methods: union.Methods})

	{
		ug.hdr.Import("unsafe")
		ug.hdr.ImportCore("gextras")

		p := pen.NewBlock()
		p.Linef("original := (*C.%s)(%s.underlying%s())", union.CType, ug.Recv(), ug.GoName)
		p.Linef("copied := C.%s(original)", copyMethod.CIdentifier)
		p.Linef("dst := (*%s)(gextras.NewStructNative(unsafe.Pointer(copied)))", ug.GoName)
		p.Linef("runtime.SetFinalizer(")
		p.Linef("  gextras.StructIntern(unsafe.Pointer(dst)),")
		p.Linef("  func(intern *struct{ C unsafe.Pointer }) {")
		p.Linef(types.RecordPrintFree(ug.gen, &typRecord, "intern.C"))
		p.Linef("},")
		p.Linef(")")
		p.Linef("return dst")

		ug.CastBlock = p.String()
	}

	for i, field := range union.Fields {
		if field.Type == nil || !field.IsReadable() {
			continue
		}

		srcVal := typeconv.ConversionValue{
			InName:      "cpy",
			OutName:     "dst",
			Direction:   typeconv.ConvertCToGo,
			InContainer: true,
			// ParameterIndex = 0 lets the type converter treat the field value
			// as a pointer rather than a field value.
			ParameterIndex: 0,
			ParameterAttrs: gir.ParameterAttrs{
				Name:    field.Name,
				AnyType: field.AnyType,
				TransferOwnership: gir.TransferOwnership{
					// Full prevents the type converter from adding a finalizer.
					// We'll write our own.
					TransferOwnership: "full",
				},
			},
		}

		converter := typeconv.NewConverter(ug.gen, &typ, []typeconv.ConversionValue{srcVal})
		converter.UseLogger(ug)

		srcRes := converter.Convert(0)
		if srcRes == nil {
			ug.Logln(logger.Skip, "field", field.Name)
			continue
		}

		// Only allow records or enums or bitfields for now.
		if !srcRes.Resolved.IsRecord() && !srcRes.Resolved.IsEnumOrBitfield() {
			ug.Logln(logger.Skip, "field", field.Name, "is unsupported type")
			continue
		}

		p := pen.NewBlock()
		ug.hdr.Import("runtime")

		// We only need to copy if this is a record, because a record is passed
		// by reference. Since enums/bitfields are copied, we don't need to copy
		// the original value.
		if srcRes.Resolved.IsRecord() {
			p.Linef(
				// The type conversion helps us ensure that copy() actually returns
				// the type that we expect it to. Otherwise, unsafe.Pointer is
				// potentialy dangerous.
				"cpy := (*C.%s)(C.%s(%s.%s.native))",
				ug.CType, copyMethod.CIdentifier, ug.Recv(), ug.ImplName,
			)
		} else {
			p.Linef("cpy := %s.%s.native", ug.Recv(), ug.ImplName)
		}

		p.Linef("var dst %s", srcRes.Out.Type)
		p.Linef(srcRes.Conversion)

		// We should free the copy when we're done if this is a record, since we
		// copied it earlier.
		if srcRes.Resolved.IsRecord() {
			ug.hdr.Import("unsafe")
			ug.hdr.ImportCore("gextras")

			p.Linef("runtime.SetFinalizer(")
			// dst is ASSUMED TO BE A POINTER.
			p.Linef("  gextras.StructIntern(unsafe.Pointer(dst)),")
			p.Linef("  func(intern *struct{ C unsafe.Pointer }) {")
			p.Linef(types.RecordPrintFree(ug.gen, &typRecord, "intern.C"))
			p.Linef("},")
			p.Linef(")")
		}

		p.Linef("runtime.KeepAlive(%s.%s)", ug.Recv(), ug.ImplName)
		p.Linef("return dst")

		ug.Fields = append(ug.Fields, UnionField{
			Field:  &union.Fields[i],
			GoName: strcases.SnakeToGo(true, field.Name),
			GoType: srcRes.Out.Type,
			Block:  p.String(),
			Record: srcRes.Resolved.IsRecord(),
		})
	}

	return true
}

func (ug *UnionGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("union %s (C.%s):", ug.GoName, ug.CType)
	ug.gen.Logln(lvl, logger.Prefix(v, p)...)
}

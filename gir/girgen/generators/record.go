package generators

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

// recordIgnoreSuffixes is a list of suffixes that structs must not have,
// otherwise they are skipped. This is mostly because these types shouldn't
// be implemented in Go like they're described.
var recordIgnoreSuffixes = []string{
	// TODO: these interfaces shouldn't actually be ignored. Instead, they
	// should contain callback functions, because it seems like they're meant
	// to be called by GTK.
	//
	// It seems like GCallback (GObject.Callback) can be interface{}, since GLib
	// can do lazy type cast on call.
	"Interface",
	"Iface",
	// No idea why these are here.
	"Class",
	"Private",
}

var recordTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	type {{ .GoName }} struct {
		nocopy gextras.NoCopy
		native *C.{{.CType}}
	}

	{{ if .GLibGetType }}
	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
		return &{{ .GoName }}{native: (*C.{{.CType}})(unsafe.Pointer(b))}, nil
	}
	{{ end }}

	{{ range .Constructors }}
	{{ if $.UseConstructor . $.GoName }}
	// {{ $.Callable.Name }} constructs a struct {{ $.GoName }}.
	func {{ $.Callable.Name }}{{ $.Callable.Tail }} {{ $.Callable.Block }}
	{{ end }}
	{{ end }}

	{{ $recv := (FirstLetter $.GoName) }}

	{{ range .Getters }}
	{{ GoDoc . 0 }}
	func ({{ $recv }} *{{ $.GoName }}) {{ .Name }}() {{ .Type }} {{ .Block }}
	{{ end }}

	{{ range .Methods }}
	{{ GoDoc . 0 }}
	func ({{ .Recv }} *{{ $.GoName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

// CanGenerateRecord returns false if the record cannot be generated.
func CanGenerateRecord(gen FileGenerator, rec *gir.Record) bool {
	log := func(v ...interface{}) {
		p := fmt.Sprintf("record %s (C.%s)", rec.Name, rec.CType)
		gen.Logln(logger.Debug, logger.Prefix(v, p)...)
	}

	if !rec.IsIntrospectable() {
		log("not introspectable")
		return false
	}

	if rec.Disguised {
		log("disguised")
		return false
	}

	// GLibIsGTypeStructFor seems to be records used in addition to classes due
	// to C? Not sure, but we likely don't need it.
	if rec.GLibIsGTypeStructFor != "" || strings.HasPrefix(rec.Name, "_") {
		log("IsGTypeStructFor or has underscore prefixed")
		return false
	}

	for _, suffix := range recordIgnoreSuffixes {
		if strings.HasSuffix(rec.Name, suffix) {
			log("contains forbidden suffix", suffix)
			return false
		}
	}

	// Ignore non-type/array fields.
	for _, field := range rec.Fields {
		if ignoreField(&field) {
			continue
		}
	}

	if types.Filter(gen, rec.Name, rec.CType) {
		log("filtered out")
		return false
	}

	return true
}

// mustIgnoreAny banished here because it disregards type renamers.
func mustIgnoreAny(gen FileGenerator, any gir.AnyType) bool {
	switch {
	case any.Type != nil:
		if types.Filter(gen, any.Type.Name, any.Type.CType) {
			return true
		}

		for _, inner := range any.Type.Types {
			if types.Filter(gen, inner.Name, inner.CType) {
				return true
			}
		}

		return false
	case any.Array != nil:
		return mustIgnoreAny(gen, gir.AnyType{Type: any.Array.Type})
	default:
		return true
	}
}

// GenerateRecord generates the records.
func GenerateRecord(gen FileGeneratorWriter, record *gir.Record) bool {
	recordGen := NewRecordGenerator(gen)
	if !recordGen.Use(record) {
		return false
	}

	writer := FileWriterFromType(gen, record)
	writer.Header().ImportCore("gextras")

	if record.GLibGetType != "" && !types.FilterCType(gen, record.GLibGetType) {
		writer.Header().AddMarshaler(record.GLibGetType, recordGen.GoName)
	}

	writer.Pen().WriteTmpl(recordTmpl, &recordGen)
	// Write the header after using the template to ensure that UseConstructor
	// registers everything.
	file.ApplyHeader(writer, &recordGen)
	return true
}

type RecordGenerator struct {
	*gir.Record
	GoName string

	// TODO: move these out of here.
	Methods []callable.Generator
	Getters []recordGetter

	// TODO: make a []callableGenerator for constructors
	Callable callable.Generator

	typ gir.TypeFindResult
	hdr file.Header
	gen FileGenerator
}

type recordGetter struct {
	InfoElements gir.InfoElements

	Name  string
	Type  string
	Block string // assume first_letter recv
}

func NewRecordGenerator(gen FileGenerator) RecordGenerator {
	return RecordGenerator{
		gen:      gen,
		Callable: callable.NewGenerator(gen),
	}
}

// hHeader returns the RecordGenerator's current file header.
func (rg *RecordGenerator) Header() *file.Header {
	return &rg.hdr
}

func (rg *RecordGenerator) Use(rec *gir.Record) bool {
	rg.hdr.Reset()

	if !CanGenerateRecord(rg.gen, rec) {
		return false
	}

	if rec.GLibGetType != "" {
		// Need this for g_value_get_boxed().
		rg.hdr.NeedsGLibObject()
	}

	rg.typ.NamespaceFindResult = rg.gen.Namespace()
	rg.typ.Type = rec

	rg.Record = rec
	rg.GoName = strcases.PascalToGo(rec.Name)
	rg.Methods = rg.methods()
	rg.Getters = rg.getters()

	return true
}

func (rg *RecordGenerator) UseConstructor(ctor *gir.Constructor, className string) bool {
	if !rg.Callable.Use(&rg.typ, &ctor.CallableAttrs) {
		return false
	}

	file.ApplyHeader(rg, &rg.Callable)
	rg.Callable.Name = strings.TrimPrefix(rg.Callable.Name, "New")
	rg.Callable.Name = "New" + rg.GoName + rg.Callable.Name

	return true
}

func (rg *RecordGenerator) methods() []callable.Generator {
	callables := callable.Grow(rg.Methods, len(rg.Record.Methods))

	for i := range rg.Record.Methods {
		method := &rg.Record.Methods[i]

		cbgen := callable.NewGenerator(rg.gen)
		if !cbgen.Use(&rg.typ, &method.CallableAttrs) {
			rg.Logln(logger.Skip, "record", rg.Name, "method", method.Name)
			continue
		}

		file.ApplyHeader(rg, &cbgen)
		callables = append(callables, cbgen)
	}

	callable.RenameGetters("", callables)
	return callables
}

func (rg *RecordGenerator) getters() []recordGetter {
	getters := rg.Getters[:0]

	// Disguised means opaque, so we're not supposed to access these fields.
	if rg.Disguised {
		return getters
	}

	methodNames := make(map[string]struct{}, len(rg.Methods))
	for _, method := range rg.Methods {
		// Fill the name map. The name we get here is the transformed name
		// (method is a callable.Generator), so we don't have to do it again.
		methodNames[method.Name] = struct{}{}
	}

	fieldCollides := func(name string) bool {
		_, collides := methodNames[strcases.SnakeToGo(true, name)]
		return collides
	}

	recv := strcases.FirstLetter(rg.GoName)
	fields := make([]typeconv.ConversionValue, 0, len(rg.Fields))

	for _, field := range rg.Fields {
		if ignoreField(&field) || mustIgnoreAny(rg.gen, field.AnyType) {
			rg.Logln(logger.Debug, "skipping field", field.Name, "after ignoreField")
			continue
		}
		if types.FilterField(rg.gen, rg.Name, &field) {
			rg.Logln(logger.Skip, "record", rg.Name, "field", field.Name)
			continue
		}

		value := typeconv.NewFieldValue(recv, "v", field)

		// Double-check if we have a method with the existing name.
		if fieldCollides(value.Name) {
			rg.Logln(logger.Debug, "colliding name", value.Name)
			continue
		}

		fields = append(fields, value)
	}

	converter := typeconv.NewConverter(rg.gen, &rg.typ, fields)
	converter.UseLogger(rg)

	for i := range fields {
		converted := converter.Convert(i)
		if converted == nil {
			rg.Logln(logger.Skip, "record", rg.Name, "field", fields[i].Name)
			continue
		}

		file.ApplyHeader(rg, converted)

		b := pen.NewBlock()
		b.Linef(converted.Out.Declare)
		b.Linef(converted.Conversion)
		b.Linef("return v")

		getters = append(getters, recordGetter{
			Name:  strcases.SnakeToGo(true, converted.Name),
			Type:  converted.Out.Type,
			Block: b.String(),
			InfoElements: gir.InfoElements{
				DocElements: gir.DocElements{Doc: fields[i].Doc},
			},
		})
	}

	return getters
}

// ignoreField returns true if the given field should be ignored.
func ignoreField(field *gir.Field) bool {
	// For "Bits > 0", we can't safely do this in Go (and probably not CGo
	// either?) so we're not doing it.
	return field.Private || field.Bits > 0 || !field.IsReadable()
}

func (rg *RecordGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("record %s (C.%s):", rg.GoName, rg.CType)
	rg.gen.Logln(lvl, logger.Prefix(v, p)...)
}

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
		native C.{{ .CType }}
	}

	// Wrap{{ .GoName }} wraps the C unsafe.Pointer to be the right type. It is
	// primarily used internally.
	func Wrap{{ .GoName }}(ptr unsafe.Pointer) *{{ .GoName }} {
		return (*{{ .GoName }})(ptr)
	}

	{{ if .GLibGetType }}
	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
		return (*{{ .GoName }})(unsafe.Pointer(b)), nil
	}
	{{ end }}

	{{ range .Constructors }}
	{{ if $.UseConstructor . $.GoName }}
	// {{ $.Callable.Name }} constructs a struct {{ $.GoName }}.
	func {{ $.Callable.Name }}{{ $.Callable.Tail }} {{ $.Callable.Block }}
	{{ end }}
	{{ end }}

	{{ $recv := (FirstLetter $.GoName) }}

	// Native returns the underlying C source pointer.
	func ({{ $recv }} *{{ .GoName }}) Native() unsafe.Pointer {
		return unsafe.Pointer(&{{ FirstLetter .GoName }}.native)
	}

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

		// Check the type against the ignored list, since ignores are usually
		// important, and CGo might still try to resolve an ignored type.
		if mustIgnoreAny(gen, field.AnyType) {
			log("ignored because field", field.Name)
			return false
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
		return types.Filter(gen, any.Type.Name, any.Type.CType)
	case any.Array != nil:
		return mustIgnoreAny(gen, any.Array.AnyType)
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
	// Need unsafe for the wrapper.
	writer.Header().Import("unsafe")

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

	rg.Record = rec
	rg.GoName = strcases.PascalToGo(rec.Name)
	rg.Methods = rg.methods()
	rg.Getters = rg.getters()

	return true
}

func (rg *RecordGenerator) UseConstructor(ctor *gir.Constructor, className string) bool {
	if !rg.Callable.Use(&ctor.CallableAttrs) {
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
		if !cbgen.Use(&method.CallableAttrs) {
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
		methodNames[method.Name] = struct{}{}
	}

	recv := strcases.FirstLetter(rg.GoName)
	fields := make([]typeconv.ConversionValue, 0, len(rg.Fields))

	for _, field := range rg.Fields {
		// For "Bits > 0", we can't safely do this in Go (and probably not CGo
		// either?) so we're not doing it.
		if ignoreField(&field) {
			continue
		}

		// Check if we have a method with the existing name.
		if _, collides := methodNames[strcases.SnakeToGo(true, field.Name)]; collides {
			// Skip generating the getter if we have a colliding method.
			continue
		}

		fields = append(fields, typeconv.NewFieldValue(recv, "v", field))
	}

	converter := typeconv.NewConverter(rg.gen, fields)
	converter.UseLogger(rg)

	for i := range fields {
		converted := converter.Convert(i)
		if converted == nil {
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
	return field.Private || field.Bits > 0 || !field.IsReadable() || !field.Writable
}

func (rg *RecordGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("record %s (C.%s):", rg.GoName, rg.CType)
	rg.gen.Logln(lvl, logger.Prefix(v, p)...)
}

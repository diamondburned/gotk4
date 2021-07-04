package generators

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/core/pen"
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
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
	type {{ .GoName }} C.{{ .CType }}

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
		return unsafe.Pointer({{ FirstLetter .GoName }})
	}

	{{ range .Getters }}
	// {{ .GoName }} gets the field inside the struct.
	func ({{ $recv }} *{{ $.GoName }}) {{ .GoName }}() {{ .GoType }} {{ .Block }}
	{{ end }}

	{{ range .Methods }}
	{{ GoDoc . 0 }}
	func ({{ .Recv }} *{{ $.GoName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

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
	GoName string
	GoType string
	Block  string // assume first_letter recv
}

func NewRecordGenerator(gen FileGenerator) RecordGenerator {
	return RecordGenerator{
		gen:      gen,
		Callable: callable.NewGenerator(gen),
	}
}

// canRecord returns true if this record is allowed.
func canRecord(gen FileGenerator, rec *gir.Record) bool {
	if !rec.IsIntrospectable() || types.Filter(gen, rec.Name, rec.CType) {
		return false
	}

	// GLibIsGTypeStructFor seems to be records used in addition to classes due
	// to C? Not sure, but we likely don't need it.
	if rec.GLibIsGTypeStructFor != "" || strings.HasPrefix(rec.Name, "_") {
		return false
	}

	for _, suffix := range recordIgnoreSuffixes {
		if strings.HasSuffix(rec.Name, suffix) {
			return false
		}
	}

	// Ignore non-type/array fields.
	for _, field := range rec.Fields {
		// Check the type against the ignored list, since ignores are usually
		// important, and CGo might still try to resolve an ignored type.
		if mustIgnoreAny(gen, field.AnyType) {
			gen.Logln(logger.Debug, "ignored because field", field.Name)
			return false
		}
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

// hHeader returns the RecordGenerator's current file header.
func (rg *RecordGenerator) Header() *file.Header {
	return &rg.hdr
}

func (rg *RecordGenerator) Use(rec *gir.Record) bool {
	if !canRecord(rg.gen, rec) {
		return false
	}

	rg.hdr.Reset()
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

	rg.Callable.Name = strings.TrimPrefix(rg.Callable.Name, "New")
	rg.Callable.Name = "New" + rg.GoName + rg.Callable.Name

	return true
}

func (rg *RecordGenerator) methods() []callable.Generator {
	callables := callable.Grow(rg.Methods, len(rg.Record.Methods))

	for _, method := range rg.Record.Methods {
		cbgen := callable.NewGenerator(rg.gen)
		if !cbgen.Use(&method.CallableAttrs) {
			continue
		}

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
		if field.Private || field.Bits > 0 || !field.Readable && !field.Writable {
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

	// Add all imports once we've gone over conversion of all fields.
	defer file.ApplyHeader(rg, converter)

	for i := range fields {
		converted := converter.Convert(i)
		if converted == nil {
			continue
		}

		b := pen.NewBlock()
		b.Linef(converted.OutDeclare)
		b.Linef(converted.Conversion)
		b.Linef("return v")

		getters = append(getters, recordGetter{
			GoName: strcases.SnakeToGo(true, converted.Name),
			GoType: converted.OutType,
			Block:  b.String(),
		})
	}

	return getters
}

func (rg *RecordGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("record %s (C.%s):", rg.GoName, rg.CType)
	rg.gen.Logln(lvl, logger.Prefix(v, p))
}

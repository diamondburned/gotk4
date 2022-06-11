package generators

import (
	"fmt"
	"strings"
	"text/tabwriter"

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

var recordNeedCDeclSuffixes = []string{
	"Interface",
	"Iface",
	"Class",
}

var recordTmpl = gotmpl.NewGoTemplate(`
	{{ $impl := UnexportPascal .GoName }}

	{{ GoDoc . 0 (AdditionalString "An instance of this type is always passed by reference.") }}
	type {{ .GoName }} struct {
		*{{ $impl }}
	}

	// {{ $impl }} is the struct that's finalized.
	type {{ $impl }} struct {
		native {{ .CGoPtrType }}
	}

	{{ if .Marshaler }}
	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
		return &{{ .GoName }}{&{{ $impl }}{({{ .CGoPtrType }})(b)}}, nil
	}
	{{ end }}

	{{ range .Constructors }}
	{{ if $.UseConstructor . }}
	// {{ $.Callable.Name }} constructs a struct {{ $.GoName }}.
	func {{ $.Callable.Name }}{{ $.Callable.Tail }} {{ $.Callable.Block }}
	{{ end }}
	{{ else }}
	{{ with .ManualConstructor }}
	// New{{ $.GoName }} creates a new {{ $.GoName }} instance from the given
	// fields.
	func New{{ $.GoName }}({{ .Fields }}) {{ $.GoName }} {{ .Body }}
	{{ end }}
	{{ end }}

	{{ $recv := (FirstLetter $.GoName) }}

	{{ range .Getters }}
	{{ GoDoc . 0 }}
	func ({{ $recv }} *{{ $.GoName }}) {{ .Name }}() {{ .Type }} {{ .Block }}
	{{ end }}

	{{ range .Setters }}
	{{ GoDoc . 0 }}
	func ({{ $recv }} *{{ $.GoName }}) Set{{ .Name }}({{ .Param }}) {{ .Block }}
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
		if ignoreField(field) {
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

	if record.GLibGetType != "" && !types.FilterCType(gen, record.GLibGetType) {
		recordGen.Marshaler = true
		writer.Header().NeedsExternGLib()
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
	GoName    string
	Marshaler bool

	// TODO: move these out of here.
	Methods []callable.Generator
	Getters []recordGetter
	Setters []recordSetter

	// ManualConstructor is the body function of the manually-generated
	// constructor. Skip if empty.
	ManualConstructor *RecordConstructor

	// TODO: make a []callableGenerator for constructors
	Callable callable.Generator

	typ gir.TypeFindResult
	hdr file.Header
	gen FileGenerator
}

// RecordConstructor describes a manual record constructor.
type RecordConstructor struct {
	Fields string
	Body   string // return struct value
}

type recordGetter struct {
	InfoElements gir.InfoElements

	Name  string
	Type  string
	Block string // assume first_letter recv
}

type recordSetter struct {
	InfoElements gir.InfoElements

	Name  string // not prefixed with Set
	Param string
	Block string
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
	rg.Marshaler = false

	if !CanGenerateRecord(rg.gen, rec) {
		return false
	}

	rg.typ.NamespaceFindResult = rg.gen.Namespace()
	rg.typ.Type = rec

	rg.Record = rec
	rg.GoName = strcases.PascalToGo(rec.Name)
	rg.methods()
	rg.getters()

	return true
}

func (rg *RecordGenerator) UseConstructor(ctor *gir.Constructor) bool {
	if types.FilterSub(rg.gen, rg.Name, ctor.Name, ctor.CIdentifier) {
		return false
	}

	if !rg.Callable.Use(&rg.typ, &ctor.CallableAttrs) {
		return false
	}

	file.ApplyHeader(rg, &rg.Callable)
	rg.Callable.Name = strings.TrimPrefix(rg.Callable.Name, "New")
	rg.Callable.Name = "New" + rg.GoName + rg.Callable.Name

	return true
}

func (rg *RecordGenerator) methods() {
	callables := callable.Grow(rg.Methods, len(rg.Record.Methods))

	for i := range rg.Record.Methods {
		method := &rg.Record.Methods[i]

		if types.FilterMethod(rg.gen, rg.Name, method) {
			rg.Logln(logger.Debug, "filtered method", method.CIdentifier)
			continue
		}

		cbgen := callable.NewGenerator(rg.gen)
		if !cbgen.Use(&rg.typ, &method.CallableAttrs) {
			rg.Logln(logger.Skip, "record", rg.Name, "method", method.Name)
			continue
		}

		file.ApplyHeader(rg, &cbgen)
		callables = append(callables, cbgen)
	}

	callable.RenameGetters("", callables)
	rg.Methods = callables
}

func (rg *RecordGenerator) getters() {
	rg.Getters = rg.Getters[:0]
	rg.Setters = rg.Setters[:0]

	// Disguised means opaque, so we're not supposed to access these fields.
	if rg.Disguised {
		return
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
	values := make([]typeconv.ConversionValue, 0, len(rg.Fields))
	fields := make([]string, 0, len(rg.Fields))

	// Do a constructor when the record has none.
	willDoConstructor := len(rg.Fields) > 0 && len(rg.Record.Constructors) == 0

	for _, field := range rg.Fields {
		if ignoreField(field) || mustIgnoreAny(rg.gen, field.AnyType) {
			rg.Logln(logger.Debug, "skipping field", field.Name, "after ignoreField")
			willDoConstructor = false
			continue
		}
		if types.FilterField(rg.gen, rg.Name, &field) {
			rg.Logln(logger.Skip, "record", rg.Name, "field", field.Name)
			willDoConstructor = false
			continue
		}
		// We can't generate bitfield accesses if we're in runtime mode.
		if field.Bits > 0 && rg.gen.LinkMode() == types.RuntimeLinkMode {
			continue
		}

		// Check type for constructor.
		if willDoConstructor {
			if field.Type == nil ||
				types.GIRPrimitiveGo(field.Type.Name) == "" ||
				types.AnyTypeIsPtr(field.AnyType) {

				// Field is more than a primitive. Skip the constructor.
				willDoConstructor = false
			}
		}

		value := typeconv.NewFieldValue("valptr", "v", field)

		// Double-check if we have a method with the existing name.
		if fieldCollides(value.Name) {
			rg.Logln(logger.Debug, "colliding name", value.Name)
			continue
		}

		values = append(values, value)
		fields = append(fields, field.Name)

		// Generate a SetX if no methods collide with the name.
		if field.Writable && !fieldCollides("set_"+value.Name) {
			in := strcases.SnakeToGo(false, field.Name)
			values = append(values, typeconv.NewFieldSetValue(in, "_valptr", field))
			fields = append(fields, field.Name)
		}
	}

	converter := typeconv.NewConverter(rg.gen, &rg.typ, values)
	converter.MustCast = rg.gen.LinkMode() == types.RuntimeLinkMode
	converter.UseLogger(rg)

	for i := range values {
		converted := converter.Convert(i)
		if converted == nil {
			rg.Logln(logger.Skip, "record", rg.Name, "field", values[i].Name)
			willDoConstructor = false
			continue
		}

		info := gir.InfoElements{
			DocElements: gir.DocElements{Doc: values[i].Doc},
		}

		// TODO: handle setters: we currently cannot yet do this, because we
		// don't have the freeing code separated in typeconv. This is in the
		// list of things to do.

		switch converted.Direction {
		case typeconv.ConvertCToGo: // getter
			b := pen.NewBlock()

			// Runtime-link mode assumes a hard-coded valptr name.
			if rg.gen.LinkMode() == types.RuntimeLinkMode {
				rg.hdr.ImportCore("girepository")
				b.Linef("offset := girepository.MustFind(%q, %q).StructFieldOffset(%q)", rg.typ.Namespace.Name, rg.typ.Name(), fields[i])
				b.Linef("valptr := unsafe.Add(unsafe.Pointer(%s), offset)", recv)
			} else {
				b.Linef("valptr := %s.native.%s", recv, strcases.CGoField(fields[i]))
			}

			b.Linef(converted.Out.Declare)
			b.Linef(converted.Conversion)
			b.Linef("return v")

			file.ApplyHeader(rg, converted)

			rg.Getters = append(rg.Getters, recordGetter{
				Name:         strcases.SnakeToGo(true, converted.Name),
				Type:         converted.Out.Type,
				Block:        b.String(),
				InfoElements: info,
			})
		}
		// case typeconv.ConvertGoToC: // setter
		// 	b := pen.NewBlock()
		// 	b.Linef(converted.Conversion)

		// 	rg.Setters = append(rg.Setters, recordSetter{
		// 		Name:         strcases.SnakeToGo(true, converted.Name),
		// 		Param:        converted.InName + " " + converted.In.Type,
		// 		Block:        b.String(),
		// 		InfoElements: info,
		// 	})
		// }
	}

	if willDoConstructor {
		rg.genManualConstructor()
	}
}

// ignoreField returns true if the given field should be ignored.
func ignoreField(field gir.Field) bool {
	// For "Bits > 0", we can't safely do this in Go (and probably not CGo
	// either?) so we're not doing it.
	return field.Private || field.Bits > 0 || !field.IsReadable()
}

func (rg *RecordGenerator) genManualConstructor() {
	params := pen.NewJoints(", ", len(rg.Fields))
	convts := make([]typeconv.ConversionValue, 0, len(rg.Fields))

	for i, field := range rg.Fields {
		name := strcases.SnakeToGo(false, field.Name)
		// Try and coalesce the type, if we can.
		param := name
		// We can do this only if the field is not the last one and that the
		// type matches. When either of those fails, we have to add the type.
		if i == len(rg.Fields)-1 || field.Type.Name != rg.Fields[i+1].Type.Name {
			param += " " + types.GIRPrimitiveGo(field.Type.Name)
		}

		params.Add(param)
		// always Go to C.
		convts = append(convts, typeconv.NewFieldSetValue(name, fmt.Sprintf("f%d", i), field))
	}

	conv := typeconv.NewConverter(rg.gen, &rg.typ, convts)
	conv.UseLogger(rg)

	results := conv.ConvertAll()
	if results == nil {
		rg.ManualConstructor = nil
		return
	}

	p := pen.NewBlockSections(512, 512)
	p.Linef(1, "v := C.%s{", rg.CType)

	for i, result := range results {
		if result.Direction != typeconv.ConvertGoToC {
			continue
		}

		p.Line(0, result.Out.Declare)
		p.Line(0, result.Conversion)
		p.Linef(1, "%s: %s,", rg.Fields[i].Name, result.OutName)
	}

	p.Linef(1, "}")
	p.Linef(1, "")

	rg.hdr.Import("unsafe")
	rg.hdr.ImportCore("gextras")

	// No finalizers needed, since the struct is completely allocated on the Go
	// heap.
	p.Linef(1, "return *(*%s)(gextras.NewStructNative(unsafe.Pointer(&v)))", rg.GoName)

	rg.ManualConstructor = &RecordConstructor{
		Fields: params.Join(),
		Body:   p.String(),
	}
}

func (rg *RecordGenerator) CGoPtrType() string {
	switch rg.gen.LinkMode() {
	case types.DynamicLinkMode:
		return "*C." + rg.Record.CType
	case types.RuntimeLinkMode:
		return "unsafe.Pointer"
	default:
		panic("unreachable")
	}
}

func (rg *RecordGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("record %s (C.%s):", rg.GoName, rg.Record.CType)
	rg.gen.Logln(lvl, logger.Prefix(v, p)...)
}

// GenerateCPrimitiveRecord generates C struct code with primitive C types.
func GenerateCPrimitiveRecord(gen types.FileGenerator, rec *gir.Record) string {
	var b strings.Builder

	w := tabwriter.NewWriter(&b, 0, 4, 1, ' ', 0)

	for i, field := range rec.Fields {
		if i != 0 {
			fmt.Fprintln(w)
		}

		if field.Callback != nil {
			// Any pointer works.
			fmt.Fprintf(w, "    void*\t%s;", field.Name)
			continue
		}

		fmt.Fprint(w, "    ", types.AnyTypeCPrimitive(gen, field.AnyType), "\t", field.Name)
		if field.Bits > 0 {
			fmt.Fprint(w, "\t : ", field.Bits)
		}
		fmt.Fprint(w, ";")
	}

	if err := w.Flush(); err != nil {
		panic("cannot flush tabwriter: " + err.Error())
	}

	return fmt.Sprintf("struct %s {\n%s\n};", rec.Name, b.String())
}

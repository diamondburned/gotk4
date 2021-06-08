package girgen

import (
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
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

var recordTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .GoName }}
	type {{ .GoName }} struct {
		native C.{{ .CType }}
	}

	// Wrap{{ .GoName }} wraps the C unsafe.Pointer to be the right type. It is
	// primarily used internally.
	func Wrap{{ .GoName }}(ptr unsafe.Pointer) *{{ .GoName }} {
		if ptr == nil {
			return nil
		}

		return (*{{ .GoName }})(ptr)
	}

	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
		return Wrap{{ .GoName }}(unsafe.Pointer(b)), nil
	}

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
	// {{ .GoName }} gets the field inside the struct.
	func ({{ $recv }} *{{ $.GoName }}) {{ .GoName }}() {{ .GoType }} {{ .Block }}
	{{ end }}

	{{ range .Methods }}
	{{ GoDoc .Doc 0 .Name }}
	func ({{ .Recv }} *{{ $.GoName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

type recordGenerator struct {
	gir.Record
	GoName  string
	Methods []callableGenerator
	Getters []recordGetter

	Callable callableGenerator

	fg *FileGenerator
	ng *NamespaceGenerator
}

type recordGetter struct {
	GoName string
	GoType string
	Block  string // assume first_letter recv
}

func newRecordGenerator(ng *NamespaceGenerator) *recordGenerator {
	return &recordGenerator{
		Callable: newCallableGenerator(ng),
		ng:       ng,
	}
}

// canRecord returns true if this record is allowed.
func canRecord(ng *NamespaceGenerator, rec gir.Record, logger TypeResolver) bool {
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
		if ng.mustIgnoreAny(field.AnyType) {
			tryLogln(logger, LogSkip, "record", rec.Name, "skipped, field", field.Name)
			return false
		}
	}

	return true
}

func (rg *recordGenerator) Use(rec gir.Record) bool {
	rg.fg = rg.ng.FileFromSource(rec.SourcePosition)

	if !canRecord(rg.ng, rec, rg.fg) {
		return false
	}

	rg.Record = rec
	rg.GoName = PascalToGo(rec.Name)
	rg.Methods = rg.methods()
	rg.Getters = rg.getters()

	return true
}

func (rg *recordGenerator) UseConstructor(ctor gir.Constructor, className string) bool {
	if !rg.Callable.Use(ctor.CallableAttrs) {
		return false
	}

	rg.Callable.Name = strings.TrimPrefix(rg.Callable.Name, "New")
	rg.Callable.Name = "New" + rg.GoName + rg.Callable.Name

	return true
}

func (rg *recordGenerator) methods() []callableGenerator {
	callables := callableGrow(rg.Methods, len(rg.Record.Methods))

	for _, method := range rg.Record.Methods {
		cbgen := newCallableGenerator(rg.ng)
		if !cbgen.Use(method.CallableAttrs) {
			continue
		}

		callables = append(callables, cbgen)
	}

	callableRenameGetters(callables)
	return callables
}

func (rg *recordGenerator) getters() []recordGetter {
	getters := rg.Getters[:0]

	// Disguised means opaque, so we're not supposed to access these fields.
	if rg.Disguised {
		return getters
	}

	methodNames := make(map[string]struct{}, len(rg.Methods))
	for _, method := range rg.Methods {
		methodNames[method.Name] = struct{}{}
	}

	recv := FirstLetter(rg.GoName)
	fields := make([]CValueProp, len(rg.Fields))

	for i, field := range rg.Fields {
		// For "Bits > 0", we can't safely do this in Go (and probably not CGo
		// either?) so we're not doing it.
		if field.Private || field.Bits > 0 || !field.Readable && !field.Writable {
			continue
		}

		goName := SnakeToGo(true, field.Name)

		// Check if we have a method with the existing name.
		if _, collides := methodNames[goName]; collides {
			// Skip generating the getter if we have a colliding method.
			continue
		}

		fields[i] = CValueProp{
			ValueProp: NewValuePropField(recv, "v", field),
		}
	}

	conversion := rg.fg.CGoConverter(rg.Name, fields)

	for i, field := range fields {
		if field.ValueProp.IsZero() {
			continue
		}

		converted := conversion.Convert(i)
		if converted == nil {
			continue
		}

		converted.Apply(rg.fg)

		b := pen.NewBlock()
		b.Linef(converted.OutDeclare)
		b.Linef(converted.Conversion)
		b.Linef("return v")

		getters = append(getters, recordGetter{
			GoName: SnakeToGo(true, rg.Fields[i].Name),
			GoType: converted.OutType,
			Block:  b.String(),
		})
	}

	return getters
}

func (ng *NamespaceGenerator) generateRecords() {
	rg := newRecordGenerator(ng)

	for _, record := range ng.current.Namespace.Records {
		if ng.mustIgnore(record.Name, record.CType) {
			continue
		}

		if !rg.Use(record) {
			continue
		}

		rg.fg.addImport("unsafe")
		rg.fg.needsGLibObject()

		if record.GLibGetType != "" && !ng.mustIgnoreC(record.GLibGetType) {
			rg.fg.addMarshaler(record.GLibGetType, rg.GoName)
		}

		rg.fg.pen.WriteTmpl(recordTmpl, &rg)
	}
}

package girgen

import (
	"strings"

	"github.com/diamondburned/gotk4/gir"
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
	// No idea why this is here.
	"Class",
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

	// Native returns the underlying C source pointer.
	func ({{ FirstLetter .GoName }} *{{ .GoName }}) Native() unsafe.Pointer {
		return unsafe.Pointer(&{{ FirstLetter .GoName }}.native)
	}

	{{ range .Constructors }}
	{{ if $.UseConstructor . $.GoName }}
	// {{ $.Callable.Name }} constructs a struct {{ $.GoName }}.
	func {{ $.Callable.Name }}{{ $.Callable.Tail }} {{ $.Callable.Block }}
	{{ end }}
	{{ end }}

	{{ range .Getters }}
	// {{ .GoName }} gets the field inside the struct.
	func ({{ FirstLetter $.GoName }} *{{ $.GoName }}) {{ .GoName }}() {{ .GoType }} {
		var ret {{ .GoType }}
		{{ .Convert }}
		return ret
	}
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
	gir.Field
	Name    string
	GoName  string
	GoType  string
	Convert string // assume first_letter -> "ret"
}

func newRecordGenerator(ng *NamespaceGenerator) *recordGenerator {
	return &recordGenerator{
		Callable: newCallableGenerator(ng),
		ng:       ng,
	}
}

// canRecord returns true if this record is allowed.
func canRecord(logger TypeResolver, rec gir.Record, log bool) bool {
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
		if field.Type == nil && field.Array == nil {
			tryLogln(logger, LogSkip, "record", rec.Name, "skipped, field", field.Name)
			return false
		}
	}

	return true
}

func (rg *recordGenerator) Use(rec gir.Record) bool {
	if !canRecord(rg.ng, rec, true) {
		return false
	}

	rg.fg = rg.ng.FileFromSource(rec.SourcePosition)

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

	var ignores ignoreIxs

	methodNames := make(map[string]struct{}, len(rg.Methods))
	for _, method := range rg.Methods {
		methodNames[method.Name] = struct{}{}
	}

	recv := FirstLetter(rg.GoName)
	fieldAt := func(i int) string {
		return recv + ".native." + cgoField(rg.Fields[i].Name)
	}

	for i, field := range rg.Fields {
		ignores.fieldIgnore(field)

		// For "Bits > 0", we can't safely do this in Go (and probably not CGo
		// either?) so we're not doing it.

		if field.Private || field.Bits > 0 || ignores.ignore(i) {
			continue
		}
		if field.Readable != nil && !*field.Readable {
			continue
		}

		goName := SnakeToGo(true, field.Name)

		// Check if we have a method with the existing name.
		if _, collides := methodNames[goName]; collides {
			// Skip generating the getter if we have a colliding method.
			continue
		}

		typ, ok := GoAnyType(rg.fg, field.AnyType, true)
		if !ok {
			continue
		}

		convert := rg.fg.CGoConverter(TypeConversionToGo{
			TypeConversion: TypeConversion{
				Value:  fieldAt(i),
				Target: "ret",
				Type:   field.AnyType,
				ArgAt:  fieldAt,
				Owner: gir.TransferOwnership{
					// Assume we have the ownership of the C value, because we do.
					TransferOwnership: "none",
				},
				ParentName: rg.CType,
			},
		})
		if convert == "" {
			continue
		}

		getters = append(getters, recordGetter{
			Field:   field,
			Name:    cgoField(field.Name),
			GoName:  goName,
			GoType:  typ,
			Convert: convert,
		})
	}

	return getters
}

// needsNative returns true if the record needs a private C field for
// referencing.
func (rg *recordGenerator) needsNative() bool {
	for _, field := range rg.Fields {
		if field.Private {
			return true
		}
	}

	return len(rg.Fields) == 0
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

		rg.fg.pen.BlockTmpl(recordTmpl, &rg)
	}
}

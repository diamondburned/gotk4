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
		return Wrap{{ .GoName }}(unsafe.Pointer(b))
	}

	// Native returns the underlying C source pointer.
	func ({{ FirstLetter .GoName }} *{{ .GoName }}) Native() unsafe.Pointer {
		return unsafe.Pointer(&{{ FirstLetter .GoName }}.native)
	}

	{{ range .Constructors }}
	{{ if $.Callable.UseConstructor . }}
	// {{ $.Callable.Name }} constructs a struct {{ $.GoName }}.
	func {{ $.Callable.Name }}{{ $.Callable.Tail }} {{ $.Callable.Block }}
	{{ end }}
	{{ end }}

	{{ range .Getters }}
	// {{ .GoName }} gets the field inside the struct.
	func ({{ FirstLetter .GoName }} *{{ $.GoName }}) {{ .GoName }}() {{ .GoType }} {
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

	Ng *NamespaceGenerator
}

type recordGetter struct {
	gir.Field
	Name    string
	GoName  string
	GoType  string
	Convert string // assume first_letter -> "ret"
}

// canRecord returns true if this record is allowed.
func (ng *NamespaceGenerator) canRecord(rec gir.Record) bool {
	// GLibIsGTypeStructFor seems to be records used in addition to classes due
	// to C? Not sure, but we likely don't need it.
	if rec.Disguised || rec.GLibIsGTypeStructFor != "" || strings.HasPrefix(rec.Name, "_") {
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
			return false
		}
	}

	return true
}

func (rg *recordGenerator) Use(rec gir.Record) bool {
	if !rg.Ng.canRecord(rec) {
		return false
	}

	rg.Record = rec
	rg.GoName = PascalToGo(rec.Name)
	rg.Methods = rg.methods()
	rg.Getters = rg.getters()

	return true
}

func (rg *recordGenerator) CtorName(ctor gir.Constructor) string {
	name := SnakeToGo(true, ctor.Name)
	name = strings.TrimPrefix(name, "New")
	return "New" + rg.GoName + name
}

func (rg *recordGenerator) methods() []callableGenerator {
	callables := rg.Methods[:0]

	for _, method := range rg.Methods {
		cbgen := newCallableGenerator(rg.Ng)
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

		if field.Private || ignores.ignore(i) {
			continue
		}

		goName := SnakeToGo(true, field.Name)

		// Check if we have a method with the existing name.
		if _, collides := methodNames[goName]; collides {
			// Skip generating the getter if we have a colliding method.
			continue
		}

		typ, ok := rg.Ng.ResolveAnyType(field.AnyType, true)
		if !ok {
			continue
		}

		convert := rg.Ng.CGoConverter(TypeConversionToGo{
			TypeConversion: TypeConversion{
				Value:  fieldAt(i),
				Target: "ret",
				Type:   field.AnyType,
				ArgAt:  fieldAt,
				Owner: gir.TransferOwnership{
					// Assume we have the ownership of the C value, because we do.
					TransferOwnership: "none",
				},
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
	rg := recordGenerator{
		Callable: callableGenerator{Ng: ng},
		Ng:       ng,
	}
	imported := false

	for _, record := range ng.current.Namespace.Records {
		if !rg.Use(record) {
			ng.logln(logInfo, "record", record.Name, "skipped")
			continue
		}

		// Add the needed imports once.
		if !imported {
			rg.Ng.addImport("unsafe")
			rg.Ng.addImport("runtime")
			imported = true
		}

		ng.pen.BlockTmpl(recordTmpl, &rg)

		// TODO: record Native() uintptr
		// TODO: record methods.
		// TODO: handle transfer-ownership
		// TODO: free method
	}
}

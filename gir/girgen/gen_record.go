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
		{{ if .PublicFields -}}
		{{ range .PublicFields -}}
		{{ GoDoc .Doc 1 .GoName }}
		{{ .GoName }} {{ .GoType }}
		{{ end }}

		{{ end }}
		native *C.{{ .CType }}
	}

	func wrap{{ .GoName }}(p *C.{{ .CType }}) *{{ .GoName }} {
		{{ if .NeedsNative -}}
		v := {{ .GoName }}{native: p}
		{{ else -}}
		var v {{ .GoName }}
		{{ end }}

		{{ range .PublicFields -}}
		{{ $.Convert . "p" "v" }}
		{{ end }}

		{{ if .NeedsNative }}
		runtime.SetFinalizer(&v, nil)
		runtime.SetFinalizer(&v, (*{{ .GoName }}).free)
		{{ end }}
		return &v
	}

	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
		c := (*C.{{ .CType }})(unsafe.Pointer(b))

		return wrap{{ .GoName }}(c)
	}

	{{ $recv := (FirstLetter .GoName) }}

	{{ if .NeedsNative }}
	func ({{ $recv }} *{{ .GoName }}) free() {}
	{{ end }}

	// Native returns the pointer to *C.{{ .CType }}. The caller is expected to
	// cast.
	func ({{ $recv }} *{{ .GoName }}) Native() unsafe.Pointer {
		return unsafe.Pointer({{ $recv }}.native)
	}

	{{ range .Constructors }}
	{{ with $tail := ($.Ng.FnCall .CallableAttrs) }}
	func New{{ $.GoName }}{{ $tail }}
	{{ end }}
	{{ end }}
`)

type recordGenerator struct {
	gir.Record
	GoName       string
	NeedsNative  bool
	PublicFields []recordField

	Ng *NamespaceGenerator
}

type recordField struct {
	gir.Field
	Name   string
	GoName string
	GoType string
}

func (rg *recordGenerator) Use(rec gir.Record) bool {
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

	rg.Record = rec
	rg.GoName = PascalToGo(rec.Name)
	rg.NeedsNative = rg.needsNative()
	rg.PublicFields = rg.publicFields()

	return true
}

func (rg *recordGenerator) Convert(field recordField, cStruct, goStruct string) string {
	return rg.Ng.CGoConverter(
		cStruct+"."+field.Name,
		goStruct+"."+field.GoName,
		field.AnyType,
		func(i int) string { return cStruct + "." + cgoField(rg.Fields[i].Name) },
	)
}

func (rg *recordGenerator) publicFields() []recordField {
	fields := make([]recordField, 0, len(rg.Fields))
	var ignores ignoreIxs

	for i, field := range rg.Fields {
		ignores.fieldIgnore(field)

		if field.Private || ignores.ignore(i) {
			continue
		}

		typ, ok := rg.Ng.ResolveAnyType(field.AnyType, true)
		if !ok {
			continue
		}

		fields = append(fields, recordField{
			Field:  field,
			Name:   cgoField(field.Name),
			GoName: SnakeToGo(true, field.Name),
			GoType: typ,
		})
	}

	return fields
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
	rg := recordGenerator{Ng: ng}
	imported := false

	for _, record := range ng.current.Namespace.Records {
		if !rg.Use(record) {
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

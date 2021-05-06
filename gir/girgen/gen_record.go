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

func ignoreRecord(rec gir.Record) bool {
	if rec.Disguised || strings.HasPrefix(rec.Name, "_") {
		return true
	}

	for _, suffix := range recordIgnoreSuffixes {
		if strings.HasSuffix(rec.Name, suffix) {
			return true
		}
	}

	return false
}

var recordTmpl = newGoTemplate(`
	{{ $type := (PascalToGo .Name) }}

	{{ GoDoc .Doc 0 $type }}
	type {{ $type }} struct {
		{{ range .Fields -}}
		{{ $.Field . }}
		{{ end -}}

		{{ if .NeedsNative -}}
		native *C.{{ .CType }}
		{{ end -}}
	}
`)

type recordGenerator struct {
	gir.Record
	Ng *NamespaceGenerator
}

func (rg recordGenerator) NeedsNative() bool {
	for _, field := range rg.Fields {
		if field.Private {
			return true
		}
	}

	return len(rg.Fields) == 0
}

func (rg recordGenerator) Field(field gir.Field) string {
	if field.Private {
		return ""
	}

	typ := rg.Ng.resolveAnyType(field.AnyType)
	if typ == "" {
		return ""
	}

	name := SnakeToGo(true, field.Name)

	return GoDoc(field.Doc, 1, name) + "\n" + name + " " + typ
}

func (rg recordGenerator) ResolveType(typ gir.Type) {
}

func (ng *NamespaceGenerator) generateRecords() {
	_ = gir.Field{}

	for _, record := range ng.current.Namespace.Records {
		if ignoreRecord(record) {
			continue
		}

		ng.pen.BlockTmpl(recordTmpl, recordGenerator{
			Record: record,
			Ng:     ng,
		})
	}
}

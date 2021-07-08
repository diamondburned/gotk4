package generators

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/types"

	ifacegen "github.com/diamondburned/gotk4/gir/girgen/generators/iface"
)

var classInterfaceTmpl = gotmpl.NewGoTemplate(`
	{{ if .Virtuals }}
	// {{ .InterfaceName }}Overrider contains methods that are overridable.
	//
	// As of right now, interface overriding and subclassing is not supported
	// yet, so the interface currently has no use.
	type {{ .InterfaceName }}Overrider interface {
		{{ range .Virtuals }}
		{{- GoDoc . 1 TrailingNewLine -}}
		{{- .Name }}{{ .Tail }}
		{{ end }}
	}
	{{ end }}

	{{ GoDoc . 0 }}
	type {{ .InterfaceName }} interface {
		gextras.Objector

		{{ range .Methods }}
		{{- GoDoc . 1 TrailingNewLine -}}
		{{- .Name }}{{ .Tail }}
		{{ else }}
		private{{ .StructName }}()
		{{ end }}
	}

	// {{ .StructName }} implements the {{ .InterfaceName }} interface.
	type {{ .StructName }} struct {
		{{ range .Tree.ImplTypes -}}
		{{ . }}
		{{ end }}
	}

	var _ {{ .InterfaceName }} = (*{{ .StructName }})(nil)

	func wrap{{ .InterfaceName }}(obj *externglib.Object) {{ .InterfaceName }} {
		return {{ .Tree.Wrap "obj" }}
	}

	{{ if .GLibGetType }}
	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return wrap{{ .InterfaceName }}(obj), nil
	}
	{{ end }}

	{{ range .Constructors }}
	{{ GoDoc . 0 }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}

	{{ range .Methods }}
	{{ GoDoc . 0 }}
	func ({{ .Recv }} *{{ $.StructName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ else }}
	func (*{{ .StructName }}) private{{ $.StructName }}() {}
	{{ end }}
`)

// GenerateInterface generates a public interface declaration, optionally
// another one for overriding, and the private struct that implements the
// interface specifically for wrapping opaque C interfaces.
func GenerateInterface(gen FileGeneratorWriter, iface *gir.Interface) bool {
	igen := ifacegen.NewGenerator(gen)
	if !igen.Use(iface) {
		return false
	}

	generateInterfaceGenerator(gen, &igen)
	return true
}

// GenerateClass generates the given class into files.
func GenerateClass(gen FileGeneratorWriter, class *gir.Class) bool {
	igen := ifacegen.NewGenerator(gen)
	if !igen.Use(class) {
		return false
	}

	generateInterfaceGenerator(gen, &igen)
	return true
}

func generateInterfaceGenerator(gen FileGeneratorWriter, igen *ifacegen.Generator) {
	writer := FileWriterFromType(gen, igen)
	writer.Header().NeedsExternGLib()
	writer.Header().ImportCore("gextras")

	// Import for implementation types.
	for _, parent := range igen.Tree.Requires {
		writer.Header().ImportImpl(parent.Resolved)
	}

	if igen.GLibGetType != "" && !types.FilterCType(gen, igen.GLibGetType) {
		writer.Header().AddMarshaler(igen.GLibGetType, igen.InterfaceName)
	}

	file.ApplyHeader(writer, igen)
	writer.Pen().WriteTmpl(classInterfaceTmpl, igen)
}

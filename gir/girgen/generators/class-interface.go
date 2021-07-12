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
	// {{ .StructName }}Overrider contains methods that are overridable.
	//
	// As of right now, interface overriding and subclassing is not supported
	// yet, so the interface currently has no use.
	type {{ .StructName }}Overrider interface {
		{{ range .Virtuals -}}
		{{- GoDoc . 1 TrailingNewLine -}}
		{{- .Name }}{{ .Tail }}
		{{ end -}}
	}
	{{ end }}

	// {{ .InterfaceName }} describes {{ .StructName }}'s methods.
	type {{ .InterfaceName }} interface {
		{{ range .Methods -}}
		{{- Synopsis . 1 TrailingNewLine -}}
		{{- .Name }}{{ .Tail }}
		{{ else }}
		private{{ .StructName }}()
		{{ end -}}
	}

	{{ GoDoc . 0 (OverrideSelfName .StructName) }}
	type {{ .StructName }} struct {
		{{ index .Tree.ImplTypes 0 }}

		{{ range (slice .Tree.ImplTypes 1) -}}
		{{ . }}
		{{ end }}
	}

	var (
		_ {{ .InterfaceName }} = (*{{ .StructName }})(nil)
		_ gextras.Nativer = (*{{ .StructName }})(nil)
	)

	{{ $wrapper := .Tree.WrapName false }}
	func {{ $wrapper }}(obj *externglib.Object) *{{ .StructName }} {
		return {{ .Wrap "obj" }}
	}

	{{ if .HasMarshaler }}
	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return {{ $wrapper }}(obj), nil
	}
	{{ end }}

	{{ range .Constructors }}
	{{ GoDoc . 0 }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}

	{{ if .Tree.HasAmbiguousSelector }}
	{{ $field := (.Tree.FirstGObjectSelector "v") }}
	// Native implements gextras.Nativer. It returns the underlying GObject
	// field.
	func (v *{{ $.StructName }}) Native() uintptr {
		return {{ if $field -}}
			{{ $field }}.Native()
		{{- else -}}
			uintptr(unsafe.Pointer(v))
		{{- end }}
	}
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

type ifacegenData struct {
	*ifacegen.Generator
	HasMarshaler bool
}

func generateInterfaceGenerator(gen FileGeneratorWriter, igen *ifacegen.Generator) {
	writer := FileWriterFromType(gen, igen)
	writer.Header().NeedsExternGLib()
	writer.Header().ImportCore("gextras")

	// Import for implementation types.
	for _, parent := range igen.Tree.Requires {
		writer.Header().ImportImpl(parent.Resolved)
	}

	data := ifacegenData{
		Generator:    igen,
		HasMarshaler: false,
	}

	if igen.GLibGetType != "" && !types.FilterCType(gen, igen.GLibGetType) {
		data.HasMarshaler = true
		writer.Header().AddMarshaler(igen.GLibGetType, igen.InterfaceName)
	}

	writer.Pen().WriteTmpl(classInterfaceTmpl, data)
	file.ApplyHeader(writer, igen)
}

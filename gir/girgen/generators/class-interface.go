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
	// {{ .InterfaceName }}Overrider contains methods that are overridable .
	//
	// As of right now, interface overriding and subclassing is not supported
	// yet, so the interface currently has no use.
	type {{ .InterfaceName }}Overrider interface {
		{{ range .Virtuals -}}
		{{ GoDoc . 1 }}
		{{ .Name }}{{ .Tail }}
		{{ end }}
	}
	{{ end }}

	{{ GoDoc . 0 }}
	type {{ .InterfaceName }} interface {
		gextras.Objector

		{{ range .Implements }}
		// As{{.Name}} casts the class to the {{.Type}} interface.
		As{{.Name}}() {{ .Type -}}
		{{ end }}

		{{ range .InheritedMethods }}
		{{ GoDoc . 1 (AdditionalString (printf "This method is inherited from %s" .Parent)) }}
		{{ .Name }}{{ .Tail -}}
		{{ end }}

		{{ range .Methods -}}
		{{ GoDoc . 1 }}
		{{ .Name }}{{ .Tail }}
		{{ end }}
	}

	// {{ .StructName }} implements the {{ .InterfaceName }} interface.
	type {{ .StructName }} struct {
		*externglib.Object
	}

	var _ {{ .InterfaceName }} = (*{{ .StructName }})(nil)

	// Wrap{{ .InterfaceName }} wraps a GObject to a type that implements
	// interface {{ .InterfaceName }}. It is primarily used internally.
	func Wrap{{ .InterfaceName }}(obj *externglib.Object) {{ .InterfaceName }} {
		return {{ .StructName }}{obj}
	}

	{{ if .GLibGetType }}
	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return Wrap{{ .InterfaceName }}(obj), nil
	}
	{{ end }}

	{{ range .Constructors }}
	{{ GoDoc . 0 }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}

	{{ $recv := (FirstLetter .StructName) }}
	{{ range .Implements }}
	func ({{ $recv }} {{ $.StructName }}) As{{.Name}}() {{ .Type }} {
		return {{ .Wrapper }}(gextras.InternObject({{ $recv }}))
	}
	{{ end }}

	{{ range .InheritedMethods }}
	func ({{ .Recv }} {{ $.StructName }}) {{ .Name }}{{ .Tail }} {
		{{ if .Return }} return {{ end -}}
		{{ .Wrapper }}(gextras.InternObject({{ .Recv }})).{{ .Name }}({{ .CallParams }})
	}
	{{ end }}

	{{ range .Methods }}
	func ({{ .Recv }} {{ $.StructName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
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

	writer := FileWriterFromType(gen, iface)
	generateInterfaceGenerator(gen, writer, &igen)
	return true
}

// GenerateClass generates the given class into files.
func GenerateClass(gen FileGeneratorWriter, class *gir.Class) bool {
	igen := ifacegen.NewGenerator(gen)
	if !igen.Use(class) {
		return false
	}

	writer := FileWriterFromType(gen, class)
	generateInterfaceGenerator(gen, writer, &igen)
	return true
}

func generateInterfaceGenerator(gen FileGenerator, writer FileWriter, igen *ifacegen.Generator) {
	writer.Header().NeedsExternGLib()

	if igen.GLibGetType != "" && !types.FilterCType(gen, igen.GLibGetType) {
		writer.Header().AddMarshaler(igen.GLibGetType, igen.InterfaceName)
	}

	file.ApplyHeader(writer, igen)
	writer.Pen().WriteTmpl(classInterfaceTmpl, igen)
}

package generators

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"

	ifacegen "github.com/diamondburned/gotk4/gir/girgen/generators/iface"
)

var classInterfaceTmpl = gotmpl.NewGoTemplate(`
	{{ $wrapper := .Tree.WrapName false }}

	{{ if .GLibTypeStruct }}
	{{ if .IsClass }}
	// {{ .StructName }}Overrides contains methods that are overridable.
	type {{ .StructName }}Overrides struct {
		{{ range .GLibTypeStruct.VirtualMethods -}}
		{{ if $.IsInSameFile . -}}
		{{- GoDoc .Go 1 TrailingNewLine -}}
		{{- .Go.Name }} func{{ .Go.Tail }}
		{{ end -}}
		{{ end -}}
	}

	func default{{ .StructName }}Overrides(v *{{ .StructName }}) {{ .StructName }}Overrides{
		return {{ .StructName }}Overrides{
			{{ range .GLibTypeStruct.VirtualMethods -}}
			{{ .Go.Name }}: v.{{ UnexportPascal .Go.Name }},
			{{ end -}}
		}
	}
	{{ else }}
	// {{ .StructName }}Overrider contains methods that are overridable.
	type {{ .StructName }}Overrider interface {
		{{ range .GLibTypeStruct.VirtualMethods -}}
		{{ if $.IsInSameFile . -}}
		{{- GoDoc .Go 1 TrailingNewLine -}}
		{{- .Go.Name }}{{ .Go.Tail }}
		{{ end -}}
		{{ end -}}
	}
	{{ end }}
	{{ end }}

	{{ GoDoc . 0 (OverrideSelfName .StructName) }}
	{{ if not .IsClass -}}
	//
	// {{ .StructName }} wraps an interface. This means the user can get the
	// underlying type by calling Cast().
	{{ end -}}
	type {{ .StructName }} struct {
		_ [0]func() // equal guard
		{{ index .Tree.ImplTypes 0 }}

		{{ range (slice .Tree.ImplTypes 1) -}}
		{{ . }}
		{{ end }}
	}

	var (
		{{ range .ImplInterfaces -}}
		_ {{ . }} = (*{{ $.StructName }})(nil)
		{{ end }}
	)

	{{ $needsPrivate := false }}

	{{ if .Abstract }}

	{{ if .IsClass }}
	// {{ .InterfaceName }} describes types inherited from class {{ .StructName }}.
	{{ $needsPrivate = true -}}
	//
	// To get the original type, the caller must assert this to an interface or
	// another type.
	type {{ .InterfaceName }} interface {
		coreglib.Objector
		base{{ .StructName }}() *{{ .StructName }}
	}
	{{ else }}
	// {{ .InterfaceName }} describes {{ .StructName }}'s interface methods.
	type {{ .InterfaceName }} interface {
		coreglib.Objector

		{{ if .Methods -}}

		{{ range .Methods }}
		{{ if $.IsInSameFile . -}}
		{{- Synopsis . 1 TrailingNewLine }}
		{{- .Name }}{{ .Tail }}
		{{- end }}
		{{- end }}

		{{ range .Signals }}
		{{ Synopsis . 1 TrailingNewLine }}
		{{- .GoName }}(func{{ .GoTail }}) coreglib.SignalHandle
		{{- end }}

		{{- end}}

		{{ if not .Methods -}}
		{{ $needsPrivate = true -}}
		base{{ .StructName }}() *{{ .StructName }}
		{{ end -}}
	}
	{{ end }}

	var _ {{ .InterfaceName }} = (*{{ .StructName }})(nil)
	{{ end }}

	{{ if .GLibTypeStruct }}
	{{ if .IsClass }}

	{{ Import . "unsafe" }}

	func init() {
		coreglib.RegisterClassInfo[*{{ .StructName }}, *{{ .GLibTypeStruct.Name }}, {{ .StructName }}Overrides](
			GType{{ .StructName }},
			init{{ .GLibTypeStruct.Name }},
			{{ $wrapper }},
			default{{ .StructName }}Overrides,
		)
	}

	func init{{ .GLibTypeStruct.Name }}(gclass unsafe.Pointer, overrides {{ .StructName }}Overrides, classInitFunc func(*{{ .GLibTypeStruct.Name }})) {
		{{- if .GLibTypeStruct.VirtualMethods }}
		{{- if .IsRuntimeLinkMode }}
			{{- ImportCore . "girepository" -}}
			classt := girepository.MustFind({{ Quote .Root.Namespace.Name }}, {{ Quote .GLibTypeStruct.Record.Name }})
			pclass := unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GType{{ .StructName }})))

			{{ range .GLibTypeStruct.VirtualMethods }}
			if overrides.{{ .Go.Name }} != nil {
				o := classt.StructFieldOffset({{ Quote .FieldName }})
				*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(pclass), o)) = unsafe.Pointer(C.{{ .C.Name }})
			}
			{{ end -}}
		{{- else }}
			pclass := (*C.{{ .GLibTypeStruct.CType }})(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GType{{ .StructName }}))))
			{{ range .GLibTypeStruct.VirtualMethods }}
			if overrides.{{ .Go.Name }} != nil {
				pclass.{{ CGoField .FieldName }} = (*[0]byte)(C.{{ .C.Name }})
			}
			{{ end }}
		{{ end -}}
		{{ end -}}

		if classInitFunc != nil {
			{{-  ImportCore . "gextras"  -}}
			class := (*{{ .GLibTypeStruct.Name }})(gextras.NewStructNative(gclass))
			classInitFunc(class)
		}
	}

	{{ else }}
	func ifaceInit{{ .InterfaceName }}(gifacePtr, data C.gpointer) {
		{{- Import . "unsafe" -}}
		{{- if .GLibTypeStruct.VirtualMethods }}

		{{- if .IsRuntimeLinkMode}}
			iface := girepository.MustFind({{ Quote .Root.Namespace.Name }}, {{ Quote .GLibTypeStruct.Record.Name }})
			{{- range .GLibTypeStruct.VirtualMethods}}
			*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset({{ Quote .FieldName }}))) = unsafe.Pointer(C.{{ .C.Name }})
			{{- end }}
		{{- else }}
			iface := (*C.{{ .GLibTypeStruct.CType }})(unsafe.Pointer(gifacePtr))
			{{- range .GLibTypeStruct.VirtualMethods }}
			iface.{{ CGoField .FieldName }} = (*[0]byte)(C.{{ .C.Name }})
			{{- end }}
		{{- end }}

		{{- end }}
	}
	{{ end }}
	{{ end }}

	func {{ $wrapper }}(obj *coreglib.Object) *{{ .StructName }} {
		return {{ .Wrap "obj" }}
	}

	{{ if .HasMarshaler }}
	func marshal{{ .StructName }}(p uintptr) (interface{}, error) {
		{{- Import . "unsafe" -}}
		return {{ $wrapper }}(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
	}
	{{ end }}

	{{ if $needsPrivate }}
	func ({{ .Recv }} *{{ .StructName }}) base{{ .StructName }}() *{{ .StructName }} {
		return {{ .Recv }}
	}

	// Base{{ .StructName }} returns the underlying base object.
	func Base{{ .StructName }}(obj {{ .InterfaceName }}) *{{ .StructName }} {
		return obj.base{{ .StructName }}()
	}
	{{ end }}

	{{ range .Signals }}
	{{ GoDoc . 0 (OverrideSelfName .GoName) }}
	func ({{ $.Recv }} *{{ $.StructName }}) {{ .GoName }}(f func{{ .GoTail }}) coreglib.SignalHandle {
		return coreglib.ConnectGeneratedClosure({{ $.Recv }}, {{ Quote .Name }}, false, unsafe.Pointer(C.{{ .CGoName }}), f)
	}
	{{ end }}
`)

var constructorInterfaceImpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
`)

// methodInterfaceTmpl needs the following type:
//
//    struct {
//        Method
//        StructName string
//    }
//
var methodInterfaceTmpl = gotmpl.NewGoTemplate(`
	{{ with .Method }}
	{{ GoDoc . 0 }}
	func ({{ .Recv }} *{{ $.StructName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

var signalInterfaceTmpl = gotmpl.NewGoTemplate(`
	//export {{ .CGoName }}
	func {{ .CGoName }}{{ .CGoTail }} {{ .Block }}
`)

var virtualExportedInterfaceTmpl = gotmpl.NewGoTemplate(`
	//export {{ .C.Name }}
	func {{ .C.Name }}{{ .C.Tail }} {{ .C.Block }}
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
	ImplInterfaces []string
	HasMarshaler   bool

	gen FileGeneratorWriter
}

func (d ifacegenData) Recv() string {
	if len(d.Methods) > 0 {
		return d.Methods[0].Recv
	}
	return "v"
}

func (d ifacegenData) IsRuntimeLinkMode() bool {
	return d.gen.LinkMode() == types.RuntimeLinkMode
}

func generateInterfaceGenerator(gen FileGeneratorWriter, igen *ifacegen.Generator) {
	writer := FileWriterFromType(gen, igen)
	writer.Header().NeedsExternGLib()
	// TOOD: add gbox

	// Import for implementation types.
	for _, parent := range igen.Tree.Requires {
		parent.Resolved.ImportImpl(gen, writer.Header())
	}

	// These conditions should follow what's in the template.
	if igen.GLibTypeStruct != nil {
		if gen.LinkMode() == types.RuntimeLinkMode && len(igen.GLibTypeStruct.VirtualMethods) > 0 {
			writer.Header().ImportCore("girepository")
		}
	}

	data := ifacegenData{
		Generator:      igen,
		ImplInterfaces: igen.ImplInterfaces(),
		HasMarshaler:   false,
		gen:            gen,
	}

	if gtype, ok := GenerateGType(gen, igen.Name, igen.GLibGetType); ok {
		data.HasMarshaler = true
		writer.Header().AddMarshaler(gtype.GetType, igen.StructName)
	}

	writer.Pen().WriteTmpl(classInterfaceTmpl, data)
	file.ApplyHeader(writer, igen)

	for _, ctor := range igen.Constructors {
		writer := FileWriterFromType(gen, ctor)
		writer.Header().ApplyFrom(&ctor.Header)
		writer.Pen().WriteTmpl(constructorInterfaceImpl, ctor)
	}

	for _, method := range igen.Methods {
		writer := FileWriterFromType(gen, method)
		writer.Header().ApplyFrom(&method.Header)
		writer.Pen().WriteTmpl(methodInterfaceTmpl, gotmpl.M{
			"Method":     method,
			"StructName": igen.StructName,
		})
	}

	for _, virtual := range igen.VirtualMethods {
		writer := FileWriterFromType(gen, virtual)
		writer.Header().ApplyFrom(&virtual.Header)

		// Unexport virtual method calls.
		virtual.Name = strcases.UnexportPascal(virtual.Name)

		writer.Pen().WriteTmpl(methodInterfaceTmpl, gotmpl.M{
			"Method":     virtual,
			"StructName": igen.StructName,
		})
	}

	if igen.GLibTypeStruct != nil {
		for _, virtual := range igen.GLibTypeStruct.VirtualMethods {
			writer := FileWriterExportedFromType(gen, virtual)
			writer.Header().ApplyFrom(virtual.C.Header)
			writer.Pen().WriteTmpl(virtualExportedInterfaceTmpl, virtual)
		}
	}

	for _, signal := range igen.Signals {
		writer := FileWriterExportedFromType(gen, signal)
		writer.Header().ApplyFrom(signal.Header)
		writer.Pen().WriteTmpl(signalInterfaceTmpl, signal)
	}
}

package generators

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// TODO: unexported type implementation
// TODO: methods for implementation

var interfaceTmpl = gotmpl.NewGoTemplate(`
	{{ if .Virtuals }}
	// {{ .InterfaceName }}Overrider contains methods that are overridable. This
	// interface is a subset of the interface {{ .InterfaceName }}.
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
		{{ range .TypeTree.PublicEmbeds -}}
		{{ . }}
		{{ end }}

		{{ range .Methods -}}
		{{ GoDoc . 1 }}
		{{ .Name }}{{ .Tail }}
		{{ end }}
	}

	// {{ .StructName }} implements the {{ .InterfaceName }} interface.
	type {{ .StructName }} struct {
		{{ range .TypeTree.PublicEmbeds -}}
		{{ . }}
		{{ end }}
	}

	var _ {{ .InterfaceName }} = (*{{ .StructName }})(nil)

	// Wrap{{ .InterfaceName }} wraps a GObject to a type that implements
	// interface {{ .InterfaceName }}. It is primarily used internally.
	func Wrap{{ .InterfaceName }}(obj *externglib.Object) {{ .InterfaceName }} {
		return {{ .Wrap "obj" }}
	}

	{{ if .GLibGetType }}
	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return Wrap{{ .InterfaceName }}(obj), nil
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
	igen := NewInterfaceGenerator(gen)
	if !igen.Use(iface) {
		return false
	}

	writer := FileWriterFromType(gen, iface)

	if iface.GLibGetType != "" && !types.FilterCType(gen, iface.GLibGetType) {
		writer.Header().AddMarshaler(iface.GLibGetType, igen.InterfaceName)
	}

	file.ApplyHeader(writer, &igen)
	writer.Pen().WriteTmpl(interfaceTmpl, &igen)
	return true
}

type InterfaceGenerator struct {
	*gir.Interface
	InterfaceName string
	StructName    string

	TypeTree types.Tree
	Virtuals []callable.Generator // for overrider
	Methods  []callable.Generator // for big interface

	source *gir.NamespaceFindResult
	header file.Header
	gen    FileGenerator
}

// NewInterfaceGenerator creates a new interface generator instance.
func NewInterfaceGenerator(gen FileGenerator) InterfaceGenerator {
	return InterfaceGenerator{
		gen:      gen,
		TypeTree: types.NewTree(gen),
	}
}

func (g *InterfaceGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := logger.Prefix(v, fmt.Sprintf("interface %s (C.%s):", g.InterfaceName, g.CType))
	g.gen.Logln(lvl, p)
}

// Reset resets the callback generator.
func (g *InterfaceGenerator) Reset() {
	*g = InterfaceGenerator{
		gen: g.gen,
	}
}

// Header returns the callback generator's current header.
func (g *InterfaceGenerator) Header() *file.Header {
	return &g.header
}

func (g *InterfaceGenerator) Use(iface *gir.Interface) bool {
	g.Reset()

	if !iface.IsIntrospectable() || types.Filter(g.gen, iface.Name, iface.CType) {
		return false
	}

	g.Interface = iface
	g.InterfaceName = strcases.PascalToGo(iface.Name)
	g.StructName = strcases.UnexportPascal(g.InterfaceName)
	g.source = g.gen.Namespace()

	if !g.TypeTree.Resolve(iface.Name) {
		g.Logln(logger.Debug, "cannot be type-resolved")
		return false
	}

	for _, imp := range g.TypeTree.Requires {
		// Import everything for the embedded types inside the interface.
		g.header.ImportPubl(imp.Resolved)
	}

	g.updateMethods()
	return true
}

// UseMethods sets only the VirtualMethods and Methods fields from the given
// interface belonging to the given namespace.. It skips the type resolving
// steps.
func (g *InterfaceGenerator) UseMethods(iface *gir.Interface, n *gir.NamespaceFindResult) {
	g.TypeTree.Reset()

	g.Interface = iface
	g.InterfaceName = strcases.PascalToGo(iface.Name)
	g.StructName = strcases.UnexportPascal(g.InterfaceName)
	g.source = n

	g.updateMethods()
}

func (g *InterfaceGenerator) updateMethods() {
	g.Methods = callable.Grow(g.Methods, len(g.Interface.Methods))
	g.Virtuals = callable.Grow(g.Virtuals, len(g.Interface.VirtualMethods))

	for _, vmethod := range g.Interface.VirtualMethods {
		gen := callable.NewGenerator(headeredFileGenerator{
			FileGenerator: g.gen,
			Headerer:      g,
		})
		if !gen.UseFromNamespace(&vmethod.CallableAttrs, g.source) {
			continue
		}

		file.ApplyHeader(g, &gen)
		g.Virtuals = append(g.Virtuals, gen)
	}

	for _, method := range g.Interface.Methods {
		gen := callable.NewGenerator(headeredFileGenerator{
			FileGenerator: g.gen,
			Headerer:      g,
		})
		if !gen.UseFromNamespace(&method.CallableAttrs, g.source) {
			continue
		}

		file.ApplyHeader(g, &gen)
		g.Methods = append(g.Methods, gen)
	}

	callable.RenameGetters(g.InterfaceName, g.Methods)
	callable.RenameGetters(g.InterfaceName, g.Virtuals)
}

// Wrap returns a wrapper block that wraps around the given *glib.Object
// variable name.
func (g *InterfaceGenerator) Wrap(obj string) string {
	return g.TypeTree.WrapInterface(obj)
}

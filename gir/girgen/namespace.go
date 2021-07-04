package girgen

import (
	"fmt"
	"path"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/generators"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// NamespaceGenerator manages generation of a namespace. A namespace contains
// various files, which are created using the FileWriter method.
type NamespaceGenerator struct {
	*Generator
	PkgPath    string
	PkgName    string
	PkgVersion string

	current *gir.NamespaceFindResult
	files   map[string]*FileGenerator
}

var (
	_ types.FileGenerator            = (*NamespaceGenerator)(nil)
	_ generators.FileGenerator       = (*NamespaceGenerator)(nil)
	_ generators.FileGeneratorWriter = (*NamespaceGenerator)(nil)
)

// NewNamespaceGenerator creates a new NamespaceGenerator from the given
// generator and namespace.
func NewNamespaceGenerator(g *Generator, n *gir.NamespaceFindResult) *NamespaceGenerator {
	return &NamespaceGenerator{
		Generator:  g,
		PkgPath:    g.ModPath(n.Namespace),
		PkgName:    gir.GoNamespace(n.Namespace),
		PkgVersion: gir.MajorVersion(n.Namespace.Version),
		current:    n,
		files:      map[string]*FileGenerator{},
	}
}

// Namespace returns the generator's namespace that includes the repository it's
// in.
func (n *NamespaceGenerator) Namespace() *gir.NamespaceFindResult {
	return n.current
}

func (n *NamespaceGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("package %s/v%s", n.PkgName, gir.MajorVersion(n.current.Namespace.Version))
	n.Generator.Logln(lvl, logger.Prefix(v, p))
}

// CanGenerate checks if a type can be generated or not.
func (n *NamespaceGenerator) CanGenerate(r *types.Resolved) bool {
	if r.Extern == nil {
		return true // built-in
	}

	if !r.Extern.IsIntrospectable() {
		return false
	}

	// TODO: figure out a more optimized way, probably by generating schemas
	// first and having a store of what's already generated. This will do for
	// now.

	stub := generators.StubFileGeneratorWriter(n)

	switch v := r.Extern.Type.(type) {
	case *gir.Alias:
		return generators.GenerateAlias(stub, v)
	case *gir.Bitfield:
		return generators.GenerateBitfield(stub, v)
	case *gir.Callback:
		return generators.GenerateCallback(stub, v)
	case *gir.Enum:
		return generators.GenerateEnum(stub, v)
	case *gir.Function:
		return generators.GenerateFunction(stub, v)
	case *gir.Interface:
		return generators.GenerateInterface(stub, v)
	case *gir.Record:
		return generators.GenerateRecord(stub, v)
	}

	// Default to true.
	return true
}

// Pkgconfig returns the current repository's pkg-config names.
func (n *NamespaceGenerator) Pkgconfig() []string {
	foundRoot := false
	pkgs := make([]string, 0, len(n.current.Repository.Packages)+1)

	for _, pkg := range n.current.Repository.Packages {
		if pkg.Name == n.current.Repository.Pkg {
			foundRoot = true
		}

		pkgs = append(pkgs, pkg.Name)
	}

	if !foundRoot {
		pkgs = append(pkgs, n.current.Repository.Pkg)
	}

	return pkgs
}

// FileWriter returns the respective file writer from the given InfoFields.
func (n *NamespaceGenerator) FileWriter(info cmt.InfoFields) generators.FileWriter {
	if info.Elements == nil {
		return n.makeFile("")
	}

	var filename string

	switch {
	case info.Elements.SourcePosition != nil:
		filename = info.Elements.SourcePosition.Filename
	case info.Elements.Doc != nil:
		filename = info.Elements.Doc.Filename
	default:
		return n.makeFile("")
	}

	if info.Attrs != nil && info.Attrs.Version != "" {
		filename += info.Attrs.Version // ex: gtk3.2.go
	}

	return n.makeFile(swapFileExt(filename, ".go"))
}

// swapFileExt returns the base name of the given filepath with its file
// extension replaced. The given file extension should contain a dot if it's not
// empty.
func swapFileExt(filepath, ext string) string {
	filename := path.Base(filepath)
	return strings.Split(filename, ".")[0] + ext
}

func (n *NamespaceGenerator) makeFile(filename string) *FileGenerator {
	if filename == "" {
		filename = n.PkgName + ".go"
	}

	f, ok := n.files[filename]
	if ok {
		return f
	}

	f = NewFileGenerator(n)
	n.files[filename] = f
	return f
}

// GenerateAll generates everything in the current namespace into files. The
// returned map maps the filename to the raw file content.
func (n *NamespaceGenerator) GenerateAll() map[string][]byte {
	for _, v := range n.current.Namespace.Aliases {
		n.logIfSkipped(generators.GenerateAlias(n, &v), v.Name)
	}
	for _, v := range n.current.Namespace.Enums {
		n.logIfSkipped(generators.GenerateEnum(n, &v), v.Name)
	}
	for _, v := range n.current.Namespace.Bitfields {
		n.logIfSkipped(generators.GenerateBitfield(n, &v), v.Name)
	}
	for _, v := range n.current.Namespace.Callbacks {
		n.logIfSkipped(generators.GenerateCallback(n, &v), v.Name)
	}
	for _, v := range n.current.Namespace.Functions {
		n.logIfSkipped(generators.GenerateFunction(n, &v), v.Name)
	}
	for _, v := range n.current.Namespace.Interfaces {
		n.logIfSkipped(generators.GenerateInterface(n, &v), v.Name)
	}
	for _, v := range n.current.Namespace.Classes {
		n.logIfSkipped(generators.GenerateClass(n, &v), v.Name)
	}
	for _, v := range n.current.Namespace.Records {
		n.logIfSkipped(generators.GenerateRecord(n, &v), v.Name)
	}

}

func (n *NamespaceGenerator) logIfSkipped(skipped bool, what string) {
	if skipped {
		n.Logln(logger.Skip, what)
	}
}

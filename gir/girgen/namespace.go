package girgen

import (
	"fmt"
	"path"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/generators"
	"github.com/diamondburned/gotk4/gir/girgen/generators/iface"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/pkg/errors"
)

// NamespaceGenerator manages generation of a namespace. A namespace contains
// various files, which are created using the FileWriter method.
type NamespaceGenerator struct {
	*Generator
	PkgPath    string
	PkgName    string
	PkgVersion string

	current    *gir.NamespaceFindResult
	files      map[string]*FileGenerator
	canResolve map[string]bool
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
		canResolve: map[string]bool{},
	}
}

// Namespace returns the generator's namespace that includes the repository it's
// in.
func (n *NamespaceGenerator) Namespace() *gir.NamespaceFindResult {
	return n.current
}

func (n *NamespaceGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("package %s/v%s:", n.PkgName, n.PkgVersion)
	n.Generator.Logln(lvl, logger.Prefix(v, p)...)
}

// CanGenerate checks if a type can be generated or not.
func (n *NamespaceGenerator) CanGenerate(r *types.Resolved) bool {
	if r.Extern == nil {
		return true // built-in
	}

	if !r.Extern.IsIntrospectable() {
		return false
	}

	publType := types.GoPublicType(n, r)

	// Cache the output of this to both avoid infinite recursions and improve
	// the performance.
	canResolve, ok := n.canResolve[publType]
	if ok {
		return canResolve
	}

	// Mark the type as resolveable to prevent infinite recursions when the
	// generator functions call CanGenerate on its own.
	n.canResolve[publType] = true

	switch v := r.Extern.Type.(type) {
	// Fast checks.
	case *gir.Alias:
		canResolve = generators.CanGenerateAlias(n, v)
	case *gir.Bitfield:
		canResolve = generators.CanGenerateBitfield(n, v)
	case *gir.Enum:
		canResolve = generators.CanGenerateEnum(n, v)
	case *gir.Record:
		canResolve = generators.CanGenerateRecord(n, v)
	case *gir.Class, *gir.Interface:
		canResolve = iface.CanGenerate(n, v)
	// Slow checks.
	case *gir.Callback:
		canResolve = generators.GenerateCallback(generators.StubFileGeneratorWriter(n), v)
	case *gir.Function:
		canResolve = generators.GenerateFunction(generators.StubFileGeneratorWriter(n), v)
	}

	// Actually store the correct value once we're done.
	n.canResolve[publType] = canResolve

	return canResolve
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
	var isRoot bool
	if filename == "" {
		filename = n.PkgName + ".go"
		isRoot = true
	}

	f, ok := n.files[filename]
	if ok {
		return f
	}

	f = NewFileGenerator(n, filename, isRoot)
	n.files[filename] = f
	return f
}

// Generate generates everything in the current namespace into files. The
// returned map maps the filename to the raw file content.
func (n *NamespaceGenerator) Generate() (map[string][]byte, error) {
	// TODO: constants
	// TODO: unions

	generateFunctions := func(parent string, fns []gir.Function) {
		for _, f := range fns {
			if parent != "" {
				f.Name = parent + "_" + f.Name
			}

			if !generators.GenerateFunction(n, &f) {
				prefix := "function " + f.Name
				if parent != "" {
					prefix = "parent " + parent + " " + prefix
				}

				n.logIfSkipped(false, prefix)
			}
		}
	}

	for _, v := range n.current.Namespace.Aliases {
		n.logIfSkipped(generators.GenerateAlias(n, &v), "alias"+v.Name)
	}
	for _, v := range n.current.Namespace.Enums {
		n.logIfSkipped(generators.GenerateEnum(n, &v), "enum "+v.Name)
		generateFunctions(v.Name, v.Functions)
	}
	for _, v := range n.current.Namespace.Bitfields {
		n.logIfSkipped(generators.GenerateBitfield(n, &v), "bitfield "+v.Name)
		generateFunctions(v.Name, v.Functions)
	}
	for _, v := range n.current.Namespace.Callbacks {
		n.logIfSkipped(generators.GenerateCallback(n, &v), "callback "+v.Name)
	}
	for _, v := range n.current.Namespace.Functions {
		n.logIfSkipped(generators.GenerateFunction(n, &v), "function "+v.Name)
	}
	for _, v := range n.current.Namespace.Interfaces {
		n.logIfSkipped(generators.GenerateInterface(n, &v), "interface "+v.Name)
		generateFunctions(v.Name, v.Functions)
	}
	for _, v := range n.current.Namespace.Classes {
		n.logIfSkipped(generators.GenerateClass(n, &v), "class "+v.Name)
		generateFunctions(v.Name, v.Functions)
	}
	for _, v := range n.current.Namespace.Records {
		n.logIfSkipped(generators.GenerateRecord(n, &v), "record "+v.Name)
		generateFunctions(v.Name, v.Functions)
	}

	files := make(map[string][]byte, len(n.files))

	var firstErr error

	for name, file := range n.files {
		if file.IsEmpty() {
			continue
		}

		b, err := file.Generate()
		files[name] = b

		if err != nil && firstErr == nil {
			firstErr = errors.Wrapf(err, "%s/v%s/%s", n.PkgName, n.PkgVersion, name)
		}
	}

	return files, firstErr
}

func (n *NamespaceGenerator) logIfSkipped(generated bool, what string) {
	if !generated {
		n.Logln(logger.Skip, what)
	}
}

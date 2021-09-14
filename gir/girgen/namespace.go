package girgen

import (
	"fmt"
	"path"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators"
	"github.com/diamondburned/gotk4/gir/girgen/generators/iface"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/pkg/errors"
)

// Postprocessor describes a processor function that modifies a namespace. It is
// called right before files are finalized within the namespace generator.
type Postprocessor func(n *NamespaceGenerator) error

// File is a union type that can either be a *FileGenerator or a *RawFile. See
// the original declaration in file.go.
type File interface {
	Pen() *pen.Pen
	Header() *file.Header
	IsEmpty() bool
	Generate() ([]byte, error)
}

// NamespaceGenerator manages generation of a namespace. A namespace contains
// various files, which are created using the FileWriter method.
type NamespaceGenerator struct {
	*Generator
	PkgPath    string
	PkgName    string
	PkgVersion string

	Files map[string]File

	postprocs  []Postprocessor
	current    *gir.NamespaceFindResult
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
		Files:      map[string]File{},
		current:    n,
		canResolve: map[string]bool{},
	}
}

// AddPostprocessors adds the given list of postprocessors.
func (n *NamespaceGenerator) AddPostprocessors(pps []Postprocessor) {
	n.postprocs = append(n.postprocs, pps...)
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

	// Set the right namespace for the generator.
	var ngen types.FileGenerator = n
	if !r.Extern.NamespaceFindResult.Eq(n.current) {
		ngen = types.OverrideNamespace(n, r.Extern.NamespaceFindResult)
	}

	switch v := r.Extern.Type.(type) {
	// Fast checks.
	case *gir.Alias:
		canResolve = generators.CanGenerateAlias(ngen, v)
	case *gir.Bitfield:
		canResolve = generators.CanGenerateBitfield(ngen, v)
	case *gir.Enum:
		canResolve = generators.CanGenerateEnum(ngen, v)
	case *gir.Record:
		canResolve = generators.CanGenerateRecord(ngen, v)
	case *gir.Class, *gir.Interface:
		canResolve = iface.CanGenerate(ngen, v)
	// Slow checks.
	case *gir.Callback:
		canResolve = generators.GenerateCallback(generators.StubFileGeneratorWriter(ngen), v)
	case *gir.Function:
		canResolve = generators.GenerateFunction(generators.StubFileGeneratorWriter(ngen), v)
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
		return n.MakeFile("")
	}

	var filename string

	switch {
	case info.Elements.SourcePosition != nil:
		filename = info.Elements.SourcePosition.Filename
	case info.Elements.Doc != nil:
		filename = info.Elements.Doc.Filename
	default:
		return n.MakeFile("")
	}

	if info.Attrs != nil && info.Attrs.Version != "" {
		filename += info.Attrs.Version // ex: gtk3.2.go
	}

	return n.MakeFile(swapFileExt(filename, ".go"))
}

// extension replaced. The given file extension should contain a dot if it's not
// empty.
func swapFileExt(filepath, ext string) string {
	filename := path.Base(filepath)
	return strings.Split(filename, ".")[0] + ext
}

// MakeFile makes a new FileGenerator for the given filename or returns an
// existing one.
func (n *NamespaceGenerator) MakeFile(filename string) *FileGenerator {
	var isRoot bool
	if filename == "" {
		filename = n.PkgName + ".go"
		isRoot = true
	}

	f, ok := n.Files[filename]
	if ok {
		return f.(*FileGenerator)
	}

	fgen := NewFileGenerator(n, filename, isRoot)
	n.Files[filename] = fgen
	return fgen
}

// MakeRawFile makes a raw file with the given filename. If filename is empty,
// then the function panics. This is usually used to generate C stubs.
func (n *NamespaceGenerator) MakeRawFile(filename string) *RawFile {
	if filename == "" {
		panic("MakeRawFile: empty filename")
	}

	f, ok := n.Files[filename]
	if ok {
		return f.(*RawFile)
	}

	raw := NewRawFile(filename)
	n.Files[filename] = raw
	return raw
}

// Generate generates everything in the current namespace into files. The
// returned map maps the filename to the raw file content.
func (n *NamespaceGenerator) Generate() (map[string][]byte, error) {
	// TODO: constants
	// TODO: unions

	generateFunctions := func(parent string, fns []gir.Function) {
		for _, f := range fns {
			if !generators.GeneratePrefixedFunction(n, &f, parent) {
				n.logIfSkipped(false, "parent "+parent+" function "+f.Name)
			}
		}
	}

	for _, v := range n.current.Namespace.Constants {
		n.logIfSkipped(generators.GenerateConstant(n, &v), "constant "+v.Name)
	}
	for _, v := range n.current.Namespace.Aliases {
		n.logIfSkipped(generators.GenerateAlias(n, &v), "alias "+v.Name)
	}
	for _, v := range n.current.Namespace.Enums {
		if !generators.GenerateEnum(n, &v) {
			n.logIfSkipped(false, "enum "+v.Name)
			continue
		}
		generateFunctions(v.Name, v.Functions)
	}
	for _, v := range n.current.Namespace.Bitfields {
		if !generators.GenerateBitfield(n, &v) {
			n.logIfSkipped(false, "bitfield "+v.Name)
			continue
		}
		generateFunctions(v.Name, v.Functions)
	}
	for _, v := range n.current.Namespace.Callbacks {
		n.logIfSkipped(generators.GenerateCallback(n, &v), "callback "+v.Name)
	}
	for _, v := range n.current.Namespace.Functions {
		n.logIfSkipped(generators.GenerateFunction(n, &v), "function "+v.Name)
	}
	for _, v := range n.current.Namespace.Interfaces {
		if !generators.GenerateInterface(n, &v) {
			n.logIfSkipped(false, "interface "+v.Name)
			continue
		}
		generateFunctions(v.Name, v.Functions)
	}
	for _, v := range n.current.Namespace.Classes {
		if !generators.GenerateClass(n, &v) {
			n.logIfSkipped(false, "class "+v.Name)
			continue
		}
		generateFunctions(v.Name, v.Functions)
	}
	for _, v := range n.current.Namespace.Records {
		if !generators.GenerateRecord(n, &v) {
			n.logIfSkipped(false, "record "+v.Name)
			continue
		}
		generateFunctions(v.Name, v.Functions)
	}

	for _, postproc := range n.postprocs {
		if err := postproc(n); err != nil {
			return nil, err
		}
	}

	// Add the required imports into the file.
	for _, file := range n.Files {
		for _, incl := range n.current.Repository.CIncludes {
			file.Header().IncludeC(incl.Name)
		}
	}

	// Move all C headers to a separate file and import that 1 file.
	headerers := make([]file.Headerer, 0, len(n.Files))
	for _, file := range n.Files {
		headerers = append(headerers, file)
	}

	c := n.MakeRawFile("stubs.h")
	c.pen.WriteString(file.AggregateCgoStubs(headerers...))

	files := make(map[string][]byte, len(n.Files))

	var firstErr error

	for name, file := range n.Files {
		if file.IsEmpty() {
			continue
		}

		if file != c && !c.IsEmpty() {
			// Include the pkgname.h file in all Go files.
			file.Header().IncludeLocalC("stubs.h")
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

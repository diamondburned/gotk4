package girgen

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/generators"
	"github.com/diamondburned/gotk4/gir/girgen/generators/iface"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/pkg/errors"
)

// Postprocessor describes a processor function that modifies a namespace. It is
// called right before files are finalized within the namespace generator.
type Postprocessor func(n *NamespaceGenerator) error

// NamespaceGenerator manages generation of a namespace. A namespace contains
// various files, which are created using the FileWriter method.
type NamespaceGenerator struct {
	*Generator
	PkgPath    string
	PkgName    string
	PkgVersion string

	Files map[string]FileGenerator

	postprocs []Postprocessor
	current   *gir.NamespaceFindResult
	c         struct {
		h *CFileGenerator
		c *CFileGenerator
	}
	canResolve map[string]bool
	genMode    types.LinkMode
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
		Files:      map[string]FileGenerator{},
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
	case *gir.Union:
		canResolve = generators.GenerateUnion(generators.StubFileGeneratorWriter(ngen), v)
	}

	// Actually store the correct value once we're done.
	n.canResolve[publType] = canResolve

	return canResolve
}

// LinkMode implements FileGenerator.
func (n *NamespaceGenerator) LinkMode() types.LinkMode { return n.genMode }

// SetLinkMode sets the link mode for the current namespace and all its files.
// The default is RuntimeLinkMode.
func (n *NamespaceGenerator) SetLinkMode(mode types.LinkMode) {
	n.genMode = mode
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
func (n *NamespaceGenerator) FileWriter(info cmt.InfoFields, export bool) generators.FileWriter {
	if n.Generator.Opts.SingleFile || info.Elements == nil {
		if export {
			return n.MakeFile(n.PkgName + "_export.go")
		}
		return n.MakeFile("")
	}

	var filename string

	switch {
	case info.Elements.SourcePosition != nil:
		filename = info.Elements.SourcePosition.Filename
	case info.Elements.Doc != nil:
		filename = info.Elements.Doc.Filename
	default:
		if export {
			return n.MakeFile(n.PkgName + "_export.go")
		}
		return n.MakeFile("")
	}

	filename = filepath.Base(filename)

	if ext := filepath.Ext(filename); ext != "" {
		filename = strings.TrimSuffix(filename, ext)
	}

	if info.Attrs != nil && info.Attrs.Version != "" {
		filename += "_" + strings.ReplaceAll(info.Attrs.Version, ".", "_") // ex: gtk_3_2.go
	}

	if export {
		filename += "_export"
	}

	return n.MakeFile(filename + ".go")
}

// File gets an existing Go file but returns false if no such file exists. It's
// useful for postprocessors to check if generation is working as intended. If
// SingleFile is true, then File will always return the same file.
func (n *NamespaceGenerator) File(filename string) (*GoFileGenerator, bool) {
	if n.Generator.Opts.SingleFile || filename == "" {
		f, ok := n.Files[n.PkgName+".go"]
		if ok {
			goFile, ok := f.(*GoFileGenerator)
			return goFile, ok
		}
	}

	f, ok := n.Files[filename]
	if ok {
		goFile, ok := f.(*GoFileGenerator)
		return goFile, ok
	}
	return nil, false
}

// MakeFile makes a new GoFileGenerator for the given filename or returns an
// existing one.
func (n *NamespaceGenerator) MakeFile(filename string) *GoFileGenerator {
	// this should lead us down the right branch
	if n.Generator.Opts.SingleFile {
		filename = ""
	}

	isRoot := filename == n.PkgName+".go"

	if filename == "" {
		filename = n.PkgName + ".go"
		isRoot = true
	}

	f, ok := n.Files[filename]
	if ok {
		return f.(*GoFileGenerator)
	}

	goFile := NewGoFileGenerator(n, filename, isRoot)
	n.Files[filename] = goFile
	return goFile
}

func (n *NamespaceGenerator) ch() *CFileGenerator {
	if n.c.h == nil {
		n.c.h = NewCFileGenerator(n, n.PkgName+".h")
	}
	return n.c.h
}

func (n *NamespaceGenerator) cc() *CFileGenerator {
	if n.c.c == nil {
		n.c.c = NewCFileGenerator(n, n.PkgName+".c")
	}
	return n.c.c
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
	for _, v := range n.current.Namespace.Unions {
		if !generators.GenerateUnion(n, &v) {
			n.logIfSkipped(false, "union "+v.Name)
			continue
		}
		generateFunctions(v.Name, v.Functions)
	}

	// Ensure that all files explicitly import runtime/cgo to not trigger an
	// error in a compiler complaining about implicitly importing runtime/cgo.
	// https://sourcegraph.com/github.com/golang/go/-/blob/src/cmd/link/internal/ld/lib.go?L563:3.

	// We also ensure that a root file must exist so the pkg-config headers can
	// be written properly for that package.

	root := n.MakeFile("")
	root.Header().DashImport("runtime/cgo")

	for _, postproc := range n.postprocs {
		if err := postproc(n); err != nil {
			return nil, err
		}
	}

	files := make(map[string][]byte, len(n.Files))

	var firstErr error
	doFile := func(file FileGenerator) {
		b, err := file.Generate()
		files[file.Name()] = b

		if err != nil && firstErr == nil {
			firstErr = errors.Wrapf(err, "%s/v%s/%s", n.PkgName, n.PkgVersion, file.Name())
		}
	}

	for _, file := range n.Files {
		if file.IsEmpty() {
			continue
		}
		doFile(file)
	}

	if !n.c.c.IsEmpty() {
		doFile(n.cc())
	}

	if !n.c.h.IsEmpty() || !n.c.c.IsEmpty() {
		doFile(n.ch())
	}

	return files, firstErr
}

func (n *NamespaceGenerator) logIfSkipped(generated bool, what string) {
	if !generated {
		n.Logln(logger.Skip, what)
	}
}

package girgen

import (
	"path"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/pkg/errors"
)

// NamespaceGenerator is a generator for a specific namespace.
type NamespaceGenerator struct {
	files []FileGenerator

	gen     *Generator
	current *gir.NamespaceFindResult
	pkgPath string // package name
}

// Generate generates the current namespace. It returns a filesystem consisting
// of only files. For correctness, the caller should WalkDir at root.
func (ng *NamespaceGenerator) Generate() (map[string][]byte, error) {
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	ng.generateAliases()
	ng.generateEnums()
	ng.generateBitfields()
	ng.generateCallbacks()
	ng.generateFuncs()
	ng.generateIfaces()
	ng.generateClasses()
	ng.generateRecords()

	files := make(map[string][]byte, len(ng.files))

	for _, file := range ng.files {
		b, err := file.generate()
		files[file.name] = b

		if err != nil {
			return files, errors.Wrap(err, "package "+ng.PackageName())
		}
	}

	return files, nil
}

// FileFromSource returns the respective file from the given SourcePosition. If
// nil is given, the original file is returned.
func (ng *NamespaceGenerator) FileFromSource(pos *gir.SourcePosition) *FileGenerator {
	if pos == nil {
		return ng.file(ng.PackageName() + ".go")
	}

	fg := ng.file(replaceExt(path.Base(pos.Filename), ".go"))
	return fg
}

// file returns the respective file from the given filename.
func (ng *NamespaceGenerator) file(goFile string) *FileGenerator {
	for i, file := range ng.files {
		if file.name == goFile {
			return &ng.files[i]
		}
	}

	i := len(ng.files)
	ng.files = append(ng.files, *NewFileGenerator(ng, goFile))
	return &ng.files[i]
}

func (ng *NamespaceGenerator) mustIgnoreAny(any gir.AnyType) bool {
	switch {
	case any.Type != nil:
		return ng.mustIgnore(any.Type.Name, any.Type.CType)
	case any.Array != nil:
		return ng.mustIgnoreAny(any.Array.AnyType)
	default:
		return true
	}
}

// mustIgnore checks the generator's filters to see if the given girType in this
// namespace should be ignored.
func (ng *NamespaceGenerator) mustIgnore(girType, cType string) (ignore bool) {
	girType = ensureNamespace(ng.Namespace(), girType)

	for _, filter := range ng.gen.Filters {
		if !filter.Filter(ng, girType, cType) {
			// Filter returns keep=false.
			ng.Logln(LogDebug, "ignoring", girType)
			return true
		}
	}

	return false
}

// mustIgnoreC is similar to mustIgnore but only works on C types.
func (ng *NamespaceGenerator) mustIgnoreC(cType string) (ignore bool) {
	return ng.mustIgnore("\x00", cType)
}

// fullGIR returns the full GIR type name if it doesn't contain a namespace.
func (ng *NamespaceGenerator) fullGIR(girType string) string {
	// Skip builtin types.
	_, isBuiltin := girToBuiltin[girType]
	if isBuiltin {
		return girType
	}

	if !strings.Contains(girType, ".") {
		return ng.current.Namespace.Name + "." + girType
	}
	return girType
}

// pkgconfig returns the current repository's pkg-config names.
func (ng *NamespaceGenerator) pkgconfig() []string {
	foundRoot := false
	pkgs := make([]string, 0, len(ng.current.Repository.Packages)+1)

	for _, pkg := range ng.current.Repository.Packages {
		if pkg.Name == ng.current.Repository.Pkg {
			foundRoot = true
		}

		pkgs = append(pkgs, pkg.Name)
	}

	if !foundRoot {
		pkgs = append(pkgs, ng.current.Repository.Pkg)
	}

	return pkgs
}

// PackageName returns the current namespace's package name.
func (ng *NamespaceGenerator) PackageName() string {
	return gir.GoNamespace(ng.current.Namespace)
}

// Namespace returns the generator's namespace that includes the repository it's
// in.
func (ng *NamespaceGenerator) Namespace() *gir.NamespaceFindResult {
	return ng.current
}

// Repositories returns all known repositories outside of this namespace
// generator.
func (ng *NamespaceGenerator) Repositories() gir.Repositories {
	return ng.gen.Repos
}

func (ng *NamespaceGenerator) Logln(level LogLevel, v ...interface{}) {
	prefix := []interface{}{"package", ng.current.Namespace.Name + ":"}
	prefix = append(prefix, v...)

	ng.gen.Logln(level, prefix...)
}

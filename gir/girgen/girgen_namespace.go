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

	// Scan to see if there's a file that needs CallbackDelete. We should put
	// its definition in the root file, if possible.
	for _, file := range ng.files {
		if file.CallbackDelete {
			root := ng.FileFromSource(gir.DocElements{})
			root.CallbackDelete = true
			break
		}
	}

	files := make(map[string][]byte, len(ng.files))

	for _, file := range ng.files {
		b, err := file.generate()
		files[file.name] = b

		if err != nil {
			pkg := ng.PackageName() + "/v" + gir.MajorVersion(ng.current.Namespace.Version)
			return files, errors.Wrap(err, "package "+pkg)
		}
	}

	return files, nil
}

// FileFromSource returns the respective file from the given SourcePosition. If
// nil is given, the original file is returned.
func (ng *NamespaceGenerator) FileFromSource(doc gir.DocElements) *FileGenerator {
	var filename string

	switch {
	case doc.SourcePosition != nil:
		filename = doc.SourcePosition.Filename
	case doc.Doc != nil:
		filename = doc.Doc.Filename
	default:
		filename = ng.PackageName()
	}

	fg := ng.file(replaceExt(path.Base(filename), ".go"))
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

// mustIgnore checks the generator's filters to see if the given girType in this
// namespace should be ignored.
func (ng *NamespaceGenerator) mustIgnore(girName, cType *string) (ignore bool) {
	girType := ensureNamespace(ng.Namespace(), *girName)
	hadNamespace := girType == *girName

	names := FilterTypeName{girType, *cType}

	for _, filter := range ng.gen.Filters {
		// Filter returns keep=false.
		if !filter.Filter(ng, &names) {
			ng.Logln(LogDebug, "ignoring", girType)
			return true
		}
	}

	if hadNamespace {
		*girName = names.GIRType
	} else {
		*girName = names.Name()
	}

	*cType = names.CType

	return false
}

// mustIgnoreC is similar to mustIgnore but only works on C types.
func (ng *NamespaceGenerator) mustIgnoreC(cType string) (ignore bool) {
	nul := "\x00"
	return ng.mustIgnore(&nul, &cType)
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
	v = append(v, nil)
	copy(v[1:], v) // shift rightwards once
	v[0] = "package " + gir.VersionedNamespace(ng.current.Namespace) + ":"

	ng.gen.Logln(level, v...)
}

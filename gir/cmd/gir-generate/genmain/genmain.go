package genmain

import (
	"flag"
	"log"
	"os"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

var (
	Output  string
	Verbose bool
	ListPkg bool
	CgoLink bool
)

func init() {
	flag.StringVar(&Output, "o", "", "output directory to mkdir in")
	flag.BoolVar(&Verbose, "v", Verbose, "log verbosely (debug mode)")
	flag.BoolVar(&ListPkg, "l", ListPkg, "only list packages and exit")
	flag.BoolVar(&CgoLink, "cgo-link", CgoLink, "generate everything to link using cgo instead of girepository")
}

// ParseFlag calls flag.Parse() and initializes external global options.
func ParseFlag() {
	flag.Parse()

	if !ListPkg && Output == "" {
		log.Fatalln("Missing -o output directory.")
	}

	if Verbose {
		girgen.DefaultOpts.LogLevel = logger.Debug
	}
}

type Package struct {
	// Name is the pkg-config name.
	Name string
	// Namespaces is the possible namespaces within it. Refer to
	// ./cmd/gir_namespaces.
	Namespaces []string
}

// HasNamespace returns true if the package allows all namespaces or has the
// given namespace in the list.
func (pkg *Package) HasNamespace(n *gir.Namespace) bool {
	if pkg.Namespaces == nil {
		return true
	}

	namespace := gir.VersionedNamespace(n)
	for _, name := range pkg.Namespaces {
		if name == namespace {
			return true
		}
	}

	return false
}

// Data contains generation data that genmain uses to generate.
type Data struct {
	// Module is the Go Module name that the generator is running for. An
	// example is "github.com/diamondburned/gotk4/pkg".
	Module string
	// Packages lists pkg-config packages and optionally the namespaces to be
	// generated. If the list of namespaces is nil, then everything is
	// generated.
	Packages []Package
	// KnownPackages is similar to Packages, but no packages in this list will
	// be used to generate code. This list automatically includes Packages.
	KnownPackages []Package
	// ImportOverrides is the list of imports to defer to another library,
	// usually because it's tedious or impossible to generate.
	//
	// Not included: coreglib (gotk3/gotk3/glib).
	ImportOverrides map[string]string
	// ExternOverrides adds into ImportOverrides packages that were generated
	// from the given GIR repositories, with the map key being the Go module
	// root for those packages. It internally invokes LoadExternOverrides.
	ExternOverrides map[string]gir.Repositories
	// PkgExceptions contains a list of file names that won't be deleted off of
	// pkg/.
	PkgExceptions []string
	// GenerateExceptions contains the keys of the underneath ImportOverrides
	// map.
	GenerateExceptions []string
	// PkgGenerated contains a list of file names that are packages generated
	// using the given Packages list. It is manually updated.
	PkgGenerated []string
	// Preprocessors defines a list of preprocessors that the main generator
	// will use. It's mostly used for renaming colliding types/identifiers.
	Preprocessors []types.Preprocessor
	// Postprocessors is a map of versioned namespace names to a list of
	// functions that are called to modify any file before it is written out.
	Postprocessors map[string][]girgen.Postprocessor
	// ExtraGoContents contains the contents of files that are appended into
	// generated outputs. It is used to add custom implementations of missing
	// functions. It is a simpler version of Postprocessors.
	ExtraGoContents map[string]string
	// Filters defines a list of GIR types to be filtered. The map key is the
	// namespace, and the values are list of names.
	Filters []types.FilterMatcher
	// ProcessConverters is a list of things that can override a type converter.
	ProcessConverters []typeconv.ConversionProcessor
	// DynamicLinkNamespaces lists namespaces that should be generated directly
	// using Cgo. It includes important core packages as well as packages that
	// are small but performance-sensitive.
	DynamicLinkNamespaces []string
}

// Overlay joins the given list of data into a single Data. The last Data in the
// list will be used for generation.
func Overlay(data ...Data) Data {
	overlay := data[:len(data)-1]
	last := data[len(data)-1]

	for _, datum := range overlay {
		if last.ExternOverrides == nil {
			last.ExternOverrides = make(map[string]gir.Repositories)
		}

		last.KnownPackages = append(last.KnownPackages, datum.Packages...)
		last.ExternOverrides[datum.Module] = MustLoadPackages(datum.Packages)
		last.Preprocessors = append(last.Preprocessors, datum.Preprocessors...)
		last.Filters = append(last.Filters, datum.Filters...)
		last.ProcessConverters = append(last.ProcessConverters, datum.ProcessConverters...)
		last.DynamicLinkNamespaces = append(last.DynamicLinkNamespaces, datum.DynamicLinkNamespaces...)
	}

	return last
}

// Run runs the application.
func Run(data Data) {
	ParseFlag()

	repos := MustLoadPackages(data.Packages)
	MustAddPackages(&repos, data.KnownPackages)
	PrintAddedPkgs(repos)

	if ListPkg {
		return
	}

	Generate(repos, data)
}

// Generate generates the packages based on the given data.
func Generate(repos gir.Repositories, data Data) {
	overrides := data.ImportOverrides
	if overrides == nil {
		overrides = map[string]string{}
	}

	for mod, extern := range data.ExternOverrides {
		for k, v := range LoadExternOverrides(mod, extern) {
			overrides[k] = v
		}
	}

	gen := girgen.NewGenerator(repos, ModulePath(data.Module, overrides))
	gen.Logger = log.New(os.Stderr, "girgen: ", log.Lmsgprefix)
	gen.ApplyPreprocessors(data.Preprocessors)
	gen.AddPostprocessors(data.Postprocessors)
	gen.AddFilters(data.Filters)
	gen.AddProcessConverters(data.ProcessConverters)

	if !CgoLink {
		gen.DynamicLinkNamespaces(data.DynamicLinkNamespaces)
	}

	if err := CleanDirectory(Output, data.PkgExceptions); err != nil {
		log.Fatalln("failed to clean output directory:", err)
	}

	genErrs := GeneratePackages(gen, Output, data.Packages, data.GenerateExceptions)
	if len(genErrs) > 0 {
		for _, err := range genErrs {
			log.Println("generation error:", err)
		}
		os.Exit(1)
	}

	if err := AppendGoFiles(Output, data.ExtraGoContents); err != nil {
		log.Fatalln("failed to append files post-generation:", err)
	}

	finalFiles := [][]string{data.PkgExceptions, data.PkgGenerated}
	if err := EnsureDirectory(Output, finalFiles...); err != nil {
		log.Fatalln("error verifying generation:", err)
	}
}

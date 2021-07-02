package types

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
)

// ModulePathFunc returns the Go module import path from the given namespace.
// See Generator.ModPath for mroe information.
type ModulePathFunc func(*gir.Namespace) string

// FileGenerator defines a generator instance.
type FileGenerator interface {
	logger.LineLogger
	// Filters returns the list of matchers that the current generator has.
	Filters() []FilterMatcher
	// ModPath crafts an import path from the given GIR namespace. The import
	// path is assumed to have the same package name as the base file, but major
	// versions are exempted as an edge case.
	ModPath(*gir.Namespace) string
	// Repositories returns the list of known repositories inside the generator.
	Repositories() gir.Repositories
	// Namespace returns the generator's current namespace.
	Namespace() *gir.NamespaceFindResult
}

type wrappedGenerator struct {
	FileGenerator
	n *gir.NamespaceFindResult
}

func (w wrappedGenerator) Namespace() *gir.NamespaceFindResult { return w.n }

// OverrideNamespace returns a new generator that overrides a generator's
// current namespace.
func OverrideNamespace(gen FileGenerator, nsp *gir.NamespaceFindResult) FileGenerator {
	return wrappedGenerator{gen, nsp}
}

// Find finds the given GIR type from the given generator.
func Find(gen FileGenerator, girType string) *gir.TypeFindResult {
	return gen.Repositories().FindType(gen.Namespace(), girType)
}

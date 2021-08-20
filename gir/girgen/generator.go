package girgen

import (
	"log"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

// Generator is a big generator that manages multiple repositories.
type Generator struct {
	Logger   *log.Logger
	LogLevel logger.Level

	repos   gir.Repositories
	modPath types.ModulePathFunc

	postmap   map[string][]Postprocessor
	filters   []types.FilterMatcher
	convProcs []typeconv.ConversionProcessor
}

// NewGenerator creates a new generator with sane defaults.
func NewGenerator(repos gir.Repositories, modPath types.ModulePathFunc) *Generator {
	return &Generator{
		LogLevel: logger.Skip,

		repos:   repos,
		modPath: modPath,
		postmap: make(map[string][]Postprocessor),
		filters: append([]types.FilterMatcher(nil), types.BuiltinHandledTypes...),
	}
}

// AddFilters adds the given list of filters.
func (g *Generator) AddFilters(filters []types.FilterMatcher) {
	g.filters = append(g.filters, filters...)
}

// Filters returns the generator's list of type filters.
func (g *Generator) Filters() []types.FilterMatcher {
	return g.filters
}

// AddPostprocessor registers the given postprocessors inside a map that has
// keys matching the namespace.
func (g *Generator) AddPostprocessors(ppMap map[string][]Postprocessor) {
	for k, v := range ppMap {
		g.postmap[k] = append(g.postmap[k], v...)
	}
}

// AddProcessConverters adds the given list of conversion processors.
func (g *Generator) AddProcessConverters(processors []typeconv.ConversionProcessor) {
	g.convProcs = append(g.convProcs, processors...)
}

// ProcessConverter satisfies the typeconv.ConversionProcessor interface.
func (g *Generator) ProcessConverter(converter *typeconv.Converter) {
	for _, proc := range g.convProcs {
		proc.ProcessConverter(converter)
	}
}

// AddPreprocessors applies the given list of preprocessors.
func (g *Generator) ApplyPreprocessors(preprocs []types.Preprocessor) {
	types.ApplyPreprocessors(g.repos, preprocs)
}

// ModPath creates an import path from the user's ModulePathFunc given into the
// constructor.
func (g *Generator) ModPath(n *gir.Namespace) string { return g.modPath(n) }

// Repositories returns the generator's repositories.
func (g *Generator) Repositories() gir.Repositories { return g.repos }

// UseNamespace creates a new namespace generator using the given namespace.
func (g *Generator) UseNamespace(namespace, version string) *NamespaceGenerator {
	versioned := gir.VersionedName(namespace, version)

	res := g.repos.FindNamespace(versioned)
	if res == nil {
		return nil
	}

	nsgen := NewNamespaceGenerator(g, res)

	if pps, ok := g.postmap[versioned]; ok {
		nsgen.AddPostprocessors(pps)
	}

	return nsgen
}

// Logln writes a log line into the internal logger.
func (g *Generator) Logln(lvl logger.Level, v ...interface{}) {
	logger.Stdlog(g.Logger, g.LogLevel, lvl, v...)
}

package girgen

import (
	"log"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// Generator is a big generator that manages multiple repositories.
type Generator struct {
	Logger   *log.Logger
	LogLevel logger.Level

	repos   gir.Repositories
	modPath types.ModulePathFunc
	filters []types.FilterMatcher
}

// NewGenerator creates a new generator with sane defaults.
func NewGenerator(repos gir.Repositories, modPath types.ModulePathFunc) *Generator {
	return &Generator{
		LogLevel: logger.Skip,

		repos:   repos,
		modPath: modPath,
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

// ModPath creates an import path from the user's ModulePathFunc given into the
// constructor.
func (g *Generator) ModPath(n *gir.Namespace) string { return g.modPath(n) }

// Repositories returns the generator's repositories.
func (g *Generator) Repositories() gir.Repositories { return g.repos }

// UseNamespace creates a new namespace generator using the given namespace.
func (g *Generator) UseNamespace(namespace, version string) *NamespaceGenerator {
	res := g.repos.FindNamespace(gir.VersionedName(namespace, version))
	if res == nil {
		return nil
	}

	return NewNamespaceGenerator(g, res)
}

// Logln writes a log line into the internal logger.
func (g *Generator) Logln(lvl logger.Level, v ...interface{}) {
	logger.Stdlog(g.Logger, g.LogLevel, lvl, v)
}

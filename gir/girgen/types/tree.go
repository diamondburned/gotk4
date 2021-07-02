package types

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/diamondburned/gotk4/core/pen"
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
)

// Tree is a structure for a type that is resolved to the lowest level of
// inheritance.
type Tree struct {
	*Resolved
	gen FileGenerator

	// Requires contains the direct dependencies of the current type. It may
	// contain interfaces that are also in other interfaces, which will not
	// build.
	Requires []Tree
	// Embeds contains the filtered dependencies, that is, ones that will not
	// collide when built.
	Embeds []Tree

	// Level sets the maximum recursion level to go. It only applies if set
	// to something more than -1.
	Level int
}

// NewTree creates a new empty type tree for resolving.
func NewTree(gen FileGenerator) Tree {
	return Tree{gen: gen, Level: -1}
}

func (tree *Tree) Reset() {
	// Zero out the fields to prevent dangling pointers.
	for i := range tree.Requires {
		tree.Requires[i] = Tree{}
	}

	tree.Resolved = nil
	tree.Embeds = tree.Embeds[:0]
	tree.Requires = tree.Requires[:0]
}

// Resolve resolves the given toplevel type into the Tree, overriding the
// Resolved and Requires fields. True is returned if the tree is successfully
// resolved.
func (tree *Tree) Resolve(toplevel string) bool {
	resolved := ResolveName(tree.gen, toplevel)
	if resolved == nil {
		tree.Reset()
		return false
	}

	return tree.ResolveFromType(resolved)
}

// ResolveFromType is like Resolve, but the caller directly supplies the
// resolved top-level type.
func (tree *Tree) ResolveFromType(toplevel *Resolved) bool {
	tree.Reset()
	tree.Resolved = toplevel

	if tree.Level == 0 {
		// No omit, since we added nothing.
		return true
	}

	// Edge cases for builtin types.
	if tree.Resolved.Builtin != nil {
		switch {
		case toplevel.IsExternGLib("InitiallyUnowned"):
			return tree.resolveParents(externGLibType("*Object", gir.Type{}, "GObject*"))
		case toplevel.IsExternGLib("Object"):
			return true
		}

		tree.omitAmbiguous(false)
		return true
	}

	if !tree.Resolved.Extern.IsIntrospectable() {
		return false
	}

	switch v := tree.Resolved.Extern.Type.(type) {
	case *gir.Class:
		// All classes have a GObject parent, so an empty parent is invalid.
		parent := v.Parent
		if parent == "" {
			parent = "GObject.Object"
		}

		// Resolving the parent type is crucial to make the class working, so if
		// this fails, halt and bail.
		if !tree.resolveName(parent) {
			tree.gen.Logln(logger.Debug, "can't resolve parent", parent, "for class", v.Name)
			return false
		}

		for _, impl := range v.Implements {
			if !tree.resolveName(impl.Name) {
				tree.gen.Logln(logger.Debug, "can't resolve impl", impl.Name, "for class", v.Name)
			}
		}

		tree.omitAmbiguous(true)

	case *gir.Interface:
		for _, prereq := range v.Prerequisites {
			// Like class parents, interface prerequisites are important.
			if !tree.resolveName(prereq.Name) {
				tree.gen.Logln(logger.Debug,
					"can't resolve prerequisite", prereq.Name, "for interface", v.Name)
				return false
			}
		}

		if len(tree.Requires) == 0 {
			// All interfaces are derived from GObjects, so we override the list
			// if it's empty.
			if !tree.resolveParents(externGLibType("*Object", gir.Type{}, "GObject*")) {
				tree.gen.Logln(logger.Debug,
					"can't resolve fallback prerequisite *GObject for interface", v.Name)
				return false
			}
		}

		tree.omitAmbiguous(true)

	default:
		tree.omitAmbiguous(false)
	}

	return true
}

// omitAmbiguous omits current-level ambiguous types.
func (tree *Tree) omitAmbiguous(actually bool) {
	// Try and reuse the backing array, but regrow it ourselves if there's not
	// enough space.
	tree.Embeds = tree.Embeds[:0]
	if cap(tree.Embeds) < len(tree.Requires) {
		tree.Embeds = make([]Tree, 0, len(tree.Requires))
	}

	// No ambiguity if there's less than 2 parents.
	if !actually || len(tree.Requires) < 2 {
		tree.Embeds = append(tree.Embeds, tree.Requires...)
		return
	}

	// This loop seems to be quite expensive. We can likely make it much
	// faster by using a hashmap of prereq names at the current level.
addLoop:
	for i, prereqTree := range tree.Requires {
		for j, otherTree := range tree.Requires {
			// Skip the same value.
			if i == j {
				continue
			}

			if requiresHasGType(otherTree.Requires, prereqTree.Resolved) {
				// prereqTree is already inside another interface that we
				// inherit from, so we don't add it to prevent ambiguity.
				continue addLoop
			}
		}

		tree.Embeds = append(tree.Embeds, prereqTree)
	}
}

// requiresHasGType scans over the list of resolved types and returns true if
// any of the resolved types implements the given resolved type.
func requiresHasGType(requires []Tree, typ *Resolved) bool {
	return requiresHasGTyperec(requires, typ.PublicType(true))
}

func requiresHasGTyperec(requires []Tree, publType string) bool {
	for _, req := range requires {
		if req.Resolved.PublicType(true) == publType {
			return true
		}
		if requiresHasGTyperec(req.Requires, publType) {
			return true
		}
	}
	return false
}

func (tree *Tree) parentLevel() int {
	if tree.Level <= 0 {
		return tree.Level
	}
	return tree.Level - 1
}

// resolveName resolves and adds the resolved type into the Tree.
func (tree *Tree) resolveName(name string) bool {
	parent := Tree{
		gen:   tree.gen,
		Level: tree.parentLevel(),
	}

	if !parent.Resolve(name) {
		return false
	}

	tree.Requires = append(tree.Requires, parent)
	return true
}

// resolveParents manually adds the given parents and resolve them to be added
// into the Tree.
func (tree *Tree) resolveParents(parents ...*Resolved) bool {
	for _, parent := range parents {
		parentTree := Tree{
			gen:   tree.gen,
			Level: tree.parentLevel(),
		}

		if !parentTree.ResolveFromType(parent) {
			// This shouldn't happen, unless the parent type made above is
			// invalid.
			return false
		}

		tree.Requires = append(tree.Requires, parentTree)
	}

	return true
}

// PublicEmbeds returns the list of the toplevel type's children as Go exported
// type names. The namespaces are appropriately prepended if needed.
func (tree *Tree) PublicEmbeds() []string {
	names := make([]string, len(tree.Embeds))

	for i, req := range tree.Embeds {
		namespace := req.Resolved.NeedsNamespace(tree.gen.Namespace())
		names[i] = req.Resolved.PublicType(namespace)
	}

	return names
}

// WalkPublInterfaces walks the tree for all embedded interfaces.
func (tree *Tree) WalkPublInterfaces(f func(*Resolved)) {
	for _, impl := range tree.Embeds {
		if impl.IsInterface() {
			f(impl.Resolved)
		}

		if len(impl.Embeds) > 0 {
			impl.WalkPublInterfaces(f)
		}
	}
}

// WrapClass creates a wrapper that uses public fields to create code that wraps
// the type tree to the top-level type. It only generates wrappers for the
// parent field. The field is assumed to be public (exported) types. Types are
// assumed to all have valid wrap functions, so no nested wraps will actually be
// done.
//
// Wrapper functions for all types are assumed to follow this format:
//
//    func WrapTypeName(obj *externglib.Object) TypeName
//
func (tree *Tree) WrapClass(obj string) string {
	return tree.wrap(obj, true)
}

// WrapInterface creates a wrappper for an interface instead.
func (tree *Tree) WrapInterface(obj string) string {
	return tree.wrap(obj, false)
}

func (tree *Tree) wrap(obj string, class bool) string {
	p := pen.NewPiece()
	p.Write(tree.Resolved.ImplType(false)).Char('{')
	p.EmptyLine()

	iter := tree.Embeds
	if class {
		iter = tree.Requires // get the Parent from Requires
	}

	for _, typ := range iter {
		if typ.Resolved.Builtin != nil {
			// If these cases hit, then the type is an Objector (as deefined by
			// gextras.Objector), so obj satisfies it.
			switch {
			case typ.Resolved.IsExternGLib("InitiallyUnowned"):
				fallthrough
			case typ.Resolved.IsExternGLib("Object"):
				p.Linef("Objector: %s,", obj)
			default:
				tree.gen.Logln(logger.Debug, "unknown builtin wrap:", spew.Sdump(typ.Resolved))
			}
		} else {
			// Extern types are generated by us, so the wrapper guarantee is
			// provided.
			namespace := typ.Resolved.NeedsNamespace(tree.gen.Namespace())

			p.Linef(
				"%s: %s(%s),",
				typ.Resolved.PublicType(false), typ.Resolved.WrapName(namespace), obj,
			)
		}

		if class {
			break
		} else {
			continue
		}
	}

	p.Char('}')
	return p.String()
}

package girgen

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/diamondburned/gotk4/core/pen"
	"github.com/diamondburned/gotk4/gir"
)

// TypeTree is a structure for a type that is resolved to the lowest level of
// inheritance.
type TypeTree struct {
	*ResolvedType

	// Requires contains the direct dependencies of the current type. It may
	// contain interfaces that are also in other interfaces, which will not
	// build.
	Requires []TypeTree
	// Embeds contains the filtered dependencies, that is, ones that will not
	// collide when built.
	Embeds []TypeTree

	// Level sets the maximum recursion level to go. It only applies if set
	// to something more than -1.
	Level int

	res interface {
		TypeResolver
		LineLogger
	}
}

// TypeTree returns a new type tree for resolving.
func (ng *NamespaceGenerator) TypeTree() TypeTree {
	return TypeTree{res: ng, Level: -1}
}

func (tree *TypeTree) Reset() {
	// Zero out the fields to prevent dangling pointers.
	for i := range tree.Requires {
		tree.Requires[i] = TypeTree{}
	}

	tree.ResolvedType = nil
	tree.Embeds = tree.Embeds[:0]
	tree.Requires = tree.Requires[:0]
}

// Resolve resolves the given toplevel type into the TypeTree, overriding the
// Resolved and Requires fields. True is returned if the tree is successfully
// resolved.
func (tree *TypeTree) Resolve(toplevel string) bool {
	resolved := ResolveTypeName(tree.res, toplevel)
	if resolved == nil {
		tree.Reset()
		return false
	}

	return tree.ResolveFromType(resolved)
}

// ResolveFromType is like Resolve, but the caller directly supplies the
// resolved top-level type.
func (tree *TypeTree) ResolveFromType(toplevel *ResolvedType) bool {
	tree.Reset()
	tree.ResolvedType = toplevel

	if tree.Level == 0 {
		// No omit, since we added nothing.
		return true
	}

	// Edge cases for builtin types.
	if tree.ResolvedType.Builtin != nil {
		switch {
		case toplevel.IsExternGLib("InitiallyUnowned"):
			return tree.resolveParents(externGLibType("*Object", gir.Type{}, "GObject*"))
		case toplevel.IsExternGLib("Object"):
			return true
		}

		tree.omitAmbiguous(false)
		return true
	}

	if !tree.ResolvedType.Extern.Result.IsIntrospectable() {
		return false
	}

	switch {
	case tree.ResolvedType.Extern.Result.Class != nil:
		// All classes have a GObject parent, so an empty parent is invalid.
		parent := tree.ResolvedType.Extern.Result.Class.Parent
		if parent == "" {
			parent = "GObject.Object"
		}

		// Resolving the parent type is crucial to make the class working, so if
		// this fails, halt and bail.
		if !tree.resolveName(parent) {
			tryLogln(tree.res, LogUnknown,
				"can't resolve parent", tree.ResolvedType.Extern.Result.Class.Parent,
				"for class", tree.ResolvedType.Extern.Result.Class.Name,
			)
			return false
		}

		for _, impl := range tree.ResolvedType.Extern.Result.Class.Implements {
			if !tree.resolveName(impl.Name) {
				tryLogln(tree.res, LogUnknown,
					"can't resolve impl", impl.Name,
					"for class", tree.ResolvedType.Extern.Result.Class.Name,
				)
			}
		}

		tree.omitAmbiguous(true)

	case tree.ResolvedType.Extern.Result.Interface != nil:
		for _, prereq := range tree.ResolvedType.Extern.Result.Interface.Prerequisites {
			// Like class parents, interface prerequisites are important.
			if !tree.resolveName(prereq.Name) {
				tryLogln(tree.res, LogUnknown,
					"can't resolve prerequisite", prereq.Name,
					"for interface", tree.ResolvedType.Extern.Result.Interface.Name,
				)
				return false
			}
		}

		if len(tree.Requires) == 0 {
			// All interfaces are derived from GObjects, so we override the list
			// if it's empty.
			if !tree.resolveParents(externGLibType("*Object", gir.Type{}, "GObject*")) {
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
func (tree *TypeTree) omitAmbiguous(actually bool) {
	// Try and reuse the backing array, but regrow it ourselves if there's not
	// enough space.
	tree.Embeds = tree.Embeds[:0]
	if cap(tree.Embeds) < len(tree.Requires) {
		tree.Embeds = make([]TypeTree, 0, len(tree.Requires))
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

			if requiresHasGType(otherTree.Requires, prereqTree.ResolvedType) {
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
func requiresHasGType(requires []TypeTree, typ *ResolvedType) bool {
	return requiresHasGTyperec(requires, typ.PublicType(true))
}

func requiresHasGTyperec(requires []TypeTree, publType string) bool {
	for _, req := range requires {
		if req.ResolvedType.PublicType(true) == publType {
			return true
		}
		if requiresHasGTyperec(req.Requires, publType) {
			return true
		}
	}
	return false
}

func (tree *TypeTree) parentLevel() int {
	if tree.Level <= 0 {
		return tree.Level
	}
	return tree.Level - 1
}

// resolveName resolves and adds the resolved type into the TypeTree.
func (tree *TypeTree) resolveName(name string) bool {
	parent := TypeTree{
		res:   tree.res,
		Level: tree.parentLevel(),
	}

	if !parent.Resolve(name) {
		return false
	}

	tree.Requires = append(tree.Requires, parent)
	return true
}

// resolveParents manually adds the given parents and resolve them to be added
// into the TypeTree.
func (tree *TypeTree) resolveParents(parents ...*ResolvedType) bool {
	for _, parent := range parents {
		parentTree := TypeTree{
			res:   tree.res,
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
func (tree *TypeTree) PublicEmbeds() []string {
	names := make([]string, len(tree.Embeds))

	for i, req := range tree.Embeds {
		namespace := req.ResolvedType.NeedsNamespace(tree.res.Namespace())
		names[i] = req.ResolvedType.PublicType(namespace)
	}

	return names
}

// WalkPublInterfaces walks the tree for all embedded interfaces.
func (tree *TypeTree) WalkPublInterfaces(f func(*ResolvedType)) {
	for _, impl := range tree.Embeds {
		if impl.IsInterface() {
			f(impl.ResolvedType)
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
func (tree *TypeTree) WrapClass(obj string) string {
	return tree.wrap(obj, true)
}

// WrapInterface creates a wrappper for an interface instead.
func (tree *TypeTree) WrapInterface(obj string) string {
	return tree.wrap(obj, false)
}

func (tree *TypeTree) wrap(obj string, class bool) string {
	p := pen.NewPiece()
	p.Write(tree.ResolvedType.ImplType(false)).Char('{')
	p.EmptyLine()

	iter := tree.Embeds
	if class {
		iter = tree.Requires // get the Parent from Requires
	}

	for _, typ := range iter {
		if typ.ResolvedType.Builtin != nil {
			// If these cases hit, then the type is an Objector (as deefined by
			// gextras.Objector), so obj satisfies it.
			switch {
			case typ.ResolvedType.IsExternGLib("InitiallyUnowned"):
				fallthrough
			case typ.ResolvedType.IsExternGLib("Object"):
				p.Linef("Objector: %s,", obj)
			default:
				tryLogln(tree.res, LogUnknown, "builtin wrapping:", spew.Sdump(typ.ResolvedType))
			}
		} else {
			// Extern types are generated by us, so the wrapper guarantee is
			// provided.
			namespace := typ.ResolvedType.NeedsNamespace(tree.res.Namespace())

			p.Linef(
				"%s: %s(%s),",
				typ.ResolvedType.PublicType(false), typ.ResolvedType.WrapName(namespace), obj,
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

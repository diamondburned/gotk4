package types

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
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

	if toplevel.Extern != nil {
		// Ensure the origin namespace is set correctly.
		tree.gen = OverrideNamespace(tree.gen, toplevel.Extern.NamespaceFindResult)
	}

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

		tree.omitRedundant(true)

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

		tree.omitRedundant(false)
	}

	return true
}

// omitRedundant searches for each parent inherited field and omits the ones
// that are already inside others.
func (tree *Tree) omitRedundant(ignoreFirst bool) {
	cleaned := make([]Tree, 0, cap(tree.Requires))
	if ignoreFirst {
		cleaned = append(cleaned, tree.Requires[0])
	}

	for i, req := range tree.Requires {
		if i == 0 && ignoreFirst {
			continue
		}

		if !findInTree(tree.Requires, req.FullGType(), i) {
			cleaned = append(cleaned, req)
		}
	}

	tree.Requires = cleaned
}

func findInTree(reqs []Tree, girType string, ignore int) bool {
	for i, req := range reqs {
		if ignore == i {
			continue
		}

		if req.FullGType() == girType {
			return true
		}

		if findInTree(req.Requires, girType, -1) {
			return true
		}
	}

	return false
}

// HasAmbiguousSelector returns true if the GObject methods cannot be accessed
// normally.
func (tree *Tree) HasAmbiguousSelector() bool {
	if len(tree.Requires) == 1 {
		return false
	}

	// If the fields already have a GInitiallyUnowned field, then we're fine.
	for _, req := range tree.Requires {
		if req.IsExternGLib("Object") {
			return false
		}
	}

	depths := make(map[int]struct{}, 7) // arbitrarily 7 depth
	return gobjectDepth(tree, depths, 0)
}

func gobjectDepth(tree *Tree, depths map[int]struct{}, current int) bool {
	field := current + 1

	for _, req := range tree.Requires {
		switch {
		case req.IsExternGLib("Object"):
			_, ok := depths[field]
			if ok {
				return true
			}
			depths[field] = struct{}{}
			continue
		case req.IsExternGLib("InitiallyUnowned"):
			// Account for the current level and the children level as well,
			// since InitiallyUnowned already contains an Object.
			_, ok1 := depths[field]
			_, ok2 := depths[field+1]
			if ok1 || ok2 {
				return true
			}
			depths[field] = struct{}{}
			depths[field+1] = struct{}{}
			continue
		}

		if gobjectDepth(&req, depths, field) {
			return true
		}
	}

	// No GObject found.
	return false
}

// FirstGObjectSelector returns the selector path to the firts GObject field.
// The returning selector should have the *Object type. The selectors will all
// use the implementation type for the name.
//
// If the object does not contain the Object field somewhere, then "nil" is
// returned.
func (tree *Tree) FirstGObjectSelector(v string) string {
	sel, ok := firstGObjectSelector(tree.Requires, v)
	if !ok {
		return ""
	}
	return sel
}

func firstGObjectSelector(nodes []Tree, sel string) (string, bool) {
	for _, node := range nodes {
		fieldSel := sel + "." + node.ImplName()

		if node.IsExternGLib("Object") {
			return fieldSel, true
		}

		if sel, ok := firstGObjectSelector(node.Requires, fieldSel); ok {
			return sel, true
		}
	}

	return sel, false
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

// WithoutGObject returns the Requires list without the GObject item.
func (tree *Tree) WithoutGObject() []Tree {
	for _, req := range tree.Requires {
		if req.IsExternGLib("InitiallyUnowned") || req.IsExternGLib("Object") {
			goto hasGObject
		}
	}

	return tree.Requires

hasGObject:
	without := make([]Tree, 0, len(tree.Requires))
	for _, req := range tree.Requires {
		if req.IsExternGLib("InitiallyUnowned") || req.IsExternGLib("Object") {
			continue
		}
		without = append(without, req)
	}
	return without
}

// ImplTypes returns the sorted list of the toplevel type's children as Go
// implementation type names. The namespaces are appropriately prepended if
// needed.
func (tree *Tree) ImplTypes() []string {
	names := make([]string, len(tree.Requires))

	for i, req := range tree.Requires {
		namespace := req.Resolved.NeedsNamespace(tree.gen.Namespace())
		names[i] = req.Resolved.ImplType(namespace)
	}

	return names
}

// Walk walks the tree recursively on what the callback returns. The callback
// will be called on the root (receiver) tree; it should then return the list of
// parents of that tree. This process is then repeated for each tree returned.
func (tree *Tree) Walk(f func(t *Tree, root bool) (traversed []Tree)) {
	tree.walk(f, true)
}

func (tree *Tree) walk(f func(*Tree, bool) (traversed []Tree), isRoot bool) {
	// Get the list of traversed tree nodes.
	traversed := f(tree, isRoot)
	// Traverse each of the nodes' own nodes.
	for i := range traversed {
		traversed[i].walk(f, false)
	}
}

// ImplImporter is an interface that describes file.Header.
type ImplImporter interface {
	ImportImpl(*Resolved)
}

// Wrap generates the wrapper for the implementation struct.
func (tree *Tree) Wrap(obj string, h ImplImporter) string {
	return wrapRef(obj, tree.wrap(obj, h, tree.gen))
}

// WrapInNamespace wraps with the given current namespace.
func (tree *Tree) WrapInNamespace(obj string, h ImplImporter, n *gir.NamespaceFindResult) string {
	return wrapRef(obj, tree.wrap(obj, h, OverrideNamespace(tree.gen, n)))
}

func wrapRef(obj, wrap string) string {
	if wrap == obj {
		// GObject is already a pointer.
		return obj
	}
	return "&" + wrap
}

func (tree *Tree) wrap(obj string, h ImplImporter, gen FileGenerator) string {
	if tree.Resolved.Builtin != nil {
		switch {
		case tree.Resolved.IsExternGLib("Object"):
			return "obj"
		case tree.Resolved.IsExternGLib("InitiallyUnowned"):
			h.ImportImpl(tree.Resolved)
			return fmt.Sprintf("externglib.InitiallyUnowned{\nObject: %s,\n}", obj)
		default:
			tree.gen.Logln(logger.Debug, "unknown builtin wrap:", spew.Sdump(tree.Resolved))
			return fmt.Sprintf("nil /* unknown type %s */", tree.Resolved.ImplType(true))
		}
	}

	needsNamespace := tree.Resolved.NeedsNamespace(gen.Namespace())
	if needsNamespace {
		h.ImportImpl(tree.Resolved)
	}

	typ := tree.Resolved.ImplType(needsNamespace)

	p := pen.NewPiece()
	p.Write(strings.TrimPrefix(typ, "*")).Char('{')
	p.EmptyLine()

	for _, typ := range tree.Requires {
		// Recursively resolve the wrapper.
		typ := typ
		p.Linef("%s: %s,", typ.Resolved.Name(), typ.wrap(obj, h, gen))
	}

	p.Char('}')
	return p.String()
}

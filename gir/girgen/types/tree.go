package types

import (
	"fmt"
	"sort"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
)

// Tree is a structure for a type that is resolved to the lowest level of
// inheritance.
type Tree struct {
	*Resolved
	gen FileGenerator
	src FileGenerator

	// Requires contains the direct dependencies of the current type. It may
	// contain interfaces that are also in other interfaces, which will not
	// build.
	Requires []Tree
}

// NewTree creates a new empty type tree for resolving.
func NewTree(gen FileGenerator) Tree {
	return Tree{gen: gen, src: gen}
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

var gobjectTreeType = externGLibType("*Object", gir.Type{
	Name:  "GObject.Object",
	CType: "GObject*",
}, "GObject*")

// ResolveFromType is like Resolve, but the caller directly supplies the
// resolved top-level type.
func (tree *Tree) ResolveFromType(toplevel *Resolved) bool {
	tree.Reset()
	tree.Resolved = toplevel

	// Edge cases for builtin types.
	if tree.Resolved.Builtin != nil {
		switch {
		case toplevel.IsExternGLib("InitiallyUnowned"):
			return tree.resolveParents(gobjectTreeType)
		case toplevel.IsExternGLib("Object"):
			return true
		}

		return false
	}

	if !tree.Resolved.Extern.IsIntrospectable() {
		return false
	}

	// Ensure that the namespace is correct.
	tree.gen = OverrideNamespace(tree.src, toplevel.Extern.NamespaceFindResult)

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
			tree.gen.Logln(logger.Debug, "can't resolve parent", parent, "for class", v.Name,
				"namespace")
			return false
		}

		for _, impl := range v.Implements {
			if !tree.resolveName(impl.Name) {
				tree.gen.Logln(logger.Debug, "can't resolve impl", impl.Name, "for class", v.Name)
			}
		}

		tree.omitRedundant(true)
		tree.sortRequires(true)

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
			if !tree.resolveFromType(gobjectTreeType) {
				tree.gen.Logln(logger.Debug,
					"can't resolve fallback prerequisite *GObject for interface", v.Name)
				return false
			}
		}

		tree.omitRedundant(false)
		tree.sortRequires(false)
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

	// Resolve additional ambiguous types after omitting.
	for _, ambiguity := range tree.AmbiguousSelectorTypes() {
		tree.resolveFromType(ambiguity)
	}
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

// AmbiguousSelectorTypes scans the whole tree and finds all types whose
// selectors are ambiguous if directly referenced from the top-level type.
func (tree *Tree) AmbiguousSelectorTypes() []*Resolved {
	if len(tree.Requires) <= 1 {
		return nil
	}

	depths := make(depthMap, len(tree.Requires)*2) // arbitrary growing
	searchFullDepth(tree, depths, 0)
	return depths.conflictingGTypes(tree)
}

// searchFullDepth searches the whole tree for all same-level depths.
func searchFullDepth(tree *Tree, depths depthMap, current int) {
	field := current + 1

	for _, req := range tree.Requires {
		switch {
		case req.IsExternGLib("Object"):
			depths.add(req.Resolved, field)
			continue
		case req.IsExternGLib("InitiallyUnowned"):
			// Since InitiallyUnowned already contains an Object, skip counting
			// it and count the object directly. This is the only built-in that
			// isn't Object that we'd care about having an embedded types.
			depths.add(req.Resolved, field)
			depths.add(gobjectTreeType, field+1)
			continue
		default:
			if req.Underlying().Extern == nil {
				// No other builtins have methods. Skip.
				continue
			}
		}

		// Register the current entity if it has a method. This isn't
		// necessarily a class.
		if req.IsClass() || req.IsInterface() || req.IsRecord() {
			depths.add(req.Underlying(), field)
		}

		searchFullDepth(&req, depths, field)
	}
}

// depthMap keeps track of the depths of all embedded classes/structs to resolve
// for conflicts.
type depthMap map[string]depthNode

type depthNode struct {
	*Resolved // do not assume equality
	depths    map[int]int
}

func (depths depthMap) conflictingGTypes(toplevel *Tree) []*Resolved {
	resolved := make([]*Resolved, 0, len(depths))

	requires := make(map[string]struct{}, len(toplevel.Requires))
	for _, req := range toplevel.Requires {
		requires[req.GType] = struct{}{}
	}

	for _, node := range depths {
		var repeating bool

		for d, repeats := range node.depths {
			// If the current depth is deeper than the top-level (that is, if
			// it's inside a field of the current struct), and that this same
			// depth on the same type occurs more than once, then count it as
			// repeating.
			if d > 1 && repeats > 1 {
				repeating = true
				break
			}
		}

		if !repeating {
			continue
		}

		// Exclude structs that are already top-level.
		_, isToplevel := requires[node.GType]
		if isToplevel {
			continue
		}

		resolved = append(resolved, node.Resolved)
	}

	return resolved
}

// add adds the given class and depth into the depths map. False is returned if
// the depth is already in.
func (depths depthMap) add(r *Resolved, depth int) {
	gtype := r.GType

	i, ok := depths[gtype].depths[depth]
	if ok {
		depths[gtype].depths[depth] = i + 1
		return
	}

	depthMap, ok := depths[gtype]
	if !ok {
		depthMap = depthNode{
			Resolved: r,
			depths:   make(map[int]int),
		}
		depths[gtype] = depthMap
	}

	depthMap.depths[depth] = 1
	return
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

func (tree *Tree) sortRequires(skipFirst bool) {
	req := tree.Requires
	if skipFirst {
		req = req[1:]
	}

	sort.Slice(req, func(i, j int) bool {
		return req[i].ImplType(true) < req[j].ImplType(true)
	})
}

// resolveName resolves and adds the resolved type into the Tree.
func (tree *Tree) resolveName(name string) bool {
	parent := Tree{
		gen: tree.gen,
		src: tree.src,
	}

	if !parent.Resolve(name) {
		return false
	}

	tree.Requires = append(tree.Requires, parent)
	return true
}

// resolveFromType is a resolveName variant that accepts a resolved type.
func (tree *Tree) resolveFromType(t *Resolved) bool {
	parent := Tree{
		gen: tree.gen,
		src: tree.src,
	}

	if !parent.ResolveFromType(t) {
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
			gen: tree.gen,
			src: tree.src,
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

// ImplInterfaces returns the sorted list of all Go interfaces that this type
// implements. The namespaces are appropriately prepended if needed.
//
// At the moment, the method will only care about abstract classes.
func (tree *Tree) ImplInterfaces() []*Resolved {
	gtypes := make(map[string]struct{}, len(tree.Requires))
	gtypes[tree.GType] = struct{}{}
	return tree.implInterfaces(make([]*Resolved, 0, len(tree.Requires)), gtypes)
}

func (tree *Tree) implInterfaces(in []*Resolved, gtypes map[string]struct{}) []*Resolved {
	for _, req := range tree.Requires {
		if req.IsExternGLib("Object") || (req.IsClass() && req.IsAbstract()) {
			_, dupe := gtypes[req.GType]
			if dupe {
				continue
			}
			in = append(in, req.Underlying())
			gtypes[req.GType] = struct{}{}
		}
	}

	for _, req := range tree.Requires {
		if req.IsClass() {
			// Skip to avoid infinitely recursing.
			_, dupe := gtypes[req.GType]
			if dupe {
				continue
			}
			in = req.implInterfaces(in, gtypes)
			gtypes[req.GType] = struct{}{}
		}
	}

	return in
}

// Walk walks the tree recursively on what the callback returns. The callback
// will be called on the root (receiver) tree; it should then return the list of
// parents of that tree. This process is then repeated for each tree returned.
//
// Example:
//
//    tree.Walk(func(t *Tree, root bool) []Tree {
//        log.Println("currently at", t.Resolved.PublName())
//        return t.Requires
//    })
//
func (tree *Tree) Walk(f func(t *Tree, root bool) (traversed []Tree)) {
	tree.walk(f, true)
}

func (tree *Tree) walk(f func(*Tree, bool) []Tree, isRoot bool) {
	// Get the list of traversed tree nodes.
	traversed := f(tree, isRoot)
	// Traverse each of the nodes' own nodes.
	for i := range traversed {
		traversed[i].walk(f, false)
	}
}

// Wrap generates the wrapper for the implementation struct.
func (tree *Tree) Wrap(obj string, h *file.Header) string {
	return wrapRef(obj, tree.wrap(obj, h, tree.gen))
}

// WrapInNamespace wraps with the given current namespace.
func (tree *Tree) WrapInNamespace(obj string, h *file.Header, n *gir.NamespaceFindResult) string {
	return wrapRef(obj, tree.wrap(obj, h, OverrideNamespace(tree.gen, n)))
}

func wrapRef(obj, wrap string) string {
	if wrap == obj {
		// GObject is already a pointer.
		return obj
	}
	return "&" + wrap
}

func (tree *Tree) wrap(obj string, h *file.Header, gen FileGenerator) string {
	if tree.Resolved.Builtin != nil {
		switch {
		case tree.Resolved.IsExternGLib("Object"):
			return "obj"
		case tree.Resolved.IsExternGLib("InitiallyUnowned"):
			tree.Resolved.ImportImpl(h)
			return fmt.Sprintf("coreglib.InitiallyUnowned{\nObject: %s,\n}", obj)
		default:
			tree.gen.Logln(logger.Debug, "unknown builtin wrap:", spew.Sdump(tree.Resolved))
			return fmt.Sprintf("nil /* unknown type %s */", tree.Resolved.ImplType(true))
		}
	}

	needsNamespace := tree.Resolved.NeedsNamespace(gen.Namespace())
	if needsNamespace {
		tree.Resolved.ImportImpl(h)
	}

	typName := tree.Resolved.ImplType(needsNamespace)

	p := pen.NewPiece()
	p.Write(strings.TrimPrefix(typName, "*")).Char('{')
	p.EmptyLine()

	for _, typ := range tree.Requires {
		// Recursively resolve the wrapper.
		typ := typ
		p.Linef("%s: %s,", typ.Resolved.Name(), typ.wrap(obj, h, gen))
	}

	p.Char('}')
	return p.String()
}

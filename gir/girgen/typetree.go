package girgen

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

// TypeTree is a structure for a type that is resolved to the lowest level of
// inheritance.
type TypeTree struct {
	Resolved *ResolvedType
	Requires []TypeTree

	ng *NamespaceGenerator
}

// TypeTree returns a new type tree for resolving.
func (ng *NamespaceGenerator) TypeTree() *TypeTree {
	return &TypeTree{ng: ng}
}

func (tree *TypeTree) reset() {
	// Zero out the fields to prevent dangling pointers.
	for i := range tree.Requires {
		tree.Requires[i] = TypeTree{}
	}

	tree.Resolved = nil
	tree.Requires = tree.Requires[:0]
}

// Resolve resolves the given toplevel type into the TypeTree, overriding the
// Resolved and Requires fields. True is returned if the tree is successfully
// resolved.
func (tree *TypeTree) Resolve(toplevel string) bool {
	resolved := tree.ng.ResolveTypeName(toplevel)
	if resolved == nil {
		tree.reset()
		return false
	}

	return tree.ResolveFromType(resolved)
}

// ResolveFromType is like Resolve, but the caller directly supplies the
// resolved top-level type.
func (tree *TypeTree) ResolveFromType(toplevel *ResolvedType) bool {
	tree.reset()
	tree.Resolved = toplevel

	// Edge cases for builtin types.
	if tree.Resolved.Builtin != nil {
		// This should be changed to a slice, just in case there's a builtin
		// type that has more than 1 inheritance, but there's none so far.
		var parent *ResolvedType

		switch {
		case toplevel.IsExternGLib("InitiallyUnowned"):
			parent = externGLibType("*Object", gir.Type{}, "*GObject")
		}

		if parent != nil {
			parentTree := TypeTree{ng: tree.ng}
			if !parentTree.ResolveFromType(parent) {
				// This shouldn't happen, unless the parent type made above is
				// invalid.
				return false
			}

			tree.Requires = append(tree.Requires, parentTree)
		}

		return true
	}

	switch {
	case tree.Resolved.Extern.Result.Class != nil:
		parent := TypeTree{ng: tree.ng}
		if !parent.Resolve(tree.Resolved.Extern.Result.Class.Parent) {
			return false
		}

		tree.Requires = append(tree.Requires, parent)

	case tree.Resolved.Extern.Result.Interface != nil:
		for _, prereq := range tree.Resolved.Extern.Result.Interface.Prerequisites {
			parent := TypeTree{ng: tree.ng}
			if !parent.Resolve(prereq.Name) {
				return false
			}

			tree.Requires = append(tree.Requires, parent)
		}
	}

	return true
}

// PublicChildren returns the list of the toplevel type's children as Go
// exported type names. The namespaces are appropriately prepended if needed.
func (tree *TypeTree) PublicChildren() []string {
	names := make([]string, len(tree.Requires))

	for i, req := range tree.Requires {
		namespace := req.Resolved.NeedsNamespace(tree.ng.current)
		names[i] = req.Resolved.PublicType(namespace)
	}

	return names
}

// Wrap creates a wrapper that uses public fields to create code that wraps the
// type tree to the top-level type. The fields are assumed to be public
// (exported) types. Types are assumed to all have valid wrap functions, so no
// nested wraps will actually be done.
//
// Wrapper functions for all types are assumed to follow this format:
//
//    func WrapTypeName(obj *externglib.Object) TypeName
//
func (tree *TypeTree) Wrap(obj string) string {
	p := pen.NewPiece()
	p.Write(tree.Resolved.PublicType(false)).Char('{')
	p.EmptyLine()

	for _, typ := range tree.Requires {
		namespace := typ.Resolved.NeedsNamespace(tree.ng.current)

		p.Linef(
			"%s: %s(%s),",
			typ.Resolved.PublicType(namespace),
			typ.Resolved.WrapName(namespace),
			obj,
		)
	}

	p.Char('}')
	return p.String()
}

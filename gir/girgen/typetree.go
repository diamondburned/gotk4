package girgen

import (
	"github.com/diamondburned/gotk4/core/pen"
	"github.com/diamondburned/gotk4/gir"
)

// TypeTree is a structure for a type that is resolved to the lowest level of
// inheritance.
type TypeTree struct {
	Resolved *ResolvedType
	Requires []TypeTree

	// Level sets the maximum recursion level to go. It only applies if set
	// to something more than -1.
	Level int

	res interface {
		TypeResolver
		LineLogger
	}
}

func NewTypeTree(res interface {
	TypeResolver
	LineLogger
}) TypeTree {
	return TypeTree{res: res, Level: -1}
}

func (tree *TypeTree) reset() {
	// Zero out the fields to prevent dangling pointers.
	for i := range tree.Requires {
		tree.Requires[i] = TypeTree{}
	}

	tree.Resolved = nil
	tree.Requires = tree.Requires[:0]
}

// IsInterface returns true if the current type in the tree is an interface.
func (tree *TypeTree) IsInterface() bool {
	return tree.Resolved.Extern != nil && tree.Resolved.Extern.Result.Interface != nil
}

// Resolve resolves the given toplevel type into the TypeTree, overriding the
// Resolved and Requires fields. True is returned if the tree is successfully
// resolved.
func (tree *TypeTree) Resolve(toplevel string) bool {
	resolved := ResolveTypeName(tree.res, toplevel)
	if resolved == nil {
		tree.reset()
		return false
	}

	return tree.ResolveFromType(resolved)
}

// treeExternGObject is for use in TypeTree ONLY.
var treeExternGObject = externGLibType("*Object", gir.Type{}, "GObject*")

// ResolveFromType is like Resolve, but the caller directly supplies the
// resolved top-level type.
func (tree *TypeTree) ResolveFromType(toplevel *ResolvedType) bool {
	tree.reset()
	tree.Resolved = toplevel

	if tree.Level == 0 {
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

	if !tree.Resolved.Extern.Result.IsIntrospectable() {
		return false
	}

	switch {
	case tree.Resolved.Extern.Result.Class != nil:
		// Resolving the parent type is crucial to make the class working, so if
		// this fails, halt and bail.
		if !tree.resolveType(tree.Resolved.Extern.Result.Class.Parent) {
			tryLogln(tree.res, LogUnknown,
				"can't resolve parent", tree.Resolved.Extern.Result.Class.Parent,
				"for class", tree.Resolved.Extern.Result.Class.Name,
			)
			return false
		}

		for _, impl := range tree.Resolved.Extern.Result.Class.Implements {
			if !tree.resolveType(impl.Name) {
				tryLogln(tree.res, LogUnknown,
					"can't resolve impl", impl.Name,
					"for class", tree.Resolved.Extern.Result.Class.Name,
				)
			}
		}

	case tree.Resolved.Extern.Result.Interface != nil:
		// All interfaces are derived from GObjects, so we override the list if
		// it's empty.
		if len(tree.Resolved.Extern.Result.Interface.Prerequisites) == 0 {
			return tree.resolveParents(treeExternGObject)
		}

		for _, prereq := range tree.Resolved.Extern.Result.Interface.Prerequisites {
			// Like class parents, interface prerequisites are important.
			if !tree.resolveType(prereq.Name) {
				tryLogln(tree.res, LogUnknown,
					"can't resolve prerequisite", prereq.Name,
					"for interface", tree.Resolved.Extern.Result.Interface.Name,
				)
				return false
			}
		}
	}

	return true
}

func (tree *TypeTree) parentLevel() int {
	if tree.Level <= 0 {
		return tree.Level
	}
	return tree.Level - 1
}

// resolveType resolves and adds the resolved type into the TypeTree.
func (tree *TypeTree) resolveType(name string) bool {
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

// ImportChildren imports the type tree's public children into the given file
// generator.
func (tree *TypeTree) ImportChildren(ng *NamespaceGenerator) {
	for _, req := range tree.Requires {
		ng.importResolved(req.Resolved)
	}
}

// Children returns the list of the toplevel type's children as Go
// exported type names. The namespaces are appropriately prepended if needed.
func (tree *TypeTree) Children() []string {
	names := make([]string, len(tree.Requires))

	for i, req := range tree.Requires {
		namespace := req.Resolved.NeedsNamespace(tree.res.Namespace())
		names[i] = req.Resolved.GoType(namespace)
	}

	return names
}

type WrapOutput struct {
	Wrapper string
	Imports map[string]string
}

func (out *WrapOutput) importResolved(res TypeResolver, typ *ResolvedType) {
	if out.Imports == nil {
		out.Imports = map[string]string{}
	}

	if typ.NeedsNamespace(res.Namespace()) {
		out.Imports[typ.Import.Path] = typ.Import.Package
	}
}

// ApplySideEffects applies the side effects from the wrap output to the given
// side effects ptr.
func (out *WrapOutput) ApplySideEffects(dst *SideEffects) {
	for path, alias := range out.Imports {
		dst.addImportAlias(path, alias)
	}
}

// Wrap creates a wrapper that uses public fields to create code that wraps the
// type tree to the top-level type.
func (tree *TypeTree) Wrap(objOrPtr string) WrapOutput {
	var out WrapOutput
	out.Wrapper = tree.wrap(objOrPtr, &out)
	return out
}

func (tree *TypeTree) wrap(objOrPtr string, out *WrapOutput) string {
	needsNamespace := tree.Resolved.NeedsNamespace(tree.res.Namespace())
	if needsNamespace {
		out.importResolved(tree.res, tree.Resolved)
	}

	p := pen.NewPiece()
	p.Write(tree.Resolved.GoType(needsNamespace)).Char('{')
	p.EmptyLine()

	for _, typ := range tree.Requires {
		switch {
		case typ.Resolved.IsExternGLib("Object"):
			out.importResolved(tree.res, treeExternGObject)

			p.Linef(
				"Object: &%s.Object{%[1]s.ToGObject(%s)},",
				typ.Resolved.Import.Package, objOrPtr,
			)

		case typ.Resolved.IsRecord():
			needsNamespace := typ.Resolved.NeedsNamespace(tree.res.Namespace())
			if needsNamespace {
				out.importResolved(tree.res, typ.Resolved)
			}

			p.Linef(
				"%s: (*%s)(unsafe.Pointer(%s)),",
				typ.Resolved.GoType(false), typ.Resolved.GoType(needsNamespace), objOrPtr,
			)

		default:
			// Recursively generate the wrapper for each subtype.
			p.Linef("%s: %s,", typ.Resolved.GoType(false), typ.wrap(objOrPtr, out))
		}
	}

	p.Char('}')
	return p.String()
}

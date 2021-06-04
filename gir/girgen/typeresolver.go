package girgen

import (
	"path"
	"strings"
	"unicode"

	"github.com/diamondburned/gotk4/gir"
)

// ResolvedType is a resolved type from a given gir.Type.
type ResolvedType struct {
	// either or
	Extern  *ExternType // optional
	Builtin *string     // optional

	Import  string // Full import path.
	Package string // Package name, also import alias.

	GType string
	CType string
	Ptr   uint8
}

// ExternType is an externally resolved type.
type ExternType struct {
	Result *gir.TypeFindResult
}

// builtinType is a convenient function to make a new resolvedType.
func builtinType(imp, typ string, girType gir.Type) *ResolvedType {
	var pkg string

	if imp != "" {
		// Create the actual type if there's an import path.
		pkg = path.Base(imp)
		typ = pkg + "." + typ
	}

	return &ResolvedType{
		Builtin: &typ,
		Import:  imp,
		Package: pkg,
		GType:   girType.Name,
		CType:   ctypeFallback(girType.CType, girType.Name),
		Ptr:     countPtrs(girType, nil),
	}
}

// externGLibType returns an external GLib type from gotk3.
func externGLibType(goType string, typ gir.Type, ctyp string) *ResolvedType {
	if typ.CType != "" {
		ctyp = typ.CType
	}

	ptrs := strings.Count(goType, "*")
	goType = strings.Repeat("*", ptrs) + "externglib." + strings.TrimPrefix(goType, "*")

	return &ResolvedType{
		Builtin: &goType,
		Import:  "github.com/gotk3/gotk3/glib",
		Package: "externglib",
		GType:   typ.Name,
		CType:   ctyp,
		Ptr:     uint8(ptrs),
	}
}

// typeFromResult creates a resolved type from the given type result.
func typeFromResult(gen *Generator, typ gir.Type, result *gir.TypeFindResult) *ResolvedType {
	if typ.CType == "" {
		gen.Logln(LogUnusuality, "type", typ.Name, "missing CType")
		return nil
	}

	return &ResolvedType{
		Extern: &ExternType{
			Result: result,
		},
		Import:  gen.ModPath(result.Namespace),
		Package: gir.GoNamespace(result.Namespace),
		GType:   typ.Name,
		CType:   typ.CType,
		Ptr:     countPtrs(typ, result),
	}
}

// TypeFromResult creates a new ResolvedType from the given type find result.
// This function is mostly useful for generating from an existing GIR value.
func TypeFromResult(ng *NamespaceGenerator, res gir.TypeFindResult) *ResolvedType {
	res.NamespaceFindResult = ng.Namespace()
	name, ctype := res.Info()
	return typeFromResult(ng.gen, gir.Type{Name: name, CType: ctype}, &res)
}

// IsExternGLib checks that the ResolvedType is exactly the gotk3/glib type with
// the given name. Pointers are not compared.
func (typ *ResolvedType) IsExternGLib(glibType string) bool {
	if typ.Builtin == nil || typ.Import != "github.com/gotk3/gotk3/glib" {
		return false
	}

	thisType := *typ.Builtin
	thisType = strings.ReplaceAll(thisType, "*", "")
	thisType = strings.TrimPrefix(thisType, typ.Package)
	thisType = strings.TrimPrefix(thisType, ".")

	return thisType == glibType
}

// IsPrimitive returns true if the resolved type is a builtin type that can be
// directly casted to an equivalent C type.
func (typ *ResolvedType) IsPrimitive() bool {
	return typ.Builtin != nil &&
		typ.Package == "" &&
		typ.Import == "" &&
		*typ.Builtin != "string"
}

// GoImplType is a convenient function around ResolvedType.ImplType.
func GoImplType(resolver TypeResolver, resolved *ResolvedType) string {
	return resolved.ImplType(resolved.NeedsNamespace(resolver.Namespace()))
}

// GoPublicType is a convenient function around ResolvedType.ImplType.
func GoPublicType(resolver TypeResolver, resolved *ResolvedType) string {
	return resolved.PublicType(resolved.NeedsNamespace(resolver.Namespace()))
}

// NeedsNamespace returns true if the returned Go type needs a namespace to be
// referenced properly.
func (typ *ResolvedType) NeedsNamespace(current *gir.NamespaceFindResult) bool {
	if typ.Extern == nil {
		return false
	}

	return !typ.Extern.Result.Eq(current)
}

// ImplType returns the implementation type. This is only different to
// PublicType as far as classes go: the returned type is the unexported
// implementation type.
func (typ *ResolvedType) ImplType(needsNamespace bool) string {
	if typ.Builtin != nil {
		return *typ.Builtin
	}

	name, _ := typ.Extern.Result.Info()
	name = PascalToGo(name)

	if typ.Extern.Result.Class != nil {
		name = UnexportPascal(name)
	}

	ptr := strings.Repeat("*", int(typ.Ptr))

	if !needsNamespace {
		return ptr + name
	}

	return ptr + typ.Package + "." + name
}

// PublicType returns the public type. If the resolved type is a class, then the
// interface type is returned.
func (typ *ResolvedType) PublicType(needsNamespace bool) string {
	switch {
	case
		typ.IsExternGLib("InitiallyUnowned"),
		typ.IsExternGLib("Object"):

		// TODO: there should be a better way to do this; one that adds imports.
		return "gextras.Objector"
	}

	if typ.Builtin != nil {
		return *typ.Builtin
	}

	name, _ := typ.Extern.Result.Info()
	name = PascalToGo(name)

	ptr := typ.Ptr
	// Since classes are implemented as interfaces, we have to dereference a
	// pointer if we have any.
	if typ.Extern.Result.Class != nil && ptr > 0 {
		ptr--
	}

	ptrStr := strings.Repeat("*", int(ptr))

	if !needsNamespace {
		return ptrStr + name
	}

	return ptrStr + typ.Package + "." + name
}

// WrapName returns the name of the wrapper function. It only works for external
// types; calling this on a built-in ResolvedType will return an empty string.
func (typ *ResolvedType) WrapName(needsNamespace bool) string {
	if typ.Extern == nil {
		return ""
	}

	name, _ := typ.Extern.Result.Info()
	name = PascalToGo(name)

	wrapName := "Wrap" + name
	if needsNamespace {
		wrapName = typ.Package + "." + wrapName
	}

	return wrapName
}

// CGoType returns the CGo type.
func (typ *ResolvedType) CGoType() string {
	return movePtr(typ.CType, "C."+cleanCType(typ.CType))
}

// TypeResolver describes a generator that can resolve a GIR type.
type TypeResolver interface {
	// ResolveType resolves the given type from the GIR type field. It returns
	// nil if the type is not known. It does not recursively traverse the type.
	ResolveType(gir.Type) *ResolvedType
	// FindType finds the given GIR type name.
	FindType(gir string) *gir.TypeFindResult
	// Namespace returns the generator's namespace that includes the repository
	// it's in.
	Namespace() *gir.NamespaceFindResult
}

var (
	_ TypeResolver = (*FileGenerator)(nil)
	_ TypeResolver = (*NamespaceGenerator)(nil)
)

// FindType finds the given GIR type.
func (ng *NamespaceGenerator) FindType(girType string) *gir.TypeFindResult {
	return ng.gen.Repos.FindType(ng.current, girType)
}

// ensureNamespace ensures that exported, non-primitive types have the namespace
// prepended. This is useful for matching hard-coded types.
func ensureNamespace(nsp *gir.NamespaceFindResult, girType string) string {
	// Special cases, because GIR is very unusual.
	switch girType {
	case "GType":
		return girType
	}

	if strings.Contains(girType, ".") {
		return girType
	}

	caps := strings.IndexFunc(girType, unicode.IsUpper)
	// First letter isn't capitalized; this isn't exported.
	if caps != 0 {
		return girType
	}

	return nsp.Namespace.Name + "." + girType
}

// unsupportedCTypes is the list of unsupported C types, either because it is
// not yet supported or will never be supported due to redundancy or else.
var unsupportedCTypes = []string{
	"tm*", // requires time.h
	"va_list",
}

// ResolveType resolves the given type from the GIR type field. It returns nil
// if the type is not known. It does not recursively traverse the type.
func (ng *NamespaceGenerator) ResolveType(typ gir.Type) *ResolvedType {
	if typ.Name == "" {
		// empty gir type
		return nil
	}

	for _, unsupported := range unsupportedCTypes {
		if unsupported == typ.CType {
			return nil
		}
	}

	if prim, ok := girToBuiltin[typ.Name]; ok {
		return builtinType("", prim, typ)
	}

	// Resolve the unknown namespace that is GLib and primitive types.
	switch ensureNamespace(ng.Namespace(), typ.Name) {
	// TODO: ignore field
	// TODO: aaaaaaaaaaaaaaaaaaaaaaa
	case "gpointer":
		return builtinType("", "interface{}", typ)
	case "GObject.Error":
		return builtinType("", "error", typ)
	case "GObject.Type", "GType":
		return externGLibType("Type", typ, "GType")
	case "GObject.Value": // inconsistency???
		return externGLibType("*Value", typ, "GValue*")
	case "GObject.Object":
		return externGLibType("*Object", typ, "GObject*")
	case "GObject.InitiallyUnowned":
		return externGLibType("InitiallyUnowned", typ, "GInitiallyUnowned*")
	}

	// CType is required here so we can properly account for pointers.
	if typ.CType == "" {
		ng.Logln(LogUnusuality, "type name", typ.Name, "missing CType")
		return nil
	}

	// Pretend that ignored types don't exist.
	if ng.mustIgnore(typ.Name, typ.CType) {
		return nil
	}

	result := ng.FindType(typ.Name)
	if result == nil {
		warnUnknownType(ng, typ.Name)
		return nil
	}

	// Check for edge cases.
	switch {
	case result.Record != nil:
		if !canRecord(nil, *result.Record) {
			return nil
		}
	}

	return typeFromResult(ng.gen, typ, result)
}

// FindType finds the given GIR type.
func (fg *FileGenerator) FindType(girType string) *gir.TypeFindResult {
	return fg.parent.gen.Repos.FindType(fg.parent.current, girType)
}

// ResolveType resolves the GIR type and adds it to the import header.
func (fg *FileGenerator) ResolveType(typ gir.Type) *ResolvedType {
	resolved := fg.parent.ResolveType(typ)
	fg.addResolvedImport(resolved)
	return resolved
}

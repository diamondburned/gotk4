package girgen

import (
	"path"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

// ResolvedType is a resolved type from a given gir.Type.
type ResolvedType struct {
	// either or
	Extern  *ExternType // optional
	Builtin *string     // optional

	Import ResolvedTypeImport
	CType  string
	GType  string
	Ptr    uint8 // used ONLY for the Go type.
}

// ExternType is an externally resolved type.
type ExternType struct {
	Result *gir.TypeFindResult
}

// ResolvedTypeImport is a single import for the resolved type.
type ResolvedTypeImport struct {
	Path    string // full path
	Package string // package name, import alias
}

// These types contain an internal pointer in Go, so the pointer count
// should be decreased.
var goContainerTypes = []string{
	"error",
	"string",
}

// builtinType is a convenient function to make a new resolvedType.
func builtinType(imp, typ string, girType gir.Type) *ResolvedType {
	var pkg string

	if imp != "" {
		// Create the actual type if there's an import path.
		pkg = path.Base(imp)
		typ = pkg + "." + typ
	}

	ptr := countPtrs(girType, nil)

	if ptr > 0 {
		for _, iface := range goContainerTypes {
			if iface == typ {
				ptr--
				break
			}
		}
	}

	return &ResolvedType{
		Builtin: &typ,
		Import: ResolvedTypeImport{
			Path:    imp,
			Package: pkg,
		},
		GType: girType.Name,
		CType: ctypeFallback(girType.CType, girType.Name),
		Ptr:   ptr,
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
		Import: ResolvedTypeImport{
			Path:    "github.com/gotk3/gotk3/glib",
			Package: "externglib",
		},
		GType: typ.Name,
		CType: ctyp,
		Ptr:   uint8(ptrs),
	}
}

// typeFromResult creates a resolved type from the given type result.
func typeFromResult(gen *Generator, typ gir.Type, result *gir.TypeFindResult) *ResolvedType {
	if typ.CType == "" {
		gen.Logln(LogUnusuality, "type", typ.Name, "missing CType")
		return nil
	}

	ptr := countPtrs(typ, result)

	// Always use internal types (like GVariant) by reference and not value,
	// since Go will refuse to allocate them.
	if result.Record != nil && recordIsOpaque(*result.Record) && ptr == 0 {
		ptr++
	}

	return &ResolvedType{
		Extern: &ExternType{Result: result},
		Import: ResolvedTypeImport{
			Path:    gen.ModPath(result.Namespace),
			Package: gir.GoNamespace(result.Namespace),
		},
		GType: typ.Name,
		CType: typ.CType,
		Ptr:   ptr,
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
	// Use ImplImport for comparison, so we're not comparing gextras types.
	if typ.Builtin == nil || typ.Import.Path != "github.com/gotk3/gotk3/glib" {
		return false
	}

	thisType := *typ.Builtin
	thisType = strings.ReplaceAll(thisType, "*", "")
	thisType = strings.TrimPrefix(thisType, typ.Import.Package)
	thisType = strings.TrimPrefix(thisType, ".")

	return thisType == glibType
}

// IsCallback returns true if the current ResolvedType is a callback.
func (typ *ResolvedType) IsCallback() bool {
	return typ.Extern != nil && typ.Extern.Result.Callback != nil
}

// IsRecord returns true if the current ResolvedType is a record.
func (typ *ResolvedType) IsRecord() bool {
	return typ.Extern != nil && typ.Extern.Result.Record != nil
}

// IsInterface returns true if the current ResolvedType is an interface.
func (typ *ResolvedType) IsInterface() bool {
	return typ.Extern != nil && typ.Extern.Result.Interface != nil
}

// IsClass returns true if the current ResolvedType is a class.
func (typ *ResolvedType) IsClass() bool {
	return typ.Extern != nil && typ.Extern.Result.Class != nil
}

// IsPrimitive returns true if the resolved type is a builtin type that can be
// directly casted to an equivalent C type OR a record..
func (typ *ResolvedType) IsPrimitive() bool {
	if typ.Builtin == nil {
		return false
	}

	if typ.HasImport() {
		return false
	}

	for _, ctype := range goContainerTypes {
		if ctype == *typ.Builtin {
			return false
		}
	}

	return true
}

// CanCast returns true if the resolved type is a builtin type that can be
// directly casted to an equivalent C type OR a record..
func (typ *ResolvedType) CanCast() bool {
	return typ.IsPrimitive() || typ.IsRecord()
}

// IsBuiltin is a convenient function to compare the builtin type.
func (typ *ResolvedType) IsBuiltin(builtin string) bool {
	return typ.Builtin != nil && *typ.Builtin == builtin
}

// HasImport returns true if the ResolvedType has an import.
func (typ *ResolvedType) HasImport() bool {
	return typ.Import != ResolvedTypeImport{}
}

// NeedsNamespace returns true if the returned Go type needs a namespace to be
// referenced properly.
func (typ *ResolvedType) NeedsNamespace(current *gir.NamespaceFindResult) bool {
	switch {
	case typ.Builtin != nil:
		return typ.HasImport()
	case typ.Extern != nil:
		return !typ.Extern.Result.Eq(current)
	default:
		return false // unreachable
	}
}

func (typ *ResolvedType) ptr(sub1 bool) string {
	ptr := typ.Ptr
	if sub1 && ptr > 0 {
		ptr--
	}
	return strings.Repeat("*", int(ptr))
}

// GoType returns the Go type. If the resolved type is a class, then the
// interface type is returned.
func (typ *ResolvedType) GoType(needsNamespace bool) string {
	if typ.Builtin != nil {
		return typ.ptr(false) + *typ.Builtin
	}

	name, _ := typ.Extern.Result.Info()
	name = PascalToGo(name)

	if !needsNamespace {
		return typ.ptr(false) + name
	}

	return typ.ptr(false) + typ.Import.Package + "." + name
}

// CGoType returns the CGo type. Its pointer count does not follow Ptr.
func (typ *ResolvedType) CGoType() string {
	return cgoTypeFromC(typ.CType)
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

var _ TypeResolver = (*NamespaceGenerator)(nil)

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
	if typ.Name == "" || !typ.IsIntrospectable() {
		return nil
	}

	for _, unsupported := range unsupportedCTypes {
		if unsupported == typ.CType {
			return nil
		}
	}

	// Try and dig out the CType if we have none.
	if typ.CType == "" {
		if result := ng.FindType(typ.Name); result != nil {
			typ.CType = result.CType()
		}
		// Last resort.
		if typ.CType == "" {
			typ.CType = ctypeFallback("", typ.Name)
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
	case "GLib.Error":
		return builtinType("", "error", typ)
	case "GLib.List":
		return externGLibType("*List", typ, "GList*")
	case "GLib.SList":
		return externGLibType("*SList", typ, "GSList*")
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

	// Pretend that ignored types don't exist. typ is a copy, so we can do this.
	if ng.mustIgnore(&typ.Name, &typ.CType) {
		return nil
	}

	result := ng.FindType(typ.Name)
	if result == nil {
		warnUnknownType(ng, typ.Name)
		return nil
	}

	// TODO: these checks shouldn't use as much load as they should, since that
	// would lengthen generation time by a lot, which isn't a huge concern, but
	// it's still one. Perhaps we could generate types separately first, and
	// then generate methods and functions afterwards.

	switch {
	case result.Callback != nil:
		cbgen := newCallbackGenerator(ng)
		if !cbgen.Use(*result.Callback) {
			return nil
		}
	case result.Record != nil:
		if !canRecord(ng, *result.Record, nil) {
			return nil
		}
	}

	return typeFromResult(ng.gen, typ, result)
}

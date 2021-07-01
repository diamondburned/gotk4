package types

import (
	"path"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
)

// InternalImportPath is the path to the core import path.
const InternalImportPath = "github.com/diamondburned/gotk4/core"

// Resolved is a resolved type from a given gir.Type.
type Resolved struct {
	// either or
	Extern  *gir.TypeFindResult // optional
	Builtin *string             // optional

	PublImport ResolvedImport
	ImplImport ResolvedImport

	CType string
	GType string
	Ptr   uint8 // used ONLY for the Go type.
}

// ResolvedImport is a single import for the resolved type.
type ResolvedImport struct {
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
func builtinType(imp, typ string, girType gir.Type) *Resolved {
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

	resolvedImport := ResolvedImport{
		Path:    imp,
		Package: pkg,
	}

	return &Resolved{
		Builtin:    &typ,
		PublImport: resolvedImport,
		ImplImport: resolvedImport,
		GType:      girType.Name,
		CType:      CTypeFallback(girType.CType, girType.Name),
		Ptr:        ptr,
	}
}

// externGLibType returns an external GLib type from gotk3.
func externGLibType(goType string, typ gir.Type, ctyp string) *Resolved {
	if typ.CType != "" {
		ctyp = typ.CType
	}

	implImport := ResolvedImport{
		Path:    "github.com/gotk3/gotk3/glib",
		Package: "externglib",
	}
	publImport := implImport

	switch strings.ReplaceAll(goType, "*", "") {
	case "InitiallyUnowned", "Object":
		publImport = ResolvedImport{
			Path:    InternalImportPath + "/gextras",
			Package: "gextras",
		}
	}

	goType = "externglib." + strings.TrimPrefix(goType, "*")

	return &Resolved{
		Builtin:    &goType,
		ImplImport: implImport,
		PublImport: publImport,
		GType:      typ.Name,
		CType:      ctyp,
		Ptr:        uint8(dereferenceOffset(int(countPtrs(typ, nil)), goType)),
	}
}

// dereferenceOffset subtracts 1 from ptrs if the ctype does not have a pointer.
// It is better explained with an example:
//
// If the C type is *GObject, then this wouldn't subtract anything, but if the C
// type is a gpointer, then we'd be subtracting 1. This code is similar to the
// one in TypeResolver.
func dereferenceOffset(ptrs int, typ string) int {
	if ptrs == 0 {
		return ptrs
	}

	count := strings.Count(typ, "*")
	if count > 1 {
		count = 1
	}

	return ptrs - (1 - count)
}

// typeFromResult creates a resolved type from the given type result.
func typeFromResult(gen Generator, typ gir.Type, result *gir.TypeFindResult) *Resolved {
	if typ.CType == "" {
		// Try to fill this back with the right CType. This might be wrong in
		// some cases, but it should work in cases where CType isn't there.
		typ.CType = result.CType()
	}
	if typ.CType == "" {
		gen.Logln(logger.Error, "type", typ.Name, "missing CType")
		return nil
	}

	resolvedImport := ResolvedImport{
		Path:    gen.ModPath(result.Namespace),
		Package: gir.GoNamespace(result.Namespace),
	}

	ptr := countPtrs(typ, result)

	// Always use internal types (like GVariant) by reference and not value,
	// since Go will refuse to allocate them.
	if record, ok := result.Type.(*gir.Record); ok {
		if RecordIsOpaque(*record) && ptr == 0 {
			ptr++
		}
	}

	return &Resolved{
		Extern:     result,
		ImplImport: resolvedImport,
		PublImport: resolvedImport,
		GType:      typ.Name,
		CType:      typ.CType,
		Ptr:        ptr,
	}
}

// func TypeFromResult(ns *gir.NamespaceFindResult, res gir.TypeFindResult) *ResolvedType {
// 	name, ctype := res.Info()
// 	return typeFromResult(ng.gen, gir.Type{Name: name, CType: ctype}, &res)
// }

// IsExternGLib checks that the ResolvedType is exactly the gotk3/glib type with
// the given name. Pointers are not compared.
func (typ *Resolved) IsExternGLib(glibType string) bool {
	// Use ImplImport for comparison, so we're not comparing gextras types.
	if typ.Builtin == nil || typ.ImplImport.Path != "github.com/gotk3/gotk3/glib" {
		return false
	}

	thisType := *typ.Builtin
	thisType = strings.ReplaceAll(thisType, "*", "")
	thisType = strings.TrimPrefix(thisType, typ.ImplImport.Package)
	thisType = strings.TrimPrefix(thisType, ".")

	return thisType == glibType
}

// IsCallback returns true if the current ResolvedType is a callback.
func (typ *Resolved) IsCallback() bool {
	if typ.Extern == nil {
		return false
	}
	_, ok := typ.Extern.Type.(*gir.Callback)
	return ok
}

// IsRecord returns true if the current ResolvedType is a record.
func (typ *Resolved) IsRecord() bool {
	if typ.Extern == nil {
		return false
	}
	_, ok := typ.Extern.Type.(*gir.Record)
	return ok
}

// IsInterface returns true if the current ResolvedType is an interface.
func (typ *Resolved) IsInterface() bool {
	if typ.Extern == nil {
		return false
	}
	_, ok := typ.Extern.Type.(*gir.Interface)
	return ok
}

// IsClass returns true if the current ResolvedType is a class.
func (typ *Resolved) IsClass() bool {
	if typ.Extern == nil {
		return false
	}
	_, ok := typ.Extern.Type.(*gir.Class)
	return ok
}

// IsPrimitive returns true if the resolved type is a builtin type that can be
// directly casted to an equivalent C type OR a record..
func (typ *Resolved) IsPrimitive() bool {
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
func (typ *Resolved) CanCast() bool {
	return typ.IsPrimitive() || typ.IsRecord()
}

// IsBuiltin is a convenient function to compare the builtin type.
func (typ *Resolved) IsBuiltin(builtin string) bool {
	return typ.Builtin != nil && *typ.Builtin == builtin
}

// HasImport returns true if the ResolvedType has an import.
func (typ *Resolved) HasImport() bool {
	var zeroi ResolvedImport
	return typ.ImplImport != zeroi || typ.PublImport != zeroi
}

// HasPointer returns true if the type being resolved has a pointer. This is
// useful for array passing from Go memory to C memory.
func (typ *Resolved) HasPointer(gen Generator) bool {
	if typ == nil {
		// Probably unknown.
		return true
	}

	if typ.Builtin != nil {
		return !typ.IsPrimitive()
	}

	switch v := typ.Extern.Type.(type) {
	case *gir.Alias:
		return ResolveName(gen, v.Name).HasPointer(gen)

	case
		*gir.Union, // TODO: handle unions
		*gir.Class,
		*gir.Callback,
		*gir.Function,
		*gir.Interface:

		return true

	case
		*gir.Enum,
		*gir.Bitfield:

		return false

	case *gir.Record:
		for _, field := range v.Fields {
			// If field is not a regular type, then it's probably an array or
			// whatever, which means a pointer.
			if field.Type == nil {
				return true
			}

			if Resolve(gen, *field.Type).HasPointer(gen) {
				return true
			}
		}

		return false
	}

	// Unknown type; assume there's a pointer.
	return true
}

// GoImplType is a convenient function around ResolvedType.ImplType.
func GoImplType(gen Generator, resolved *Resolved) string {
	return resolved.ImplType(resolved.NeedsNamespace(gen.Namespace()))
}

// GoPublicType is a convenient function around ResolvedType.ImplType.
func GoPublicType(gen Generator, resolved *Resolved) string {
	return resolved.PublicType(resolved.NeedsNamespace(gen.Namespace()))
}

// NeedsNamespace returns true if the returned Go type needs a namespace to be
// referenced properly.
func (typ *Resolved) NeedsNamespace(current *gir.NamespaceFindResult) bool {
	switch {
	case typ.Builtin != nil:
		return typ.HasImport()
	case typ.Extern != nil:
		return !typ.Extern.Eq(current)
	default:
		return false // unreachable
	}
}

func (typ *Resolved) ptr(sub1 bool) string {
	ptr := typ.Ptr
	if sub1 && ptr > 0 {
		ptr--
	}
	return strings.Repeat("*", int(ptr))
}

// ImplType returns the implementation type. This is only different to
// PublicType as far as classes go: the returned type is the unexported
// implementation type.
func (typ *Resolved) ImplType(needsNamespace bool) string {
	if typ.Builtin != nil {
		return *typ.Builtin
	}

	name := typ.Extern.Name()
	name = strcases.PascalToGo(name)

	switch typ.Extern.Type.(type) {
	case *gir.Class, *gir.Interface:
		name = strcases.UnexportPascal(name)
	}

	if !needsNamespace {
		return typ.ptr(false) + name
	}

	return typ.ptr(false) + typ.ImplImport.Package + "." + name
}

// PublicType returns the public type. If the resolved type is a class, then the
// interface type is returned.
func (typ *Resolved) PublicType(needsNamespace bool) string {
	switch {
	case
		typ.IsExternGLib("InitiallyUnowned"),
		typ.IsExternGLib("Object"):

		// TODO: there should be a better way to do this; one that adds imports.
		return typ.ptr(true) + "gextras.Objector"
	}

	if typ.Builtin != nil {
		return typ.ptr(false) + *typ.Builtin
	}

	name := typ.Extern.Name()
	name = strcases.PascalToGo(name)

	// Classes have a pointer in C, but we implement it as an interface in Go.
	ptrStr := typ.ptr(typ.IsClass() || typ.IsInterface())

	if !needsNamespace {
		return ptrStr + name
	}

	return ptrStr + typ.PublImport.Package + "." + name
}

// WrapName returns the name of the wrapper function. It only works for external
// types; calling this on a built-in ResolvedType will return an empty string.
func (typ *Resolved) WrapName(needsNamespace bool) string {
	if typ.Extern == nil {
		return ""
	}

	name := typ.Extern.Name()
	name = strcases.PascalToGo(name)

	wrapName := "Wrap" + name
	if needsNamespace {
		// The wrapper is all exported, so it's probably public. In reality it
		// doesn't matter, since all extern types will have the same imports.
		wrapName = typ.PublImport.Package + "." + wrapName
	}

	return wrapName
}

// CGoType returns the CGo type. Its pointer count does not follow Ptr.
func (typ *Resolved) CGoType() string {
	return CGoTypeFromC(typ.CType)
}

// UnsupportedCTypes is the list of unsupported C types, either because it is
// not yet supported or will never be supported due to redundancy or else.
var UnsupportedCTypes = []string{
	"tm*", // requires time.h
	"va_list",
}

// ResolveName resolves the given GIR type name. The resolved type will
// always have no pointer.
func ResolveName(gen Generator, girType string) *Resolved {
	return Resolve(gen, gir.Type{Name: girType})
}

// Resolve resolves the given type from the GIR type field. It returns nil if
// the type is not known. It does not recursively traverse the type.
func Resolve(gen Generator, typ gir.Type) *Resolved {
	if typ.Name == "" || !typ.IsIntrospectable() {
		return nil
	}

	for _, unsupported := range UnsupportedCTypes {
		if unsupported == typ.CType {
			return nil
		}
	}

	var result *gir.TypeFindResult

	// Try and dig out the CType if we have none.
	if typ.CType == "" {
		if result = Find(gen, typ.Name); result != nil {
			typ.CType = result.CType()
		}
		// Last resort.
		if typ.CType == "" {
			typ.CType = CTypeFallback("", typ.Name)
		}
	}

	if prim, ok := girToBuiltin[typ.Name]; ok {
		return builtinType("", prim, typ)
	}

	// Resolve the unknown namespace that is GLib and primitive types.
	switch EnsureNamespace(gen.Namespace(), typ.Name) {
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
		gen.Logln(logger.Unusual, "type name", typ.Name, "missing CType")
		return nil
	}

	// Pretend that ignored types don't exist. typ is a copy, so we can do this.
	if Filter(gen, typ.Name, typ.CType) {
		return nil
	}

	if result == nil {
		result = Find(gen, typ.Name)
	}

	if result == nil {
		gen.Logln(logger.Debug, "unknown type resolved", strconv.Quote(typ.Name))
		return nil
	}

	// TODO: these checks shouldn't use as much load as they should, since that
	// would lengthen generation time by a lot, which isn't a huge concern, but
	// it's still one. Perhaps we could generate types separately first, and
	// then generate methods and functions afterwards.

	switch /* v := */ result.Type.(type) {
	case *gir.Record:
		panic("TODO: resolve.go/Record")
		// if !canRecord(ng, *result.Record, nil) {
		// 	return nil
		// }

	case *gir.Callback:
		panic("TODO resolve.go/Callback")
		// cbgen := newCallbackGenerator(ng)
		// if !cbgen.Use(*result.Callback) {
		// 	return nil
		// }
		// source = v.SourcePosition
	}

	return typeFromResult(gen, typ, result)
}

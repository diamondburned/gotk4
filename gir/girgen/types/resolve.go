package types

import (
	"path"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
)

// InternalImportPath is the path to the core import path.
const InternalImportPath = "github.com/diamondburned/gotk4/pkg/core"

// Resolved is a resolved type from a given gir.Type.
type Resolved struct {
	// either or
	Extern  *gir.TypeFindResult // optional
	Builtin *string             // optional
	Aliased *Resolved           // optional

	// TODO: move file.Header over to types.Header.
	// TODO: replace {Publ,Impl}Import with types.Header to allow HashTable.

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
var goContainerTypes = map[string]struct{}{
	"error":       {},
	"string":      {},
	"interface{}": {},
}

// BuiltinType is a convenient function to make a new built-in *Resolved.
func BuiltinType(imp, typ string, girType gir.Type) *Resolved {
	return builtinType(imp, typ, girType)
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
		if strings.HasPrefix(typ, "interface{") {
			ptr--
			goto subtracted
		}
		if _, ok := goContainerTypes[typ]; ok {
			ptr--
			goto subtracted
		}
	subtracted:
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
	if typ.CType == "" {
		typ.CType = ctyp
	}

	imp := ResolvedImport{
		Path:    "github.com/diamondburned/gotk4/pkg/core/glib",
		Package: "externglib",
	}

	var ptr uint8

	if typ.CType != "" {
		ptr = countPtrs(typ, nil)
	} else {
		ptr = uint8(strings.Count(ctyp, "*"))
	}
	// Edge case.
	if ptr > 0 && goType == "AnyClosure" {
		ptr--
	}

	goType = "externglib." + strings.TrimPrefix(goType, "*")

	return &Resolved{
		Builtin:    &goType,
		ImplImport: imp,
		PublImport: imp,
		GType:      typ.Name,
		CType:      typ.CType,
		Ptr:        ptr,
	}
}

//// dereferenceOffset subtracts 1 from ptrs if the ctype does not have a pointer.
//// It is better explained with an example:
////
//// If the C type is *GObject, then this wouldn't subtract anything, but if the C
//// type is a gpointer, then we'd be subtracting 1. This code is similar to the
//// one in TypeResolver.
//func dereferenceOffset(ptrs int, typ string) int {
//	if ptrs == 0 {
//		return ptrs
//	}

//	count := strings.Count(typ, "*")
//	if count > 1 {
//		count = 1
//	}

//	return ptrs - (1 - count)
//}

// typeFromResult creates a resolved type from the given type result.
func typeFromResult(gen FileGenerator, typ gir.Type, result *gir.TypeFindResult) *Resolved {
	if typ.CType == "" {
		// Try to fill this back with the right CType. This might be wrong in
		// some cases, but it should work in cases where CType isn't there.
		typ.CType = result.CType()
	}
	if typ.CType == "" {
		gen.Logln(logger.Error, "type", typ.Name, "missing CType")
		return nil
	}

	var resolvedImport ResolvedImport
	var ignoreOpaque bool

	switch result.Namespace.Name {
	case "cairo":
		// gotk3/cairo structs already contain a pointer.
		ignoreOpaque = true
		resolvedImport = ResolvedImport{
			Path:    "github.com/diamondburned/gotk4/pkg/cairo",
			Package: "cairo",
		}
	default:
		resolvedImport = ResolvedImport{
			Path:    gen.ModPath(result.Namespace),
			Package: gir.GoNamespace(result.Namespace),
		}
	}

	ptr := countPtrs(typ, result)

	if isGpointer(typ.CType, typ.Name, int(ptr)) {
		ptr++
	}

	// Always use internal types (like GVariant) by reference and not value,
	// since Go will refuse to allocate them.
	if record, ok := result.Type.(*gir.Record); ok {
		if !ignoreOpaque && RecordIsOpaque(*record) && ptr == 0 {
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

// TypeFromResult is meant to be used by an external package to generate a
// Resolved from existing type information.
func TypeFromResult(gen FileGenerator, v interface{}) *Resolved {
	res := gir.TypeFindResult{
		NamespaceFindResult: gen.Namespace(),
		Type:                v,
	}

	typ := gir.Type{
		Name:  res.Name(),
		CType: res.CType(),
	}
	if typ.CType == "" {
		return nil
	}

	return typeFromResult(gen, typ, &res)
}

// IsExternGLib checks that the ResolvedType is exactly the gotk3/glib type with
// the given name. Pointers are not compared.
func (typ *Resolved) IsExternGLib(glibType string) bool {
	// Use ImplImport for comparison, so we're not comparing gextras types.
	if typ.Builtin == nil || typ.ImplImport.Path != "github.com/diamondburned/gotk4/pkg/core/glib" {
		return false
	}

	thisType := *typ.Builtin
	thisType = strings.ReplaceAll(thisType, "*", "")
	thisType = strings.TrimPrefix(thisType, typ.ImplImport.Package)
	thisType = strings.TrimPrefix(thisType, ".")

	return thisType == glibType
}

var gpointerTypes = map[string]struct{}{
	"gpointer":      {},
	"gconstpointer": {},
}

// IsGpointer returns true if the given type is a gpointer or a pointer to it.
func IsGpointer(ctype string) bool {
	_, is := gpointerTypes[CleanCType(ctype, true)]
	return is
}

// IsGpointer returns true if the given type is a gpointer type.
func (typ *Resolved) IsGpointer() bool {
	return isGpointer(typ.CType, typ.GType, int(typ.Ptr)-1)
}

func isGpointer(ctype, gtype string, ptr int) bool {
	_, is := gpointerTypes[CleanCType(ctype, true)]
	if ptr > 0 {
		// If the CType is a pointer, then we make sure that this isn't just a C
		// pointer type masked into a gpointer.
		_, is2 := gpointerTypes[gtype]
		is = is && is2
	}
	return is
}

// CanNil returns true if the Go type can be nil.
func (typ *Resolved) CanNil() bool {
	if typ.IsClass() || typ.IsInterface() || typ.IsCallback() {
		return true
	}

	return typ.Ptr > 0
}

func (typ *Resolved) externType() *gir.TypeFindResult {
	if typ.Aliased != nil {
		return typ.Aliased.Extern
	}
	return typ.Extern
}

// IsCallback returns true if the current ResolvedType is a callback.
func (typ *Resolved) IsCallback() bool {
	t := typ.externType()
	if t == nil {
		return false
	}

	_, ok := t.Type.(*gir.Callback)
	return ok
}

// IsUnion returns true if the current ResolvedType is a union.
func (typ *Resolved) IsUnion() bool {
	t := typ.externType()
	if t == nil {
		return false
	}

	_, ok := t.Type.(*gir.Union)
	return ok
}

// IsRecord returns true if the current ResolvedType is a record.
func (typ *Resolved) IsRecord() bool {
	t := typ.externType()
	if t == nil {
		return false
	}

	_, ok := t.Type.(*gir.Record)
	return ok
}

// IsInterface returns true if the current ResolvedType is an interface.
func (typ *Resolved) IsInterface() bool {
	t := typ.externType()
	if t == nil {
		return false
	}

	_, ok := t.Type.(*gir.Interface)
	return ok
}

// IsClass returns true if the current ResolvedType is a class.
func (typ *Resolved) IsClass() bool {
	t := typ.externType()
	if t == nil {
		return false
	}

	_, ok := t.Type.(*gir.Class)
	return ok
}

// IsAbstract returns true if the resolved type is an interface or an abstract
// class.
func (typ *Resolved) IsAbstract() bool {
	if typ.IsInterface() {
		return true
	}

	if typ.IsClass() {
		return typ.externType().Type.(*gir.Class).Abstract
	}

	return false
}

// IsEnumOrBitfield returns true if the resolved type is an external enum or
// bitfield type.
func (typ *Resolved) IsEnumOrBitfield() bool {
	t := typ.externType()
	if t == nil {
		return false
	}

	_, ok1 := t.Type.(*gir.Enum)
	_, ok2 := t.Type.(*gir.Bitfield)
	return ok1 || ok2
}

func (typ *Resolved) PublicIsInterface() bool {
	if typ.Builtin != nil {
		return typ.IsExternGLib("Object") || typ.IsExternGLib("InitiallyUnowned")
	}
	return typ.IsAbstract()
}

// IsPrimitive returns true if the resolved type is a builtin type that can be
// directly casted to an equivalent C type OR a record.
func (typ *Resolved) IsPrimitive() bool {
	if typ.Builtin == nil {
		return false
	}

	if typ.HasImport() {
		return false
	}

	_, ok := goContainerTypes[*typ.Builtin]
	if ok {
		return false
	}

	return true
}

// IsContainerBuiltin returns true if the resolved type is a built-in Go
// container type (like string, error or interface{}).
func (typ *Resolved) IsContainerBuiltin() bool {
	if typ.Builtin == nil {
		return false
	}

	_, ok := goContainerTypes[*typ.Builtin]
	if ok {
		return true
	}

	return false
}

// CanCast returns true if the resolved type is a builtin type that can be
// directly casted to an equivalent C type OR a record..
func (typ *Resolved) CanCast(gen FileGenerator) bool {
	if typ.IsPrimitive() {
		// Only allow casting if the type has the same size as the Go or C
		// equivalent.
		_, ok := noCasting[typ.GType]
		return !ok
	}

	if typ.IsEnumOrBitfield() {
		return true
	}

	// We can only directly cast the struct if it only contains primitives.
	return typ.IsRecord() && !typ.HasPointer(gen)
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
func (typ *Resolved) HasPointer(gen FileGenerator) bool {
	if typ == nil {
		// Probably unknown.
		return true
	}

	if typ.Builtin != nil {
		return !typ.IsPrimitive()
	}

	switch v := typ.Extern.Type.(type) {
	case *gir.Alias:
		if typ.Aliased != nil {
			return typ.Aliased.HasPointer(gen)
		}
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
			// whatever, which means a pointer. If it's a bitfield OR a private
			// field, then it's best that we don't touch it.
			if field.Type == nil || field.Bits != 0 || field.Private {
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

// FullGType returns the GType with the namespace.
//
// Deprecated: Use typ.GType.
func (typ *Resolved) FullGType() string {
	return typ.GType
}

// GoImplType is a convenient function around ResolvedType.ImplType.
func GoImplType(gen FileGenerator, resolved *Resolved) string {
	return resolved.ImplType(resolved.NeedsNamespace(gen.Namespace()))
}

// GoPublicType is a convenient function around ResolvedType.ImplType.
func GoPublicType(gen FileGenerator, resolved *Resolved) string {
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

// Name returns the type name without the namespace or pointer.
func (typ *Resolved) Name() string {
	return typ.ImplName()
}

// ImplName returns the implementation type name.
func (typ *Resolved) ImplName() string {
	if typ.Builtin != nil {
		parts := strings.Split(*typ.Builtin, ".")
		return strings.ReplaceAll(parts[len(parts)-1], "*", "")
	}

	name := strcases.PascalToGo(typ.Extern.Name())
	if name == "Object" {
		// Avoid collision with externglib.Object, since that might be embedded
		// into an implementation struct.
		name = "ObjectClass"
	}

	return name
}

// PublicName returns the public type name.
func (typ *Resolved) PublicName() string {
	if typ.Builtin != nil {
		switch {
		case typ.IsExternGLib("InitiallyUnowned"), typ.IsExternGLib("Object"):
			return "Objector"
		}

		parts := strings.Split(*typ.Builtin, ".")
		return strings.ReplaceAll(parts[len(parts)-1], "*", "")
	}

	name := strcases.PascalToGo(typ.Extern.Name())

	switch typ.Extern.Type.(type) {
	case *gir.Class:
		// Avoid collision with externglib.Object, since that might be embedded
		// into an implementation struct.
		if name == "Object" {
			name = "ObjectClass"
		}
		name = strcases.Interfacify(name)

	case *gir.Interface:
		if name == "Object" {
			name = "ObjectInterface"
		}
		name = strcases.Interfacify(name)
	}

	return name
}

// ImplType returns the implementation type. This is only different to
// PublicType as far as classes go: the returned type is the unexported
// implementation type.
func (typ *Resolved) ImplType(needsNamespace bool) string {
	if typ.Builtin != nil {
		// Always use a pointer for Object.
		if typ.IsExternGLib("Object") && typ.Ptr == 0 {
			return "*" + *typ.Builtin
		}

		return typ.ptr(false) + *typ.Builtin
	}

	name := typ.Name()

	if !needsNamespace {
		return typ.ptr(false) + name
	}
	return typ.ptr(false) + typ.ImplImport.Package + "." + name
}

// PublicType returns the public type. If the resolved type is a class, then the
// interface type is returned.
func (typ *Resolved) PublicType(needsNamespace bool) string {
	switch {
	case typ.IsExternGLib("InitiallyUnowned"), typ.IsExternGLib("Object"):
		if !needsNamespace {
			return typ.ptr(true) + "Objector"
		}

		return typ.ptr(true) + "externglib.Objector"
	}

	if typ.Builtin != nil {
		if !needsNamespace {
			parts := strings.Split(*typ.Builtin, ".")
			return typ.ptr(false) + parts[len(parts)-1]
		}

		return typ.ptr(false) + *typ.Builtin
	}

	name := typ.PublicName()
	// Classes have a pointer in C, but we implement it as an interface in Go.
	ptrStr := typ.ptr(typ.PublicIsInterface())

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

	wrapName := "wrap" + name
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
	"tm", // requires time.h
	"va_list",
}

// ResolveName resolves the given GIR type name. The resolved type will
// always have no pointer.
func ResolveName(gen FileGenerator, girType string) *Resolved {
	typ := gir.Type{
		Name: girType,
	}

	if result := Find(gen, girType); result != nil {
		typ.CType = result.CType()

		typ.Introspectable = new(bool)
		*typ.Introspectable = result.IsIntrospectable()
	}

	return Resolve(gen, typ)
}

// BuiltinHandledTypes contains types manually handled by Resolve and
// typeconv.Converter, as well as types that are never supposed to be handled.
var BuiltinHandledTypes = []FilterMatcher{
	// These are already manually covered in the girgen code; they are
	// provided by package gotk3/glib.
	AbsoluteFilter("GLib.Error"),
	// Ignore generating everything in GObject, but allow resolving its
	// types.
	RegexFilter("GObject..*"),
	// This is not supported by Go. We might be able to support it in
	// the future using a 16-byte data structure, but the actual size
	// isn't well defined as far as I know.
	AbsoluteFilter("*.long double"),
	// Special marking for internal types from GLib (apparently for
	// glib:get-type).
	AbsoluteFilter("C.intern"),
	// Ignore all of ByteArray's methods and functions. Use the C namespace,
	// because record methods aren't filtered properly.
	RegexFilter(`C.g_byte_array_.*`),
	// Already covered by externglib.
	AbsoluteFilter("GLib.Type"),
}

// Resolve resolves the given type from the GIR type field. It returns nil if
// the type is not known. It does not recursively traverse the type.
func Resolve(gen FileGenerator, typ gir.Type) *Resolved {
	if typ.Name == "" || !typ.IsIntrospectable() {
		return nil
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

	if typ.CType != "" {
		for _, unsupported := range UnsupportedCTypes {
			if unsupported == typ.CType {
				return nil
			}
		}
	}

	if prim, ok := girToBuiltin[typ.Name]; ok {
		if prim == "" {
			// void type, exit.
			return nil
		}
		return builtinType("", prim, typ)
	}

	// Treat actual gpointer types as the pseudo cgo.Handle type. Check both GIR
	// and C type to avoid masked gpointer types that aren't actually arbitrary.
	if typ.Name == "gpointer" && IsGpointer(typ.CType) {
		return builtinType("runtime/cgo", "Handle", typ)
	}

	// Fill namespace.
	typ.Name = EnsureNamespace(gen.Namespace(), typ.Name)

	// Resolve the unknown namespace that is GLib and primitive types.
	switch typ.Name {
	// TODO: ignore field
	case "GLib.Error":
		return builtinType("", "error", typ)
	case "GLib.List":
		return externGLibType("*List", typ, "GList*")
	case "GLib.SList":
		return externGLibType("*SList", typ, "GSList*")
	// TODO: include GLib.HashTable
	case "GObject.Type", "GType":
		return externGLibType("Type", typ, "GType")
	case "GObject.Value", "GValue": // inconsistency???
		return externGLibType("*Value", typ, "GValue*")
	case "GObject.Object":
		return externGLibType("*Object", typ, "GObject*")
	case "GObject.InitiallyUnowned":
		return externGLibType("InitiallyUnowned", typ, "GInitiallyUnowned*")
	case "GObject.Closure":
		return externGLibType("AnyClosure", typ, "GClosure*")
	}

	// CType is required here so we can properly account for pointers.
	if typ.CType == "" {
		gen.Logln(logger.Unusual, "type name", typ.Name, "missing CType")
		return nil
	}

	// Pretend that ignored types don't exist. typ is a copy, so we can do this.
	if Filter(gen, typ.Name, typ.CType) {
		gen.Logln(logger.Debug, "ignored type", typ.Name)
		return nil
	}

	if result == nil {
		result = Find(gen, typ.Name)
	}

	if result == nil {
		gen.Logln(logger.Debug, "unknown type resolved", typ.Name)
		return nil
	}

	// TODO: these checks shouldn't use as much load as they should, since that
	// would lengthen generation time by a lot, which isn't a huge concern, but
	// it's still one. Perhaps we could generate types separately first, and
	// then generate methods and functions afterwards.

	resolved := typeFromResult(gen, typ, result)

	if !gen.CanGenerate(resolved) {
		gen.Logln(logger.Debug, "cannot generate type", typ.Name, resolved.PublicType(true))
		return nil
	}

	if alias, ok := resolved.Extern.Type.(*gir.Alias); ok {
		resolved.Aliased = Resolve(gen, alias.Type)
	}

	return resolved
}

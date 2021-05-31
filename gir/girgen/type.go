package girgen

import (
	"fmt"
	"path"
	"strings"
	"sync"

	"github.com/diamondburned/gotk4/gir"
	"golang.org/x/sync/singleflight"
)

func countPtrs(typ gir.Type, result *gir.TypeFindResult) uint8 {
	ptr := uint8(strings.Count(typ.CType, "*"))

	if ptr > 0 {
		// Edge case: a string is a gchar*, so we don't need a pointer.
		if typ.Name == "utf8" {
			ptr--
			goto ret
		}

		if result != nil {
			// Edge case: interfaces must not be pointers. We should still
			// sometimes allow for pointers to interfaces, if needed, but this
			// likely won't work.
			switch {
			case result.Interface != nil:
				fallthrough
			case result.Class != nil:
				ptr--
				goto ret
			}
		}
	}

ret:
	return ptr
}

var ctypePrefixEraser = strings.NewReplacer(
	"const", "",
	"volatile", "",
)

// movePtr moves the same number of pointers from the given orig string into
// another string.
func movePtr(orig, into string) string {
	ptr := strings.Count(orig, "*")
	return strings.Repeat("*", ptr) + into
}

// anyTypeIsVoid returns true if AnyType is a void type.
func anyTypeIsVoid(any gir.AnyType) bool {
	return any.Type != nil && any.Type.Name == "none"
}

// anyTypeCGo returns the CGo type for a GIR AnyType. An empty string is
// returned if none is made.
func anyTypeCGo(any gir.AnyType) string {
	return cgoType(anyTypeC(any))
}

// anyTypeC returns the C type for a GIR AnyType. An empty string is returned if
// none is made.
func anyTypeC(any gir.AnyType) string {
	switch {
	case any.Array != nil:
		return oneOrOtherStr(any.Array.CType, any.Array.Name)
	case any.Type != nil:
		return oneOrOtherStr(any.Type.CType, any.Type.Name)
	default:
		return ""
	}
}

// oneOrOtherStr returns one or the other string, whichever is not empty.
func oneOrOtherStr(one, other string) string {
	if one != "" {
		return one
	}
	return other
}

// cgoType turns the given C type into CGo.
func cgoType(cType string) string {
	oldType := cType
	cType = ctypePrefixEraser.Replace(cType)
	cType = strings.ReplaceAll(cType, "*", "")
	cType = strings.TrimSpace(cType)

	if replace, ok := cgoPrimitiveTypes[cType]; ok {
		cType = replace
	}

	return movePtr(oldType, "C."+cType)
}

// returnIsVoid returns true if the return type is void.
func returnIsVoid(ret *gir.ReturnValue) bool {
	return ret == nil || (ret != nil && anyTypeIsVoid(ret.AnyType))
}

// girToBuiltin maps the given GIR primitive type to a Go builtin type.
var girToBuiltin = map[string]string{
	"none":        "",
	"gboolean":    "bool",
	"gfloat":      "float32",
	"gdouble":     "float64",
	"long double": "float64",
	"gint":        "int",
	"gssize":      "int",
	"gint8":       "int8",
	"gint16":      "int16",
	"gshort":      "int16",
	"gint32":      "int32",
	"glong":       "int32",
	"int32":       "int32",
	"gint64":      "int64",
	"guint":       "uint",
	"gsize":       "uint",
	"guchar":      "byte",
	"gchar":       "byte",
	"guint8":      "byte", // some weird cases
	"guint16":     "uint16",
	"gushort":     "uint16",
	"guint32":     "uint32",
	"gulong":      "uint32",
	"gunichar":    "uint32",
	"guint64":     "uint64",
	"utf8":        "string",
	"filename":    "string",
}

// girPrimitiveGo returns Go primitive types that can be copied by-value without
// doing any pointer work. It returns an empty string if there's none.
func girPrimitiveGo(typ string) string {
	gp, ok := girToBuiltin[typ]
	if !ok || gp == "string" {
		return ""
	}
	return gp
}

// cgoPrimitiveTypes contains edge cases for referencing C primitive types from
// CGo.
//
// See https://gist.github.com/zchee/b9c99695463d8902cd33.
var cgoPrimitiveTypes = map[string]string{
	"long double":  "longdouble",
	"unsigned int": "uint",
}

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

var (
	typeCache  sync.Map
	typeFlight singleflight.Group
)

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
		CType:   oneOrOtherStr(girType.CType, girType.Name),
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
		gen.logln(logError, "type", typ.Name, "missing CType")
		return nil
	}

	pkg := gir.GoNamespace(result.Namespace)

	return &ResolvedType{
		Extern: &ExternType{
			Result: result,
		},
		Import:  gen.ImportPath(pkg),
		Package: pkg,
		GType:   typ.Name,
		CType:   typ.CType,
		Ptr:     countPtrs(typ, result),
	}
}

// TypeFromResult creates a new ResolvedType from the given type find result.
// This function is mostly useful for generating from an existing GIR value.
func TypeFromResult(ng *NamespaceGenerator, res gir.TypeFindResult) *ResolvedType {
	res.NamespaceFindResult = ng.current
	name, ctype := res.Info()
	t := typeFromResult(ng.gen, gir.Type{Name: name, CType: ctype}, &res)
	if t == nil {
		ng.logln(logError, "typeFromResult returned nil for", name)
	}
	return t
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
	return cgoType(typ.CType)
}

// TypeHasPointer returns true if the type being resolved has a pointer. This is
// useful for array passing from Go memory to C memory.
func (ng *NamespaceGenerator) TypeHasPointer(typ *ResolvedType) bool {
	if typ == nil {
		// Probably unknown.
		return true
	}

	if typ.Builtin != nil {
		return !typ.IsPrimitive()
	}

	res := typ.Extern.Result

	switch {
	case res.Alias != nil:
		return ng.TypeHasPointer(ng.ResolveTypeName(res.Alias.Name))

	case
		res.Class != nil,
		res.Callback != nil,
		res.Function != nil,
		res.Interface != nil:
		return true

	case res.Union != nil:
		return true // TODO: handle unions

	case
		res.Enum != nil,
		res.Bitfield != nil:
		return false

	case res.Record != nil:
		for _, field := range res.Record.Fields {
			// If field is not a regular type, then it's probably an array or
			// whatever, which means a pointer.
			if field.Type == nil {
				return true
			}

			if ng.TypeHasPointer(ng.ResolveType(*field.Type)) {
				return true
			}
		}

		return false
	}

	// Unknown type; assume there's a pointer.
	return true
}

// PublicType returns the generated public Go type of the given resolved type.
func (ng *NamespaceGenerator) PublicType(resolved *ResolvedType) string {
	return resolved.PublicType(resolved.NeedsNamespace(ng.current))
}

// ImplType returns the generated implementation Go type of the given resolved type.
func (ng *NamespaceGenerator) ImplType(resolved *ResolvedType) string {
	return resolved.ImplType(resolved.NeedsNamespace(ng.current))
}

// arrayType generates the Go type signature for the given array.
func (ng *NamespaceGenerator) resolveArrayType(array gir.Array, pub bool) (string, bool) {
	arrayPrefix := "[]"
	if array.FixedSize > 0 {
		arrayPrefix = fmt.Sprintf("[%d]", array.FixedSize)
	}

	child, _ := ng.ResolveAnyType(array.AnyType, pub)
	// There can't be []void, so this check ensures there can only be valid
	// array types.
	if child == "" {
		return "", false
	}

	return arrayPrefix + child, true
}

// ResolveAnyType generates the Go type signature for the AnyType union. An
// empty string returned is an invalid type. If pub is true, then the returned
// string will use public interface types for classes instead of implementation
// types.
func (ng *NamespaceGenerator) ResolveAnyType(any gir.AnyType, pub bool) (string, bool) {
	switch {
	case any.Array != nil:
		return ng.resolveArrayType(*any.Array, pub)
	case any.Type != nil:
		return ng.ResolveToGoType(*any.Type, pub)
	}

	// Probably varargs, ignore because Cgo.
	return "", false
}

// ResolveToGoType is a convenient function that wraps around ResolveType and
// returns the Go type.
func (ng *NamespaceGenerator) ResolveToGoType(typ gir.Type, pub bool) (string, bool) {
	resolved := ng.ResolveType(typ)
	if resolved == nil {
		return "", false
	}

	if pub {
		return ng.PublicType(resolved), true
	}

	return ng.ImplType(resolved), true
}

// ResolveTypeName resolves the given GIR type name. The resolved type will
// always have no pointer.
func (ng *NamespaceGenerator) ResolveTypeName(girType string) *ResolvedType {
	var cType string

	// FindType is cached, so we can afford to do this.
	result := ng.gen.Repos.FindType(ng.current.Namespace.Name, girType)
	if result != nil {
		// Use the CType result ONLY. The returned Name from Info does NOT have
		// the namespace prepended.
		_, cType = result.Info()
	}

	return ng.ResolveType(gir.Type{
		Name:  girType,
		CType: cType,
	})
}

// ResolveType resolves the given type from the GIR type field. It returns nil
// if the type is not known. It does not recursively traverse the type.
func (ng *NamespaceGenerator) ResolveType(typ gir.Type) *ResolvedType {
	// Skip the cache if the type does not have a CType. We want to do this to
	// prevent filling the cache up with incomplete GIR types.
	if typ.CType == "" {
		resolved := ng.resolveTypeUncached(typ)
		ng.addResolvedImport(resolved)

		return resolved
	}

	// Build the full type name. This name should only be used for caching; it
	// does not differentiate properly built-in types and is thus unreliable.
	//
	// We also move the pointers from CType to the name to make the name unique
	// across pointers.
	fullName := movePtr(typ.CType, ng.fullGIR(typ.Name))

	v, ok := typeCache.Load(fullName)
	if ok {
		resolved := v.(*ResolvedType)
		ng.addResolvedImport(resolved)

		return resolved
	}

	// Cache miss. Use singleflight to ensure we're not looking up multiple
	// versions of the same type to prevent cache stampede.
	v, _, _ = typeFlight.Do(fullName, func() (interface{}, error) {
		resolved := ng.resolveTypeUncached(typ)

		// Save into the cache within the singleflight callback regardless if
		// the result is found or not.
		typeCache.Store(fullName, resolved)

		return resolved, nil
	})

	// may be a non-nil interface to a nil pointer.
	resolved := v.(*ResolvedType)
	ng.addResolvedImport(resolved)

	return resolved
}

func (ng *NamespaceGenerator) addResolvedImport(resolved *ResolvedType) {
	if resolved != nil && resolved.Import != "" && resolved.Import != ng.pkgPath {
		ng.addImportAlias(resolved.Import, resolved.Package)
	}
}

func (ng *NamespaceGenerator) resolveTypeUncached(typ gir.Type) *ResolvedType {
	if typ.Name == "" {
		ng.logln(logWarn, "empty gir type", typ)
		return nil
	}

	if prim, ok := girToBuiltin[typ.Name]; ok {
		return builtinType("", prim, typ)
	}

	// Resolve the unknown namespace that is GLib and primitive types.
	switch typ.Name {
	// TODO: ignore field
	// TODO: aaaaaaaaaaaaaaaaaaaaaaa
	case "gpointer":
		return builtinType("", "interface{}", typ)
	case "GLib.DestroyNotify", "DestroyNotify": // This should be handled externally.
		return builtinType("unsafe", "Pointer", typ)
	case "GType":
		return externGLibType("Type", typ, "GType")
	case "GObject.GValue", "GObject.Value": // inconsistency???
		return externGLibType("*Value", typ, "GValue")
	case "GObject.Object":
		return externGLibType("*Object", typ, "*GObject")
	case "GObject.Closure":
		return externGLibType("*Closure", typ, "*GClosure")
	case "GObject.Callback":
		// Callback is a special func(Any) Any type, so we treat it as
		// interface{} similarly to object.Connect(). We can use glib's Closure
		// APIs to parse this interface{}.
		return builtinType("", "interface{}", typ)

	case "GObject.InitiallyUnowned":
		return externGLibType("InitiallyUnowned", typ, "*GInitiallyUnowned")

	case "va_list":
		// CGo cannot handle variadic argument lists.
		return nil

	// We don't know what these types translates to.
	case "GObject.TypeModule":
		return nil
	case "GObject.ParamSpec": // this is deprecated
		return nil
	case "GObject.Parameter": // also deprecated I think
		return nil
	// TODO: Find a way to map EnumValue type.
	// TODO: Add _full function support.
	case "GObject.EnumValue":
		return nil
	}

	// CType is required here so we can properly account for pointers.
	if typ.CType == "" {
		ng.logln(logWarn, "type name", typ.Name, "missing CType")
		return nil
	}

	result := ng.gen.Repos.FindType(ng.current.Namespace.Name, typ.Name)
	if result == nil {
		ng.warnUnknownType(typ.Name)
		return nil
	}

	// Check for edge cases.
	switch {
	case result.Record != nil:
		if !ng.canRecord(*result.Record) {
			ng.logln(logInfo, "skipping", result.Record.Name, "from canRecord")
		}
	}

	return typeFromResult(ng.gen, typ, result)
}

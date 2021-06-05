package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
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

// cleanCType cleans the underlying C type of pointers and special keywords for
// CGo.
func cleanCType(cType string) string {
	cType = ctypePrefixEraser.Replace(cType)
	cType = strings.ReplaceAll(cType, "*", "")
	cType = strings.TrimSpace(cType)

	if replace, ok := cgoPrimitiveTypes[cType]; ok {
		cType = replace
	}

	return cType
}

// anyTypeIsVoid returns true if AnyType is a void type.
func anyTypeIsVoid(any gir.AnyType) bool {
	return any.Type != nil && any.Type.Name == "none"
}

// cgoTypeFromC converts a C type to a CGo type.
func cgoTypeFromC(ctype string) string {
	return movePtr(ctype, "C."+cleanCType(ctype))
}

// anyTypeCGo returns the CGo type for a GIR AnyType. An empty string is
// returned if none is made.
func anyTypeCGo(any gir.AnyType) string {
	return cgoTypeFromC(anyTypeC(any))
}

// anyTypeC returns the C type for a GIR AnyType. An empty string is returned if
// none is made.
func anyTypeC(any gir.AnyType) string {
	switch {
	case any.Array != nil:
		return ctypeFallback(any.Array.CType, any.Array.Name)
	case any.Type != nil:
		return ctypeFallback(any.Type.CType, any.Type.Name)
	default:
		return ""
	}
}

// girCTypes maps some primitive GIR types to C types, because people who write
// GIR generators are lazy fucks who can't be bothered to fill in the right
// information.
var girCTypes = map[string]string{
	"utf8":     "gchar*",
	"filename": "gchar*",
}

// ctypeFallback returns the C type OR the GIR type if it's empty.
func ctypeFallback(c, gir string) string {
	if c != "" {
		return ctypePrefixEraser.Replace(c)
	}

	// Handle edge cases with a hard-coded type map. Thanks, GIR, for evading
	// needed information.
	ctyp, ok := girCTypes[gir]
	if ok {
		return ctyp
	}

	return ctypePrefixEraser.Replace(gir)
}

// returnIsVoid returns true if the return type is void.
func returnIsVoid(ret *gir.ReturnValue) bool {
	return ret == nil || (ret != nil && anyTypeIsVoid(ret.AnyType))
}

// girToBuiltin maps the given GIR primitive type to a Go builtin type.
var girToBuiltin = map[string]string{
	"none":     "",
	"gboolean": "bool",
	"gfloat":   "float32",
	"gdouble":  "float64",
	"gint":     "int",
	"gssize":   "int",
	"gint8":    "int8",
	"gint16":   "int16",
	"gshort":   "int16",
	"gint32":   "int32",
	"glong":    "int32",
	"int32":    "int32",
	"gint64":   "int64",
	"guint":    "uint",
	"gsize":    "uint",
	"guchar":   "byte",
	"gchar":    "byte",
	"guint8":   "byte", // some weird cases
	"guint16":  "uint16",
	"gushort":  "uint16",
	"guint32":  "uint32",
	"gulong":   "uint32",
	"gunichar": "uint32",
	"guint64":  "uint64",
	"utf8":     "string",
	"filename": "string",

	// "long double": "float64",
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
	"unsigned int": "uint",

	// "long double":  "longdouble",
}

// TypeHasPointer returns true if the type being resolved has a pointer. This is
// useful for array passing from Go memory to C memory.
func TypeHasPointer(resolver TypeResolver, typ *ResolvedType) bool {
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
		return TypeHasPointer(resolver, ResolveTypeName(resolver, res.Alias.Name))

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

			if TypeHasPointer(resolver, resolver.ResolveType(*field.Type)) {
				return true
			}
		}

		return false
	}

	// Unknown type; assume there's a pointer.
	return true
}

// GoAnyType generates the Go type signature for the AnyType union. An empty
// string returned is an invalid type. If pub is true, then the returned string
// will use public interface types for classes instead of implementation types.
func GoAnyType(resolver TypeResolver, any gir.AnyType, pub bool) (string, bool) {
	switch {
	case any.Array != nil:
		return goArrayType(resolver, *any.Array, pub)
	case any.Type != nil:
		return GoType(resolver, *any.Type, pub)
	}

	// Probably varargs, ignore because Cgo.
	return "", false
}

// goArrayType generates the Go type signature for the given array.
func goArrayType(resolver TypeResolver, array gir.Array, pub bool) (string, bool) {
	arrayPrefix := "[]"
	if array.FixedSize > 0 {
		arrayPrefix = fmt.Sprintf("[%d]", array.FixedSize)
	}

	child, _ := GoAnyType(resolver, array.AnyType, pub)
	// There can't be []void, so this check ensures there can only be valid
	// array types.
	if child == "" {
		return "", false
	}

	return arrayPrefix + child, true
}

// GoType is a convenient function that wraps around ResolveType and returns the
// Go type.
func GoType(resolver TypeResolver, typ gir.Type, pub bool) (string, bool) {
	resolved := resolver.ResolveType(typ)
	if resolved == nil {
		return "", false
	}

	needsNamespace := resolved.NeedsNamespace(resolver.Namespace())

	if pub {
		return resolved.PublicType(needsNamespace), true
	}
	return resolved.ImplType(needsNamespace), true
}

// ResolveTypeName resolves the given GIR type name. The resolved type will
// always have no pointer.
func ResolveTypeName(resolver TypeResolver, girType string) *ResolvedType {
	var cType string

	// FindType is cached, so we can afford to do this.
	result := resolver.FindType(girType)
	if result != nil {
		// Use the CType result ONLY. The returned Name from Info does NOT have
		// the namespace prepended.
		_, cType = result.Info()
	}

	return resolver.ResolveType(gir.Type{
		Name:  girType,
		CType: cType,
	})
}

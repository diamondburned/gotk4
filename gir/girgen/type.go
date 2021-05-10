package girgen

import (
	"fmt"
	"log"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

type resolvedType struct {
	GoType string
	Type   gir.Type
	Import string // optional
}

// builtinType is a convenient function to make a new resolvedType.
func builtinType(goType string, typ gir.Type) *resolvedType {
	resolved := &resolvedType{GoType: goType, Type: typ}
	resolved.setPtr(nil)
	return resolved
}

// typeFromResult creates a resolved type from the given type result.
func typeFromResult(typ gir.Type, result *gir.TypeFindResult) *resolvedType {
	name, _ := result.Info()

	// same namespace, no package qual required.
	if result.SameNamespace {
		resolved := &resolvedType{name, typ, ""}
		resolved.setPtr(result)
		return resolved
	}

	// different namespace.
	pkg := gir.GoNamespace(result.Namespace)
	path := gir.ImportPath(pkg)

	resolved := &resolvedType{pkg + "." + name, typ, path}
	resolved.setPtr(result)
	return resolved
}

func (typ *resolvedType) setPtr(result *gir.TypeFindResult) {
	ptr := strings.Count(typ.Type.CType, "*")

	// Edge case: interfaces must not be pointers. We should still sometimes
	// allow for pointers to interfaces, if needed, but this likely won't work.
	if result != nil && result.Interface != nil && ptr > 0 {
		ptr--
	}
	// Edge case: a string is a gchar*, so we don't need a pointer.
	if typ.Type.Name == "utf8" && ptr > 0 {
		ptr--
	}

	typ.GoType = strings.Repeat("*", ptr) + typ.GoType
}

// CGoType returns the CGo type.
func (typ *resolvedType) CGoType() string {
	ptr := strings.Count(typ.Type.CType, "*")
	val := strings.TrimSuffix(typ.Type.CType, "*")

	return strings.Repeat("*", ptr) + "C." + val
}

// arrayType generates the Go type signature for the given array.
func (ng *NamespaceGenerator) resolveArrayType(array gir.Array) (string, bool) {
	arrayPrefix := "[]"
	if array.FixedSize > 0 {
		arrayPrefix = fmt.Sprintf("[%d]", array.FixedSize)
	}

	child, _ := ng.ResolveAnyType(array.AnyType)
	// There can't be []void, so this check ensures there can only be valid
	// array types.
	if child == "" {
		return "", false
	}

	return arrayPrefix + child, true
}

// ResolveAnyType generates the Go type signature for the AnyType union. An
// empty string returned is an invalid type.
func (ng *NamespaceGenerator) ResolveAnyType(any gir.AnyType) (string, bool) {
	switch {
	case any.Array != nil:
		return ng.resolveArrayType(*any.Array)

	case any.Type != nil:
		if r := ng.ResolveType(*any.Type); r != nil {
			return r.GoType, true
		}

	case any.VarArgs != nil:
		// CGo doesn't support variadic types.
		return "", false

	default:
		ng.debugln("anyType empty")
	}

	return "", false
}

// ResolveType resolves the given type from the GIR type field. It returns nil
// if the type is not known. It does not recursively traverse the type.
func (ng *NamespaceGenerator) ResolveType(typ gir.Type) *resolvedType {
	resolved := ng.resolveType(typ)
	if resolved != nil && resolved.Import != "" {
		ng.addImport(resolved.Import)
	}

	return resolved
}

// girPrimitiveGo converts the given GIR primitive type to a Go primitive type.
// False is returend if girType is not a primitive type.
func girPrimitiveGo(girType string) (string, bool) {
	var goType string

	switch girType {
	case "none":
		goType = ""
	case "gboolean":
		goType = "bool"
	case "gfloat":
		goType = "float32"
	case "gdouble":
		goType = "float64"
	case "long double": // pain
		goType = "float64"
	case "gint", "gssize":
		goType = "int"
	case "gint8":
		goType = "int8"
	case "gint16", "gshort":
		goType = "int16"
	case "gint32", "glong", "int32":
		goType = "int32"
	case "gint64":
		goType = "int64"
	case "guint", "gsize":
		goType = "uint"
	case "guchar", "gchar":
		goType = "byte"
	case "guint8":
		goType = "uint8"
	case "guint16", "gushort":
		goType = "uint16"
	case "guint32", "gulong", "gunichar": // pain pain pain pain
		goType = "uint32"
	case "guint64":
		goType = "uint64"
	case "utf8", "filename": // filename is probably UTF-16 hybrid ???
		goType = "string"
	case "gpointer":
		// TODO: ignore field
		// TODO: aaaaaaaaaaaaaaaaaaaaaaa
		goType = "unsafe.Pointer"
	default:
		return "", false
	}

	return goType, true
}

func (ng *NamespaceGenerator) resolveType(typ gir.Type) *resolvedType {
	if prim, ok := girPrimitiveGo(typ.Name); ok {
		return builtinType(prim, typ)
	}

	// Resolve the unknown namespace that is GLib and primitive types.
	switch typ.Name {
	case "GLib.DestroyNotify", "DestroyNotify": // This should be handled externally.
		return builtinType("unsafe.Pointer", typ)
	case "GType":
		return builtinType("glib.Type", typ)
	case "GObject.GValue", "GObject.Value": // inconsistency???
		return builtinType("*glib.Value", typ)
	case "GObject.Object":
		return builtinType("*glib.Object", typ)
	case "GObject.Closure":
		return builtinType("*glib.Closure", typ)
	case "GObject.GInitiallyUnowned":
		return builtinType("glib.InitiallyUnowned", typ)

	case "GObject.Callback":
		// Callback is a special func(Any) Any type, so we treat it as
		// interface{} similarly to object.Connect(). We can use glib's Closure
		// APIs to parse this interface{}.
		return builtinType("interface{}", typ)

	case "va_list":
		// CGo cannot handle variadic argument lists.
		return nil

	// We don't know what these types translates to.
	// TODO: Find a way to map EnumValue type.
	// TODO: Add _full function support.
	case "GObject.EnumValue":
		return nil
	}

	// Types that aren't in the switch tree that match any of these patterns are
	// types that must be in the switch tree, so them not being in there is a
	// bug.
	for _, check := range ng.gen.KnownTypes {
		if check(typ.Name) {
			log.Fatalf("missing gir type %s in the type tree\n", typ.Name)
		}
	}

	result := ng.gen.Repos.FindType(
		ng.current.Namespace.Name,
		ng.current.Namespace.Version,
		typ.Name,
	)
	if result == nil {
		ng.gen.warnUnknownType(typ.Name)
		return nil
	}

	return typeFromResult(typ, result)
}

// TODO: GoTypeConverter converts Go types to C with GIR type.

// AnyTypeConverter returns Go code that is the conversion from the given C
// value type to its respective Go value type. An empty string returned is
// invalid.
func (ng *NamespaceGenerator) AnyTypeConverter(value, target string, any gir.AnyType) string {
	switch {
	case any.Array != nil:
		return ng.arrayConverter(value, target, *any.Array)
	case any.Type != nil:
		return ng.typeConverter(value, target, *any.Type)
	}

	// Ignore VarArgs.
	return ""
}

func directCallOrCreate(value, target, typ string, create bool) string {
	var op = " = "
	if create {
		op = " := "
	}

	return target + op + typ + "(" + value + ")"
}

func (ng *NamespaceGenerator) arrayConverter(value, target string, array gir.Array) string {
	if array.Type != nil {
		ng.gen.debugln("skipping nested array", array)
		return ""
	}

	// Detect if the underlying is a compatible Go primitive type. If it is,
	// then we can directly cast a fixed-size array.
	if array.Type != nil {
		if primitiveGo, isPrimitive := girPrimitiveGo(array.Type.Name); isPrimitive {
			return fmt.Sprintf("%s = ([%d]%s)(%s)", target, array.FixedSize, primitiveGo, value)
		}
	}

	innerType := ng.ResolveType(*array.Type)
	if innerType == nil {
		return ""
	}
	innerCGoType := innerType.CGoType()

	// Generate a type converter from "src" to "dst" variables.
	innerConv := ng.typeConverter("src", "dst", *array.Type)
	if innerConv == "" {
		return ""
	}

	var b pen.Block

	switch {
	case array.FixedSize > 0:
		b.Linef("var a [%d]%s", array.FixedSize, innerType)
		// TODO: nested array support
		b.Linef("cArray := ([%d]%s)(%s)", array.FixedSize, array.Type.CType, value)
		b.EmptyLine()
		b.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		b.Linef("  src := cArray[i]")
		b.Linef("  dst := &a[i]")
		b.Linef("  " + innerConv)
		b.Linef("}")

	case array.Length != nil:
		return "" // TODO

	case array.Name == "GLib.Array": // treat as Go array
		b.Linef("var length uintptr")
		b.Linef("p := C.g_array_steal(&%s, (*C.gsize)(&length))", value)
		b.Linef("a := make([]%s, length)", innerType)
		b.Linef("for i := 0; i < length; i++ {")
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))", innerCGoType)
		b.Linef("  dst := &a[i]")
		// TODO: nested array support
		b.Linef("  " + innerConv)
		b.Linef("}")

	default: // null-terminated
		// Scan for the length.
		b.Linef("var length uint")
		b.Linef("for p := unsafe.Pointer(%s); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {", value)
		b.Linef("  length++")
		b.Linef("}")

		b.EmptyLine()

		// Preallocate the slice.
		b.Linef("a := make([]%s, length)", innerType)
		b.Linef("for i := 0; i < length; i++ {")
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(%s)) + 1))", innerCGoType, value)
		b.Linef("  dst := &a[i]")
		b.Linef("  " + innerConv)
		b.Linef("}")
	}

	return b.String()
}

func (ng *NamespaceGenerator) typeConverter(value, target string, typ gir.Type) string {
	return ng._typeConverter(value, target, typ, false)
}

func (ng *NamespaceGenerator) _typeConverter(value, target string, typ gir.Type, create bool) string {
	// Resolve primitive types.
	if prim, ok := girPrimitiveGo(typ.Name); ok {
		// Edge cases.
		switch prim {
		case "string":
			return directCallOrCreate(value, target, "C.GoString", create)
		default:
			return directCallOrCreate(value, target, prim, create)
		}
	}

	// Resolve special-case GLib types.
	switch typ.Name {
	case "GLib.DestroyNotify", "DestroyNotify":
		return ""
	case "GType":
		return ""
	case "GObject.GValue", "GObject.Value": // inconsistency???
		return ""
	case "GObject.Object":
		return directCallOrCreate(value, target, "glib.Take", create)
	case "GObject.Closure":
		return ""
	case "GObject.GInitiallyUnowned":
		return ""

	case "GObject.Callback":
		// TODO: When is this ever needed? How do I even do this?
		return ""
	case "va_list":
		// CGo cannot handle variadic argument lists.
		return ""
	case "GObject.EnumValue":
		// Refer to ResolveType.
		return ""
	}

	result := ng.gen.Repos.FindType(
		ng.current.Namespace.Name,
		ng.current.Namespace.Version,
		typ.Name,
	)
	if result == nil {
		// Probably already warned.
		return ""
	}

	resolved := typeFromResult(typ, result)

	// Resolve alias.
	if result.Alias != nil {
		b := strings.Builder{}
		b.WriteString("{\n")
		b.WriteString(ng._typeConverter(value, "tmp", result.Alias.Type, true))
		b.WriteString("\n")
		b.WriteString(directCallOrCreate("tmp", target, resolved.GoType, false))
		b.WriteString("}")
		return b.String()
	}

	// Resolve castable number types.
	if result.Enum != nil || result.Bitfield != nil {
		return directCallOrCreate(value, target, resolved.GoType, false)
	}

	// TODO: callbacks and functions are handled differently. Unsure if they're
	// doable.
	// TODO: handle unions.
	// TODO: interfaces should be wrapped by an unexported type.

	// Assume the wrap function. This so far works for classes and records.
	return directCallOrCreate(value, target, "wrap"+resolved.GoType, false)
}

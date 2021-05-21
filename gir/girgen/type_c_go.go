package girgen

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

// C to Go type conversions.

// TODO: GoTypeConverter converts Go types to C with GIR type.

// ArgAtFunc is the function to get the argument name at the given index. This
// function is primarily used for certain type conversions that need to access
// multiple variables.
type ArgAtFunc func(i int) string

// CGoConverter returns Go code that is the conversion from the given C value
// type to its respective Go value type. An empty string returned is invalid.
//
// The given argPrefix is used to get the nth parameter by concatenating the
// prefix with the index number. This is used for length parameters.
func (ng *NamespaceGenerator) CGoConverter(value, target string, any gir.AnyType, argAt ArgAtFunc) string {
	switch {
	case any.Array != nil:
		return ng.cgoArrayConverter(value, target, *any.Array, argAt)
	case any.Type != nil:
		return ng.cgoTypeConverter(value, target, *any.Type)
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

func (ng *NamespaceGenerator) cgoArrayConverter(value, target string, array gir.Array, argAt ArgAtFunc) string {
	if array.Type == nil {
		ng.gen.logln(logWarn, "skipping nested array", array)
		return ""
	}

	// Detect if the underlying is a compatible Go primitive type. If it is,
	// then we can directly cast a fixed-size array.
	if primitiveGo, ok := girPrimitiveGo[array.Type.Name]; ok {
		return fmt.Sprintf("%s = ([%d]%s)(%s)", target, array.FixedSize, primitiveGo, value)
	}

	innerResolved := ng.ResolveType(*array.Type)
	if innerResolved == nil {
		return ""
	}
	innerType := ng.PublicType(innerResolved)
	innerCGoType := innerResolved.CGoType()

	// Generate a type converter from "src" to "a[i]" variables.
	innerConv := ng.cgoTypeConverter("src", "a[i]", *array.Type)
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
		b.Linef("  " + innerConv)
		b.Linef("}")

	case array.Length != nil:
		lengthArg := argAt(*array.Length)
		b.Linef("a := make([]%s, %s)", innerType, lengthArg)
		b.Linef("for i := 0; i < %s; i++ {", lengthArg)
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))", innerCGoType)
		b.Linef("  " + innerConv)
		b.Linef("}")

	case array.Name == "GLib.Array": // treat as Go array
		b.Linef("var length uintptr")
		b.Linef("p := C.g_array_steal(&%s, (*C.gsize)(&length))", value)
		b.Linef("a := make([]%s, length)", innerType)
		b.Linef("for i := 0; i < length; i++ {")
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))", innerCGoType)
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
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(%s)) + i))", innerCGoType, value)
		b.Linef("  " + innerConv)
		b.Linef("}")
	}

	return b.String()
}

func (ng *NamespaceGenerator) cgoTypeConverter(value, target string, typ gir.Type) string {
	return ng._cgoTypeConverter(value, target, typ, false)
}

func (ng *NamespaceGenerator) _cgoTypeConverter(value, target string, typ gir.Type, create bool) string {
	// Resolve primitive types.
	if prim, ok := girPrimitiveGo[typ.Name]; ok {
		// Edge cases.
		switch prim {
		case "string":
			p := pen.NewPiece()
			p.Linef(directCallOrCreate(value, target, "C.GoString", create))
			p.Linef("defer C.free(unsafe.Pointer(%s))", value)
			return p.String()
		case "bool":
			ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
			return directCallOrCreate(value, target, "gextras.Gobool", create)
		default:
			return directCallOrCreate(value, target, prim, create)
		}
	}

	// Resolve special-case GLib types.
	switch typ.Name {
	case "gpointer":
		return directCallOrCreate(value, target, "unsafe.Pointer", create)
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
	case "GObject.InitiallyUnowned":
		return ""
	case "GObject.Callback":
		// TODO: When is this ever needed? How do I even do this?
		return ""
	case "va_list":
		// CGo cannot handle variadic argument lists.
		return ""
	case "GObject.EnumValue", "GObject.TypeModule", "GObject.ParamSpec", "GObject.Parameter":
		// Refer to ResolveType.
		return ""
	}

	result := ng.gen.Repos.FindType(ng.current.Namespace.Name, typ.Name)
	if result == nil {
		// Probably already warned.
		return ""
	}

	resolved := typeFromResult(ng.gen, typ, result)

	exportedName, _ := resolved.Extern.Result.Info()
	exportedName = PascalToGo(exportedName)

	// Resolve alias.
	if result.Alias != nil {
		var b pen.Block
		b.Line(ng._cgoTypeConverter(value, "tmp", result.Alias.Type, true))
		b.Line(directCallOrCreate("tmp", target, exportedName, false))
		return b.String()
	}

	// Resolve castable number types.
	if result.Enum != nil || result.Bitfield != nil {
		return directCallOrCreate(value, target, exportedName, false)
	}

	// TODO: callbacks and functions are handled differently. Unsure if they're
	// doable.
	// TODO: handle unions.
	// TODO: interfaces should be wrapped by an unexported type.

	// Assume the wrap function. This so far works for classes and records.
	// TODO: handle wrap functions from another package.
	return directCallOrCreate(value, target, "wrap"+exportedName, false)
}

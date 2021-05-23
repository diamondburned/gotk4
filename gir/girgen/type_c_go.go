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

// CGoConversion describes the information needed to generate Go code to convert
// C types to Go types.
type CGoConversion struct {
	Value  string
	Target string
	Type   gir.AnyType
	Owner  gir.TransferOwnership

	// ArgAt is used for array and closure generation.
	ArgAt ArgAtFunc
}

// copy creates a copy of CGoConversion with the given value and target. If the
// variables are empty, then the values aren't changed in the copy.
func (conv CGoConversion) copy(value, target string) CGoConversion {
	if value != "" {
		conv.Value = value
	}
	if target != "" {
		conv.Target = target
	}
	return conv
}

// CGoConverter returns Go code that is the conversion from the given C value
// type to its respective Go value type. An empty string returned is invalid.
//
// The given argPrefix is used to get the nth parameter by concatenating the
// prefix with the index number. This is used for length parameters.
func (ng *NamespaceGenerator) CGoConverter(conv CGoConversion) string {
	switch {
	case conv.Type.Array != nil:
		return ng.cgoArrayConverter(conv, *conv.Type.Array)
	case conv.Type.Type != nil:
		return ng.cgoTypeConverter(conv, *conv.Type.Type)
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

func (ng *NamespaceGenerator) cgoArrayConverter(conv CGoConversion, array gir.Array) string {
	if array.Type == nil {
		ng.gen.logln(logWarn, "skipping nested array", array)
		return ""
	}

	innerResolved := ng.ResolveType(*array.Type)
	if innerResolved == nil {
		return ""
	}
	innerType := ng.PublicType(innerResolved)
	innerCGoType := innerResolved.CGoType()

	// Generate a type converter from "src" to "${target}[i]" variables.
	innerConv := ng.cgoTypeConverter(conv.copy("src", conv.Target+"[i]"), *array.Type)
	if innerConv == "" {
		return ""
	}

	var b pen.Block

	switch {
	case array.FixedSize > 0:
		// Detect if the underlying is a compatible Go primitive type. If it is,
		// then we can directly cast a fixed-size array.
		if primitiveGo, ok := girPrimitiveGo[array.Type.Name]; ok {
			return fmt.Sprintf("%s = ([%d]%s)(%s)", conv.Target, array.FixedSize, primitiveGo, conv.Value)
		}

		// TODO: nested array support
		b.Linef("cArray := ([%d]%s)(%s)", array.FixedSize, array.Type.CType, conv.Value)
		b.EmptyLine()
		b.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		b.Linef("  src := cArray[i]")
		b.Linef("  " + innerConv)
		b.Linef("}")

	case array.Length != nil:
		lengthArg := conv.ArgAt(*array.Length)
		b.Linef("%s = make([]%s, %s)", conv.Target, innerType, lengthArg)
		b.Linef("for i := 0; i < uintptr(%s); i++ {", lengthArg)
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))", innerCGoType)
		b.Linef("  " + innerConv)
		b.Linef("}")

	case array.Name == "GLib.Array": // treat as Go array
		b.Linef("var length uintptr")
		b.Linef("p := C.g_array_steal(&%s, (*C.gsize)(&length))", conv.Value)
		b.Linef("%s = make([]%s, length)", conv.Target, innerType)
		b.Linef("for i := 0; i < length; i++ {")
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))", innerCGoType)
		// TODO: nested array support
		b.Linef("  " + innerConv)
		b.Linef("}")

	default: // null-terminated
		// Scan for the length.
		b.Linef("var length uint")
		b.Linef("for p := unsafe.Pointer(%s); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {", conv.Value)
		b.Linef("  length++")
		b.Linef("}")

		b.EmptyLine()

		// Preallocate the slice.
		b.Linef("%s = make([]%s, length)", conv.Target, innerType)
		b.Linef("for i := 0; i < length; i++ {")
		b.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(%s)) + i))", innerCGoType, conv.Value)
		b.Linef("  " + innerConv)
		b.Linef("}")
	}

	return b.String()
}

func (ng *NamespaceGenerator) cgoTypeConverter(conv CGoConversion, typ gir.Type) string {
	return ng._cgoTypeConverter(conv, typ, false)
}

func (ng *NamespaceGenerator) _cgoTypeConverter(conv CGoConversion, typ gir.Type, create bool) string {
	// Resolve primitive types.
	if prim, ok := girPrimitiveGo[typ.Name]; ok {
		// Edge cases.
		switch prim {
		case "string":
			p := pen.NewPiece()
			p.Linef(directCallOrCreate(conv.Value, conv.Target, "C.GoString", create))
			p.Linef("defer C.free(unsafe.Pointer(%s))", conv.Value)
			return p.String()
		case "bool":
			ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
			return directCallOrCreate(conv.Value, conv.Target, "gextras.Gobool", create)
		default:
			return directCallOrCreate(conv.Value, conv.Target, prim, create)
		}
	}

	// Resolve special-case GLib types.
	switch typ.Name {
	case "gpointer":
		return directCallOrCreate(conv.Value, conv.Target, "unsafe.Pointer", create)
	case "GLib.DestroyNotify", "DestroyNotify":
		return ""
	case "GType":
		return ""
	case "GObject.GValue", "GObject.Value": // inconsistency???
		return ""
	case "GObject.Object":
		return directCallOrCreate(conv.Value, conv.Target, "glib.Take", create)
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
		b := pen.NewBlock()
		b.Line(ng._cgoTypeConverter(conv.copy("", "tmp"), result.Alias.Type, true))
		b.Line(directCallOrCreate("tmp", conv.Target, exportedName, false))
		return b.String()
	}

	// Resolve castable number types.
	if result.Enum != nil || result.Bitfield != nil {
		return directCallOrCreate(conv.Value, conv.Target, exportedName, false)
	}

	if result.Class != nil {
		var gobjectFunction string
		switch conv.Owner.TransferOwnership {
		case "full", "container":
			// Full or container means we implicitly own the object, so we must
			// not take another reference.
			gobjectFunction = "AssumeOwnership"
		default:
			// Else the object is either unowned by us or it's a floating
			// reference. Take our own or sink the object.
			gobjectFunction = "Take"
		}

		return fmt.Sprintf(
			"%s = wrap%s(externglib.%s(unsafe.Pointer(%s)))",
			conv.Target, exportedName, gobjectFunction, conv.Value,
		)
	}

	// TODO: callbacks and functions are handled differently. Unsure if they're
	// doable.
	// TODO: handle unions.
	// TODO: interfaces should be wrapped by an unexported type.

	// Assume the wrap function. This so far works for classes and records.
	// TODO: handle wrap functions from another package.
	return directCallOrCreate(conv.Value, conv.Target, "wrap"+exportedName, false)
}

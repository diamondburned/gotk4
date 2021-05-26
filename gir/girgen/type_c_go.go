package girgen

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

// C to Go type conversions.

// TODO: GoTypeConverter converts Go types to C with GIR type.

// CGoConverter returns Go code that is the conversion from the given C value
// type to its respective Go value type. An empty string returned is invalid.
//
// The given argPrefix is used to get the nth parameter by concatenating the
// prefix with the index number. This is used for length parameters.
func (ng *NamespaceGenerator) CGoConverter(conv TypeConversion) string {
	switch {
	case conv.Type.Array != nil:
		return ng.cgoArrayConverter(conv)
	case conv.Type.Type != nil:
		return ng.cgoTypeConverter(conv)
	}

	// Ignore VarArgs.
	return ""
}

func (ng *NamespaceGenerator) cgoArrayConverter(conv TypeConversion) string {
	array := conv.Type.Array

	if array.Type == nil {
		ng.gen.logln(logWarn, "skipping nested array", array.CType)
		return ""
	}

	innerResolved := ng.ResolveType(*array.Type)
	if innerResolved == nil {
		return ""
	}
	innerType := ng.PublicType(innerResolved)
	innerCGoType := innerResolved.CGoType()

	// Generate a type converter from "src" to "${target}[i]" variables.
	innerTypeConv := conv
	innerTypeConv.Value = "src"
	innerTypeConv.Target = conv.Target + "[i]"
	innerTypeConv.Type = array.AnyType

	innerConv := ng.CGoConverter(innerTypeConv)
	if innerConv == "" {
		return ""
	}

	var b pen.Block

	switch {
	case array.FixedSize > 0:
		// Detect if the underlying is a compatible Go primitive type that isn't
		// a string. If it is, then we can directly cast a fixed-size array.
		if primitiveGo, ok := girPrimitiveGo[array.Type.Name]; ok && primitiveGo != "string" {
			return conv.callf("[%d]%s", array.FixedSize, primitiveGo)
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

func (ng *NamespaceGenerator) cgoTypeConverter(conv TypeConversion) string {
	typ := conv.Type.Type

	// Resolve primitive types.
	if prim, ok := girPrimitiveGo[typ.Name]; ok {
		// Edge cases.
		switch prim {
		case "string":
			p := pen.NewPiece()
			p.Linef("%s = C.GoString(%s)", conv.Value, conv.Target)
			p.Linef("defer C.free(unsafe.Pointer(%s))", conv.Value)
			return p.String()
		case "bool":
			ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
			return directCall(conv.Value, conv.Target, "gextras.Gobool")
		default:
			return directCall(conv.Value, conv.Target, prim)
		}
	}

	// Resolve special-case GLib types.
	switch typ.Name {
	case "gpointer":
		ng.addImport("github.com/diamondburned/gotk4/internal/box")

		if conv.BoxCast == "" {
			return fmt.Sprintf(
				"%s = box.Get(uintptr(%s))",
				conv.Target, conv.Value,
			)
		}

		return fmt.Sprintf(
			"%s = box.Get(uintptr(%s)).(%s)",
			conv.Target, conv.Value, conv.BoxCast,
		)

	case "GObject.Object", "GObject.InitiallyUnowned":
		return cgoTakeObject(conv, "")

	case "GLib.DestroyNotify", "DestroyNotify":
		// There's no Go equivalent for C's DestroyNotify; the user should never
		// see this.
		return ""

	case "GType":
		return "" // TODO
	case "GObject.GValue", "GObject.Value":
		return "" // TODO
	case "GObject.Closure":
		return "" // TODO
	case "GObject.Callback":
		return "" // TODO
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

	resolved := typeFromResult(ng.gen, *typ, result)

	// goName contains the pointer.
	goName := ng.PublicType(resolved)

	// Resolve alias.
	if result.Alias != nil {
		rootType := ng.ResolveType(result.Alias.Type)
		if rootType == nil {
			ng.logln(logWarn, "alias", result.Alias.Name, "lacks conv for", result.Alias.Type.Name)
			return ""
		}

		rootConv := conv
		rootConv.Type = gir.AnyType{Type: &result.Alias.Type}
		rootConv.Target = "tmp"

		b := pen.NewBlock()
		b.Linef("var tmp %s", ng.PublicType(rootType))
		b.Linef(ng.CGoConverter(rootConv))
		b.Linef("%s = %s(tmp)", conv.Target, goName)
		return b.String()
	}

	// Resolve castable number types.
	if result.Enum != nil || result.Bitfield != nil {
		return conv.call(goName)
	}

	if result.Class != nil || result.Record != nil {
		// externName doesn't contain the pointer.
		externName, _ := resolved.Extern.Result.Info()
		externName = PascalToGo(externName)

		wrapName := "Wrap" + externName
		if resolved.NeedsNamespace(ng.current) {
			wrapName = resolved.Package + "." + wrapName
		}

		switch {
		case result.Class != nil:
			return cgoTakeObject(conv, wrapName)
		case result.Record != nil:
			return conv.call(wrapName)
		}
	}

	if result.Callback != nil {
		ng.logln(logError, "idk what to do with C->Go callback", goName)
		return ""
	}

	// TODO: function
	// TODO: union
	// TODO: callback
	// TODO: interface

	// TODO: callbacks and functions are handled differently. Unsure if they're
	// doable.
	// TODO: handle unions.
	// TODO: interfaces should be wrapped by an unexported type.

	// Assume the wrap function. This so far works for classes and records.
	// TODO: handle wrap functions from another package.
	return ""
}

// cgoTakeObject generates a glib.Take or glib.AssumeOwnership.
func cgoTakeObject(conv TypeConversion, wrap string) string {
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

	if wrap == "" {
		return fmt.Sprintf(
			"%s = externglib.%s(unsafe.Pointer(%s.Native()))",
			conv.Target, gobjectFunction, conv.Value,
		)
	}

	return fmt.Sprintf(
		"%s = %s(externglib.%s(unsafe.Pointer(%s.Native())))",
		conv.Target, wrap, gobjectFunction, conv.Value,
	)
}

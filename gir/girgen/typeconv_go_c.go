package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/internal/pen"
)

// Go to C type conversions.

// TODO: is there a reason GoCConverter just doesn't take in ParameterAttr?

func (ng *NamespaceGenerator) GoCConverter(conv TypeConversionToC) string {
	switch {
	case conv.Type.Type != nil:
		return ng.gocTypeConverter(conv)
	case conv.Type.Array != nil:
		return ng.gocArrayConverter(conv)
	default:
		return ""
	}

}

func (ng *NamespaceGenerator) gocArrayConverter(conv TypeConversionToC) string {
	array := conv.Type.Array

	if array.Type == nil {
		ng.logln(logWarn, "skipping nested array", array.CType)
	}

	innerResolved := ng.ResolveType(*array.Type)
	if innerResolved == nil {
		return ""
	}

	outerCGoType := anyTypeCGo(conv.Type)
	innerCGoType := innerResolved.CGoType()

	// Generate a type converter from "src" to "dst[i]" variables.
	innerTypeConv := conv
	innerTypeConv.TypeConversion = conv.inner("src", "dst[i]")

	innerConv := ng.GoCConverter(innerTypeConv)
	if innerConv == "" {
		return ""
	}

	var b pen.Block

	switch {
	case array.FixedSize > 0:
		if innerResolved.IsPrimitive() {
			// We can directly use Go's array as a pointer, as long as we defer
			// properly.
			b.Linef("%s = (%s)(&%s)", conv.Target, outerCGoType, conv.Value)

			ng.addImport("runtime")
			b.Linef("defer runtime.KeepAlive(&%s)", conv.Value)
			return b.String()
		}

		// Target fixed array, so we can directly set the data over. The memory
		// is ours, and allocation is handled by Go.
		b.Linef("dst := &%s", conv.Target)
		b.EmptyLine()
		b.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		b.Linef("  src := %s[i]", conv.Value)
		b.Linef("  " + innerConv)
		b.Linef("}")

	case array.Length != nil:
		length := fmt.Sprintf("len(%s)", conv.Value)

		// Use the backing array with the appropriate transfer-ownership rule
		// for primitive types; see type_c_go.go.
		if !conv.isTransferring() && innerResolved.IsPrimitive() {
			b.Linef("%s = (%s)(&%s[0])", conv.Target, outerCGoType, conv.Value)
			b.Linef("%s = %s", conv.ArgAt(*array.Length), length)

			ng.addImport("runtime")
			b.Linef("defer runtime.KeepAlive(%s)", conv.Value)
			return b.String()
		}

		// Copying is pretty much required here, since the C code will store the
		// pointer, so we can't reliably do this with Go's memory.

		ptr := fmt.Sprintf("C.malloc(%s * len(%s))", csizeof(ng, innerResolved), conv.Value)

		b.Linef("var dst []%s", innerCGoType)
		b.Linef(goSliceFromPtr("dst", ptr, length))

		// C.malloc will allocate on the C side, so we'll have to free it.
		if !conv.isTransferring() {
			b.Linef("defer C.free(unsafe.Pointer(sliceHeader.Data))")
		}

		b.EmptyLine()
		b.Linef("for i := 0; i < %s; i++ {", length)
		b.Linef("  src := %s[i]", conv.Value)
		b.Linef("  " + innerConv)
		b.Linef("}")
		b.EmptyLine()
		b.Linef("%s = (%s)(unsafe.Pointer(sliceHeader.Data))", conv.Target, outerCGoType)
		b.Linef("%s = %s", conv.ArgAt(*array.Length), length)

	case array.Name == "GLib.Array":
		length := fmt.Sprintf("len(%s)", conv.Value)

		b.Linef(
			"%s = C.g_array_sized_new(%v, false, C.guint(%s), %s)",
			conv.Target, array.ZeroTerminated, csizeof(ng, innerResolved), length)

		b.EmptyLine()
		b.Linef("var dst []%s", innerCGoType)
		b.Linef(goSliceFromPtr("dst", conv.Target+".data", length))
		b.EmptyLine()

		b.Linef("for i := 0; i < %s; i++ {", length)
		b.Linef("  src := %s[i]", conv.Value)
		b.Linef("  " + innerConv)
		b.Linef("}")
	}

	return b.String()
}

func csizeof(ng *NamespaceGenerator, resolved *ResolvedType) string {
	if !strings.Contains(resolved.CType, "*") {
		return "C.sizeof_" + resolved.CType
	}

	ng.addImport("unsafe")
	return "unsafe.Sizeof((*struct{})(nil))"
}

func (ng *NamespaceGenerator) gocTypeConverter(conv TypeConversionToC) string {
	typ := conv.Type.Type

	if prim, ok := girToBuiltin[typ.Name]; ok {
		switch prim {
		case "string":
			p := pen.NewPiece()
			p.Linef("%s = (*C.gchar)(C.CString(%s))", conv.Target, conv.Value)

			// If we're not giving ownership this C-allocated string, then we
			// can free it once done.
			if !conv.isTransferring() {
				p.Linef("defer C.free(unsafe.Pointer(%s))", conv.Target)
			}

			return p.String()

		case "bool":
			ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
			return conv.call("gextras.Cbool")

		default:
			return conv.call(anyTypeCGo(conv.Type))
		}
	}

	switch typ.Name {
	case "gpointer":
		ng.addImport("github.com/diamondburned/gotk4/internal/box")
		return fmt.Sprintf("%s = C.gpointer(box.Assign(%s))", conv.Target, conv.Value)

	case "GLib.DestroyNotify", "DestroyNotify":
		// Use a constant C function. Value is unused.
		ng.needsCallbackDelete()
		return fmt.Sprintf("%s = (*[0]byte)(C.callbackDelete)", conv.Target)

	case "GType":
		// Just a primitive.
		return fmt.Sprintf("%s = C.GType(%s)", conv.Target, conv.Value)

	case "GObject.GValue", "GObject.Value":
		// https://pkg.go.dev/github.com/gotk3/gotk3/glib#Type
		return fmt.Sprintf("%s = (*C.GValue)(%s.GValue)", conv.Target, conv.Value)

	case "GObject.Object":
		// Use .Native() here instead of directly accessing the native pointer,
		// since Value might be an Objector interface.
		return fmt.Sprintf("%s = (*C.GObject)(%s.Native())", conv.Target, conv.Value)
	case "GObject.InitiallyUnowned":
		return fmt.Sprintf("%s = (*C.GInitiallyUnowned)(%s.Native())", conv.Target, conv.Value)

	// These are empty until they're filled out in type_c_go.go
	case "GObject.Closure":
		return ""
	case "GObject.Callback":
		return ""
	case "va_list":
		return ""
	case "GObject.EnumValue", "GObject.TypeModule", "GObject.ParamSpec", "GObject.Parameter":
		// Refer to ResolveType.
		return ""
	}

	result := ng.gen.Repos.FindType(ng.current.Namespace.Name, typ.Name)
	if result == nil {
		return ""
	}

	resolved := typeFromResult(ng.gen, *typ, result)

	exportedName, _ := resolved.Extern.Result.Info()
	exportedName = PascalToGo(exportedName)

	if result.Enum != nil || result.Bitfield != nil {
		// Direct cast-able.
		return fmt.Sprintf("%s = (%s)(%s)", conv.Target, resolved.CGoType(), conv.Value)
	}

	if result.Class != nil || result.Record != nil || result.Interface != nil {
		// gextras.Objector has Native() uintptr.
		return fmt.Sprintf(
			"%s = (%s)(%s.Native())",
			conv.Target, resolved.CGoType(), conv.Value,
		)
	}

	if result.Callback != nil {
		// Callbacks must have the closure attribute to store the closure
		// pointer.
		if conv.Closure == nil {
			ng.logln(logInfo, "skipping callback", exportedName, "since missing closure")
			return ""
		}

		ng.addImport("github.com/diamondburned/gotk4/internal/box")

		p := pen.NewPiece()

		// Return the constant function here. The function will dynamically load
		// the user_data, which will match with the "gpointer" case above.
		//
		// As for the pointer to byte array cast, see
		// https://github.com/golang/go/issues/19835.
		p.Linef("%s = (*[0]byte)(C.%s%s)", conv.Target, callbackPrefix, exportedName)
		p.Linef("%s = C.gpointer(box.Assign(%s))", conv.ArgAt(*conv.Closure), conv.Value)

		if conv.Destroy != nil {
			p.Linef("%s = (*[0]byte)(C.callbackDelete)", conv.ArgAt(*conv.Destroy))
		}

		return p.String()
	}

	// TODO
	return ""
}

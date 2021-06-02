package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/internal/pen"
)

// Go to C type conversions.

// TODO: is there a reason GoCConverter just doesn't take in ParameterAttr?

// GoCConverter generates code that converts from Go values to C.
func (fg *FileGenerator) GoCConverter(conv TypeConversionToC) string {
	switch {
	case conv.Type.Type != nil:
		return fg.gocTypeConverter(conv)
	case conv.Type.Array != nil:
		return fg.gocArrayConverter(conv)
	default:
		return ""
	}
}

func (fg *FileGenerator) gocArrayConverter(conv TypeConversionToC) string {
	array := conv.Type.Array

	if array.Type == nil {
		fg.Logln(LogSkip, "nested array", array.CType)
		return ""
	}

	innerResolved := fg.ResolveType(*array.Type)
	if innerResolved == nil {
		return ""
	}

	outerCGoType := anyTypeCGo(conv.Type)
	innerCGoType := innerResolved.CGoType()

	// Generate a type converter from "src" to "dst[i]" variables.
	innerTypeConv := conv
	innerTypeConv.TypeConversion = conv.inner("src", "dst[i]")

	innerConv := fg.GoCConverter(innerTypeConv)
	if innerConv == "" {
		return ""
	}

	switch {
	case array.FixedSize > 0:
		if !TypeHasPointer(fg, innerResolved) {
			fg.addImport("runtime")

			// We can directly use Go's array as a pointer, as long as we defer
			// properly.
			return pen.NewPiece().
				Linef("%s = (%s)(&%s)", conv.Target, outerCGoType, conv.Value).
				Linef("defer runtime.KeepAlive(&%s)", conv.Value).
				String()
		}

		// Target fixed array, so we can directly set the data over. The memory
		// is ours, and allocation is handled by Go.
		b := pen.NewBlock()
		b.Linef("dst := &%s", conv.Target)
		b.EmptyLine()
		b.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		b.Linef("  src := %s[i]", conv.Value)
		b.Linef("  " + innerConv)
		b.Linef("}")

		return b.String()

	case array.Length != nil:
		length := fmt.Sprintf("len(%s)", conv.Value)

		// Use the backing array with the appropriate transfer-ownership rule
		// for primitive types; see type_c_go.go.
		if !conv.isTransferring() && !TypeHasPointer(fg, innerResolved) {
			fg.addImport("runtime")
			fg.addImport("unsafe")

			return pen.NewPiece().
				Linef("%s = (%s)(unsafe.Pointer(&%s[0]))", conv.Target, outerCGoType, conv.Value).
				Linef("%s = %s", conv.ArgAt(*array.Length), length).
				Linef("defer runtime.KeepAlive(%s)", conv.Value).
				String()
		}

		fg.addImport("reflect")
		fg.addImport("unsafe")

		// Copying is pretty much required here, since the C code will store the
		// pointer, so we can't reliably do this with Go's memory.

		b := pen.NewBlock()
		b.Linef("var dst []%s", innerCGoType)
		b.Linef("ptr := C.malloc(%s * len(%s))", csizeof(fg, innerResolved), conv.Value)
		b.Linef(goSliceFromPtr("dst", "ptr", length))

		// C.malloc will allocate on the C side, so we'll have to free it.
		if !conv.isTransferring() {
			b.Line("defer C.free(unsafe.Pointer(ptr))")
		}

		b.EmptyLine()
		b.Linef("for i := 0; i < %s; i++ {", length)
		b.Linef("  src := %s[i]", conv.Value)
		b.Linef("  " + innerConv)
		b.Linef("}")
		b.EmptyLine()
		b.Linef("%s = (%s)(unsafe.Pointer(ptr))", conv.Target, outerCGoType)
		b.Linef("%s = %s", conv.ArgAt(*array.Length), length)

		return b.String()

	case array.Name == "GLib.Array":
		length := fmt.Sprintf("len(%s)", conv.Value)

		fg.addImport("reflect")
		fg.addImport("unsafe")

		b := pen.NewBlock()
		b.Linef(
			"%s = C.g_array_sized_new(%v, false, C.guint(%s), %s)",
			conv.Target, array.ZeroTerminated, csizeof(fg, innerResolved), length)
		b.EmptyLine()
		b.Linef("var dst []%s", innerCGoType)
		b.Linef(goSliceFromPtr("dst", conv.Target+".data", length))
		b.EmptyLine()
		b.Linef("for i := 0; i < %s; i++ {", length)
		b.Linef("  src := %s[i]", conv.Value)
		b.Linef("  " + innerConv)
		b.Linef("}")

		return b.String()

	case array.IsZeroTerminated():
		length := fmt.Sprintf("len(%s)", conv.Value)

		fg.addImport("reflect")
		fg.addImport("unsafe")

		b := pen.NewBlock()
		b.Linef("var dst []%s", innerCGoType)
		b.Linef("ptr := C.malloc(%s * (len(%s)+1))", csizeof(fg, innerResolved), conv.Value)
		b.Linef(goSliceFromPtr("dst", "ptr", length))

		// See above, in the array.Length != nil case.
		if !conv.isTransferring() {
			b.Line("defer C.free(unsafe.Pointer(ptr))")
		}

		b.EmptyLine()
		b.Linef("for i := 0; i < %s; i++ {", length)
		b.Linef("  src := %s[i]", conv.Value)
		b.Linef("  " + innerConv)
		b.Linef("}")
		b.EmptyLine()
		b.Linef("%s = (%s)(unsafe.Pointer(ptr))", conv.Target, outerCGoType)

		return b.String()
	}

	fg.Logln(LogSkip, conv.ParentName+":", "weird array type to C")
	return ""
}

func csizeof(fg *FileGenerator, resolved *ResolvedType) string {
	if !strings.Contains(resolved.CType, "*") {
		return "C.sizeof_" + resolved.CType
	}

	fg.addImport("unsafe")
	return "unsafe.Sizeof((*struct{})(nil))"
}

func (fg *FileGenerator) gocTypeConverter(conv TypeConversionToC) string {
	typ := conv.Type.Type

	for _, unsupported := range unsupportedCTypes {
		if unsupported == typ.CType {
			return ""
		}
	}

	if prim, ok := girToBuiltin[typ.Name]; ok {
		switch prim {
		case "string":
			p := pen.NewPiece()
			p.Linef("%s = (*C.gchar)(C.CString(%s))", conv.Target, conv.Value)

			// If we're not giving ownership this C-allocated string, then we
			// can free it once done.
			if !conv.isTransferring() {
				fg.addImport("unsafe")
				p.Linef("defer C.free(unsafe.Pointer(%s))", conv.Target)
			}

			return p.String()

		case "bool":
			return pen.NewPiece().
				Linef("if %s {", conv.Value).
				Linef("  %s = C.TRUE", conv.Target).
				Linef("}").
				String()

		default:
			return conv.call(anyTypeCGo(conv.Type))
		}
	}

	switch ensureNamespace(fg.Namespace(), typ.Name) {
	case "gpointer":
		fg.addImport("github.com/diamondburned/gotk4/internal/box")
		return fmt.Sprintf("%s = C.gpointer(box.Assign(%s))", conv.Target, conv.Value)

	case "GObject.Type", "GType":
		// Just a primitive.
		return fmt.Sprintf("%s = C.GType(%s)", conv.Target, conv.Value)

	case "GObject.Value":
		// https://pkg.go.dev/github.com/gotk3/gotk3/glib#Type
		return fmt.Sprintf("%s = (*C.GValue)(%s.GValue)", conv.Target, conv.Value)

	case "GObject.Object":
		// Use .Native() here instead of directly accessing the native pointer,
		// since Value might be an Objector interface.
		return fmt.Sprintf("%s = (*C.GObject)(%s.Native())", conv.Target, conv.Value)

	case "GObject.InitiallyUnowned":
		return fmt.Sprintf("%s = (*C.GInitiallyUnowned)(%s.Native())", conv.Target, conv.Value)
	}

	// Pretend that ignored types don't exist.
	if fg.parent.mustIgnore(typ.Name, typ.CType) {
		return ""
	}

	result := fg.FindType(typ.Name)
	if result == nil {
		return ""
	}

	resolved := typeFromResult(fg.parent.gen, *typ, result)

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
			fg.Logln(LogSkip, "Go->C callback", exportedName, "since missing closure")
			return ""
		}

		fg.addImport("github.com/diamondburned/gotk4/internal/box")

		p := pen.NewPiece()

		// Return the constant function here. The function will dynamically load
		// the user_data, which will match with the "gpointer" case above.
		//
		// As for the pointer to byte array cast, see
		// https://github.com/golang/go/issues/19835.
		p.Linef("%s = (*[0]byte)(C.%s%s)", conv.Target, callbackPrefix, exportedName)
		p.Linef("%s = C.gpointer(box.Assign(%s))", conv.ArgAt(*conv.Closure), conv.Value)

		if conv.Destroy != nil {
			fg.needsCallbackDelete()
			p.Linef("%s = (*[0]byte)(C.callbackDelete)", conv.ArgAt(*conv.Destroy))
		}

		return p.String()
	}

	// TODO
	return ""
}

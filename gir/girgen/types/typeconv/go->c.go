package typeconv

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// Go to C type conversions.

func (conv *Converter) gocConvert(value *ValueConverted) bool {
	switch {
	case value.AnyType.Type != nil:
		return conv.gocConverter(value)
	case value.AnyType.Array != nil:
		return conv.gocArrayConverter(value)
	default:
		return false
	}
}

func (conv *Converter) gocArrayConverter(value *ValueConverted) bool {
	if value.AnyType.Array.Type == nil {
		conv.Logln(logger.Debug, "Go->C skipping nested array", value.AnyType.Array.CType)
		return value.Optional
	}

	array := *value.AnyType.Array

	// Ensure that the array type matches the inner type. Some functions violate
	// this, e.g. g_spawn_command_line_sync().
	if array.Type.CType == "" {
		// Copy the inner type so we don't accidentally change a reference.
		typ := *array.Type
		// Dereference the inner type by 1.
		typ.CType = strings.Replace(array.CType, "*", "", 1)

		array.AnyType.Type = &typ
		value.AnyType.Array = &array
	}

	if array.CType == "void*" {
		// CGo treats void* arrays a bit weirdly: the function's input parameter
		// type is actually unsafe.Pointer, so we have to wrap it.
		value.OutCall = fmt.Sprintf("unsafe.Pointer(%s)", value.OutCall)
	}

	// This is always the same.
	value.OutType = types.AnyTypeCGo(value.AnyType)
	value.outDecl.Linef("var %s %s", value.OutName, value.OutType)

	// Length is roughly always the same as well.
	var length *ValueConverted
	if array.Length != nil {
		length = conv.convertParam(*array.Length)
		// Ensure length is present.
		if length == nil {
			return false
		}

		// Length has no input, as it's from the slice.
		value.outDecl.Linef("var %s %s", length.OutName, length.OutType)
		value.p.Linef("%s = %s(len(%s))", length.OutName, length.OutType, value.InName)
	}

	// These cases have invalid inner type names that aren't useful to us, so we
	// handle them on our own.
	switch {
	case types.CleanCType(array.CType, false) == "gchar*":
		// This is technically a []byte, and we should use a []byte, because the
		// C code may mutate the backing array. The internal type is "utf8",
		// which is false, because the C type is just a single character.
		value.InType = "[]byte"
		value.inDecl.Linef("var %s []byte", value.InName)

		// Only hand over Go memory if we're not copying.
		if !value.isTransferring() {
			value.header.Import("unsafe")
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s[0]))",
				value.OutName, value.OutType, value.InName,
			)

			return true
		}

		// Use CString, which copies. Since we're transferring ownership to C,
		// don't free it after we're done. CString is actually null-terminated,
		// but we're giving the caller the length, so an extra byte is fine.
		value.p.Linef(
			"%s = (%s)(C.CString(%s))",
			value.OutName, value.OutType, value.InName,
		)

		return true
	}

	inner := conv.convertInner(value, value.InName+"[i]", "out[i]")
	if inner == nil {
		return false
	}

	if array.FixedSize > 0 {
		value.InType = fmt.Sprintf("[%d]%s", array.FixedSize, inner.InType)
		value.inDecl.Linef("var %s %s", value.InName, value.InType)
	} else {
		value.InType = fmt.Sprintf("[]%s", inner.InType)
		value.inDecl.Linef("var %s %s", value.InName, value.InType)
	}

	switch {
	case array.FixedSize > 0:
		// Safe to do if this is a primitive AND we're not setting this inside a
		// calllback, since the callback will retain Go memory beyond its
		// lifetime which is bad.
		if !value.isTransferring() && inner.resolved.CanCast() {
			// We can directly use Go's array as a pointer, as long as we defer
			// properly.
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s))",
				value.OutName, value.OutType, value.InName)

			return true
		}

		// Target fixed array, so we can directly set the data over. The memory
		// is ours, and allocation is handled by Go.
		value.p.Descend()

		value.p.Linef("out := (*[%d]%s)(unsafe.Pointer(%s))",
			array.FixedSize, inner.OutType, value.OutName)

		value.p.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		value.p.Linef("  " + inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	case array.Length != nil:
		// Use the backing array with the appropriate transfer-ownership rule
		// for primitive types; see type_c_go.go.
		if !value.isTransferring() && inner.resolved.CanCast() {
			value.header.Import("unsafe")
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s[0]))",
				value.OutName, value.OutType, value.InName)

			return true
		}

		value.header.Import("unsafe")

		value.p.Linef(
			"%s = (%s)(%s)",
			value.OutName, value.OutType, inner.cmalloc(value.InName, false),
		)
		if !value.isTransferring() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)
		}

		value.p.Descend()

		value.p.Linef("out := unsafe.Slice(%s, len(%s))", value.OutName, value.InName)
		value.p.Linef("for i := range %s {", value.InName)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	case array.Name == "GLib.Array":
		value.header.Import("unsafe")

		// https://developer.gnome.org/glib/stable/glib-Arrays.html#g-array-sized-new
		value.p.Linef(
			"%s = C.g_array_sized_new(%t, false, C.guint(%s), C.guint(len(%s)))",
			value.OutName, array.IsZeroTerminated(), inner.csizeof(), value.InName)
		value.p.Linef(
			"%s = C.g_array_set_size(%s, C.guint(len(%s)))",
			value.OutName, value.OutName, value.InName)

		if !value.isTransferring() {
			value.p.Linef("defer C.g_array_unref(%s)", value.OutName)
		}

		value.p.Descend()

		value.p.Linef("out := unsafe.Slice(%s.data, len(%s))", value.OutName, value.InName)
		value.p.Linef("for i := range %s {", value.InName)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	case array.Name == "GLib.ByteArray":
		// We know that the type will always be []byte if this case hits, so we
		// don't need to do complicated conversions.
		value.header.Import("unsafe")

		if !value.isTransferring() {
			// No-copy path that works as long as we don't use unref, because
			// that will free the Go memory.
			value.p.Linef(
				"%s = C.g_byte_array_new_take((*C.guint8)(&%s[0]), C.size(len(%s)))",
				value.OutName, value.InName, value.InName)
			// Steal will free the GByteArray, but not the underlying array
			// itself, which is the Go memory.
			value.p.Linef("defer C.g_byte_array_steal(%s, nil)", value.InName)
			return true
		}

		// https://developer.gnome.org/glib/stable/glib-Byte-Arrays.html#g-byte-array-new
		value.p.Linef(
			"%s = C.g_byte_array_sized_new(C.guint(len(%s)))",
			value.OutName, value.InName)
		value.p.Linef(
			"%s = C.g_byte_array_append(%s, (*C.guint8)(&%s[0]), C.guint(len(%s)))",
			value.OutName, value.OutName, value.InName, value.InName)
		// unref will free the underlying array as well.
		value.p.Linef("defer C.g_byte_array_unref(%s)", value.OutName)
		return true

	case array.IsZeroTerminated():
		value.header.Import("unsafe")

		// See if we can possibly reuse the Go slice in a shorter way.
		if !value.isTransferring() && inner.resolved.CanCast() {
			value.p.Descend()
			value.p.Linef("var zero %s", inner.InType)
			value.p.Linef("%s = append(%[1]s, zero)", value.InName)
			value.p.Ascend()

			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s[0]))",
				value.OutName, value.OutType, value.InName)

			return true
		}

		value.p.Linef(
			"%s = (%s)(%s)",
			value.OutName, value.OutType, inner.cmalloc(value.InName, true))
		if !value.isTransferring() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)
		}

		value.p.Descend()

		value.p.Linef("out := unsafe.Slice(%s, len(%s))", value.OutName, value.InName)
		value.p.Linef("for i := range %s {", value.InName)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	default:
		conv.Logln(logger.Debug, "Go->C weird array type", array.Type)
	}

	return false
}

func (conv *Converter) gocConverter(value *ValueConverted) bool {
	for _, unsupported := range types.UnsupportedCTypes {
		if unsupported == value.AnyType.Type.CType {
			return false
		}
	}

	if !value.resolveType(conv) {
		return false
	}

	switch {
	case value.resolved.IsBuiltin("string"):
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 0)
		}

		value.p.Linef("%s = (%s)(C.CString(%s))", value.OutName, value.OutType, value.InName)
		// If we're not giving ownership this mallocated string, then we
		// can free it once done.
		if !value.isTransferring() {
			value.header.Import("unsafe")
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)
		}

		return true

	case value.resolved.IsBuiltin("bool"):
		if !value.isPtr(0) {
			return conv.convertRef(value, 0, 0)
		}

		switch types.CleanCType(value.resolved.CType, true) {
		case "gboolean":
			// Manually use C.TRUE.
			value.p.Linef("if %s { %s = C.TRUE }", value.InName, value.OutName)
		case "_Bool", "bool":
			// CGo supports casting the integer const to a C boolean.
			fallthrough
		default:
			value.p.Linef("if %s { %s = %s(1) }", value.InName, value.OutName, value.OutType)
		}

		return true

	case value.resolved.IsBuiltin("error"):
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 0)
		}

		value.header.ImportCore("gerror")
		value.header.Import("unsafe")

		value.p.Linef("%s = (*C.GError)(gerror.New(%s))", value.OutName, value.InName)
		if !value.isTransferring() {
			value.p.Linef("defer C.g_error_free(%s)", value.OutName)
		}
		return true

	case value.resolved.IsPrimitive():
		// Don't use the convertRef routine, because we might want to preserve
		// the pointer in case the API is weird.
		if value.resolved.Ptr > 0 {
			value.header.Import("unsafe")
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(%s))",
				value.OutName, value.OutType, value.InName,
			)
		} else {
			value.p.Linef("%s = %s(%s)", value.OutName, value.OutType, value.InName)
		}
		return true
	}

	// Only add these imports afterwards, since all imports above are manually
	// resolved.
	if value.needsNamespace {
		// We're using the PublicType, so add that import.
		value.header.ImportPubl(value.resolved)
	}

	switch types.EnsureNamespace(conv.sourceNamespace, value.AnyType.Type.Name) {
	case "gpointer":
		value.header.Import("unsafe")
		value.header.ImportCore("box")
		value.p.Linef(
			"%s = %s(box.Assign(unsafe.Pointer(%s)))",
			value.OutName, value.OutType, value.InName,
		)
		return true

	case "GObject.Type", "GType":
		value.header.NeedsGLibObject()
		// Just a primitive.
		value.p.Linef("%s = (%s)(%s)", value.OutName, value.OutType, value.InName)
		return true

	case "GObject.Value":
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 1)
		}

		value.header.NeedsGLibObject()
		// https://pkg.go.dev/github.com/gotk3/gotk3/glib#Type
		value.p.Linef(
			"%s = (%s)(unsafe.Pointer(&%s.GValue))",
			value.OutName, value.OutType, value.InName)
		return true

	case "GObject.Object", "GObject.InitiallyUnowned":
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 1)
		}

		value.header.Import("unsafe")
		value.p.Linef(
			"%s = (%s)(unsafe.Pointer(%s.Native()))",
			value.OutName, value.OutType, value.InName)
		return true
	}

	if value.resolved.Extern == nil {
		return false
	}

	switch v := value.resolved.Extern.Type.(type) {
	case *gir.Enum, *gir.Bitfield:
		if !value.isPtr(0) {
			return conv.convertRef(value, 0, 0)
		}

		value.p.Linef("%s = %s(%s)", value.OutName, value.OutType, value.InName)
		return true

	case *gir.Class, *gir.Interface:
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 0)
		}

		value.header.Import("unsafe")
		value.p.Linef(
			"%s = (%s)(unsafe.Pointer(%s.Native()))",
			value.OutName, value.OutType, value.InName,
		)
		return true

	case *gir.Record:
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 1)
		}

		value.header.Import("unsafe")
		value.p.Linef(
			"%s = (%s)(unsafe.Pointer(%s.Native()))",
			value.OutName, value.OutType, value.InName,
		)
		return true

	case *gir.Callback:
		exportedName := value.resolved.Extern.Name()
		exportedName = strcases.PascalToGo(exportedName)

		// Callbacks must have the closure attribute to store the closure
		// pointer.
		if value.Closure == nil {
			conv.Logln(logger.Debug, "Go->C callback", exportedName, "since missing closure")
			return false
		}

		value.header.ImportCore("box")
		value.header.AddCallback(v)

		// Return the constant function here. The function will dynamically load
		// the user_data, which will match with the "gpointer" case above.
		//
		// As for the pointer to byte array cast, see
		// https://github.com/golang/go/issues/19835.
		value.p.Linef("%s = (*[0]byte)(C.%s%s)", value.OutName, file.CallbackPrefix, exportedName)

		closure := conv.convertParam(*value.Closure)
		if closure == nil {
			return false
		}

		value.outDecl.Linef("var %s %s", closure.OutName, closure.OutType)
		value.p.Linef("%s = %s(box.Assign(%s))", closure.OutName, closure.OutType, value.InName)

		if value.Destroy != nil {
			value.header.CallbackDelete = true

			destroy := conv.convertParam(*value.Destroy)
			if destroy == nil {
				return false
			}

			value.outDecl.Linef("var %s %s", destroy.OutName, destroy.OutType)
			value.p.Linef(
				"%s = (%s)((*[0]byte)(C.callbackDelete))",
				destroy.OutName, destroy.OutType,
			)
		}

		return true

		// case value.AllowNone:
		// 	value.p.Linef("var %s %s // unsupported", value.OutName, value.OutType)
	}

	return false
}

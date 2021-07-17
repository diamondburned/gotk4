package typeconv

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
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
		return false
	}

	array := *value.AnyType.Array

	// Ensure that the array type matches the inner type. Some functions violate
	// this, e.g. g_spawn_command_line_sync().
	if array.Type.CType == "" {
		// Copy the inner type so we don't accidentally change a reference.
		typ := *array.Type

		if types.CleanCType(array.CType, true) == "void" {
			// We can't dereference the array type to get a void type, so we
			// have to guess from the GIR type. Thanks, GIR.
			typ.CType = types.CTypeFallback("", typ.Name)
		} else {
			// Dereference the inner type by 1.
			typ.CType = strings.Replace(array.CType, "*", "", 1)
		}

		array.AnyType.Type = &typ
		value.AnyType.Array = &array
	}

	if types.CleanCType(array.CType, true) == "void" {
		// CGo treats void* arrays a bit weirdly: the function's input parameter
		// type is actually unsafe.Pointer, so we have to wrap it.
		value.Out.Call = fmt.Sprintf("unsafe.Pointer(%s)", value.Out.Call)
	}

	// This is always the same.
	value.Out.Type = types.AnyTypeCGo(value.AnyType)
	value.outDecl.Linef("var %s %s", value.OutName, value.Out.Type)

	// Length is roughly always the same as well.
	var length *ValueConverted
	if array.Length != nil {
		length = conv.convertParam(*array.Length)
		// Ensure length is present.
		if length == nil {
			return false
		}

		// Multiple arrays may use the same length value.
		if length.finalize() {
			value.outDecl.Linef("var %s %s", length.OutName, length.Out.Type)
		}

		// Length has no input, as it's from the slice.
		value.p.Linef("%s = (%s)(len(%s))", length.Out.Set, length.Out.Type, value.InName)
	}

	inner := conv.convertInner(value, value.InName+"[i]", "out[i]")
	if inner == nil {
		return false
	}

	// Set the array value's resolved type to the inner type.
	value.Resolved = inner.Resolved
	value.NeedsNamespace = inner.NeedsNamespace

	// These cases have invalid inner type names that aren't useful to us, so we
	// handle them on our own.
	switch {
	case types.CleanCType(array.CType, false) == "gchar*":
		// This is technically a []byte, and we should use a []byte, because the
		// C code may mutate the backing array. The internal type is "utf8",
		// which is false, because the C type is just a single character.
		value.In.Type = "[]byte"
		value.inDecl.Linef("var %s []byte", value.InName)

		if array.IsZeroTerminated() {
			// Ensure that the input has a null byte.
			value.p.Linef("%s = append(%[1]s, 0)", value.InName)
		}

		// Only hand over Go memory if we don't have to reallocate.
		if !value.MustRealloc() {
			value.header.Import("unsafe")
			value.p.Linef("if len(%s) > 0 {", value.InName)
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s[0]))",
				value.Out.Set, value.Out.Type, value.InName,
			)
			value.p.Linef("}")
			return true
		}

		// Use CBytes, which copies. Since we're transferring ownership to C,
		// don't free it after we're done. CString is actually null-terminated,
		// but we're giving the caller the length, so an extra byte is fine.
		value.p.Linef(
			"%s = (%s)(C.CBytes(%s))",
			value.Out.Set, value.Out.Type, value.InName,
		)

		return true
	}

	if array.FixedSize > 0 {
		value.In.Type = fmt.Sprintf("[%d]%s", array.FixedSize, inner.In.Type)
		value.inDecl.Linef("var %s %s", value.InName, value.In.Type)
	} else {
		value.In.Type = fmt.Sprintf("[]%s", inner.In.Type)
		value.inDecl.Linef("var %s %s", value.InName, value.In.Type)
	}

	// TODO: PtrArray

	switch {
	case array.FixedSize > 0:
		// Safe to do if this is a primitive AND we're not setting this inside a
		// calllback, since the callback will retain Go memory beyond its
		// lifetime which is bad.
		if !value.MustRealloc() && inner.Resolved.CanCast() {
			// We can directly use Go's array as a pointer, as long as we defer
			// properly.
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s))",
				value.Out.Set, value.Out.Type, value.InName)

			return true
		}

		// Target fixed array, so we can directly set the data over. The memory
		// is ours, and allocation is handled by Go.
		value.header.ApplyFrom(inner.Header())
		value.p.Descend()

		if !value.MustRealloc() {
			// We can allocate this on Go's stack or heap.
			value.p.Linef("%s = &([%d]%s)", value.Out.Set, array.FixedSize, inner.Out.Type)
		} else {
			value.p.Linef("%s = %s", value.Out.Set, inner.cmalloc(value.InName, false))
			if value.ShouldFree() {
				value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.Out.Set)
			}
		}

		value.p.Linef("out := (*[%d]%s)(unsafe.Pointer(%s))",
			array.FixedSize, inner.Out.Type, value.Out.Set)

		value.p.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		value.p.Linef("  " + inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	case array.Length != nil:
		// Use the backing array with the appropriate transfer-ownership rule
		// for primitive types; see type_c_go.go.
		if !value.MustRealloc() && inner.Resolved.CanCast() {
			value.header.Import("unsafe")
			value.p.Linef("if len(%s) > 0 {", value.InName)
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s[0]))",
				value.Out.Set, value.Out.Type, value.InName)
			value.p.Linef("}")

			return true
		}

		value.header.Import("unsafe")

		value.p.Linef(
			"%s = (%s)(%s)",
			value.Out.Set, value.Out.Type, inner.cmalloc(value.InName, false),
		)
		if value.ShouldFree() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)

		} else if inner.Resolved.CanCast() {
			// Edge case: we can use the optimized copy() built-in if the type
			// is castable. This case should never be hit if ShouldFree is true,
			// because we're still using Go memory.
			value.p.Linef(
				"copy(unsafe.Slice((*%s)(%s)), %s)",
				inner.In.Type, value.OutName, value.InName,
			)
			return true
		}

		value.header.ApplyFrom(inner.Header())
		value.p.Descend()

		value.p.Linef(
			// Use the pointer to inner type here, because the array might be a
			// void pointer. GIR should use equivalent types, so this should be
			// fine.
			"out := unsafe.Slice((*%s)(%s), len(%s))",
			inner.Out.Type, value.OutName, value.InName,
		)

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
			value.Out.Set, array.IsZeroTerminated(), inner.csizeof(), value.InName)
		value.p.Linef(
			"%s = C.g_array_set_size(%s, C.guint(len(%s)))",
			value.Out.Set, value.OutName, value.InName)

		if value.ShouldFree() {
			value.p.Linef("defer C.g_array_unref(%s)", value.OutName)
		}

		value.header.ApplyFrom(inner.Header())
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

		// https://developer.gnome.org/glib/stable/glib-Byte-Arrays.html#g-byte-array-new
		value.p.Linef(
			"%s = C.g_byte_array_sized_new(C.guint(len(%s)))",
			value.Out.Set, value.InName)

		value.p.Linef("if len(%s) > 0 {", value.InName)
		value.p.Linef(
			"%s = C.g_byte_array_append(%s, (*C.guint8)(&%s[0]), C.guint(len(%s)))",
			value.Out.Set, value.OutName, value.InName, value.InName)
		value.p.Linef("}")

		// unref will free the underlying array as well.
		value.p.Linef("defer C.g_byte_array_unref(%s)", value.OutName)
		return true

	case array.IsZeroTerminated():
		value.header.Import("unsafe")

		value.p.Descend()
		defer value.p.Ascend()

		// See if we can possibly reuse the Go slice in a shorter way.
		if !value.MustRealloc() && inner.Resolved.CanCast() {
			value.p.Linef("var zero %s", inner.In.Type)
			value.p.Linef("%s = append(%[1]s, zero)", value.InName)
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s[0]))",
				value.Out.Set, value.Out.Type, value.InName,
			)
			return true
		}

		value.p.Linef(
			"%s = (%s)(%s)",
			value.Out.Set, value.Out.Type, inner.cmalloc(value.InName, true))
		if value.ShouldFree() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)
		}

		value.header.ApplyFrom(inner.Header())
		value.p.Descend()

		value.p.Linef("out := unsafe.Slice(%s, len(%s)+1)", value.OutName, value.InName)
		// malloc does not zero out the memory, so we have to zero it out
		// ourselves.
		value.p.Linef("var zero %s", inner.Out.Type)
		value.p.Linef("out[len(%s)] = zero", value.InName)
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
		if unsupported == types.CleanCType(value.AnyType.Type.CType, true) {
			return false
		}
	}

	if !value.resolveType(conv) {
		value.Logln(logger.Debug, "cannot resolve type")
		return false
	}

	switch {
	case value.Resolved.IsBuiltin("cgo.Handle"):
		value.header.Import("runtime/cgo")
		value.header.Import("unsafe")
		// unsafe.Pointer is needed for pointer to pointers, so we're playing it
		// safe.
		value.p.Linef("%s = (%s)(unsafe.Pointer(%s))", value.Out.Set, value.Out.Type, value.InName)
		return true

	case value.Resolved.IsBuiltin("string"):
		if !value.isPtr(1) {
			return false
		}

		// Cast using an unsafe.Pointer in case the output type is uchar and Go
		// refuses to compile it.
		value.header.Import("unsafe")
		value.p.Linef(
			"%s = (%s)(unsafe.Pointer(C.CString(%s)))",
			value.Out.Set, value.Out.Type, value.InName,
		)
		// If we're not giving ownership this mallocated string, then we
		// can free it once done.
		if value.ShouldFree() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)
		}

		return true

	case value.Resolved.IsBuiltin("bool"):
		switch types.CleanCType(value.Resolved.CType, true) {
		case "gboolean":
			// Manually use C.TRUE.
			value.p.LineTmpl(value, `if <.InPtr 0><.In.Name> {
				<.OutPtr 0><.Out.Set> = C.TRUE
			}`)
		case "_Bool", "bool":
			// CGo supports casting the integer const to a C boolean.
			fallthrough
		default:
			value.p.LineTmpl(value, `if <.InPtr 0><.In.Name> {
				<.OutPtr 0><.Out.Set> = <.Out.Type>(1)
			}`)
		}

		return true

	case value.Resolved.IsBuiltin("error"):
		if !value.isPtr(1) {
			return false
		}

		value.header.ImportCore("gerror")
		value.header.Import("unsafe")

		value.p.LineTmpl(value, "<.Out.Set> = <.OutCast 1>(gerror.New(<.InNamePtr 0>))")

		// TODO: figure this out.
		// if value.ShouldFree() {
		// 	value.p.Linef("if %s != nil {", value.Out.Set)
		// 	value.p.Linef("  defer C.g_error_free(%s)", value.Out.Set)
		// 	value.p.Linef("}")
		// }
		return true

	case value.Resolved.IsPrimitive():
		// Don't use the convertRef routine, because we might want to preserve
		// the pointer in case the API is weird.
		if value.Resolved.Ptr > 0 {
			value.header.Import("unsafe")
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(%s))",
				value.Out.Set, value.Out.Type, value.InName,
			)
		} else {
			value.p.Linef("%s = %s(%s)", value.Out.Set, value.Out.Type, value.InName)
		}
		return true
	}

	switch types.EnsureNamespace(conv.sourceNamespace, value.AnyType.Type.Name) {
	case "GObject.Type", "GType":
		value.header.NeedsGLibObject()
		value.p.LineTmpl(value, "<.Out.Set> = <.OutCast 0>(<.InNamePtr 0>)")
		return true

	case "GObject.Value":
		// https://pkg.go.dev/github.com/gotk3/gotk3/glib#Type
		value.header.NeedsGLibObject()
		value.p.LineTmpl(value,
			"<.Out.Set> = <.OutCast 1>(unsafe.Pointer(&<.InNamePtr 1>.GValue))")
		return true

	case "GObject.Object", "GObject.InitiallyUnowned":
		value.header.Import("unsafe")
		value.p.LineTmpl(value,
			"<.Out.Set> = <.OutCast 1>(unsafe.Pointer(<.InNamePtrPubl 1>.Native()))")
		return true

	case "cairo.Context", "cairo.Pattern", "cairo.Region", "cairo.Surface":
		value.header.Import("unsafe")
		value.p.Linef(
			"%s = (%s)(unsafe.Pointer(%s.Native()))",
			value.Out.Set, value.OutCast(1), value.InNamePtr(1),
		)
		return true
	}

	if value.Resolved.Extern == nil {
		return false
	}

	switch v := value.Resolved.Extern.Type.(type) {
	case *gir.Enum, *gir.Bitfield:
		value.p.LineTmpl(value, "<.Out.Set> = <.OutCast 0>(<.InNamePtr 0>)")
		return true

	case *gir.Class, *gir.Interface:
		value.header.Import("unsafe")

		name := value.InNamePtrPubl(1)
		if value.IsPublic {
			// Public interfaces don't have .Native, so we type-assert it.
			value.header.ImportCore("gextras")
			name = fmt.Sprintf("(%s).(gextras.Nativer)", name)
		}

		value.p.Linef("%s = %s(unsafe.Pointer(%s.Native()))", value.Out.Set, value.OutCast(1), name)
		return true

	case *gir.Record:
		value.header.Import("unsafe")
		value.header.ImportCore("gextras")
		value.vtmpl(
			"<.Out.Set> = <.OutCast 1>(gextras.StructNative(unsafe.Pointer(<.InNamePtr 1>)))",
		)

		// This code might trigger a double-free.
		// // If ShouldFree is true, then ideally, we'll be freeing the C copy of
		// // the value once we're done. However, since we're passing by reference,
		// // we can't simply do that. Instead, if we cannot free the data once
		// // we're done, then we detach the finalizer so Go can't.
		// if !value.ShouldFree() && types.RecordHasRef(v) == nil {
		// 	// value.header.Import("runtime")
		// 	value.Logln(logger.Error, "adding runtime")
		// 	value.p.Linef("runtime.SetFinalizer(%s, nil)", value.InName)
		// }
		return true

	case *gir.Callback:
		exportedName := file.CallbackExportedName(value.Resolved.Extern.NamespaceFindResult, v)

		// Callbacks must have the closure attribute to store the closure
		// pointer.
		if value.Closure == nil {
			conv.Logln(logger.Debug, exportedName, "missing closure")
			return false
		}

		closure := conv.param(*value.Closure)
		if closure == nil || closure.Type == nil {
			value.Logln(logger.Debug, exportedName, "closure", *value.Closure, "not found")
			return false
		}

		value.header.ImportCore("gbox")
		value.header.AddCallback(value.Resolved.Extern.NamespaceFindResult, v)

		// Return the constant function here. The function will dynamically load
		// the user_data, which will match with the "gpointer" case above.
		//
		// As for the pointer to byte array cast, see
		// https://github.com/golang/go/issues/19835.
		value.p.Linef("%s = (*[0]byte)(C.%s)", value.Out.Set, exportedName)

		scope := value.Scope
		if scope == "" {
			// https://wiki.gnome.org/Projects/GObjectIntrospection/Annotations
			scope = "call"
		}

		assign := "Assign"
		if scope == "async" {
			// AssignOnce will pop the callback once it's called.
			assign = "AssignOnce"
		}

		userDataType := types.AnyTypeCGo(closure.AnyType)

		value.outDecl.Linef("var %s %s", closure.OutName, userDataType)
		value.p.Linef(
			"%s = %s(gbox.%s(%s))",
			closure.Out.Set, userDataType, assign, value.InName,
		)

		switch scope {
		case "call":
			value.p.Linef("defer gbox.Delete(uintptr(%s))", closure.Out.Set)
		case "async":
			// Handled in AssignOnce.
		case "notified":
			if value.Destroy == nil {
				value.Logln(logger.Debug, "scope=notified missing destroy param")
				return false
			}

			// Check if destroy's Resolved type is nil
			destroy := conv.param(*value.Destroy)
			if destroy == nil {
				value.Logln(logger.Debug, "cannot find destroy param, allowing anyway...")
				return true
			}
			if destroy.Type == nil || destroy.Type.Name != "GLib.DestroyNotify" {
				value.Logln(logger.Debug, "unknown destroyer type, allowing anyway...")
				return true
			}

			// Mark this as done.
			if destroy.finalize() {
				value.header.CallbackDelete = true
				value.outDecl.Linef("var %s C.GDestroyNotify", destroy.OutName)
			}

			value.p.Linef(
				"%s = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))",
				destroy.OutName,
			)
		default:
			return false
		}

		return true
	}

	if value.Optional {
		value.p.Linef("var %s %s // unsupported", value.OutName, value.Out.Type)
		return true
	}

	return false
}

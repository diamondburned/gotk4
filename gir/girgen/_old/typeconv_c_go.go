package girgen

import (
	"fmt"
	"strings"
)

// C to Go type conversions.

func (conv *TypeConverter) cgoConvert(value *ValueConverted) bool {
	switch {
	case value.AnyType.Array != nil:
		return conv.cgoArrayConverter(value)
	case value.AnyType.Type != nil:
		return conv.cgoTypeConverter(value)
	default:
		return false
	}
}

func (conv *TypeConverter) cgoArrayConverter(value *ValueConverted) bool {
	if value.AnyType.Array.Type == nil {
		conv.log(LogDebug, "C->Go skipping nested array", value.AnyType.Array.CType)
		return value.Optional // ok if optional
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

	// All generators must declare src.
	inner := conv.convertInner(value, "src[i]", value.OutName+"[i]")
	if inner == nil {
		return false
	}

	value.InType = anyTypeCGo(value.AnyType)
	if value.ParameterIsOutput() {
		// Dereference the input type, as we'll be passing in references.
		value.InType = strings.TrimPrefix(value.InType, "*")
	}

	if array.FixedSize > 0 && value.outputAllocs() {
		value.inDecl.Linef("var %s [%d]%s", value.InName, array.FixedSize, value.InType)
		// We've allocated an array, so have C write to this array.
		value.InCall = fmt.Sprintf("&%s[0]", value.InName)
	} else {
		value.inDecl.Linef("var %s %s", value.InName, value.InType)
		// Slice allocations are done later, since we don't know the length yet.
		// CallerAllocates is probably impossible to do here.
		value.InCall = fmt.Sprintf("&%s", value.InName)
	}

	if array.FixedSize > 0 {
		value.OutType = fmt.Sprintf("[%d]%s", array.FixedSize, inner.OutType)
		value.outDecl.Linef("var %s %s", value.OutName, value.OutType)
	} else {
		value.OutType = fmt.Sprintf("[]%s", inner.OutType)
		value.outDecl.Linef("var %s %s", value.OutName, value.OutType)
	}

	switch {
	case array.FixedSize > 0:
		value.addImport("unsafe")

		// Detect if the underlying is a compatible Go primitive type that isn't
		// a string. If it is, then we can directly cast the fixed-size array
		// pointer.
		if inner.resolved.CanCast() {
			value.p.Linef(
				"%s = *(*%s)(unsafe.Pointer(&%s))",
				value.OutName, value.OutType, value.InName)
			return true
		}

		value.p.Descend()

		// Direct cast is not possible; make a temporary array with the CGo type
		// so we can loop over it easily.
		value.p.Linef("src := &%s", value.InName)
		value.p.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	case array.Length != nil:
		value.addImport("unsafe")

		length := conv.convertParam(*array.Length)
		if length == nil {
			return false
		}

		value.inDecl.Linef("var %s %s // in", length.InName, length.InType)
		// Length has no outDecl.

		// If we're owning the new data, then we will directly use the backing
		// array, but we can only do this if the underlying type is a primitive,
		// since those have equivalent Go representations. Any other types will
		// have to be copied or otherwise converted somehow.
		//
		// TODO: record conversion should handle ownership: if
		// transfer-ownership is none, then the native pointer should probably
		// not be freed.
		if value.isTransferring() && inner.resolved.CanCast() {
			value.addImport("runtime")

			value.p.Linef("%s = unsafe.Slice((*%s)(unsafe.Pointer(%s)), %s)",
				value.OutName, inner.OutType, value.InName, length.InName)

			// See: https://golang.org/misc/cgo/gmp/gmp.go?s=3086:3757#L87
			value.p.Linef("runtime.SetFinalizer(&%s, func(v *%s) {", value.OutName, value.OutType)
			value.p.Linef("  C.free(unsafe.Pointer(&(*v)[0]))")
			value.p.Linef("})")

			return true
		}

		// Make sure to free the input by the time we're done.
		if value.isTransferring() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.InName)
		}

		if inner.resolved.CanCast() {
			// We can cast directly, which means no conversion is needed. Use
			// the faster built-in copy() for this.
			value.p.Linef("%s = make(%s, %s)", value.OutName, value.OutType, length.InName)
			value.p.Linef(
				"copy(%s, unsafe.Slice((*%s)(unsafe.Pointer(%s)), %s))",
				value.OutName, inner.OutType, value.InName, length.InName)
			return true
		}

		value.p.Descend()
		value.p.Linef("src := unsafe.Slice(%s, %s)", value.InName, length.InName)
		value.p.Linef("%s = make(%s, %s)", value.OutName, value.OutType, length.InName)
		value.p.Linef("for i := 0; i < int(%s); i++ {", length.InName)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")
		value.p.Ascend()
		return true

	case array.Name == "GLib.Array": // treat as Go array
		value.addImport("unsafe")

		value.p.Descend()

		value.p.Linef("var len uintptr")
		value.p.Linef("p := C.g_array_steal(&%s, (*C.gsize)(&len))", value.InName)
		value.p.Linef("src := unsafe.Slice((*%s)(p), len)", inner.InType)
		value.p.Linef("%s = make(%s, len)", value.OutName, value.OutType)
		value.p.Linef("for i := 0; i < len; i++ {", value.InName)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	case array.Name == "GLib.ByteArray":
		value.addImport("unsafe")

		if value.isTransferring() {
			value.p.Descend()
			value.p.Linef("var len uintptr")
			// If we're fully getting the backing array, then we can just steal
			// it (since we own it now), which is less copying.
			value.p.Linef("p := C.g_byte_array_steal(&%s, (*C.gsize)(&len))", value.InName)
			value.p.Linef("%s = unsafe.Slice((*byte)(p), len)", value.OutName)
			value.p.Linef("runtime.SetFinalizer(&%s, func(v *[]byte) {")
			value.p.Linef("  C.free(unsafe.Pointer(&(*v)[0]))")
			value.p.Linef("})")
			value.p.Ascend()
			return true
		}

		value.p.Linef("%s = make([]byte, %s.len)", value.OutName, value.InName)
		value.p.Linef(
			// Use the built-in copy(), because it is fast.
			"copy(%s, unsafe.Slice((*byte)(%s.data), %[2]s.len))",
			value.OutName, value.InName)
		return true

	case array.IsZeroTerminated():
		value.addImport("unsafe")

		value.p.Descend()

		// Scan for the length.
		value.p.Linef("var i int")
		value.p.Linef("var z %s", inner.InType)
		value.p.Linef("for p := %s; *p != z; p = &unsafe.Slice(p, i+1)[i] {", value.InName)
		value.p.Linef("  i++")
		value.p.Linef("}")
		value.p.EmptyLine()

		value.p.Linef("src := unsafe.Slice(%s, i)", value.InName)
		value.p.Linef("%s = make(%s, i)", value.OutName, value.OutType)
		value.p.Linef("for i := range src {")
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	default:
		conv.log(LogSkip, "C->Go weird array type", array.Type)
	}

	return false
}

func (conv *TypeConverter) cgoTypeConverter(value *ValueConverted) bool {
	for _, unsupported := range unsupportedCTypes {
		if unsupported == value.AnyType.Type.Name {
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

		value.p.Linef("%s = C.GoString(%s)", value.OutName, value.InName)
		// Only free this if C is transferring ownership to us.
		if value.isTransferring() {
			value.addImport("unsafe")
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.InName)
		}
		return true

	case value.resolved.IsBuiltin("bool"):
		if !value.isPtr(0) {
			return conv.convertRef(value, 0, 0)
		}

		switch cleanCType(value.resolved.CType, true) {
		case "gboolean":
			// gboolean is resolved to C type int, so we have to do regular int
			// comparison.
			value.p.Linef("if %s != 0 { %s = true }", value.InName, value.OutName)
		case "_Bool", "bool":
			fallthrough
		default:
			// CGo supports _Bool and bool directly.
			value.p.Linef("if %s { %s = true }", value.InName, value.OutName)
		}
		return true

	case value.resolved.IsBuiltin("error"):
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 0)
		}

		value.addImportInternal("gerror")
		value.addImport("unsafe")

		value.p.Linef("%s = gerror.Take(unsafe.Pointer(%s))", value.OutName, value.InName)
		return true

	case value.resolved.IsPrimitive():
		// Don't use the convertRef routine, because we might want to preserve
		// the pointer in case the API is weird.
		if value.resolved.Ptr > 0 {
			value.addImport("unsafe")
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
		value.importPubl(value.resolved)
	}

	// Resolve special-case GLib types.
	switch ensureNamespace(conv.ng.current, value.AnyType.Type.Name) {
	case "gpointer":
		value.addImportInternal("box")
		value.p.Linef("%s = box.Get(uintptr(%s))", value.OutName, value.InName)
		return true

	case "GObject.Type", "GType":
		value.needsExternGLib()
		value.needsGLibObject()
		value.p.Linef("%s = externglib.Type(%s)", value.OutName, value.InName)
		return true

	case "GObject.Value":
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 1)
		}

		value.addImport("unsafe")
		value.needsExternGLib()
		value.needsGLibObject()

		value.p.Linef(
			"%s = externglib.ValueFromNative(unsafe.Pointer(%s))",
			value.OutName, value.InName,
		)

		// Set this to be freed if we have the ownership now.
		if value.isTransferring() {
			value.addImport("runtime")

			// https://pkg.go.dev/github.com/gotk3/gotk3/glib?utm_source=godoc#Value
			value.p.Linef("runtime.SetFinalizer(%s, func(v *externglib.Value) {", value.OutName)
			value.p.Linef("  C.g_value_unset((*C.GValue)(v.GValue))")
			value.p.Linef("})")
		}
		return true

	case "GObject.Object", "GObject.InitiallyUnowned":
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 1)
		}

		value.cgoSetObject()
		return true
	}

	// TODO: function
	// TODO: union
	// TODO: callback

	// TODO: callbacks and functions are handled differently. Unsure if they're
	// doable.
	// TODO: handle unions.

	if value.resolved.Extern == nil {
		return false
	}

	result := value.resolved.Extern.Result

	switch {
	case result.Enum != nil, result.Bitfield != nil:
		if !value.isPtr(0) {
			return conv.convertRef(value, 0, 0)
		}

		value.p.Linef("%s = %s(%s)", value.OutName, value.OutType, value.InName)
		return true

	case result.Class != nil, result.Interface != nil:
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 0)
		}

		value.cgoSetObject()
		return true

	case result.Record != nil:
		// // We can slightly cheat here. Since Go structs are declared by wrapping
		// // the C type, we can directly cast to the C type if this is an output
		// // parameter. This saves us a copy.
		// if value.resolved.Ptr < 2 && value.outputAllocs() {
		// 	value.addImport("unsafe")

		// 	value.outDecl.Reset()
		// 	value.inDecl.Reset()

		// 	// Write the Go type directly.
		// 	value.inDecl.Linef("var %s %s", value.OutName, value.OutType)
		// 	value.InCall = fmt.Sprintf("(%s)(%s)", value.InType, value.OutName)

		// 	return true
		// }

		// Require 1 pointer to avoid weird copies.
		if !value.isPtr(1) {
			return conv.convertRef(value, 1, 1)
		}

		value.addImport("unsafe")
		value.p.Linef("%s = (%s)(unsafe.Pointer(%s))", value.OutName, value.OutType, value.InName)

		if value.isTransferring() {
			value.addImport("runtime")

			value.p.Linef("runtime.SetFinalizer(&%s, func(v *%s) {", value.OutName, value.OutType)
			value.p.Linef("  C.free(unsafe.Pointer(v))")
			value.p.Linef("})")
		}
		return true

	case result.Alias != nil:
		// underlying := conv.ng.FindType(result.Alias.Name)
		// if underlying == nil {
		// 	conv.fail()
		// 	return
		// }

		// resolved := conv.ng.ResolveType(underlying)

		// TODO: find a way to construct the output wrapper. Easiest way is to
		// output to a tmp variable and convert back, but this would require
		// putting this inside a block.

		// TODO
		// return false

		// case value.AllowNone:
		// 	value.outDecl.Linef("var %s %s // unsupported", value.OutName, value.OutType)
		// 	return true
	}

	return false
}

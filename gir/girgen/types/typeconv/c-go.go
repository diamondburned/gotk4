package typeconv

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// C to Go type conversions.

func (conv *Converter) cgoParameterOverrides(value *ValueConverted) {
	// noop
}

func (conv *Converter) cgoConvert(value *ValueConverted) bool {
	switch {
	case value.AnyType.Array != nil:
		return conv.cgoArrayConverter(value)
	case value.AnyType.Type != nil:
		return conv.cgoConvertNested(value)
	default:
		return false
	}
}

func (conv *Converter) cgoArrayConverter(value *ValueConverted) bool {
	if value.AnyType.Array.Type == nil {
		conv.Logln(logger.Debug, "C->Go skipping nested array", value.AnyType.Array.CType)
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

	// All generators must declare src.
	inner := conv.convertInner(value, "src[i]", value.OutName+"[i]")
	if inner == nil {
		return false
	}

	// Set the array value's resolved type to the inner type.
	value.Resolved = inner.Resolved
	value.NeedsNamespace = inner.NeedsNamespace

	value.In.Type = types.AnyTypeCGo(value.AnyType)
	if value.ParameterIsOutput() {
		// Dereference the input type, as we'll be passing in references.
		value.In.Type = strings.TrimPrefix(value.In.Type, "*")
	}

	if array.FixedSize > 0 && value.outputAllocs() {
		value.inDecl.Linef("var %s [%d]%s", value.InName, array.FixedSize, value.In.Type)
		// We've allocated an array, so have C write to this array.
		value.In.Call = fmt.Sprintf("&%s[0]", value.InName)
	} else {
		value.inDecl.Linef("var %s %s", value.InName, value.In.Type)
		// Slice allocations are done later, since we don't know the length yet.
		// CallerAllocates is probably impossible to do here.
		value.In.Call = fmt.Sprintf("&%s", value.InName)
	}

	if array.FixedSize > 0 {
		value.Out.Type = fmt.Sprintf("[%d]%s", array.FixedSize, inner.Out.Type)
		value.outDecl.Linef("var %s %s", value.OutName, value.Out.Type)
	} else {
		value.Out.Type = fmt.Sprintf("[]%s", inner.Out.Type)
		value.outDecl.Linef("var %s %s", value.OutName, value.Out.Type)
	}

	switch {
	case array.FixedSize > 0:
		value.header.Import("unsafe")

		// Detect if the underlying is a compatible Go primitive type that isn't
		// a string. If it is, then we can directly cast the fixed-size array
		// pointer.
		if inner.Resolved.CanCast() {
			value.p.Linef(
				"%s = *(*%s)(unsafe.Pointer(&%s))",
				value.Out.Set, value.Out.Type, value.InName)
			return true
		}

		value.header.ApplyFrom(inner.Header())
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
		value.header.Import("unsafe")

		length := conv.convertParam(*array.Length)
		if length == nil {
			return false
		}

		// Multiple arrays may use the same length value.
		if length.finalize() {
			value.header.ApplyFrom(length.Header())
			value.inDecl.Linef("var %s %s // in", length.InName, length.In.Type)
			// Length has no outDecl.
		}

		// If we're owning the new data, then we will directly use the backing
		// array, but we can only do this if the underlying type is a primitive,
		// since those have equivalent Go representations. Any other types will
		// have to be copied or otherwise converted somehow.
		//
		// TODO: record conversion should handle ownership: if
		// transfer-ownership is none, then the native pointer should probably
		// not be freed.
		if !value.MustRealloc() && inner.Resolved.CanCast() {
			value.header.Import("runtime")

			value.p.Linef("%s = unsafe.Slice((*%s)(unsafe.Pointer(%s)), %s)",
				value.Out.Set, inner.Out.Type, value.InName, length.InName)

			// See: https://golang.org/misc/cgo/gmp/gmp.go?s=3086:3757#L87
			value.p.Linef("runtime.SetFinalizer(&%s, func(v *%s) {", value.OutName, value.Out.Type)
			value.p.Linef("  C.free(unsafe.Pointer(&(*v)[0]))")
			value.p.Linef("})")

			return true
		}

		// Make sure to free the input by the time we're done.
		if value.ShouldFree() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.InName)
		}

		if inner.Resolved.CanCast() {
			// We can cast directly, which means no conversion is needed. Use
			// the faster built-in copy() for this.
			value.p.Linef("%s = make(%s, %s)", value.Out.Set, value.Out.Type, length.InName)
			value.p.Linef(
				"copy(%s, unsafe.Slice((*%s)(unsafe.Pointer(%s)), %s))",
				value.OutName, inner.Out.Type, value.InName, length.InName)
			return true
		}

		value.header.ApplyFrom(inner.Header())

		value.p.Descend()
		value.p.Linef("src := unsafe.Slice(%s, %s)", value.InName, length.InName)
		value.p.Linef("%s = make(%s, %s)", value.Out.Set, value.Out.Type, length.InName)
		value.p.Linef("for i := 0; i < int(%s); i++ {", length.InName)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")
		value.p.Ascend()
		return true

	case array.Name == "GLib.Array": // treat as Go array
		value.header.Import("unsafe")
		value.header.ApplyFrom(inner.Header())

		value.p.Descend()

		value.p.Linef("var len uintptr")
		value.p.Linef("p := C.g_array_steal(&%s, (*C.gsize)(&len))", value.InName)
		value.p.Linef("src := unsafe.Slice((*%s)(p), len)", inner.In.Type)
		value.p.Linef("%s = make(%s, len)", value.Out.Set, value.Out.Type)
		value.p.Linef("for i := 0; i < len; i++ {", value.InName)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	case array.Name == "GLib.ByteArray":
		value.header.Import("unsafe")

		if !value.MustRealloc() {
			value.header.Import("runtime")

			value.p.Descend()
			value.p.Linef("var len C.gsize")
			// If we're fully getting the backing array, then we can just steal
			// it (since we own it now), which is less copying.
			value.p.Linef("p := C.g_byte_array_steal(&%s, &len)", value.InName)
			value.p.Linef("%s = unsafe.Slice((*byte)(p), uint(len))", value.Out.Set)
			value.p.Linef("runtime.SetFinalizer(&%s, func(v *[]byte) {", value.OutName)
			value.p.Linef("  C.free(unsafe.Pointer(&(*v)[0]))")
			value.p.Linef("})")
			value.p.Ascend()
			return true
		}

		value.p.Linef("%s = make([]byte, %s.len)", value.Out.Set, value.InName)
		value.p.Linef(
			// Use the built-in copy(), because it is fast.
			"copy(%s, unsafe.Slice((*byte)(%s.data), %[2]s.len))",
			value.OutName, value.InName)
		return true

	case array.IsZeroTerminated():
		value.header.Import("unsafe")
		value.header.ApplyFrom(inner.Header())

		value.p.Descend()

		// Scan for the length.
		value.p.Linef("var i int")
		value.p.Linef("var z %s", inner.In.Type)
		value.p.Linef("for p := %s; *p != z; p = &unsafe.Slice(p, i+1)[i] {", value.InName)
		value.p.Linef("  i++")
		value.p.Linef("}")
		value.p.EmptyLine()

		value.p.Linef("src := unsafe.Slice(%s, i)", value.InName)
		value.p.Linef("%s = make(%s, i)", value.Out.Set, value.Out.Type)
		value.p.Linef("for i := range src {")
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	default:
		conv.Logln(logger.Skip, "C->Go weird array type", array.Type)
	}

	return false
}

func (conv *Converter) cgoConvertNested(value *ValueConverted) bool {
	if !value.resolveType(conv) {
		value.Logln(logger.Debug, "cannot resolve type")
		return false
	}

	if value.AnyType.Type.Type == nil {
		return conv.cgoConverter(value)
	}

	// TODO: GHashTable.
	switch {
	case value.Resolved.IsExternGLib("List"):
		value.header.Import("unsafe")

		value.vtmpl("<.Out.Set> = externglib.WrapList(uintptr(unsafe.Pointer(<.InNamePtr 1>)))")

		inner := conv.convertInner(value, "src", "dst")
		if inner != nil {
			value.p.Linef("%s.DataWrapper(func(_p unsafe.Pointer) interface{} {",
				value.OutInNamePtr(1))
			value.p.Linef("  src := (%s)(_p)", inner.In.Type)
			value.p.Linef("  %s", inner.Out.Declare)
			value.p.Linef("  %s", inner.Conversion)
			value.p.Linef("  return %s", inner.Out.Name)
			value.p.Linef("})")
		}

		switch value.TransferOwnership.TransferOwnership {
		case "container":
			value.vtmpl("<.OutInNamePtr 1>.AttachFinalizer(nil)")
		case "full":
			value.p.Linef("%s.AttachFinalizer(func(v uintptr) {", value.OutInNamePtr(1))
			value.p.Linef("  " + conv.cFree(inner, "unsafe.Pointer(v)"))
			value.p.Linef("})")
		}

		// Transfer over the header in the end.
		if inner != nil {
			inner.finalize()
			value.header.ApplyFrom(inner.Header())
		}

		return true
	}

	return false
}

// cFree calls the free function of the given value on the given string
// variable. v should be of type unsafe.Pointer. If value is nil, then a generic
// C.free is returned.
func (conv *Converter) cFree(value *ValueConverted, v string) string {
	switch {
	case value == nil:
		fallthrough
	case value.Resolved.IsBuiltin("string"):
		return fmt.Sprintf("C.free(%s)", v)
	}

	ptr := v
	v = fmt.Sprintf("(%s)(%s)", value.In.Type, v)

	switch value.Resolved.GType {
	case "GObject.Value": // *externglib.Value
		return fmt.Sprintf("C.g_value_unset(%s)", v)
	case "cairo.Context":
		return fmt.Sprintf("C.cairo_destroy(%s)", v)
	case "cairo.Surface":
		return fmt.Sprintf("C.cairo_surface_destroy(%s)", v)
	case "cairo.Pattern":
		return fmt.Sprintf("C.cairo_pattern_destroy(%s)", v)
	case "cairo.Region":
		return fmt.Sprintf("C.cairo_region_destroy(%s)", v)
	}

	if value.Resolved.Extern == nil {
		return fmt.Sprintf("C.free(%s)", v)
	}

	switch typ := value.Resolved.Extern.Type.(type) {
	case *gir.Class, *gir.Interface: // GObject
		return fmt.Sprintf("C.g_object_unref(C.gpointer(uintptr(%s)))", ptr)

	case *gir.Record:
		free := types.RecordHasUnref(typ)
		if free == nil {
			free = types.RecordHasFree(typ)
		}
		if free != nil {
			return fmt.Sprintf("C.%s(%s)", free.CIdentifier, v)
		}
		return fmt.Sprintf("C.free(unsafe.Pointer(%s))", v)

	case *gir.Alias:
		result := conv.convertType(
			value,
			value.InName,
			value.OutName,
			typ.Type.AnyType,
			value.TransferOwnership.TransferOwnership,
		)
		if result != nil {
			return conv.cFree(result, v)
		}
	}

	return fmt.Sprintf("C.free(%s)", v)
}

func (conv *Converter) cgoConverter(value *ValueConverted) bool {
	// TODO: make the freeing use cFree().

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
			value.Logln(logger.Debug, "weird string pointer rule")
			return false
		}

		// Preemptively cast the value to char*, since this might be used for
		// uchar as well.
		value.header.Import("unsafe")
		value.p.Linef(
			"%s = C.GoString((*C.gchar)(unsafe.Pointer(%s)))",
			value.Out.Set, value.InName,
		)
		// Only free this if C is transferring ownership to us.
		if value.ShouldFree() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.InName)
		}
		return true

	case value.Resolved.IsBuiltin("bool"):
		switch types.CleanCType(value.Resolved.CType, true) {
		case "gboolean":
			// gboolean is resolved to C type int, so we have to do regular int
			// comparison.
			value.p.LineTmpl(value, `
				if <.InPtr 0><.In.Name> != 0 {
					<.OutPtr 0><.Out.Set> = true
				}
			`)
		case "_Bool", "bool":
			fallthrough
		default:
			// CGo supports _Bool and bool directly.
			value.p.LineTmpl(value, `
				if <.InPtr 0><.In.Name> {
					<.OutPtr 0><.Out.Set> = true
				}
			`)
		}
		return true

	case value.Resolved.IsBuiltin("error"):
		if !value.isPtr(1) {
			value.Logln(logger.Debug, "weird GError pointer")
			return false
		}

		value.header.ImportCore("gerror")
		value.header.Import("unsafe")

		value.p.LineTmpl(value, "<.Out.Set> = gerror.Take(unsafe.Pointer(<.InNamePtr 1>))")
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

	// Resolve special-case GLib types.
	switch value.Resolved.GType {
	case "GObject.Type", "GType":
		value.header.NeedsExternGLib()
		value.header.NeedsGLibObject()
		value.p.LineTmpl(value, `<.Out.Set> = <.OutPtr 0>externglib.Type(<.InNamePtr 0>)`)
		return true

	case "GObject.Value":
		value.header.Import("unsafe")
		value.header.NeedsExternGLib()
		value.header.NeedsGLibObject()

		value.p.LineTmpl(value, `<.Out.Set> =
			<- .OutPtr 1>externglib.ValueFromNative(unsafe.Pointer(<.InNamePtr 1>))`)

		// Set this to be freed if we have the ownership now.
		if value.ShouldFree() {
			value.header.Import("runtime")

			// https://pkg.go.dev/github.com/gotk3/gotk3/glib?utm_source=godoc#Value
			value.p.Linef("runtime.SetFinalizer(%s, func(v *externglib.Value) {", value.OutName)
			value.p.Linef("  C.g_value_unset((*C.GValue)(unsafe.Pointer(v.GValue)))")
			value.p.Linef("})")
		}
		return true

	case "GObject.Object", "GObject.InitiallyUnowned":
		return value.cgoSetObject(conv)

	// These 4 cairo structs are found when grep -R-ing the codebase.
	case "cairo.Context", "cairo.Pattern", "cairo.Region", "cairo.Surface":
		var ref string
		var unref string

		value.header.Import("unsafe")
		value.header.Import("runtime")

		switch value.Resolved.GType {
		case "cairo.Context":
			ref = "cairo_reference"
			unref = "cairo_destroy"
			value.p.Linef(
				"%s = cairo.WrapContext(uintptr(unsafe.Pointer(%s)))",
				value.Out.Set, value.InNamePtr(1),
			)

		case "cairo.Surface":
			ref = "cairo_surface_reference"
			unref = "cairo_surface_destroy"
			value.p.Linef(
				"%s = cairo.WrapSurface(uintptr(unsafe.Pointer(%s)))",
				value.Out.Set, value.InNamePtr(1),
			)

		case "cairo.Pattern":
			ref = "cairo_pattern_reference"
			unref = "cairo_pattern_destroy"
			value.p.Descend()
			// Hack to fit the Pattern type.
			value.p.Linef("_pp:= &struct{p unsafe.Pointer}{unsafe.Pointer(%s)}", value.InNamePtr(1))
			value.p.Linef("%s = (*cairo.Pattern)(unsafe.Pointer(_pp))", value.Out.Set)
			value.p.Ascend()

		case "cairo.Region":
			ref = "cairo_region_reference"
			unref = "cairo_region_destroy"
			value.p.Descend()
			// Hack to fit the Region type.
			value.p.Linef("_pp:= &struct{p unsafe.Pointer}{unsafe.Pointer(%s)}", value.InNamePtr(1))
			value.p.Linef("%s = (*cairo.Region)(unsafe.Pointer(_pp))", value.Out.Set)
			value.p.Ascend()
		}

		// MustRealloc can also be used to check if we need to take a reference:
		// instead of reallocating, we take our own reference.
		if value.MustRealloc() {
			value.p.Linef("C.%s(%s)", ref, value.InNamePtr(1))
		}

		value.p.Linef("runtime.SetFinalizer(%s%s, func(v %s%s) {",
			value.OutInPtr(1), value.OutName, value.OutPtr(1), value.Out.Type)
		value.p.Linef("C.%s((%s%s)(unsafe.Pointer(v.Native())))",
			unref, value.InPtr(1), value.In.Type)
		value.p.Linef("})")

		return true
	}

	// TODO: function
	// TODO: union
	// TODO: callback

	// TODO: callbacks and functions are handled differently. Unsure if they're
	// doable.
	// TODO: handle unions.

	if value.Resolved.Extern == nil {
		value.Logln(logger.Debug, "unknown built-in")
		return false
	}

	switch v := value.Resolved.Extern.Type.(type) {
	case *gir.Enum, *gir.Bitfield:
		value.vtmpl("<.Out.Set> = <.OutCast 0>(<.InNamePtr 0>)")
		return true

	case *gir.Class, *gir.Interface:
		return value.cgoSetObject(conv)

	case *gir.Record:
		value.header.Import("unsafe")
		value.header.ImportCore("gextras")

		// Require 1 pointer to avoid weird copies.
		value.vtmpl(
			"<.Out.Set> = <.OutCast 1>(gextras.NewStructNative(unsafe.Pointer(<.InNamePtr 1>)))")
		if value.fail {
			value.Logln(logger.Debug, "record set fail")
			return false
		}

		var free *gir.Method
		var unref bool

		if ref := types.RecordHasRef(v); ref != nil && value.Resolved.Ptr > 0 {
			// MustRealloc can also be used to check if we need to take a
			// reference: instead of reallocating, we take our own reference.
			if value.MustRealloc() {
				value.p.Linef("C.%s(%s)", ref.CIdentifier, value.InNamePtr(1))
			}
			unref = true
			free = types.RecordHasUnref(v)
		} else {
			free = types.RecordHasFree(v)
		}

		// We can take ownership if the type can be reference-counted anyway.
		if value.ShouldFree() || unref {
			value.header.Import("runtime")
			value.vtmpl(
				"runtime.SetFinalizer(<.OutInPtr 1><.OutName>, func(v <.OutPtr 1><.Out.Type>) {")
			if value.fail {
				value.Logln(logger.Debug, "SetFinalizer set fail")
			}

			if free != nil {
				value.p.Linef(
					"C.%s((%s%s)(gextras.StructNative(unsafe.Pointer(v))))",
					free.CIdentifier, value.OutPtr(1), value.In.Type,
				)
			} else {
				value.p.Linef(
					"C.free(gextras.StructNative(unsafe.Pointer(v)))",
				)
			}
			value.p.Linef("})")
		}

		return true

	case *gir.Alias:
		result := conv.convertType(
			value,
			value.InName,
			value.OutName,
			v.Type.AnyType,
			value.TransferOwnership.TransferOwnership,
		)
		if result != nil {
			value.header.ApplyFrom(result.Header())
			return true
		}
		return false
	}

	if value.Optional {
		value.p.Linef("var %s %s // unsupported", value.OutName, value.Out.Type)
		return true
	}

	return false
}

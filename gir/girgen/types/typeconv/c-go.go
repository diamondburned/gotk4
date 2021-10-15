package typeconv

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// C to Go type conversions.

func (conv *Converter) cgoConvert(value *ValueConverted) bool {
	if types.CountPtr(value.In.Type) > 0 && value.IsOptional() {
		value.p.Linef("if %s != nil {", value.In.Name)
		defer value.p.Ascend()
	}

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
	array := *value.AnyType.Array

	// All generators must declare src.
	inner := conv.convertInner(value, "src[i]", value.OutName+"[i]")
	if inner == nil {
		return false
	}

	// Set the array value's resolved type to the inner type.
	value.Resolved = inner.Resolved
	value.NeedsNamespace = inner.NeedsNamespace

	// TODO: maybe remove this. I can't remember why this is here. Maybe the
	// edge cases done in resolveType are not applicable, or resolveType wasn't
	// called here at all.
	value.In.Type = types.AnyTypeCGo(value.AnyType)

	switch {
	case value.ParameterIndex == ReturnValueIndex && value.In.Type == "*C.void":
		// If cgoType is a *C.void return type, then turn it into an
		// unsafe.Pointer due to a quirk in cgo. See golang/go#21878.
		value.header.Import("unsafe")
		value.In.Type = "unsafe.Pointer"
	case value.ParameterIsOutput():
		// Dereference the input type, as we'll be passing in references.
		value.In.Type = strings.TrimPrefix(value.In.Type, "*")
	}

	// The earlier ResolveType routine will be setting the inDecl, but we have
	// our own rules for that.
	value.inDecl.Reset()

	if array.FixedSize > 0 && value.outputAllocs() {
		value.inDecl.Linef("var %s [%d]%s // in", value.InName, array.FixedSize, value.In.Type)
		// We've allocated an array, so have C write to this array.
		value.In.Call = fmt.Sprintf("&%s[0]", value.InName)
	} else {
		value.inDecl.Linef("var %s %s // in", value.InName, value.In.Type)
		// Slice allocations are done later, since we don't know the length yet.
		// CallerAllocates is probably impossible to do here.
		value.In.Call = fmt.Sprintf("&%s", value.InName)
	}

	switch {
	case array.FixedSize > 0:
		value.header.Import("unsafe")

		// Detect if the underlying is a compatible Go primitive type that isn't
		// a string. If it is, then we can directly cast the fixed-size array
		// pointer.
		if inner.Resolved.CanCast(conv.fgen) {
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
			value.Logln(logger.Debug, "length parameter", *array.Length, "not found")
			return false
		}

		// Multiple arrays may use the same length value.
		if length.finalize() {
			value.header.ApplyFrom(length.Header())
			value.inDecl.Linef("var %s %s // in", length.InName, length.In.Type)
			// Length has no outDecl.
		}

		// No realloc fast path available; setting a finalizer on a slice value
		// is invalid.

		// Make sure to free the input by the time we're done.
		if value.ShouldFree() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.InName)
		}

		if inner.Resolved.CanCast(conv.fgen) {
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
		value.p.Linef("for i := 0; i < len; i++ {")
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

		if value.ShouldFree() {
			// Just a regular array pointer.
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.InName)
		}

		value.p.Descend()

		// Scan for the length.
		value.p.Linef("var i int")
		value.p.Linef("var z %s", inner.In.Type)
		value.p.Linef("for p := %s; *p != z; p = &unsafe.Slice(p, 2)[1] {", value.InName)
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
	if len(value.Inner) == 0 {
		return conv.cgoConverter(value)
	}

	// Miraculously, convertInner will handle freeing for us. The only con is
	// that this approach will not work once we change from copying to
	// container-viewing, since the routine will end up freeing on view.
	//
	// This is a concern for future me, though.

	switch value.Resolved.GType {
	case "GLib.List", "GLib.SList":
		var sizeFn string
		var moveFn string

		switch value.Resolved.GType {
		case "GLib.List":
			sizeFn = "ListSize"
			moveFn = "MoveList"
		case "GLib.SList":
			sizeFn = "SListSize"
			moveFn = "MoveSList"
		}

		inner := conv.convertInner(value, "src", "dst")
		if inner == nil {
			value.Logln(logger.Debug, "List missing inner type")
			return false
		}

		value.header.Import("unsafe")
		value.header.ImportCore("gextras")
		value.header.ApplyFrom(inner.Header())

		value.p.Linef(
			"%s = make([]%s, 0, gextras.%s(unsafe.Pointer(%s)))",
			value.Out.Set, inner.Out.Type, sizeFn, value.InNamePtr(1))

		value.p.Linef("gextras.%s(unsafe.Pointer(%s), %t, func(v unsafe.Pointer) {",
			moveFn, value.InNamePtr(1), value.ShouldFree())
		value.p.Linef("  src := (%s)(v)", inner.In.Type)
		value.p.Linef("  %s", inner.Out.Declare)
		value.p.Linef("  %s", inner.Conversion)
		value.p.Linef("  %s = append(%[1]s, %s)", value.Out.Set, inner.Out.Name)
		value.p.Linef("})")

		return true

	case "GLib.HashTable":
		kt := conv.convertType(value, "ksrc", "kdst", &value.Type.Types[0])
		vt := conv.convertType(value, "vsrc", "vdst", &value.Type.Types[1])
		if kt == nil || vt == nil {
			value.Logln(logger.Debug, "no key/value-type")
			return false
		}

		value.header.Import("unsafe")
		value.header.ImportCore("gextras")
		value.header.ApplyFrom(kt.Header())
		value.header.ApplyFrom(vt.Header())

		value.p.Linef(
			"%s = make(%s, gextras.HashTableSize(unsafe.Pointer(%s)))",
			value.Out.Set, value.Out.Type, value.InNamePtr(1))

		value.p.Linef(
			"gextras.MoveHashTable(unsafe.Pointer(%s), %t, func(k, v unsafe.Pointer) {",
			value.InNamePtr(1), value.ShouldFree())
		value.p.Linef("ksrc := *(*%s)(k)", kt.In.Type)
		value.p.Linef("vsrc := *(*%s)(v)", vt.In.Type)
		value.p.Linef(kt.Out.Declare)
		value.p.Linef(vt.Out.Declare)
		value.p.Linef(kt.Conversion)
		value.p.Linef(vt.Conversion)
		value.p.Linef("%s[kdst] = vdst", value.Out.Set)
		value.p.Linef("})")

		return true
	}

	return false
}

/*
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
		result := conv.convertType(value, value.InName, value.OutName, &typ.Type)
		if result != nil {
			return conv.cFree(result, v)
		}
	}

	return fmt.Sprintf("C.free(%s)", v)
}
*/

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

			// https://pkg.go.dev/github.com/diamondburned/gotk4/pkg/core/glib?utm_source=godoc#Value
			value.p.Linef("runtime.SetFinalizer(%s, func(v *externglib.Value) {", value.OutName)
			value.p.Linef("  C.g_value_unset((*C.GValue)(unsafe.Pointer(v.Native())))")
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
		value.vtmpl(`
			<.Out.Set> = <.OutCast 1>(gextras.NewStructNative(unsafe.Pointer(<.InNamePtr 1>)))
		`)
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
			value.vtmpl(`
				runtime.SetFinalizer(
					gextras.StructIntern(unsafe.Pointer(<.OutInPtr 1><.OutName>)),
					func(intern *struct{ C unsafe.Pointer }) {
			`)
			if value.fail {
				value.Logln(logger.Debug, "SetFinalizer set fail")
			}

			if free != nil {
				value.p.Linef(
					"C.%s((%s)(intern.C))",
					free.CIdentifier, types.AnyTypeCGo(free.Parameters.InstanceParameter.AnyType),
				)
			} else {
				value.p.Linef("C.free(intern.C)")
			}

			value.p.Linef("},")
			value.p.Linef(")")
		}

		return true

	case *gir.Alias:
		typ := types.MoveTypePtr(*value.Type, v.Type)

		result := conv.convertType(value, value.InName, value.OutName, typ)
		if result == nil {
			return false
		}

		value.p.Line(result.Conversion)
		value.header.ApplyFrom(result.Header())
		return true
	}

	if value.Optional {
		value.p.Linef("var %s %s // unsupported", value.OutName, value.Out.Type)
		return true
	}

	return false
}

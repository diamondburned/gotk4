package girgen

import (
	"fmt"
)

// Go to C type conversions.

// TypeConversionToC describes a conversion of one or more Go values to C using
// CGo.
type TypeConversionToC struct {
	conversionTo
}

// GoCConverter returns a new converter that converts a value from Go to C.
func (fg *FileGenerator) GoCConverter(parent string, values []ValueProp) *TypeConversionToC {
	conv := &TypeConversionToC{}
	conv.conversionTo = newConversionTo(fg, parent, values, conv)
	return conv
}

func (conv *TypeConversionToC) convert(value *ValueConverted) bool {
	switch {
	case value.AnyType.Type != nil:
		return conv.typeConverter(value)
	case value.AnyType.Array != nil:
		return conv.arrayConverter(value)
	default:
		return false
	}
}

func (conv *TypeConversionToC) arrayConverter(value *ValueConverted) bool {
	if value.AnyType.Array.Type == nil {
		conv.log(LogDebug, "Go->C skipping nested array", value.AnyType.Array.CType)
		return value.AllowNone
	}

	array := value.AnyType.Array

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

	// This is always the same.
	value.OutType = anyTypeCGo(value.AnyType)
	value.outDecl.Linef("var %s %s", value.OutName, value.OutType)

	switch {
	case array.FixedSize > 0:
		// Safe to do if this is a primitive AND we're not setting this inside a
		// calllback, since the callback will retain Go memory beyond its
		// lifetime which is bad.
		if !value.isTransferring() && inner.resolved.CanCast() {
			value.addImport("runtime")

			// We can directly use Go's array as a pointer, as long as we defer
			// properly.
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s))",
				value.OutName, value.OutType, value.InName,
			)

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
		length := conv.convertParam(*array.Length)
		if length == nil {
			return false
		}

		// Length has no input, as it's from the slice.
		value.outDecl.Linef("var %s %s", length.OutName, length.OutType)
		value.p.Linef("%s = %s(len(%s))", length.OutName, length.OutType, value.InName)

		// Use the backing array with the appropriate transfer-ownership rule
		// for primitive types; see type_c_go.go.
		if !value.isTransferring() && inner.resolved.CanCast() {
			value.addImport("unsafe")
			value.addImport("runtime")

			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s[0]))",
				value.OutName, value.OutType, value.InName,
			)

			return true
		}

		value.addImport(importInternal("ptr"))
		value.addImport("unsafe")

		value.p.Linef(
			"%s = (%s)(%s)",
			value.OutName, value.OutType, inner.malloc(value.InName, false),
		)
		if !value.isTransferring() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)
			value.p.EmptyLine()
		}

		value.p.Descend()

		value.p.Linef("var out []%s", inner.OutType)
		value.p.Linef(goSliceFromPtr("out", value.OutName, fmt.Sprintf("len(%s)", value.InName)))
		value.p.EmptyLine()

		value.p.Linef("for i := range %s {", value.InName)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	case array.Name == "GLib.Array":
		value.addImport(importInternal("ptr"))
		value.addImport("unsafe")

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

		value.p.Linef("var out []%s", inner.OutType)
		value.p.Linef(
			goSliceFromPtr("out", value.OutName+".data", fmt.Sprintf("len(%s)", value.InName)),
		)
		value.p.EmptyLine()

		value.p.Linef("for i := range %s {", value.InName)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	case array.IsZeroTerminated():
		value.addImport(importInternal("ptr"))
		value.addImport("unsafe")

		value.p.Linef(
			"%s = (%s)(%s)",
			value.OutName, value.OutType, inner.malloc(value.InName, true))
		if !value.isTransferring() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)
			value.p.EmptyLine()
		}

		value.p.Descend()

		value.p.Linef("var out []%s", inner.OutType)
		value.p.Linef(goSliceFromPtr("dst", value.OutName, fmt.Sprintf("len(%s)", value.InName)))
		value.p.EmptyLine()

		value.p.Linef("for i := range %s {", value.InName)
		value.p.Linef(inner.Conversion)
		value.p.Linef("}")

		value.p.Ascend()
		return true

	default:
		conv.log(LogDebug, "Go->C weird array type", array.Type)
	}

	return false
}

func (conv *TypeConversionToC) typeConverter(value *ValueConverted) bool {
	for _, unsupported := range unsupportedCTypes {
		if unsupported == value.AnyType.Type.CType {
			return false
		}
	}

	if !value.resolveType(&conv.conversionTo, false) {
		return false
	}

	switch {
	case value.resolved.IsBuiltin("string"):
		value.p.Linef("%s = (%s)(C.CString(%s))", value.OutName, value.OutType, value.InName)
		// If we're not giving ownership this mallocated string, then we
		// can free it once done.
		if !value.isTransferring() {
			value.addImport("unsafe")
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)
		}
		return true

	case value.resolved.IsBuiltin("bool"):
		value.p.Linef("if %s { %s = %s(1) }", value.InName, value.OutName, value.OutType)
		return true

	case value.resolved.IsBuiltin("error"):
		value.addImport(importInternal("gerror"))
		value.p.Linef("%s = (*C.GError)(gerror.New(unsafe.Pointer(%s)))", value.OutName, value.InName)
		if !value.isTransferring() {
			value.p.Linef("defer C.g_error_free(%s)", value.OutName)
		}
		return true

	case value.resolved.IsPrimitive():
		value.p.Linef("%s = %s(%s)", value.OutName, value.OutType, value.InName)
		return true
	}

	switch ensureNamespace(conv.ng.current, value.AnyType.Type.Name) {
	case "gpointer":
		value.addImport("unsafe")
		value.addImport(importInternal("box"))
		value.p.Linef(
			"%s = %s(box.Assign(unsafe.Pointer(%s)))",
			value.OutName, value.OutType, value.InName,
		)
		return true

	case "GObject.Type", "GType":
		value.NeedsGLibObject = true
		// Just a primitive.
		value.p.Linef("%s = C.GType(%s)", value.OutName, value.InName)
		return true

	case "GObject.Value":
		value.NeedsGLibObject = true
		// https://pkg.go.dev/github.com/gotk3/gotk3/glib#Type
		value.p.Linef("%s = (*C.GValue)(%s.GValue)", value.OutName, value.InName)
		return true

	case "GObject.Object":
		value.addImport("unsafe")
		value.p.Linef("%s = (*C.GObject)(unsafe.Pointer(%s.Native()))", value.OutName, value.InName)
		return true

	case "GObject.InitiallyUnowned":
		value.addImport("unsafe")
		value.p.Linef(
			"%s = (*C.GInitiallyUnowned)(unsafe.Pointer(%s.Native()))",
			value.OutName, value.InName)
		return true
	}

	// Pretend that ignored types don't exist.
	if conv.ng.mustIgnore(value.AnyType.Type.Name, value.AnyType.Type.CType) {
		return false
	}

	if value.resolved.Extern == nil {
		return false
	}

	result := value.resolved.Extern.Result

	switch {
	case result.Enum != nil, result.Bitfield != nil:
		value.p.Linef("%s = (%s)(%s)", value.OutName, value.OutType, value.InName)
		return true

	case result.Class != nil, result.Record != nil, result.Interface != nil:
		value.p.Linef(
			"%s = (%s)(unsafe.Pointer(%s.Native()))",
			value.OutName, value.OutType, value.InName,
		)
		return true

	case result.Callback != nil:
		exportedName, _ := result.Info()
		exportedName = PascalToGo(exportedName)

		// Callbacks must have the closure attribute to store the closure
		// pointer.
		if value.Closure == nil {
			conv.log(LogDebug, "Go->C callback", exportedName, "since missing closure")
			return false
		}

		value.addImport(importInternal("box"))
		value.addCallback(result.Callback)

		// Return the constant function here. The function will dynamically load
		// the user_data, which will match with the "gpointer" case above.
		//
		// As for the pointer to byte array cast, see
		// https://github.com/golang/go/issues/19835.
		value.p.Linef("%s = (*[0]byte)(C.%s%s)", value.OutName, callbackPrefix, exportedName)

		closure := conv.convertParam(*value.Closure)
		if closure == nil {
			return false
		}

		value.outDecl.Linef("var %s %s", closure.OutName, closure.OutType)
		value.p.Linef("%s = %s(box.Assign(%s))", closure.OutName, closure.OutType, value.InName)

		if value.Destroy != nil {
			value.CallbackDelete = true

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

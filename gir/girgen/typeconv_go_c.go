package girgen

import (
	"fmt"
	"log"
	"strings"
)

// Go to C type conversions.

// TODO: is there a reason GoCConverter just doesn't take in ParameterAttr?

// GoValueProp describes a Go variable.
type GoValueProp struct {
	ValueProp
}

// inner is used only for arrays.
func (prop *GoValueProp) inner(in, out string) *GoValueProp {
	return &GoValueProp{
		ValueProp: prop.ValueProp.inner(in, out),
	}
}

// TypeConversionToC describes a conversion of one or more Go values to C using
// CGo.
type TypeConversionToC struct {
	values  []GoValueProp
	ignores map[int]struct{}
	parent  string // for debugging

	conversionTo
}

// GoCConverter returns a new converter that converts a value from Go to C.
func (fg *FileGenerator) GoCConverter(parent string, values []GoValueProp) *TypeConversionToC {
	ignores := make(map[int]struct{}, 10)
	for _, value := range values {
		value.loadIgnore(ignores)
	}

	return &TypeConversionToC{
		values:       values,
		ignores:      ignores,
		parent:       parent,
		conversionTo: newConversionTo(fg, parent),
	}
}

// ConvertAll converts all values.
func (conv *TypeConversionToC) ConvertAll() []TypeConverted {
	return ConvertAllValues(conv, len(conv.values))
}

// Convert converts the value at the given index.
func (conv *TypeConversionToC) Convert(i int) *TypeConverted {
	// Bound check.
	if i >= len(conv.values) {
		return nil
	}

	value := conv.values[i]
	value.init()

	// Ignored values are manually obtained in the conversion process, so we
	// don't convert them here.
	if value.ParameterIndex != nil {
		_, ignore := conv.ignores[*value.ParameterIndex]
		if ignore {
			return &TypeConverted{}
		}
	}

	// Reset the state when done. The returns all copy the internal state, so
	// we're fine.
	defer conv.reset()

	conv.gocConverter(&value)
	if conv.failed {
		conv.logFail(LogDebug, "Go->C cannot convert type", anyTypeC(conv.values[i].AnyType))
		return nil
	}

	if value.InType == "" || value.OutType == "" {
		log.Panicln("missing CGoType or GoType for value", conv.parent, i)
	}

	c := value.TypeConverted
	c.finalize()
	c.ConversionSideEffects = conv.sides

	return &c
}

func (conv *TypeConversionToC) valueAt(at int) *GoValueProp {
	for _, value := range conv.values {
		if value.ParameterIndex != nil && *value.ParameterIndex == at {
			value.init()
			return &value
		}
	}

	conv.logFail(LogError, "Go->C conversion arg not found at", at)

	prop := GoValueProp{ValueProp: errorValueProp}
	prop.init()

	return &prop
}

func (conv *TypeConversionToC) gocConverter(value *GoValueProp) {
	if value.ValueProp == errorValueProp {
		conv.fail()
		return
	}

	switch {
	case value.AnyType.Type != nil:
		conv.gocTypeConverter(value)
	case value.AnyType.Array != nil:
		conv.gocArrayConverter(value)
	default:
		conv.fail()
	}
}

func (conv *TypeConversionToC) gocArrayConverter(value *GoValueProp) {
	if value.AnyType.Array.Type == nil {
		if !value.AllowNone {
			conv.logFail(LogSkip, "nested array", value.AnyType.Array.CType)
		}
		return
	}

	array := value.AnyType.Array

	inner := value.inner(value.InName+"[i]", "out[i]")
	conv.gocConverter(inner)

	if conv.failed {
		return
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
			conv.sides.addImport("runtime")

			// We can directly use Go's array as a pointer, as long as we defer
			// properly.
			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s))",
				value.OutName, value.OutType, value.InName,
			)
			value.p.Linef("defer runtime.KeepAlive(&%s)", value.OutName)

			return
		}

		// Target fixed array, so we can directly set the data over. The memory
		// is ours, and allocation is handled by Go.
		value.p.Descend()

		value.p.Linef("out := (*[%d]%s)(unsafe.Pointer(%s))",
			array.FixedSize, inner.OutType, value.OutName)

		value.p.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		value.p.Linef("  " + inner.p.String())
		value.p.Linef("}")

		value.p.Ascend()

	case array.Length != nil:
		length := conv.valueAt(*array.Length)
		conv.gocConverter(length)
		// Length has no input, as it's from the slice.

		value.outDecl.Linef("var %s %s", length.OutName, length.OutType)
		value.p.Linef("%s = %s(len(%s))", length.OutName, length.OutType, value.InName)

		// Use the backing array with the appropriate transfer-ownership rule
		// for primitive types; see type_c_go.go.
		if !value.isTransferring() && inner.resolved.CanCast() {
			conv.sides.addImport("unsafe")
			conv.sides.addImport("runtime")

			value.p.Linef(
				"%s = (%s)(unsafe.Pointer(&%s[0]))",
				value.OutName, value.OutType, value.InName,
			)
			value.p.Linef("defer runtime.KeepAlive(%s)", value.OutName)

			return
		}

		conv.sides.addImport(importInternal("ptr"))
		conv.sides.addImport("unsafe")

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
		value.p.Linef(inner.p.String())
		value.p.Linef("}")

		value.p.Ascend()

	case array.Name == "GLib.Array":
		conv.sides.addImport(importInternal("ptr"))
		conv.sides.addImport("unsafe")

		// https://developer.gnome.org/glib/stable/glib-Arrays.html#g-array-sized-new
		value.p.Linef(
			"%s = C.g_array_sized_new(%t, false, C.guint(%s), C.guint(len(%s)))",
			value.OutName, array.IsZeroTerminated(), csizeof(inner.resolved), value.InName)
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
		value.p.Linef(inner.p.String())
		value.p.Linef("}")

		value.p.Ascend()

	case array.IsZeroTerminated():
		conv.sides.addImport(importInternal("ptr"))
		conv.sides.addImport("unsafe")

		value.p.Linef("%s = (%s)(%s)", value.OutName, value.OutType, inner.malloc(value.InName, true))
		if !value.isTransferring() {
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)
			value.p.EmptyLine()
		}

		value.p.Descend()

		value.p.Linef("var out []%s", inner.OutType)
		value.p.Linef(goSliceFromPtr("dst", value.OutName, fmt.Sprintf("len(%s)", value.InName)))
		value.p.EmptyLine()

		value.p.Linef("for i := range %s {", value.InName)
		value.p.Linef(inner.p.String())
		value.p.Linef("}")

		value.p.Ascend()

	default:
		conv.logFail(LogSkip, "weird array type to C")
	}
}

func csizeof(resolved *ResolvedType) string {
	if !strings.Contains(resolved.CType, "*") {
		return "C.sizeof_" + resolved.CType
	}

	// Size of an integer is the same as uintptr.
	return "unsafe.Sizeof(int(0))"
}

func (conv *TypeConversionToC) gocTypeConverter(value *GoValueProp) {
	for _, unsupported := range unsupportedCTypes {
		if unsupported == value.AnyType.Type.CType {
			conv.fail()
			return
		}
	}

	if !value.resolveType(&conv.conversionTo, false) {
		conv.fail()
		return
	}

	switch {
	case value.resolved.IsBuiltin("string"):
		value.p.Linef("%s = (%s)(C.CString(%s))", value.OutName, value.OutType, value.InName)
		// If we're not giving ownership this mallocated string, then we
		// can free it once done.
		if !value.isTransferring() {
			conv.sides.addImport("unsafe")
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.OutName)
		}
		return

	case value.resolved.IsBuiltin("bool"):
		value.p.Linef("if %s { %s = %s(1) }", value.InName, value.OutName, value.OutType)
		return

	case value.resolved.IsBuiltin("error"):
		conv.sides.addImport(importInternal("gerror"))
		value.p.Linef("%s = (*C.GError)(gerror.New(unsafe.Pointer(%s)))", value.OutName, value.InName)
		if !value.isTransferring() {
			value.p.Linef("defer C.g_error_free(%s)", value.OutName)
		}
		return

	case value.resolved.IsPrimitive():
		value.p.Linef("%s = %s(%s)", value.OutName, value.OutType, value.InName)
		return
	}

	switch ensureNamespace(conv.ng.current, value.AnyType.Type.Name) {
	case "gpointer":
		conv.sides.addImport(importInternal("box"))

		value.p.Linef(
			"%s = %s(box.Assign(unsafe.Pointer(%s)))",
			value.OutName, value.OutType, value.InName,
		)
		return

	case "GObject.Type", "GType":
		conv.sides.NeedsGLibObject = true

		// Just a primitive.
		value.p.Linef("%s = C.GType(%s)", value.OutName, value.InName)
		return

	case "GObject.Value":
		conv.sides.NeedsGLibObject = true

		// https://pkg.go.dev/github.com/gotk3/gotk3/glib#Type
		value.p.Linef("%s = (*C.GValue)(%s.GValue)", value.OutName, value.InName)
		return

	case "GObject.Object":
		value.p.Linef("%s = (*C.GObject)(unsafe.Pointer(%s.Native()))", value.OutName, value.InName)
		return

	case "GObject.InitiallyUnowned":
		value.p.Linef(
			"%s = (*C.GInitiallyUnowned)(unsafe.Pointer(%s.Native()))",
			value.OutName, value.InName)
		return
	}

	// Pretend that ignored types don't exist.
	if conv.ng.mustIgnore(value.AnyType.Type.Name, value.AnyType.Type.CType) {
		conv.fail()
		return
	}

	if value.resolved.Extern == nil {
		conv.fail()
		return
	}

	result := value.resolved.Extern.Result

	switch {
	case result.Enum != nil, result.Bitfield != nil:
		value.p.Linef("%s = (%s)(%s)", value.OutName, value.OutType, value.InName)

	case result.Class != nil, result.Record != nil, result.Interface != nil:
		value.p.Linef(
			"%s = (%s)(unsafe.Pointer(%s.Native()))",
			value.OutName, value.OutType, value.InName,
		)

	case result.Callback != nil:
		exportedName, _ := result.Info()
		exportedName = PascalToGo(exportedName)

		// Callbacks must have the closure attribute to store the closure
		// pointer.
		if value.Closure == nil {
			conv.logFail(LogSkip, "Go->C callback", exportedName, "since missing closure")
			return
		}

		conv.sides.addImport(importInternal("box"))
		conv.sides.addCallback(result.Callback)

		// Return the constant function here. The function will dynamically load
		// the user_data, which will match with the "gpointer" case above.
		//
		// As for the pointer to byte array cast, see
		// https://github.com/golang/go/issues/19835.
		value.p.Linef("%s = (*[0]byte)(C.%s%s)", value.OutName, callbackPrefix, exportedName)

		closure := conv.valueAt(*value.Closure)
		conv.gocConverter(closure)

		value.outDecl.Linef("var %s %s", closure.OutName, closure.OutType)
		value.p.Linef("%s = %s(box.Assign(%s))", closure.OutName, closure.OutType, value.InName)

		if value.Destroy != nil {
			conv.sides.CallbackDelete = true

			destroy := conv.valueAt(*value.Destroy)
			conv.gocConverter(destroy)

			value.outDecl.Linef("var %s %s", destroy.OutName, destroy.OutType)
			value.p.Linef(
				"%s = (%s)((*[0]byte)(C.callbackDelete))",
				destroy.OutName, destroy.OutType,
			)
		}

	case value.AllowNone:
		value.p.Linef("var %s %s // unsupported", value.OutName, value.OutType)

	default:
		conv.fail()
	}
}

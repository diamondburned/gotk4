package girgen

import (
	"fmt"
	"log"

	"github.com/diamondburned/gotk4/internal/pen"
)

// C to Go type conversions.

// CTypeProp describes a C variable.
type CValueProp struct {
	ValueProp

	// BoxCast is an optional Go type that the boxed value should be casted to,
	// but only if the Type is a gpointer. This is only useful to convert from C
	// to Go.
	// BoxCast string
}

// inner is used only for arrays.
func (prop *CValueProp) inner(in, out string) *CValueProp {
	return &CValueProp{
		ValueProp: prop.ValueProp.inner(in, out),
	}
}

// TypeConversionToGo describes the conversion context, that is all the values
// that may or may not be converted for the conversion routine to use.
type TypeConversionToGo struct {
	values  []CValueProp
	ignores map[int]struct{}
	parent  string

	conversionTo
}

// CGoConverter returns a new converter that converts a value from C to Go.
func (fg *FileGenerator) CGoConverter(parent string, values []CValueProp) *TypeConversionToGo {
	ignores := make(map[int]struct{}, 10)
	for _, value := range values {
		value.loadIgnore(ignores)
	}

	return &TypeConversionToGo{
		values:  values,
		ignores: ignores,
		parent:  parent,

		conversionTo: newConversionTo(fg, parent),
	}
}

// WriteAll writes all conversionsto the given sections.
func (conv *TypeConversionToGo) WriteAll(in, out, con *pen.BlockSection) bool {
	// Get the FileGenerator out of nowhere.
	fg := conv.logger.(*FileGenerator)

	for i := 0; i < len(conv.values); i++ {
		converted := conv.Convert(i)
		if converted == nil {
			conv.logFail(LogDebug, "C->Go cannot convert type", anyTypeC(conv.values[i].Type))
			return false
		}

		converted.Apply(fg)
		converted.WriteAll(in, out, con)
	}

	return true
}

// Convert converts the value at the given index.
func (conv *TypeConversionToGo) Convert(i int) *TypeConverted {
	// Bound check.
	if i >= len(conv.values) {
		return nil
	}

	// Make a shallow copy of the value.
	value := conv.values[i]
	value.initState()

	// Ignored values are manually obtained in the conversion process, so we
	// don't convert them here. Zero values are not important.
	if value.ParameterIndex != nil {
		_, ignore := conv.ignores[*value.ParameterIndex]
		if ignore {
			return &TypeConverted{}
		}
	}

	// Reset the state when done. The returns all copy the internal state, so
	// we're fine.
	defer conv.reset()

	conv.cgoConverter(&value)
	if conv.failed {
		return nil
	}

	if value.InType == "" || value.OutType == "" {
		log.Panicln("missing CGoType or GoType for value", conv.parent, i)
	}

	return &TypeConverted{
		ValueProp:             &value.ValueProp,
		InDeclare:             value.inDecl.String(),
		OutDeclare:            value.outDecl.String(),
		Conversion:            value.p.String(),
		ConversionSideEffects: conv.sides,
	}
}

func (conv *TypeConversionToGo) reset() {
	conv.conversionTo.reset()
}

// valueAt returns a copy of the value at the given parameter index.
func (conv *TypeConversionToGo) valueAt(at int) *CValueProp {
	for i, value := range conv.values {
		if value.ParameterIndex != nil && *value.ParameterIndex == at {
			value := conv.values[i]
			value.initState()
			return &value
		}
	}

	conv.logFail(LogError, "C->Go conversion arg not found at", at)
	return &CValueProp{ValueProp: errorValueProp}
}

func (conv *TypeConversionToGo) cgoConverter(value *CValueProp) {
	switch {
	case value.Type.Array != nil:
		conv.cgoArrayConverter(value)
	case value.Type.Type != nil:
		conv.cgoTypeConverter(value)
	default:
		conv.fail()
	}
}

func (conv *TypeConversionToGo) cgoArrayConverter(value *CValueProp) {
	if value.Type.Array.Type == nil {
		if !value.AllowNone {
			conv.logFail(LogSkip, "nested array", value.Type.Array.CType)
		}
		return
	}

	array := value.Type.Array

	// All generators must declare src.
	inner := value.inner("src", value.Out+"[i]")
	conv.cgoConverter(inner)

	if conv.failed {
		return
	}

	if array.FixedSize > 0 {
		value.OutType = fmt.Sprintf("[%d]%s", array.FixedSize, inner.OutType)
		value.outDecl.Linef("var %s %s", value.Out, value.OutType)

		value.InType = anyTypeCGo(value.Type)
		value.inDecl.Linef("var %s [%d]%s", value.In, array.FixedSize, inner.InType)
	} else {
		value.OutType = fmt.Sprintf("[]%s", inner.OutType)
		value.outDecl.Linef("var %s %s", value.Out, value.OutType)

		value.InType = anyTypeCGo(value.Type)
		value.inDecl.Linef("var %s %s", value.In, value.InType)
	}

	switch {
	case array.FixedSize > 0:
		// Detect if the underlying is a compatible Go primitive type that isn't
		// a string. If it is, then we can directly cast a fixed-size array.
		if inner.resolved.IsPrimitive() {
			value.p.Linef("%s = (%s)(%s)", value.Out, value.OutType, value.In)
			return
		}

		value.p.Descend()

		// Direct cast is not possible; make a temporary array with the CGo type
		// so we can loop over it easily.
		value.p.Linef("tmp := *(*%s)(unsafe.Pointer(&%s))", value.OutType, value.In)
		value.p.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		value.p.Linef("  src := tmp[i]")
		value.p.Linef("  " + inner.p.String())
		value.p.Linef("}")

		value.p.Ascend()

	case array.Length != nil:
		conv.sides.addImport("unsafe")
		conv.sides.addImport(importInternal("ptr"))

		length := conv.valueAt(*array.Length)
		value.inDecl.Linef("var %s %s", length.In, anyTypeCGo(length.Type))
		// Length has no outDecl.

		// If we're owning the new data, then we will directly use the backing
		// array, but we can only do this if the underlying type is a primitive,
		// since those have equivalent Go representations. Any other types will
		// have to be copied or otherwise converted somehow.
		//
		// TODO: record conversion should handle ownership: if
		// transfer-ownership is none, then the native pointer should probably
		// not be freed.
		if value.isTransferring() && inner.resolved.IsPrimitive() {
			conv.sides.addImport("runtime")

			value.p.Linef(goSliceFromPtr(value.Out, value.In, length.In))

			// See: https://golang.org/misc/cgo/gmp/gmp.go?s=3086:3757#L87
			value.p.Linef("runtime.SetFinalizer(&%s, func(v *%s) {", value.Out, value.OutType)
			value.p.Linef("  C.free(ptr.Slice(unsafe.Pointer(v)))")
			value.p.Linef("})")

			return
		}

		value.p.Linef("%s = make(%s, %s)", value.Out, value.OutType, length.In)
		value.p.Linef("for i := 0; i < uintptr(%s); i++ {", length.In)
		value.p.Linef("  src := (%s)(ptr.Add(unsafe.Pointer(%s), i))", inner.InType, value.In)
		value.p.Linef("  " + inner.p.String())
		value.p.Linef("}")

	case array.Name == "GLib.Array": // treat as Go array
		conv.sides.addImport("unsafe")
		conv.sides.addImport(importInternal("ptr"))

		value.p.Descend()

		value.p.Linef("var len uintptr")
		value.p.Linef("p := C.g_array_steal(&%s, (*C.gsize)(&len))", value.In)

		value.p.Linef("%s = make(%s, len)", value.Out, value.OutType)
		value.p.Linef("for i := 0; i < len; i++ {", value.In)
		value.p.Linef("  src := (%s)(ptr.Add(unsafe.Pointer(p), i))", value.In)
		value.p.Linef("  " + inner.p.String())
		value.p.Linef("}")

		value.p.Ascend()

	case array.IsZeroTerminated():
		conv.sides.addImport("unsafe")
		conv.sides.addImport(importInternal("ptr"))

		value.p.Descend()

		// Get the size of the type so we know how much to increment when
		// scanning.
		sizeof := csizeof(inner.resolved)

		// Scan for the length.
		value.p.Linef("var length int")
		value.p.Linef("for p := %s; *p != 0; p = (%s)(ptr.Add(unsafe.Pointer(p), %s)) {",
			value.In, value.InType, sizeof)
		value.p.Linef("  length++")
		value.p.Linef("  if length < 0 { panic(`length overflow`) }")
		value.p.Linef("}")
		value.p.EmptyLine()

		// Preallocate the slice.
		value.p.Linef("%s = make(%s, length)", value.Out, value.OutType)
		value.p.Linef("for i := uintptr(0); i < uintptr(length); i += %s {", sizeof)
		value.p.Linef("  src := (%s)(ptr.Add(unsafe.Pointer(%s), i))", inner.InType, value.In)
		value.p.Linef("  " + inner.p.String())
		value.p.Linef("}")

		value.p.Ascend()

	default:
		conv.logFail(LogSkip, "weird array type to Go")
	}
}

func (conv *TypeConversionToGo) cgoTypeConverter(value *CValueProp) {
	for _, unsupported := range unsupportedCTypes {
		if unsupported == value.Type.Type.Name {
			conv.fail()
			return
		}
	}

	if !value.resolveType(&conv.conversionTo, true) {
		conv.fail()
		return
	}

	switch {
	case value.resolved.IsBuiltin("string"):
		value.p.Linef("%s = C.GoString(%s)", value.Out, value.In)
		// Only free this if C is transferring ownership to us.
		if value.isTransferring() {
			conv.sides.addImport("unsafe")
			value.p.Linef("defer C.free(unsafe.Pointer(%s))", value.In)
		}
		return

	case value.resolved.IsBuiltin("bool"):
		conv.sides.NeedsStdBool = true
		value.p.Linef("%s = C.bool(%s) != C.false", value.Out, value.In)
		return

	case value.resolved.IsBuiltin("error"):
		conv.sides.addImport(importInternal("gerror"))
		value.p.Linef("%s = gerror.Take(unsafe.Pointer(%s))", value.Out, value.In)
		return

	case value.resolved.IsPrimitive():
		value.p.Linef("%s = %s(%s)", value.Out, value.InType, value.In)
		return
	}

	// Resolve special-case GLib types.
	switch ensureNamespace(conv.ng.current, value.Type.Type.Name) {
	case "gpointer":
		conv.sides.addImport(importInternal("box"))

		value.p.Linef("%s = box.Get(uintptr(%s))", value.Out, value.In)
		return

	case "GObject.Object", "GObject.InitiallyUnowned":
		conv.sides.addImport("unsafe")
		conv.sides.addImport(importInternal("gextras"))
		conv.sides.addGLibImport()

		value.p.Line(value.cgoSetObject("gextras.Objector"))
		return

	case "GObject.Type", "GType":
		conv.sides.addGLibImport()

		value.p.Linef("%s = externglib.Type(%s)", value.Out, value.In)
		return

	case "GObject.Value":
		conv.sides.addGLibImport()
		conv.sides.addImport("unsafe")

		value.p.Linef("%s = externglib.ValueFromNative(unsafe.Pointer(%s))", value.Out, value.In)
		// Set this to be freed if we have the ownership now.
		if value.isTransferring() {
			// https://pkg.go.dev/github.com/gotk3/gotk3/glib?utm_source=godoc#Value
			value.p.Linef("runtime.SetFinalizer(%s, func(v *externglib.Value) {", value.Out)
			value.p.Linef("  C.g_value_unset((*C.GValue)(v.GValue))")
			value.p.Linef("})")
		}
		return
	}

	// Pretend that ignored types don't exist.
	if conv.ng.mustIgnore(value.Type.Type.Name, value.Type.Type.CType) {
		conv.fail()
		return
	}

	// TODO: function
	// TODO: union
	// TODO: callback

	// TODO: callbacks and functions are handled differently. Unsure if they're
	// doable.
	// TODO: handle unions.

	if value.resolved.Extern == nil {
		conv.fail()
		return
	}

	result := value.resolved.Extern.Result

	switch {
	case result.Enum != nil, result.Bitfield != nil:
		// Resolve castable number types.
		value.p.Linef("%s = %s(%s)", value.Out, value.OutType, value.In)

	case result.Class != nil, result.Interface != nil:
		conv.sides.addImport("unsafe")
		conv.sides.addImport(importInternal("gextras"))

		value.p.Line(value.cgoSetObject(value.OutType))

	case result.Record != nil:
		// We should only use the concrete wrapper for the record, since the
		// returned type is concretely known here.
		wrapName := value.resolved.WrapName(value.needsNamespace)
		valueIn := value.In

		value.p.Linef("%s = %s(unsafe.Pointer(%s))", value.Out, wrapName, valueIn)

		if value.isTransferring() {
			conv.sides.addImport("runtime")

			value.p.Linef("runtime.SetFinalizer(%s, func(v %s) {", value.Out, value.OutType)
			value.p.Linef("  C.free(unsafe.Pointer(v.Native()))")
			value.p.Linef("})")
		}

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
		conv.fail()

	case value.AllowNone:
		value.outDecl.Linef("var %s %s // unsupported", value.Out, value.OutType)
	default:
		conv.fail()
	}
}

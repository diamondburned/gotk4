package girgen

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

// C to Go type conversions.

// TypeConversionToGo describes a conversion of one or more C values to Go.
type TypeConversionToGo struct {
	Values []CValueProp
	// Parent is used for debugging.
	Parent string
}

// CTypeProp describes a C variable.
type CValueProp struct {
	ValueProp

	// BoxCast is an optional Go type that the boxed value should be casted to,
	// but only if the Type is a gpointer. This is only useful to convert from C
	// to Go.
	BoxCast string

	// OutputParam makes the conversion treat this value like an output
	// parameter. The declarations will be written to the output's Preamble.
	OutputParam bool
}

// inner is used only for arrays.
func (prop *CValueProp) inner(in, out string) *CValueProp {
	return &CValueProp{
		ValueProp:   prop.ValueProp.inner(in, out),
		OutputParam: false,
	}
}

type conversionToGo struct {
	conversionTo
	pre     *pen.Paper // ONLY USE FOR OutputParam.
	convert *TypeConversionToGo
}

func (conv *conversionToGo) valueAt(i int) *CValueProp {
	if i < len(conv.convert.Values) {
		return &conv.convert.Values[i]
	}

	conv.logFail(LogError, "C->Go callback out-of-bound closure arg", i)
	return &CValueProp{ValueProp: errorValueProp}
}

type TypeConvertedToGo struct {
	// Preamble contains `values[i].In' declarations for OutputParams.
	Preamble   string
	Conversion string
	ConversionSideEffects
}

// CGoConverter returns Go code that is the conversion from the given C value
// type to its respective Go value type. An empty string returned is invalid.
//
// The given argPrefix is used to get the nth parameter by concatenating the
// prefix with the index number. This is used for length parameters.
func (fg *FileGenerator) CGoConverter(conv TypeConversionToGo) *TypeConvertedToGo {
	state := conversionToGo{
		conversionTo: newConversionTo(fg, conv.Parent),
		pre:          pen.NewPaperSize(1024), // 1KB
		convert:      &conv,
	}

	ignores := make([]bool, len(conv.Values))
	returns := make([]string, len(conv.Values))

	for _, value := range conv.Values {
		value.loadIgnore(&ignores)
	}

	for i := range conv.Values {
		// Ignored values are manually obtained in the conversion process, so we
		// don't convert them here.
		if ignores[i] {
			continue
		}

		returns[i] = conv.Values[i].Out

		state.cgoConverter(&conv.Values[i])
		state.p.EmptyLine()

		if state.failed {
			return nil
		}
	}

	return &TypeConvertedToGo{
		Preamble:              state.pre.String(),
		Conversion:            state.p.String(),
		ConversionSideEffects: state.sides,
	}
}

func (conv *conversionToGo) cgoConverter(value *CValueProp) {
	switch {
	case value.Type.Array != nil:
		conv.cgoArrayConverter(value)
	case value.Type.Type != nil:
		conv.cgoTypeConverter(value)
	default:
		conv.fail()
	}
}

func (conv *conversionToGo) cgoArrayConverter(value *CValueProp) {
	if value.Type.Array.Type == nil {
		if !value.AllowNone {
			conv.logFail(LogSkip, "nested array", value.Type.Array.CType)
		}
		return
	}

	array := value.Type.Array

	innerResolved := conv.ng.ResolveType(*array.AnyType.Type)
	if innerResolved == nil {
		conv.fail()
		return
	}

	innerGoType := GoPublicType(conv.ng, innerResolved)
	innerCGoType := innerResolved.CGoType()

	if value.OutputParam {
		// Non-fixed arrays are allocated on the C side, so we only need a
		// pointer.
		cgoType := anyTypeCGo(value.Type)
		conv.pre.Linef("var %s %s", value.In, cgoType)
	}

	if value.OutputParam && array.FixedSize == 0 {
		// Everything that's not FixedSize have this segment in common.
		conv.pre.Linef("var %s %s", value.In, anyTypeCGo(value.Type))
	}

	switch {
	case array.FixedSize > 0:
		if value.OutputParam {
			// Fixed size arrays are preallocated on the stack, so we have to
			// pass in the pointer to the array.
			conv.pre.Linef("var %s_array [%d]%s", value.In, array.FixedSize, innerCGoType)
			conv.pre.Linef("%s := &%s_array[0]", value.In)
		}

		// Detect if the underlying is a compatible Go primitive type that isn't
		// a string. If it is, then we can directly cast a fixed-size array.
		if !innerResolved.IsPrimitive() {
			conv.p.Linef("%s := ([%d]%s)(%s)", value.Out, array.FixedSize, innerGoType, value.In)
			return
		}

		conv.p.Descend()
		// Direct cast is not possible; make a temporary array with the CGo type
		// so we can loop over it.
		conv.p.Linef("tmp := ([%d]%s)(%s)", value.Out, array.FixedSize, innerCGoType, value.In)

		// TODO: nested array support
		conv.p.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		conv.cgoConverter(value.inner("tmp[i]", value.In+"[i]"))
		conv.p.Linef("}")

		conv.p.Ascend()

	case array.Length != nil:
		lengthArg := conv.valueAt(*array.Length)
		if value.OutputParam {
			conv.pre.Linef("var %s %s", lengthArg.In, anyTypeCGo(lengthArg.Type))
		}

		conv.sides.addImport("unsafe")

		// If we're owning the new data, then we will directly use the backing
		// array, but we can only do this if the underlying type is a primitive,
		// since those have equivalent Go representations. Any other types will
		// have to be copied or otherwise converted somehow.
		//
		// TODO: record conversion should handle ownership: if
		// transfer-ownership is none, then the native pointer should probably
		// not be freed.
		if value.isTransferring() && innerResolved.IsPrimitive() {
			conv.sides.addImport("runtime")
			conv.sides.addImport("reflect")

			conv.p.Linef("var %s []%s", value.Out, innerGoType)
			conv.p.Linef(goSliceFromPtr(value.Out, value.In, lengthArg.In))

			// See: https://golang.org/misc/cgo/gmp/gmp.go?s=3086:3757#L87
			conv.p.Linef("runtime.SetFinalizer(&%s, func(v *[]%s) {", value.Out, innerGoType)
			conv.p.Linef("  sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(v))")
			conv.p.Linef("  C.free(unsafe.Pointer(slideHeader.Data))")
			conv.p.Linef("})")

			return
		}

		conv.p.Linef("%s := make([]%s, %s)", value.Out, innerGoType, lengthArg.In)
		conv.p.Linef("for i := 0; i < uintptr(%s); i++ {", lengthArg)
		conv.p.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))", innerCGoType)
		conv.cgoConverter(value.inner("src", value.Out+"[i]"))
		conv.p.Linef("}")

	case array.Name == "GLib.Array": // treat as Go array
		conv.sides.addImport("unsafe")

		conv.p.Linef("var %s []%s", value.Out, innerGoType)

		conv.p.Descend()

		conv.p.Linef("var len uintptr")
		conv.p.Linef("p := C.g_array_steal(&%s, (*C.gsize)(&len))", value.In)

		conv.p.Linef("%s = make([]%s, len)", value.Out, innerGoType)
		conv.p.Linef("for i := 0; i < len; i++ {", value.In)
		conv.p.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + i))", innerCGoType)
		conv.cgoConverter(value.inner("src", value.Out+"[i]"))
		conv.p.Linef("}")

		conv.p.Ascend()

	case array.IsZeroTerminated():
		conv.sides.addImport("unsafe")

		conv.p.Linef("var %s []%s", value.Out, innerGoType)

		conv.p.Descend()

		// Scan for the length.
		conv.p.Linef("var length uint")
		conv.p.Linef("for p := unsafe.Pointer(%s); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {", value.In)
		conv.p.Linef("  length++")
		conv.p.Linef("}")

		conv.p.EmptyLine()

		// Preallocate the slice.
		conv.p.Linef("%s = make([]%s, length)", value.Out, innerGoType)
		conv.p.Linef("for i := 0; i < length; i++ {")
		conv.p.Linef("  src := (%s)(unsafe.Pointer(uintptr(unsafe.Pointer(%s)) + i))", innerCGoType, value.In)
		conv.cgoConverter(value.inner("src", value.Out+"[i]"))
		conv.p.Linef("}")

		conv.p.Ascend()

	default:
		conv.logFail(LogSkip, "weird array type to Go")
		return
	}
}

func (conv *conversionToGo) cgoTypeConverter(value *CValueProp) {
	for _, unsupported := range unsupportedCTypes {
		if unsupported == value.Type.Type.Name {
			conv.fail()
			return
		}
	}

	if value.OutputParam {
		// Declare the variable, even when it's as the original alias type.
		// We'll wrap this variable with our own conversions.
		cgoType := anyTypeCGo(value.Type)
		conv.pre.Linef("var %s %s", value.In, cgoType)
	}

	var (
		result   *gir.TypeFindResult
		resolved *ResolvedType

		cgoType     string
		unwrapStack []string
	)

	girType := value.Type.Type
	for {
		result = conv.ng.FindType(girType.Name)
		resolved = conv.ng.ResolveType(*girType)

		if result == nil || result.Alias == nil {
			break
		}

		// Wrap the CGo input value.
		cgoType = resolved.CGoType()
		value.In = fmt.Sprintf("(%s)(%s)", cgoType, value.In)

		// Add the Go type to the unwrap stack.
		needsNsp := resolved.NeedsNamespace(conv.ng.current)
		unwrapStack = append(unwrapStack, resolved.PublicType(needsNsp))

		// Resolve again with the target alias type.
		girType = &result.Alias.Type
	}

	if resolved == nil {
		conv.fail()
		return
	}

	// TODO: find a way to construct the output wrapper. Easiest way is to
	// output to a tmp variable and convert back, but this would require putting
	// this inside a block.

	// Resolve primitive types.
	if prim, ok := girToBuiltin[girType.Name]; ok {
		// Edge cases.
		switch prim {
		case "string":
			conv.p.Linef("%s := C.GoString(%s)", value.Out, value.In)

			// Only free this if C is transferring ownership to us.
			if value.isTransferring() {
				conv.sides.addImport("unsafe")
				conv.p.Linef("defer C.free(unsafe.Pointer(%s))", value.In)
			}
		case "bool":
			conv.sides.NeedsStdBool = true
			conv.p.Linef("%s := C.bool(%s) != C.false", value.Out, value.In)
		default:
			conv.p.Linef("%s := %s(%s)", value.Out, cgoType, value.In)
		}

		return
	}

	// Resolve special-case GLib types.
	switch ensureNamespace(conv.ng.current, girType.Name) {
	case "gpointer":
		conv.sides.addImport("github.com/diamondburned/gotk4/internal/box")
		if value.BoxCast != "" && value.BoxCast != "interface{}" {
			conv.p.Linef("%s := box.Get(uintptr(%s)).(%s)", value.Out, value.In, value.BoxCast)
		} else {
			conv.p.Linef("%s := box.Get(uintptr(%s))", value.Out, value.In)
		}
		return

	case "GObject.Object", "GObject.InitiallyUnowned":
		conv.sides.addImport("unsafe")
		conv.sides.addImport("github.com/diamondburned/gotk4/internal/gextras")
		conv.p.Line(value.cgoCreateObject("gextras.Objector"))
		return

	case "GObject.Type", "GType":
		conv.sides.addGLibImport()
		conv.p.Linef("%s := externglib.Type(%s)", value.Out, value.In)
		return

	case "GObject.Value":
		conv.sides.addGLibImport()
		conv.sides.addImport("unsafe")
		conv.p.Linef("%s := externglib.ValueFromNative(unsafe.Pointer(%s))", value.Out, value.In)
		// Set this to be freed if we have the ownership now.
		if value.isTransferring() {
			conv.p.Linef("runtime.SetFinalizer(%s, func(v *externglib.Value) {", value.Out)
			conv.p.Linef("  C.g_value_unset((*C.GValue)(v.GValue))")
			conv.p.Linef("})")
		}
		return

	case "GObject.Error":
		return
	}

	// Pretend that ignored types don't exist.
	if conv.ng.mustIgnore(girType.Name, girType.CType) {
		conv.fail()
		return
	}

	if value.OutputParam && resolved.Ptr > 0 {
		resolved.Ptr--
	}

	// goName contains the pointer.
	goName := GoPublicType(conv.ng, resolved)

	// TODO: function
	// TODO: union
	// TODO: callback

	// TODO: callbacks and functions are handled differently. Unsure if they're
	// doable.
	// TODO: handle unions.

	switch {
	case result.Enum != nil, result.Bitfield != nil:
		// Resolve castable number types.
		conv.p.Linef("%s := %s(%s)", value.Out, goName, value.In)

	case result.Class != nil, result.Interface != nil:
		conv.sides.addImport("unsafe")
		conv.sides.addImport("github.com/diamondburned/gotk4/internal/gextras")
		conv.p.Line(value.cgoCreateObject(goName))

	case result.Record != nil:
		// We should only use the concrete wrapper for the record, since the
		// returned type is concretely known here.
		wrapName := resolved.WrapName(resolved.NeedsNamespace(conv.ng.current))
		valueIn := value.In

		if resolved.Ptr == 0 {
			wrapName = "*" + wrapName
			valueIn = "&" + valueIn
		}

		conv.p.Linef("%s := %s(unsafe.Pointer(%s))", value.Out, wrapName, valueIn)

		if !value.isTransferring() {
			return
		}

		// If ownership is being transferred to us on the Go side, then
		// we should free.
		conv.sides.addImport("runtime")

		arg := value.Out
		typ := goName

		if resolved.Ptr == 1 {
			arg = "&" + arg
			typ = "*" + typ
		}

		conv.p.Linef("runtime.SetFinalizer(%s, func(v %s) {", arg, typ)
		conv.p.Linef("  C.free(unsafe.Pointer(v.Native()))")
		conv.p.Linef("})")

	case value.AllowNone:
		conv.p.Linef("var %s %s // unsupported", value.Out, goName)

	default:
		conv.fail()
	}
}

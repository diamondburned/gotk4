package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

// Go to C type conversions.

// TODO: is there a reason GoCConverter just doesn't take in ParameterAttr?

// TypeConversionToC describes a conversion of one or more Go values to C using
// CGo.
type TypeConversionToC struct {
	Values []GoValueProp
	Parent string // for debugging
}

// GoValueProp describes a Go variable.
type GoValueProp struct {
	ValueProp
}

// inner is used only for arrays.
func (prop *GoValueProp) inner(in, out string) *GoValueProp {
	return &GoValueProp{ValueProp: prop.ValueProp.inner(in, out)}
}

type conversionToC struct {
	conversionTo
	convert *TypeConversionToC
}

func (conv *conversionToC) valueAt(i int) *GoValueProp {
	if i < len(conv.convert.Values) {
		return &conv.convert.Values[i]
	}

	conv.logFail(LogError, "Go->C callback out-of-bound closure arg", i)
	return &GoValueProp{ValueProp: errorValueProp}
}

// TypeConvertedToC is the result of the type conversion. It describes the
// function call arguments and the declarations/conversions.
//
// Note that the function parameters assume that the caller completely fills out
// the parameters such as Closure and Destroy, since those give information to
// fields not given to the Values slice in the input struct.
type TypeConvertedToC struct {
	Parameters []string
	Conversion string
	ConversionSideEffects
}

// GoCConverter generates code that converts from Go values to C.
func (fg *FileGenerator) GoCConverter(conv TypeConversionToC) *TypeConvertedToC {
	state := conversionToC{
		conversionTo: newConversionTo(fg, conv.Parent),
		convert:      &conv,
	}

	ignores := make([]bool, len(conv.Values))
	params := make([]string, len(conv.Values))

	for _, value := range conv.Values {
		value.loadIgnore(&ignores)
	}

	for i := range conv.Values {
		// Write the parameter fully regardless if we have a conversion or not.
		params[i] = conv.Values[i].Out

		if ignores[i] {
			continue
		}

		state.gocConverter(&conv.Values[i])
		state.p.EmptyLine()

		if state.failed {
			return nil
		}
	}

	return &TypeConvertedToC{
		Parameters:            params,
		Conversion:            state.p.String(),
		ConversionSideEffects: state.sides,
	}
}

func (conv *conversionToC) gocConverter(value *GoValueProp) {
	switch {
	case value.Type.Type != nil:
		conv.gocTypeConverter(value)
	case value.Type.Array != nil:
		conv.gocArrayConverter(value)
	default:
		conv.fail()
	}
}

func (conv *conversionToC) gocArrayConverter(value *GoValueProp) {
	if value.Type.Array.Type == nil {
		if !value.AllowNone {
			conv.logFail(LogSkip, "nested array", value.Type.Array.CType)
		}
		return
	}

	array := value.Type.Array

	// Use parent's ResolveType to not add imports.
	innerResolved := conv.ng.ResolveType(*array.Type)
	if innerResolved == nil {
		conv.fail()
		return
	}

	outerCGoType := anyTypeCGo(value.Type)
	innerCGoType := innerResolved.CGoType()

	switch {
	case array.FixedSize > 0:
		// Resolve using parent to not have side effects.
		if !conv.typeHasPtr(innerResolved) {
			conv.sides.addImport("runtime")

			// We can directly use Go's array as a pointer, as long as we defer
			// properly.
			conv.p.Linef("%s := (%s)(&%s)", value.Out, outerCGoType, value.In)
			conv.p.Linef("defer runtime.KeepAlive(&%s)", value.Out)

			return
		}

		// Target fixed array, so we can directly set the data over. The memory
		// is ours, and allocation is handled by Go.
		conv.p.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		conv.gocConverter(value.inner(value.In+"[i]", value.Out+"[i]"))
		conv.p.Linef("}")

	case array.Length != nil:
		length := fmt.Sprintf("len(%s)", value.In)
		lengthOut := conv.valueAt(*array.Length)
		lengthCGoType := anyTypeCGo(lengthOut.Type)

		// Use the backing array with the appropriate transfer-ownership rule
		// for primitive types; see type_c_go.go.
		if !value.isTransferring() && !conv.typeHasPtr(innerResolved) {
			conv.sides.addImport("runtime")
			conv.sides.addImport("unsafe")

			conv.p.Linef("%s := (%s)(unsafe.Pointer(&%s[0]))", value.Out, outerCGoType, value.In)
			conv.p.Linef("%s := %s(%s)", lengthOut.Out, lengthCGoType, length)
			conv.p.Linef("defer runtime.KeepAlive(%s)", value.Out)

			return
		}

		conv.sides.addImport("reflect")
		conv.sides.addImport("unsafe")

		// Copying is pretty much required here, since the C code will store the
		// pointer, so we can't reliably do this with Go's memory.

		conv.p.Linef("var %s %s", value.Out, outerCGoType)
		conv.p.Linef("%s := %s(%s)", lengthOut.Out, lengthCGoType, length)

		conv.p.Descend()

		conv.p.Linef("ptr := C.malloc(int(%s) * %s)", csizeof(innerResolved), length)
		// C.malloc will allocate on the C side, so we'll have to free it.
		if !value.isTransferring() {
			conv.p.Line("defer C.free(unsafe.Pointer(ptr))")
			conv.p.EmptyLine()
		}

		conv.p.Linef("var tmp []%s", innerCGoType)
		conv.p.Linef(goSliceFromPtr("dst", "tmp", length))
		conv.p.EmptyLine()

		conv.p.Linef("for i := 0; i < %s; i++ {", length)
		conv.gocConverter(value.inner(value.In+"[i]", "tmp[i]"))
		conv.p.Linef("}")

		conv.p.Linef("%s := (%s)(unsafe.Pointer(ptr_%[1]s))", value.Out, outerCGoType)
		conv.p.Linef("%s := %s(%s)", lengthOut.Out, lengthCGoType, length)

		conv.p.Ascend()

	case array.Name == "GLib.Array":
		length := fmt.Sprintf("len(%s)", value.In)

		conv.sides.addImport("reflect")
		conv.sides.addImport("unsafe")

		conv.p.Linef(
			"%s := C.g_array_sized_new(%t, false, C.guint(%s), %s)",
			value.Out, array.IsZeroTerminated(), csizeof(innerResolved), length,
		)

		conv.p.Descend()

		conv.p.Linef("var tmp []%s", innerCGoType)
		conv.p.Linef(goSliceFromPtr("tmp", value.Out+".data", length))
		conv.p.EmptyLine()

		conv.p.Linef("for i := 0; i < %s; i++ {", length)
		conv.gocConverter(value.inner(value.In+"[i]", "tmp[i]"))
		conv.p.Linef("}")

		conv.p.Ascend()

	case array.IsZeroTerminated():
		length := fmt.Sprintf("len(%s)", value.In)

		conv.sides.addImport("reflect")
		conv.sides.addImport("unsafe")

		conv.p.Linef("var %s %s", value.Out, outerCGoType)

		conv.p.Descend()

		conv.p.Linef("ptr := C.malloc(%s * (%s+1))", value.Out, csizeof(innerResolved), length)
		// See above in the array.Length != nil case.
		if !value.isTransferring() {
			conv.p.Line("defer C.free(unsafe.Pointer(ptr))")
			conv.p.EmptyLine()
		}

		conv.p.Linef("var tmp []%s", innerCGoType)
		conv.p.Linef(goSliceFromPtr("dst", "ptr_"+value.Out, length))
		conv.p.EmptyLine()

		conv.p.Linef("for i := 0; i < %s; i++ {", length)
		conv.gocConverter(value.inner(value.In+"[i]", "tmp[i]"))
		conv.p.Linef("}")

		conv.p.Ascend()

		conv.p.Linef("%s := (%s)(unsafe.Pointer(ptr_%[1]s))", value.Out, outerCGoType)

	default:
		conv.logFail(LogSkip, "weird array type to C")
	}
}

func csizeof(resolved *ResolvedType) string {
	if !strings.Contains(resolved.CType, "*") {
		return "C.sizeof_" + resolved.CType
	}

	return "unsafe.Sizeof((*[0]byte)(nil))"
}

func (conv *conversionToC) gocTypeConverter(value *GoValueProp) {
	for _, unsupported := range unsupportedCTypes {
		if unsupported == value.Type.Type.CType {
			conv.fail()
			return
		}
	}

	// Resolve for type alias.
	var result *gir.TypeFindResult
	girType := value.Type.Type

	// TODO: preserve the original cast.
	for {
		result = conv.ng.FindType(girType.Name)
		if result == nil || result.Alias == nil {
			break
		}
		// Set girType and resolve again.
		girType = &result.Alias.Type
	}

	cgoType := anyTypeCGo(value.Type)

	if prim, ok := girToBuiltin[girType.Name]; ok {
		switch prim {
		case "string":
			conv.p.Linef("%s := (%s)(C.CString(%s))", value.Out, cgoType, value.In)
			// If we're not giving ownership this C-allocated string, then we
			// can free it once done.
			if !value.isTransferring() {
				conv.sides.addImport("unsafe")
				conv.p.Linef("defer C.free(unsafe.Pointer(%s))", value.Out)
			}
		case "bool":
			conv.p.Linef("var %s %s", value.Out, cgoType)
			conv.p.Linef("if %s { %s = %s(1) }", value.In, value.Out, cgoType)
		default:
			conv.p.Linef("%s := %s(%s)", value.Out, cgoType, value.In)
		}
		return
	}

	switch ensureNamespace(conv.ng.current, girType.Name) {
	case "gpointer":
		conv.sides.addImport("github.com/diamondburned/gotk4/internal/box")
		conv.p.Linef("%s := %s(box.Assign(unsafe.Pointer(%s)))", value.Out, cgoType, value.In)
		return

	case "GObject.Type", "GType":
		// Just a primitive.
		conv.p.Linef("%s := C.GType(%s)", value.Out, value.In)
		return

	case "GObject.Value":
		// https://pkg.go.dev/github.com/gotk3/gotk3/glib#Type
		conv.p.Linef("%s := (*C.GValue)(%s.GValue)", value.Out, value.In)
		return

	case "GObject.Object":
		// Use .Native() here instead of directly accessing the native pointer,
		// since Value might be an Objector interface.
		conv.p.Linef(
			"%s := (*C.GObject)(unsafe.Pointer(%s.Native()))",
			value.Out, value.In,
		)
		return

	case "GObject.InitiallyUnowned":
		conv.p.Linef(
			"%s := (*C.GInitiallyUnowned)(unsafe.Pointer(%s.Native()))",
			value.Out, value.In,
		)
		return
	}

	switch {
	case conv.ng.mustIgnore(girType.Name, girType.CType):
		// Pretend that ignored types don't exist.
		conv.fail()

	case result.Enum != nil, result.Bitfield != nil:
		// Direct cast-able.
		conv.p.Linef("%s := (%s)(%s)", value.Out, cgoType, value.In)

	case result.Class != nil, result.Record != nil, result.Interface != nil:
		// gextras.Objector has Native() uintptr.
		conv.p.Linef("%s := (%s)(unsafe.Pointer(%s.Native()))", value.Out, cgoType, value.In)

	case result.Callback != nil:
		exportedName, _ := result.Info()
		exportedName = PascalToGo(exportedName)

		// Callbacks must have the closure attribute to store the closure
		// pointer.
		if value.Closure == nil {
			conv.logFail(LogSkip, "Go->C callback", exportedName, "since missing closure")
			return
		}

		conv.sides.addImport("github.com/diamondburned/gotk4/internal/box")
		conv.sides.addCallback(result.Callback)

		// Return the constant function here. The function will dynamically load
		// the user_data, which will match with the "gpointer" case above.
		//
		// As for the pointer to byte array cast, see
		// https://github.com/golang/go/issues/19835.
		conv.p.Linef("%s := (*[0]byte)(C.%s%s)", value.Out, callbackPrefix, exportedName)

		closureArg := conv.valueAt(*value.Closure)
		conv.p.Linef("%s := %s(box.Assign(%s))", closureArg.Out, cgoType, value.In)

		if value.Destroy != nil {
			conv.sides.CallbackDelete = true

			destroyArg := conv.valueAt(*value.Destroy)
			conv.p.Linef("%s := (*[0]byte)(C.callbackDelete)", destroyArg.Out)
		}

	case value.AllowNone:
		conv.p.Linef("var %s %s // unsupported", value.Out, cgoType)

	default:
		conv.fail()
	}
}

package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/internal/pen"
)

// Go to C type conversions.

func (ng *NamespaceGenerator) GoCConverter(conv TypeConversion) string {
	switch {
	case conv.Type.Type != nil:
		return ng.gocTypeConverter(conv)
	case conv.Type.Array != nil:
		return ng.gocArrayConverter(conv)
	default:
		return ""
	}

}

func (ng *NamespaceGenerator) gocArrayConverter(conv TypeConversion) string {
	array := conv.Type.Array

	if array.Type == nil {
		ng.logln(logWarn, "skipping nested array", array.CType)
	}

	innerResolved := ng.ResolveType(*array.Type)
	if innerResolved == nil {
		return ""
	}
	// innerType := ng.PublicType(innerResolved)
	innerCGoType := innerResolved.CGoType()

	// Generate a type converter from "src" to "${target}[i]" variables.
	innerTypeConv := conv
	innerTypeConv.Value = "src"
	innerTypeConv.Target = "dst"
	innerTypeConv.Type = array.AnyType

	innerConv := ng.GoCConverter(innerTypeConv)
	if innerConv == "" {
		return ""
	}

	var b pen.Block

	switch {
	case array.FixedSize > 0:
		if t, ok := girPrimitiveGo[array.Type.Name]; ok && t != "string" {
			return conv.callf("[%d]%s", array.FixedSize, innerCGoType)
		}

		b.Linef("goArray := ([%d]%s)(%s)", array.FixedSize)
		b.EmptyLine()
		b.Linef("for i := 0; i < %d; i++ {", array.FixedSize)
		b.Linef("  src := goArray[i]")
		b.Linef("  var dst %s", innerCGoType)
		b.Linef("  " + innerConv)
		b.Linef("  %s[i] = dst", conv.Target)
		b.Linef("}")

	case array.Length != nil:
		lengthArg := conv.ArgAt(*array.Length)
		b.Linef("%s = (*%s)(C.malloc(%s * %s))", conv.Target, innerCGoType, csizeof(ng, innerResolved), lengthArg)
		b.Linef("for i := 0; i < uintptr(%s); i++ {", lengthArg)
		b.Linef("  src := %s[i]", conv.Value)
		b.Linef("  var dst %s", innerCGoType)
		b.Linef("  " + innerConv)
		b.Linef("  *(*%s)(unsafe.Pointer(uintptr(unsafe.Pointer(%s))+i)) = dst", innerCGoType, conv.Target)
		b.Linef("}")
	}

	return b.String()
}

func csizeof(ng *NamespaceGenerator, resolved *ResolvedType) string {
	if !strings.Contains(resolved.CType, "*") {
		return "C.sizeof_" + resolved.CType
	}

	ng.addImport("unsafe")
	return "unsafe.Sizeof((*struct{})(nil))"
}

func (ng *NamespaceGenerator) gocTypeConverter(conv TypeConversion) string {
	typ := conv.Type.Type

	if prim, ok := girPrimitiveGo[typ.Name]; ok {
		switch prim {
		case "string":
			p := pen.NewPiece()
			p.Linef("%s = (*C.gchar)(C.CString(%s))", conv.Target, conv.Value)
			p.Linef("defer C.free(unsafe.Pointer(%s))", conv.Value)
			return p.String()

		case "bool":
			ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
			return conv.call("gextras.Cbool")

		default:
			return conv.call(anyTypeCGo(conv.Type))
		}
	}

	switch typ.Name {
	case "gpointer":
		ng.addImport("github.com/diamondburned/gotk4/internal/box")
		return fmt.Sprintf("%s = C.gpointer(box.Assign(%s))", conv.Target, conv.Value)

	case "GLib.DestroyNotify", "DestroyNotify":
		// This should never be called, because the caller should never see a
		// DestroyNotify, so there's no use to convert from Go to C.
		ng.logln(logError, "unexpected DestroyNotify conversion from Go to C")
		return ""

	case "GType":
		// Just a primitive.
		return fmt.Sprintf("%s = C.GType(%s)", conv.Target, conv.Value)

	case "GObject.GValue", "GObject.Value":
		// https://pkg.go.dev/github.com/gotk3/gotk3/glib#Type
		return fmt.Sprintf("%s = (*C.GValue)(%s.GValue)", conv.Target, conv.Value)

	case "GObject.Object":
		// Use .Native() here instead of directly accessing the native pointer,
		// since Value might be an Objector interface.
		return fmt.Sprintf("%s = (*C.GObject)(%s.Native())", conv.Target, conv.Value)
	case "GObject.InitiallyUnowned":
		return fmt.Sprintf("%s = (*C.GInitiallyUnowned)(%s.Native())", conv.Target, conv.Value)

	// These are empty until they're filled out in type_c_go.go
	case "GObject.Closure":
		return ""
	case "GObject.Callback":
		return ""
	case "va_list":
		return ""
	case "GObject.EnumValue", "GObject.TypeModule", "GObject.ParamSpec", "GObject.Parameter":
		// Refer to ResolveType.
		return ""
	}

	result := ng.gen.Repos.FindType(ng.current.Namespace.Name, typ.Name)
	if result == nil {
		return ""
	}

	resolved := typeFromResult(ng.gen, *typ, result)

	exportedName, _ := resolved.Extern.Result.Info()
	exportedName = PascalToGo(exportedName)

	if result.Enum != nil || result.Bitfield != nil {
		// Direct cast-able.
		return fmt.Sprintf("%s = (%s)(%s)", conv.Target, resolved.CGoType(), conv.Value)
	}

	if result.Class != nil || result.Record != nil {
		// gextras.Objector has Native() uintptr.
		return fmt.Sprintf(
			"%s = (%s)(%s.Native())",
			conv.Target, resolved.CGoType(), conv.Value,
		)
	}

	if result.Callback != nil {
		// Return the constant function here. The function will dynamically load
		// the user_data, which will match with the "gpointer" case above.
		return fmt.Sprintf("%s = (*[0]byte)(C.%s%s)", conv.Target, callbackPrefix, exportedName)
	}

	// TODO
	return ""
}

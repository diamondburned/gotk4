package girgen

import (
	"fmt"

	"github.com/diamondburned/gotk4/internal/pen"
)

// Go to C type conversions.

func (ng *NamespaceGenerator) GoCConverter(conv TypeConversion) string {
	switch {
	case conv.Type.Type != nil:
		return ng.gocTypeConverter(conv)
	default:
		return ""
	}

}

func (ng *NamespaceGenerator) gocArrayConverter(conv TypeConversion) string {
	array := conv.Type.Array

	if array.Type == nil {
		ng.logln(logWarn, "skipping nested array", array.CType)
	}

	return ""
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
			return conv.call("C." + typ.CType)
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

	case "GObject.Object", "GObject.InitiallyUnowned":
		// Use .Native() here instead of directly accessing the native pointer,
		// since Value might be an Objector interface.
		return fmt.Sprintf("%s = (%s)(%s.Native())", conv.Target, typ.CType, conv.Value)

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

	if result.Class != nil {
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

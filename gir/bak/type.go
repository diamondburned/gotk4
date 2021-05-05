package gir

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Type struct {
	XMLName        xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 type"`
	Name           string   `xml:"name,attr"`
	CType          string   `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Introspectable *bool    `xml:"introspectable,attr"`
}

func (t Type) IsPtr() bool {
	return len(t.CType) > 0 && t.CType[len(t.CType)-1] == '*'
}

func (t Type) IsConst() bool {
	return strings.HasPrefix(t.CType, "const")
}

func (t Type) IsInterface() bool {
	// TODO: find a better way
	return strings.HasSuffix(t.CType, "able*")
}

// CGoType returns the C type in CGo.
func (t Type) CGoType() string {
	return CGoType(t.CType)
}

func (t Type) GenCGoType() *jen.Statement {
	var stmt = new(jen.Statement)
	if t.IsPtr() {
		stmt.Op("*")
	}

	var ctype = t.CType
	if t.IsConst() {
		ctype = strings.TrimSpace(strings.TrimPrefix(ctype, "const"))
	}

	stmt.Qual("C", strings.TrimSuffix(ctype, "*"))

	return stmt
}

func CGoType(ctype string) (gotype string) {
	var ptr = len(ctype) > 0 && ctype[len(ctype)-1] == '*'
	var typ = fmt.Sprintf("C.%s", strings.TrimSuffix(ctype, "*"))
	if ptr {
		return "*" + typ
	}
	return typ
}

// GoType returns the type in Go.
func (t Type) GoType() string {
	return t.Type().GoString()
}

// TypeParam returns a type specifically used for interface
func (t Type) TypeParam() *jen.Statement {
	if t.IsInterface() {
		return jen.Id(InterfaceName(t.GoType()))
	}

	return t.Type()
}

// Type returns the generated Go type in Go code.
func (t Type) Type() *jen.Statement {
	return t.Map()
}

// TypeMap maps the type to a Go type in Go code.
func TypeMap(typeName string) *jen.Statement {
	return (Type{Name: typeName}).Map()
}

// Map maps the type from C to a Go type in Go code.
func (t Type) Map() *jen.Statement {
	switch t.Name {
	case "void", "none":
		return &jen.Statement{}
	case "gboolean":
		return jen.Bool()
	case "gfloat":
		return jen.Float32()
	case "gdouble":
		return jen.Float64()
	case "gint":
		return jen.Int()
	case "gint8":
		return jen.Int8()
	case "gint16":
		return jen.Int16()
	case "gint32":
		return jen.Int32()
	case "gint64":
		return jen.Int64()
	case "guint":
		return jen.Uint()
	case "guint8":
		return jen.Uint8()
	case "guint16":
		return jen.Uint16()
	case "guint32":
		return jen.Uint32()
	case "guint64":
		return jen.Uint64()
	case "utf8":
		return jen.String()
	case "gpointer":
		// TODO: ignore field
		// TODO: aaaaaaaaaaaaaaaaaaaaaaa
		return jen.Qual("unsafe", "Pointer")

	case "GLib.DestroyNotify":
		return jen.Id("DestroyNotify")
	case "GType":
		return jen.Qual("github.com/gotk3/gotk3/glib", "Type")
	case "GObject.GValue":
		return jen.Op("*").Qual("github.com/gotk3/gotk3/glib", "Value")
	case "GObject.Object":
		return jen.Op("*").Qual("github.com/gotk3/gotk3/glib", "Object")
	case "GObject.GInitiallyUnowned":
		return jen.Qual("github.com/gotk3/gotk3/glib", "InitiallyUnowned")
	case "C.gpointer":
		return jen.Uintptr()

	// We don't know what these types translates to.
	case
		// TODO: Find a way to map EnumValue type.
		"GObject.EnumValue",
		"Gtk.Buildable", "Gtk.TargetList",
		"Gio.LoadableIcon", "Gio.Cancellable", "Gio.AsyncResult":
		return nil
	}

	if parts := strings.Split(t.Name, "."); len(parts) == 2 {
		var stmt = new(jen.Statement)
		if typeMapInterface(parts) && t.IsPtr() {
			stmt = jen.Op("*")
		}

		switch parts[0] {
		case "Atk": // no idea what this translates to
			return nil
		case "Gtk":
			return stmt.Qual("github.com/gotk3/gotk3/gtk", parts[1])
		case "Gdk", "GdkPixbuf":
			return stmt.Qual("github.com/gotk3/gotk3/gdk", parts[1])
		case "GObject", "Gio", "GLib":
			return stmt.Qual("github.com/gotk3/gotk3/glib", parts[1])
		case "Pango":
			return stmt.Qual("github.com/gotk3/gotk3/pango", parts[1])
		case "Cairo":
			return stmt.Qual("github.com/gotk3/gotk3/cairo", parts[1])
		}
	}

	// Is this an interface? If yes, don't treat them as a pointer.
	if !t.IsInterface() && t.IsPtr() {
		return jen.Op("*").Id(t.Name)
	}

	return jen.Id(t.Name)
}

func (t Type) ZeroValue() *jen.Statement {
	if t.IsPtr() || t.IsFunc() {
		return jen.Nil()
	}

	switch t.Name {
	case "gboolean":
		return jen.False()
	case "gfloat", "gdouble":
		return jen.Lit(0.0)
	case "gint", "guint", "GType": // TODO: gint* and guint*
		return jen.Lit(0)
	case "gpointer":
		return jen.Qual("C", "gpointer").Call(jen.Uintptr().Call(jen.Lit(0)))
	case "utf8":
		return jen.Lit("")
	case "GLib.DestroyNotify":
		return jen.Nil()
	}

	return t.Type().Values()
}

func typeMapInterface(parts []string) (ptr bool) {
	switch parts[0] {
	case "Gtk":
		switch parts[1] {
		case "Widget":
			parts[1] = "IWidget"
			return false
		}
	}

	return true
}

// GenCaster generates the type or function to be used to cast or convert C to
// Go types.
func (t Type) GenCaster(tmpVar, value *jen.Statement) *jen.Statement {
	var stmt = tmpVar.Clone().Op(":=")
	var goType = t.GoType()

	switch goType {
	case "bool":
		stmt.Id("gobool")
	case "string":
		stmt.Qual("C", "GoString")
	case "uintptr":
		return stmt.Qual("unsafe", "Pointer").Call(jen.Uintptr().Call(value))

	// Handle IWidget separately.
	case "gtk.IWidget":
		stmt = jen.List(tmpVar, jen.Err()).Op(":=").Id("castWidget").Call(value)
		stmt.Line()
		stmt.If(jen.Err().Op("!=").Nil()).Block(
			jen.Panic(
				jen.Lit(fmt.Sprintf("cast widget %s failed: ", t.CGoType())).
					Op("+").
					Err().Dot("Error").Call(),
			),
		)

		return stmt

	// Handle glib.Object separately.
	case "glib.Object", "*glib.Object":
		// TODO: see if this leaks.
		return stmt.Add(genObjTake(value))

	// Handle *glib.SList separately.
	case "glib.SList", "*glib.SList":
		return stmt.Qual("github.com/gotk3/gotk3/glib", "WrapSList").Call(
			jen.Uintptr().Call(jen.Qual("unsafe", "Pointer").Call(value)),
		)

	// Handle glib.ListModel separately. TODO: handle all glib types that
	// embed *glib.Object.
	case "glib.ListModel", "*glib.ListModel":
		// Enforce a non-pointer when using resolveWrapValues.
		return genObjectCtor(value, tmpVar).Op("&").Add(resolveWrapValues("glib.ListModel"))

	case "glib.Value", "*glib.Value":
		return stmt.Qual("github.com/gotk3/gotk3/glib", "ValueFromNative").Call(
			jen.Call(jen.Qual("unsafe", "Pointer").Call(value)),
		)

	default:
		switch {
		case t.IsFunc():
			log.Panicln("Unsure GenCaster for func type", t.Name)
		case t.IsEnum():
			break
		case t.IsInterface():
			if t := EmbeddedFieldNoPanic(goType); t != "" {
				return genObjectCtor(value, tmpVar).Op("&").Add(resolveWrapValues(goType))
			}

			return stmt.Id(goType).Values(jen.Line().
				Id("Object").Op(":").Add(genObjTake(value)).Op(",").
				Line(),
			)

		default:
			// Is this a known class? If yes, then use its wrap function.
			for _, class := range activeNamespace.Classes {
				if class.Name == t.Name {
					return stmt.Add(jen.Id(class.WrapperFnName())).Call(
						jen.Qual("unsafe", "Pointer").Call(value),
					)
				}
			}

			// See if any of our types are wrappable. Ignore pointers.
			var derefType = strings.TrimPrefix(goType, "*")

			if t := EmbeddedFieldNoPanic(derefType); t != "" {
				return genObjectCtor(value, tmpVar).Op("&").Add(resolveWrapValues(derefType))
			}
		}

		if t.IsPtr() {
			stmt.Parens(t.Type())
		} else {
			stmt.Add(t.Type())
		}
	}

	stmt.Call(value)

	return stmt
}

func genObjectCtor(value, tmpVar *jen.Statement) *jen.Statement {
	stmt := jen.Id("obj").Op(":=").Add(genObjTake(value))
	stmt.Line()
	stmt.Add(tmpVar).Op(":=")
	return stmt
}

func genObjTake(objVar *jen.Statement) *jen.Statement {
	return jen.Qual("github.com/gotk3/gotk3/glib", "Take").Call(
		jen.Qual("unsafe", "Pointer").Call(objVar),
	)
}

// CNeedsFree returns true if the generated value from GenCCaster needs freeing.
// Pay attention to transfer-ownership when doing this.
func (t Type) CNeedsFree() bool {
	return t.GoType() == "string"
}

// IsFunc returns true if the given type is a callback.
func (t Type) IsFunc() bool {
	// TODO: find a better way to check a func callback.
	return strings.HasSuffix(t.Name, "Func")
}

// IsNamespaceFunc returns true if the current type is a callback that belongs
// to the current namespace.
func (t Type) IsNamespaceFunc() bool {
	if !t.IsFunc() {
		return false
	}

	for _, callback := range activeNamespace.Callbacks {
		if callback.Name == t.Name {
			return true
		}
	}

	return false
}

// knownEnums is the list of known external enums. This list exists to lazily
// detect enum types without needing to keep multiple gir namespaces.
//
// TODO: allow for namespace detection
var knownEnums = []string{
	"Gtk.Orientation",
	"Gtk.IconSize",
	"Gtk.PackType",
	"Pango.EllipsizeMode",
}

func (t Type) IsEnum() bool {
	for _, enum := range activeNamespace.Enums {
		if enum.Name == t.Name {
			return true
		}
	}

	for _, enum := range knownEnums {
		if enum == t.Name {
			return true
		}
	}

	return false
}

// GenCCaster generates a function or type cast to convert Go values to C.
func (t Type) GenCCaster(value *jen.Statement) *jen.Statement {
	// TODO: account for enums

	switch goType := t.GoType(); goType {
	case "bool":
		return jen.Id("cbool").Call(value)
	case "float32":
		return jen.Qual("C", "gfloat").Call(value)
	case "float64":
		return jen.Qual("C", "gdouble").Call(value)
	case "int":
		return jen.Qual("C", "gint").Call(value)
	case "uint":
		return jen.Qual("C", "guint").Call(value)
	case "unsafe.Pointer":
		return jen.Qual("C", "gpointer").Call(jen.Uintptr().Call(value))
	case "string":
		return jen.Qual("C", "CString").Call(value)
	case "gtk.IWidget":
		return jen.Id("cwidget").Call(value)
	case "glib.Type":
		return jen.Qual("C", "GType").Call(value)
	case "*gdk.Rectangle":
		return jen.Parens(jen.Op("*").Qual("C", "GdkRectangle")).Call(
			jen.Qual("unsafe", "Pointer").Call(jen.Op("&").Add(value).Dot("GdkRectangle")),
		)
	default:
		switch {
		// Handle int and uint specifically.
		case strings.HasPrefix(goType, "int"), strings.HasPrefix(goType, "uint"):
			return jen.Qual("C", fmt.Sprintf("g%s", goType)).Call(value)
		case t.IsFunc():
			return ZeroByteCast(jen.Qual("C", CallbackExternCName(t.Name)))
		case t.IsEnum():
			return t.GenCGoType().Call(value)
		}

		// See if any of our types are wrappable. Ignore pointers.
		var derefType = strings.TrimPrefix(goType, "*")

		switch {
		case EmbeddedFieldCheck(derefType, "gtk.Widget") && !t.IsInterface():
			// Avoid an ambiguous selector.
			value = value.Clone().Dot("Widget")
		}

		return jen.Parens(t.GenCGoType()).Call(
			jen.Qual("unsafe", "Pointer").Call(value.Clone().Op(".").Id("Native").Call()),
		)
	}
}

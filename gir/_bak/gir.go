// Package gir provides a roughly-written gir generator.
package gir

import (
	"encoding/xml"
	"log"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/dave/jennifer/jen"
)

func firstChar(str string) string {
	r, _ := utf8.DecodeRune([]byte(str))
	return string(unicode.ToLower(r))
}

var (
	snakeRegex = regexp.MustCompile(`_\w`)
	snakeRepl  = strings.NewReplacer(
		"Xalign", "XAlign",
		"Yalign", "YAlign",
		"Id", "ID",
	)
)

// GoNamer is the interface for structs that can output idiomatic Go type names.
type GoNamer interface {
	GoName() string
}

func snakeToGo(pascal bool, snakeString string) string {
	if pascal {
		snakeString = "_" + snakeString
	}

	snakeString = snakeRegex.ReplaceAllStringFunc(snakeString,
		func(orig string) string {
			return string(unicode.ToUpper(rune(orig[1])))
		},
	)

	return snakeRepl.Replace(snakeString)
}

func NewGotk3Generator(name string) *jen.File {
	f := jen.NewFile(name)
	f.ImportName("github.com/gotk3/gotk3/gtk", "gtk")
	f.ImportName("github.com/gotk3/gotk3/gdk", "gdk")
	f.ImportName("github.com/gotk3/gotk3/glib", "glib")
	f.ImportName("github.com/gotk3/gotk3/pango", "pango")
	f.ImportName("github.com/gotk3/gotk3/cairo", "cairo")
	f.ImportName("github.com/diamondburned/handy/internal/callback", "callback")
	f.CgoPreamble("#cgo pkg-config: libhandy-1 gtk+-3.0 glib-2.0 gio-2.0 glib-2.0 gobject-2.0")
	f.CgoPreamble("#cgo CFLAGS: -Wno-deprecated-declarations")
	f.CgoPreamble("#include <handy.h>")
	f.CgoPreamble("#include <gtk/gtk.h>")
	f.CgoPreamble("#include <gio/gio.h>")
	f.CgoPreamble("#include <glib.h>")
	f.CgoPreamble("#include <glib-object.h>")
	f.CgoPreamble("extern void callbackDelete(gpointer ptr);")

	f.Comment("//export callbackDelete")
	f.Func().Id("callbackDelete").Params(jen.Id("ptr").Qual("C", "gpointer")).Block(
		jen.Qual("github.com/diamondburned/handy/internal/callback", "Delete").Call(
			jen.Uintptr().Call(jen.Id("ptr")),
		),
	)
	f.Line()

	f.Comment("objector is used internally for other interfaces.")
	f.Type().Id("objector").Interface(
		jen.Qual("github.com/gotk3/gotk3/glib", "IObject"),
		jen.Id("Connect").
			Call(jen.String(), jen.Interface()).
			Qual("github.com/gotk3/gotk3/glib", "SignalHandle"),
		jen.Id("ConnectAfter").
			Call(jen.String(), jen.Interface()).
			Qual("github.com/gotk3/gotk3/glib", "SignalHandle"),
		jen.Id("GetProperty").
			Call(jen.Id("name").String()).
			Parens(jen.List(jen.Interface(), jen.Error())),
		jen.Id("SetProperty").
			Call(jen.Id("name").String(), jen.Id("value").Interface()).
			Parens(jen.Error()),
		jen.Id("Native").
			Call().
			Parens(jen.Uintptr()),
	)
	f.Line()

	f.Comment("asserting objector interface")
	f.Var().Id("_").Id("objector").Op("=").
		Parens(jen.Op("*").Qual("github.com/gotk3/gotk3/glib", "Object")).
		Call(jen.Nil())
	f.Line()

	f.Add(GenCasterInterface())
	f.Line()

	return f
}

type Annotation struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type CInclude struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Name    string   `xml:"name,attr"`
}

type Include struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	Name    string   `xml:"name,attr"`
	Version *string  `xml:"version,attr"`
}

type Package struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 package"`
	Name    string   `xml:"name,attr"`
}

type TransferOwnership struct {
	TransferOwnership *string `xml:"transfer-ownership,attr"`
}

// EmbeddedFieldCheck checks if the given goType embeds the required
// containsType.
func EmbeddedFieldCheck(goType, containsType string) bool {
	for {
		switch goType = EmbeddedFieldNoPanic(goType); goType {
		case containsType:
			return true
		case "*glib.Object":
			fallthrough
		case "":
			return false
		}
	}
}

func EmbeddedField(goType string) string {
	var t = EmbeddedFieldNoPanic(goType)
	if t == "" {
		log.Panicln("Unknown type:", goType)
	}
	return t
}

func EmbeddedFieldNoPanic(goType string) string {
	switch goType {
	case "gtk.ApplicationWindow": // TODO: handle interfaces inside here
		return "gtk.Window"

	case "gtk.Window":
		fallthrough
	case "gtk.ListBoxRow":
		fallthrough
	case "gtk.EventBox":
		return "gtk.Bin"

	case "gtk.HeaderBar":
		fallthrough
	case "gtk.Stack":
		fallthrough
	case "gtk.Bin":
		return "gtk.Container"

	case "gtk.Entry":
		fallthrough
	case "gtk.Container":
		fallthrough
	case "gtk.DrawingArea":
		return "gtk.Widget"

	case "gtk.Widget":
		return "glib.InitiallyUnowned"

	case "gdk.Pixbuf":
		fallthrough
	case "glib.MenuModel":
		fallthrough
	case "glib.Icon":
		fallthrough
	case "glib.InitiallyUnowned":
		fallthrough
	case "glib.ListModel":
		return "*glib.Object"
	}

	for _, class := range activeNamespace.Classes {
		if class.GoName() == goType {
			return class.ParentInstanceType()
		}
	}

	for _, iface := range activeNamespace.Interfaces {
		if iface.GoName() == goType {
			if iface.RequiresWidget() {
				return "gtk.Widget"
			}
			return "*glib.Object"
		}
	}

	return ""
}

func GenMarshalerFnName(typeName string) *jen.Statement {
	return jen.Id("marshal" + typeName)
}

// GenMarshalerFn generates the marshal function's surroundings. The generated
// function would be named typeName prefixed by "marshal". The uintptr argument
// is named p.
func GenMarshalerFn(typeName string, body *jen.Statement) *jen.Statement {
	return jen.
		Func().
		Add(GenMarshalerFnName(typeName)).
		Params(jen.Id("p").Uintptr()).
		Params(jen.Interface(), jen.Error()).
		Block(body).
		Line()
}

func GenMarshalerItem(getType, typeName string) *jen.Statement {
	return jen.Values(
		jen.Qual("github.com/gotk3/gotk3/glib", "Type").Call(jen.Qual("C", getType).Call()),
		GenMarshalerFnName(typeName),
	)
}

// resolveWrapValues resolves embedded struct fields. Note that only
// *glib.Object is allowed to be a pointer in childType.
func resolveWrapValues(childType string, implements ...Implements) *jen.Statement {
	return resolveWrapValueField(childType, "", implements...)
}

// resolveWrapValueField does what resolveWrapValues does but iwth a custom
// field name.
func resolveWrapValueField(childType, fieldN string, implements ...Implements) *jen.Statement {
	switch childType {
	case "*glib.Object":
		return jen.Id("obj")
	case "":
		return nil
	}

	var embedT = EmbeddedField(childType)
	if fieldN == "" {
		fieldN = fieldNameFromType(embedT)
	}

	// Treat interfaces specially.
	if iface := activeNamespace.FindInterface(childType); iface != nil {
		// TODO: confirm this is not needed.
		if len(implements) > 0 {
			log.Panicf("Interface %s shouldn't have implements\n", iface.Name)
		}

		return GenInterfaceWrapper(iface.GoName(), iface.RequiresWidget())
	}

	var values = jen.Statement{
		jen.Id(fieldN).Op(":").Add(resolveWrapValues(embedT)).Op(",").Line(),
	}

	for _, impls := range implements {
		if iface := activeNamespace.FindInterface(impls.Name); iface != nil {
			values.Add(jen.Id(iface.InterfaceName()).Op(":").Op("&").Add(
				GenInterfaceWrapper(iface.Name, iface.RequiresWidget()),
			))
			values.Op(",").Line()
			continue
		}

		// Try and make a rough guess.
		if t := TypeMap(impls.Name); t != nil {
			var goType = t.GoString()
			values.Add(jen.Id(fieldNameFromType(goType)).Op(":").Add(
				jen.Id(goType).Values(resolveWrapValues("*glib.Object")),
			))
			values.Op(",").Line()
		}
	}

	return jen.Id(childType).Values(jen.Line().Add(values...).Line())
}

func fieldNameFromType(typeName string) string {
	var fields = strings.Split(typeName, ".")
	var fieldN = strings.TrimPrefix(fields[len(fields)-1], "*")
	return fieldN
}

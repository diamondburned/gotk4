package gir

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Implements struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 implements"`
	Name    string   `xml:"name,attr"`
}

type Prerequisite struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 prerequisite"`
	Name    string   `xml:"name,attr"`
}

// GenInterfaceWrappper generates the interface struct. This function assumes
// the object is named "obj".
func GenInterfaceWrapper(ifaceGObjName string, widget bool) *jen.Statement {
	// Ignore pointers.
	ifaceGObjName = strings.TrimSuffix(ifaceGObjName, "*")

	var stmt = TypeMap(ifaceGObjName)
	if !widget {
		return stmt.Values(jen.Id("obj"))
	}

	return stmt.Values(
		jen.Line().Id("Caster").Op(":").Op("&").Add(
			resolveWrapValues("gtk.Widget").Op(",").Line(),
		),
	)
}

// GenCasterInterface generates the constant caster interface that helps conceal
// the underlying Widget.
func GenCasterInterface() *jen.Statement {
	stmt := jen.Comment("Caster is the interface that allows casting objects to widgets.")
	stmt.Line()
	stmt.Type().Id("Caster").Interface(
		jen.Id("objector"),
		jen.Id("Cast").Params().Params(
			jen.Qual("github.com/gotk3/gotk3/gtk", "IWidget"),
			jen.Error(),
		),
	)
	stmt.Line()

	return stmt
}

type Interface struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 interface"`
	Name    string   `xml:"name,attr"`

	CType         string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CSymbolPrefix string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`

	GLibTypeName   string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType    string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GLibTypeStruct string `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`

	Functions     []Function     `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Methods       []Method       `xml:"http://www.gtk.org/introspection/core/1.0 method"` // translated to Go fns
	Prerequisites []Prerequisite `xml:"http://www.gtk.org/introspection/core/1.0 prerequisite"`

	// Constructor    *Constructor `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	// Implementses   []Implements    `xml:"http://www.gtk.org/introspection/core/1.0 implements"`
	// Fields         []Field         `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	// Callbacks      []Callback      `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	// Constants      []Constant      `xml:"http://www.gtk.org/introspection/core/1.0 constant"`
	// VirtualMethods []VirtualMethod `xml:"http://www.gtk.org/introspection/core/1.0 virtual-method"`
}

func (i Interface) GenerateAll() *jen.Statement {
	s := new(jen.Statement)
	s.Add(i.GenInterface())
	s.Line()
	s.Add(i.GenType())
	s.Line()
	s.Add(i.GenNative())
	s.Line()
	s.Add(i.GenMethods())
	return s
}

func (i Interface) GenNative() *jen.Statement {
	c := firstChar(i.Name)
	p := jen.Id(c).Op("*").Id(i.GoName())

	f := jen.Add(GenCommentReflowLines("native", fmt.Sprintf(
		"turns the current *%s into the native C pointer type.",
		i.GoName(),
	)))

	f.Func().Params(p).Id("native").Params().Id("*" + i.CGoType())

	return f.Block(
		jen.Return(
			jen.Parens(jen.Op("*").Qual("C", i.CType)).Call(
				jen.Qual("unsafe", "Pointer").Call(
					jen.Id(c).Dot("Native").Call(),
				),
			),
		),
	)
}

func (i Interface) CGoType() string {
	return CGoType(i.CType)
}

func (i Interface) GenType() *jen.Statement {
	var stmt = jen.Type().Id(i.GoName())
	if i.RequiresWidget() {
		stmt.Struct(jen.Id("Caster"))
	} else {
		stmt.Struct(TypeMap("GObject.Object"))
	}
	return stmt
}

func (i Interface) GenMethods() *jen.Statement {
	var name = i.GoName()
	var stmt = new(jen.Statement)

	for _, m := range i.Methods {
		if m.IsIgnored() {
			continue
		}

		stmt.Add(m.GenFunc(name))
		stmt.Line()
	}

	return stmt
}

// RequiresWidget returns true if the interface requires a widget.
func (i Interface) RequiresWidget() bool {
	// TODO: convert to Go type and assert nested structs.
	for _, prereq := range i.Prerequisites {
		if prereq.Name == "Gtk.Widget" {
			return true
		}
	}

	return false
}

func (i Interface) GoName() string {
	return snakeToGo(true, i.Name)
}

// InterfaceName generates an idiomatic Go interface name.
func (i Interface) InterfaceName() string {
	return InterfaceName(i.Name)
}

// InterfaceName turns glib interface naming conventions to Go.
func InterfaceName(ifaceName string) string {
	var name = strings.TrimSuffix(ifaceName, "able")
	if !strings.HasSuffix(name, "e") {
		name += "er"
	} else {
		name += "r"
	}

	return name
}

func (i Interface) GenInterface() *jen.Statement {
	var name = i.InterfaceName()
	var methods = jen.Statement{}

	// Always implement either a base GObject interface or the widget caster
	// interface.

	if i.RequiresWidget() {
		methods = append(methods, jen.Id("Caster"))
	} else {
		methods = append(methods, jen.Id("objector"))
	}

	for _, m := range i.Methods {
		if m.IsIgnored() {
			continue
		}

		var stmt = new(jen.Statement)
		if m.Doc != nil {
			stmt.Add(m.Doc.GenGoComments(name, m.GoName()))
		}

		var parm = []Parameter{}
		if m.Parameters != nil {
			parm = m.Parameters.Parameters
		}
		var args = make(map[string]*jen.Statement, len(parm)+1)

		// Generate the parameters in the function signature.
		stmt.Id(m.GoName()).ParamsFunc(func(g *jen.Group) {
			for _, param := range parm {
				if param.IsIgnored() {
					continue
				}

				n := jen.Id(param.GoName())
				args[param.Name] = n

				g.Add(n, param.Type.Type())
			}
		})

		if m.ReturnValue != nil {
			stmt.Add(m.ReturnValue.Type.Type())
		}

		methods = append(methods, stmt)
	}

	return jen.Type().Id(name).Interface(methods...).Line()
}

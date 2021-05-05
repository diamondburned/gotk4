package gir

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/dave/jennifer/jen"
)

type Class struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 class"`
	Name    string   `xml:"name,attr"`

	CType         string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CSymbolPrefix string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`

	Parent string `xml:"parent,attr"`

	GLibTypeName   string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType    string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GLibTypeStruct string `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`

	Implements   []Implements  `xml:"http://www.gtk.org/introspection/core/1.0 implements"`
	Constructors []Constructor `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Methods      []Method      `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Fields       []Field       `xml:"http://www.gtk.org/introspection/core/1.0 field"`

	// Functions    []Function    `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	// Callbacks    []Callback    `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
}

func (c Class) FnWithC(CIdentifier string) interface{} {
	for _, method := range c.Methods {
		if method.CIdentifier == CIdentifier {
			return method
		}
	}

	for _, constructor := range c.Constructors {
		if constructor.CIdentifier == CIdentifier {
			return constructor
		}
	}

	return nil
}

func (c Class) GenerateAll() *jen.Statement {
	f := new(jen.Statement)
	f.Add(c.GenType())
	f.Line()
	f.Add(c.GenWrapper())
	f.Line()
	f.Add(c.GenMarshaler())
	f.Line()
	f.Add(c.GenConstructors())
	f.Line()
	f.Add(c.GenNative())
	f.Line()
	f.Add(c.GenMethods())
	return f
}

func (c Class) GenType() *jen.Statement {
	var fields = jen.Statement{
		c.GenParentInstanceType(),
	}

	var ifaceFields jen.Statement

	for _, impls := range c.Implements {
		if iface := activeNamespace.FindInterface(impls.Name); iface != nil {
			// Use interfaces to conceal the underlying GObject methods which
			// avoids ambiguous selectors.
			// TODO: make gotk3 do this too. This is useless as long as gotk3's
			// interfaces are still bad.
			ifaceFields.Id(iface.InterfaceName())

			// ifaceFields.Id(iface.GoName())
			continue
		}

		// Try and make a rough guess.
		if t := TypeMap(impls.Name); t != nil {
			ifaceFields.Add(t)
		}
	}

	if len(ifaceFields) > 0 {
		fields.Line().Comment("Interfaces")
	}

	return jen.Type().Id(c.Name).Struct(append(fields, ifaceFields...)...)
}

func (c Class) WrapperFnName() string {
	return "wrap" + c.GoName()
}

func (c Class) GenWrapper() *jen.Statement {
	var name = c.WrapperFnName()
	var gtyp = c.GoName()

	s := GenCommentReflowLines(
		name,
		fmt.Sprintf("wraps the given pointer to *%s.", gtyp),
	)

	s.Func().Id(name).Params(jen.Id("ptr").Qual("unsafe", "Pointer")).Op("*").Id(gtyp).Block(
		jen.Id("obj").Op(":=").Qual("github.com/gotk3/gotk3/glib", "Take").Call(jen.Id("ptr")),
		jen.Return(jen.Op("&").Add(resolveWrapValues(gtyp, c.Implements...))),
	)

	s.Line()

	return s
}

func (c Class) GenMarshalerItem() *jen.Statement {
	return GenMarshalerItem(c.GLibGetType, c.GoName())
}

func (c Class) GenMarshaler() *jen.Statement {
	var goName = c.GoName()
	var wrapFn = c.WrapperFnName()

	return GenMarshalerFn(
		goName,
		jen.Return(
			jen.Id(wrapFn).Call(
				jen.Qual("unsafe", "Pointer").Call(
					jen.Qual("C", "g_value_get_object").Call(
						jen.Parens(jen.Op("*").Qual("C", "GValue")).Call(
							jen.Qual("unsafe", "Pointer").Call(jen.Id("p")),
						),
					),
				),
			),
			jen.Nil(),
		),
	)
}

func (c Class) GenConstructors() *jen.Statement {
	var stmt = make(jen.Statement, 0, len(c.Constructors)*2)
	for _, ctor := range c.Constructors {
		if ctor.IsIgnored() {
			continue
		}

		stmt.Add(ctor.GenFunc(c))
		stmt.Line()
	}

	return &stmt
}

func (c Class) GenNative() *jen.Statement {
	i := firstChar(c.Name)
	p := jen.Id(i).Op("*").Id(c.GoName())

	f := jen.Add(GenCommentReflowLines("native", fmt.Sprintf(
		"turns the current *%s into the native C pointer type.",
		c.GoName(),
	)))

	f.Func().Params(p).Id("native").Params().Id("*" + c.CGoType())

	switch goType := c.GoName(); {
	// We can only use gwidget() if the class inherits gtk.Widget.
	case EmbeddedFieldCheck(goType, "gtk.Widget"):
		// Call the struct name to avoid ambiguities.
		var prnt = fieldNameFromType(c.ParentInstanceType())

		f.Block(
			jen.Return(
				jen.Parens(jen.Op("*").Qual("C", c.CType)).Call(
					jen.Id("gwidget").Call(jen.Op("&").Id(i).Dot(prnt)),
				),
			),
		)

	// We can use Native() with an Object.
	case EmbeddedFieldCheck(goType, "*glib.Object"):
		f.Block(
			jen.Return(
				jen.Parens(jen.Op("*").Qual("C", c.CType)).Call(
					jen.Qual("unsafe", "Pointer").Call(
						jen.Id(i).Dot("Native").Call(),
					),
				),
			),
		)

	default:
		log.Panicln("Unknown Native of type", goType)
	}

	f.Line()
	return f
}

func (c Class) GenMethods() *jen.Statement {
	var stmt = make(jen.Statement, 0, len(c.Methods)*3)
	for _, method := range c.Methods {
		if method.IsIgnored() {
			continue
		}

		stmt.Add(method.GenFunc(c.Name))
		stmt.Line()
	}

	return &stmt
}

func (c Class) GenParentInstanceType() *jen.Statement {
	var t = TypeMap(c.Parent)
	if t.GoString() == "glib.InitiallyUnowned" {
		return jen.Qual("github.com/gotk3/gotk3/glib", "InitiallyUnowned")
	}

	return t
}

func (c Class) ParentInstanceType() string {
	return c.GenParentInstanceType().GoString()
}

func (c Class) CGoType() string {
	return CGoType(c.CType)
}

func (c Class) GoName() string {
	return snakeToGo(true, c.Name)
}

package gir

import (
	"encoding/xml"

	"github.com/dave/jennifer/jen"
)

type Member struct {
	XMLName     xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 member"`
	Name        string   `xml:"name,attr"`
	Value       int      `xml:"value,attr"`
	CIdentifier string   `xml:"http://www.gtk.org/introspection/c/1.0 identifer,attr"`

	Doc *Doc
}

func (m Member) GoName() string {
	return snakeToGo(true, m.Name)
}

type Enum struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 enumeration"`
	Name    string   `xml:"name,attr"` // Go case
	Version string   `xml:"version,attr"`

	Doc *Doc

	GLibTypeName string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType  string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`

	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	Members []Member `xml:"http://www.gtk.org/introspection/core/1.0 member"`
}

func (e Enum) GoName() string {
	return snakeToGo(true, e.Name)
}

func (e Enum) GenerateAll() *jen.Statement {
	f := new(jen.Statement)
	f.Add(e.GenType())
	f.Line()
	f.Add(e.GenMarshaler())
	f.Line()
	f.Add(e.GenConsts())
	return f
}

func (e Enum) GenMarshalerItem() *jen.Statement {
	return GenMarshalerItem(e.GLibGetType, e.GoName())
}

func (e Enum) GenMarshaler() *jen.Statement {
	var goName = e.GoName()
	return GenMarshalerFn(goName,
		jen.Return(
			jen.Id(goName).Call(
				jen.Qual("C", "g_value_get_enum").Call(
					jen.Parens(jen.Op("*").Qual("C", "GValue")).Call(
						jen.Qual("unsafe", "Pointer").Call(jen.Id("p")),
					),
				),
			),
			jen.Nil(),
		),
	)
}

func (e Enum) GenType() *jen.Statement {
	var s = new(jen.Statement)
	if e.Doc != nil {
		s.Add(e.Doc.GenGoComments("", e.GoName()))
	}

	return s.Type().Id(e.GoName()).Int()
}

func (e Enum) GenConsts() *jen.Statement {
	var enumName = e.GoName()

	return jen.Const().DefsFunc(func(g *jen.Group) {
		for _, member := range e.Members {
			var fullName = enumName + member.GoName()

			var s = new(jen.Statement)
			if member.Doc != nil {
				s.Add(member.Doc.GenGoCommentsIndent(1, "", fullName))
			}

			s.Id(fullName).Id(enumName).Op("=").Lit(member.Value)
			g.Add(s)
		}
	})
}

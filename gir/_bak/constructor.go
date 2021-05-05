package gir

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Constructor struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	CallableAttrs
}

func (c Constructor) TypeName() string {
	// Slice away the namespace.
	id := c.CIdentifier[strings.Index(c.CIdentifier, "_")+1:]

	// Slice away the name.
	id = strings.TrimSuffix(id, fmt.Sprintf("_%s", c.Name))

	return snakeToGo(true, id)
}

func (c Constructor) GoName() string {
	return c.TypeName() + snakeToGo(true, c.Name)
}

func (c Constructor) GenFunc(class Class) *jen.Statement {
	var s = new(jen.Statement)
	if c.Doc != nil {
		s.Add(c.Doc.GenGoComments("", c.GoName()))
	} else {
		s.Add(GenCommentReflowLines(
			c.GoName(),
			fmt.Sprintf("creates a new %s.", class.GoName()),
		))
	}

	var parm = []Parameter{}
	if c.Parameters != nil {
		parm = c.Parameters.Parameters
	}
	var args = make(map[string]*jen.Statement, len(parm)+1)

	s.Func().Id(c.GoName())
	s.ParamsFunc(func(g *jen.Group) {
		for _, param := range parm {
			if param.IsIgnored() {
				continue
			}

			n := jen.Id(param.GoName())
			args[param.Name] = n

			g.Add(n, param.Type.Type())
		}
	})

	s.Op("*").Id(class.GoName())

	var cargs = make(map[string]*jen.Statement, len(parm)+1)

	return s.BlockFunc(func(g *jen.Group) {
		for i, param := range parm {
			if arg, hasArgument := args[param.Name]; hasArgument {
				var valueVar = jen.Id(fmt.Sprintf("v%d", i+1))
				cargs[param.Name] = valueVar

				g.Add(param.GenValueCall(arg, valueVar))
			}
		}

		if len(parm) > 1 {
			g.Line()
		}

		g.Return(
			jen.Id(class.WrapperFnName()).Call(
				jen.Qual("unsafe", "Pointer").Call(
					jen.Qual("C", c.CIdentifier).ParamsFunc(func(g *jen.Group) {
						for _, param := range parm {
							a, ok := cargs[param.Name]
							if ok {
								g.Add(a)
							} else {
								// Add as a constant to allow implicit type casting.
								g.Add(param.Type.ZeroValue())
							}
						}
					}),
				),
			),
		)
	})

	// return s.Op("*").Id(class.GoName()).Block(
	// 	jen.Id("v").Op(":=").Qual("C", c.CIdentifier).Call(),
	// 	jen.Return(jen.Id(class.WrapperFnName()).Call(
	// 		jen.Qual("unsafe", "Pointer").Call(jen.Id("v")),
	// 	)),
	// )
}

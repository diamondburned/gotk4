package gir

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/dave/jennifer/jen"
)

type Method struct {
	XMLName     xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Name        string   `xml:"name,attr"`
	CIdentifier string   `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	CallableAttrs
}

func (m Method) GenFunc(parentType string) *jen.Statement {
	stmt := m.genFunc(parentType)
	if stmt == nil {
		log.Printf("skipping function %s.%s", parentType, m.GoName())
		return &jen.Statement{}
	}

	return stmt
}

func (m Method) genFunc(parentType string) *jen.Statement {
	i := firstChar(parentType)
	p := jen.Id(i).Op("*").Id(parentType)

	var stmt = new(jen.Statement)
	if m.Doc != nil {
		stmt.Add(m.Doc.GenGoComments(i, m.GoName()))
	}

	if m.IsDeprecated() {
		if m.Doc != nil {
			stmt.Comment("")
			stmt.Line()
		}
		stmt.Commentf("This method is deprecated since version %s.", m.DeprecatedVersion)
		stmt.Line()
	}

	stmt.Func().Params(p).Id(m.GoName())

	var parm = []Parameter{}
	if m.Parameters != nil {
		parm = m.Parameters.Parameters
	}

	for _, param := range parm {
		if param.IsBlockedType() {
			return nil
		}
	}

	var args = make(map[string]*jen.Statement, len(parm)+1)

	// Generate the parameters in the function signature.
	stmt.ParamsFunc(func(g *jen.Group) {
		for _, param := range parm {
			if param.IsIgnored() {
				continue
			}

			n := jen.Id(param.GoName())
			args[param.Name] = n

			g.Add(n, param.Type.TypeParam())
		}
	})

	if m.ReturnValue != nil && m.ReturnValue.Type != nil {
		if m.ReturnValue.Type.Map() == nil {
			return nil
		}
		stmt.Add(m.ReturnValue.Type.TypeParam())
	}

	// List of arguments to call the C function. Not to be confused with the
	// above list of arguments to call the current Go function.
	var cargs = make(map[string]*jen.Statement, len(parm)+1)

	// Generate the value type converters in the function body.
	stmt.BlockFunc(func(g *jen.Group) {
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

		g.Add(m.ReturnValue.GenReturnFunc(
			jen.Qual("C", m.CIdentifier).ParamsFunc(func(g *jen.Group) {
				if m.HasInstanceParameter() {
					g.Add(jen.Id(i).Op(".").Id("native").Call())
				}

				for i, param := range parm {
					switch arg, hasCArgument := cargs[param.Name]; {
					case hasCArgument:
						g.Add(arg)

					// Treat UserData and UserDataFreeFunc specially.
					case param.IsUserData():
						if callback := m.UserDataParameter(i); callback != nil {
							g.Add(CallbackGenAssign(args[callback.Name]))
						} else {
							g.Add(param.Type.ZeroValue())
						}

					case param.IsUserDataFreeFunc():
						g.Add(CallbackGenDelete())

					default:
						// Add as a constant to allow implicit type casting.
						g.Add(param.Type.ZeroValue())
					}
				}
			}),
		))
	})

	return stmt
}

func (m Method) GoName() string {
	return snakeToGo(true, m.Name)
}

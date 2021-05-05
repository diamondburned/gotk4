package gir

import (
	"encoding/xml"
	"fmt"

	"github.com/dave/jennifer/jen"
)

type Function struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	CallableAttrs
}

func (f Function) GoName() string {
	return snakeToGo(true, f.Name)
}

func (f Function) GenFunc() *jen.Statement {
	var stmt = new(jen.Statement)
	if f.Doc != nil {
		stmt.Add(f.Doc.GenGoComments("", f.GoName()))
	}

	stmt.Func().Id(f.GoName())

	var parm = []Parameter{}
	if f.Parameters != nil {
		parm = f.Parameters.Parameters
	}
	var args = make(map[string]*jen.Statement, len(parm))

	// Generate the parameters in the function signature.
	stmt.ParamsFunc(func(g *jen.Group) {
		for _, param := range parm {
			if param.IsIgnored() {
				continue
			}

			n := jen.Id(param.GoName())
			args[param.Name] = n

			// Is this an interface? If yes, then treat it specially.
			g.Add(n, param.Type.TypeParam())
		}
	})

	if f.ReturnValue != nil {
		stmt.Add(f.ReturnValue.Type.TypeParam())
	}

	// List of arguments to call the C function. Not to be confused with the
	// above list of arguments to call the current Go function.
	var cargs = make(map[string]*jen.Statement, len(parm))

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

		g.Add(f.ReturnValue.GenReturnFunc(
			jen.Qual("C", f.CIdentifier).ParamsFunc(func(g *jen.Group) {
				for i, param := range parm {
					switch arg, hasCArgument := cargs[param.Name]; {
					case hasCArgument:
						g.Add(arg)

					// Treat UserData and UserDataFreeFunc specially.
					case param.IsUserData():
						if callback := f.UserDataParameter(i); callback != nil {
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

package gir

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Callback struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	CallableAttrs
}

func (c Callback) GoName() string {
	return snakeToGo(true, c.Name)
}

func (c Callback) GenGoType() *jen.Statement {
	var s = new(jen.Statement)
	if c.Doc != nil {
		s.Add(c.Doc.GenGoComments("", c.GoName()))
	}

	s.Type().Id(c.GoName()).Func()

	s.ParamsFunc(func(g *jen.Group) {
		if c.Parameters == nil {
			return
		}

		for _, param := range c.Parameters.Parameters {
			if param.IsIgnored() {
				continue
			}

			g.Add(jen.Id(param.GoName()), param.Type.Type())
		}
	})

	if !c.ReturnValue.IsVoid() {
		s.Add(c.ReturnValue.Type.Type())
	}

	return s
}

// GenGlobalGoFunction generates a Go function with the export comment. This
// function is used to be called from C. It triggers the callback inside the
// map.
func (c Callback) GenGlobalGoFunction() *jen.Statement {
	s := jen.Comment("//export callback" + c.Name)
	s.Line()
	s.Func().Id("callback" + c.Name)

	s.ParamsFunc(func(g *jen.Group) {
		if c.Parameters == nil {
			return
		}

		for _, param := range c.Parameters.Parameters {
			g.Add(jen.Id(param.GoName()), param.Type.GenCGoType())
		}
	})

	if !c.ReturnValue.IsVoid() {
		s.Add(c.ReturnValue.Type.GenCGoType())
	}

	s.BlockFunc(func(g *jen.Group) {
		if c.Parameters == nil {
			return
		}

		// Get the callback closure from the global map, if there's a userData
		// argument.
		var userData = c.Parameters.SearchUserData()
		if userData == nil {
			return
		}

		g.Id("fn").Op(":=").Qual("github.com/diamondburned/handy/internal/callback", "Get").Call(
			jen.Uintptr().Call(jen.Id(userData.GoName())),
		)

		// TODO: is this panic worthy?
		g.If(jen.Id("fn").Op("==").Nil()).Block(
			jen.Panic(jen.Lit(fmt.Sprintf("callback for %s not found", c.Name))),
		)

		g.Line()

		var goargs map[string]*jen.Statement
		if c.Parameters != nil {
			goargs = make(map[string]*jen.Statement, len(c.Parameters.Parameters))
		}

		// Convert C arguments to Go variables.
		for i, param := range c.Parameters.Parameters {
			if param.IsIgnored() {
				continue
			}

			v := jen.Id(fmt.Sprintf("arg%d", i))
			goargs[param.Name] = v

			g.Add(param.Type.GenCaster(v, jen.Id(param.GoName())))
		}

		g.Line()

		var fn = new(jen.Statement)
		if !c.ReturnValue.IsVoid() {
			fn.Id("v").Op(":=")
		}

		g.Add(fn.Id("fn").Assert(jen.Id(c.Name)).CallFunc(func(g *jen.Group) {
			for _, param := range c.Parameters.Parameters {
				if goarg, ok := goargs[param.Name]; ok {
					g.Add(goarg)
				}
			}
		}))

		if !c.ReturnValue.IsVoid() {
			// See if the return value is a pointer. If yes, then we must add a
			// reference: one for Gtk and the other for Go's GC.

			// This check is meant to check for objects, but there isn't a good
			// way to. The commonly known exception to pointers not to an object
			// is *gchar aka utf8, so we filter for that.

			if c.ReturnValue.Type.IsPtr() && c.ReturnValue.Type.Name != "utf8" {
				g.If(jen.Id("v").Op("!=").Nil()).Block(
					jen.Id("v").Dot("Ref").Call(),
				)
			}

			g.Return(c.ReturnValue.Type.GenCCaster(jen.Id("v")))
		}
	})

	return s
}

func (c Callback) ExternCName() string {
	return CallbackExternCName(c.Name)
}

func CallbackExternCName(callbackName string) string {
	return fmt.Sprintf("callback%s", callbackName)
}

func (c Callback) GenExternC() string {
	s := strings.Builder{}
	s.WriteString("extern ")
	s.WriteString(c.ReturnValue.Type.CType)
	s.WriteString(" ")
	s.WriteString(c.ExternCName())

	s.WriteString("(")

	if c.Parameters != nil {
		var params = make([]string, len(c.Parameters.Parameters))
		for i, param := range c.Parameters.Parameters {
			params[i] = fmt.Sprintf("%s v%d", param.Type.CType, i)
		}
		s.WriteString(strings.Join(params, ", "))
	}

	s.WriteString(");")

	return s.String()
}

// TODO: gen assign fn

// GenCGoFunc generates the CGo function to be used in the arguments.
func (c Callback) GenCGoFunc() *jen.Statement {
	return jen.Qual("C", c.ExternCName())
}

// CallbackGenAssign generates a call to callback.Assign with the given fnValue.
func CallbackGenAssign(fnValue *jen.Statement) *jen.Statement {
	return jen.Qual("C", "gpointer").Call(
		jen.Qual("github.com/diamondburned/handy/internal/callback", "Assign").Call(
			fnValue,
		),
	)
}

func CallbackGenDelete() *jen.Statement {
	return ZeroByteCast(jen.Qual("C", "callbackDelete"))
}

// ZeroByteCast works around Go misdetecting C function pointers as (*[0]byte).
// Refer to https://github.com/golang/go/issues/19835.
func ZeroByteCast(caller *jen.Statement) *jen.Statement {
	return jen.Parens(jen.Op("*").Index(jen.Lit(0)).Byte()).Call(caller)
}

package gir

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Namespace struct {
	XMLName            xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 namespace"`
	Version            string   `xml:"version,attr"`
	SharedLibrary      string   `xml:"shared-library,attr"`
	IdentifierPrefixes string   `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	SymbolPrefixes     string   `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`

	Classes     []Class      `xml:"http://www.gtk.org/introspection/core/1.0 class"`
	Enums       []Enum       `xml:"http://www.gtk.org/introspection/core/1.0 enumeration"`
	Functions   []Function   `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Callbacks   []Callback   `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	Interfaces  []Interface  `xml:"http://www.gtk.org/introspection/core/1.0 interface"`
	Annotations []Annotation `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
}

type namespaceGenerator struct {
	*Namespace
}

// FnWithC searches the entire namespace for anything with the given C
// identifier. The returned type mayy be Method, Constructor or Function.
func (n namespaceGenerator) FnWithC(CIdentifier string) interface{} {
	for _, class := range n.Classes {
		if v := class.FnWithC(CIdentifier); v != nil {
			return v
		}
	}

	for _, function := range n.Functions {
		if function.CIdentifier == CIdentifier {
			return function
		}
	}

	return nil
}

func (n namespaceGenerator) FindInterface(ifaceName string) *Interface {
	for _, iface := range n.Interfaces {
		if iface.Name == ifaceName {
			return &iface
		}
	}

	return nil
}

func (n namespaceGenerator) GenerateToFile(f *jen.File) {
	f.CgoPreamble(n.GenCallbackPreamble())
	f.Add(n.GenerateAll())
}

func (n namespaceGenerator) GenCallbackPreamble() string {
	var preambles = make([]string, 0, len(n.Callbacks))
	for _, callback := range n.Callbacks {
		preambles = append(preambles, fmt.Sprintf("// %s", callback.GenExternC()))
	}

	return strings.Join(preambles, "\n")
}

func (n namespaceGenerator) GenerateAll() *jen.Statement {
	f := new(jen.Statement)
	f.Add(n.GenInit())
	f.Add(n.GenEnums())
	f.Add(n.GenInterfaces())
	f.Add(n.GenCallbacks())
	f.Add(n.GenFunctions())
	f.Add(n.GenClasses())
	return f
}

func (n namespaceGenerator) GenEnums() *jen.Statement {
	var f = new(jen.Statement)

	for _, enum := range n.Enums {
		f.Add(enum.GenerateAll())
		f.Line()
	}

	return f
}

func (n namespaceGenerator) GenInterfaces() *jen.Statement {
	var f = new(jen.Statement)

	for _, iface := range n.Interfaces {
		f.Add(iface.GenerateAll())
		f.Line()
	}

	return f
}

func (n namespaceGenerator) GenCallbacks() *jen.Statement {
	var f = new(jen.Statement)

	for _, callback := range n.Callbacks {
		f.Add(callback.GenGoType())
		f.Line()
		f.Add(callback.GenGlobalGoFunction())
		f.Line()
	}

	return f
}

func (n namespaceGenerator) GenFunctions() *jen.Statement {
	var f = new(jen.Statement)

	for _, function := range n.Functions {
		f.Add(function.GenFunc())
		f.Line()
	}

	return f
}

func (n namespaceGenerator) GenClasses() *jen.Statement {
	var f = new(jen.Statement)

	for _, class := range n.Classes {
		f.Add(class.GenerateAll())
		f.Line()
	}

	return f
}

func (n namespaceGenerator) GenInit() *jen.Statement {
	return jen.Func().Id("init").Params().Block(n.genMarshalers()).Line()
}

func (n namespaceGenerator) genMarshalers() *jen.Statement {
	return jen.Qual("github.com/gotk3/gotk3/glib", "RegisterGValueMarshalers").Call(
		jen.Index().Qual("github.com/gotk3/gotk3/glib", "TypeMarshaler").BlockFunc(
			n.genMarshalersList,
		),
	)
}

func (n namespaceGenerator) genMarshalersList(g *jen.Group) {
	g.Comment("Enums")
	for _, enum := range n.Enums {
		g.Add(enum.GenMarshalerItem()).Op(",")
	}

	g.Line()

	g.Comment("Objects/Classes")
	for _, class := range n.Classes {
		g.Add(class.GenMarshalerItem()).Op(",")
	}
}

package iface

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

type Method struct {
	InfoElements  *gir.InfoElements
	InfoAttrs     *gir.InfoAttrs
	CallableAttrs *gir.CallableAttrs

	Recv       string
	Name       string
	Tail       string
	Block      string
	Header     file.Header
	Results    []typeconv.ValueConverted
	ParamDocs  []cmt.ParamDoc
	ReturnDocs []cmt.ParamDoc
}

type ParamDoc struct {
	Name         string
	InfoElements gir.InfoElements
}

func newMethod(cgen *callable.Generator) Method {
	return Method{
		InfoElements:  &cgen.CallableAttrs.InfoElements,
		InfoAttrs:     &cgen.CallableAttrs.InfoAttrs,
		CallableAttrs: cgen.CallableAttrs,

		Recv:       cgen.Recv(),
		Name:       cgen.Name,
		Tail:       callable.CoalesceTail(cgen.Tail),
		Block:      cgen.Block,
		Header:     file.CopyHeader(cgen.Header()),
		Results:    cgen.Results,
		ParamDocs:  cgen.ParamDocs,
		ReturnDocs: cgen.ReturnDocs,
	}
}

type Methods []Method

func (m Methods) find(goName string, this int) bool {
	for i, callable := range m {
		if callable.Name == goName && this != i {
			return true
		}
	}
	return false
}

func (m *Methods) reset(capacity int) {
	if cap(*m) >= capacity {
		*m = (*m)[:0]
	}

	*m = make(Methods, 0, capacity*2)
}

func (m *Methods) setMethods(g *Generator, methods []gir.Method) {
	m.reset(len(methods))
	g.methods = make([]gir.Method, 0, len(methods))

	for i := range methods {
		if types.FilterMethod(g.gen, g.Name, &methods[i]) {
			g.Logln(logger.Debug, "filtered method", methods[i].CIdentifier)
			continue
		}

		g.cgen.Preamble = callable.CallablePreamble

		if !g.cgen.Use(&g.Root, &methods[i].CallableAttrs) {
			g.cgen.Logln(logger.Debug, "setMethods skipped", methods[i].CIdentifier)
			continue
		}

		*m = append(*m, newMethod(&g.cgen))
		g.methods = append(g.methods, methods[i])
	}
}

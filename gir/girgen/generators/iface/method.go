package iface

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

type Method struct {
	InfoElements *gir.InfoElements
	InfoAttrs    *gir.InfoAttrs

	Recv  string
	Name  string
	Tail  string
	Block string
}

func newMethod(cgen *callable.Generator) Method {
	return Method{
		InfoElements: &cgen.InfoElements,
		InfoAttrs:    &cgen.InfoAttrs,

		Recv:  cgen.Recv(),
		Name:  cgen.Name,
		Tail:  callable.CoalesceTail(cgen.Tail),
		Block: cgen.Block,
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

	for i := range methods {
		if types.FilterMethod(g.gen, g.Name, &methods[i]) {
			g.Logln(logger.Debug, "filtered method", methods[i].CIdentifier)
			continue
		}

		if !g.cgen.Use(&g.Root, &methods[i].CallableAttrs) {
			g.cgen.Logln(logger.Debug, "setMethods skipped", methods[i].CIdentifier)
			continue
		}

		g.header.ApplyFrom(g.cgen.Header())
		*m = append(*m, newMethod(&g.cgen))
	}
}

func (m *Methods) setVirtuals(g *Generator, virtuals []gir.VirtualMethod) {
	m.reset(len(virtuals))

	for i := range virtuals {
		if types.FilterSub(g.gen, g.Name, virtuals[i].Name, virtuals[i].CIdentifier) {
			g.Logln(logger.Debug, "filtered method", virtuals[i].CIdentifier)
			continue
		}

		if !g.cgen.Use(&g.Root, &virtuals[i].CallableAttrs) {
			g.cgen.Logln(logger.Debug, "setVirtuals skipped", virtuals[i].CIdentifier)
			continue
		}

		// Don't apply the headers naively. We only import the types, since
		// we're not yet converting these.
		for _, result := range g.cgen.Results {
			if result.NeedsNamespace {
				g.header.ImportPubl(result.Resolved)
			}
		}

		*m = append(*m, newMethod(&g.cgen))
	}
}

type InterfaceImplements struct {
	Name    string
	Type    string
	Wrapper string
}

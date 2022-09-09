package iface

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
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

func (m *Methods) setVirtuals(g *Generator, virtuals []gir.VirtualMethod) {
	m.reset(len(virtuals))
	g.virtuals = make([]gir.VirtualMethod, 0, len(virtuals))

	for i := range virtuals {
		if types.FilterSub(g.gen, g.Name, virtuals[i].Name, virtuals[i].CIdentifier) {
			g.Logln(logger.Debug, "filtered method", virtuals[i].CIdentifier)
			continue
		}

		g.cgen.Preamble = func(cgen *callable.Generator, sect *pen.BlockSection) (string, bool) {
			if g.GLibTypeStruct == nil {
				return "", false
			}

			sect.Linef("gclass := (*C.%s)(coreglib.PeekParentClass(%s))", g.GLibTypeStruct.CType, cgen.Recv())
			sect.Linef("fnarg := gclass.%s", strcases.CGoField(cgen.CallableAttrs.Name))

			// Add the fnarg argumnet.
			g.cgen.ExtraArgs = [2][]string{
				{"unsafe.Pointer(fnarg)"}, // front
				{},                        // back
			}

			// Add the virtual method function type in the form of
			// _gotk4_gtk_Widget_virtual_size_allocate().
			ccall := file.ExportedName(g.gen.Namespace(), g.StructName, "virtual", cgen.CallableAttrs.Name)
			cgen.Header().AddCallbackHeader(types.CgoFuncBridge(g.gen, ccall, cgen.CallableAttrs))

			return "C." + ccall, true
		}

		if !g.cgen.Use(&g.Root, &virtuals[i].CallableAttrs) {
			g.cgen.Logln(logger.Debug, "setVirtuals skipped", virtuals[i].CIdentifier)
			continue
		}

		// Don't apply the headers naively. We only import the types, since
		// we're not yet converting these.
		// for _, result := range g.cgen.Results {
		// 	if result.NeedsNamespace {
		// 		result.Resolved.ImportPubl(g.gen, &g.header)
		// 	}
		// }

		*m = append(*m, newMethod(&g.cgen))
		g.virtuals = append(g.virtuals, virtuals[i])
	}
}

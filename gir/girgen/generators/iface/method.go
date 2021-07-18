package iface

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
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
		Tail:  cgen.Tail,
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

type generateParams struct {
	self   string
	cgen   *callable.Generator
	header *file.Header
	source *gir.NamespaceFindResult
}

func (m *Methods) setMethods(g *Generator, methods []gir.Method) {
	m.reset(len(methods))

	for i := range methods {
		if !g.cgen.Use(&g.Root, &methods[i].CallableAttrs) {
			g.cgen.Logln(logger.Debug, "setMethods skipped", methods[i].CIdentifier)
			continue
		}

		if types.FilterMethod(g.gen, g.Name, &methods[i]) {
			g.cgen.Logln(logger.Debug, "filtered method", methods[i].CIdentifier)
			continue
		}

		g.header.ApplyFrom(g.cgen.Header())
		*m = append(*m, newMethod(&g.cgen))
	}
}

func (m *Methods) setVirtuals(g *Generator, virtuals []gir.VirtualMethod) {
	m.reset(len(virtuals))

	for i := range virtuals {
		if !g.cgen.Use(&g.Root, &virtuals[i].CallableAttrs) {
			g.cgen.Logln(logger.Debug, "setVirtuals skipped", virtuals[i].CIdentifier)
			continue
		}

		// Don't apply the headers naively. We only import the types, since
		// we're not yet converting these.
		for _, result := range g.cgen.Results {
			g.header.ImportPubl(result.Resolved)
		}

		*m = append(*m, newMethod(&g.cgen))
	}
}

type InterfaceImplements struct {
	Name    string
	Type    string
	Wrapper string
}

// // InheritedMethod describes a method inherited from an interface (using
// // the implements thing). The Use function must add imports for the wrappers if
// // needed.
// type InheritedMethod struct {
// 	Method
// 	Parent     string
// 	Wrapper    string
// 	CallParams string
// 	Return     bool
// }

// type InheritedMethods []InheritedMethod

// func (m InheritedMethods) find(goName string) int {
// 	for i, callable := range m {
// 		if callable.Name == goName {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func (m *InheritedMethods) reset() {
// 	*m = (*m)[:0]
// }

// // inheritMethods appends the given interface type's methods into dst. It is
// // optional to fill up p.self.
// func (m *InheritedMethods) add(g *Generator, t *types.Resolved) {
// 	if t.Builtin != nil {
// 		// Skip generating builtin types.
// 		return
// 	}

// 	var methods []gir.Method
// 	switch typ := t.Extern.Type.(type) {
// 	case *gir.Class:
// 		methods = typ.Methods
// 	case *gir.Interface:
// 		methods = typ.Methods
// 	default:
// 		return
// 	}

// 	needsNamespace := t.NeedsNamespace(g.cgen.FileGenerator().Namespace())
// 	parentName := t.PublicType(needsNamespace)
// 	wrapper := t.WrapName(needsNamespace)

// 	for i := range methods {
// 		// Conveniently avoid deprecated methods.
// 		if methods[i].InfoAttrs.Deprecated {
// 			continue
// 		}

// 		if !g.cgen.UseFromNamespace(&methods[i].CallableAttrs, t.Extern.NamespaceFindResult) {
// 			continue
// 		}

// 		// Check for collision and restore the getter name if it's colliding.
// 		if g.hasMethodField(g.cgen.Name) {
// 			g.cgen.RestoreName()
// 		}
// 		// Recheck again. If this is colliding, then ignore.
// 		if g.hasMethodField(g.cgen.Name) {
// 			continue
// 		}

// 		// Import the parameters and returns only; don't import everything in
// 		// the callback generator.
// 		for _, result := range g.cgen.Results {
// 			g.header.ImportPubl(result.Resolved)
// 		}

// 		// Parse the parameter values out of the function in a pretty hacky
// 		// way by extracting the types out.
// 		params := append([]string(nil), g.cgen.GoArgs.Joints()...)
// 		for i, word := range params {
// 			params[i] = strings.SplitN(word, " ", 2)[0]
// 		}

// 		*m = append(*m, InheritedMethod{
// 			Method:     newMethod(&g.cgen),
// 			Parent:     parentName,
// 			Wrapper:    wrapper,
// 			CallParams: strings.Join(params, ", "),
// 			Return:     g.cgen.GoRets.Len() > 0,
// 		})
// 	}
// }

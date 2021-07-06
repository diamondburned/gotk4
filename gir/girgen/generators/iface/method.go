package iface

import (
	"strings"

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

func (m Methods) find(goName string) int {
	for i, callable := range m {
		if callable.Name == goName {
			return i
		}
	}
	return -1
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

func (m *Methods) setMethods(methods []gir.Method, p generateParams) {
	m.reset(len(methods))

	for i := range methods {
		if !p.cgen.UseFromNamespace(&methods[i].CallableAttrs, p.source) {
			p.cgen.Logln(logger.Debug, "setMethods skipped", methods[i].CIdentifier)
			continue
		}

		p.cgen.Header().ApplyHeader(p.header)
		*m = append(*m, newMethod(p.cgen))
	}

	m.renameGetters(p.self)
}

func (m *Methods) setVirtuals(virtuals []gir.VirtualMethod, p generateParams) {
	m.reset(len(virtuals))

	for i := range virtuals {
		if !p.cgen.UseFromNamespace(&virtuals[i].CallableAttrs, p.source) {
			p.cgen.Logln(logger.Debug, "setVirtuals skipped", virtuals[i].CIdentifier)
			continue
		}

		p.cgen.Header().ApplyHeader(p.header)
		*m = append(*m, newMethod(p.cgen))
	}

	m.renameGetters(p.self)
}

func (m Methods) renameGetters(parentName string) {
	for i, call := range m {
		newName, _ := callable.RenameGetter(call.Name)

		// Avoid duplicating method names with Objector.
		// TODO: account for other interfaces as well.
		objectorMethod := parentName != "" && types.IsObjectorMethod(newName)
		if objectorMethod {
			newName += parentName
		}

		if m.find(newName) > -1 {
			if !objectorMethod {
				continue
			}

			// We cannot not rename this method if it's an objectorMethod.
			newName += "_"
		}

		m[i].Name = newName
	}
}

type InterfaceImplements struct {
	Name    string
	Type    string
	Wrapper string
}

// InheritedMethod describes a method inherited from an interface (using
// the implements thing). The Use function must add imports for the wrappers if
// needed.
type InheritedMethod struct {
	Method
	Parent     string
	Wrapper    string
	CallParams string
	Return     bool
}

type InheritedMethods []InheritedMethod

func (m *InheritedMethods) reset() {
	*m = (*m)[:0]
}

// inheritMethods appends the given interface type's methods into dst. It is
// optional to fill up p.self.
func (m *InheritedMethods) add(cgen *callable.Generator, h *file.Header, t *types.Resolved) {
	if t.Builtin != nil {
		// Skip generating builtin types.
		return
	}

	var methods []gir.Method
	switch typ := t.Extern.Type.(type) {
	case *gir.Class:
		methods = typ.Methods
	case *gir.Interface:
		methods = typ.Methods
	default:
		return
	}

	needsNamespace := t.NeedsNamespace(cgen.FileGenerator().Namespace())
	parentName := t.PublicType(needsNamespace)
	wrapper := t.WrapName(needsNamespace)

	for i := range methods {
		if !cgen.UseFromNamespace(&methods[i].CallableAttrs, t.Extern.NamespaceFindResult) {
			continue
		}

		cgen.Header().ApplyHeader(h)

		// Parse the parameter values out of the function in a pretty hacky
		// way by extracting the types out.
		params := append([]string(nil), cgen.GoArgs.Joints()...)
		for i, word := range params {
			params[i] = strings.SplitN(word, " ", 2)[0]
		}

		*m = append(*m, InheritedMethod{
			Method:     newMethod(cgen),
			Parent:     parentName,
			Wrapper:    wrapper,
			CallParams: strings.Join(params, ", "),
			Return:     cgen.GoRets.Len() > 0,
		})
	}
}

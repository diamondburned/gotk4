package iface

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callback"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

type TypeStruct struct {
	*gir.Record
	Methods []TypeStructMethod

	ns   *gir.NamespaceFindResult
	igen *Generator
}

type TypeStructMethod struct {
	Header    file.Header
	FieldName string

	// Go func
	Go *Method

	C struct {
		// Name is the name of the wrapper function.
		Name  string
		Tail  string
		Block string
	}
}

func newTypeStruct(g *Generator, result *gir.TypeFindResult) *TypeStruct {
	record, ok := result.Type.(*gir.Record)
	if !ok {
		g.Logln(logger.Skip, "type-struct skipped since not *gir.Record")
		return nil
	}

	ts := &TypeStruct{
		Record: record,
		ns:     result.NamespaceFindResult,
		igen:   g,
	}

	if !ts.init() {
		return nil
	}

	return ts
}

func (ts *TypeStruct) WrapperFuncName(field string) string {
	return file.ExportedName(ts.ns, ts.Record.Name+"_"+field)
}

func (ts *TypeStruct) init() bool {
	for i := range ts.igen.Virtuals {
		virtual := &ts.igen.virtuals[i]
		vmethod := &ts.igen.Virtuals[i]

		method, ok := ts.newTypeStructMethod(virtual, vmethod)
		if ok {
			ts.Methods = append(ts.Methods, method)
			ts.igen.header.ApplyFrom(&method.Header)
		} else {
			ts.Logln(logger.Skip, "virtual method", virtual.Name)
		}
	}

	// If this is an interface, then we must implement all methods. If any of
	// them cannot be converted, then we failed.
	if ts.igen.IsInterface() && len(ts.igen.virtuals) != len(ts.Methods) {
		return false
	}

	return true
}

func (ts *TypeStruct) newTypeStructMethod(virtual *gir.VirtualMethod, vmethod *Method) (TypeStructMethod, bool) {
	field := ts.findTypeStructField(virtual)
	if field == nil {
		return TypeStructMethod{}, false
	}

	method := TypeStructMethod{
		FieldName: field.Name,
		Go:        vmethod,
	}

	_, isKeyword := strcases.GoKeywords[method.FieldName]
	if isKeyword {
		method.FieldName = "_" + method.FieldName
	}

	method.Header.NeedsExternGLib()
	method.Header.Import("unsafe")

	// TODO: verify first parameter.
	// TODO: method.Go.Tail can be obtained inside the callback generator
	// instead of passing it through a callable generator first.

	// A few side notes:
	//   - Using the callback generator's
	cbgen := callback.NewGenerator(ts.igen.gen)
	cbgen.Parent = ts.igen.Root.Type
	cbgen.Preamble = func(cbgen *callback.Generator, p *pen.BlockSection) (string, bool) {
		p.Linef("goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))")
		if ts.igen.IsClass() {
			p.Linef("iface := goval.(interface{ %s%s })", method.Go.Name, method.Go.Tail)
		} else {
			p.Linef("iface := goval.(%s)", ts.igen.OverriderName())
		}
		return "iface." + method.Go.Name, true
	}

	if !cbgen.Use(&virtual.CallableAttrs) {
		return TypeStructMethod{}, false
	}

	method.C.Block = cbgen.Block
	method.C.Name = file.ExportedName(ts.ns, ts.Record.Name, field.Name)
	method.C.Tail = cbgen.CGoTail
	method.Header.ApplyFrom(cbgen.Header())
	types.AddCallableHeader(types.OverrideNamespace(ts.igen.gen, ts.ns), &method.Header, method.C.Name, vmethod.CallableAttrs)
	// types.AddCallableHeader(ts.igen.gen, &method.Header, method.C.Name, vmethod.CallableAttrs)

	return method, true
}

func (ts *TypeStruct) findTypeStructField(virtual *gir.VirtualMethod) *gir.Field {
	name := virtual.Name
	if virtual.Invoker != "" {
		name = virtual.Invoker
	}

	for i, field := range ts.Record.Fields {
		if field.Name == name {
			return &ts.Record.Fields[i]
		}
	}

	return nil
}

func (ts *TypeStruct) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("typestruct %s (C.%s):", ts.Record.Name, ts.Record.CType)
	ts.igen.Logln(lvl, logger.Prefix(v, p)...)
}

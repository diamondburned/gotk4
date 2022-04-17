package iface

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callback"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
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

	// This commented out snippet makes the code more correct, but:
	//
	//   1. We don't have subclassing support yet, so it's useless to generate
	//      something that requires implementing everything, and
	//   2. This is also used for the return type, but since some methods cannot
	//      be generated in some cases, we still want it to be useful.
	//

	// If this is an interface, then we must implement all methods. If any of
	// them cannot be converted, then we failed.
	// if ts.igen.IsInterface() && len(ts.igen.virtuals) != len(ts.Methods) {
	// 	ts.Logln(logger.Debug, "generated TypeStruct is missing some virtual methods")
	// 	return false
	// }

	return true
}

func (ts *TypeStruct) newTypeStructMethod(virtual *gir.VirtualMethod, vmethod *Method) (TypeStructMethod, bool) {
	field := ts.findTypeStructField(virtual)
	if field == nil {
		ts.Logln(logger.Debug, fmt.Sprintf(
			"cannot find virtual with name %q (invoker %q)",
			virtual.Name, virtual.Invoker,
		))
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
		p.Linef("goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))")
		if ts.igen.IsClass() {
			p.Linef("iface := goval.(interface{ %s%s })", method.Go.Name, method.Go.Tail)
		} else {
			p.Linef("iface := goval.(%s)", ts.igen.OverriderName())
		}
		return "iface." + method.Go.Name, true
	}

	if !cbgen.Use(&virtual.CallableAttrs) {
		ts.Logln(logger.Debug, "virtual", virtual.Name, "cannot be generated")
		return TypeStructMethod{}, false
	}

	method.C.Block = cbgen.Block
	method.C.Name = file.ExportedName(ts.ns, ts.Record.Name, field.Name)
	method.C.Tail = cbgen.CGoTail
	method.Header.ApplyFrom(cbgen.Header())
	method.Header.AddCallable(ts.ns, method.C.Name, vmethod.CallableAttrs)

	return method, true
}

func (ts *TypeStruct) findTypeStructField(virtual *gir.VirtualMethod) *gir.Field {
	var field *gir.Field
	// if virtual.Invoker != "" {
	// 	field = ts.findTypeStructFieldName(virtual.Invoker)
	// }
	if field == nil {
		field = ts.findTypeStructFieldName(virtual.Name)
	}
	return field
}

func (ts *TypeStruct) findTypeStructFieldName(name string) *gir.Field {
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

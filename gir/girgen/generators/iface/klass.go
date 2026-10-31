package iface

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"

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
	Name           string          // TODO
	VirtualMethods []VirtualMethod // TODO: move this out

	ns   *gir.NamespaceFindResult
	igen *Generator
}

type VirtualMethod struct {
	Header    file.Header
	FieldName string

	// Go func
	Go *Method

	C struct {
		// Name is the name of the wrapper function.
		Name   string
		Tail   string
		Block  string
		Header *file.Header
	}
}

func newTypeStruct(g *Generator, result *gir.TypeFindResult) *TypeStruct {
	record, ok := result.Type.(*gir.Record)
	if !ok {
		g.Logln(logger.Skip, "TypeStruct", result.Name, "skipped since not *gir.Record")
		return nil
	}

	// We're usually in the same namespace.
	if !g.gen.CanGenerate(types.TypeFromResult(g.gen, record)) {
		g.Logln(logger.Skip, "TypeStruct", result.Name, "skipped since cannot generate")
		return nil
	}

	ts := &TypeStruct{
		Record: record,
		Name:   strcases.PascalToGo(record.Name),
		ns:     result.NamespaceFindResult,
		igen:   g,
	}

	return ts
}

func (ts *TypeStruct) WrapperFuncName(field string) string {
	return file.ExportedName(ts.ns, ts.Record.Name+"_"+field)
}

func (ts *TypeStruct) Init() bool {
	fields := make(map[string]bool, len(ts.igen.VirtualMethods))

	for i := range ts.igen.VirtualMethods {
		virtual := &ts.igen.virtuals[i]
		vmethod := &ts.igen.VirtualMethods[i]

		method, ok := ts.virtualMethod(virtual, vmethod)
		if !ok {
			ts.Logln(logger.Skip, "virtual method", virtual.Name)
			continue
		}

		// Multiple virtual methods may resolve to the same field, such as
		// GtkEditable's delete_text and do_delete_text.
		if fields[method.FieldName] {
			ts.Logln(logger.Skip, "virtual method", virtual.Name, "has a duplicate field", method.FieldName)
			continue
		}
		fields[method.FieldName] = true

		ts.VirtualMethods = append(ts.VirtualMethods, method)
		ts.igen.header.ApplyFrom(&method.Header)
	}

	return true
}

func (ts *TypeStruct) virtualMethod(virtual *gir.VirtualMethod, vmethod *Method) (VirtualMethod, bool) {
	field := ts.findTypeStructField(virtual)
	if field == nil {
		return VirtualMethod{}, false
	}

	method := VirtualMethod{
		FieldName: field.Name,
		Go:        vmethod,
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
		h := cbgen.Header()
		h.NeedsExternGLib()
		h.Import("unsafe")

		p.Linef("instance0 := coreglib.Take(unsafe.Pointer(arg0))")
		p.Linef("overrides := coreglib.OverridesFromGType[%sOverrides](instance0)", ts.igen.StructName)
		p.Linef("if overrides.%s == nil {", method.Go.Name)
		p.Linef(`  panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected %sOverrides.%s, got none")`, ts.igen.StructName, method.Go.Name)
		p.Linef("}")
		return "overrides." + method.Go.Name, true
	}

	if !cbgen.Use(&virtual.CallableAttrs) {
		return VirtualMethod{}, false
	}

	// The callback must call the Overrides field with the same signature that
	// the field is declared with.
	if !sameTailTypes(cbgen.GoTail, vmethod.Tail) {
		ts.Logln(logger.Debug, "virtual method", virtual.Name, "has mismatching signatures", cbgen.GoTail, "and", vmethod.Tail)
		return VirtualMethod{}, false
	}

	method.C.Block = cbgen.Block
	method.C.Name = file.ExportedName(ts.ns, ts.Record.Name, field.Name)
	method.C.Tail = cbgen.CGoTail
	method.C.Header = cbgen.Header()

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

// sameTailTypes returns true if the two function tails have the same parameter
// and return types, ignoring their names.
func sameTailTypes(tail1, tail2 string) bool {
	types1, ok1 := tailTypes(tail1)
	types2, ok2 := tailTypes(tail2)
	return ok1 && ok2 && types1 == types2
}

func tailTypes(tail string) (string, bool) {
	expr, err := parser.ParseExpr("func" + tail)
	if err != nil {
		return "", false
	}

	fn, ok := expr.(*ast.FuncType)
	if !ok {
		return "", false
	}

	var b strings.Builder
	fset := token.NewFileSet()

	writeFields := func(fields *ast.FieldList) {
		if fields == nil {
			return
		}
		for _, field := range fields.List {
			for range max(1, len(field.Names)) {
				printer.Fprint(&b, fset, field.Type)
				b.WriteString(", ")
			}
		}
	}

	writeFields(fn.Params)
	b.WriteString("| ")
	writeFields(fn.Results)

	return b.String(), true
}

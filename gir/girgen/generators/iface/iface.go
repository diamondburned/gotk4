// Package iface provides an interface generator.
package iface

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

type Generator struct {
	Name         string
	CType        string
	GLibGetType  string
	InfoAttrs    *gir.InfoAttrs
	InfoElements *gir.InfoElements

	InterfaceName   string
	StructName      string
	ParentInterface string

	Implements       []InterfaceImplements
	InheritedMethods InheritedMethods
	Virtuals         Methods // for overrider
	Methods          Methods // for big interface
	Constructors     Methods

	Tree types.Tree

	source *gir.NamespaceFindResult
	header file.Header
	gen    types.FileGenerator
	cgen   callable.Generator
}

var _ = cmt.EnsureInfoFields((*Generator)(nil))

// NewGenerator creates a new interface generator instance.
func NewGenerator(gen types.FileGenerator) Generator {
	return Generator{
		gen:  gen,
		cgen: callable.NewGenerator(gen),
		Tree: types.NewTree(gen),
	}
}

func (g *Generator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("interface %s (C.%s):", g.InterfaceName, g.CType)
	g.gen.Logln(lvl, logger.Prefix(v, p)...)
}

// Reset resets the callback generator.
func (g *Generator) Reset() {
	g.header.Reset()
	g.cgen.Reset()

	g.Tree.Reset()
	g.Methods.reset(0)
	g.Virtuals.reset(0)
	g.Implements = g.Implements[:0]
	g.InheritedMethods.reset()
}

// Header returns the callback generator's current header.
func (g *Generator) Header() *file.Header {
	return &g.header
}

// Use accepts either a *gir.Class or a *gir.Interface; any other type will make
// it panic.
func (g *Generator) Use(typ interface{}) bool {
	g.Reset()

	var methods []gir.Method
	var virtuals []gir.VirtualMethod

	switch typ := typ.(type) {
	case *gir.Class:
		g.Name = typ.Name
		g.CType = typ.CType
		g.GLibGetType = typ.GLibGetType
		g.InfoAttrs = &typ.InfoAttrs
		g.InfoElements = &typ.InfoElements

		g.InterfaceName = strcases.PascalToGo(typ.Name)
		methods = typ.Methods
		virtuals = typ.VirtualMethods

		if !typ.IsIntrospectable() || types.Filter(g.gen, typ.Name, typ.CType) {
			return false
		}

		if !g.Tree.Resolve(typ.Name) {
			g.Logln(logger.Debug, "class cannot be type-resolved")
			return false
		}

		g.ParentInterface = types.GoPublicType(g.gen, g.Tree.Requires[0].Resolved)
		g.header.ImportPubl(g.Tree.Requires[0].Resolved)

		for _, ctor := range typ.Constructors {
			// Copy and bodge this so the constructors and stuff are named properly.
			// This copies things safely, so class is not modified.
			ctor := bodgeClassCtor(typ, ctor)
			if !g.cgen.UseConstructor(&ctor.CallableAttrs) {
				continue
			}

			file.ApplyHeader(g, &g.cgen)
			g.Constructors = append(g.Constructors, newMethod(&g.cgen))
		}

	case *gir.Interface:
		g.Name = typ.Name
		g.CType = typ.CType
		g.GLibGetType = typ.GLibGetType
		g.InfoAttrs = &typ.InfoAttrs
		g.InfoElements = &typ.InfoElements

		g.InterfaceName = strcases.PascalToGo(typ.Name)
		methods = typ.Methods
		virtuals = typ.VirtualMethods

		if !typ.IsIntrospectable() || types.Filter(g.gen, typ.Name, typ.CType) {
			return false
		}

		if !g.Tree.Resolve(typ.Name) {
			g.Logln(logger.Debug, "interface cannot be type-resolved")
			return false
		}

		g.ParentInterface = "gextras.Objector"
		g.header.NeedsExternGLib()
	}

	g.StructName = strcases.UnexportPascal(g.InterfaceName)
	g.source = g.gen.Namespace()
	g.Implements = g.Implements[:0]
	g.header.NeedsExternGLib()

	for _, imp := range g.Tree.WithoutGObject() {
		// Import everything for the embedded types inside the interface.
		g.header.ImportPubl(imp.Resolved)
		// Import gextras for InternObject.
		g.header.ImportCore("gextras")

		needsNamespace := imp.NeedsNamespace(g.gen.Namespace())
		wrapper := imp.WrapName(needsNamespace)

		g.Implements = append(g.Implements, InterfaceImplements{
			Name:    imp.PublicType(false),
			Type:    imp.PublicType(needsNamespace),
			Wrapper: wrapper,
		})
	}

	genParams := generateParams{
		self:   g.InterfaceName,
		cgen:   &g.cgen,
		header: &g.header,
		source: g.source,
	}

	g.Methods.setMethods(methods, genParams)
	g.Virtuals.setVirtuals(virtuals, genParams)

	g.Tree.Walk(func(t *types.Tree, isRoot bool) []types.Tree {
		if !isRoot {
			g.InheritedMethods.add(&g.cgen, &g.header, t.Resolved)
		}

		return t.WithoutGObject()
	})

	return true
}

// Wrap returns a wrapper block that wraps around the given *glib.Object
// variable name.
func (g *Generator) Wrap(obj string) string {
	return g.Tree.WrapInterface(obj)
}

// bodgeClassCtor bodges the given constructor to return exactly the class type
// instead of any other. It returns the original ctor if the conditions don't
// match for bodging.
//
// We have to do this to work around some cases where widget constructors would
// return the widget class instead of the actual class.
func bodgeClassCtor(class *gir.Class, ctor gir.Constructor) gir.Constructor {
	if ctor.ReturnValue == nil || ctor.ReturnValue.Type == nil {
		return ctor
	}

	retVal := *ctor.ReturnValue
	retTyp := *retVal.AnyType.Type

	// Note: this has caused me quite a lot of trouble. It's probably wrong as
	// well. The whole point is to work around the C API's weird class typing.
	retTyp.CType = types.MoveCPtr(class.CType, retTyp.CType)

	retTyp.Name = class.Name
	retTyp.Introspectable = class.Introspectable
	retTyp.AnyType = gir.AnyType{}

	retVal.AnyType.Type = &retTyp
	ctor.ReturnValue = &retVal

	ctor.Name = strings.TrimPrefix(ctor.Name, "new")
	ctor.Name = strings.TrimPrefix(ctor.Name, "_")
	if ctor.Name != "" {
		ctor.Name = "_" + ctor.Name
	}

	ctor.Name = "new_" + class.Name + ctor.Name

	return ctor
}

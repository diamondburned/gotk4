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

	InterfaceName string
	StructName    string

	Virtuals     Methods // for overrider
	Methods      Methods // for big interface
	Constructors Methods

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
	// g.Implements = g.Implements[:0]
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
		g.StructName = g.InterfaceName + "Class"

		methods = typ.Methods
		virtuals = typ.VirtualMethods

		if !typ.IsIntrospectable() || types.Filter(g.gen, typ.Name, typ.CType) {
			return false
		}

		if !g.Tree.Resolve(typ.Name) {
			g.Logln(logger.Debug, "class cannot be type-resolved")
			return false
		}

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
		g.StructName = g.InterfaceName + "Interface"

		methods = typ.Methods
		virtuals = typ.VirtualMethods

		if !typ.IsIntrospectable() || types.Filter(g.gen, typ.Name, typ.CType) {
			return false
		}

		if !g.Tree.Resolve(typ.Name) {
			g.Logln(logger.Debug, "interface cannot be type-resolved")
			return false
		}
	}

	g.source = g.gen.Namespace()
	g.header.NeedsExternGLib()

	g.Methods.setMethods(g, methods)
	g.Virtuals.setVirtuals(g, virtuals)

	g.renameGetters(g.Methods)
	g.renameGetters(g.Virtuals)

	return true
}

func (g *Generator) renameGetters(methods Methods) {
	for i, method := range methods {
		newName, ok := callable.RenameGetter(method.Name)

		if g.hasMethod(newName) || methods.find(newName, i) {
			if ok {
				// Duplicate getter field; just skip renaming.
				continue
			}

			// Field is duplicate even when it wasn't renamed; work around this.
			newName += g.InterfaceName
		}

		methods[i].Name = newName
	}
}

func (g *Generator) hasMethod(name string) bool {
	for _, parent := range g.Tree.Requires {
		if parent.Name() == name {
			return true
		}
	}

	// Never override Object's methods.
	if types.IsObjectorMethod(name) {
		return true
	}

	return false
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

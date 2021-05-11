package girgen

import (
	"github.com/diamondburned/gotk4/gir"
)

var classTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .GoName }}
	type {{ .GoName }} struct {
		{{ index .TypeTree 0 }}
	}
`)

type classGenerator struct {
	gir.Class
	GoName   string
	TypeTree []string

	ng *NamespaceGenerator
}

func (cg *classGenerator) Use(class gir.Class) bool {
	cg.TypeTree = cg.TypeTree[:0]

	// Loop to resolve the parent type, the parent type of that parent type, and
	// so on.
	if class.Parent != "" {
		parent := class.Parent
		for {
			parentType := cg.ng.ResolveTypeName(parent)
			if parentType == nil {
				return false
			}

			goType := parentType.GoType(parentType.NeedsNamespace(cg.ng.current))
			cg.TypeTree = append(cg.TypeTree, goType)

			// We've resolved as deep as we can, so bail. This check works
			// because non-class types don't have parent classes.
			if parentType.Extern == nil || parentType.Extern.Result.Class == nil {
				break
			}

			// Use the parent class' parent type.
			parent = parentType.Extern.Result.Class.Parent
			if parent == "" {
				break
			}
		}
	} else {
		// TODO: check what happens if a class has no parent. It should have a
		// GObject parent, usually.
		return false
	}

	cg.Class = class
	cg.GoName = PascalToGo(class.Name)

	return true
}

func (ng *NamespaceGenerator) generateClasses() {
	cg := classGenerator{
		TypeTree: make([]string, 15),
		ng:       ng,
	}

	for _, class := range ng.current.Namespace.Classes {
		if !cg.Use(class) {
			continue
		}

		ng.pen.BlockTmpl(classTmpl, cg)
	}
}

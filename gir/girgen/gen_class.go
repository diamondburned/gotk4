package girgen

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

var classTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .GoName }}
	type {{ .GoName }} struct {
		{{ $.Ng.GoType (index .TypeTree 0) }}
	}

	func wrap{{ .GoName }}(obj *externglib.Object) *{{ .GoName }} {
		return {{ .Wrap "obj" }}
	}

	func marshal{{ .GoName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return wrapWidget(obj), nil
	}

	{{ range .Constructors }}
	{{ with $tail := ($.CtorCall .CallableAttrs) }}
	func New{{ $.GoName }}{{ $tail }}
	{{ end }}
	{{ end }}
`)

type classGenerator struct {
	gir.Class
	GoName   string
	TypeTree []*ResolvedType

	Ng *NamespaceGenerator
}

func (cg *classGenerator) Use(class gir.Class) bool {
	cg.TypeTree = cg.TypeTree[:0]

	// Loop to resolve the parent type, the parent type of that parent type, and
	// so on.
	if class.Parent != "" {
		parent := class.Parent
		for {
			parentType := cg.Ng.ResolveTypeName(parent)
			if parentType == nil {
				return false
			}

			cg.TypeTree = append(cg.TypeTree, parentType)

			if parentType.Parent == "" {
				break
			}

			// Use the parent class' parent type.
			parent = parentType.Parent
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

// CtorCall generates a FnCall for a constructor.
func (cg *classGenerator) CtorCall(attrs gir.CallableAttrs) string {
	args, ok := cg.Ng.FnArgs(attrs)
	if !ok {
		return ""
	}

	return "(" + args + ") *" + cg.GoName
}

// Wrap returns the wrap string around the given variable name of type
// *glib.Object.
func (cg *classGenerator) Wrap(objName string) string {
	var p pen.Piece
	p.Char('&').Write(cg.GoName).Char('{')

	for _, typ := range cg.TypeTree {
		p.Writef("%s{", typ.GoType(false))
	}

	p.Write(objName)

	for range cg.TypeTree {
		p.Char('}')
	}

	p.Char('}')

	return p.String()
}

func (ng *NamespaceGenerator) generateClasses() {
	cg := classGenerator{
		TypeTree: make([]*ResolvedType, 15),
		Ng:       ng,
	}

	for _, class := range ng.current.Namespace.Classes {
		if !cg.Use(class) {
			continue
		}

		ng.pen.BlockTmpl(classTmpl, &cg)
	}
}

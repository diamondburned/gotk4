package girgen

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

var classTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .InterfaceName }}
	type {{ .InterfaceName }} interface {
		{{ $.Ng.PublicType (index .TypeTree 0) }}
		{{ range .Methods }}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Call -}}
		{{ end }}
	}

	type {{ .StructName }} struct {
		{{ $.Ng.ImplType (index .TypeTree 0) }}
	}

	func wrap{{ .InterfaceName }}(obj *externglib.Object) {{ .InterfaceName }} {
		return {{ .Wrap "obj" }}
	}

	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return wrapWidget(obj), nil
	}

	{{ range .Constructors }}
	{{ with $tail := ($.CtorCall .CallableAttrs) }}
	func New{{ $.InterfaceName }}{{ $tail }}
	{{ end }}
	{{ end }}

	{{ range .Methods }}
	func ({{ FirstLetter $.StructName }} {{ $.StructName }}) {{ .Call }}
	{{ end }}
`)

type classGenerator struct {
	gir.Class
	StructName    string
	InterfaceName string

	TypeTree []*ResolvedType
	Methods  []classMethod

	Ng *NamespaceGenerator
}

type classMethod struct {
	*gir.Method
	Name string
	Call string
}

func (cg *classGenerator) Use(class gir.Class) bool {
	cg.TypeTree = cg.TypeTree[:0]
	cg.Methods = cg.Methods[:0]

	if class.Parent == "" {
		// TODO: check what happens if a class has no parent. It should have a
		// GObject parent, usually.
		return false
	}

	// Loop to resolve the parent type, the parent type of that parent type, and
	// so on.
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

	for i, method := range class.Methods {
		call := cg.Ng.FnCall(method.CallableAttrs)
		if call == "" {
			continue
		}

		name := SnakeToGo(true, method.Name)

		cg.Methods = append(cg.Methods, classMethod{
			Method: &class.Methods[i],
			Name:   name,
			Call:   name + call,
		})
	}

	cg.Class = class
	cg.InterfaceName = PascalToGo(class.Name)
	cg.StructName = UnexportPascal(cg.InterfaceName)

	return true
}

// CtorCall generates a FnCall for a constructor.
func (cg *classGenerator) CtorCall(attrs gir.CallableAttrs) string {
	args, ok := cg.Ng.FnArgs(attrs)
	if !ok {
		return ""
	}

	return "(" + args + ") " + cg.InterfaceName
}

// Wrap returns the wrap string around the given variable name of type
// *glib.Object.
func (cg *classGenerator) Wrap(objName string) string {
	var p pen.Piece
	p.Char('&').Write(cg.StructName).Char('{')

	for _, typ := range cg.TypeTree {
		p.Writef("%s{", cg.Ng.ImplType(typ))
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

		ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
		ng.pen.BlockTmpl(classTmpl, &cg)
	}
}

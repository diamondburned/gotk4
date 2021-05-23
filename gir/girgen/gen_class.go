package girgen

import (
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

var classTmpl = newGoTemplate(`
	{{ GoDoc .Doc 0 .InterfaceName }}
	type {{ .InterfaceName }} interface {
		{{ $.Ng.PublicType (index .TypeTree 0) }}
		{{ range .Methods }}
		{{ GoDoc .Doc 1 .Name }}
		{{ .Name }}{{ .Tail -}}
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
	{{ $tail := ($.CtorCall .CallableAttrs) }}
	{{ if $tail }}
	func {{ $.CtorName . }}{{ $tail }}
	{{ end }}
	{{ end }}

	{{ range .Methods }}
	func ({{ FirstLetter $.StructName }} {{ $.StructName }}) {{ .Name }}{{ .Tail }}
	{{ end }}
`)

type classGenerator struct {
	gir.Class
	StructName    string
	InterfaceName string

	TypeTree []*ResolvedType
	Methods  []callableGenerator

	Ng *NamespaceGenerator
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
			cg.Ng.logln(logWarn, "failed to resolve parent", parent, "for", class.Name)
			return false
		}

		cg.TypeTree = append(cg.TypeTree, parentType)

		if parentType.Parent == "" {
			break
		}

		// Use the parent class' parent type.
		parent = parentType.Parent
	}

	methodNames := make(map[string]struct{}, len(class.Methods))
	for _, method := range class.Methods {
		cbgen := newCallableGenerator(cg.Ng)
		if !cbgen.Use(method.CallableAttrs) {
			continue
		}

		cg.Methods = append(cg.Methods, cbgen)
		methodNames[cbgen.Name] = struct{}{}
	}

	// Use Go-idiomatic getter names, unless there's a duplicate.
	for i, cbgen := range cg.Methods {
		if !strings.HasPrefix(cbgen.Name, "Get") || cbgen.Name == "Get" {
			continue
		}

		newName := strings.TrimPrefix(cbgen.Name, "Get")
		_, dup := methodNames[newName]
		if dup {
			cg.Ng.logln(logInfo, "not renaming cbgen", cbgen.Name, "in class", class.Name)
			continue // skip
		}

		delete(methodNames, cbgen.Name)
		methodNames[newName] = struct{}{}

		cg.Methods[i].Name = newName
	}

	cg.Class = class
	cg.InterfaceName = PascalToGo(class.Name)
	cg.StructName = UnexportPascal(cg.InterfaceName)

	return true
}

func (cg *classGenerator) CtorName(ctor gir.Constructor) string {
	name := SnakeToGo(true, ctor.Name)
	name = strings.TrimPrefix(name, "New")
	return "New" + cg.InterfaceName + name
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
	p.Write(cg.StructName).Char('{')

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
			ng.logln(logInfo, "skipping class", class.Name)
			continue
		}

		ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
		ng.pen.BlockTmpl(classTmpl, &cg)
	}
}

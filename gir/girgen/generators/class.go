package generators

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var classTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	type {{ .InterfaceName }} interface {
		{{ .ParentInterface }}

		{{ range .Implements }}
		// As{{.Name}} casts the class to the {{.Type}} interface.
		As{{.Name}}() {{ .Type -}}
		{{ end }}

		{{ range .Methods }}
		{{ GoDoc . 1 }}
		{{ .Name }}{{ .Tail -}}
		{{ end }}
	}

	// {{ .StructName }} implements the {{ .InterfaceName }} class.
	type {{ .StructName }} struct {
		{{ .ParentInterface }}
	}

	// Wrap{{ .InterfaceName }} wraps a GObject to the right type. It is
	// primarily used internally.
	func Wrap{{ .InterfaceName }}(obj *externglib.Object) {{ .InterfaceName }} {
		return {{ .Tree.WrapClass "obj" }}
	}

	{{ if .GLibGetType }}
	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
		obj := externglib.Take(unsafe.Pointer(val))
		return Wrap{{ .InterfaceName }}(obj), nil
	}
	{{ end }}

	{{ range .Constructors }}
	{{ GoDoc . 0 }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}

	{{ range .Methods }}
	func ({{ .Recv }} {{ $.StructName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}

	{{ $recv := (FirstLetter .StructName) }}
	{{ range .Implements }}
	func ({{ $recv }} {{ $.StructName }}) As{{.Name}}() {{ .Type }} {
		return {{ .Wrapper }}(gextras.InternObject({{ $recv }}))
	}
	{{ end }}
`)

// GenerateClass generates the given class into files.
func GenerateClass(gen FileGeneratorWriter, class *gir.Class) bool {
	classGen := NewClassGenerator(gen)
	if !classGen.Use(class) {
		return false
	}

	writer := FileWriterFromType(gen, class)

	if class.GLibGetType != "" && !types.FilterCType(gen, class.GLibGetType) {
		writer.Header().AddMarshaler(class.GLibGetType, classGen.InterfaceName)
	}

	// Need for the object wrapper.
	writer.Header().NeedsExternGLib()

	file.ApplyHeader(writer, &classGen)
	writer.Pen().WriteTmpl(classTmpl, &classGen)
	return true
}

type classCallable struct {
	InfoElements *gir.InfoElements
	InfoAttrs    *gir.InfoAttrs

	Recv  string
	Name  string
	Tail  string
	Block string
}

type classImplements struct {
	Name    string
	Type    string
	Wrapper string
}

func newClassCallable(cgen *callable.Generator) classCallable {
	return classCallable{
		InfoElements: &cgen.InfoElements,
		InfoAttrs:    &cgen.InfoAttrs,

		Recv:  cgen.Recv(),
		Name:  cgen.Name,
		Tail:  cgen.Tail,
		Block: cgen.Block,
	}
}

type ClassGenerator struct {
	*gir.Class
	StructName      string
	InterfaceName   string
	ParentInterface string

	Implements   []classImplements
	Constructors []classCallable
	Methods      []classCallable

	Tree types.Tree // starts from current resolved type

	header file.Header
	gen    FileGenerator
	cgen   callable.Generator
}

// NewClassGenerator creates a new class generator.
func NewClassGenerator(gen FileGenerator) ClassGenerator {
	classGenerator := ClassGenerator{
		gen:  gen,
		Tree: types.NewTree(gen),
	}
	classGenerator.cgen = callable.NewGenerator(headeredFileGenerator{
		FileGenerator: gen,
		Headerer:      &classGenerator,
	})

	return classGenerator
}

func (cg *ClassGenerator) Logln(lvl logger.Level, v ...interface{}) {
	p := fmt.Sprintf("class %s (C.%s):", cg.InterfaceName, cg.CType)
	cg.gen.Logln(lvl, logger.Prefix(v, p)...)
}

// Reset resets the callback generator.
func (g *ClassGenerator) Reset() {
	g.cgen.Reset()
	g.Tree.Reset()

	*g = ClassGenerator{
		gen:  g.gen,
		cgen: g.cgen,
		Tree: g.Tree,
	}
}

// Header returns the callback generator's current header.
func (g *ClassGenerator) Header() *file.Header {
	return &g.header
}

func (cg *ClassGenerator) Use(class *gir.Class) bool {
	cg.Reset()
	cg.Class = class
	cg.InterfaceName = strcases.PascalToGo(class.Name)
	cg.StructName = strcases.UnexportPascal(cg.InterfaceName)

	if !class.IsIntrospectable() || types.Filter(cg.gen, class.Name, class.CType) {
		return false
	}

	// Skip generating built-in types as well as orphaned types.
	if !cg.Tree.Resolve(class.Name) || cg.Tree.Builtin != nil {
		cg.Logln(logger.Debug, "because unknown parent type", class.Parent)
		return false
	}

	cg.ParentInterface = types.GoPublicType(cg.gen, cg.Tree.Requires[0].Resolved)
	implements := cg.Tree.Requires[1:]

	// Add imports for the parent class and implemented interfaces.
	for _, imp := range cg.Tree.Requires {
		cg.header.ImportResolvedType(imp.PublImport)
	}

	cg.Methods = classCallableGrow(cg.Methods, len(class.Methods))
	cg.Constructors = classCallableGrow(cg.Constructors, len(class.Constructors))
	cg.Implements = cg.Implements[:0]

	cg.cgen.Constructor = true

	for _, ctor := range class.Constructors {
		// Copy and bodge this so the constructors and stuff are named properly.
		// This copies things safely, so class is not modified.
		ctor := bodgeClassCtor(class, ctor)
		if !cg.cgen.Use(&ctor.CallableAttrs) {
			continue
		}

		file.ApplyHeader(cg, &cg.cgen)
		cg.Constructors = append(cg.Constructors, newClassCallable(&cg.cgen))
	}

	cg.cgen.Constructor = false

	for _, impl := range implements {
		needsNamespace := impl.NeedsNamespace(cg.gen.Namespace())
		wrapper := impl.WrapName(needsNamespace)

		cg.Implements = append(cg.Implements, classImplements{
			Name:    impl.PublicType(false),
			Type:    impl.PublicType(needsNamespace),
			Wrapper: wrapper,
		})
	}

	if len(cg.Implements) > 0 {
		// Import gextras for InternObject.
		cg.header.ImportCore("gextras")
	}

	// Only generate the methods after we've generated the inherited methods,
	// because we want the class to satisfy the inherited interfaces as a first
	// priority, and then generate our methods down below. This way, we can also
	// avoid method collisions.
	for i := range class.Methods {
		method := &class.Methods[i]

		if types.FilterMethod(cg.gen, cg.Class.Name, method) {
			continue
		}

		if !cg.cgen.Use(&method.CallableAttrs) {
			continue
		}

		file.ApplyHeader(cg, &cg.cgen)
		cg.Methods = append(cg.Methods, newClassCallable(&cg.cgen))
	}

	// Rename all methods to have idiomatic and non-colliding names, if possible.
	for i, method := range cg.Methods {
		newName, isGetter := callable.RenameGetter(method.Name)
		isDuplicate := cg.hasField(newName)

		// Avoid duplicating field names with inherited interfaces, including
		// Objector, but only this if we're not renaming a getter (since that's
		// not important).
		if isDuplicate && !isGetter {
			newName += cg.InterfaceName
			isDuplicate = cg.hasField(newName)
		}

		if !isDuplicate {
			cg.Methods[i].Name = newName
		}
	}

	return true
}

func (cg *ClassGenerator) hasField(goName string) bool {
	if types.IsObjectorMethod(goName) || cg.ParentInterface == goName {
		return true
	}
	for _, impl := range cg.Implements {
		if "As"+impl.Name == goName {
			return true
		}
	}
	for _, callable := range cg.Methods {
		if callable.Name == goName {
			return true
		}
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

func classCallableGrow(callables []classCallable, n int) []classCallable {
	if cap(callables) <= n {
		return callables[:0]
	}
	return make([]classCallable, 0, n*2)
}

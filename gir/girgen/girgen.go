package girgen

import (
	"fmt"
	"io"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
	"github.com/fatih/color"
)

func newGoTemplate(block string) *template.Template {
	_, file, _, _ := runtime.Caller(1)
	base := filepath.Base(file)

	t := template.New(base)
	t.Funcs(template.FuncMap{
		"PascalToGo":     PascalToGo,
		"UnexportPascal": UnexportPascal,
		"SnakeToGo":      SnakeToGo,
		"FirstLetter":    FirstLetter,
		"GoDoc":          GoDoc,
	})
	t = template.Must(t.Parse(block))
	return t
}

// Generator is a big generator that manages multiple repositories.
type Generator struct {
	// KnownTypes contains a list of type checks that return true if the given
	// type matches a known type.
	KnownTypes []func(string) bool
	// NoMarshalPkgs contains a list of namespace package names that don't have
	// marshalers generated. GLib is by default in this list, because it
	// shouldn't have any marshalers.
	NoMarshalPkgs []string

	Repos      gir.Repositories
	RootModule string

	logger *log.Logger
	color  bool
}

// NewGenerator creates a new generator with sane defaults.
func NewGenerator(repos gir.Repositories, root string) *Generator {
	return &Generator{
		Repos:      repos,
		RootModule: root,
	}
}

// WithLogger sets the generator's logger.
func (g *Generator) WithLogger(logger *log.Logger, color bool) {
	g.logger = logger
	g.color = color
}

type logLevel uint8

const (
	logInfo logLevel = iota
	logSummary
	logWarn
	logError
)

func (lvl logLevel) prefix() string {
	switch lvl {
	case logInfo:
		return "info:"
	case logSummary:
		return "summary:"
	case logWarn:
		return "warning:"
	case logError:
		return "error:"
	default:
		return ""
	}
}

func (lvl logLevel) colorf(f string, v ...interface{}) string {
	switch lvl {
	case logInfo:
		return color.HiBlueString(f, v...)
	case logSummary:
		return color.HiGreenString(f, v...)
	case logWarn:
		return color.HiYellowString(f, v...)
	case logError:
		return color.HiRedString(f, v...)
	default:
		return fmt.Sprintf(f, v...)
	}
}

func (g *Generator) logln(level logLevel, v ...interface{}) {
	if g.logger == nil {
		return
	}

	prefix := level.prefix()
	if g.color {
		prefix = level.colorf(prefix)
	}

	v = append([]interface{}{prefix}, v...)
	g.logger.Println(v...)
}

func (g *Generator) ImportPath(pkgPath string) string {
	return gir.ImportPath(g.RootModule, pkgPath)
}

// UseNamespace creates a new namespace generator using the given namespace.
func (g *Generator) UseNamespace(namespace string) *NamespaceGenerator {
	res := g.Repos.FindNamespace(namespace)
	if res == nil {
		return nil
	}

	return &NamespaceGenerator{
		pen: pen.NewPaper(),          // 4MB
		pre: pen.NewPaperSize(10240), // 10KB
		cgo: pen.NewPaperSize(10240), // 10KB

		imports: map[string]string{},
		pkgPath: g.ImportPath(gir.GoNamespace(res.Namespace)),
		gen:     g,
		current: res,
	}
}

// ignoreIxs is a map of indexes to ignore in the function signature.
type ignoreIxs map[int]struct{}

func (ig *ignoreIxs) init() {
	if *ig == nil {
		*ig = map[int]struct{}{}
	}
}

func (ig *ignoreIxs) set(i int) {
	ig.init()
	(*ig)[i] = struct{}{}
}

func (ig *ignoreIxs) fieldIgnore(field gir.Field) {
	ig.typeIgnore(field.AnyType)
}

func (ig *ignoreIxs) paramsIgnore(params *gir.Parameters) {
	if params == nil {
		return
	}
	for _, param := range params.Parameters {
		ig.paramIgnore(param)
	}
}

func (ig *ignoreIxs) paramIgnore(param gir.Parameter) {
	if param.Closure != nil {
		ig.set(*param.Closure)
	}
	if param.Destroy != nil {
		ig.set(*param.Destroy)
	}

	ig.typeIgnore(param.AnyType)
}

func (ig *ignoreIxs) typeIgnore(typ gir.AnyType) {
	if typ.Array != nil {
		if typ.Array.Length != nil {
			ig.set(*typ.Array.Length)
		}
	}
}

func (ig ignoreIxs) ignore(i int) bool {
	_, ignore := ig[i]
	return ignore
}

// NamespaceGenerator is a generator for a specific namespace.
type NamespaceGenerator struct {
	pen *pen.Paper // body
	pre *pen.Paper
	cgo *pen.Paper

	pkgPath string            // package name
	imports map[string]string // optional alias value

	gen     *Generator
	current *gir.NamespaceFindResult

	// inserted keeps track of what was inserted once.
	inserted struct {
		CallbackDelete bool
	}
}

// Generate generates the current namespace into the given writer.
func (ng *NamespaceGenerator) Generate(w io.Writer) error {
	pkgs := []string{"#cgo pkg-config:"}
	for _, pkg := range ng.current.Repository.Packages {
		pkgs = append(pkgs, pkg.Name)
	}

	ng.cgo.Words(pkgs...)
	ng.cgo.Words("#cgo CFLAGS: -Wno-deprecated-declarations")
	for _, cIncl := range ng.current.Repository.CIncludes {
		ng.cgo.Wordf("#include <%s>", cIncl.Name)
	}
	ng.cgo.Line()

	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	ng.generateInit()
	ng.generateAliases()
	ng.generateEnums()
	ng.generateBitfields()
	ng.generateCallbacks()
	ng.generateFuncs()
	ng.generateIfaces()
	ng.generateRecords()
	ng.generateClasses()

	if err := pen.Flush(ng.pen, ng.cgo, ng.pre); err != nil {
		ng.logln(logError, "generation error:", err)
		return err
	}

	pen := pen.New(w)
	pen.Words("// Code generated by girgen. DO NOT EDIT.")
	pen.Line()

	pen.Words("package", ng.PackageName())
	pen.Line()

	if len(ng.imports) > 0 {
		pen.Words("import (")
		for imp, alias := range ng.imports {
			// Only use the import alias if it's provided and does not match the
			// base name of the import path for idiomaticity.
			if alias != "" && alias != path.Base(imp) {
				pen.Words(alias, "", strconv.Quote(imp))
			} else {
				pen.Words(strconv.Quote(imp))
			}
		}
		pen.Block(")")
		pen.Line()
	}

	for _, line := range strings.Split(ng.cgo.String(), "\n") {
		pen.Words("//", line)
	}
	pen.Words(`import "C"`)
	pen.Line()

	pen.WriteString(ng.pre.String())
	pen.WriteString(ng.pen.String())

	if err := pen.Flush(); err != nil {
		ng.logln(logError, "final file write error:", err)
		return err
	}

	return nil
}

func (ng *NamespaceGenerator) needsCallbackDelete() {
	if ng.inserted.CallbackDelete {
		return
	}

	ng.inserted.CallbackDelete = true
	ng.addImport("github.com/diamondburned/gotk4/internal/box")

	ng.cgo.Words("// extern void callbackDelete(gpointer);")

	ng.pre.Words("//export callbackDelete")
	ng.pre.Words("func callbackDelete(ptr C.gpointer) {")
	ng.pre.Words("  box.Delete(box.Callback, uintptr(ptr))")
	ng.pre.Words("}")

	ng.pre.Line()
}

// fullGIR returns the full GIR type name if it doesn't contain a namespace.
func (ng *NamespaceGenerator) fullGIR(girType string) string {
	// Skip builtin types.
	_, isBuiltin := girToBuiltin[girType]
	if isBuiltin {
		return girType
	}

	if !strings.Contains(girType, ".") {
		return ng.current.Namespace.Name + "." + girType
	}
	return girType
}

// PackageName returns the current namespace's package name.
func (ng *NamespaceGenerator) PackageName() string {
	return gir.GoNamespace(ng.current.Namespace)
}

// Namespace returns the generator's namespace.
func (ng *NamespaceGenerator) Namespace() *gir.Namespace {
	return ng.current.Namespace
}

// Repository returns the generator's repository.
func (ng *NamespaceGenerator) Repository() *gir.PkgRepository {
	return ng.current.Repository
}

func (ng *NamespaceGenerator) logln(level logLevel, v ...interface{}) {
	prefix := []interface{}{"package", ng.current.Namespace.Name + ":"}
	prefix = append(prefix, v...)

	ng.gen.logln(level, prefix...)
}

func (ng *NamespaceGenerator) warnUnknownType(typ string) {
	ng.logln(logWarn, "unknown gir type", strconv.Quote(typ))
}

func (ng *NamespaceGenerator) addImport(pkgPath string) {
	ng.addImportAlias(pkgPath, "")
}

func (ng *NamespaceGenerator) addImportAlias(pkgPath, alias string) {
	_, ok := ng.imports[pkgPath]
	if ok {
		return
	}

	ng.imports[pkgPath] = alias
}

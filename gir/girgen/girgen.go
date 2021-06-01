package girgen

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
	"github.com/fatih/color"
	"github.com/pkg/errors"
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
	Repos   gir.Repositories
	ModPath ModulePathFunc
	Filters []FilterMatcher

	logger *log.Logger
	color  bool
}

// ModulePathFunc returns the Go module import path from the given namespace.
type ModulePathFunc func(res *gir.NamespaceFindResult) string

// NewGenerator creates a new generator with sane defaults.
func NewGenerator(repos gir.Repositories, modPath ModulePathFunc) *Generator {
	return &Generator{
		Repos:   repos,
		ModPath: modPath,
		Filters: []FilterMatcher{
			// These are already manually covered in the girgen code; they are
			// provided by package gotk3/glib.
			AbsoluteFilter("GObject.Type"),
			AbsoluteFilter("GObject.Value"),
			AbsoluteFilter("GObject.Object"),
			AbsoluteFilter("GObject.InitiallyUnowned"),
		},
	}
}

// AddFilters adds the given list of filters.
func (g *Generator) AddFilters(filters []FilterMatcher) {
	g.Filters = append(g.Filters, filters...)
}

// WithLogger sets the generator's logger.
func (g *Generator) WithLogger(logger *log.Logger, color bool) {
	g.logger = logger
	g.color = color
}

type logLevel uint8

const (
	logDebug logLevel = iota
	// logUnsupported is used for logs that report conditions impossible to do
	// in Go properly.
	logUnsupported
	// logUnknown is reserved for logging down unknown types or types that
	// cannot be resolved.
	logUnknown
	// logSkip is reserved for logging down skipped types.
	logSkip
	// logUnusuality is reserved for logging down unexpected GIR values or
	// formats. It may be used to log things yet to be supported but can be.
	logUnusuality
	// logError is reserved for fatal and/or unexpected errors.
	logError
)

func (lvl logLevel) prefix() string {
	switch lvl {
	case logDebug:
		return "debug:"
	case logUnsupported:
		return "unsupported:"
	case logUnknown:
		return "unknown type:"
	case logSkip:
		return "skipped:"
	case logUnusuality:
		return "unusuality:"
	case logError:
		return "error:"
	default:
		return ""
	}
}

func (lvl logLevel) colorf(f string, v ...interface{}) string {
	switch lvl {
	case logUnsupported:
		return color.YellowString(f, v...)
	case logUnknown:
		return color.BlueString(f, v...)
	case logSkip:
		return color.GreenString(f, v...)
	case logUnusuality:
		return color.RedString(f, v...)
	case logError:
		return color.New(color.Bold, color.FgHiRed).Sprintf(f, v...)
	case logDebug:
		fallthrough
	default:
		return color.New(color.Faint).Sprintf(f, v...)
	}
}

func (g *Generator) logln(level logLevel, v ...interface{}) {
	if g.logger == nil {
		return
	}

	prefix := level.prefix()
	if prefix != "" {
		if g.color {
			prefix = level.colorf(prefix)
		}
		v = append([]interface{}{prefix}, v...)
	}

	g.logger.Println(v...)
}

// UseNamespace creates a new namespace generator using the given namespace.
func (g *Generator) UseNamespace(namespace, version string) *NamespaceGenerator {
	res := g.Repos.FindNamespace(gir.VersionedName(namespace, version))
	if res == nil {
		return nil
	}

	return &NamespaceGenerator{
		pen: pen.NewPaper(),          // 4MB
		pre: pen.NewPaperSize(10240), // 10KB
		cgo: pen.NewPaperSize(10240), // 10KB

		imports: map[string]string{},
		pkgPath: g.ModPath(res),
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

	pkgPath    string            // package name
	imports    map[string]string // optional alias value
	marshalers []string

	gen     *Generator
	current *gir.NamespaceFindResult

	// inserted keeps track of what was inserted once.
	inserted struct {
		CallbackDelete bool
	}
}

// Generate generates the current namespace into the given writer.
func (ng *NamespaceGenerator) Generate() ([]byte, error) {
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

	ng.marshalers = make([]string, 0, 0+
		len(ng.current.Namespace.Enums)+
		len(ng.current.Namespace.Bitfields)+
		len(ng.current.Namespace.Records)+
		len(ng.current.Namespace.Classes)+
		len(ng.current.Namespace.Interfaces),
	)

	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	ng.generateAliases()
	ng.generateEnums()
	ng.generateBitfields()
	ng.generateCallbacks()
	ng.generateFuncs()
	ng.generateIfaces()
	ng.generateRecords()
	ng.generateClasses()

	if len(ng.marshalers) > 0 {
		ng.addImportAlias("github.com/gotk3/gotk3/glib", "externglib")
	}

	pen.Flush(ng.pen, ng.cgo, ng.pre)

	var out bytes.Buffer
	// Preallocate 10KB + existing buffers.
	out.Grow(10*1024 + ng.pen.Size() + ng.cgo.Size() + ng.pre.Size())

	pen := pen.New(&out)
	pen.Words("// Code generated by girgen. DO NOT EDIT.")
	pen.Line()

	pen.Words("package", ng.PackageName())
	pen.Line()

	if len(ng.imports) > 0 {
		builtin := make([]string, len(ng.imports))
		externs := make([]string, len(ng.imports))

		for path, alias := range ng.imports {
			if !strings.Contains(path, "/") {
				builtin = append(builtin, makeImport(path, alias))
			} else {
				externs = append(externs, makeImport(path, alias))
			}
		}

		sort.Strings(builtin)
		sort.Strings(externs)

		pen.Words("import (")

		for _, str := range builtin {
			pen.Words(str)
		}
		if len(builtin) > 0 && len(externs) > 0 {
			pen.Line()
		}
		for _, str := range externs {
			pen.Words(str)
		}

		pen.Block(")")
		pen.Line()
	}

	for _, line := range strings.Split(ng.cgo.String(), "\n") {
		pen.Words("//", line)
	}
	pen.Words(`import "C"`)
	pen.Line()

	if len(ng.marshalers) > 0 {
		pen.Words("func init() {")
		pen.Words("  externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{")

		for _, marshaler := range ng.marshalers {
			pen.Words("      " + marshaler)
		}

		pen.Words("  })")
		pen.Words("}")
		pen.Line()
	}

	pen.WriteString(ng.pre.String())
	pen.WriteString(ng.pen.String())

	pen.Flush()

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		ng.logln(logError, "failed to fmt final source code")
		return out.Bytes(), errors.Wrap(err, "failed to fmt")
	}

	return formatted, nil
}

func makeImport(importPath, alias string) string {
	pathBase := path.Base(importPath)

	// Check if the base is a version part.
	if strings.HasPrefix(pathBase, "v") {
		_, err := strconv.Atoi(strings.TrimPrefix(pathBase, "v"))
		if err == nil {
			// Valid version part. Trim it.
			importPath = path.Dir(importPath)
		}
	}

	if alias == "" || alias == path.Base(importPath) {
		return strconv.Quote(importPath)
	}

	// Only use the import alias if it's provided and does not match the base
	// name of the import path for idiomaticity.
	return alias + " " + strconv.Quote(importPath)
}

func (ng *NamespaceGenerator) addMarshaler(glibGetType, goName string) {
	ng.marshalers = append(ng.marshalers, fmt.Sprintf(
		`{T: externglib.Type(C.%s()), F: marshal%s},`, glibGetType, goName,
	))
}

// mustIgnore checks the generator's filters to see if the given girType in this
// namespace should be ignored.
func (ng *NamespaceGenerator) mustIgnore(girType, cType string) (ignore bool) {
	for _, filter := range ng.gen.Filters {
		if !filter.Filter(ng, girType, cType) {
			// Filter returns keep=false.
			return true
		}
	}

	return false
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
	ng.logln(logUnknown, strconv.Quote(typ))
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

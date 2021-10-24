package types

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
)

// TODO: refactor to add method accuracy

// Preprocessor describes something that can preprocess anything in the given
// list of repositories. This is useful for renaming functions, classes or
// anything else.
type Preprocessor interface {
	// Preprocess goes over the given list of repos, changing what's necessary.
	Preprocess(repos gir.Repositories)
}

// ApplyPreprocessors applies the given list of preprocessors onto the given
// list of GIR repositories.
func ApplyPreprocessors(repos gir.Repositories, preprocs []Preprocessor) {
	for _, preproc := range preprocs {
		preproc.Preprocess(repos)
	}
}

// PreprocessorFunc is a helper function to satisfy the Preprocessor interface.
type PreprocessorFunc func(gir.Repositories)

// Preprocess calls f.
func (f PreprocessorFunc) Preprocess(repos gir.Repositories) {
	f(repos)
}

// RemovePkgconfig removes the given pkgconfig packages from the given
// repository.
func RemovePkgconfig(girFile string, pkgs ...string) Preprocessor {
	matchers := makeMatchers(pkgs)

	return PreprocessorFunc(func(repos gir.Repositories) {
		repo := repos.FromGIRFile(girFile)
		if repo == nil {
			log.Printf("RemoveCIncludes: gir file %s not found", girFile)
			return
		}

		packages := repo.Packages[:0]

	findPkg:
		for _, pkg := range repo.Packages {
			for _, match := range matchers {
				if match(pkg.Name) {
					continue findPkg
				}
			}
			packages = append(packages, pkg)
		}

		repo.Packages = packages
	})
}

// RemoveCIncludes removes the given C includes from the given repository.
func RemoveCIncludes(girFile string, cincls ...string) Preprocessor {
	matchers := makeMatchers(cincls)

	return PreprocessorFunc(func(repos gir.Repositories) {
		repo := repos.FromGIRFile(girFile)
		if repo == nil {
			log.Printf("RemoveCIncludes: gir file %s not found", girFile)
			return
		}

		cIncludes := repo.CIncludes[:0]

	findCIncl:
		for _, cincl := range repo.CIncludes {
			for _, match := range matchers {
				if match(cincl.Name) {
					continue findCIncl
				}
			}
			cIncludes = append(cIncludes, cincl)
		}

		repo.CIncludes = cIncludes
	})
}

func makeMatchers(inputs []string) []func(string) bool {
	fns := make([]func(string) bool, len(inputs))

	for i, input := range inputs {
		if strings.HasPrefix(input, "/") && strings.HasSuffix(input, "/") {
			input = strings.Trim(input, "/")
			re := regexp.MustCompile(input)
			fns[i] = re.MatchString
		} else {
			input := input
			fns[i] = func(v string) bool { return v == input }
		}
	}

	return fns
}

func girTypeMustBeVersioned(girType string) {
	namespace, _ := gir.SplitGIRType(girType)

	// Verify that the namespace is present.
	_, version := gir.ParseVersionName(namespace)
	if version == "" {
		log.Panicf("girType %q missing version", girType)
	}
}

// PreserveGetName matches a type and prepends "get_" or "Get" into it to
// preserve the getter name in case of collision.
func PreserveGetName(girType string) Preprocessor {
	girTypeMustBeVersioned(girType)
	return PreprocessorFunc(func(repos gir.Repositories) {
		result := repos.FindFullType(girType)
		if result == nil {
			log.Printf("GIR type %q not found", girType)
		}

		if name := result.Name(); strcases.GuessSnake(name) {
			result.SetName("get_" + name)
		} else {
			result.SetName("Get" + name)
		}
	})
}

// RenameEnumMembers renames all members of the matched enums. It is primarily
// used to avoid collisions.
func RenameEnumMembers(enum, regex, replace string) Preprocessor {
	girTypeMustBeVersioned(enum)
	re := regexp.MustCompile(regex)
	return PreprocessorFunc(func(repos gir.Repositories) {
		result := repos.FindFullType(enum)
		if result == nil {
			log.Printf("GIR enum %q not found", enum)
			return
		}

		enum, ok := result.Type.(*gir.Enum)
		if !ok {
			log.Panicf("GIR type %T is not enum", result.Type)
		}

		for i, member := range enum.Members {
			parts := strings.SplitN(member.CIdentifier, "_", 2)
			parts[1] = re.ReplaceAllString(parts[1], replace)
			enum.Members[i].CIdentifier = parts[0] + "_" + parts[1]
		}
	})
}

type typeRenamer struct {
	from, to string
}

// TypeRenamer creates a new filter matcher that renames a type. The given GIR
// type must contain the versioned namespace, like "Gtk3.Widget" but the given
// name must not. The GIR type is absolutely matched, similarly to
// AbsoluteFilter.
func TypeRenamer(girType, newName string) Preprocessor {
	girTypeMustBeVersioned(girType)
	return typeRenamer{
		from: girType,
		to:   newName,
	}
}

func (ren typeRenamer) Preprocess(repos gir.Repositories) {
	result := repos.FindFullType(ren.from)
	if result == nil {
		log.Printf("GIR type %q not found", ren.from)
		return
	}

	oldName := result.Name()
	result.SetName(ren.to)

	if info := cmt.GetInfoFields(result.Type); info.Elements != nil {
		changedMsg := fmt.Sprintf("This type has been renamed from %s.", oldName)
		if info.Elements.Doc != nil {
			info.Elements.Doc.String += "\n\n" + changedMsg
		}
	}
}

// RemoveRecordFields removes the given fields from the record with the given
// full GIR type. The fields must be cased as they appear in the GIR file.
func RemoveRecordFields(girType string, fields ...string) Preprocessor {
	return PreprocessorFunc(func(repos gir.Repositories) {
		res := repos.FindFullType(girType)
		if res == nil {
			log.Printf("GIR type %q not found", girType)
			return
		}

		record, ok := res.Type.(*gir.Record)
		if !ok {
			log.Panicf("RemoveRecordFields: GIR type %q is not a record", girType)
			return
		}

		fieldMap := make(map[string]struct{}, len(fields))
		for _, field := range fields {
			fieldMap[field] = struct{}{}
		}

		filtered := record.Fields[:0]
		for _, field := range record.Fields {
			if _, ok := fieldMap[field.Name]; !ok {
				filtered = append(filtered, field)
			}
		}
		record.Fields = filtered
	})
}

type modifyCallable struct {
	girType string
	modFunc func(*gir.CallableAttrs)
}

// ModifyCallable is a preprocessor that modifies an existing callable. It only
// does Function or Callback.
func ModifyCallable(girType string, f func(c *gir.CallableAttrs)) Preprocessor {
	girTypeMustBeVersioned(girType)
	return modifyCallable{
		girType: girType,
		modFunc: f,
	}
}

// RenameCallable renames a callable using ModifyCallable.
func RenameCallable(girType, newName string) Preprocessor {
	return ModifyCallable(girType, func(c *gir.CallableAttrs) {
		c.Name = newName
	})
}

// ModifyParamDirections wraps ModifyCallable to conveniently override the
// parameters' directions.
func ModifyParamDirections(girType string, dirOverrides map[string]string) Preprocessor {
	return ModifyCallable(girType, func(c *gir.CallableAttrs) {
		for name, dir := range dirOverrides {
			param := FindParameter(c, name)
			if param == nil {
				log.Panicf("cannot find parameter %s for %s", name, girType)
			}
			param.Direction = dir
		}
	})
}

func (m modifyCallable) Preprocess(repos gir.Repositories) {
	threeParts := strings.SplitN(m.girType, ".", 3)
	girType := strings.Join(threeParts[:2], ".")

	result := repos.FindFullType(girType)
	if result == nil {
		log.Printf("GIR type %q not found", m.girType)
		return
	}

	switch v := result.Type.(type) {
	case *gir.Function:
		m.modFunc(&v.CallableAttrs)
		return
	case *gir.Callback:
		m.modFunc(&v.CallableAttrs)
		return
	}

	if len(threeParts) != 3 {
		goto notFound
	}

	switch v := result.Type.(type) {
	case *gir.Class:
		for i, ctor := range v.Constructors {
			if ctor.Name == threeParts[2] {
				m.modFunc(&v.Constructors[i].CallableAttrs)
				return
			}
		}
		for i, method := range v.Methods {
			if method.Name == threeParts[2] {
				m.modFunc(&v.Methods[i].CallableAttrs)
				return
			}
		}
	case *gir.Record:
		for i, ctor := range v.Constructors {
			if ctor.Name == threeParts[2] {
				m.modFunc(&v.Constructors[i].CallableAttrs)
				return
			}
		}
		for i, method := range v.Methods {
			if method.Name == threeParts[2] {
				m.modFunc(&v.Methods[i].CallableAttrs)
				return
			}
		}
	case *gir.Interface:
		for i, method := range v.Methods {
			if method.Name == threeParts[2] {
				m.modFunc(&v.Methods[i].CallableAttrs)
				return
			}
		}
		for i, method := range v.VirtualMethods {
			if method.Name == threeParts[2] {
				m.modFunc(&v.VirtualMethods[i].CallableAttrs)
				return
			}
		}
	}

notFound:
	log.Panicf("GIR type %q has no callable", m.girType)
}

// FilterMatcher describes a filter for a GIR type.
type FilterMatcher interface {
	// Filter matches for the girType within the given namespace from the
	// namespace generator. The GIR type will never have a namespace prefix.
	Filter(gen FileGenerator, gir, c string) (omit bool)

	// TODO: use this API.
	// Filter(gen FileGenerator, res *gir.TypeFindResult) (omit bool)
}

// Filter returns true if the given GIR and/or C type should be omitted from the
// given generator.
func Filter(gen FileGenerator, gir, c string) (omit bool) {
	gir = EnsureNamespace(gen.Namespace(), gir)

	for _, filter := range gen.Filters() {
		if filter.Filter(gen, gir, c) {
			return true
		}
	}

	return false
}

// FilterCType filters only the C type or identifier. It is useful for filtering
// C functions and such.
func FilterCType(gen FileGenerator, c string) (omit bool) {
	return Filter(gen, "\x00", c)
}

// FilterSub filters a field or method inside a parent.
func FilterSub(gen FileGenerator, parent, sub, cType string) (omit bool) {
	if cType == "" {
		// If the method is missing a C identifier for some dumb reason, we
		// should ensure that it will never be matched incorrectly.
		cType = "\x00"
	}
	girName := strcases.Dots(gen.Namespace().Namespace.Name, parent, sub)
	return Filter(gen, girName, cType)
}

// FilterMethod filters a method similarly to Filter.
func FilterMethod(gen FileGenerator, parent string, method *gir.Method) (omit bool) {
	return FilterSub(gen, parent, method.Name, method.CIdentifier)
}

// FilterField filters a field similarly to Filter.
func FilterField(gen FileGenerator, parent string, field *gir.Field) (omit bool) {
	return FilterSub(gen, parent, field.Name, "")
}

type absoluteFilter struct {
	namespace string
	matcher   string
}

// AbsoluteFilter matches the names absolutely.
func AbsoluteFilter(abs string) FilterMatcher {
	parts := strings.SplitN(abs, ".", 2)
	if len(parts) != 2 {
		log.Panicf("missing namespace for AbsoluteFilter %q", abs)
	}

	return absoluteFilter{parts[0], parts[1]}
}

func (abs absoluteFilter) Filter(gen FileGenerator, gir, c string) (omit bool) {
	if abs.namespace == "C" {
		return c == abs.matcher
	}

	typ, eq := EqNamespace(abs.namespace, gir)
	return eq && typ == abs.matcher
}

type regexFilter struct {
	namespace string
	matcher   *regexp.Regexp
}

// RegexFilter returns a regex filter for FilterMatcher. A regex filter's format
// is a string consisting of two parts joined by a period: a namespace and a
// matcher. The only regex part is the matcher.
func RegexFilter(matcher string) FilterMatcher {
	parts := strings.SplitN(matcher, ".", 2)
	if len(parts) != 2 {
		log.Panicf("invalid regex filter format %q", matcher)
	}

	regex := regexp.MustCompile(wholeMatchRegex(parts[1]))

	return &regexFilter{
		namespace: parts[0],
		matcher:   regex,
	}
}

func wholeMatchRegex(regex string) string {
	// special regex
	if strings.Contains(regex, "(?") {
		return regex
	}

	// must whole match
	return "^" + regex + "$"
}

// Filter implements FilterMatcher.
func (rf *regexFilter) Filter(gen FileGenerator, gir, c string) (omit bool) {
	switch rf.namespace {
	case "C":
		return rf.matcher.MatchString(c)
	case "*":
		return rf.matcher.MatchString(gir)
	}

	typ, eq := EqNamespace(rf.namespace, gir)
	return eq && rf.matcher.MatchString(typ)
}

// EqNamespace is used for FilterMatchers to compare types and namespaces.
func EqNamespace(nsp, girType string) (typ string, ok bool) {
	namespace, typ := gir.SplitGIRType(girType)

	n, ver := gir.ParseVersionName(nsp)
	if ver != "" {
		// Wanted namespace has a version. If the type does not have a version,
		// then pop the version off the wanted namespace.
		_, version := gir.ParseVersionName(namespace)
		if version == "" {
			nsp = n
		}
	}

	return typ, namespace == nsp
}

type fileFilter struct {
	match string
}

// FileFilter filters based on the source position.
func FileFilter(contains string) FilterMatcher {
	return fileFilter{contains}
}

func (ff fileFilter) Filter(gen FileGenerator, girT, cT string) (omit bool) {
	res := Find(gen, girT)
	if res == nil {
		return false
	}

	return TypeIsInFile(res.Type, ff.match)
}

// TypeIsInFile returns true if the given type was declared in the given
// filename. The filename shouldn't contain the file extension.
func TypeIsInFile(typ interface{}, file string) bool {
	info := cmt.GetInfoFields(typ)
	if info.Elements == nil {
		log.Panicf("type %T missing info.Elements", typ)
	}

	if info.Elements.SourcePosition != nil {
		if strings.Contains(info.Elements.SourcePosition.Filename, file) {
			return true
		}
	}

	if info.Elements.Doc != nil {
		if strings.Contains(info.Elements.Doc.Filename, file) {
			return true
		}
	}

	return false
}

package types

import (
	"log"
	"regexp"
	"strings"

	"github.com/diamondburned/gotk4/gir"
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

type typeRenamer struct {
	from, to string
}

// TypeRenamer creates a new filter matcher that renames a type. The given GIR
// type must contain the versioned namespace, like "Gtk3.Widget" but the given
// name must not. The GIR type is absolutely matched, similarly to
// AbsoluteFilter.
func TypeRenamer(girType, newName string) Preprocessor {
	_, version := gir.ParseVersionName(girType)
	if version == "" {
		log.Panicf("girType %q missing version", girType)
	}

	return typeRenamer{
		from: girType,
		to:   newName,
	}
}

func (ren typeRenamer) Preprocess(repos gir.Repositories) {
	result := repos.FindFullType(ren.from)
	if result == nil {
		log.Printf("GIR type %q not found", ren.from)
	}

	result.SetName(ren.to)
}

// FilterMatcher describes a filter for a GIR type.
type FilterMatcher interface {
	// Filter matches for the girType within the given namespace from the
	// namespace generator. The GIR type will never have a namespace prefix.
	Filter(gen FileGenerator, gir, c string) (omit bool)
}

// Filter returns true if the given GIR and/or C type should be omitted from the
// given generator.
func Filter(gen FileGenerator, gir, c string) (omit bool) {
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

// FilterMethod filters a method similarly to Filter.
func FilterMethod(gen FileGenerator, parent string, method *gir.Method) (omit bool) {
	girName := strcases.Dots(gen.Namespace().Namespace.Name, parent, method.Name)

	cType := method.CIdentifier
	if cType == "" {
		// If the method is missing a C identifier for some dumb reason, we
		// should ensure that it will never be matched incorrectly.
		cType = "\x00"
	}

	return Filter(gen, girName, cType)
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
	return typ, namespace == nsp
}

type fileFilter struct {
	match *regexp.Regexp
}

// FileFilter filters based on the source position.
func FileFilter(regex string) FilterMatcher {
	return fileFilter{regexp.MustCompile(regex)}
}

func (ff fileFilter) Filter(gen FileGenerator, girT, cT string) (omit bool) {
	res := Find(gen, girT)
	if res == nil {
		return false
	}

	var source *gir.SourcePosition

	switch v := res.Type.(type) {
	case *gir.Alias:
		source = v.SourcePosition
	case *gir.Class:
		source = v.SourcePosition
	case *gir.Interface:
		source = v.SourcePosition
	case *gir.Record:
		source = v.SourcePosition
	case *gir.Enum:
		source = v.SourcePosition
	case *gir.Function:
		source = v.SourcePosition
	case *gir.Union:
		source = v.SourcePosition
	case *gir.Bitfield:
		source = v.SourcePosition
	case *gir.Callback:
		source = v.SourcePosition
	}

	if source == nil {
		return false
	}

	return !ff.match.MatchString(source.Filename)
}

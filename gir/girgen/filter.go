package girgen

import (
	"log"
	"regexp"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

// FilterMatcher describes a filter for a GIR type.
type FilterMatcher interface {
	// Filter matches for the girType within the given namespace from the
	// namespace generator. The GIR type will never have a namespace prefix.
	Filter(ng *NamespaceGenerator, girType, cType string) (keep bool)
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

func (abs absoluteFilter) Filter(ng *NamespaceGenerator, girType, cType string) (keep bool) {
	if abs.namespace == "C" {
		return cType != abs.matcher
	}

	typ, eq := EqNamespace(abs.namespace, ng.Namespace(), girType)
	return !eq || typ != abs.matcher
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
func (rf *regexFilter) Filter(ng *NamespaceGenerator, girType, cType string) (keep bool) {
	if rf.namespace == "C" {
		return !rf.matcher.MatchString(cType)
	}

	typ, eq := EqNamespace(rf.namespace, ng.Namespace(), girType)
	if !eq {
		return true
	}

	return !rf.matcher.MatchString(typ)
}

// EqNamespace is used for FilterMatchers to compare types and namespaces.
func EqNamespace(nsp string, current *gir.Namespace, girType string) (typ string, ok bool) {
	fullName, typ := gir.SplitGIRType(girType)
	namespace, _ := gir.ParseVersionName(fullName)

	// Type has namespace; ensure the namespace matches the matcher's.
	if namespace != "" && namespace != nsp {
		return "", false
	}

	// Type has no namespace; check the current generator's namespace instead.
	if namespace == "" && current.Name != nsp {
		return "", false
	}

	return typ, true
}

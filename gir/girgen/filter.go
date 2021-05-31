package girgen

import (
	"log"
	"regexp"
	"strings"
	"sync"
)

// FilterMatcher describes a filter for a GIR type.
type FilterMatcher interface {
	// Filter matches for the girType within the given namespace from the
	// namespace generator. The GIR type will never have a namespace prefix.
	Filter(ng *NamespaceGenerator, girType string) (keep bool)
}

// RegexFilter is a regex filter for FilterMatcher. A regex filter's format is a
// string consisting of two parts joined by a period: a namespace and a matcher.
// The only regex part is the matcher.
type RegexFilter string

var _ FilterMatcher = RegexFilter("")

var (
	regexCache map[string]*regexp.Regexp
	regexMutex sync.Mutex
)

// parts parses the RegexFilter into 2 parts and guarantees its format.
func (rf RegexFilter) parts() []string {
	parts := strings.SplitN(string(rf), ".", 2)
	if len(parts) != 2 {
		log.Panicf("invalid regex filter format %q", string(rf))
	}
	return parts
}

func (rf RegexFilter) namespace() string {
	return rf.parts()[0]
}

func (rf RegexFilter) regex() (regex *regexp.Regexp) {
	regexStr := rf.parts()[1]
	var ok bool

	regexMutex.Lock()
	defer regexMutex.Unlock()

	regex, ok = regexCache[regexStr]
	if !ok {
		regex = regexp.MustCompile(regexStr)
		regexCache[regexStr] = regex
	}

	return
}

// Filter implements FilterMatcher.
func (rf RegexFilter) Filter(ng *NamespaceGenerator, girType string) (keep bool) {
	if ng.Namespace().Name != rf.namespace() {
		return true
	}

	return !rf.regex().MatchString(girType)
}

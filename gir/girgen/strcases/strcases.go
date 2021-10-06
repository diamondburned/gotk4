// Package strcases provides helper functions to convert between string cases,
// such as Pascal Case, snake_case and Go's Mixed Caps, along with various
// special cases.
package strcases

import (
	"log"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	_ "embed"
)

//go:embed capitalized.txt
var capitalizedTXT string

//go:embed replaced.txt
var replacedTXT string

var (
	snakeRegex     = regexp.MustCompile(`[_0-9]+\w`)
	pascalSpecials = strings.Split(capitalizedTXT, "\n")
	pascalWords    = map[string]string{}

	pascalRegex        *regexp.Regexp
	pascalPostReplacer *strings.Replacer
)

func initPascalWords() {
	for _, line := range strings.Split(replacedTXT, "\n") {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		words := strings.Split(line, "->")
		if len(words) != 2 {
			log.Fatalf("invalid replace %q", line)
		}

		words[0] = strings.TrimSpace(words[0])
		words[1] = strings.TrimSpace(words[1])
		pascalWords[words[0]] = words[1]
	}
}

func initPascalRegex() {
	fullRegex := strings.Builder{}
	fullRegex.Grow(256)
	fullRegex.WriteByte('(')

	for i, special := range pascalSpecials {
		if special == "" {
			continue
		}
		if i > 0 {
			fullRegex.WriteByte('|')
		}
		fullRegex.WriteString(special)
	}

	fullRegex.WriteByte(')')

	// Must account for the next character being either EOF or a capitalized
	// letter to avoid cases like "IDentifier".
	fullRegex.WriteString("([A-Z0-9]|$)")

	pascalRegex = regexp.MustCompile(fullRegex.String())
}

// AddPascalSpecials adds the given list of regexes into the list of cases that
// will be fully capitalized during case conversion to Go.
func AddPascalSpecials(regexes []string) {
	pascalSpecials = append(pascalSpecials, regexes...)
	initPascalRegex()
}

func initPascalPostReplacer() {
	postReplacerArgs := make([]string, len(pascalWords)*2)
	for from, to := range pascalWords {
		postReplacerArgs = append(postReplacerArgs, from, to)
	}

	pascalPostReplacer = strings.NewReplacer(postReplacerArgs...)
}

// SetPascalWords sets the given map of words to be replaced after the pascal
// specials stage as a method of fixing edge cases.
func SetPascalWords(words map[string]string) {
	for from, to := range words {
		pascalWords[from] = to
	}
	initPascalPostReplacer()
}

func init() {
	initPascalWords()
	initPascalRegex()
	initPascalPostReplacer()
}

// Dots is a helper function to join strings in dots for debugging.
func Dots(parts ...string) string {
	nonEmptyParts := parts[:0]

	for _, part := range parts {
		if part == "" {
			continue
		}

		if strings.Contains(part, "*") {
			part = "(" + part + ")"
		}

		nonEmptyParts = append(nonEmptyParts, part)
	}

	return strings.Join(nonEmptyParts, ".")
}

// IsLower returns true if the string is all lower-cased.
func IsLower(s string) bool {
	return strings.IndexFunc(s, unicode.IsUpper) == -1
}

// GuessSnake guesses if the given name is snake-cased or not.
func GuessSnake(name string) (snake bool) {
	return strings.Contains(name, "_") || IsLower(name)
}

// Go converts either pascal or snake case to the Go name. The original casing
// is inferred from the given name.
func Go(name string) string {
	if GuessSnake(name) {
		return SnakeToGo(true, name)
	} else {
		return PascalToGo(name)
	}
}

// PascalToGo converts regular Pascal case to Go.
func PascalToGo(pascal string) string {
	// Force constructors to have a New prefix instead of suffix.
	if strings.HasSuffix(pascal, "New") {
		pascal = "New" + strings.TrimSuffix(pascal, "New")
	}

	pascal = pascalRegex.ReplaceAllStringFunc(pascal, strings.ToUpper)
	pascal = pascalPostReplacer.Replace(pascal)

	return pascal
}

// FirstLetter returns the first letter in lower-case.
func FirstLetter(p string) string {
	r, sz := utf8.DecodeRuneInString(p)
	if sz > 0 && r != utf8.RuneError {
		return string(unicode.ToLower(r))
	}

	return string(p[0]) // fallback
}

// UnexportPascal converts the PascalToGo string to be unexported.
func UnexportPascal(pascal string) string {
	runes := []rune(pascal)
	if len(runes) < 1 {
		return SnakeNoGo(strings.ToLower(pascal))
	}

	var i int
	for i < len(runes) && unicode.IsUpper(runes[i]) {
		i++
	}

	if i > 1 {
		i--
	}

	pascal = strings.ToLower(string(runes[:i])) + string(runes[i:])
	pascal = SnakeNoGo(pascal)

	return pascal
}

// SnakeToGo converts snake case to Go's special case. If Pascal is true, then
// the first letter is capitalized.
func SnakeToGo(pascal bool, snakeString string) string {
	if pascal {
		snakeString = "_" + snakeString
	}

	snakeString = snakeRegex.ReplaceAllStringFunc(snakeString,
		func(orig string) string {
			orig = strings.ToUpper(orig)
			orig = strings.Replace(orig, "_", "", 2)
			return orig
		},
	)

	if !pascal {
		return SnakeNoGo(snakeString)
	}

	return PascalToGo(snakeString)
}

// KebabToGo converts kebab case to Go's special case. See SnakeToGo.
func KebabToGo(pascal bool, kebabString string) string {
	return SnakeToGo(pascal, strings.ReplaceAll(kebabString, "-", "_"))
}

// GoKeywords includes Go keywords. This is primarily to prevent collisions with
// meaningful Go words.
var GoKeywords = map[string]string{
	// Keywords.
	"break":       "",
	"default":     "",
	"func":        "fn",
	"interface":   "iface",
	"select":      "sel",
	"case":        "",
	"defer":       "",
	"go":          "",
	"map":         "",
	"struct":      "",
	"chan":        "ch",
	"else":        "",
	"goto":        "",
	"package":     "pkg",
	"switch":      "",
	"const":       "",
	"fallthrough": "",
	"if":          "",
	"range":       "",
	"type":        "typ",
	"continue":    "",
	"for":         "",
	"import":      "",
	"return":      "ret",
	"var":         "",
}

// GoBuiltinTypes contains Go built-in types.
var GoBuiltinTypes = map[string]string{
	// Types.
	"bool":       "",
	"byte":       "",
	"complex128": "cmplx",
	"complex64":  "cmplx",
	"error":      "err",
	"float32":    "",
	"float64":    "",
	"int":        "",
	"int16":      "",
	"int32":      "",
	"int64":      "",
	"int8":       "",
	"rune":       "",
	"string":     "str",
	"uint":       "",
	"uint16":     "",
	"uint32":     "",
	"uint64":     "",
	"uint8":      "",
	"uintptr":    "",
}

// CGoField formats the C field name to not be confused with a Go keyword.
// See https://golang.org/cmd/cgo/#hdr-Go_references_to_C.
func CGoField(field string) string {
	_, keyword := GoKeywords[field]
	if keyword {
		return "_" + field
	}
	return field
}

// SnakeNoGo ensures the snake-case string is never a Go keyword.
func SnakeNoGo(snake string) string {
	s, isKeyword := GoKeywords[snake]
	if isKeyword {
		if s != "" {
			return s
		}
		return "_" + snake
	}

	s, isType := GoBuiltinTypes[snake]
	if isType {
		if s != "" {
			return s
		}
		return "_" + snake
	}

	return snake
}

var vowels = [255]bool{
	'a': true,
	'i': true,
	'u': true,
	'e': true,
	'o': true,
}

// Interfacify appends the -er suffix into the given word to idiomatically
// adhere to Go's interface naming convention. If the word already ends with an
// -er suffix, then another suffix will be added.
func Interfacify(word string) string {
	// https://www.englishclub.com/spelling/rules-add-er-est.htm
	// https://www.thefreedictionary.com/Commonly-Confused-Suffixes-er-or-ar.htm
	// https://ginsengenglish.com/blog/cvc-words
	switch {
	case wordConsonantAndSuffix(word, 'e'):
		fallthrough
	case wordConsonantAndSuffix(word, 'a'):
		fallthrough
	case wordConsonantAndSuffix(word, 'o'):
		return word + "r"

	case wordConsonantAndSuffix(word, 'y'):
		return word[:len(word)-1] + "ier"

	case strings.HasSuffix(word, "it") && !wordIsOrException(word):
		fallthrough
	case strings.HasSuffix(word, "ct"):
		return word + "or"

	// CVC form is bad. It's ugly.
	case wordIsCVC(word) && !strings.HasSuffix(word, "er"):
		return word + string(word[len(word)-1]) + "er"

	case wordEndsInConsonant(word):
		fallthrough
	default:
		return word + "er"
	}
}

var orExceptions = []string{"delimit", "profit", "recruit"}

func wordIsOrException(word string) bool {
	for _, exc := range orExceptions {
		if strings.EqualFold(exc, word) {
			return true
		}
	}
	return false
}

// wordConsonantAndSuffix returns true if the word ends with a consonant and the
// given suffix character.
func wordConsonantAndSuffix(word string, char byte) bool {
	if len(word) < 2 {
		return false
	}

	last2 := word[len(word)-2:]
	return !vowels[last2[0]] && last2[1] == char
}

// wordEndsInConsonant returns true if the word ends with a consonant.
func wordEndsInConsonant(word string) bool {
	return len(word) > 1 && !vowels[word[len(word)-1]]
}

var cvcExceptions = [255]bool{
	'w': true,
	'x': true,
	'y': true,
}

// wordIsCVC returns true if the given word follows the C+V+C form.
func wordIsCVC(word string) bool {
	if len(word) < 3 {
		return false
	}

	last3 := word[len(word)-3:]

	return !cvcExceptions[last3[2]] &&
		!vowels[last3[0]] && vowels[last3[1]] && !vowels[last3[2]]
}

package girgen

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	snakeRegex     = regexp.MustCompile(`[_0-9]+\w`)
	pascalSpecials = []string{
		"Id",
		"Io",
		"Uri",
		"Url",
		"Css",
		"Dpi",
		"Ltr",
		"Rtl",
		"Gtk",
		"Nul",
		"Eof",
		"Md5",
		"Dmy",
		"Nfc",
		"Nfd",
		"Nfkc",
		"Nfkd",
		"Simd",
		"Hmac",
		"Ascii",
		"Toupper",
		"Tolower",
		`Sha(\d+)?`,
		`Utf(\d+)?`,
		`[XY][^aiueo]`,
		`(|S|R)[XYZxyz]{3}`,
	}
	pascalRegex *regexp.Regexp
)

func init() {
	fullRegex := strings.Builder{}
	fullRegex.Grow(256)
	fullRegex.WriteByte('(')

	for i, special := range pascalSpecials {
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

// See Go specs, section Keywords.
var goKeywords = map[string]struct{}{
	"break":       {},
	"default":     {},
	"func":        {},
	"interface":   {},
	"select":      {},
	"case":        {},
	"defer":       {},
	"go":          {},
	"map":         {},
	"struct":      {},
	"chan":        {},
	"else":        {},
	"goto":        {},
	"package":     {},
	"switch":      {},
	"const":       {},
	"fallthrough": {},
	"if":          {},
	"range":       {},
	"type":        {},
	"continue":    {},
	"for":         {},
	"import":      {},
	"return":      {},
	"var":         {},
}

// cgoField formats the C field name to not be confused with a Go keyword.
// See https://golang.org/cmd/cgo/#hdr-Go_references_to_C.
func cgoField(field string) string {
	_, keyword := goKeywords[field]
	if keyword {
		return "_" + field
	}
	return field
}

// PascalToGo converts regular Pascal case to Go.
func PascalToGo(pascal string) string {
	// Force constructors to have a New prefix instead of suffix.
	if strings.HasSuffix(pascal, "New") {
		pascal = "New" + strings.TrimSuffix(pascal, "New")
	}

	return pascalRegex.ReplaceAllStringFunc(pascal, strings.ToUpper)
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
			orig = strings.ReplaceAll(orig, "_", "")
			return orig
		},
	)

	snakeString = PascalToGo(snakeString)

	if !pascal {
		// Special cases.
		_, isKeyword := goKeywords[snakeString]
		if isKeyword {
			snakeString = "_" + snakeString
		}
	}

	return snakeString
}

// FirstChar returns the first character.
func FirstChar(str string) string {
	r, _ := utf8.DecodeRune([]byte(str))
	return string(unicode.ToLower(r))
}

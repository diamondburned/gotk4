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
		"Ltr",
		"Rtl",
		"Gtk",
		"Nul",
		"Eof",
		"Md5",
		"Nfc",
		"Nfd",
		"Nfkc",
		"Nfkd",
		`Sha(\d+)?`,
		`Utf(\d+)?`,
		`[XY][^aiueo]`,
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
	fullRegex.WriteString("([A-Z]|$)")

	pascalRegex = regexp.MustCompile(fullRegex.String())
}

// PascalToGo converts regular Pascal case to Go.
func PascalToGo(pascal string) string {
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

	return PascalToGo(snakeString)
}

// FirstChar returns the first character.
func FirstChar(str string) string {
	r, _ := utf8.DecodeRune([]byte(str))
	return string(unicode.ToLower(r))
}

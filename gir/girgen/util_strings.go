package girgen

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	snakeRegex    = regexp.MustCompile(`[_0-9]+\w`)
	snakeXYRegex  = regexp.MustCompile(`[XY][^aiueo]`)
	snakeSpecials = strings.NewReplacer(
		"Id", "ID",
		"Gtk", "GTK",
	)
)

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

	snakeString = snakeSpecials.Replace(snakeString)

	// Capitalize Xalign, Ytilt, etc.
	snakeString = snakeXYRegex.ReplaceAllStringFunc(snakeString,
		func(orig string) string {
			return strings.ToUpper(orig)
		},
	)

	return snakeString
}

// FirstChar returns the first character.
func FirstChar(str string) string {
	r, _ := utf8.DecodeRune([]byte(str))
	return string(unicode.ToLower(r))
}

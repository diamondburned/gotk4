package girgen

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// string utilities

var (
	snakeRegex = regexp.MustCompile(`_\w`)
	snakeRepl  = strings.NewReplacer(
		"Xalign", "XAlign",
		"Yalign", "YAlign",
		"Id", "ID",
	)
)

func snakeToGo(pascal bool, snakeString string) string {
	if pascal {
		snakeString = "_" + snakeString
	}

	snakeString = snakeRegex.ReplaceAllStringFunc(snakeString,
		func(orig string) string {
			return string(unicode.ToUpper(rune(orig[1])))
		},
	)

	return snakeRepl.Replace(snakeString)
}

func firstChar(str string) string {
	r, _ := utf8.DecodeRune([]byte(str))
	return string(unicode.ToLower(r))
}

// Package cmt provides functions that parse and render GIR comments into nice
// and conventional Go comments.
package cmt

import (
	"go/doc"
	"html"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/diamondburned/gotk4/gir"
)

const (
	CommentsColumnLimit = 80 - 3 // account for prefix "// "
	CommentsTabWidth    = 4
)

var (
	cmtNamespaceRegex = regexp.MustCompile(`#[A-Z]\w+?[A-Z]`)
	cmtArgumentRegex  = regexp.MustCompile(`@\w+`)
	cmtPrimitiveRegex = regexp.MustCompile(`%\w+`)
	cmtFunctionRegex  = regexp.MustCompile(`\w+\(\)`)
	cmtHeadingRegex   = regexp.MustCompile(`\n*#+ (.*?)(?: ?#+ ?\{#.*?\})?\n+`)
	cmtCodeblockRegex = regexp.MustCompile(`(?ms)\n*\|\[(?:<!--.*-->)?\n(.*?)\n\]\|\n*`)
	cmtHyperlinkRegex = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
)

// GoDoc renders a Go documentation string from the given GIR doc. If doc is
// nil, then an empty string is returned.
func GoDoc(doc *gir.Doc, indentLvl int, self string) string {
	if doc == nil {
		return ""
	}
	return CommentReflowLinesIndent(indentLvl, self, doc.String)
}

// nthWord returns the nth word, or an empty string if none.
func nthWord(paragraph string, n int) string {
	words := strings.SplitN(paragraph, " ", n+2)
	if len(words) < n+2 {
		return ""
	}
	return words[n]
}

// nthWordSimplePresent checks if the second word has a trailing "s".
func nthWordSimplePresent(paragraph string, n int) bool {
	word := nthWord(paragraph, n)
	return !strings.EqualFold(word, "this") && strings.HasSuffix(word, "s")
}

// lowerFirstLetter lower-cases the first letter in the paragraph.
func lowerFirstWord(paragraph string) string {
	r, sz := utf8.DecodeRuneInString(paragraph)
	if sz > 0 {
		return string(unicode.ToLower(r)) + paragraph[sz:]
	}
	return string(paragraph)
}

// CommentReflowLinesIndent reflows the given cmt paragraphs into idiomatic Go
// comment strings. It is automatically indented.
func CommentReflowLinesIndent(indentLvl int, self, cmt string) string {
	cmt = html.UnescapeString(cmt)

	switch {
	case strings.HasPrefix(cmt, "#") && nthWordSimplePresent(cmt, 1):
		// Trim the first word away and replace it with the Go name.
		cmt = self + " " + strings.SplitN(cmt, " ", 2)[1]

	case nthWordSimplePresent(cmt, 0):
		cmt = self + " " + lowerFirstWord(cmt)

	default:
		// Trim the word "this" away to make the sentence gramatically correct.
		cmt = strings.TrimPrefix(cmt, "this ")
		cmt = self + ": " + lowerFirstLetter(cmt)
	}

	// Fix up the codeblocks and render it using GoDoc format.
	cmt = cmtCodeblockRegex.ReplaceAllStringFunc(cmt, func(match string) string {
		matches := cmtCodeblockRegex.FindStringSubmatch(match)

		lines := strings.Split(matches[1], "\n")
		for i, line := range lines {
			lines[i] = "   " + line
		}

		// Use our own new lines.
		return "\n\n" + strings.Join(lines, "\n") + "\n\n"
	})

	// Fix up headers in the preprocessing stage. We also sanitize the trailing
	// new lines here.
	cmt = cmtHeadingRegex.ReplaceAllString(cmt, "\n\n$1\n\n")

	// Wipe the namespace prefix syntax.
	cmt = cmtNamespaceRegex.ReplaceAllStringFunc(cmt, func(str string) string {
		// Replace "#?" with just "?".
		return str[len(str)-1:]
	})

	// Undo all hyperlinks.
	cmt = cmtHyperlinkRegex.ReplaceAllString(cmt, "$1 ($2)")

	// Fix up new lines before we throw this into ToText so to not confuse it.
	cmt = tidyParagraphs(cmt)

	// TODO: Replace snake-cased functions with known ones in the namespace.
	// Prepend a C prefix otherwise.

	// cmt = cmtFunctionsRegex.ReplaceAllStringFunc(cmt, func(str string) string {
	// 	fnName := strings.TrimSuffix(str, "()")
	// 	result := ns.gen.Repos.FindCType(fnName)
	//
	// 	if result.Method != nil {
	// 		if fn.Parameters.HasInstanceParameter() {
	// 			return fmt.Sprintf(
	// 				"(%s).%s()",
	// 				fn.Parameters.InstanceParameter.Type.GoType(),
	// 				fn.GoName(),
	// 			)
	// 		}
	// 	} else {
	// 		return fmt.Sprintf("%s()", fn.GoName())
	// 	}
	//
	// 	return fmt.Sprintf("C.%s()", fnName)
	// })

	// Replace C primitives with Go's.
	cmt = cmtPrimitiveRegex.ReplaceAllStringFunc(cmt, func(str string) string {
		// [:1] trims the % away.
		switch str = str[1:]; str {
		case "NULL":
			return "nil"
		case "TRUE":
			return "true"
		case "FALSE":
			return "false"
		default:
			return str
		}
	})

	// Account for the indentation in the column limit.
	col := CommentsColumnLimit - (CommentsTabWidth * indentLvl)

	cmt = docText(cmt, col)

	ident := strings.Repeat("\t", indentLvl)
	lines := strings.Split(cmt, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		// Trim the trailing empty line, if any.
		lines = lines[:len(lines)-1]
	}

	for i, line := range lines {
		lines[i] = ident + "// " + line
	}

	return strings.Join(lines, "\n")
}

// tidyParagraphs cleans up new lines without touching codeblocks.
func tidyParagraphs(text string) string {
	paragraphs := strings.Split(text, "\n\n")

	for i, paragraph := range paragraphs {
		if strings.HasPrefix(paragraph, " ") {
			continue
		}

		paragraphs[i] = strings.ReplaceAll(paragraph, "\n", " ")
	}

	return strings.Join(paragraphs, "\n\n")
}

func docText(p string, col int) string {
	builder := strings.Builder{}
	builder.Grow(len(p) + 64)

	doc.ToText(&builder, p, "", "   ", col)
	return builder.String()
}

func openOrCloseCodeblock(paragraph string) bool {
	return strings.HasPrefix(paragraph, "|[") || strings.HasSuffix(paragraph, "]|")
}

func lowerFirstLetter(p string) string {
	if p == "" {
		return ""
	}

	runes := []rune(p)
	if len(runes) < 2 {
		return string(unicode.ToLower(runes[0]))
	}

	// Edge case: gTK, etc.
	if unicode.IsUpper(runes[1]) {
		return p
	}

	return string(unicode.ToLower(runes[0])) + string(runes[1:])
}

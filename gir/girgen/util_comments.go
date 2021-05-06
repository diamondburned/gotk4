package girgen

import (
	"fmt"
	"go/doc"
	"regexp"
	"strings"
	"unicode"

	"github.com/diamondburned/gotk4/gir"
)

const (
	CommentsColumnLimit = 80 - 3 // account for prefix "// "
	CommentsTabWidth    = 4
)

var (
	cmtNamespaceRegex = regexp.MustCompile(`#[A-Z]\w+?[A-Z]`)
	cmtArgumentsRegex = regexp.MustCompile(`@\w+`)
	cmtPrimitiveRegex = regexp.MustCompile(`%\w+`)
	cmtFunctionsRegex = regexp.MustCompile(`\w+\(\)`)
	cmtMdHeadingRegex = regexp.MustCompile(`#+ `)
	cmtOpenBlockRegex = regexp.MustCompile(`(?ms)\|\[(?:&lt;!--.*--&gt;\n)?(.*)(?:\]\|)?`)
	cmtWhitespaceProc = strings.NewReplacer(
		"\n\n", "\n\n",
		"\n", " ",
	)
)

// GoDoc renders a Go documentation string from the given GIR doc. If doc is
// nil, then an empty string is returned.
func GoDoc(doc *gir.Doc, indentLvl int, self string) string {
	if doc == nil {
		return ""
	}
	return CommentReflowLinesIndent(indentLvl, self, doc.String)
}

// CommentReflowLinesIndent reflows the given cmt paragraphs into idiomatic Go
// comment strings. It is automatically indented.
func CommentReflowLinesIndent(indentLvl int, self, cmt string) string {
	// Trim the word "this" away to make the sentence gramatically correct.
	cmt = strings.TrimPrefix(cmt, "this ")
	cmt = strings.TrimPrefix(cmt, "#")

	// Wipe the namespace.
	cmt = cmtNamespaceRegex.ReplaceAllStringFunc(cmt, func(str string) string {
		// Return only the last character.
		return str[len(str)-1:]
	})

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

	// Split into paragraphs and parse each of them.
	paragraphs := strings.Split(cmt, "\n\n")
	codeblock := false

	for i, paragraph := range paragraphs {
		codeblockChange := openOrCloseCodeblock(paragraph)

		switch {
		// Codeblock:
		case codeblock || codeblockChange:
			// Toggle codeblock state.
			if codeblockChange {
				codeblock = !codeblock
			}

			// Parse the codeblock.
			var lines []string
			if m := cmtOpenBlockRegex.FindStringSubmatch(paragraph); m != nil {
				lines = strings.Split(m[1], "\n")
			} else {
				paragraph = strings.TrimSuffix(paragraph, "]|")
				lines = strings.Split(paragraph, "\n")
			}

			// Edge case w/ GVariant's comment.
			if len(lines) > 0 && lines[len(lines)-1] == "]|" {
				lines = lines[:len(lines)-1]
				codeblock = false
			}

			// Render the codeblock in GoDoc markup.
			for i, line := range lines {
				lines[i] = "  " + line
			}

			paragraphs[i] = strings.Join(lines, "\n")

		// Headings, but account for any number of hashes.
		case strings.HasPrefix(strings.SplitN(paragraph, " ", 2)[0], "#"):
			// Go's heading syntax doesn't require the hash.
			paragraph = cmtMdHeadingRegex.ReplaceAllString(paragraph, "")
			// Ensure there's no period.
			paragraphs[i] = strings.TrimSuffix(paragraph, ".")

		case !codeblock:
			fallthrough

		// Normal paragraphs.
		default:
			paragraph = strings.TrimSpace(cmtWhitespaceProc.Replace(paragraph))
			if i == 0 {
				// Prefix the paragraph with the current entity.
				paragraph = fmt.Sprintf("%s: %s", self, lowerFirstLetter(paragraph))
			}
			paragraphs[i] = paragraph
		}
	}

	// Account for the indentation in the column limit.
	col := CommentsColumnLimit - (CommentsTabWidth * indentLvl)

	cmt = strings.Join(paragraphs, "\n\n")
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

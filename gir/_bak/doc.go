package gir

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/dave/jennifer/jen"
	"github.com/mitchellh/go-wordwrap"
)

type Doc struct {
	XMLName  xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 doc"`
	Filename string   `xml:"filename,attr"`
	Line     int      `xml:"line,attr"`
	String   string   `xml:",innerxml"`
}

// Lower returns the comment with the first character lower-cased.
func (d Doc) Lower() string {
	r, sz := utf8.DecodeRuneInString(d.String)
	if sz > 0 {
		return string(unicode.ToLower(r)) + d.String[sz:]
	}

	return d.String
}

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

// GenGoComments generates comments in idiomatic Go style. The given selfName
// replaces @self with the given receiver.
func (d Doc) GenGoComments(selfName, prefix string) *jen.Statement {
	return d.GenGoCommentsIndent(0, selfName, prefix)
}

func (d Doc) GenGoCommentsIndent(indentLvl uint, selfName, prefix string) *jen.Statement {
	if d.String == "" {
		return nil
	}

	cmt := d.Lower()

	// Replace @self with the given receiver name.
	cmt = cmtArgumentsRegex.ReplaceAllStringFunc(cmt, func(str string) string {
		if str == "@self" && selfName != "" {
			return selfName
		}
		return str[1:]
	})

	return GenCommentReflowLinesIndent(indentLvl, prefix, cmt)
}

func (d Doc) GenComments() *jen.Statement {
	return jen.Comment(d.String).Line()
}

const (
	CommentsColumnLimit = 80 - 3 // account for prefix "// "
	CommentsTabWidth    = 4
)

// CommentReflowLinesIndent reflows the given cmt paragraphs into idiomatic Go
// comment strings. It is automatically indented.
func CommentReflowLinesIndent(indentLvl uint, self, cmt string) string {
	// Account for the indentation in the column limit.
	columns := CommentsColumnLimit - (CommentsTabWidth * indentLvl)

	// Trim the word "this" away to make the sentence gramatically correct.
	cmt = strings.TrimPrefix(cmt, "this ")

	// Wipe the namespace.
	cmt = cmtNamespaceRegex.ReplaceAllStringFunc(cmt, func(str string) string {
		// Return only the last character.
		return str[len(str)-1:]
	})

	// Replace snake-cased functions with known ones in the namespace. Prepend a
	// C prefix otherwise.
	cmt = cmtFunctionsRegex.ReplaceAllStringFunc(cmt, func(str string) string {
		var fnName = strings.TrimSuffix(str, "()")

		switch fn := activeNamespace.FnWithC(fnName); fn := fn.(type) {
		case Method:
			if fn.Parameters.HasInstanceParameter() {
				return fmt.Sprintf(
					"(%s).%s()",
					fn.Parameters.InstanceParameter.Type.GoType(),
					fn.GoName(),
				)
			}
		case GoNamer:
			return fmt.Sprintf("%s()", fn.GoName())
		}

		return fmt.Sprintf("C.%s()", fnName)
	})

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
	var paragraphs = strings.Split(cmt, "\n\n")
	var codeblock = false

	for i, paragraph := range paragraphs {
		switch codeblockChange := openOrCloseCodeblock(paragraph); {
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

			// Render the codeblock in GoDoc markup.
			for i, line := range lines {
				lines[i] = "   " + line
			}
			paragraphs[i] = strings.Join(lines, "\n")

		// Headings;
		case strings.HasPrefix(paragraph, "#"):
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
				paragraph = fmt.Sprintf("%s %s", self, paragraph)
			}
			paragraph = wordwrap.WrapString(paragraph, columns)
			paragraphs[i] = paragraph
		}
	}

	cmt = strings.Join(paragraphs, "\n\n")

	lines := strings.Split(cmt, "\n")
	for i, line := range lines {
		lines[i] = "// " + line
	}

	return strings.Join(lines, "\n")
}

func openOrCloseCodeblock(paragraph string) bool {
	return strings.HasPrefix(paragraph, "|[") || strings.HasSuffix(paragraph, "]|")
}

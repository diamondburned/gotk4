// Package cmt provides functions that parse and render GIR comments into nice
// and conventional Go comments.
package cmt

import (
	"fmt"
	"go/doc"
	"html"
	"log"
	"reflect"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
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

// InfoFields contains common fields that a GIR schema type may contain. These
// fields should rarely be used, because they're too magical. All fields inside
// InfoFields are optional.
type InfoFields struct {
	Name     *string
	Attrs    *gir.InfoAttrs
	Elements *gir.InfoElements
}

func getField(value reflect.Value, field string) interface{} {
	v := value.FieldByName(field)
	if v == (reflect.Value{}) {
		return nil
	}
	if v.Type().Kind() == reflect.Ptr {
		return v.Interface()
	}
	if v.CanAddr() {
		return v.Addr().Interface()
	}
	cpy := reflect.New(v.Type())
	cpy.Elem().Set(v)
	return cpy.Interface()
}

// GetInfoFields gets the InfoFields from the given value.
func GetInfoFields(v interface{}) InfoFields {
	value := reflect.Indirect(reflect.ValueOf(v))
	if value.Kind() != reflect.Struct {
		panic("given value is not a struct")
	}

	var inf InfoFields

	inf.Name, _ = getField(value, "Name").(*string)
	inf.Attrs, _ = getField(value, "InfoAttrs").(*gir.InfoAttrs)
	inf.Elements, _ = getField(value, "InfoElements").(*gir.InfoElements)

	return inf
}

// EnsureInfoFields ensures that the given type contains all fields inside
// InfoFields.
func EnsureInfoFields(v interface{}) struct{} {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	mustField := func(name string, fieldTyp reflect.Type) {
		field, ok := typ.FieldByName(name)
		if !ok {
			log.Panicf("type %v missing field %s", typ, name)
		}
		if field.Type.Kind() == reflect.Ptr {
			field.Type = field.Type.Elem()
		}
		if field.Type != fieldTyp {
			log.Panicf("type %v field %s got type %v, expected %v", typ, name, field.Type, fieldTyp)
		}
	}

	mustField("Name", reflect.TypeOf(""))
	mustField("InfoAttrs", reflect.TypeOf(gir.InfoAttrs{}))
	mustField("InfoElements", reflect.TypeOf(gir.InfoElements{}))
	return struct{}{}
}

var (
	_ = EnsureInfoFields((*gir.Alias)(nil))
	_ = EnsureInfoFields((*gir.Bitfield)(nil))
	_ = EnsureInfoFields((*gir.Callback)(nil))
	_ = EnsureInfoFields((*gir.Class)(nil))
	_ = EnsureInfoFields((*gir.Enum)(nil))
	_ = EnsureInfoFields((*gir.Function)(nil))
	_ = EnsureInfoFields((*gir.Interface)(nil))
	_ = EnsureInfoFields((*gir.Record)(nil))
	_ = EnsureInfoFields((*gir.Union)(nil))
)

// Option defines possible options for GoDoc.
type Option interface{ opts() }

type (
	overrideSelfName string
	originalTypeName string
	additionalString string
	trailingNewLine  struct{}
	synopsize        struct{}
)

func (overrideSelfName) opts() {}
func (originalTypeName) opts() {}
func (additionalString) opts() {}
func (trailingNewLine) opts()  {}
func (synopsize) opts()        {}

func searchOpts(opts []Option, f func(opt Option) bool) bool {
	for _, opt := range opts {
		if f(opt) {
			return true
		}
	}
	return false
}

// OverrideSelfName overrides the Go type name that's implicitly converted
// automatically by GoDoc.
func OverrideSelfName(self string) Option { return overrideSelfName(self) }

// AdditionalString adds the given string into the tail of the command as
// another paragraph.
func AdditionalString(str string) Option { return additionalString(str) }

// TrailingNewLine adds a trailing new line during documentation generation.
func TrailingNewLine() Option { return trailingNewLine{} }

func isLower(s string) bool {
	return strings.IndexFunc(s, unicode.IsUpper) == -1
}

// Synopsis renders the synopsis of the documentation.
func Synopsis(v interface{}, indentLvl int, opts ...Option) string {
	return goDoc(v, indentLvl, append(opts, synopsize{}))
}

// GoDoc renders a Go documentation string from the given struct. The struct
// should contain at least the field Name, InfoAttrs and InfoElements.
func GoDoc(v interface{}, indentLvl int, opts ...Option) string {
	return goDoc(v, indentLvl, opts)
}

func goDoc(v interface{}, indentLvl int, opts []Option) string {
	inf := GetInfoFields(v)

	var self string
	var orig string

	if inf.Name != nil {
		orig = *inf.Name

		if strings.Contains(orig, "_") || isLower(orig) {
			self = strcases.SnakeToGo(true, orig)
		} else {
			self = strcases.PascalToGo(orig)
		}
	}

	var docBuilder strings.Builder
	if inf.Elements != nil && inf.Elements.Doc != nil {
		docBuilder.WriteString(inf.Elements.Doc.String)
	}

	if inf.Attrs != nil && inf.Attrs.Deprecated {
		if docBuilder.Len() > 0 {
			docBuilder.WriteString("\n\n")
		}

		switch {
		case inf.Elements != nil && inf.Elements.DocDeprecated != nil:
			v := strings.TrimSuffix(inf.Elements.DocDeprecated.String, ".")
			fmt.Fprintf(&docBuilder, "Deprecated: %s.", v)
		case inf.Attrs.DeprecatedVersion != "":
			v := strings.TrimSuffix(inf.Attrs.DeprecatedVersion, ".")
			fmt.Fprintf(&docBuilder, "Deprecated: since version %s.", v)
		default:
			fmt.Fprintf(&docBuilder, "Deprecated.")
		}
	}

	for _, opt := range opts {
		switch opt := opt.(type) {
		case overrideSelfName:
			self = string(opt)

		case additionalString:
			if docBuilder.Len() > 0 {
				docBuilder.WriteString("\n\n")
			}
			docBuilder.WriteString(string(opt))
		}
	}

	opts = append(opts, originalTypeName(orig))
	cmt := format(self, docBuilder.String(), opts)

	synopsize := searchOpts(opts, func(opt Option) bool {
		_, ok := opt.(synopsize)
		return ok
	})
	if synopsize {
		cmt = doc.Synopsis(cmt)
	}

	cmt = ReflowLinesIndent(indentLvl, cmt, opts...)

	return cmt
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

// popFirstWord pops the first word off.
func popFirstWord(paragraph string) string {
	parts := strings.SplitN(paragraph, " ", 2)
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}

func format(self, cmt string, opts []Option) string {
	if cmt == "" {
		return ""
	}

	cmt = html.UnescapeString(cmt)

	if self != "" {
		switch strings.ToLower(nthWord(cmt, 0)) {
		case "a", "an", "the":
			cmt = popFirstWord(cmt)
		}

		var typeNamed bool
	optLoop:
		for _, opt := range opts {
			switch opt := opt.(type) {
			case originalTypeName:
				if opt != "" && strings.HasSuffix(nthWord(cmt, 0), string(opt)) {
					typeNamed = true
					break optLoop
				}
			}
		}

		switch {
		case typeNamed:
			fallthrough
		case strings.HasPrefix(cmt, "#") && nthWord(cmt, 1) != "":
			// Trim the first word away and replace it with the Go name.
			cmt = self + " " + popFirstWord(cmt)
		case nthWordSimplePresent(cmt, 0):
			cmt = self + " " + lowerFirstWord(cmt)
		default:
			// Trim the word "this" away to make the sentence gramatically
			// correct.
			cmt = strings.TrimPrefix(cmt, "this ")
			cmt = self + ": " + lowerFirstLetter(cmt)
		}
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

	return cmt
}

// ReflowLinesIndent reflows the given cmt paragraphs into idiomatic Go comment
// strings. It is automatically indented.
func ReflowLinesIndent(indentLvl int, cmt string, opts ...Option) string {
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

	complete := strings.Join(lines, "\n")

	for _, opt := range opts {
		if _, ok := opt.(trailingNewLine); ok {
			complete += "\n"
			break
		}
	}

	return complete
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

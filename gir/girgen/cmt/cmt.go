// Package cmt provides functions that parse and render GIR comments into nice
// and conventional Go comments.
package cmt

import (
	"fmt"
	"go/doc"
	"go/doc/comment"
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
	CommentsTabWidth = 4
)

var (
	cmtListRegex      = regexp.MustCompile(`(?m)\n?\n^- `)
	cmtCodeSpanRegex  = regexp.MustCompile("`.*?`")
	cmtReferenceRegex = regexp.MustCompile(`\[(\w+)@(.*?)\]`)
	cmtNamespaceRegex = regexp.MustCompile(`#[A-Z]\w+?[A-Z]`)
	cmtArgumentRegex  = regexp.MustCompile(`@\w+`)
	cmtPrimitiveRegex = regexp.MustCompile(`%\w+`)
	// cmtFunctionRegex  = regexp.MustCompile(`\w+\(\)`)
	cmtHeadingRegex   = regexp.MustCompile(`\n*#+ (.*?)(?: ?#+ ?\{#.*?\})?\n+`)
	cmtHyperlinkRegex = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	cmtCodeblockRegex = regexp.MustCompile(`(?ms)\n*\|\[(?:<!--.*-->)?\n(.*?)\n *\]\|\n*`)

	mdCodeblockRegex = regexp.MustCompile(`(?ms)\n*\x60\x60\x60\w*\n(.*?)\x60\x60\x60\n*`)

	mdEndRegex = regexp.MustCompile(`(?m)\s*(` + strings.Join([]string{
		``,
		"`",
		"^ *```+",
		`^\]\|`,
	}, "|") + `)\z`)
)

// ParamDoc is the documentation for one parameter. It is added by callable
// generators for GoDoc to generate the rest.
type ParamDoc struct {
	Name         string
	Optional     bool
	InfoElements gir.InfoElements
}

// InfoFields contains common fields that a GIR schema type may contain. These
// fields should rarely be used, because they're too magical. All fields inside
// InfoFields are optional.
type InfoFields struct {
	Name       *string
	Attrs      *gir.InfoAttrs
	Elements   *gir.InfoElements
	ParamDocs  []ParamDoc
	ReturnDocs []ParamDoc
}

func getField(value reflect.Value, field string) interface{} {
	v := value.FieldByName(field)
	if v == (reflect.Value{}) {
		return nil
	}

	switch v.Type().Kind() {
	case reflect.Ptr, reflect.Slice:
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
	inf.ParamDocs, _ = getField(value, "ParamDocs").([]ParamDoc)
	inf.ReturnDocs, _ = getField(value, "ReturnDocs").([]ParamDoc)

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
	additionalPrefix string
	paragraphIndent  int
	trailingNewLine  struct{}
	synopsize        struct{}
)

func (overrideSelfName) opts() {}
func (originalTypeName) opts() {}
func (additionalString) opts() {}
func (additionalPrefix) opts() {}
func (paragraphIndent) opts()  {}
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

func searchOptsBool(opts []Option, b Option) bool {
	t := reflect.TypeOf(b)
	for _, opt := range opts {
		if reflect.TypeOf(opt) == t {
			return true
		}
	}
	return false
}

// OverrideSelfName overrides the Go type name that's implicitly converted
// automatically by GoDoc.
func OverrideSelfName(self string) Option { return overrideSelfName(self) }

// AdditionalString adds the given string into the tail of the comment as
// another paragraph.
func AdditionalString(str string) Option { return additionalString(str) }

// AdditionalPrefix adds the given prefix string into the head of the comment.
func AdditionalPrefix(pfx string) Option { return additionalPrefix(pfx) }

// ParagraphIndent indents the inner comment paragraphs.
func ParagraphIndent(indent int) Option { return paragraphIndent(indent) }

// TrailingNewLine adds a trailing new line during documentation generation.
func TrailingNewLine() Option { return trailingNewLine{} }

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
		self = strcases.Go(orig)
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

	synopsize := searchOptsBool(opts, synopsize{})

	docStr := preprocessMarkdown(self, docBuilder.String(), append(opts, originalTypeName(orig)))
	if synopsize && docStr == "" {
		return ""
	}
	docBuilder.Reset()
	docBuilder.WriteString(docStr)

	if len(inf.ParamDocs) > 0 {
		writeParamDocs(&docBuilder,
			"\n\nThe function takes the following parameters:\n",
			inf.ParamDocs)
	}
	if len(inf.ReturnDocs) > 0 {
		writeParamDocs(&docBuilder,
			"\n\nThe function returns the following values:\n",
			inf.ReturnDocs)
	}

	cmt := convertMarkdownStringToGoDoc(docBuilder.String())

	if synopsize {
		printer := &comment.Printer{
			TextWidth: -1, // don't wrap yet
		}
		cmtStr := string(printer.Text(cmt))
		cmtStr = new(doc.Package).Synopsis(cmtStr)
		cmtStr = addPeriod(cmtStr)
		cmt = new(comment.Parser).Parse(cmtStr)
	}

	printer := &comment.Printer{
		// Don't use "\t" in TextPrefix because when calculating
		// .PrintWidth the printer only counts tabs as width=1.
		// Instead use CommentsTabWidth spaces, and count on the final
		// gofmt step to turn them into tabs.
		TextPrefix:     strings.Repeat(" ", CommentsTabWidth*indentLvl) + "// ",
		TextCodePrefix: strings.Repeat(" ", CommentsTabWidth*indentLvl) + "//\t",
	}
	cmtStr := string(printer.Text(cmt))
	cmtStr = transformLines(cmtStr, func(n, d int, line string) string {
		if line == "" && n+1 != d {
			return printer.TextPrefix
		}
		return line
	})

	if !searchOptsBool(opts, trailingNewLine{}) {
		cmtStr = strings.TrimSuffix(cmtStr, "\n")
	}

	return cmtStr
}

func addPeriod(md string) string {
	var suffix string
	if idx := mdEndRegex.FindStringIndex(md); idx != nil && (cmtCodeblockRegex.MatchString(md) || !strings.HasSuffix(md[idx[0]:], "]|")) {
		suffix = md[idx[0]:]
		md = md[:idx[0]]
	}
	if md != "" && !strings.HasSuffix(md, ".") {
		md += "."
	}
	md += suffix
	return md
}

func writeParamDocs(tail *strings.Builder, label string, params []ParamDoc) {
	line1Indent := " - "
	lineNIndent := "   "

	written := false

	for _, param := range params {
		name := param.Name
		if param.Optional {
			name += " (optional)"
		}

		var doc string
		if param.InfoElements.Doc != nil {
			doc = preprocessMarkdown(name, param.InfoElements.Doc.String, []Option{
				originalTypeName(param.Name),
			})
			// Insert a dash space into the lines.
			doc = transformLines(doc, func(i, _ int, line string) string {
				if i == 0 {
					return line1Indent + line
				} else {
					return lineNIndent + line
				}
			})
		} else {
			// If we only have 1 param with no doc, then just give up.
			if len(params) < 2 {
				return
			}
			doc = line1Indent + name
		}

		if !written {
			tail.WriteString(label)
			written = true
		}
		tail.WriteString("\n")
		tail.WriteString(doc)
	}
}

// nthWord returns the nth word, or an empty string if none.
func nthWord(paragraph string, n int) string {
	words := strings.SplitN(paragraph, " ", n+2)
	if len(words) < n+2 {
		return ""
	}
	return words[n]
}

// wordIsSimplePresent returns whether word is a simple present tense verb
// (e.g.: "is", "takes", "emits").
func wordIsSimplePresent(word string) bool {
	return !strings.EqualFold(word, "this") && strings.HasSuffix(word, "s")
}

// trimFirstWord trims the first word off.
func trimFirstWord(paragraph string) string {
	parts := strings.SplitN(paragraph, " ", 2)
	if len(parts) < 2 {
		return ""
	}
	return parts[1]
}

// preprocessMarkdown takes a (GTK-Doc-flavored / GI-DocGen-flavored)
// markdown string makes some stylistic changes to it.  The return
// value is still markdown.
func preprocessMarkdown(self, cmt string, opts []Option) string {
	if cmt == "" {
		return ""
	}

	cmt = html.UnescapeString(cmt)

	if self != "" {
		switch strings.ToLower(nthWord(cmt, 0)) {
		case "a", "an", "the":
			cmt = trimFirstWord(cmt)
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

		firstWord := nthWord(cmt, 0)
		switch {
		case strings.EqualFold(firstWord, "is"), strings.EqualFold(firstWord, "will"):
			// Turn "Will frob the fnord" into "{name} will frob the fnord".
			cmt = self + " " + lowerFirstLetter(cmt)
		case strings.EqualFold(firstWord, "emitted"):
			// Turn "emitted by Emitter" into "{name} is emitted by Emitter".
			cmt = self + " is " + lowerFirstLetter(cmt)
		case typeNamed, strings.HasPrefix(cmt, "#") && nthWord(cmt, 1) != "":
			// Turn "#CName does stuff" into "GoName does stuff"
			cmt = self + " " + trimFirstWord(cmt)
		case wordIsSimplePresent(firstWord):
			// Turn "Frobs the fnord" into "{name} frobs the fnord".
			cmt = self + " " + lowerFirstLetter(cmt)
		default:
			// Trim the word "this" away to make the sentence gramatically
			// correct.
			cmt = self + ": " + lowerFirstLetter(strings.TrimPrefix(cmt, "this "))
		}
	}

	for _, opt := range opts {
		switch opt := opt.(type) {
		case additionalPrefix:
			cmt = string(opt) + cmt
		}
	}

	cmt = strings.TrimSpace(cmt)
	cmt = addPeriod(cmt)

	return cmt
}

// convertMarkdownStringToGoDoc converts a (GTK-Doc-flavored /
// GI-DocGen-flavored) markdown string into a parsed Go
// [*comment.Doc].
func convertMarkdownStringToGoDoc(cmt string) *comment.Doc {
	// Fix up the codeblocks and render it using GoDoc format.
	codeblockFunc := func(re *regexp.Regexp, match string) string {
		matches := re.FindStringSubmatch(match)

		lines := strings.Split(matches[1], "\n")
		for i, line := range lines {
			lines[i] = "   " + line
		}

		// Use our own new lines.
		return "\n\n" + strings.Join(lines, "\n") + "\n\n"
	}
	cmt = cmtCodeblockRegex.ReplaceAllStringFunc(cmt, func(match string) string {
		return codeblockFunc(cmtCodeblockRegex, match)
	})
	cmt = mdCodeblockRegex.ReplaceAllStringFunc(cmt, func(match string) string {
		return codeblockFunc(mdCodeblockRegex, match)
	})

	// Fix up headers in the preprocessing stage. We also sanitize the trailing
	// new lines here.
	cmt = cmtHeadingRegex.ReplaceAllString(cmt, "\n\n$1\n\n")

	// Wipe the namespace prefix syntax.
	cmt = cmtNamespaceRegex.ReplaceAllStringFunc(cmt, func(str string) string {
		// Replace "#?" with just "?".
		return str[len(str)-1:]
	})

	// Put list entries into their own paragraphs.
	cmt = cmtListRegex.ReplaceAllStringFunc(cmt, func(in string) string {
		if strings.HasPrefix(in, "\n\n") {
			return in
		}
		return "\n" + in
	})

	// Unwrap all hyperlinks.
	cmt = cmtHyperlinkRegex.ReplaceAllString(cmt, "$1 ($2)")

	// Fix up new lines before we throw this into ToText so to not confuse it.
	cmt = tidyParagraphs(cmt)

	cmt = cmtCodeSpanRegex.ReplaceAllStringFunc(cmt, func(in string) string {
		return strings.Trim(in, "`")
	})

	cmt = cmtReferenceRegex.ReplaceAllStringFunc(cmt, func(in string) string {
		matches := cmtReferenceRegex.FindStringSubmatch(in)

		rtype := matches[1]
		if rtype == "id" {
			// Keep as original
			return matches[2]
		}

		words := strings.Split(matches[2], ".")
		if len(words) > 0 {
			// Package namespace.
			words[0] = strings.ToLower(words[0])
		}
		if len(words) > 1 {
			// Class/Record type or whatever.
			words[1] = strcases.Go(words[1])
			// Indicate that this is a function if possible.
			switch rtype {
			case "func":
				words[1] += "()"
			}
		}
		if len(words) > 2 {
			// Method name.
			words[2] = strcases.SnakeToGo(true, words[2])
			// Indicate that this is a function if possible.
			switch rtype {
			case "method", "vfunc":
				words[2] += "()"
			}
		}

		return strings.Join(words, ".")
	})

	cmt = cmtPrimitiveRegex.ReplaceAllStringFunc(cmt, func(in string) string { return in[1:] })
	cmt = cmtArgumentRegex.ReplaceAllStringFunc(cmt, func(in string) string { return in[1:] })

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

	return new(comment.Parser).Parse(cmt)
}

func transformLines(cmt string, f func(int, int, string) string) string {
	lines := strings.Split(cmt, "\n")
	for i, line := range lines {
		lines[i] = f(i, len(lines), line)
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

// lowerFirstLetter lower-cases the first letter in the paragraph.
func lowerFirstLetter(p string) string {
	if p == "" {
		return ""
	}

	r1, r1w := utf8.DecodeRuneInString(p)
	if unicode.IsLower(r1) {
		return p
	}
	if r1w == len(p) {
		return strings.ToLower(p)
	}

	// Edge case: gTK, etc.
	if r2, _ := utf8.DecodeRuneInString(p[r1w:]); unicode.IsUpper(r2) {
		return p
	}

	return strings.ToLower(p[:r1w]) + p[r1w:]
}

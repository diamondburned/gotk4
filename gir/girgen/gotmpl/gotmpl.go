// Package gotmpl provides abstractions around text/template to better generate
// Go files.
package gotmpl

import (
	"io"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/gocode"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
)

func NewGoTemplate(block string) *template.Template {
	_, file, _, _ := runtime.Caller(1)
	base := filepath.Base(file)

	t := template.New(base)
	t.Funcs(template.FuncMap{
		"Interfacify":    strcases.Interfacify,
		"PascalToGo":     strcases.PascalToGo,
		"UnexportPascal": strcases.UnexportPascal,
		"KebabToGo":      strcases.KebabToGo,
		"SnakeToGo":      strcases.SnakeToGo,
		"FirstLetter":    strcases.FirstLetter,

		"GoDoc":            cmt.GoDoc,
		"Synopsis":         cmt.Synopsis,
		"OverrideSelfName": cmt.OverrideSelfName,
		"AdditionalString": cmt.AdditionalString,
		"AdditionalPrefix": cmt.AdditionalPrefix,
		"ParagraphIndent":  cmt.ParagraphIndent,
		"TrailingNewLine":  cmt.TrailingNewLine,

		"Quote": strconv.Quote,

		"CoalesceTail": gocode.CoalesceTail,
		"FormatReturn": gocode.FormatReturn,
		"ExtractDefer": gocode.ExtractDefer,
	})
	t = template.Must(t.Parse(block))
	return t
}

var (
	renderTmpls = map[string]*template.Template{}
	tmplMutex   sync.Mutex
)

// M describes a key-value map for a template render.
type M = map[string]any

// Render renders the given template string with the given key-value pair.
func Render(w io.Writer, tmpl string, v any) {
	tmpl = strings.TrimSpace(tmpl) + "\n"

	tmplMutex.Lock()
	renderTmpl, ok := renderTmpls[tmpl]
	if !ok {
		renderTmpl = template.New("(anonymous)")
		renderTmpl = renderTmpl.Delims("<", ">")
		renderTmpl = template.Must(renderTmpl.Parse(tmpl))
		renderTmpls[tmpl] = renderTmpl
	}
	tmplMutex.Unlock()

	if err := renderTmpl.ExecuteTemplate(w, "(anonymous)", v); err != nil {
		log.Panicln("inline render fail:", err)
	}
}

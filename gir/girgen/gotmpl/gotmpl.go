// Package gotmpl provides abstractions around text/template to better generate
// Go files.
package gotmpl

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
)

func NewGoTemplate(block string) *template.Template {
	_, srcFile, _, _ := runtime.Caller(1)
	base := filepath.Base(srcFile)

	t := template.New(base)
	t.Funcs(template.FuncMap{
		"Interfacify":    strcases.Interfacify,
		"PascalToGo":     strcases.PascalToGo,
		"UnexportPascal": strcases.UnexportPascal,
		"KebabToGo":      strcases.KebabToGo,
		"SnakeToGo":      strcases.SnakeToGo,
		"FirstLetter":    strcases.FirstLetter,
		"CGoField":       strcases.CGoField,

		"GoDoc":            cmt.GoDoc,
		"Synopsis":         cmt.Synopsis,
		"OverrideSelfName": cmt.OverrideSelfName,
		"AdditionalString": cmt.AdditionalString,
		"AdditionalPrefix": cmt.AdditionalPrefix,
		"ParagraphIndent":  cmt.ParagraphIndent,
		"TrailingNewLine":  cmt.TrailingNewLine,

		"Quote": func(strs ...interface{}) string {
			return strconv.Quote(fmt.Sprint(strs...))
		},

		"Import":             importFunc((*file.Header).Import),
		"ImportCore":         importFunc((*file.Header).ImportCore),
		"ImportResolvedType": importFunc((*file.Header).ImportResolvedType),
		"DashImport":         importFunc((*file.Header).DashImport),
	})
	t = template.Must(t.Parse(block))
	return t
}

func importFunc[ArgT any](method func(*file.Header, ArgT)) func(file.Headerer, ArgT) string {
	return func(headerer file.Headerer, arg ArgT) string {
		method(headerer.Header(), arg)
		return ""
	}
}

var (
	renderTmpls = map[string]*template.Template{}
	tmplMutex   sync.Mutex
)

// M describes a key-value map for a template render.
type M = map[string]interface{}

// Render renders the given template string with the given key-value pair.
func Render(w io.Writer, tmpl string, v interface{}) {
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

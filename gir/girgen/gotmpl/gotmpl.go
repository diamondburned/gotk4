// Package gotmpl provides abstractions around text/template to better generate
// Go files.
package gotmpl

import (
	"html/template"
	"path/filepath"
	"runtime"

	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
)

func newGoTemplate(block string) *template.Template {
	_, file, _, _ := runtime.Caller(1)
	base := filepath.Base(file)

	t := template.New(base)
	t.Funcs(template.FuncMap{
		"PascalToGo":     strcases.PascalToGo,
		"UnexportPascal": strcases.UnexportPascal,
		"SnakeToGo":      strcases.SnakeToGo,
		"FirstLetter":    strcases.FirstLetter,
		"GoDoc":          cmt.GoDoc,
	})
	t = template.Must(t.Parse(block))
	return t
}

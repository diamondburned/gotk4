package pen

import (
	"fmt"
	"io"
	"log"
	"strings"
	"text/template"

	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
)

// Pen wraps a Pen and its own buffer.
type Pen struct {
	PenWriter
}

// PenWriter describes an interface that a pen can write to.
type PenWriter interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
}

type noopWriter struct{}

func (noopWriter) Write(b []byte) (int, error)       { return len(b), nil }
func (noopWriter) WriteByte(b byte) error            { return nil }
func (noopWriter) WriteString(s string) (int, error) { return len(s), nil }

// NewPen creates a new Pen that preallocates 1KB.
func NewPen(w PenWriter) *Pen {
	return &Pen{w}
}

// NoopPen is a pen that does nothing.
var NoopPen = NewPen(noopWriter{})

// Words writes a list of words into a single line.
func (p *Pen) Words(words ...any) {
	for i, word := range words {
		if i != 0 {
			p.WriteByte(' ')
		}

		switch word := word.(type) {
		case string:
			p.WriteString(word)
		case []string:
			p.WriteString(strings.Join(word, " "))
		default:
			log.Panicf("unknown type %T given", word)
		}
	}

	p.EmptyLine()
}

// Printf writes a Sprintf-formatted string.
func (p *Pen) Printf(f string, v ...any) {
	if len(v) == 0 {
		p.WriteString(f)
	} else {
		fmt.Fprintf(p.PenWriter, f, v...)
	}
}

// Lines writes multiple lines.
func (p *Pen) Lines(lines []string) {
	for _, line := range lines {
		p.Line(line)
	}
}

// Line writes a single line.
func (p *Pen) Line(line string) { p.Linef(line) }

// Linef writes a Sprintf-formatted line.
func (p *Pen) Linef(f string, v ...any) {
	p.Printf(f, v...)
	p.EmptyLine()
}

// EmptyLine adds an empty line.
func (p *Pen) EmptyLine() {
	p.WriteByte('\n')
}

func (p *Pen) Descend() { p.Line("{") }
func (p *Pen) Ascend()  { p.Line("}") }

// WritTmpl writes a template into the pen.
func (p *Pen) WriteTmpl(tmpl *template.Template, args any) {
	if err := tmpl.Execute(p.PenWriter, args); err != nil {
		log.Panicln("template error:", err)
	}
	p.EmptyLine()
}

// LineTmpl writes an inline template with the delimiter "{" and "}".
func (p *Pen) LineTmpl(v any, tmpl string) {
	gotmpl.Render(p.PenWriter, tmpl, v)
}

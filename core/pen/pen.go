package pen

import (
	"fmt"
	"io"
	"log"
	"strings"
	"text/template"
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
func (p *Pen) Words(words ...interface{}) {
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

// Lines writes multiple lines.
func (p *Pen) Lines(lines []string) {
	for _, line := range lines {
		p.Line(line)
	}
}

// Line writes a single line.
func (p *Pen) Line(line string) { p.Linef(line) }

// Linef writes a Sprintf-formatted line.
func (p *Pen) Linef(f string, v ...interface{}) {
	if len(v) == 0 {
		p.WriteString(f)
	} else {
		fmt.Fprintf(p.PenWriter, f, v...)
	}
	p.EmptyLine()
}

// EmptyLine adds an empty line.
func (p *Pen) EmptyLine() {
	p.WriteByte('\n')
}

func (p *Pen) Descend() { p.Line("{") }
func (p *Pen) Ascend()  { p.Line("}") }

// WritTmpl writes a template into the pen.
func (p *Pen) WriteTmpl(tmpl *template.Template, args interface{}) {
	if err := tmpl.Execute(p.PenWriter, args); err != nil {
		log.Panicln("template error:", err)
	}
	p.EmptyLine()
}

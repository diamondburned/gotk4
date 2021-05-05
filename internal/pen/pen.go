package pen

import (
	"bufio"
	"io"
	"strings"
	"text/template"
)

// Pen is an utility writer.
type Pen struct {
	bufio.Writer
	err error
}

// New creates a new Pen.
func New(w io.Writer) *Pen {
	bufw, ok := w.(*bufio.Writer)
	if !ok {
		bufw = bufio.NewWriter(w)
	}

	return &Pen{*bufw, nil}
}

// Block writes a (whitespace-trimmed) block of text and inserts 2 lines.
func (p *Pen) Block(block string) {
	if p.err != nil {
		return
	}

	p.WriteString(strings.TrimSpace(block))
	p.Line()
	p.Line()
}

// BlockTmpl writes a template into the pen.
func (p *Pen) BlockTmpl(tmpl *template.Template, args interface{}) {
	if p.err != nil {
		return
	}

	p.err = tmpl.Execute(&p.Writer, args)
	p.Line()
}

// Words writes a list of words into a single line.
func (p *Pen) Words(words ...string) {
	if p.err != nil {
		return
	}

	for i, word := range words {
		if i != 0 {
			p.WriteByte(' ')
		}

		_, p.err = p.WriteString(word)
		if p.err != nil {
			return
		}
	}

	p.Line()
}

// Line adds a line.
func (p *Pen) Line() {
	if p.err != nil {
		return
	}

	p.err = p.WriteByte('\n')
	return
}

// Flush flushes the internal buffer into the writer.
func (p *Pen) Flush() error {
	if p.err != nil {
		return p.err
	}

	return p.Writer.Flush()
}

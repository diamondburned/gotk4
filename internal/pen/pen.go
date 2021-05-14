// Package pen contains helper functions to work with strings and code
// generation.
package pen

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"text/template"
)

// Piece is a simple string builder.
type Piece struct {
	str strings.Builder
}

// Writef writes using Printf.
func (p *Piece) Writef(f string, v ...interface{}) *Piece {
	fmt.Fprintf(&p.str, f, v...)
	return p
}

// Write writes using Print.
func (p *Piece) Write(v ...interface{}) *Piece {
	if len(v) == 1 {
		if str, ok := v[0].(string); ok {
			p.str.WriteString(str)
			return p
		}
	}

	fmt.Fprint(&p.str, v...)
	return p
}

// Char writes a single ASCII character.
func (p *Piece) Char(b byte) *Piece {
	p.str.WriteByte(b)
	return p
}

// String returns the inner string block.
func (p *Piece) String() string { return p.str.String() }

// Block writes a scoped Go block.
type Block struct {
	str strings.Builder
}

func (b *Block) EmptyLine() { b.Line("") }

// Line writes a line.
func (b *Block) Line(line string) {
	b.str.WriteString(line)
	b.str.WriteByte('\n')
}

// Linef writes a line using Printf.
func (b *Block) Linef(f string, v ...interface{}) {
	fmt.Fprintf(&b.str, f, v...)
	b.str.WriteByte('\n')
}

// String returns the block.
func (b *Block) String() string {
	return "{\n" + b.str.String() + "\n}"
}

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

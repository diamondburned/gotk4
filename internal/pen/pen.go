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

// Joints is a string builder that joins using a separator.
type Joints struct {
	sep  string
	strs []string
}

// NewJoints creates a new Joints instance.
func NewJoints(sep string, cap int) *Joints {
	return &Joints{
		sep:  sep,
		strs: make([]string, 0, cap),
	}
}

// Add adds a new joint.
func (j *Joints) Add(str string) { j.strs = append(j.strs, str) }

// Addf adds a new joint with Sprintf.
func (j *Joints) Addf(f string, v ...interface{}) {
	j.Add(fmt.Sprintf(f, v...))
}

// Join joins the joints.
func (j *Joints) Join() string { return strings.Join(j.strs, j.sep) }

// Piece is a simple string builder with easy chaining.
type Piece struct {
	str strings.Builder
}

// NewPiece returns a new piece.
func NewPiece() *Piece {
	return &Piece{}
}

func (p *Piece) ensureCap() {
	if p.str.Cap() < 4096 {
		p.str.Grow(4096)
	}
}

// Writef writes using Printf.
func (p *Piece) Writef(f string, v ...interface{}) *Piece {
	p.ensureCap()
	fmt.Fprintf(&p.str, f, v...)
	return p
}

// Write writes using Print.
func (p *Piece) Write(v ...interface{}) *Piece {
	p.ensureCap()

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
	p.ensureCap()
	p.str.WriteByte(b)
	return p
}

// Line writes a line.
func (p *Piece) Line(line string) *Piece {
	p.ensureCap()
	p.str.WriteString(line)
	p.str.WriteByte('\n')
	return p
}

// Linef writes a line using Sprintf.
func (p *Piece) Linef(f string, v ...interface{}) {
	p.ensureCap()
	p.Writef(f, v...)
	p.str.WriteByte('\n')
}

// String returns the inner string block.
func (p *Piece) String() string {
	return strings.TrimSuffix(p.str.String(), "\n")
}

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
	return "{\n" + strings.TrimSuffix(b.str.String(), "\n") + "\n}"
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

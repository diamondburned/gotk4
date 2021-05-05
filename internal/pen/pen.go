package pen

import (
	"bufio"
	"io"
	"strings"
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
	p.WriteString(strings.TrimSpace(block))
	p.Line()
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

func (p *Pen) Line() {
	if p.err != nil {
		return
	}

	p.err = p.WriteByte('\n')
	return
}

func (p *Pen) Flush() error {
	if p.err != nil {
		return p.err
	}

	return p.Writer.Flush()
}

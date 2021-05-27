package pen

import (
	"fmt"
	"strings"
)

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
	if len(v) == 0 {
		p.str.WriteString(f)
	} else {
		fmt.Fprintf(&p.str, f, v...)
	}
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

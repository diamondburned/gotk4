package pen

import (
	"fmt"
	"strings"
)

// Block writes a scoped Go block.
type Block struct {
	str  strings.Builder
	nest []Block
}

// NewBlock creates a new preallocated block.
func NewBlock() *Block {
	b := Block{}
	b.str.Grow(4096)
	return &b
}

func (b *Block) EmptyLine() { b.Line("") }

// Line writes a line.
func (b *Block) Line(line string) {
	b.str.WriteString(line)
	b.str.WriteByte('\n')
}

// Linef writes a line using Printf. If v is none, then f is taken literally.
func (b *Block) Linef(f string, v ...interface{}) {
	if len(v) == 0 {
		b.str.WriteString(f)
	} else {
		fmt.Fprintf(&b.str, f, v...)
	}
	b.str.WriteByte('\n')
}

// String returns the block.
func (b *Block) String() string {
	return "{\n" + strings.TrimSuffix(b.str.String(), "\n") + "\n}"
}

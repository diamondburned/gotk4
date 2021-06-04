package pen

import (
	"strings"
)

// Paper wraps a Pen and its own buffer.
type Paper struct {
	Pen
}

// NewPaper creates a new Paper that preallocates 1KB.
func NewPaper() *Paper {
	return NewPaperSize(10240) // 10KB
}

// NewPaperSize creates a new Paper with the given size to preallocate.
func NewPaperSize(size int) *Paper {
	builder := strings.Builder{}
	builder.Grow(size)

	return &Paper{Pen{&builder}}
}

func (p *Paper) builder() *strings.Builder { return p.PenWriter.(*strings.Builder) }

// Len returns the internal length of the buffer.
func (p *Paper) Len() int {
	return p.builder().Len()
}

// String returns the final string written from the Pen.
func (p *Paper) String() string {
	return strings.TrimSuffix(p.builder().String(), "\n")
}

// IsEmpty returns true if the buffer is empty.
func (p *Paper) IsEmpty() bool {
	return p.builder().Len() == 0
}

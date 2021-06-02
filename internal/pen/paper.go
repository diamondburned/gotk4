package pen

import "strings"

// Paper wraps a Pen and its own buffer.
type Paper struct {
	Pen
	buf strings.Builder
}

// NewPaper creates a new Paper that preallocates 2MB.
func NewPaper() *Paper {
	return NewPaperSize(2 * 1024 * 1024) // 2MB
}

// NewPaperSize creates a new Paper with the given size to preallocate.
func NewPaperSize(size int) *Paper {
	p := Paper{}
	p.buf.Grow(size)
	p.Pen = *New(&p.buf)
	return &p
}

// String returns the final string written from the Pen.
func (p *Paper) String() string { return strings.TrimSuffix(p.buf.String(), "\n") }

// IsEmpty returns true if the buffer is empty.
func (p *Paper) IsEmpty() bool { return p.buf.Len() == 0 }

package pen

import (
	"bytes"
	"strings"
)

// PaperString wraps a Pen and its own buffer.
type PaperString struct {
	Pen
}

// NewPaperString creates a new Paper that preallocates 1KB.
func NewPaperString() *PaperString {
	return NewPaperStringSize(10240) // 10KB
}

// NewPaperStringSize creates a new Paper with the given size to preallocate.
func NewPaperStringSize(size int) *PaperString {
	builder := strings.Builder{}
	builder.Grow(size)

	return &PaperString{Pen{&builder}}
}

func (p *PaperString) builder() *strings.Builder { return p.PenWriter.(*strings.Builder) }

// Reset clears the internal buffer.
func (p *PaperString) Reset() {
	p.builder().Reset()
}

// Len returns the internal length of the buffer.
func (p *PaperString) Len() int {
	return p.builder().Len()
}

// String returns the final string written from the Pen.
func (p *PaperString) String() string {
	return strings.TrimSuffix(p.builder().String(), "\n")
}

// IsEmpty returns true if the buffer is empty.
func (p *PaperString) IsEmpty() bool {
	return p.builder().Len() == 0
}

// PaperBuffer wraps a Pen and its own buffer.
type PaperBuffer struct {
	Pen
}

// NewPaperBuffer creates a new Paper that preallocates 1KB.
func NewPaperBuffer() *PaperBuffer {
	return NewPaperBufferSize(10240) // 10KB
}

// NewPaperBufferSize creates a new Paper with the given size to preallocate.
func NewPaperBufferSize(size int) *PaperBuffer {
	builder := bytes.Buffer{}
	builder.Grow(size)

	return &PaperBuffer{Pen{&builder}}
}

func (p *PaperBuffer) builder() *bytes.Buffer { return p.PenWriter.(*bytes.Buffer) }

// Reset clears the internal buffer.
func (p *PaperBuffer) Reset() {
	p.builder().Reset()
}

// Len returns the internal length of the buffer.
func (p *PaperBuffer) Len() int {
	return p.builder().Len()
}

// String returns the copied final string written from the Pen.
func (p *PaperBuffer) String() string {
	return strings.TrimSuffix(p.builder().String(), "\n")
}

// Bytes returns the internal byte slice.
func (p *PaperBuffer) Bytes() []byte {
	return bytes.TrimSuffix(p.builder().Bytes(), []byte("\n"))
}

// IsEmpty returns true if the buffer is empty.
func (p *PaperBuffer) IsEmpty() bool {
	return p.builder().Len() == 0
}

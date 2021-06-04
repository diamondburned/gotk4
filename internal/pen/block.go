package pen

import (
	"fmt"
	"log"
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

// maxBlockSections is the arbitrary maximum number of block sections to be
// made.
const maxBlockSections = 8

// BlockSections is a section writer for writing multiple sections of a block at
// once. It can write 8 sections maximum.
type BlockSections struct {
	strs [maxBlockSections]strings.Builder
	len  int
}

// NewBlockSections preallocates backing arrays for BlockSections for a maximum
// of 8.
func NewBlockSections(preallocs ...int) *BlockSections {
	if len(preallocs) > maxBlockSections {
		log.Panicf("trying to alloc %d sections, only %d allowed", len(preallocs), maxBlockSections)
	}

	var sections BlockSections

	for i, prealloc := range preallocs {
		sections.strs[i].Grow(prealloc)
	}

	return &sections
}

func (sects *BlockSections) grow(sect int) {
	sect++ // len
	if sects.len >= sect {
		return
	}

	if sects.len > maxBlockSections {
		log.Panicf("trying to request %d sections, only %d allowed", sect, maxBlockSections)
	}
	sects.len = sect
}

// Linef writes a Sprintf'd line.
func (sects *BlockSections) Linef(sect int, f string, v ...interface{}) {
	sects.grow(sect)
	if len(v) == 0 {
		sects.strs[sect].WriteString(f)
	} else {
		fmt.Fprintf(&sects.strs[sect], f, v...)
	}
	sects.strs[sect].WriteByte('\n')
}

// Line writes a single line into the given section.
func (sects *BlockSections) Line(sect int, line string) { sects.Linef(sect, line) }

// EmptyLine writes an empty line into the given section.
func (sects *BlockSections) EmptyLine(sect int) {
	sects.grow(sect)
	sects.strs[sect].WriteByte('\n')
}

// Section returns a single section within the block sections.
func (sects *BlockSections) Section(sect int) *BlockSection {
	return &BlockSection{sects, sect}
}

// finalLen estimates the final buffer length. It typically overestimates.
func (sects *BlockSections) finalLen() int {
	var sum int
	for i := 0; i < sects.len; i++ {
		sum += sects.strs[i].Len()
	}

	// Account for new lines.
	if sects.len > 0 {
		sum += (sects.len - 1) * 2
	}

	sum += 4 // {\n\n}

	return sum
}

// String joins the sections together.
func (sects *BlockSections) String() string {
	joined := strings.Builder{}
	joined.Grow(sects.finalLen())

	joined.WriteByte('{')
	joined.WriteByte('\n')

	first := true

	for i := 0; i < sects.len; i++ {
		// Ignore empty blocks.
		if sects.strs[i].Len() == 0 {
			continue
		}

		if !first {
			joined.WriteByte('\n')
			joined.WriteByte('\n')
		} else {
			first = false
		}

		joined.WriteString(
			strings.TrimSuffix(sects.strs[i].String(), "\n"),
		)
	}

	joined.WriteByte('\n')
	joined.WriteByte('}')

	return joined.String()
}

// BlockSection is a section writer that writes into a single section from
// BlockSections.
type BlockSection struct {
	p *BlockSections
	n int
}

// Linef writes a Sprintf'd line.
func (sect *BlockSection) Linef(f string, v ...interface{}) { sect.p.Linef(sect.n, f, v...) }

// Line writes a single line into the section.
func (sect *BlockSection) Line(line string) { sect.p.Line(sect.n, line) }

// EmptyLine write an empty line into the section.
func (sect *BlockSection) EmptyLine() { sect.p.EmptyLine(sect.n) }

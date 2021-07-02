package generators

import (
	"github.com/diamondburned/gotk4/core/pen"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// FileGenerator describes the interface of a file generator. A FileGenerator is
// immutable: anything that takes in a FileGenerator cannot mutate the state of
// it.
type FileGenerator interface {
	types.FileGenerator
}

// FileGeneratorWriter is a FileGenerator that can be mutated or written to.
type FileGeneratorWriter interface {
	FileGenerator
	file.Headerer

	// Pen returns the generator's file writer.
	Pen() *pen.Pen
}

// headeredFileGenerator is used to overried a Header to be used inside
// callable.FileGenerator.
type headeredFileGenerator struct {
	types.FileGenerator
	file.Headerer
}

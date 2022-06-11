package generators

import (
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
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
	// FileWriter returns the file writer for the given source position.
	FileWriter(cmt.InfoFields) FileWriter
	// CHeaderFile returns the C header file.
	CHeaderFile() FileWriter
}

// FileWriterFromType is a convenient function that returns the FileWriter from
// the given GIR type.
func FileWriterFromType(w FileGeneratorWriter, v interface{}) FileWriter {
	return w.FileWriter(cmt.GetInfoFields(v))
}

// FileWriter describes a file that generators can write into and change its
// file header.
type FileWriter interface {
	file.Headerer
	// Pen returns the generator's file writer.
	Pen() *pen.Pen
}

// // headeredFileGenerator is used to overried a Header to be used inside
// // callable.FileGenerator.
// type headeredFileGenerator struct {
// 	types.FileGenerator
// 	file.Headerer
// }

// StubFileGeneratorWriterWriter wraps an existing FileGenerator around a stub
// file writer. This is useful for using existing functions that expect to write
// something, but only the checks are wanted.
func StubFileGeneratorWriter(gen FileGenerator) FileGeneratorWriter {
	return stubFileGeneratorWriter{gen}
}

type (
	stubFileGeneratorWriter struct{ FileGenerator }
	stubFileWriter          struct{}
)

func (s stubFileGeneratorWriter) FileWriter(cmt.InfoFields) FileWriter { return stubFileWriter{} }
func (s stubFileGeneratorWriter) CHeaderFile() FileWriter              { return stubFileWriter{} }

func (s stubFileWriter) Header() *file.Header { return file.NoopHeader }
func (s stubFileWriter) Pen() *pen.Pen        { return pen.NoopPen }

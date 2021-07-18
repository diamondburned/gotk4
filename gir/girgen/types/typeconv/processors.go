package typeconv

import (
	"log"
	"strings"

	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// ConversionProcessor is a processor that can override an entire converter.
type ConversionProcessor interface {
	ProcessConverter(conv *Converter)
}

// ConversionProcessFunc is a function that satisfies the ConversionProcessor
// interface.
type ConversionProcessFunc func(conv *Converter)

// Process calls f.
func (f ConversionProcessFunc) ProcessConverter(conv *Converter) {
	f(conv)
}

// ProcessCallback creates a new conversion processor that calls f on every
// callback type matching the given girType. The GIR type is matched absolutely.
func ProcessCallback(girType string, f func(conv *Converter)) ConversionProcessor {
	parts := strings.SplitN(girType, ".", 2)
	if len(parts) != 2 {
		log.Panicf("missing namespace for AbsoluteFilter %q", girType)
	}

	return ConversionProcessFunc(func(conv *Converter) {
		t, ok := types.EqNamespace(parts[0], conv.Parent.VersionedNamespaceType())
		if ok && t == parts[1] {
			f(conv)
		}
	})
}

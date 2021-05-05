package gir

import (
	"encoding/xml"
	"os"

	"github.com/pkg/errors"
)

var repository struct {
	Includes   []Include   `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	CIncludes  []CInclude  `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Namespaces []Namespace `xml:"http://www.gtk.org/introspection/core/1.0 namespace"`
}

func ParseRepositoryFile(path string) error {
	f, err := os.Open("./Handy-1.gir")
	if err != nil {
		return errors.Wrap(err, "Failed to open file")
	}
	defer f.Close()

	if err := xml.NewDecoder(f).Decode(&repository); err != nil {
		return errors.Wrap(err, "Failed to decode gir XML")
	}

	return nil
}

func Namespaces() []Namespace {
	return repository.Namespaces
}

var activeNamespace namespaceGenerator

// SetActiveNamespace sets the active global namespace to generate from. This
// method is not thread safe. As such, only ONE namespace can be generated at a
// time.
func SetActiveNamespace(i int) namespaceGenerator {
	activeNamespace = namespaceGenerator{&repository.Namespaces[i]}
	return activeNamespace
}

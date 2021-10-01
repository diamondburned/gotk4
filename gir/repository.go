package gir

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/pkg/errors"
)

// Repository represents a GObject Introspection Repository, which contains the
// includes, C includes and namespaces of a single gir file.
type Repository struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 repository"`

	Version             string `xml:"version,attr"`
	CIdentifierPrefixes string `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	CSymbolPrefixes     string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`

	Includes   []Include   `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	CIncludes  []CInclude  `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Packages   []Package   `xml:"http://www.gtk.org/introspection/core/1.0 package"`
	Namespaces []Namespace `xml:"http://www.gtk.org/introspection/core/1.0 namespace"`
}

// ParseRepository parses a repository from the given file path.
func ParseRepository(file string) (*Repository, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer f.Close()

	return ParseRepositoryFromReader(f)
}

// ParseRepositoryFromReader parses a repository from the given reader.
func ParseRepositoryFromReader(r io.Reader) (*Repository, error) {
	var repo Repository

	if err := xml.NewDecoder(r).Decode(&repo); err != nil {
		return nil, errors.Wrap(err, "failed to decode gir XML")
	}

	return &repo, nil
}

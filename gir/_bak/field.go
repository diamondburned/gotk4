package gir

import "encoding/xml"

type Field struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	Name    string   `xml:"name,attr"`
	Type    Type
	Doc     *Doc
}

// TypeName returns the type name or an empty string if Field is nil.
func (f *Field) TypeName() (t string) {
	if f == nil {
		return ""
	}
	return f.Type.Name
}

func (f *Field) GoName() string {
	return f.Type.Map().GoString()
}

package gir

import "encoding/xml"

type Annotation struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Array struct {
	XMLName        xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 array"`
	Length         int      `xml:"http://www.gtk.org/introspection/core/1.0 length"`
	ZeroTerminated int      `xml:"http://www.gtk.org/introspection/core/1.0 zero-terminated"`

	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	Type Type `xml:"http://www.gtk.org/introspection/core/1.0 type"`
}

type CInclude struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Name    string   `xml:"name,attr"`
}

type CallableAttrs struct {
	Name        string `xml:"name,attr"`
	CType       string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CIdentifier string `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`

	Deprecated        int    `xml:"deprecated,attr"`
	DeprecatedVersion string `xml:"deprecated-version,attr"`

	Parameters  *Parameters
	ReturnValue *ReturnValue `xml:"http://www.gtk.org/introspection/core/1.0 return-value"`
	Doc         *Doc
}

type Callback struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	CallableAttrs
}

type Class struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 class"`
	Name    string   `xml:"name,attr"`

	CType         string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CSymbolPrefix string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`

	Parent string `xml:"parent,attr"`

	GLibTypeName   string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType    string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GLibTypeStruct string `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`

	Implements   []Implements  `xml:"http://www.gtk.org/introspection/core/1.0 implements"`
	Constructors []Constructor `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Methods      []Method      `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Fields       []Field       `xml:"http://www.gtk.org/introspection/core/1.0 field"`
}

type Constructor struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	CallableAttrs
}

type Doc struct {
	XMLName  xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 doc"`
	Filename string   `xml:"filename,attr"`
	Line     int      `xml:"line,attr"`
	String   string   `xml:",innerxml"`
}

type Enum struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 enumeration"`
	Name    string   `xml:"name,attr"` // Go case
	Version string   `xml:"version,attr"`

	Doc *Doc

	GLibTypeName string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType  string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`

	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	Members []Member `xml:"http://www.gtk.org/introspection/core/1.0 member"`
}

type Field struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	Name    string   `xml:"name,attr"`
	Type    Type
	Doc     *Doc
}

type Function struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	CallableAttrs
}

type Implements struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 implements"`
	Name    string   `xml:"name,attr"`
}

type Include struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 include"`
	Name    string   `xml:"name,attr"`
	Version *string  `xml:"version,attr"`
}

type InstanceParameter struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 instance-parameter"`
	ParameterAttrs
}

type Interface struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 interface"`
	Name    string   `xml:"name,attr"`

	CType         string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CSymbolPrefix string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`

	GLibTypeName   string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType    string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GLibTypeStruct string `xml:"http://www.gtk.org/introspection/glib/1.0 type-struct,attr"`

	Functions     []Function     `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Methods       []Method       `xml:"http://www.gtk.org/introspection/core/1.0 method"` // translated to Go fns
	Prerequisites []Prerequisite `xml:"http://www.gtk.org/introspection/core/1.0 prerequisite"`
}

type Member struct {
	XMLName     xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 member"`
	Name        string   `xml:"name,attr"`
	Value       int      `xml:"value,attr"`
	CIdentifier string   `xml:"http://www.gtk.org/introspection/c/1.0 identifer,attr"`

	Doc *Doc
}

type Method struct {
	XMLName     xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Name        string   `xml:"name,attr"`
	CIdentifier string   `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	CallableAttrs
}

type Namespace struct {
	XMLName             xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 namespace"`
	Name                string   `xml:"name,attr"`
	Version             string   `xml:"version,attr"`
	SharedLibrary       string   `xml:"shared-library,attr"`
	CIdentifierPrefixes string   `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	CSymbolPrefixes     string   `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`

	Classes     []Class      `xml:"http://www.gtk.org/introspection/core/1.0 class"`
	Records     []Record     `xml:"http://www.gtk.org/introspection/core/1.0 record"`
	Enums       []Enum       `xml:"http://www.gtk.org/introspection/core/1.0 enumeration"`
	Functions   []Function   `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Callbacks   []Callback   `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	Interfaces  []Interface  `xml:"http://www.gtk.org/introspection/core/1.0 interface"`
	Annotations []Annotation `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
}

type Package struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 package"`
	Name    string   `xml:"name,attr"`
}

type Parameter struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 parameter"`
	ParameterAttrs
}

type ParameterAttrs struct {
	Name      string `xml:"name,attr"`
	AllowNone int    `xml:"allow-none,attr"` // 1 == true?
	TransferOwnership
	Type Type
	Doc  *Doc
}

type Parameters struct {
	XMLName           xml.Name           `xml:"http://www.gtk.org/introspection/core/1.0 parameters"`
	InstanceParameter *InstanceParameter `xml:"http://www.gtk.org/introspection/core/1.0 instance-parameter"`
	Parameters        []Parameter        `xml:"http://www.gtk.org/introspection/core/1.0 parameter"`
}

type Prerequisite struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 prerequisite"`
	Name    string   `xml:"name,attr"`
}

type Record struct {
	XMLName              xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 record"`
	Name                 string   `xml:"name,attr"`
	CType                string   `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Disguised            bool     `xml:"disguised,attr"`
	GLibTypeName         string   `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType          string   `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	CSymbolPrefix        string   `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	Foreign              bool     `xml:"foreign,attr"`
	GLibIsGTypeStructFor string   `xml:"http://www.gtk.org/introspection/glib/1.0 is-gtype-struct-for,attr"`
	Introspectable       bool     `xml:"introspectable,attr"`
	Deprecated           string   `xml:"deprecated,attr"`
	DeprecatedVersion    string   `xml:"deprecated-version,attr"`
	Version              string   `xml:"version,attr"`
	Stability            string   `xml:"stability,attr"`

	Fields       []Field       `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	Functions    []Function    `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Unions       []Union       `xml:"http://www.gtk.org/introspection/core/1.0 union"`
	Methods      []Method      `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Constructors []Constructor `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Properties   []Property    `xml:"http://www.gtk.org/introspection/core/1.0 property"`
	Annotations  []Annotation  `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
}

type Union struct{}

type Property struct{}

type ReturnValue struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 return-value"`
	TransferOwnership
	Doc *Doc

	// Possible enums.
	Type  *Type
	Array *Array
}

type TransferOwnership struct {
	TransferOwnership *string `xml:"transfer-ownership,attr"`
}

type Type struct {
	XMLName        xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 type"`
	Name           string   `xml:"name,attr"`
	CType          string   `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Introspectable *bool    `xml:"introspectable,attr"`
}

package gir

import "encoding/xml"

type Alias struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 alias"`

	Name  string `xml:"name,attr"`
	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	InfoAttrs
	InfoElements

	Type Type
}

type AnyType struct {
	// Possible variants.
	Type    *Type    `xml:"http://www.gtk.org/introspection/core/1.0 type"`
	Array   *Array   `xml:"http://www.gtk.org/introspection/core/1.0 array"`
	VarArgs *VarArgs `xml:"http://www.gtk.org/introspection/core/1.0 varargs"`
}

type Annotation struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Array struct {
	XMLName        xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 array"`
	Name           string   `xml:"name,attr"`
	ZeroTerminated *bool    `xml:"zero-terminated,attr"`
	FixedSize      int      `xml:"fixed-size,attr"`
	Introspectable bool     `xml:"introspectable,attr"`
	Length         *int     `xml:"length,attr"` // ix of .Parameters
	CType          string   `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	AnyType
}

// IsZeroTerminated returns true if the Array is zero-terminated. It accounts
// for edge cases of the structure.
func (a Array) IsZeroTerminated() bool {
	return a.Name == "" && (a.ZeroTerminated == nil || *a.ZeroTerminated)
}

type Bitfield struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 bitfield"`

	Name         string `xml:"name,attr"` // Go case
	CType        string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GLibTypeName string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType  string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`

	Members   []Member   `xml:"http://www.gtk.org/introspection/core/1.0 member"`
	Functions []Function `xml:"http://www.gtk.org/introspection/core/1.0 function"`

	InfoAttrs
	InfoElements
}

type Boxed struct{}

type CInclude struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/c/1.0 include"`
	Name    string   `xml:"name,attr"`
}

type CallableAttrs struct {
	Name        string       `xml:"name,attr"`
	CIdentifier string       `xml:"http://www.gtk.org/introspection/c/1.0 identifier,attr"`
	ShadowedBy  string       `xml:"shadowed-by,attr"`
	Shadows     string       `xml:"shadows,attr"`
	Throws      bool         `xml:"throws,attr"`
	MovedTo     string       `xml:"moved-to,attr"`
	Parameters  *Parameters  `xml:"http://www.gtk.org/introspection/core/1.0 parameters"`
	ReturnValue *ReturnValue `xml:"http://www.gtk.org/introspection/core/1.0 return-value"`
	InfoAttrs
	InfoElements
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

	InfoAttrs
	InfoElements

	Implements     []Implements    `xml:"http://www.gtk.org/introspection/core/1.0 implements"`
	Constructors   []Constructor   `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Methods        []Method        `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	VirtualMethods []VirtualMethod `xml:"http://www.gtk.org/introspection/core/1.0 virtual-method"`
	Fields         []Field         `xml:"http://www.gtk.org/introspection/core/1.0 field"`
}

type Constant struct{}

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

type DocElements struct {
	Doc            *Doc
	SourcePosition *SourcePosition
}

type Enum struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 enumeration"`

	Name            string `xml:"name,attr"` // Go case
	CType           string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	GLibTypeName    string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType     string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`
	GLibErrorDomain string `xml:"http://www.gtk.org/introspection/glib/1.0 error-domain,attr"`

	Members   []Member   `xml:"http://www.gtk.org/introspection/core/1.0 member"`
	Functions []Function `xml:"http://www.gtk.org/introspection/core/1.0 function"`

	InfoAttrs
	InfoElements
}

type Field struct {
	XMLName  xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	Name     string   `xml:"name,attr"`
	Private  bool     `xml:"private,attr"`
	Writable bool     `xml:"writable,attr"`
	Readable bool     `xml:"readable,attr"`
	Bits     int      `xml:"bits,attr"`
	AnyType
	Doc *Doc
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
	Version string   `xml:"version,attr"`
}

type InfoAttrs struct {
	Introspectable    *bool  `xml:"introspectable,attr"` // default true
	Deprecated        string `xml:"deprecated,attr"`
	DeprecatedVersion string `xml:"deprecated-version,attr"`
	Version           string `xml:"version,attr"`
	Stability         string `xml:"stability,attr"`
}

// IsIntrospectable returns true if the InfoAttrs indicates that the type is
// introspectable.
func (inf InfoAttrs) IsIntrospectable() bool {
	return inf.Introspectable == nil || *inf.Introspectable
}

type InfoElements struct {
	DocElements
	Annotations []Annotation `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
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

	Functions      []Function      `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Methods        []Method        `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	VirtualMethods []VirtualMethod `xml:"http://www.gtk.org/introspection/core/1.0 virtual-method"`
	Prerequisites  []Prerequisite  `xml:"http://www.gtk.org/introspection/core/1.0 prerequisite"`

	InfoAttrs
	InfoElements
}

type Member struct {
	XMLName     xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 member"`
	Name        string   `xml:"name,attr"`
	Value       string   `xml:"value,attr"`
	CIdentifier string   `xml:"http://www.gtk.org/introspection/c/1.0 identifer,attr"`
	GLibNick    string   `xml:"http://www.gtk.org/introspection/glib/1.0 nick,attr"`

	InfoAttrs
	InfoElements
}

type Method struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	CallableAttrs
}

type Namespace struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 namespace"`

	Name                string `xml:"name,attr"`
	Version             string `xml:"version,attr"`
	CIdentifierPrefixes string `xml:"http://www.gtk.org/introspection/c/1.0 identifier-prefixes,attr"`
	CSymbolPrefixes     string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefixes,attr"`
	Prefix              string `xml:"http://www.gtk.org/introspection/c/1.0 prefix,attr"`
	SharedLibrary       string `xml:"shared-library,attr"`

	Aliases     []Alias      `xml:"http://www.gtk.org/introspection/core/1.0 alias"`
	Classes     []Class      `xml:"http://www.gtk.org/introspection/core/1.0 class"`
	Interfaces  []Interface  `xml:"http://www.gtk.org/introspection/core/1.0 interface"`
	Records     []Record     `xml:"http://www.gtk.org/introspection/core/1.0 record"`
	Enums       []Enum       `xml:"http://www.gtk.org/introspection/core/1.0 enumeration"`
	Functions   []Function   `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Unions      []Union      `xml:"http://www.gtk.org/introspection/core/1.0 union"`
	Bitfields   []Bitfield   `xml:"http://www.gtk.org/introspection/core/1.0 bitfield"`
	Callbacks   []Callback   `xml:"http://www.gtk.org/introspection/core/1.0 callback"`
	Constants   []Constant   `xml:"http://www.gtk.org/introspection/core/1.0 constant"`
	Annotations []Annotation `xml:"http://www.gtk.org/introspection/core/1.0 attribute"`
	Boxeds      []Boxed      `xml:"http://www.gtk.org/introspection/core/1.0 boxed"`
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
	Name            string `xml:"name,attr"`
	Direction       string `xml:"direction,attr"`
	CallerAllocates bool   `xml:"caller-allocates,attr"`
	Closure         *int   `xml:"closure,attr"`
	Destroy         *int   `xml:"destroy,attr"`
	Scope           string `xml:"scope,attr"`
	Skip            bool   `xml:"skip,attr"`
	Nullable        bool   `xml:"nullable,attr"`
	Optional        bool   `xml:"optional,attr"`

	TransferOwnership
	AnyType
	Doc *Doc
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

type Property struct{}

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

	Fields       []Field       `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	Functions    []Function    `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Unions       []Union       `xml:"http://www.gtk.org/introspection/core/1.0 union"`
	Methods      []Method      `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Constructors []Constructor `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Properties   []Property    `xml:"http://www.gtk.org/introspection/core/1.0 property"`

	InfoAttrs
	InfoElements
}

type ReturnValue struct {
	XMLName        xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 return-value"`
	Introspectable bool     `xml:"introspectable,attr"`
	Nullable       bool     `xml:"nullable,attr"`
	Scope          string   `xml:"scope,attr"`
	Skip           bool     `xml:"skip,attr"`
	AllowNone      bool     `xml:"allow-none,attr"`
	Closure        *int     `xml:"closure,attr"`
	Destroy        *int     `xml:"destroy,attr"`
	TransferOwnership
	DocElements
	AnyType
}

type SourcePosition struct {
	XMLName  xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 source-position"`
	Filename string   `xml:"filename,attr"`
	Line     int      `xml:"line,attr"`
	Column   int      `xml:"column,attr"`
}

type TransferOwnership struct {
	TransferOwnership string `xml:"transfer-ownership,attr"`
}

type Type struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 type"`

	Name           string `xml:"name,attr"`
	CType          string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	Introspectable *bool  `xml:"introspectable,attr"`

	DocElements
	AnyType
}

func (typ Type) IsIntrospectable() bool {
	return typ.Introspectable == nil || *typ.Introspectable
}

type Union struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 union"`

	Name          string `xml:"name,attr"` // Go case
	CType         string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`
	CSymbolPrefix string `xml:"http://www.gtk.org/introspection/c/1.0 symbol-prefix,attr"`
	GLibTypeName  string `xml:"http://www.gtk.org/introspection/glib/1.0 type-name,attr"`
	GLibGetType   string `xml:"http://www.gtk.org/introspection/glib/1.0 get-type,attr"`

	InfoAttrs
	InfoElements

	Fields       []Field       `xml:"http://www.gtk.org/introspection/core/1.0 field"`
	Constructors []Constructor `xml:"http://www.gtk.org/introspection/core/1.0 constructor"`
	Methods      []Method      `xml:"http://www.gtk.org/introspection/core/1.0 method"`
	Functions    []Function    `xml:"http://www.gtk.org/introspection/core/1.0 function"`
	Records      []Record      `xml:"http://www.gtk.org/introspection/core/1.0 record"`
}

type VarArgs struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 varargs"`
}

type VirtualMethod struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 virtual-method"`

	Invoker string `xml:"invoker,attr"`
	CallableAttrs
}

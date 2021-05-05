package gir

import (
	"encoding/xml"

	"github.com/dave/jennifer/jen"
)

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

// UserDataParameter returns a non-nil parameter if the given index points to a
// valid user data parameter.
func (c CallableAttrs) UserDataParameter(i int) *Parameter {
	if i < 1 || c.Parameters == nil {
		return nil
	}

	var callback = c.Parameters.Parameters[i-1]
	if callback.Type.IsNamespaceFunc() {
		return &callback
	}

	return nil
}

func (c CallableAttrs) HasInstanceParameter() bool {
	return c.Parameters != nil && c.Parameters.HasInstanceParameter()
}

// HasArrayParameter returns true if the function has an array as a parameter or
// return value.
func (c CallableAttrs) HasArrayParameter() bool {
	if c.ReturnValue == nil {
		return false
	}

	return c.ReturnValue.Array != nil
}

// IsDeprecated returns true if the current object is deprecated.
func (c CallableAttrs) IsDeprecated() bool {
	return c.Deprecated == 1
}

// IsVariadic returns true if the current function is variadic.
func (c CallableAttrs) IsVariadic() bool {
	return c.Parameters != nil && c.Parameters.IsVariadic()
}

// IsBlocked returns true if the current function contains a blocked parameter
// type.
func (c CallableAttrs) IsBlocked() bool {
	return c.Parameters != nil && c.Parameters.HasBlockedType()
}

var ignoredCallables = []func(CallableAttrs) bool{
	CallableAttrs.IsBlocked,
	CallableAttrs.IsVariadic,
	CallableAttrs.HasArrayParameter, // TODO support arrays
}

func (c CallableAttrs) IsIgnored() bool {
	for _, isIgnored := range ignoredCallables {
		if isIgnored(c) {
			return true
		}
	}

	return false
}

type Parameters struct {
	XMLName           xml.Name           `xml:"http://www.gtk.org/introspection/core/1.0 parameters"`
	InstanceParameter *InstanceParameter `xml:"http://www.gtk.org/introspection/core/1.0 instance-parameter"`
	Parameters        []Parameter        `xml:"http://www.gtk.org/introspection/core/1.0 parameter"`
}

// HasInstanceParameter returns true if p and InstanceParameter are not nil.
func (p *Parameters) HasInstanceParameter() bool {
	return p != nil && p.InstanceParameter != nil && !p.InstanceParameter.IsIgnored()
}

// IsVariadic returns if the list of parameters contain a variadic parameter.
func (p Parameters) IsVariadic() bool {
	for _, param := range p.Parameters {
		if param.Name == "..." {
			return true
		}
	}

	return false
}

// HasBlockedType returns if the list of parameters contain a parameter of
// blocked types.
func (p Parameters) HasBlockedType() bool {
	for _, param := range p.Parameters {
		if param.IsBlockedType() {
			return true
		}
	}

	return false
}

// SearchUserData searches for the UserData parameter. It returns nil if p is
// nil or if userData is not in the list of parameters.
func (p *Parameters) SearchUserData() *Parameter {
	if p == nil {
		return nil
	}

	for _, param := range p.Parameters {
		if param.IsUserData() {
			return &param
		}
	}

	return nil
}

type InstanceParameter struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 instance-parameter"`
	ParameterAttrs
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

func (p ParameterAttrs) GoName() string {
	return snakeToGo(false, p.Name)
}

// func (p ParameterAttrs) IsInterface() bool {}

// IsVariadic returns true if the current parameter is variadic.
func (p ParameterAttrs) IsVariadic() bool {
	return p.Name == "..."
}

var ignoredParams = []func(ParameterAttrs) bool{
	ParameterAttrs.IsVariadic,
	ParameterAttrs.IsUserData,
	ParameterAttrs.IsUserDataFreeFunc,
	ParameterAttrs.IsNotNamespaceFunc,
}

func (p ParameterAttrs) IsIgnored() bool {
	for _, isIgnored := range ignoredParams {
		if isIgnored(p) {
			return true
		}
	}

	return false
}

func (p ParameterAttrs) IsUserData() bool {
	return p.Name == "user_data" && p.Type.Name == "gpointer"
}

func (p ParameterAttrs) IsUserDataFreeFunc() bool {
	return p.Type.Name == "GLib.DestroyNotify"
}

// IsNotNamespaceFunc returns true if the parameter is a function and is not the
// current namespace's callback.
func (p ParameterAttrs) IsNotNamespaceFunc() bool {
	return p.Type.IsFunc() && !p.Type.IsNamespaceFunc()
}

func (p ParameterAttrs) IsDestroyNotifyFunc() bool {
	return p.Type.Name == "GLib.DestroyNotify"
}

// IsBlockedType returns true if the parameter has a blocked type.
func (p ParameterAttrs) IsBlockedType() bool {
	return p.Type.Map() == nil
}

// GenValueCall generates a value conversion call from the given names. If
// argName is nil, then the returned value will be a zero-value. Specifically
// for UserData, a callback should be passed in as an argument.
func (p ParameterAttrs) GenValueCall(argName, valueName *jen.Statement) *jen.Statement {
	// Filter out ignored parameters.
	if argName == nil {
		return nil
	}

	var stmt = jen.Add(valueName).Op(":=")

	// Handle special cases.
	switch {
	case p.IsUserDataFreeFunc():
		return stmt.Add(CallbackGenDelete())
	case p.IsUserData():
		return stmt.Add(CallbackGenAssign(argName))
	case p.IsNotNamespaceFunc():
		return stmt.Nil()
	default:
		stmt.Add(p.Type.GenCCaster(argName))
	}

	if p.Type.CNeedsFree() {
		stmt.Line()
		stmt.Defer().Qual("C", "free").Call(jen.Qual("unsafe", "Pointer").Call(valueName))
	}

	// if p.Type.IsFunc() {
	// }

	return stmt
}

type ReturnValue struct {
	XMLName xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 return-value"`
	TransferOwnership
	Doc *Doc

	// Possible enums.
	Type  *Type
	Array *Array
}

type Array struct {
	XMLName        xml.Name `xml:"http://www.gtk.org/introspection/core/1.0 array"`
	Length         int      `xml:"http://www.gtk.org/introspection/core/1.0 length"`
	ZeroTerminated int      `xml:"http://www.gtk.org/introspection/core/1.0 zero-terminated"`

	CType string `xml:"http://www.gtk.org/introspection/c/1.0 type,attr"`

	Type Type `xml:"http://www.gtk.org/introspection/core/1.0 type"`
}

// IsVoid returns true if the type name is "none" or if *ReturnValue is nil.
func (r *ReturnValue) IsVoid() bool {
	if r == nil {
		return true
	}

	return r.Type.Name == "none"
}

// GenReturn generates a statement with the return token.
func (r *ReturnValue) GenReturnFunc(call *jen.Statement) *jen.Statement {
	if r.IsVoid() {
		return call
	}

	v := jen.Id("r")
	return jen.Add(r.Type.GenCaster(v, call)).Line().Return(v)
}

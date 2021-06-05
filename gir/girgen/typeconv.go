package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

// TODO:
//   - support OutputIsParameter
//   - support CallerAllocates

// ValueProp describes the generic properties of a Go or C value for conversion.
type ValueProp struct {
	In    string
	Out   string
	Type  gir.AnyType
	Owner gir.TransferOwnership

	// Closure marks the user_data argument. If this is provided, then the
	// conversion function will set the parameter to the callback ID. The caller
	// is responsible for skipping conversion of these indices.
	Closure *int

	// Destroy marks the callback to destroy the user_data argument. If this is
	// provided, then callbackDelete will be set along with Closure.
	Destroy *int

	// ParameterIndex explicitly gives this value an index used for matching
	// with the given index clues from the GIR files, such as closure, destroy
	// or length.
	ParameterIndex *int

	// OutputIsParameter makes the conversion treat this value like an output
	// parameter. Specifically, the output will be dereferenced when it is set.
	OutputIsParameter bool

	// CallerAllocates determines if the converter should take care of
	// allocating the type or not.
	CallerAllocates bool

	// AllowNone, if true, will allow types that cannot be converted to stay.
	AllowNone bool

	// internal state
	*valuePropState
}

// errorValueProp is an invalid GoValueProp returned when valueAt errors out.
var errorValueProp = ValueProp{
	In:  "NotAvailable",
	Out: "NotAvailable",
	Type: gir.AnyType{
		Type: &gir.Type{
			Name:  "none",
			CType: "void",
		},
	},
}

// NewValuePropParam creates a ValueProp from the given parameter attribute.
func NewValuePropParam(in, out string, i *int, param gir.ParameterAttrs) ValueProp {
	return ValueProp{
		In:                in,
		Out:               out,
		Type:              param.AnyType,
		Owner:             param.TransferOwnership,
		Closure:           param.Closure,
		Destroy:           param.Destroy,
		AllowNone:         param.AllowNone,
		CallerAllocates:   param.CallerAllocates,
		OutputIsParameter: param.Direction == "out",
		ParameterIndex:    i,
	}
}

// NewValuePropReturn creates a new ValueProp from teh given return attribute.
func NewValuePropReturn(in, out string, ret gir.ReturnValue) ValueProp {
	return ValueProp{
		In:        in,
		Out:       out,
		Type:      ret.AnyType,
		Owner:     ret.TransferOwnership,
		AllowNone: ret.Skip || ret.AllowNone,
		Closure:   ret.Closure,
		Destroy:   ret.Destroy,
	}
}

// IsZero returns true if ValueProp is empty.
func (value *ValueProp) IsZero() bool {
	return value.In == "" || value.Out == ""
}

func (value *ValueProp) loadIgnore(ignores map[int]struct{}) {
	// These are handled below.
	if value.Closure != nil {
		ignores[*value.Closure] = struct{}{}
	}
	if value.Destroy != nil {
		ignores[*value.Destroy] = struct{}{}
	}
	if value.Type.Array != nil && value.Type.Array.Length != nil {
		ignores[*value.Type.Array.Length] = struct{}{}
	}
}

// outDeref returns true if the output is dereferenced.
func (value *ValueProp) outDeref() bool {
	return strings.HasPrefix(value.Out, "*")
}

// inner is used only for arrays.
func (value ValueProp) inner(in, out string) ValueProp {
	if value.Type.Array == nil {
		return value
	}

	// If the array's ownership is ONLY container, then we must not take over
	// the inner values. Therefore, we only generate the appropriate code.
	if value.Owner.TransferOwnership == "container" {
		value.Owner.TransferOwnership = "none"
	}

	prop := ValueProp{
		In:    in,
		Out:   out,
		Type:  value.Type.Array.AnyType,
		Owner: value.Owner,
	}
	prop.initState()

	return prop
}

// valuePropState wraps around ValueProp for internal use.
type valuePropState struct {
	// ConversionExtras are filled up after Convert. This is useful for the
	// caller to reuse information derived during conversion.
	ConversionExtras

	resolved       *ResolvedType // only for type conversions
	needsNamespace bool

	p       *pen.PaperString
	inDecl  *pen.PaperString // ONLY USE FOR OutputParam.
	outDecl *pen.PaperString
}

// ConversionExtras contains extra information obtained during conversion.
type ConversionExtras struct {
	InType  string
	OutType string
}

func (value *ValueProp) initState() {
	value.valuePropState = &valuePropState{
		p:       pen.NewPaperStringSize(2048), // 2KB
		inDecl:  pen.NewPaperStringSize(128),  // 0.1KB
		outDecl: pen.NewPaperStringSize(128),  // 0.1KB
	}
}

// resolveType resolves the value type to the resolved field. If inputC is true,
// then the input type is set to the CGo type, otherwise the Go type is set.
func (value *ValueProp) resolveType(conv *conversionTo, inputC bool) bool {
	if value.Type.Type == nil {
		return false
	}

	value.resolved = conv.ng.ResolveType(*value.Type.Type)
	if value.resolved == nil {
		return false
	}

	value.needsNamespace = value.resolved.NeedsNamespace(conv.ng.current)
	if value.needsNamespace {
		conv.sides.addImportAlias(value.resolved.Import, value.resolved.Package)
	}

	cgoType := value.resolved.CGoType()
	if !inputC && value.outDeref() {
		// Output is dereferenced; trim a pointer from the CGo type.
		cgoType = strings.TrimPrefix(cgoType, "*")
	}

	goType := value.resolved.PublicType(value.needsNamespace)

	if inputC {
		value.InType = cgoType
		value.OutType = goType
	} else {
		value.OutType = cgoType
		value.InType = goType
	}

	value.inDecl.Linef("var %s %s", value.In, value.InType)
	value.outDecl.Linef("var %s %s", value.Out, value.OutType)

	return true
}

// isTransferring is true when the ownership is either full or container. If the
// converter code isn't generating for an array, then distinguishing this
// doesn't matter. If the caller hasn't set the ownership yet, then it is
// assumed that we're not getting the ownership, therefore false is returned.
//
// If the generating code is an array, and the conversion is being passed into
// the same generation routine for the inner type, then the ownership should be
// turned into "none" just for that inner type. See TypeConversion.inner().
func (prop ValueProp) isTransferring() bool {
	return false ||
		prop.Owner.TransferOwnership == "full" ||
		prop.Owner.TransferOwnership == "container"
}

// cgoSetObject generates a glib.Take or glib.AssumeOwnership into a new
// function.
func (prop ValueProp) cgoSetObject(ifaceType string) string {
	var gobjectFunction string
	if prop.isTransferring() {
		// Full or container means we implicitly own the object, so we must
		// not take another reference.
		gobjectFunction = "AssumeOwnership"
	} else {
		// Else the object is either unowned by us or it's a floating
		// reference. Take our own or sink the object.
		gobjectFunction = "Take"
	}

	return fmt.Sprintf(
		"%s = gextras.CastObject(externglib.%s(unsafe.Pointer(%s.Native()))).(%s)",
		prop.Out, gobjectFunction, prop.In, ifaceType,
	)
}

// TypeConverted is the result of conversion for a single value.
type TypeConverted struct {
	*ValueProp

	InDeclare  string
	OutDeclare string
	Conversion string
	ConversionSideEffects
}

func (conved *TypeConverted) WriteAll(in, out, conv *pen.BlockSection) {
	if in != nil {
		in.Linef(conved.InDeclare)
	}
	if out != nil {
		out.Linef(conved.OutDeclare)
	}
	if conv != nil {
		conv.Linef(conved.Conversion)
	}
}

// ConversionSideEffects describes the side effects of the conversion, such as
// importing new things or modifying the Cgo preamble.
type ConversionSideEffects struct {
	Imports        map[string]string
	Callbacks      []string
	CallbackDelete bool
	NeedsStdBool   bool
}

func (sides *ConversionSideEffects) addImport(path string) {
	sides.addImportAlias(path, "")
}

func (sides *ConversionSideEffects) addImportAlias(path, alias string) {
	if sides.Imports == nil {
		sides.Imports = map[string]string{}
	}

	sides.Imports[path] = alias
}

func (sides *ConversionSideEffects) addGLibImport() {
	resolved := externGLibType("", gir.Type{}, "")
	sides.addImportAlias(resolved.Import, resolved.Package)
}

func (sides ConversionSideEffects) addCallback(callback *gir.Callback) {
	sides.Callbacks = append(sides.Callbacks, CallbackCHeader(callback))
}

// Apply applies the side effects of the conversion. The caller has control over
// calling this.
func (sides ConversionSideEffects) Apply(fg *FileGenerator) {
	if sides.CallbackDelete {
		fg.needsCallbackDelete()
	}
	if sides.NeedsStdBool {
		fg.needsStdbool()
	}
	for path, alias := range sides.Imports {
		fg.addImportAlias(path, alias)
	}
	for _, callback := range sides.Callbacks {
		fg.addCallbackHeader(callback)
	}
}

type conversionTo struct {
	ng     *NamespaceGenerator
	logger lineLogger
	parent string

	// conversion state
	sides  ConversionSideEffects
	failed bool
}

func newConversionTo(fg *FileGenerator, parent string) conversionTo {
	return conversionTo{
		ng:     fg.parent,
		logger: fg,
		parent: parent,
	}
}

func (conv *conversionTo) reset() {
	conv.sides = ConversionSideEffects{}
	conv.failed = false
}

func (conv *conversionTo) fail() { conv.failed = true }

func (conv *conversionTo) logFail(lvl LogLevel, v ...interface{}) {
	if conv.parent != "" {
		v2 := make([]interface{}, 0, 2+len(v))
		v2 = append(v2, "in", conv.parent)
		v2 = append(v2, v...)

		v = v2
	}

	conv.logger.Logln(lvl, v...)
	conv.fail()
}

func (conv *conversionTo) typeHasPtr(typ *ResolvedType) bool {
	// use .parent to prevent importing
	return TypeHasPointer(conv.ng, typ)
}

// goSliceFromPtr crafts a typ slice from the given ptr as the backing array
// with the given len, then set it into target. typ should be innerType. A
// temporary variable named sliceHeader is made.
//
// Imports needed: github.com/diamondburned/gotk4/internal/ptr.
func goSliceFromPtr(target, ptr, len string) string {
	return fmt.Sprintf(
		"ptr.SetSlice(unsafe.Pointer(&%s), unsafe.Pointer(%s), int(%s))",
		target, ptr, len,
	)
}

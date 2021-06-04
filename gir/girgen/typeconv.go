package girgen

import (
	"fmt"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

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

	// AllowNone, if true, will allow types that cannot be converted to stay.
	AllowNone bool
}

var (
	// errorValueProp is an invalid GoValueProp returned when valueAt errors out.
	errorValueProp = ValueProp{
		In:  "NotAvailable",
		Out: "NotAvailable",
		Type: gir.AnyType{
			Type: &gir.Type{
				Name:  "none",
				CType: "void",
			},
		},
	}
)

// NewValuePropParam creates a ValueProp from the given parameter attribute.
func NewValuePropParam(in, out string, param gir.ParameterAttrs) ValueProp {
	return ValueProp{
		In:        in,
		Out:       out,
		Type:      param.AnyType,
		Owner:     param.TransferOwnership,
		Closure:   param.Closure,
		Destroy:   param.Destroy,
		AllowNone: param.AllowNone,
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

	return ValueProp{
		In:    in,
		Out:   out,
		Type:  value.Type.Array.AnyType,
		Owner: value.Owner,
	}
}

func (value *ValueProp) loadIgnore(ignores *[]bool) {
	// These are handled below.
	if value.Closure != nil {
		(*ignores)[*value.Closure] = true
	}
	if value.Destroy != nil {
		(*ignores)[*value.Destroy] = true
	}
	if value.Type.Array != nil && value.Type.Array.Length != nil {
		(*ignores)[*value.Type.Array.Length] = true
	}
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

// cgoCreateObject generates a glib.Take or glib.AssumeOwnership into a new
// function.
func (prop ValueProp) cgoCreateObject(ifaceType string) string {
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
		"%s := gextras.CastObject(externglib.%s(unsafe.Pointer(%s.Native()))).(%s)",
		prop.Out, gobjectFunction, prop.In, ifaceType,
	)
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
	p      *pen.Paper
	logger lineLogger
	parent string

	// conversion state
	sides  ConversionSideEffects
	failed bool
}

func newConversionTo(fg *FileGenerator, parent string) conversionTo {
	return conversionTo{
		ng:     fg.parent,
		p:      pen.NewPaperSize(10 * 1024), // 10KB
		logger: fg,
		parent: parent,
	}
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
// Imports needed: reflect, unsafe.
func goSliceFromPtr(target, ptr, len string) string {
	return pen.NewPiece().
		Linef("sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&%s))", target).
		Linef("sliceHeader.Data = uintptr(unsafe.Pointer(%s))", ptr).
		Linef("sliceHeader.Len = %s", len).
		Linef("sliceHeader.Cap = %s", len).
		String()
}

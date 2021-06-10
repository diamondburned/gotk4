package girgen

import (
	"fmt"
	"log"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

// TODO:
//   - support ParameterIsOutput
//   - support CallerAllocates

// ValueProp describes the generic properties of a Go or C value for conversion.
type ValueProp struct {
	InName    string
	OutName   string
	AnyType   gir.AnyType
	Ownership gir.TransferOwnership

	// Closure marks the user_data argument. If this is provided, then the
	// conversion function will set the parameter to the callback ID. The caller
	// is responsible for skipping conversion of these indices.
	Closure *int

	// Destroy marks the callback to Destroy the user_data argument. If this is
	// provided, then callbackDelete will be set along with Closure.
	Destroy *int

	// ParameterIndex explicitly gives this value an index used for matching
	// with the given index clues from the GIR files, such as closure, destroy
	// or length.
	ParameterIndex *int

	// ParameterIsOutput makes the conversion treat this value like an output
	// parameter. Specifically, the output will be dereferenced when it is set.
	ParameterIsOutput bool

	// CallerAllocates determines if the converter should take care of
	// allocating the type or not.
	CallerAllocates bool

	// AllowNone, if true, will allow types that cannot be converted to stay.
	AllowNone bool
}

// NewValuePropParam creates a ValueProp from the given parameter attribute.
func NewValuePropParam(in, out string, i int, param gir.ParameterAttrs) ValueProp {
	value := ValueProp{
		InName:            in,
		OutName:           out,
		AnyType:           param.AnyType,
		Ownership:         param.TransferOwnership,
		Closure:           param.Closure,
		Destroy:           param.Destroy,
		AllowNone:         param.AllowNone,
		CallerAllocates:   param.CallerAllocates,
		ParameterIsOutput: param.Direction == "out",
	}
	if i > -1 {
		value.ParameterIndex = &i
	}
	return value
}

// NewValuePropReturn creates a new ValueProp from the given return attribute.
func NewValuePropReturn(in, out string, ret gir.ReturnValue) ValueProp {
	return ValueProp{
		InName:    in,
		OutName:   out,
		AnyType:   ret.AnyType,
		Ownership: ret.TransferOwnership,
		AllowNone: ret.Skip || ret.AllowNone,
		Closure:   ret.Closure,
		Destroy:   ret.Destroy,
	}
}

// NewValuePropField creates a new ValueProp from the given field. The struct is
// assumed to have a native field.
func NewValuePropField(recv, out string, field gir.Field) ValueProp {
	return ValueProp{
		InName:  fmt.Sprintf("%s.native.%s", recv, cgoField(field.Name)),
		OutName: out,
		AnyType: field.AnyType,
	}
}

// NewThrowValue creates a new GError value.
func NewThrowValue(in, out string) ValueProp {
	return ValueProp{
		InName:  in,
		OutName: out,
		AnyType: gir.AnyType{
			Type: &gir.Type{
				Name:  "GLib.Error",
				CType: "GError*",
			},
		},
		AllowNone:         true,
		ParameterIsOutput: true,
	}
}

// IsZero returns true if ValueProp is empty.
func (value *ValueProp) IsZero() bool {
	return value.InName == "" || value.OutName == ""
}

// outputAllocs returns true if the parameter is a value we need to allocate
// ourselves.
func (value *ValueProp) outputAllocs() bool {
	return value.ParameterIsOutput && (value.CallerAllocates || value.isTransferring())
}

func (value *ValueProp) loadIgnore(ignores map[int]struct{}) {
	// These are handled below.
	if value.Closure != nil {
		ignores[*value.Closure] = struct{}{}
	}
	if value.Destroy != nil {
		ignores[*value.Destroy] = struct{}{}
	}
	if value.AnyType.Array != nil && value.AnyType.Array.Length != nil {
		ignores[*value.AnyType.Array.Length] = struct{}{}
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
func (prop *ValueProp) isTransferring() bool {
	return false ||
		prop.Ownership.TransferOwnership == "full" ||
		prop.Ownership.TransferOwnership == "container"
}

type valueConverter interface {
	convert(*ValueConverted) bool
}

type conversionTo struct {
	converter valueConverter

	ng     *NamespaceGenerator
	logger lineLogger
	parent string

	results []ValueConverted
	ignores map[int]struct{}
}

func newConversionTo(
	fg *FileGenerator, parent string, values []ValueProp, converter valueConverter) conversionTo {

	conv := conversionTo{
		converter: converter,

		ng:     fg.parent,
		logger: fg,
		parent: parent,

		results: make([]ValueConverted, len(values)),
		ignores: map[int]struct{}{},
	}

	for _, value := range values {
		value.loadIgnore(conv.ignores)
	}

	for i := range conv.results {
		conv.results[i] = newValueConverted(&values[i])
		result := &conv.results[i]

		if !result.Skip && result.ParameterIndex != nil {
			_, ignore := conv.ignores[*result.ParameterIndex]
			if ignore {
				result.Skip = true
			}
		}
	}

	return conv
}

// AddCCallParam adds call parameters for C functions.
func AddCCallParam(params *pen.Joints, goc *TypeConversionToC, cgo *TypeConversionToGo) {
	for _, result := range goc.results {
		params.Add(result.OutCall)
	}
	for _, result := range cgo.results {
		if result.ParameterIsOutput {
			params.Add(result.InCall)
		}
	}
}

// ConvertAll converts all values.
func (conv *conversionTo) ConvertAll() []ValueConverted {
	results := make([]ValueConverted, 0, len(conv.results))

	// Convert everything in one go.
	for i := range conv.results {
		if !conv.convert(&conv.results[i]) {
			return nil
		}
	}

	for _, result := range conv.results {
		if result.Skip {
			continue
		}

		results = append(results, result)
	}

	return results
}

// Convert converts the value at the given index.
func (conv *conversionTo) Convert(i int) *ValueConverted {
	// Bound check.
	if i >= len(conv.results) {
		return nil
	}

	result := &conv.results[i]
	if !conv.convert(result) || result.Skip {
		return nil
	}

	return result
}

func (conv *conversionTo) convert(result *ValueConverted) bool {
	if result.p == nil {
		return true
	}

	if !conv.converter.convert(result) {
		conv.log(LogDebug, "C->Go cannot convert type", anyTypeC(result.AnyType))
		return false
	}

	if result.InType == "" || result.OutType == "" {
		log.Panicf(
			"missing CGoType or GoType for parent %s, %s -> %s",
			conv.parent, result.InName, result.OutName,
		)
	}

	// Only finalize when succeeded.
	result.finalize()
	return true
}

// convertInner is used while converting arrays; it returns the result of the
// inner value converted.
func (conv *conversionTo) convertInner(of *ValueConverted, in, out string) *ValueConverted {
	if of.AnyType.Array == nil {
		return nil
	}

	owner := of.Ownership

	// If the array's ownership is ONLY container, then we must not take over
	// the inner values. Therefore, we only generate the appropriate code.
	if owner.TransferOwnership == "container" {
		owner.TransferOwnership = "none"
	}

	result := newValueConverted(&ValueProp{
		InName:    of.InName,
		OutName:   of.OutName,
		AnyType:   of.AnyType.Array.AnyType,
		Ownership: owner,
	})

	if !conv.convert(&result) {
		return nil
	}

	return &result
}

// convertParam converts the parameter at the given index. This parameter index
// is different from indexing the values slice.
func (conv *conversionTo) convertParam(at int) *ValueConverted {
	for i := range conv.results {
		result := &conv.results[i]
		if result.ParameterIndex == nil || *result.ParameterIndex != at {
			continue
		}

		if !conv.convert(result) {
			return nil
		}

		return result
	}

	conv.log(LogError, "C->Go conversion arg not found at", at)
	return nil
}

func (conv *conversionTo) log(lvl LogLevel, v ...interface{}) {
	if conv.parent != "" {
		v2 := make([]interface{}, 0, 2+len(v))
		v2 = append(v2, "in", conv.parent)
		v2 = append(v2, v...)

		v = v2
	}

	conv.logger.Logln(lvl, v...)
}

// ValueConverted is the result of conversion for a single value.
//
// Quick convention note:
//
//    - {In,Out}Name is for the original name with no modifications.
//    - {In,Out}Call is used for the C or Go function arguments.
//    - {in,out}Set is used for setting the variables.
//
// Usually, these are the same, but they're sometimes different depending on the
// edge case.
type ValueConverted struct {
	*ValueProp // original

	InCall    string // use for calls
	InType    string
	InDeclare string

	OutCall    string // use for calls
	OutType    string
	OutDeclare string

	Conversion string
	ConversionSideEffects

	resolved       *ResolvedType // only for type conversions
	needsNamespace bool

	p       *pen.PaperString
	inDecl  *pen.PaperString
	outDecl *pen.PaperString

	// Skip is true if the type value should be skipped for
	Skip bool
}

func newValueConverted(value *ValueProp) ValueConverted {
	return ValueConverted{
		ValueProp: value,
		InCall:    value.InName,
		OutCall:   value.OutName,

		p:       pen.NewPaperStringSize(1024), // 1KB
		inDecl:  pen.NewPaperStringSize(128),  // 0.1KB
		outDecl: pen.NewPaperStringSize(128),  // 0.1KB
	}
}

func (value *ValueConverted) finalize() {
	value.InDeclare = value.inDecl.String()
	value.OutDeclare = value.outDecl.String()
	value.Conversion = value.p.String()

	// Allow GC to collect the internal buffers.
	value.inDecl = nil
	value.outDecl = nil
	value.p = nil
}

// resolveType resolves the value type to the resolved field. If inputC is true,
// then the input type is set to the CGo type, otherwise the Go type is set.
func (value *ValueConverted) resolveType(conv *conversionTo, inputC bool) bool {
	if value.AnyType.Type == nil {
		return false
	}

	if value.resolved != nil {
		// already resolved
		return true
	}

	value.resolved = conv.ng.ResolveType(*value.AnyType.Type)
	if value.resolved == nil {
		return false
	}

	// If this is the output parameter, then the pointer count should be less.
	// This only affects the Go type.
	if value.ParameterIsOutput {
		value.resolved.Ptr--
	}

	value.needsNamespace = value.resolved.NeedsNamespace(conv.ng.current)
	if value.needsNamespace {
		value.addImportAlias(value.resolved.Import, value.resolved.Package)
	}

	cgoType := value.resolved.CGoType()
	goType := value.resolved.PublicType(value.needsNamespace)

	if inputC {
		value.InType = cgoType
		value.OutType = goType
	} else {
		value.OutType = cgoType
		value.InType = goType
	}

	if inputC && value.outputAllocs() {
		value.InCall = "&" + value.InCall
		value.InType = strings.TrimPrefix(value.InType, "*")
	}

	value.outDecl.Linef("var %s %s", value.OutName, value.OutType)
	value.inDecl.Linef("var %s %s", value.InName, value.InType)

	return true
}

// cgoSetObject generates a glib.Take or glib.AssumeOwnership into a new
// function.
func (value *ValueConverted) cgoSetObject() {
	var gobjectFunction string
	if value.isTransferring() {
		// Full or container means we implicitly own the object, so we must
		// not take another reference.
		gobjectFunction = "AssumeOwnership"
	} else {
		// Else the object is either unowned by us or it's a floating
		// reference. Take our own or sink the object.
		gobjectFunction = "Take"
	}

	value.addGLibImport()
	value.addImport(importInternal("gextras"))
	value.addImport("unsafe")

	value.p.Linef(
		"%s = gextras.CastObject(externglib.%s(unsafe.Pointer(%s.Native()))).(%s)",
		value.OutName, gobjectFunction, value.InName, value.OutType,
	)
}

func (value *ValueConverted) malloc(lenOf string, add1 bool) string {
	lenOf = "len(" + lenOf + ")"
	if add1 {
		lenOf = "(" + lenOf + "+1)"
	}

	return fmt.Sprintf("C.malloc(%s * %s)", lenOf, value.csizeof())
}

func (value *ValueConverted) csizeof() string {
	// Arrays are lists of pointers.
	if strings.Contains(anyTypeC(value.AnyType), "*") {
		return value.ptrsz()
	}

	if value.resolved == nil {
		// Erroneous case.
		return value.ptrsz()
	}

	if value.resolved.IsRecord() {
		return "C.sizeof_struct_" + value.resolved.CType
	}

	return "C.sizeof_" + value.resolved.CType
}

func (value *ValueConverted) ptrsz() string {
	value.addImport("unsafe")
	// Size of a pointer is the same as int.
	return "unsafe.Sizeof(int(0))"
}

// ConversionSideEffects describes the side effects of the conversion, such as
// importing new things or modifying the Cgo preamble.
type ConversionSideEffects struct {
	Imports         map[string]string
	Callbacks       []string
	CallbackDelete  bool
	NeedsStdBool    bool
	NeedsGLibObject bool
}

// applySideEffects applies all side effects of the given list of type converted
// results.
func applySideEffects(fg *FileGenerator, results []ValueConverted) {
	for _, result := range results {
		result.Apply(fg)
	}
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
	if sides.NeedsGLibObject {
		fg.needsGLibObject()
	}
	for path, alias := range sides.Imports {
		fg.addImportAlias(path, alias)
	}
	for _, callback := range sides.Callbacks {
		fg.addCallbackHeader(callback)
	}
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

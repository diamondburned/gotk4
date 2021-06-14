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

	// Optional, if true, will allow types that cannot be converted to stay.
	Optional bool

	// Nullable, if true, will preserve the pointer type to preserve
	// nullability.
	Nullable bool // TODO

	// Scope determines the asynchronous rules for the value.
	//
	//    - notified: valid until a GDestroyNotify argument is called
	//    - async: only valid for the duration of the first callback invocation
	//      (can only be called once)
	//    - call: only valid for the duration of the call, can be called
	//      multiple times during the call
	//
	Scope string
}

// NewValuePropParam creates a ValueProp from the given parameter attribute.
func NewValuePropParam(in, out string, i int, param gir.ParameterAttrs) ValueProp {
	value := ValueProp{
		InName:            in,
		OutName:           out,
		AnyType:           param.AnyType,
		Ownership:         param.TransferOwnership,
		Scope:             param.Scope,
		Closure:           param.Closure,
		Destroy:           param.Destroy,
		Optional:          param.Optional || param.Skip,
		Nullable:          param.Nullable,
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
		Scope:     ret.Scope,
		Optional:  ret.Skip,
		Nullable:  ret.Nullable,
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
		Ownership: gir.TransferOwnership{
			TransferOwnership: "full",
		},
		AnyType: gir.AnyType{
			Type: &gir.Type{
				Name: "GLib.Error",
				// Function parameter type is technically a double-pointer here.
				CType: "GError**",
			},
		},
		Optional:          true,
		Nullable:          true,
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
		InName:            in,
		OutName:           out,
		AnyType:           of.AnyType.Array.AnyType,
		Ownership:         owner,
		ParameterIsOutput: of.ParameterIsOutput,
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
	SideEffects

	p       *pen.PaperString
	inDecl  *pen.PaperString
	outDecl *pen.PaperString

	resolved       *ResolvedType // only for type conversions
	needsNamespace bool

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

	// Pretend that ignored types don't exist.
	if conv.ng.mustIgnore(value.AnyType.Type.Name, value.AnyType.Type.CType) {
		return false
	}
	// ResolveType already checks this, but we can early bail.
	if !value.AnyType.Type.IsIntrospectable() {
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

	if value.resolved.IsCallback() {
		value.addCallback(value.resolved.Extern.Result.Callback)
	}

	// If this is the output parameter, then the pointer count should be less.
	// This only affects the Go type.
	if inputC && value.ParameterIsOutput && value.resolved.Ptr > 0 {
		value.resolved.Ptr--
	}

	value.needsNamespace = value.resolved.NeedsNamespace(conv.ng.current)
	if value.needsNamespace {
		// We're using the PublicType, so add that import.
		value.importPubl(value.resolved)
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

	if inputC && value.ParameterIsOutput {
		value.InCall = "&" + value.InCall
		value.InType = strings.TrimPrefix(value.InType, "*")
	}

	value.inDecl.Linef("var %s %s // in", value.InName, value.InType)
	value.outDecl.Linef("var %s %s // out", value.OutName, value.OutType)

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
	value.addImportInternal("gextras")
	value.addImport("unsafe")

	value.p.Linef(
		"%s = gextras.CastObject(externglib.%s(unsafe.Pointer(%s.Native()))).(%s)",
		value.OutName, gobjectFunction, value.InName, value.OutType,
	)
}

func (value *ValueConverted) cmalloc(lenOf string, add1 bool) string {
	lenOf = "len(" + lenOf + ")"
	if add1 {
		lenOf = "(" + lenOf + "+1)"
	}

	return fmt.Sprintf("C.malloc(C.ulong(%s) * C.ulong(%s))", lenOf, value.csizeof())
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

	// 	if value.resolved.IsRecord() {
	// 		return "C.sizeof_struct_" + value.resolved.CType
	// 	}

	return "C.sizeof_" + value.resolved.CType
}

func (value *ValueConverted) ptrsz() string {
	value.addImport("unsafe")
	// Size of a pointer is the same as uint.
	return "unsafe.Sizeof(uint(0))"
}

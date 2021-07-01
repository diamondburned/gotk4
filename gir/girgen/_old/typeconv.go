package girgen

import (
	"fmt"
	"log"
	"strings"

	"github.com/diamondburned/gotk4/core/pen"
	"github.com/diamondburned/gotk4/gir"
)

// ConversionDirection is the conversion direction between Go and C.
type ConversionDirection uint8

const (
	_ ConversionDirection = iota
	ConvertGoToC
	ConvertCToGo
)

// ConversionValueIndex describes an overloaded index type that reserves its
// negative values for special values.
type ConversionValueIndex int

const (
	_ ConversionValueIndex = -iota // 0
	UnknownValueIndex
	ErrorValueIndex
	ReturnValueIndex
)

// Index returns the actual underlying index if any, or it returns -1.
func (ix ConversionValueIndex) Index() int {
	if ix > UnknownValueIndex {
		return int(ix)
	}
	return -1
}

// Is checks that the index matches. This method should be used as it guarantees
// that the given index isn't special.
func (ix ConversionValueIndex) Is(at int) bool {
	if at < 0 {
		log.Panicln("given index", at, "is invalid")
	}
	return ix.Index() == at
}

// ConversionValue describes the generic properties of a Go or C value for
// conversion.
type ConversionValue struct {
	gir.ParameterAttrs

	InName  string
	OutName string

	// WrapObject, if not empty, will make the converter directly wrap to the
	// object type instead of wrapping it in a Box. This should only be used for
	// converting within types of the same package.
	WrapObject string

	// Direction is the direction of conversion.
	Direction ConversionDirection

	// ParameterIndex explicitly gives this value an index used for matching
	// with the given index clues from the GIR files, such as closure, destroy
	// or length.
	ParameterIndex ConversionValueIndex
}

// NewConversionValue creates a new ConversionValue from the given parameter
// attributes.
func NewConversionValue(
	in, out string, i int, dir ConversionDirection, param gir.ParameterAttrs) ConversionValue {

	value := ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      dir,
		ParameterIndex: UnknownValueIndex,
		ParameterAttrs: param,
	}
	if i > -1 {
		value.ParameterIndex = ConversionValueIndex(i)
	}

	return value
}

// NewConversionValueReturn creates a new ConversionValue from the given return
// attribute.
func NewConversionValueReturn(
	in, out string, dir ConversionDirection, ret gir.ReturnValue) ConversionValue {

	return ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      dir,
		ParameterIndex: ReturnValueIndex,
		ParameterAttrs: gir.ParameterAttrs{
			Closure:  ret.Closure,
			Destroy:  ret.Destroy,
			Scope:    ret.Scope,
			Skip:     ret.Skip,
			Nullable: ret.Nullable,
			Optional: ret.Skip,

			TransferOwnership: ret.TransferOwnership,
			AnyType:           ret.AnyType,
			Doc:               ret.Doc,
		},
	}
}

// NewConversionValueField creates a new ConversionValue from the given C struct
// field. The struct is assumed to have a native field.
func NewConversionValueField(recv, out string, field gir.Field) ConversionValue {
	return ConversionValue{
		InName:         fmt.Sprintf("%s.%s", recv, cgoField(field.Name)),
		OutName:        out,
		Direction:      ConvertCToGo,
		ParameterIndex: UnknownValueIndex,
		ParameterAttrs: gir.ParameterAttrs{
			Name:    field.Name,
			Skip:    field.Private || !field.Readable || field.Bits > 0,
			AnyType: field.AnyType,
			Doc:     field.Doc,
		},
	}
}

// NewThrowValue creates a new GError value. Thrown values are always assumed
// to be conversions from C to Go. Errors should ALWAYS go AFTER the return!
func NewThrowValue(in, out string) ConversionValue {
	return ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      ConvertCToGo,
		ParameterIndex: ErrorValueIndex,
		ParameterAttrs: gir.ParameterAttrs{
			TransferOwnership: gir.TransferOwnership{
				TransferOwnership: "full",
			},
			AnyType: gir.AnyType{
				Type: &gir.Type{
					Name: "GLib.Error",
					// Function parameter type is technically a double-pointer
					// here.
					CType: "GError**",
				},
			},
			Optional:  true,
			Nullable:  true,
			Direction: "out",
		},
	}
}

// IsZero returns true if ConversionValue is empty.
func (value *ConversionValue) IsZero() bool {
	return value.InName == "" || value.OutName == ""
}

// ParameterIsOutput returns true if the direction is out.
func (value *ConversionValue) ParameterIsOutput() bool {
	return value.ParameterAttrs.Direction == "out"
}

// outputAllocs returns true if the parameter is a value we need to allocate
// ourselves.
func (value *ConversionValue) outputAllocs() bool {
	return value.ParameterIsOutput() && (value.CallerAllocates || value.isTransferring())
}

// isTransferring is true when the ownership is either full or container. If the
// converter code isn't generating for an array, then distinguishing this
// doesn't matter. If the caller hasn't set the ownership yet, then it is
// assumed that we're not getting the ownership, therefore false is returned.
//
// If the generating code is an array, and the conversion is being passed into
// the same generation routine for the inner type, then the ownership should be
// turned into "none" just for that inner type. See TypeConversion.inner().
func (prop *ConversionValue) isTransferring() bool {
	return false ||
		prop.TransferOwnership.TransferOwnership == "full" ||
		prop.TransferOwnership.TransferOwnership == "container"
}

type TypeConverter struct {
	results []ValueConverted
	ng      *NamespaceGenerator
	logger  LineLogger
}

// NewTypeConverter creates a new type converter from the given file generator.
// The converter will add no side effects to the given file generator.
func NewTypeConverter(ng *NamespaceGenerator, values []ConversionValue) *TypeConverter {
	conv := TypeConverter{
		ng:      ng,
		results: make([]ValueConverted, len(values)),
	}

	// paramAt gets the parameter at the given index.
	paramAt := func(at int) *ConversionValue {
		for i, value := range values {
			if value.ParameterIndex.Is(at) {
				return &values[i]
			}
		}
		return nil
	}

	// skip marks the value at the given parameter index to be skipped.
	skip := func(at int) {
		if value := paramAt(at); value != nil {
			value.Skip = true
		}
	}

	// isSameDirection checks that the parameter at the given index has the same
	// direction. In some cases like g_main_context_query, the output parameter
	// type is handled weirdly with an opposite direction length input, and
	// there's no good way to handle that in Go, so we skip.
	isSameDirection := func(of *ConversionValue, at int) bool {
		value := paramAt(at)
		if value != nil {
			return value.ParameterAttrs.Direction == of.ParameterAttrs.Direction
		}
		return true
	}

	for _, value := range values {
		// Ensure the direction is valid.
		if value.Direction == 0 {
			conv.log(LogError, "value", value.InName, "->", value.OutName, "has invalid direction")
			return nil
		}

		if value.ParameterAttrs.Direction == "out" && !anyTypeIsPtr(value.AnyType) {
			// Output direction but not pointer parameter is invalid; bail.
			conv.log(LogUnusuality, "value type", anyTypeC(value.AnyType), "is output but no ptr")
			return nil
		}

		if value.Closure != nil {
			skip(*value.Closure)
		}
		if value.Destroy != nil {
			skip(*value.Destroy)
		}

		if value.AnyType.Array != nil && value.AnyType.Array.Length != nil {
			if !isSameDirection(&value, *value.AnyType.Array.Length) {
				return nil
			}

			skip(*value.AnyType.Array.Length)
		}
	}

	for i := range conv.results {
		// Fill up the results list after transforming the values.
		conv.results[i] = newValueConverted(&values[i])
	}

	return &conv
}

// AddCCallParam adds call parameters for C functions.
func AddCCallParam(converter *TypeConverter) []string {
	if converter == nil {
		return nil
	}

	// TODO: find a less awful hack, which is a very non-trivial hack, because
	// we're splitting input and output parameters onto its own slice and
	// convert them separately.
	//
	// A good way to solve this would be to combine them into the same routine,
	// probably by using a list of []TypeConversion interfaces and invoke
	// different routines depending on the type.

	params := make([]string, 0, len(converter.results))

	for _, result := range converter.results {
		switch result.Direction {
		case ConvertGoToC:
			params = append(params, result.OutCall)
		case ConvertCToGo:
			if result.ParameterIsOutput() {
				params = append(params, result.InCall)
			}
		}
	}

	return params
}

// UseLogger sets the logger to be used instead of the given NamespaceGenrator.
func (conv *TypeConverter) UseLogger(logger LineLogger) {
	if conv != nil {
		conv.logger = logger
	}
}

// ConvertAll converts all values.
func (conv *TypeConverter) ConvertAll() []ValueConverted {
	// Allow calling with a nil TypeConverter to allow the constructor to return
	// a nil, but make it convenient enough that the caller wouldn't have to
	// check.
	if conv == nil {
		return nil
	}

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
func (conv *TypeConverter) Convert(i int) *ValueConverted {
	if conv == nil {
		return nil
	}

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

func (conv *TypeConverter) convert(result *ValueConverted) bool {
	if result.p == nil {
		// result is already finalized, skip.
		return true
	}

	switch result.Direction {
	case ConvertCToGo:
		if !conv.cgoConvert(result) {
			conv.log(LogDebug, "C->Go cannot convert type", anyTypeC(result.AnyType))
			return false
		}
	case ConvertGoToC:
		if !conv.gocConvert(result) {
			conv.log(LogDebug, "Go->C cannot convert type", anyTypeC(result.AnyType))
			return false
		}
	default:
		return false
	}

	if result.InType == "" || result.OutType == "" {
		conv.log(LogError, "missing CGoType or GoType for", result.InName, "->", result.OutName)
		panic("see above")
	}

	// Only finalize when succeeded.
	result.finalize()
	return true
}

// convertInner is used while converting arrays; it returns the result of the
// inner value converted.
func (conv *TypeConverter) convertInner(of *ValueConverted, in, out string) *ValueConverted {
	if of.AnyType.Array == nil {
		return nil
	}

	attrs := of.ParameterAttrs
	attrs.AnyType = of.AnyType.Array.AnyType

	// If the array's ownership is ONLY container, then we must not take over
	// the inner values. Therefore, we only generate the appropriate code.
	if attrs.TransferOwnership.TransferOwnership == "container" {
		attrs.TransferOwnership.TransferOwnership = "none"
	}

	result := newValueConverted(&ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      of.Direction,
		ParameterIndex: UnknownValueIndex,
		ParameterAttrs: attrs,
	})

	if !conv.convert(&result) {
		return nil
	}

	return &result
}

// convertParam converts the parameter at the given index. This parameter index
// is different from indexing the values slice. If inherit is given (not nil),
// then several attributes such as the direction is brought over.
func (conv *TypeConverter) convertParam(at int) *ValueConverted {
	convert := func(result *ValueConverted) *ValueConverted {
		if !conv.convert(result) {
			return nil
		}
		return result
	}

	// Fast path.
	if at < len(conv.results) {
		result := &conv.results[at]
		if result.ParameterIndex.Is(at) {
			return convert(result)
		}
	}

	for i := range conv.results {
		result := &conv.results[i]
		if result.ParameterIndex.Is(at) {
			return convert(result)
		}
	}

	conv.log(LogUnusuality, "conversion arg not found at", at)
	return nil
}

const convertRefTmp = "refTmp"

// TODO: realistically, the difference between the expected poiner and what C
// wants is only 1. We can work around this.
//
// TODO: ideally, we should treat all foreign pointers as arrays, because they
// usually are. It would also allow the caller to allocate a sized array, as
// they could read the comments.
//
// TODO: there's a way to guess the pointer offset without switch-casing on
// every type. We can do this with IsPrimitive and IsClass fairly easily. We
// will have to account for Go type edge cases, however.

// convertRef generates weird code that references and dereferences values as
// needed, then call the conversion routine on it again. This shold only be used
// in type conversions, not array conversions.
//
// WARNING: the caller MUST call this AFTER value.resolveType!
func (conv *TypeConverter) convertRef(value *ValueConverted, wantC, wantGo int) bool {
	// ensure we're working with types only.
	if value.AnyType.Type == nil {
		return false
	}

	value.p.Descend()
	defer value.p.Ascend()

	// Make a copy of value so we can directly plug it back in.
	refValue := *value

	var (
		// positive needs dereferencing, negative needs referencing
		inDiff  int
		outDiff int
	)

	// Account for pointer types.
	if isGPointer(value.resolved.CType) {
		if wantC > 0 {
			wantC--
		} else {
			wantGo++
		}
	}

	// Prefer the implementation Go type instead of the public type.
	goType := value.resolved.ImplType(value.needsNamespace)

	// Change the pointer types into what the converter wants.
	switch value.Direction {
	case ConvertCToGo:
		refValue.InType, inDiff = forceGoPtr(refValue.InType, wantC)
		refValue.OutType, outDiff = forceGoPtr(goType, wantGo)
	case ConvertGoToC:
		refValue.InType, inDiff = forceGoPtr(goType, wantGo)
		refValue.OutType, outDiff = forceGoPtr(refValue.OutType, wantC)
	}

	refValue.InName = convertRefTmp + "In"
	refValue.OutName = convertRefTmp + "Out"

	refValue.p.Linef("var %s %s", refValue.InName, refValue.InType)
	refValue.p.Linef("var %s %s", refValue.OutName, refValue.OutType)
	refValue.p.EmptyLine()

	if inDiff > 0 {
		// Dereferencing can be a one-liner.
		refValue.p.Linef("%s = %s%s", refValue.InName, strings.Repeat("*", inDiff), value.InName)
	} else {
		// Use the original input value as name for the first variable to be
		// referenced.
		current := value.InName
		for i := 0; i < -inDiff; i++ {
			refValue.p.Linef("in%d := &%s", i, current)
			current = fmt.Sprintf("in%d", i)
		}
		// Copy the temporary variable into the refValue input.
		refValue.p.Linef("%s = %s", refValue.InName, current)
	}

	refValue.p.EmptyLine()

	var ok bool
	switch value.Direction {
	case ConvertCToGo:
		ok = conv.cgoTypeConverter(&refValue)
	case ConvertGoToC:
		ok = conv.gocTypeConverter(&refValue)
	}
	if !ok {
		return false
	}

	refValue.p.EmptyLine()

	// outDiff is the opposite for some reason, so negate it before using.
	if outDiff = -outDiff; outDiff > 0 {
		// Dereferencing can be a one-liner.
		refValue.p.Linef("%s = %s%s", value.OutName, strings.Repeat("*", outDiff), refValue.OutName)
	} else {
		// Use the new output value as name for the first variable to be
		// referenced.
		current := refValue.OutName
		for i := 0; i < -outDiff; i++ {
			refValue.p.Linef("out%d := &%s", i, current)
			current = fmt.Sprintf("out%d", i)
		}
		// Copy the temporary variable into the original output.
		refValue.p.Linef("%s = %s", value.OutName, current)
	}

	return true
}

func forceGoPtr(goType string, want int) (string, int) {
	current := strings.Count(goType, "*")
	return strings.Repeat("*", want) + strings.ReplaceAll(goType, "*", ""), current - want
}

func (conv *TypeConverter) log(lvl LogLevel, v ...interface{}) {
	if conv.logger == nil {
		conv.ng.Logln(lvl, v...)
	} else {
		conv.logger.Logln(lvl, v...)
	}
}

// ValueConverted is the result of conversion for a single value.
//
// Quick convention note:
//
//    - {In,Out}Name is for the original name with no modifications.
//    - {In,Out}Call is used for the C or Go function arguments.
//
// Usually, these are the same, but they're sometimes different depending on the
// edge case.
type ValueConverted struct {
	ConversionValue // original

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
}

func newValueConverted(value *ConversionValue) ValueConverted {
	return ValueConverted{
		ConversionValue: *value,
		InCall:          value.InName,
		OutCall:         value.OutName,

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
func (value *ValueConverted) resolveType(conv *TypeConverter) bool {
	if value.AnyType.Type == nil {
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

	// Copy Type for mutation.
	typ := *value.AnyType.Type

	// Proritize hard-coded types over ignored types.
	value.resolved = conv.ng.ResolveType(typ)
	if value.resolved == nil {
		conv.log(LogDebug, value.logPrefix(), "can't resolve", anyTypeCGo(value.AnyType), typ.Name)
		return false
	}

	// Apply type renamers only. Filtered types will be ignored in ResolveType.
	conv.ng.mustIgnore(&typ.Name, &typ.CType)

	// Set the type back for use. We're setting the AnyType struct, which is a
	// copy, so it's fine.
	value.AnyType.Type = &typ

	if value.resolved.IsCallback() {
		value.addCallback(value.resolved.Extern.Result.Callback)
	}

	value.needsNamespace = value.resolved.NeedsNamespace(conv.ng.current)

	// If this is the output parameter, then the pointer count should be less.
	// This only affects the Go type.
	if value.Direction == ConvertCToGo && value.ParameterIsOutput() && value.resolved.Ptr > 0 {
		value.resolved.Ptr--
	}

	cgoType := value.resolved.CGoType()
	goType := value.resolved.PublicType(value.needsNamespace)

	if value.Direction == ConvertCToGo {
		value.InType = cgoType
		value.OutType = goType
	} else {
		value.OutType = cgoType
		value.InType = goType
	}

	if value.Direction == ConvertCToGo && value.ParameterIsOutput() {
		value.InCall = "&" + value.InCall
		value.InType = strings.TrimPrefix(value.InType, "*")
	}

	value.inDecl.Linef("var %s %s // in", value.InName, value.InType)
	value.outDecl.Linef("var %s %s // out", value.OutName, value.OutType)

	return true
}

// cgoSetObject generates a glib.Take or glib.AssumeOwnership into a new
// function. This should only be used for C to Go conversion.
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

	value.needsExternGLib()
	value.addImport("unsafe")

	if value.WrapObject != "" {
		// This is only ever used for local constructors, so we don't need to
		// mess with namespaces.
		// TODO: maybe not make such a bad assumption.
		value.p.Linef(
			"%s = %s(externglib.%s(unsafe.Pointer(%s)))",
			value.OutName, value.WrapObject, gobjectFunction, value.InName,
		)
		return
	}

	value.addImportInternal("gextras")
	value.p.Linef(
		"%s = gextras.CastObject(externglib.%s(unsafe.Pointer(%s))).(%s)",
		value.OutName, gobjectFunction, value.InName, value.OutType,
	)
}

func (value *ValueConverted) cmalloc(lenOf string, add1 bool) string {
	lenOf = "len(" + lenOf + ")"
	if add1 {
		lenOf += "+1"
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

func (value *ValueConverted) logPrefix() string {
	switch value.Direction {
	case ConvertCToGo:
		return "C->Go"
	case ConvertGoToC:
		return "Go->C"
	default:
		return ""
	}
}

// isPtr checks pointer coherency for C types and Go types. It's mostly used to
// guarantee that conversion routines get what they expect.
func (value *ValueConverted) isPtr(wantC int) bool {
	// See this same piece of code in convertRef for more information.
	if isGPointer(value.resolved.CType) && wantC > 0 {
		wantC--
	}

	switch value.Direction {
	case ConvertCToGo:
		return strings.Count(value.InType, "*") == wantC
	case ConvertGoToC:
		return strings.Count(value.OutType, "*") == wantC
	default:
		return false
	}

	// Rationale for not verifying Go pointer offset is that the pointer offset
	// is already determined in the type resolver routine, so repeating that
	// information is redundant.
	//
	// Edit: this rationale does NOT work because the type resolver only has the
	// wanted Go pointer information up to the point of creating a new
	// ResolvedType, and there's no way we can get it back. This routine may not
	// need to verify the Go pointer, but the conversiron routine will.
}

// ref wraps the given name and references or dereferences the name to be the
// wanted pointer.
func ref(v string, is, want uint8) string {
	return (reffer{is, want}).fwd() + v
}

type reffer struct {
	is   uint8
	want uint8
}

func (r reffer) difference() uint8 {
	if r.is > r.want {
		return r.is - r.want
	}
	return r.want - r.is
}

// fwd creates the pointer prefix to invoke the call.
func (r reffer) fwd() string {
	switch {
	case r.is > r.want:
		return strings.Repeat("*", int(r.is-r.want))
	case r.is < r.want:
		return strings.Repeat("&", int(r.want-r.is))
	default:
		return ""
	}
}

// bwd creates the pointer prefix to undo the pointer by referencing.
func (r reffer) bwd() string {
	// TODO: this is dumb; it doesn't work if the pointer difference is more
	// than 1.
	switch {
	case r.is > r.want:
		return strings.Repeat("&", int(r.is-r.want))
	case r.is < r.want:
		return strings.Repeat("*", int(r.want-r.is))
	default:
		return ""
	}
}

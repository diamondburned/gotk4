package typeconv

import (
	"fmt"
	"log"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// ConversionDirection is the conversion direction between Go and C.
type ConversionDirection uint8

const (
	_ ConversionDirection = iota
	ConvertGoToC
	ConvertCToGo
	ConvertGoToCToGo
)

// ConversionValueIndex describes an overloaded index type that reserves its
// negative values for special values.
type ConversionValueIndex int8

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

	// Direction is the direction of conversion.
	Direction ConversionDirection

	// ParameterIndex explicitly gives this value an index used for matching
	// with the given index clues from the GIR files, such as closure, destroy
	// or length.
	ParameterIndex ConversionValueIndex

	// KeepType overrides the abstract type if true.
	KeepType bool
}

// NewValue creates a new ConversionValue from the given parameter attributes.
func NewValue(
	in, out string, i int, dir ConversionDirection, param gir.Parameter) ConversionValue {

	// https://wiki.gnome.org/Projects/GObjectIntrospection/Annotations
	if param.TransferOwnership.TransferOwnership == "" {
		switch param.Direction {
		case "in":
			param.TransferOwnership.TransferOwnership = "full"
		case "out", "inout":
			if param.CallerAllocates {
				param.TransferOwnership.TransferOwnership = "none"
			}
		}
	}

	value := ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      dir,
		ParameterIndex: UnknownValueIndex,
		ParameterAttrs: param.ParameterAttrs,
	}
	if i > -1 {
		value.ParameterIndex = ConversionValueIndex(i)
	}

	return value
}

// NewReceiverValue creates a new ConversionValue specifically for the method
// receiver.
func NewReceiverValue(
	in, out string, dir ConversionDirection, param *gir.InstanceParameter) ConversionValue {

	return ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      dir,
		ParameterIndex: UnknownValueIndex,
		ParameterAttrs: param.ParameterAttrs,
		KeepType:       true, // concrete method receivers
	}
}

// NewReturnValue creates a new ConversionValue from the given return attribute.
func NewReturnValue(in, out string, dir ConversionDirection, ret gir.ReturnValue) ConversionValue {
	if ret.TransferOwnership.TransferOwnership == "" {
		if strings.Contains(types.AnyTypeC(ret.AnyType), "const") {
			ret.TransferOwnership.TransferOwnership = "none"
		} else {
			ret.TransferOwnership.TransferOwnership = "full"
		}
	}

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

// NewFieldValue creates a new ConversionValue from the given C struct field.
// The struct is assumed to have a native field.
func NewFieldValue(recv, out string, field gir.Field) ConversionValue {
	return ConversionValue{
		InName:         fmt.Sprintf("%s.native.%s", recv, strcases.CGoField(field.Name)),
		OutName:        out,
		Direction:      ConvertCToGo,
		ParameterIndex: UnknownValueIndex,
		ParameterAttrs: gir.ParameterAttrs{
			Name:    field.Name,
			Skip:    field.Private || !field.IsReadable() || field.Bits > 0,
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
	switch value.ParameterAttrs.Direction {
	case "out", "inout":
		return true
	case "in":
		return false
	}

	// GIR is fucking miserable. Why can't they get these properly?
	if value.Doc != nil && strings.HasPrefix(value.Doc.String, "Return location") {
		return true
	}

	return false
}

// outputAllocs returns true if the parameter is a value we need to allocate
// ourselves.
func (value *ConversionValue) outputAllocs() bool {
	return value.ParameterIsOutput() && (value.CallerAllocates || value.ownershipIsTransferring())
}

// isTransferring is true when the ownership is either full or container. If the
// converter code isn't generating for an array, then distinguishing this
// doesn't matter. If the caller hasn't set the ownership yet, then it is
// assumed that we're not getting the ownership, therefore false is returned.
//
// If the generating code is an array, and the conversion is being passed into
// the same generation routine for the inner type, then the ownership should be
// turned into "none" just for that inner type. See TypeConversion.inner().
func (prop *ConversionValue) ownershipIsTransferring() bool {
	return false ||
		prop.TransferOwnership.TransferOwnership == "full" ||
		prop.TransferOwnership.TransferOwnership == "container"
}

// ShouldFree returns true if the C value must be freed once we're done.
func (prop *ConversionValue) ShouldFree() bool {
	// goReceiving is true when we're receiving the C value.
	goReceiving := prop.ParameterIndex == ReturnValueIndex || prop.ParameterIsOutput()
	// If we're not converting C to Go, then we're probably in a callback, so
	// the ownership is flipped.
	if prop.Direction != ConvertCToGo {
		goReceiving = !goReceiving
	}

	if goReceiving {
		return prop.ownershipIsTransferring()
	}

	return !prop.ownershipIsTransferring()
}

// MustRealloc returns true if we need to malloc our values to give it to C.
// Generally, if a conversion routine has a no-alloc path, it should check
// MustRealloc first. If MustRealloc is true, then it must check ShouldFree.
//
//    if prop.MustAlloc() {
//        v = &oldValue
//    } else {
//        v = malloc()
//        if prop.ShouldFree() {
//            defer free(v)
//        }
//    }
//
func (prop *ConversionValue) MustRealloc() bool {
	// goGiving is true when we're giving the C value.
	goGiving := prop.ParameterIndex > -1 && !prop.ParameterIsOutput()
	// If we're not converting Go to C, then we're probably in a callback, so
	// the ownership is flipped.
	if prop.Direction != ConvertGoToC {
		goGiving = !goGiving
	}

	if goGiving {
		return prop.ownershipIsTransferring()
	}

	return !prop.ownershipIsTransferring()
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

	In  ValueName
	Out ValueName

	Conversion     string
	PostConversion string // only for two-stage conversions

	// internal states
	Resolved       *types.Resolved // only for type conversions
	IsPublic       bool
	NeedsNamespace bool

	log func(lvl logger.Level, v ...interface{})

	// output writers
	p       *pen.PaperString
	inDecl  *pen.PaperString
	outDecl *pen.PaperString
	header  file.Header

	fail  bool
	final bool
}

// ValueName contains the different names to use during and after conversion.
type ValueName struct {
	Name    string
	Call    string
	Set     string
	Type    string
	Declare string
}

func newValueConverted(conv *Converter, value *ConversionValue) ValueConverted {
	switch value.Direction {
	case ConvertCToGo, ConvertGoToC, ConvertGoToCToGo:
		// ok
	default:
		log.Panicf("unknown conversion direction %v", value.Direction)
	}

	return ValueConverted{
		ConversionValue: *value,
		In: ValueName{
			Name: value.InName,
			Call: value.InName,
			Set:  value.InName,
		},
		Out: ValueName{
			Name: value.OutName,
			Call: value.OutName,
			Set:  value.OutName,
		},
		log:     conv.Logln,
		p:       pen.NewPaperStringSize(1024), // 1KB
		inDecl:  pen.NewPaperStringSize(128),  // 0.1KB
		outDecl: pen.NewPaperStringSize(128),  // 0.1KB
	}
}

// Header returns the header of the current value.
func (value *ValueConverted) Header() *file.Header { return &value.header }

func (value *ValueConverted) isDone() bool { return value.final }

// finalize finalizes the value and returns true, or false is returned if the
// value is already finalized.
func (value *ValueConverted) finalize() bool {
	if value.final {
		return false
	}

	value.flush()

	// Allow GC to collect the internal buffers.
	value.inDecl = nil
	value.outDecl = nil
	value.p = nil

	value.final = true
	return true
}

// flush commits the writers.
func (value *ValueConverted) flush() {
	value.In.Declare = value.inDecl.String()
	value.Out.Declare = value.outDecl.String()
	value.Conversion = value.p.String()
}

func (value *ValueConverted) Logln(lvl logger.Level, v ...interface{}) {
	value.log(lvl, logger.Prefix(v, value.logPrefix())...)
}

// resolveType resolves the value type to the resolved field. If inputC is true,
// then the input type is set to the CGo type, otherwise the Go type is set.
func (value *ValueConverted) resolveType(conv *Converter) bool {
	if value.AnyType.Type == nil {
		return false
	}

	// ResolveType already checks this, but we can early bail.
	if !value.AnyType.Type.IsIntrospectable() {
		return false
	}

	if value.Resolved != nil {
		// already resolved
		return true
	}

	// Proritize hard-coded types over ignored types.
	resolveNamespace := types.OverrideNamespace(conv.fgen, conv.sourceNamespace)
	value.Resolved = types.Resolve(resolveNamespace, *value.AnyType.Type)
	if value.Resolved == nil {
		conv.Logln(logger.Debug, "can't resolve", types.AnyTypeCGo(value.AnyType))
		return false
	}

	if value.Resolved.IsCallback() {
		value.header.AddCallback(value.Resolved.Extern.Type.(*gir.Callback))
	}

	// If this is the output parameter, then the pointer count should be less.
	// This only affects the Go type.
	if value.ParameterIsOutput() && value.Resolved.Ptr > 0 {
		value.Resolved.Ptr--
	}

	value.NeedsNamespace = value.Resolved.NeedsNamespace(conv.currentNamespace)

	cgoType := value.Resolved.CGoType()

	switch value.Direction {
	case ConvertCToGo:
		value.In.Type = cgoType
		// Go output can be the implementation type.
		value.Out.Type = value.Resolved.ImplType(value.NeedsNamespace)

	case ConvertGoToC, ConvertGoToCToGo:
		value.Out.Type = cgoType
		if !value.KeepType && value.Resolved.IsAbstract() {
			value.In.Type = value.Resolved.PublicType(value.NeedsNamespace)
			value.IsPublic = true
		} else {
			value.In.Type = value.Resolved.ImplType(value.NeedsNamespace)
		}
	}

	if value.NeedsNamespace {
		if value.IsPublic {
			value.header.ImportPubl(value.Resolved)
		} else {
			value.header.ImportImpl(value.Resolved)
		}
	}

	if value.ParameterIsOutput() {
		switch value.Direction {
		case ConvertCToGo:
			value.In.Call = "&" + value.In.Call
			value.In.Type = strings.TrimPrefix(value.In.Type, "*")
		case ConvertGoToC:
			value.Out.Set = "*" + value.Out.Set
			value.Out.Type = strings.TrimPrefix(value.Out.Type, "*")
		case ConvertGoToCToGo: // in is Go, out is C
			value.In.Call = "&" + value.In.Call
			value.In.Type = strings.TrimPrefix(value.In.Type, "*")
		}
	}

	value.inDecl.Linef("var %s %s // in", value.InName, value.In.Type)
	value.outDecl.Linef("var %s %s // out", value.OutName, value.Out.Type)

	return true
}

// cgoSetObject generates a glib.Take or glib.AssumeOwnership into a new
// function. This should only be used for C to Go conversion.
func (value *ValueConverted) cgoSetObject(conv *Converter) bool {
	var gobjectFunction string
	if value.ownershipIsTransferring() {
		// Full or container means we implicitly own the object, so we must
		// not take another reference.
		gobjectFunction = "AssumeOwnership"
	} else {
		// Else the object is either unowned by us or it's a floating
		// reference. Take our own or sink the object.
		gobjectFunction = "Take"
	}

	value.header.NeedsExternGLib()
	value.header.Import("unsafe")

	m := gotmpl.M{
		"Value": value,
		"Func":  gobjectFunction,
	}

	if value.Resolved.IsExternGLib("Object") {
		// Shortcut for GObject.
		value.p.LineTmpl(m, `
			<.Value.Out.Set> = externglib.<.Func>(unsafe.Pointer(<.Value.InPtr 1><.Value.InName>))
		`)
		return true
	}

	if value.IsPublic {
		// Require the abstract cast if we have an abstract type.
		goto abstract
	}

	if !value.NeedsNamespace {
		value.p.LineTmpl(m, `
			<.Value.Out.Set> = <.Value.OutPtr 1><.Value.Resolved.WrapName false ->
				(externglib.<.Func>(unsafe.Pointer(<.Value.InPtr 1><.Value.InName>)))
		`)
		return true
	}

	if tree := types.NewTree(conv.fgen); tree.ResolveFromType(value.Resolved) {
		wrap := tree.WrapInNamespace("obj", &value.header, conv.sourceNamespace)
		if value.OutPtr(1) == "*" {
			// Dereference the wrapped struct value by removing the &.
			wrap = strings.TrimPrefix(wrap, "&")
		}
		m["Wrap"] = wrap

		value.p.LineTmpl(m, `{
			obj := externglib.<.Func>(unsafe.Pointer(<.Value.InPtr 1><.Value.InName>))
			<.Value.Out.Set> = <.Wrap>
		}`)

		return true
	}

abstract:
	value.header.ImportCore("gextras")
	value.p.LineTmpl(m,
		"<.Value.Out.Set> = (< .Value.OutPtr 1 ->gextras.CastObject(externglib.<.Func>("+
			"unsafe.Pointer(<.Value.InPtr 1><.Value.InName>)))).(<.Value.Out.Type>)")

	return true
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
	if strings.Contains(types.AnyTypeC(value.AnyType), "*") {
		return value.ptrsz()
	}

	if value.Resolved == nil {
		// Erroneous case.
		return value.ptrsz()
	}

	return "C.sizeof_" + types.CleanCType(value.Resolved.CType, true)
}

func (value *ValueConverted) ptrsz() string {
	value.header.Import("unsafe")
	// Size of a pointer is the same as uint.
	return "unsafe.Sizeof(uint(0))"
}

func (value *ValueConverted) logPrefix() string {
	var prefix string
	switch value.Direction {
	case ConvertCToGo:
		prefix = fmt.Sprintf("C %s -> Go %s", value.InName, value.OutName)
	case ConvertGoToC:
		prefix = fmt.Sprintf("Go %s -> C %s", value.InName, value.OutName)
	case ConvertGoToCToGo:
		prefix = fmt.Sprintf("Go %s -> C %s -> Go", value.InName, value.OutName)
	}

	if value.Resolved != nil {
		prefix += fmt.Sprintf(" (%s)", value.Resolved.PublicType(true))
	}

	return prefix
}

// isPtr checks pointer coherency for C types and Go types. It's mostly used to
// guarantee that conversion routines get what they expect.
func (value *ValueConverted) isPtr(wantC int) bool {
	// See this same piece of code in convertRef for more information.
	if value.Resolved.IsGpointer() && wantC > 0 {
		wantC--
	}

	switch value.Direction {
	case ConvertCToGo:
		return strings.Count(value.In.Type, "*") == wantC
	case ConvertGoToC, ConvertGoToCToGo: // TODO: doubt
		return strings.Count(value.Out.Type, "*") == wantC
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

// InNamePtrPubl adds in an edge case if the value being inputted is possibly a
// Go interface.
func (value *ValueConverted) InNamePtrPubl(want int) string {
	if want > 0 && value.IsPublic {
		want--
	}
	return value.InNamePtr(want)
}

// InNamePtr returns the name with the pointer prefixed using InPtr.
func (value *ValueConverted) InNamePtr(want int) string {
	ptr := value.InPtr(want)
	if ptr == "" {
		return value.InName
	}
	return fmt.Sprintf("(%s%s)", ptr, value.InName)
}

func (value *ValueConverted) InPtr(want int) string {
	// Account for gpointer.
	has := strings.Count(value.In.Type, "*")
	if value.Direction == ConvertCToGo && value.Resolved.IsGpointer() {
		has++
	}

	return value._ptr(has, want)
}

// OutCast returns the left-hand side consisting of the pointer dereference and
// the pointer-prefixed type.
func (value *ValueConverted) OutCast(want int) string {
	ptr := value.OutPtr(want)
	if ptr == "" && !strings.Contains(value.Out.Type, "*") {
		return value.Out.Type
	}
	return fmt.Sprintf("%s(%s%s)", ptr, ptr, value.Out.Type)
}

// OutInPtr returns the left-hand side for the output name and type SPECIFICALLY
// for inputting the output name elsewhere, like giving it to SetFinalizer.
func (value *ValueConverted) OutInPtr(want int) string {
	has := strings.Count(value.Out.Type, "*")
	return value._ptr(has, want)
}

func (value *ValueConverted) OutPtr(want int) string {
	has := strings.Count(value.Out.Type, "*")
	// Account for gpointer.
	if value.Resolved.IsGpointer() {
		switch value.Direction {
		case ConvertGoToC, ConvertGoToCToGo:
			has++
		}
	}

	ptr := value._ptr(want, has)
	if ptr == "&" {
		// Refuse to reference the value we converted, since that requires a
		// temporary variable.
		value.fail = true
		value.Logln(logger.Debug, "OutPtr refusing to reference, has", has, "want", want)
		return ""
	}

	return ptr
}

func (value *ValueConverted) _ptr(has, want int) string {
	if difference(has, want) > 1 {
		value.fail = true
		value.Logln(logger.Debug, "pointer difference too high, has", has, "want", want)
		return ""
	}

	switch {
	case has < want:
		return "&"
	case has > want:
		return "*"
	default:
		return ""
	}
}

func difference(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}

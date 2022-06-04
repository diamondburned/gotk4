// Package typeconv provides conversions between C and Go types.
package typeconv

import (
	"log"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// // ValueProcessor is a processor that can override each conversion value.
// type ValueProcessor interface {
// 	Process(value *ValueConverted)
// }

type Converter struct {
	Parent   *gir.TypeFindResult
	Results  []ValueConverted
	Callback bool
	// MustCast, if true, will force the converter to generate code that
	// force-casts InName and OutName to the right type.
	MustCast bool

	fgen   types.FileGenerator
	logger logger.LineLogger
	header file.Header

	final []ValueConverted
}

// NewConverter creates a new type converter from the given file generator.
// The converter will add no side effects to the given file generator.
func NewConverter(
	fgen types.FileGenerator, parent *gir.TypeFindResult, values []ConversionValue) *Converter {

	if fgen == nil {
		panic("missing fgen")
	}
	if parent == nil || parent.NamespaceFindResult == nil || parent.Type == nil {
		panic("missing parent")
	}

	conv := Converter{
		Parent:  parent,
		Results: make([]ValueConverted, len(values)),
		fgen:    fgen,
	}

	for i := range conv.Results {
		// Fill up the results list after transforming the values.
		conv.Results[i] = newValueConverted(&conv, &values[i])
	}

	// skip marks the value at the given parameter index to be skipped.
	skip := func(at int) {
		if value := conv.param(at); value != nil {
			value.Skip = true
		}
	}

	for _, value := range conv.Results {
		// Ensure the direction is valid.
		if value.Direction == 0 {
			value.Logln(logger.Error, "invalid direction")
			return nil
		}

		if value.ParameterAttrs.Direction == "out" && !types.AnyTypeIsPtr(value.AnyType) {
			// Output direction but not pointer parameter is invalid; bail.
			value.Logln(logger.Error, "is output but no ptr")
			return nil
		}

		// Only skip the parameter's closure index if the parameter itself is
		// a callback. Sometimes, the user_data parameter will flag the callback
		// as a closure argument, which messes up the generator.
		if value.Closure != nil && !types.IsGpointer(types.AnyTypeC(value.AnyType)) {
			skip(*value.Closure)
		}
		if value.Destroy != nil {
			skip(*value.Destroy)
		}

		if value.AnyType.Array != nil && value.AnyType.Array.Length != nil {
			skip(*value.Array.Length)
		}
	}

	return &conv
}

// CCallParams generates the call parameters for calling the C function.
func (conv *Converter) CCallParams() []string {
	if conv == nil {
		return nil
	}

	params := make([]string, 0, len(conv.Results))

	for _, result := range conv.Results {
		switch result.Direction {
		case ConvertGoToC:
			params = append(params, result.Out.Call)
		case ConvertCToGo:
			if result.ParameterIsOutput() {
				params = append(params, result.In.Call)
			}
		}
	}

	return params
}

// UseLogger sets the logger to be used instead of the given NamespaceGenrator.
func (conv *Converter) UseLogger(logger logger.LineLogger) {
	if conv != nil {
		conv.logger = logger
	}
}

// Header returns the header of all converted values. This method should only be
// used once ConvertAll or Convert has been called.
func (conv *Converter) Header() *file.Header {
	return &conv.header
}

// ConvertAll converts all values.
func (conv *Converter) ConvertAll() []ValueConverted {
	// Allow calling with a nil Converter to allow the constructor to return
	// a nil, but make it convenient enough that the caller wouldn't have to
	// check.
	if conv == nil {
		return nil
	}

	if conv.final != nil {
		return conv.final
	}

	// Convert everything in one go.
	for i := range conv.Results {
		result := &conv.Results[i]

		if !conv.convert(result) || result.fail {
			// final is true if the value is already manually handled.
			// Otherwise, exit.
			if !conv.Results[i].final {
				result.Logln(logger.Debug, "no conversion")
				return nil
			}
		}

		if !result.Skip {
			// Prevent duplicated conversions when the parameter isn't skipped
			// (thus written automatically).
			result.finalize()
		}
	}

	if proc, ok := conv.fgen.(ConversionProcessor); ok {
		proc.ProcessConverter(conv)
	}

	conv.final = make([]ValueConverted, 0, len(conv.Results))

	for i := range conv.Results {
		// Finalize all results.
		result := &conv.Results[i]
		result.finalize()

		if result.Skip {
			continue
		}

		// if types.TypeIsInFile(conv.Parent.Type, "gsocketservice.") {
		// 	for path := range result.header.Imports {
		// 		if strings.Contains(path, "gextras") {
		// 			result.Logln(logger.Debug, "imports gextras")
		// 		}
		// 	}
		// }

		file.ApplyHeader(conv, &conv.Results[i])
		conv.final = append(conv.final, *result)
	}

	return conv.final
}

// Convert converts the value at the given index.
func (conv *Converter) Convert(i int) *ValueConverted {
	if conv == nil {
		return nil
	}

	// Bound check.
	if i >= len(conv.Results) {
		return nil
	}

	// Ensure that all values are converted.
	if conv.ConvertAll() == nil {
		return nil
	}

	return &conv.Results[i]
}

func (conv *Converter) convert(result *ValueConverted) bool {
	if result.isDone() {
		// result is already finalized, skip.
		return !result.fail
	}

	if !result.resolveType(conv) {
		result.Logln(logger.Debug, "cannot resolve type")
		return false
	}

	switch result.Direction {
	case ConvertCToGo:
		if !conv.cgoConvert(result) || result.fail {
			result.Logln(logger.Debug, "C->Go cannot convert type", types.AnyTypeC(result.AnyType))
			result.fail = true
			return false
		}
	case ConvertGoToC:
		if !conv.gocConvert(result) || result.fail {
			result.Logln(logger.Debug, "Go->C cannot convert type", types.AnyTypeC(result.AnyType))
			result.fail = true
			return false
		}
	default:
		log.Panicf("unknown conversion direction %d", result.Direction)
		return false
	}

	if !result.Optional {
		if result.In.Type == "" || result.Out.Type == "" {
			result.Logln(logger.Error, "missing CGoType or GoType")
			panic("see above")
		}
		if result.Resolved == nil {
			result.Logln(logger.Error, "missing Resolved type")
			panic("see above")
		}
	}

	result.flush()
	return true
}

// convertInner is used while converting arrays; it returns the result of the
// inner value converted.
func (conv *Converter) convertInner(of *ValueConverted, in, out string) *ValueConverted {
	var inner *gir.Type

	switch {
	case of.Array != nil:
		inner = of.Array.Type
	case len(of.Type.Types) > 0:
		inner = &of.Type.Types[0]
	}

	if inner == nil {
		return nil
	}

	var existing *ValueType
	if len(of.Inner) > 0 {
		existing = &of.Inner[0]
	}

	value := conv.convertTypeExisting(of, in, out, inner, existing)
	if value == nil {
		switch {
		case of.Array != nil:
			of.Logln(logger.Debug, "convertInner fail on array", of.Array.Type.Name)
		case len(of.Type.Types) > 0:
			of.Logln(logger.Debug, "convertInner fail on inner types", of.Type.Types[0].Name)
		}
	}

	return value
}

// convertType converts a manually-crafted value with the given type.
func (conv *Converter) convertType(
	of *ValueConverted, in, out string, typ *gir.Type) *ValueConverted {

	return conv.convertTypeExisting(of, in, out, typ, nil)
}

func (conv *Converter) convertTypeExisting(
	of *ValueConverted, in, out string, typ *gir.Type, existing *ValueType) *ValueConverted {
	// If the array's ownership is ONLY container, then we must not take over
	// the inner values. Therefore, we only generate the appropriate code.
	owner := of.TransferOwnership.TransferOwnership
	if owner == "container" {
		owner = "none"
	}

	attrs := of.ParameterAttrs
	attrs.Nullable = false
	attrs.Optional = false
	attrs.AnyType = gir.AnyType{Type: typ}
	attrs.TransferOwnership.TransferOwnership = owner

	result := newValueConverted(conv, &ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      of.Direction,
		ParameterIndex: of.ParameterIndex,
		ParameterAttrs: attrs,
		InContainer:    of.Type != nil && len(of.Type.Types) > 0, // is container type
	})

	if of.Array != nil || result.InContainer {
		// inArray is used for deciding whether or not record should be a
		// pointer, and it shouldn't be inside an array.
		result.inArray = true
	}

	if existing != nil {
		result.ValueType = *existing
	}

	// If the value is in a container, then its direction is always in. This is
	// because container-inner types are converted in a callback.
	if result.InContainer {
		result.ParameterAttrs.Direction = "in"
	}

	if !conv.convert(&result) {
		return nil
	}

	result.Resolved.ImportImpl(conv.fgen, &of.header)
	return &result
}

func (conv *Converter) isRuntimeLinking() bool {
	return conv.fgen.LinkMode() == types.RuntimeLinkMode
}

// param returns the unconverted value.
func (conv *Converter) param(at int) *ValueConverted {
	for i := range conv.Results {
		result := &conv.Results[i]

		if result.ParameterIndex.Is(at) {
			return result
		}
	}
	return nil
}

// convertParam converts the parameter at the given index. This parameter index
// is different from indexing the values slice. If inherit is given (not nil),
// then several attributes such as the direction is brought over.
func (conv *Converter) convertParam(at int) *ValueConverted {
	convert := func(result *ValueConverted) *ValueConverted {
		if !conv.convert(result) {
			return nil
		}
		return result
	}

	// Fast path.
	if at < len(conv.Results) {
		for i := at; i < at+2 && i < len(conv.Results); i++ {
			result := &conv.Results[i]
			if result.ParameterIndex.Is(at) {
				return convert(result)
			}
		}
	}

	for i := range conv.Results {
		result := &conv.Results[i]
		if result.ParameterIndex.Is(at) {
			return convert(result)
		}
	}

	conv.Logln(logger.Debug, "conversion arg not found at", at)
	return nil
}

func (conv *Converter) convertIx(ix ConversionValueIndex) *ValueConverted {
	for i, res := range conv.Results {
		if res.ParameterIndex != ix {
			continue
		}

		if !conv.convert(&conv.Results[i]) {
			return nil
		}

		return &conv.Results[i]
	}

	return nil
}

// rgen returns the generator with the source namespace. It ensures that the
// correct namespace is searched from.
func (conv *Converter) rgen() types.FileGenerator {
	return types.OverrideNamespace(conv.fgen, conv.Parent.NamespaceFindResult)
}

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

func (conv *Converter) Logln(lvl logger.Level, v ...interface{}) {
	if conv.logger == nil {
		conv.fgen.Logln(lvl, v...)
	} else {
		conv.logger.Logln(lvl, v...)
	}
}

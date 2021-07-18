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
	Parent  *gir.TypeFindResult
	Results []ValueConverted

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

	// isSameDirection checks that the parameter at the given index has the same
	// direction. In some cases like g_main_context_query, the output parameter
	// type is handled weirdly with an opposite direction length input, and
	// there's no good way to handle that in Go, so we skip.
	isSameDirection := func(of *ConversionValue, at int) bool {
		if value := conv.param(at); value != nil {
			return value.ParameterAttrs.Direction == of.ParameterAttrs.Direction
		}
		return true
	}

	for _, value := range values {
		// Ensure the direction is valid.
		if value.Direction == 0 {
			conv.Logln(logger.Error,
				"value", value.InName, "->", value.OutName, "has invalid direction")
			return nil
		}

		if value.ParameterAttrs.Direction == "out" && !types.AnyTypeIsPtr(value.AnyType) {
			// Output direction but not pointer parameter is invalid; bail.
			conv.Logln(logger.Error,
				"value type", types.AnyTypeC(value.AnyType), "is output but no ptr")
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
			if !isSameDirection(&value, *value.AnyType.Array.Length) {
				return nil
			}

			skip(*value.AnyType.Array.Length)
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
		if !conv.convert(&conv.Results[i]) || conv.Results[i].fail {
			// final is true if the value is already manually handled.
			// Otherwise, exit.
			if !conv.Results[i].final {
				return nil
			}
		}
	}

	if proc, ok := conv.fgen.(ConversionProcessor); ok {
		proc.ProcessConverter(conv)
	}

	conv.final = make([]ValueConverted, 0, len(conv.Results))

	for i, result := range conv.Results {
		// Finalize all results.
		result.finalize()

		if result.Skip {
			continue
		}

		file.ApplyHeader(conv, &conv.Results[i])
		conv.final = append(conv.final, result)
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
	var inner *gir.AnyType

	switch {
	case of.AnyType.Array != nil:
		inner = &of.AnyType.Array.AnyType
	case of.AnyType.Type.Type != nil:
		inner = &of.AnyType.Type.AnyType
	}

	if inner == nil || inner.Type == nil {
		return nil
	}

	// If the array's ownership is ONLY container, then we must not take over
	// the inner values. Therefore, we only generate the appropriate code.
	owner := of.TransferOwnership.TransferOwnership
	if owner == "container" {
		owner = "none"
	}

	return conv.convertType(of, in, out, *inner, owner)
}

// convertType converts a manually-crafted value with the given type.
func (conv *Converter) convertType(
	of *ValueConverted, in, out string, typ gir.AnyType, owner string) *ValueConverted {

	attrs := of.ParameterAttrs
	attrs.AnyType = typ
	attrs.TransferOwnership.TransferOwnership = owner

	result := newValueConverted(conv, &ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      of.Direction,
		ParameterIndex: UnknownValueIndex,
		ParameterAttrs: attrs,
		InContainer:    of.Type != nil && of.Type.Type != nil, // is container type
	})

	// If the value is in a container, then its direction is always in. This is
	// because container-inner types are converted in a callback.
	if result.InContainer {
		result.ParameterAttrs.Direction = "in"
	}

	if !conv.convert(&result) {
		return nil
	}

	of.header.ImportImpl(result.Resolved)
	return &result
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

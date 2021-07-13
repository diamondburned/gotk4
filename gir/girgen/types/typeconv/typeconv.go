// Package typeconv provides conversions between C and Go types.
package typeconv

import (
	"log"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

type Converter struct {
	results []ValueConverted
	fgen    types.FileGenerator
	logger  logger.LineLogger
	header  file.Header

	// CurrentNamespace is the current namespace that this converter is
	// generating for.
	currentNamespace *gir.NamespaceFindResult
	// SourceNamespace is the namespace that the values are from.
	sourceNamespace *gir.NamespaceFindResult
}

// NewConverter creates a new type converter from the given file generator.
// The converter will add no side effects to the given file generator.
func NewConverter(fgen types.FileGenerator, values []ConversionValue) *Converter {
	conv := Converter{
		fgen:             fgen,
		results:          make([]ValueConverted, len(values)),
		sourceNamespace:  fgen.Namespace(),
		currentNamespace: fgen.Namespace(),
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
			conv.Logln(logger.Error,
				"value", value.InName, "->", value.OutName, "has invalid direction")
			return nil
		}

		if value.ParameterAttrs.Direction == "out" && !types.AnyTypeIsPtr(value.AnyType) {
			// Output direction but not pointer parameter is invalid; bail.
			conv.Logln(logger.Debug,
				"value type", types.AnyTypeC(value.AnyType), "is output but no ptr")
			return nil
		}

		// Only skip the parameter's closure index if the parameter itself is
		// a callback. Sometimes, the user_data parameter will flag the callback
		// as a closure argument, which messes up the generator.
		if value.Scope != "" && value.Closure != nil {
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
		conv.results[i] = newValueConverted(&conv, &values[i])
	}

	return &conv
}

// CCallParams generates the call parameters for calling the C function.
func (conv *Converter) CCallParams() []string {
	if conv == nil {
		return nil
	}

	params := make([]string, 0, len(conv.results))

	for _, result := range conv.results {
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

// SetSourceNamespace sets the converter's source namespace, which is the
// namespace that all the values originate from.
func (conv *Converter) SetSourceNamespace(ns *gir.NamespaceFindResult) {
	if conv != nil {
		conv.sourceNamespace = ns
	}
}

// SetSCurrentNamespace sets the converter's current namespace, which is the
// namespace that the converter is generating for.
func (conv *Converter) SetCurrentNamespace(ns *gir.NamespaceFindResult) {
	if conv != nil {
		conv.currentNamespace = ns
	}
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

	results := make([]ValueConverted, 0, len(conv.results))

	// Convert everything in one go.
	for i := range conv.results {
		if !conv.convert(&conv.results[i]) || conv.results[i].fail {
			return nil
		}
	}

	for i, result := range conv.results {
		if result.Skip {
			continue
		}

		result.finalize()
		file.ApplyHeader(conv, &conv.results[i])
		results = append(results, result)
	}

	return results
}

// Convert converts the value at the given index.
func (conv *Converter) Convert(i int) *ValueConverted {
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

	result.finalize()
	file.ApplyHeader(conv, result)
	return result
}

func (conv *Converter) convert(result *ValueConverted) bool {
	if result.isDone() {
		// result is already finalized, skip.
		return !result.fail
	}

	switch result.Direction {
	case ConvertCToGo:
		if !conv.cgoConvert(result) || result.fail {
			conv.Logln(logger.Debug, "C->Go cannot convert type", types.AnyTypeC(result.AnyType))
			result.fail = true
			return false
		}
	case ConvertGoToC:
		if !conv.gocConvert(result) || result.fail {
			conv.Logln(logger.Debug, "Go->C cannot convert type", types.AnyTypeC(result.AnyType))
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
	if of.AnyType.Array == nil {
		return nil
	}

	// If the array's ownership is ONLY container, then we must not take over
	// the inner values. Therefore, we only generate the appropriate code.
	owner := of.TransferOwnership.TransferOwnership
	if owner == "container" {
		owner = "none"
	}

	result := conv.convertType(of, in, out, of.AnyType.Array.AnyType, owner)
	if result == nil {
		return nil
	}

	// Set the array value's resolved type to the inner type.
	of.Resolved = result.Resolved
	of.NeedsNamespace = result.NeedsNamespace

	return result
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
	})

	if !conv.convert(&result) {
		return nil
	}

	of.header.ImportImpl(result.Resolved)
	return &result
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

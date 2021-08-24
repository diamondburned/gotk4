package typeconv

import (
	"log"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/types"
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

	// InContainer, if true, will increment the C type pointer by 1 to indicate
	// that the type is stored indirectly as a pointer in a container. A
	// container has to be a <type> wrapped inside another <type>.
	InContainer bool
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

// TODO: add field index support into ParameterIndex.

// NewFieldValue creates a new ConversionValue from the given C struct field.
// The struct is assumed to have a native field.
func NewFieldValue(in, out string, field gir.Field) ConversionValue {
	return ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      ConvertCToGo,
		ParameterIndex: ReturnValueIndex,
		ParameterAttrs: gir.ParameterAttrs{
			Name:    field.Name,
			Skip:    field.Private || !field.IsReadable() || field.Bits > 0,
			AnyType: field.AnyType,
			Doc:     field.Doc,
		},
	}
}

// NewFieldSetValue creates a new ConversionValue from the given C struct field
// that converts a Go to C type.
func NewFieldSetValue(in, out string, field gir.Field) ConversionValue {
	return ConversionValue{
		InName:         in,
		OutName:        out,
		Direction:      ConvertGoToC,
		ParameterIndex: ReturnValueIndex,
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
	// Containers are handled separately.
	if value.InContainer {
		return false
	}

	if value.ParameterIndex.Index() == -1 {
		// Not a parameter, but allow ErrorValue, since it is always an output
		// parameter.
		return value.ParameterIndex == ErrorValueIndex
	}

	param := gir.Parameter{
		ParameterAttrs: value.ParameterAttrs,
	}

	return types.GuessParameterOutput(&param) == "out"
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
	// if prop.Direction != ConvertCToGo {
	// 	goReceiving = !goReceiving
	// }

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
	// if prop.Direction != ConvertGoToC {
	// 	goGiving = !goGiving
	// }

	if goGiving {
		return prop.ownershipIsTransferring()
	}

	return !prop.ownershipIsTransferring()
}

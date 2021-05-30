package girgen

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
)

// ArgAtFunc is the function to get the argument name at the given index. This
// function is primarily used for certain type conversions that need to access
// multiple variables.
type ArgAtFunc func(i int) string

// TypeConversion describes the type information to convert from and to.
type TypeConversion struct {
	Value  string
	Target string
	Type   gir.AnyType
	Owner  gir.TransferOwnership

	// ArgAt is used for array and closure generation.
	ArgAt ArgAtFunc
	// ParentName is used primarily for debugging.
	ParentName string
}

// TypeConversionToC contains type information that is only useful when
// converting from Go to C.
type TypeConversionToC struct {
	TypeConversion

	// Closure marks the user_data argument. If this is provided, then the
	// conversion function will set the parameter to the callback ID. The caller
	// is responsible for skipping conversion of these indices.
	Closure *int
	// Destroy marks the callback to destroy the user_data argument. If this is
	// provided, then callbackDelete will be set along with Closure.
	Destroy *int
}

// TypeConversionToGo contains type information that is only useful when
// converting from C to Go.
type TypeConversionToGo struct {
	TypeConversion

	// BoxCast is an optional Go type that the boxed value should be casted to,
	// but only if the Type is a gpointer. This is only useful to convert from C
	// to Go.
	BoxCast string
}

// isTransferring is true when the ownership is either full or container. If the
// converter code isn't generating for an array, then distinguishing this
// doesn't matter. If the caller hasn't set the ownership yet, then it is
// assumed that we're not getting the ownership, therefore false is returned.
//
// If the generating code is an array, and the conversion is being passed into
// the same generation routine for the inner type, then the ownership should be
// turned into "none" just for that inner type. See TypeConversion.inner().
func (conv TypeConversion) isTransferring() bool {
	return false ||
		conv.Owner.TransferOwnership == "full" ||
		conv.Owner.TransferOwnership == "container"
}

// inner generates the proper type conversion for the underlying type, assuming
// the current TypeConversion is an array. It returns conv if the current type
// is not.
func (conv TypeConversion) inner(val, target string) TypeConversion {
	if conv.Type.Array == nil {
		return conv
	}

	conv.Value = val
	conv.Target = target
	conv.Type = conv.Type.Array.AnyType

	// If the array's ownership is ONLY container, then we must not take over
	// the inner values. Therefore, we only generate the appropriate code.
	if conv.Owner.TransferOwnership == "container" {
		conv.Owner.TransferOwnership = "none"
	}

	return conv
}

// call is a helper function around directCall.
func (conv TypeConversion) call(typ string) string {
	return directCall(conv.Value, conv.Target, typ)
}

// callf is a helper function around directCall and Sprintf.
func (conv TypeConversion) callf(typf string, typv ...interface{}) string {
	if len(typv) == 0 {
		return conv.call(typf)
	}
	return conv.call(fmt.Sprintf(typf, typv...))
}

// directCall generates a Go function call or type conversion that is
//
//    value = typ(target)
//
func directCall(value, target, typ string) string {
	if strings.Contains(typ, "*") {
		typ = "(" + typ + ")"
	}

	return target + " = " + typ + "(" + value + ")"
}

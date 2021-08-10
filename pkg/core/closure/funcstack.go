package closure

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// FrameSize is the number of frames that FuncStack should trace back from.
const FrameSize = 3

// FuncStack wraps a function value and provides function frames containing the
// caller trace for debugging.
type FuncStack struct {
	Func   interface{}
	Frames []uintptr
}

// NewFuncStack creates a new FuncStack. It panics if fn is not a function. The
// given frameSkip is added 2, meaning the first frame from 0 will start from
// the caller of NewFuncStack.
func NewFuncStack(fn interface{}, frameSkip int) *FuncStack {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic("closure value is not a func")
	}

	return newFuncStack(fn, frameSkip)
}

func newFuncStack(fn interface{}, frameSkip int) *FuncStack {
	frames := make([]uintptr, FrameSize)
	frames = frames[:runtime.Callers(frameSkip+3, frames)]

	return &FuncStack{
		Func:   fn,
		Frames: frames,
	}
}

// NewIdleFuncStack works akin to NewFuncStack, but it also validates the given
// function type for the correct acceptable signatures for SourceFunc while also
// caching the checks.
func NewIdleFuncStack(fn interface{}, frameSkip int) *FuncStack {
	switch fn.(type) {
	case func(), func() bool:
		return newFuncStack(fn, frameSkip)

	default:
		fs := newFuncStack(fn, frameSkip)
		fs.Panicf("unexpected func type %T, expected func() (error?)", fn)
	}

	return nil
}

// Func returns the function as a reflect.Value.
func (fs *FuncStack) Value() reflect.Value {
	return reflect.ValueOf(fs.Func)
}

// IsValid returns true if the given FuncStack is not a zero-value i.e.  valid.
func (fs *FuncStack) IsValid() bool {
	return fs != nil && fs.Frames != nil
}

const headerSignature = "closure error: "

// Panicf panics with the given FuncStack printed to standard error.
func (fs *FuncStack) Panicf(msgf string, v ...interface{}) {
	msg := strings.Builder{}
	msg.WriteString(headerSignature)
	fmt.Fprintf(&msg, msgf, v...)

	msg.WriteString("\n\nClosure added at:")

	frames := runtime.CallersFrames(fs.Frames)
	for {
		frame, more := frames.Next()
		msg.WriteString("\n\t")
		msg.WriteString(frame.Function)
		msg.WriteString(" at ")
		msg.WriteString(frame.File)
		msg.WriteByte(':')
		msg.WriteString(strconv.Itoa(frame.Line))

		if !more {
			break
		}
	}

	panic(msg.String())
}

// TryRepanic attempts to recover a panic. If successful, it will re-panic with
// the trace, or none if there is already one.
func (fs *FuncStack) TryRepanic() {
	panicking := recover()
	if panicking == nil {
		return
	}

	if msg, ok := panicking.(string); ok {
		if strings.Contains(msg, headerSignature) {
			// We can just repanic as-is.
			panic(msg)
		}
	}

	fs.Panicf("unexpected panic caught: %v", panicking)
}

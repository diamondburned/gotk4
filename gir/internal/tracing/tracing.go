package tracing

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// Traceable is a struct that holds a stack trace.
type Traceable struct {
	frames []uintptr
}

// NewTraces returns a new Traceable struct with the stack trace of the
// calling function.
func NewTraces(skip int) Traceable {
	var traces [4]uintptr
	n := runtime.Callers(2+skip, traces[:])
	return Traceable{traces[:n]}
}

// Frames returns the stack trace as a slice of uintptrs.
func (f Traceable) Frames() []uintptr {
	return f.frames
}

// Trace returns the trace as a string.
func (f Traceable) Trace() string {
	traces := make([]string, 0, len(f.frames))
	for _, frame := range f.frames {
		pc := runtime.FuncForPC(frame)
		if pc == nil {
			continue
		}

		name := pc.Name()
		if strings.HasPrefix(name, "runtime.") {
			continue
		}

		file, line := pc.FileLine(frame)
		file = filepath.Base(file)

		traces = append(traces, fmt.Sprintf("%s @ %s:%d", name, file, line))
	}
	return strings.Join(traces, ", ")
}

// Tracer is an interface for anything that has a stack trace pointing to
// where it was created.
type Tracer interface {
	Trace() string
}

// Trace returns the stack trace of the given value. If the value does not
// implement the Tracer interface, an empty string is returned.
func Trace(v any) string {
	if t, ok := v.(Tracer); ok {
		return t.Trace()
	}
	return ""
}

// IsTraceable returns true if the given value implements the Tracer
// interface.
func IsTraceable(v any) bool {
	_, ok := v.(Tracer)
	return ok
}

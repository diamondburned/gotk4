//go:build gotk4pprof

package intern

import (
	"runtime/pprof"
	"unsafe"
)

var profiler = pprof.NewProfile("gotk4-object-box")

func profileRecordObject(ptr unsafe.Pointer, skip int) {
	profiler.Add(ptr, skip+1)
}

func profileRemoveObject(ptr unsafe.Pointer) {
	profiler.Remove(ptr)
}

//go:build !gotk4pprof

package intern

import "unsafe"

func profileRecordObject(ptr unsafe.Pointer, skip int) {}

func profileRemoveObject(ptr unsafe.Pointer) {}

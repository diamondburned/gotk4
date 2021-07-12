package gbox

import (
	"sync"

	"github.com/diamondburned/gotk4/pkg/core/slab"
)

//export callbackDelete
func callbackDelete(ptr uintptr) {
	Delete(ptr)
}

var registry struct {
	sync.RWMutex
	slab slab.Slab
}

// Assign assigns the given value and returns the fake pointer.
func Assign(v interface{}) uintptr {
	registry.Lock()
	defer registry.Unlock()

	return registry.slab.Put(v)
}

// Get gets the value from the given fake pointer. The context must match the
// given value in Assign.
func Get(ptr uintptr) interface{} {
	registry.RLock()
	defer registry.RUnlock()

	return registry.slab.Get(ptr)
}

// Delete deletes a boxed value.
func Delete(ptr uintptr) {
	Pop(ptr)
}

// Pop gets a value and deletes it atomically.
func Pop(ptr uintptr) interface{} {
	registry.Lock()
	defer registry.Unlock()

	return registry.slab.Pop(ptr)
}

package box

import (
	"sync"

	"github.com/diamondburned/gotk4/internal/slab"
)

type Context uint8

const (
	// Boxed stores user data.
	Boxed Context = iota
	// Callback stores Go closures.
	Callback

	max
)

var registries [max]struct {
	sync.RWMutex
	slab slab.Slab
}

// Assign assigns the given value and returns the fake pointer.
func Assign(ctx Context, v interface{}) uintptr {
	registries[ctx].Lock()
	defer registries[ctx].Unlock()

	return registries[ctx].slab.Put(v)
}

// Get gets the value from the given fake pointer. The context must match the
// given value in Assign.
func Get(ctx Context, ptr uintptr) interface{} {
	registries[ctx].RLock()
	defer registries[ctx].RUnlock()

	return registries[ctx].slab.Get(ptr)
}

// Delete deletes a boxed value.
func Delete(ctx Context, ptr uintptr) {
	GetAndDelete(ctx, ptr)
}

// GetAndDelete gets a value and deletes it atomically.
func GetAndDelete(ctx Context, ptr uintptr) interface{} {
	registries[ctx].Lock()
	defer registries[ctx].Unlock()

	return registries[ctx].slab.Pop(ptr)
}

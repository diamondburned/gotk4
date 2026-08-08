package closure

import (
	"sync"
	"unsafe"
)

// Registry describes the local closure registry of each object.
type Registry struct {
	reg sync.Map // unsafe.Pointer(*C.GClosure) -> registeredClosure
}

type registeredClosure struct {
	callback *FuncStack
	signal   bool
}

// NewRegistry creates an empty closure registry.
func NewRegistry() *Registry {
	return &Registry{}
}

// Register registers the given GClosure callback.
func (r *Registry) Register(gclosure unsafe.Pointer, callback *FuncStack) {
	r.reg.Store(uintptr(gclosure), registeredClosure{callback: callback})
}

// RegisterSignal registers a callback attached to a GObject signal. Signal
// closures can be disconnected during explicit object cleanup; closures used
// by other APIs must be released by those APIs instead.
func (r *Registry) RegisterSignal(gclosure unsafe.Pointer, callback *FuncStack) {
	r.reg.Store(uintptr(gclosure), registeredClosure{callback: callback, signal: true})
}

// Load loads the given GClosure's callback. Nil is returned if it's not found.
func (r *Registry) Load(gclosure unsafe.Pointer) *FuncStack {
	fs, ok := r.reg.Load(uintptr(gclosure))
	if !ok {
		return nil
	}
	return fs.(registeredClosure).callback
}

// Delete deletes the given GClosure callback.
func (r *Registry) Delete(gclosure unsafe.Pointer) {
	r.reg.Delete(uintptr(gclosure))
}

// RangeSignals calls f for every registered signal closure. Returning false
// stops the iteration. The callback may delete the current closure.
func (r *Registry) RangeSignals(f func(unsafe.Pointer, *FuncStack) bool) {
	r.reg.Range(func(key, value any) bool {
		entry := value.(registeredClosure)
		if !entry.signal {
			return true
		}
		return f(unsafe.Pointer(key.(uintptr)), entry.callback)
	})
}

// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"

//export _gotk4_gtk4_SpinButton_ConnectChangeValue
func _gotk4_gtk4_SpinButton_ConnectChangeValue(arg0 C.gpointer, arg1 C.GtkScrollType, arg2 C.guintptr) {
	var f func(scroll ScrollType)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(scroll ScrollType))
	}

	var _scroll ScrollType // out

	_scroll = ScrollType(arg1)

	f(_scroll)
}

//export _gotk4_gtk4_SpinButton_ConnectOutput
func _gotk4_gtk4_SpinButton_ConnectOutput(arg0 C.gpointer, arg1 C.guintptr) (cret C.gboolean) {
	var f func() (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func() (ok bool))
	}

	ok := f()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk4_SpinButton_ConnectValueChanged
func _gotk4_gtk4_SpinButton_ConnectValueChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gtk4_SpinButton_ConnectWrapped
func _gotk4_gtk4_SpinButton_ConnectWrapped(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}
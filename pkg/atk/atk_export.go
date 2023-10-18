// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
import "C"

//export _gotk4_atk1_Function
func _gotk4_atk1_Function(arg1 C.gpointer) (cret C.gboolean) {
	var fn Function
	{
		v := gbox.Get(uintptr(arg1))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(Function)
	}

	ok := fn()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_KeySnoopFunc
func _gotk4_atk1_KeySnoopFunc(arg1 *C.AtkKeyEventStruct, arg2 C.gpointer) (cret C.gint) {
	var fn KeySnoopFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(KeySnoopFunc)
	}

	var _event *KeyEventStruct // out

	_event = (*KeyEventStruct)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	gint := fn(_event)

	var _ int

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_Component_ConnectBoundsChanged
func _gotk4_atk1_Component_ConnectBoundsChanged(arg0 C.gpointer, arg1 *C.AtkRectangle, arg2 C.guintptr) {
	var f func(arg1 *Rectangle)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 *Rectangle))
	}

	var _arg1 *Rectangle // out

	_arg1 = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_arg1)
}

//export _gotk4_atk1_Document_ConnectLoadComplete
func _gotk4_atk1_Document_ConnectLoadComplete(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Document_ConnectLoadStopped
func _gotk4_atk1_Document_ConnectLoadStopped(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Document_ConnectPageChanged
func _gotk4_atk1_Document_ConnectPageChanged(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(pageNumber int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(pageNumber int))
	}

	var _pageNumber int // out

	_pageNumber = int(arg1)

	f(_pageNumber)
}

//export _gotk4_atk1_Document_ConnectReload
func _gotk4_atk1_Document_ConnectReload(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Hypertext_ConnectLinkSelected
func _gotk4_atk1_Hypertext_ConnectLinkSelected(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(arg1 int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 int))
	}

	var _arg1 int // out

	_arg1 = int(arg1)

	f(_arg1)
}

//export _gotk4_atk1_Selection_ConnectSelectionChanged
func _gotk4_atk1_Selection_ConnectSelectionChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Table_ConnectColumnDeleted
func _gotk4_atk1_Table_ConnectColumnDeleted(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) {
	var f func(arg1, arg2 int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1, arg2 int))
	}

	var _arg1 int // out
	var _arg2 int // out

	_arg1 = int(arg1)
	_arg2 = int(arg2)

	f(_arg1, _arg2)
}

//export _gotk4_atk1_Table_ConnectColumnInserted
func _gotk4_atk1_Table_ConnectColumnInserted(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) {
	var f func(arg1, arg2 int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1, arg2 int))
	}

	var _arg1 int // out
	var _arg2 int // out

	_arg1 = int(arg1)
	_arg2 = int(arg2)

	f(_arg1, _arg2)
}

//export _gotk4_atk1_Table_ConnectColumnReordered
func _gotk4_atk1_Table_ConnectColumnReordered(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Table_ConnectModelChanged
func _gotk4_atk1_Table_ConnectModelChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Table_ConnectRowDeleted
func _gotk4_atk1_Table_ConnectRowDeleted(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) {
	var f func(arg1, arg2 int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1, arg2 int))
	}

	var _arg1 int // out
	var _arg2 int // out

	_arg1 = int(arg1)
	_arg2 = int(arg2)

	f(_arg1, _arg2)
}

//export _gotk4_atk1_Table_ConnectRowInserted
func _gotk4_atk1_Table_ConnectRowInserted(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) {
	var f func(arg1, arg2 int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1, arg2 int))
	}

	var _arg1 int // out
	var _arg2 int // out

	_arg1 = int(arg1)
	_arg2 = int(arg2)

	f(_arg1, _arg2)
}

//export _gotk4_atk1_Table_ConnectRowReordered
func _gotk4_atk1_Table_ConnectRowReordered(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Text_ConnectTextAttributesChanged
func _gotk4_atk1_Text_ConnectTextAttributesChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Text_ConnectTextCaretMoved
func _gotk4_atk1_Text_ConnectTextCaretMoved(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(arg1 int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 int))
	}

	var _arg1 int // out

	_arg1 = int(arg1)

	f(_arg1)
}

//export _gotk4_atk1_Text_ConnectTextChanged
func _gotk4_atk1_Text_ConnectTextChanged(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) {
	var f func(arg1, arg2 int)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1, arg2 int))
	}

	var _arg1 int // out
	var _arg2 int // out

	_arg1 = int(arg1)
	_arg2 = int(arg2)

	f(_arg1, _arg2)
}

//export _gotk4_atk1_Text_ConnectTextInsert
func _gotk4_atk1_Text_ConnectTextInsert(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 *C.gchar, arg4 C.guintptr) {
	var f func(arg1, arg2 int, arg3 string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1, arg2 int, arg3 string))
	}

	var _arg1 int    // out
	var _arg2 int    // out
	var _arg3 string // out

	_arg1 = int(arg1)
	_arg2 = int(arg2)
	_arg3 = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))

	f(_arg1, _arg2, _arg3)
}

//export _gotk4_atk1_Text_ConnectTextRemove
func _gotk4_atk1_Text_ConnectTextRemove(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 *C.gchar, arg4 C.guintptr) {
	var f func(arg1, arg2 int, arg3 string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1, arg2 int, arg3 string))
	}

	var _arg1 int    // out
	var _arg2 int    // out
	var _arg3 string // out

	_arg1 = int(arg1)
	_arg2 = int(arg2)
	_arg3 = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))

	f(_arg1, _arg2, _arg3)
}

//export _gotk4_atk1_Text_ConnectTextSelectionChanged
func _gotk4_atk1_Text_ConnectTextSelectionChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Value_ConnectValueChanged
func _gotk4_atk1_Value_ConnectValueChanged(arg0 C.gpointer, arg1 C.gdouble, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(value float64, text string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(value float64, text string))
	}

	var _value float64 // out
	var _text string   // out

	_value = float64(arg1)
	_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_value, _text)
}

//export _gotk4_atk1_Window_ConnectActivate
func _gotk4_atk1_Window_ConnectActivate(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Window_ConnectCreate
func _gotk4_atk1_Window_ConnectCreate(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Window_ConnectDeactivate
func _gotk4_atk1_Window_ConnectDeactivate(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Window_ConnectDestroy
func _gotk4_atk1_Window_ConnectDestroy(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Window_ConnectMaximize
func _gotk4_atk1_Window_ConnectMaximize(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Window_ConnectMinimize
func _gotk4_atk1_Window_ConnectMinimize(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Window_ConnectMove
func _gotk4_atk1_Window_ConnectMove(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Window_ConnectResize
func _gotk4_atk1_Window_ConnectResize(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_Window_ConnectRestore
func _gotk4_atk1_Window_ConnectRestore(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_HyperlinkClass_get_end_index
func _gotk4_atk1_HyperlinkClass_get_end_index(arg0 *C.AtkHyperlink) (cret C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[HyperlinkOverrides](instance0)
	if overrides.EndIndex == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected HyperlinkOverrides.EndIndex, got none")
	}

	gint := overrides.EndIndex()

	var _ int

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_HyperlinkClass_get_n_anchors
func _gotk4_atk1_HyperlinkClass_get_n_anchors(arg0 *C.AtkHyperlink) (cret C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[HyperlinkOverrides](instance0)
	if overrides.NAnchors == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected HyperlinkOverrides.NAnchors, got none")
	}

	gint := overrides.NAnchors()

	var _ int

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_HyperlinkClass_get_object
func _gotk4_atk1_HyperlinkClass_get_object(arg0 *C.AtkHyperlink, arg1 C.gint) (cret *C.AtkObject) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[HyperlinkOverrides](instance0)
	if overrides.GetObject == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected HyperlinkOverrides.GetObject, got none")
	}

	var _i int // out

	_i = int(arg1)

	object := overrides.GetObject(_i)

	var _ *AtkObject

	cret = (*C.AtkObject)(unsafe.Pointer(coreglib.InternObject(object).Native()))

	return cret
}

//export _gotk4_atk1_HyperlinkClass_get_start_index
func _gotk4_atk1_HyperlinkClass_get_start_index(arg0 *C.AtkHyperlink) (cret C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[HyperlinkOverrides](instance0)
	if overrides.StartIndex == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected HyperlinkOverrides.StartIndex, got none")
	}

	gint := overrides.StartIndex()

	var _ int

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_HyperlinkClass_get_uri
func _gotk4_atk1_HyperlinkClass_get_uri(arg0 *C.AtkHyperlink, arg1 C.gint) (cret *C.gchar) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[HyperlinkOverrides](instance0)
	if overrides.URI == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected HyperlinkOverrides.URI, got none")
	}

	var _i int // out

	_i = int(arg1)

	utf8 := overrides.URI(_i)

	var _ string

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

//export _gotk4_atk1_HyperlinkClass_is_selected_link
func _gotk4_atk1_HyperlinkClass_is_selected_link(arg0 *C.AtkHyperlink) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[HyperlinkOverrides](instance0)
	if overrides.IsSelectedLink == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected HyperlinkOverrides.IsSelectedLink, got none")
	}

	ok := overrides.IsSelectedLink()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_HyperlinkClass_is_valid
func _gotk4_atk1_HyperlinkClass_is_valid(arg0 *C.AtkHyperlink) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[HyperlinkOverrides](instance0)
	if overrides.IsValid == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected HyperlinkOverrides.IsValid, got none")
	}

	ok := overrides.IsValid()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_atk1_HyperlinkClass_link_activated
func _gotk4_atk1_HyperlinkClass_link_activated(arg0 *C.AtkHyperlink) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[HyperlinkOverrides](instance0)
	if overrides.LinkActivated == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected HyperlinkOverrides.LinkActivated, got none")
	}

	overrides.LinkActivated()
}

//export _gotk4_atk1_HyperlinkClass_link_state
func _gotk4_atk1_HyperlinkClass_link_state(arg0 *C.AtkHyperlink) (cret C.guint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[HyperlinkOverrides](instance0)
	if overrides.LinkState == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected HyperlinkOverrides.LinkState, got none")
	}

	guint := overrides.LinkState()

	var _ uint

	cret = C.guint(guint)

	return cret
}

//export _gotk4_atk1_Hyperlink_ConnectLinkActivated
func _gotk4_atk1_Hyperlink_ConnectLinkActivated(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_MiscClass_threads_enter
func _gotk4_atk1_MiscClass_threads_enter(arg0 *C.AtkMisc) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[MiscOverrides](instance0)
	if overrides.ThreadsEnter == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected MiscOverrides.ThreadsEnter, got none")
	}

	overrides.ThreadsEnter()
}

//export _gotk4_atk1_MiscClass_threads_leave
func _gotk4_atk1_MiscClass_threads_leave(arg0 *C.AtkMisc) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[MiscOverrides](instance0)
	if overrides.ThreadsLeave == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected MiscOverrides.ThreadsLeave, got none")
	}

	overrides.ThreadsLeave()
}

//export _gotk4_atk1_ObjectClass_active_descendant_changed
func _gotk4_atk1_ObjectClass_active_descendant_changed(arg0 *C.AtkObject, arg1 *C.gpointer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.ActiveDescendantChanged == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.ActiveDescendantChanged, got none")
	}

	var _child *unsafe.Pointer // out

	if arg1 != nil {
		_child = (*unsafe.Pointer)(unsafe.Pointer(arg1))
	}

	overrides.ActiveDescendantChanged(_child)
}

//export _gotk4_atk1_ObjectClass_children_changed
func _gotk4_atk1_ObjectClass_children_changed(arg0 *C.AtkObject, arg1 C.guint, arg2 C.gpointer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.ChildrenChanged == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.ChildrenChanged, got none")
	}

	var _changeIndex uint            // out
	var _changedChild unsafe.Pointer // out

	_changeIndex = uint(arg1)
	_changedChild = (unsafe.Pointer)(unsafe.Pointer(arg2))

	overrides.ChildrenChanged(_changeIndex, _changedChild)
}

//export _gotk4_atk1_ObjectClass_focus_event
func _gotk4_atk1_ObjectClass_focus_event(arg0 *C.AtkObject, arg1 C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.FocusEvent == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.FocusEvent, got none")
	}

	var _focusIn bool // out

	if arg1 != 0 {
		_focusIn = true
	}

	overrides.FocusEvent(_focusIn)
}

//export _gotk4_atk1_ObjectClass_get_description
func _gotk4_atk1_ObjectClass_get_description(arg0 *C.AtkObject) (cret *C.gchar) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.Description == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.Description, got none")
	}

	utf8 := overrides.Description()

	var _ string

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_atk1_ObjectClass_get_index_in_parent
func _gotk4_atk1_ObjectClass_get_index_in_parent(arg0 *C.AtkObject) (cret C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.IndexInParent == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.IndexInParent, got none")
	}

	gint := overrides.IndexInParent()

	var _ int

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_ObjectClass_get_layer
func _gotk4_atk1_ObjectClass_get_layer(arg0 *C.AtkObject) (cret C.AtkLayer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.Layer == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.Layer, got none")
	}

	layer := overrides.Layer()

	var _ Layer

	cret = C.AtkLayer(layer)

	return cret
}

//export _gotk4_atk1_ObjectClass_get_mdi_zorder
func _gotk4_atk1_ObjectClass_get_mdi_zorder(arg0 *C.AtkObject) (cret C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.MDIZOrder == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.MDIZOrder, got none")
	}

	gint := overrides.MDIZOrder()

	var _ int

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_ObjectClass_get_n_children
func _gotk4_atk1_ObjectClass_get_n_children(arg0 *C.AtkObject) (cret C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.NChildren == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.NChildren, got none")
	}

	gint := overrides.NChildren()

	var _ int

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_ObjectClass_get_name
func _gotk4_atk1_ObjectClass_get_name(arg0 *C.AtkObject) (cret *C.gchar) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.Name == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.Name, got none")
	}

	utf8 := overrides.Name()

	var _ string

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_atk1_ObjectClass_get_object_locale
func _gotk4_atk1_ObjectClass_get_object_locale(arg0 *C.AtkObject) (cret *C.gchar) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.ObjectLocale == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.ObjectLocale, got none")
	}

	utf8 := overrides.ObjectLocale()

	var _ string

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_atk1_ObjectClass_get_parent
func _gotk4_atk1_ObjectClass_get_parent(arg0 *C.AtkObject) (cret *C.AtkObject) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.Parent == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.Parent, got none")
	}

	object := overrides.Parent()

	var _ *AtkObject

	cret = (*C.AtkObject)(unsafe.Pointer(coreglib.InternObject(object).Native()))

	return cret
}

//export _gotk4_atk1_ObjectClass_get_role
func _gotk4_atk1_ObjectClass_get_role(arg0 *C.AtkObject) (cret C.AtkRole) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.Role == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.Role, got none")
	}

	role := overrides.Role()

	var _ Role

	cret = C.AtkRole(role)

	return cret
}

//export _gotk4_atk1_ObjectClass_initialize
func _gotk4_atk1_ObjectClass_initialize(arg0 *C.AtkObject, arg1 C.gpointer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.Initialize == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.Initialize, got none")
	}

	var _data unsafe.Pointer // out

	_data = (unsafe.Pointer)(unsafe.Pointer(arg1))

	overrides.Initialize(_data)
}

//export _gotk4_atk1_ObjectClass_property_change
func _gotk4_atk1_ObjectClass_property_change(arg0 *C.AtkObject, arg1 *C.AtkPropertyValues) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.PropertyChange == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.PropertyChange, got none")
	}

	var _values *PropertyValues // out

	_values = (*PropertyValues)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.PropertyChange(_values)
}

//export _gotk4_atk1_ObjectClass_ref_relation_set
func _gotk4_atk1_ObjectClass_ref_relation_set(arg0 *C.AtkObject) (cret *C.AtkRelationSet) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.RefRelationSet == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.RefRelationSet, got none")
	}

	relationSet := overrides.RefRelationSet()

	var _ *RelationSet

	cret = (*C.AtkRelationSet)(unsafe.Pointer(coreglib.InternObject(relationSet).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(relationSet).Native()))

	return cret
}

//export _gotk4_atk1_ObjectClass_ref_state_set
func _gotk4_atk1_ObjectClass_ref_state_set(arg0 *C.AtkObject) (cret *C.AtkStateSet) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.RefStateSet == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.RefStateSet, got none")
	}

	stateSet := overrides.RefStateSet()

	var _ *StateSet

	cret = (*C.AtkStateSet)(unsafe.Pointer(coreglib.InternObject(stateSet).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(stateSet).Native()))

	return cret
}

//export _gotk4_atk1_ObjectClass_remove_property_change_handler
func _gotk4_atk1_ObjectClass_remove_property_change_handler(arg0 *C.AtkObject, arg1 C.guint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.RemovePropertyChangeHandler == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.RemovePropertyChangeHandler, got none")
	}

	var _handlerId uint // out

	_handlerId = uint(arg1)

	overrides.RemovePropertyChangeHandler(_handlerId)
}

//export _gotk4_atk1_ObjectClass_set_description
func _gotk4_atk1_ObjectClass_set_description(arg0 *C.AtkObject, arg1 *C.gchar) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.SetDescription == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.SetDescription, got none")
	}

	var _description string // out

	_description = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	overrides.SetDescription(_description)
}

//export _gotk4_atk1_ObjectClass_set_name
func _gotk4_atk1_ObjectClass_set_name(arg0 *C.AtkObject, arg1 *C.gchar) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.SetName == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.SetName, got none")
	}

	var _name string // out

	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	overrides.SetName(_name)
}

//export _gotk4_atk1_ObjectClass_set_parent
func _gotk4_atk1_ObjectClass_set_parent(arg0 *C.AtkObject, arg1 *C.AtkObject) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.SetParent == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.SetParent, got none")
	}

	var _parent *AtkObject // out

	_parent = wrapObject(coreglib.Take(unsafe.Pointer(arg1)))

	overrides.SetParent(_parent)
}

//export _gotk4_atk1_ObjectClass_set_role
func _gotk4_atk1_ObjectClass_set_role(arg0 *C.AtkObject, arg1 C.AtkRole) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.SetRole == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.SetRole, got none")
	}

	var _role Role // out

	_role = Role(arg1)

	overrides.SetRole(_role)
}

//export _gotk4_atk1_ObjectClass_state_change
func _gotk4_atk1_ObjectClass_state_change(arg0 *C.AtkObject, arg1 *C.gchar, arg2 C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.StateChange == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.StateChange, got none")
	}

	var _name string   // out
	var _stateSet bool // out

	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	if arg2 != 0 {
		_stateSet = true
	}

	overrides.StateChange(_name, _stateSet)
}

//export _gotk4_atk1_ObjectClass_visible_data_changed
func _gotk4_atk1_ObjectClass_visible_data_changed(arg0 *C.AtkObject) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AtkObjectOverrides](instance0)
	if overrides.VisibleDataChanged == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AtkObjectOverrides.VisibleDataChanged, got none")
	}

	overrides.VisibleDataChanged()
}

//export _gotk4_atk1_Object_ConnectActiveDescendantChanged
func _gotk4_atk1_Object_ConnectActiveDescendantChanged(arg0 C.gpointer, arg1 *C.gpointer, arg2 C.guintptr) {
	var f func(arg1 *AtkObject)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 *AtkObject))
	}

	var _arg1 *AtkObject // out

	_arg1 = wrapObject(coreglib.Take(unsafe.Pointer(arg1)))

	f(_arg1)
}

//export _gotk4_atk1_Object_ConnectChildrenChanged
func _gotk4_atk1_Object_ConnectChildrenChanged(arg0 C.gpointer, arg1 C.guint, arg2 *C.gpointer, arg3 C.guintptr) {
	var f func(arg1 uint, arg2 *AtkObject)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 uint, arg2 *AtkObject))
	}

	var _arg1 uint       // out
	var _arg2 *AtkObject // out

	_arg1 = uint(arg1)
	_arg2 = wrapObject(coreglib.Take(unsafe.Pointer(arg2)))

	f(_arg1, _arg2)
}

//export _gotk4_atk1_Object_ConnectFocusEvent
func _gotk4_atk1_Object_ConnectFocusEvent(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) {
	var f func(arg1 bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 bool))
	}

	var _arg1 bool // out

	if arg1 != 0 {
		_arg1 = true
	}

	f(_arg1)
}

//export _gotk4_atk1_Object_ConnectPropertyChange
func _gotk4_atk1_Object_ConnectPropertyChange(arg0 C.gpointer, arg1 *C.gpointer, arg2 C.guintptr) {
	var f func(arg1 *PropertyValues)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 *PropertyValues))
	}

	var _arg1 *PropertyValues // out

	_arg1 = (*PropertyValues)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_arg1)
}

//export _gotk4_atk1_Object_ConnectStateChange
func _gotk4_atk1_Object_ConnectStateChange(arg0 C.gpointer, arg1 *C.gchar, arg2 C.gboolean, arg3 C.guintptr) {
	var f func(arg1 string, arg2 bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(arg1 string, arg2 bool))
	}

	var _arg1 string // out
	var _arg2 bool   // out

	_arg1 = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	if arg2 != 0 {
		_arg2 = true
	}

	f(_arg1, _arg2)
}

//export _gotk4_atk1_Object_ConnectVisibleDataChanged
func _gotk4_atk1_Object_ConnectVisibleDataChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_atk1_ObjectFactoryClass_invalidate
func _gotk4_atk1_ObjectFactoryClass_invalidate(arg0 *C.AtkObjectFactory) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ObjectFactoryOverrides](instance0)
	if overrides.Invalidate == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ObjectFactoryOverrides.Invalidate, got none")
	}

	overrides.Invalidate()
}

//export _gotk4_atk1_PlugClass_get_object_id
func _gotk4_atk1_PlugClass_get_object_id(arg0 *C.AtkPlug) (cret *C.gchar) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[PlugOverrides](instance0)
	if overrides.ObjectID == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected PlugOverrides.ObjectID, got none")
	}

	utf8 := overrides.ObjectID()

	var _ string

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))

	return cret
}

//export _gotk4_atk1_SocketClass_embed
func _gotk4_atk1_SocketClass_embed(arg0 *C.AtkSocket, arg1 *C.gchar) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[SocketOverrides](instance0)
	if overrides.Embed == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected SocketOverrides.Embed, got none")
	}

	var _plugId string // out

	_plugId = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	overrides.Embed(_plugId)
}
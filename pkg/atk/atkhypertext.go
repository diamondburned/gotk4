// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
// extern AtkHyperlink* _gotk4_atk1_HypertextIface_get_link(AtkHypertext*, gint);
// extern gint _gotk4_atk1_HypertextIface_get_link_index(AtkHypertext*, gint);
// extern gint _gotk4_atk1_HypertextIface_get_n_links(AtkHypertext*);
// extern void _gotk4_atk1_HypertextIface_link_selected(AtkHypertext*, gint);
// extern void _gotk4_atk1_Hypertext_ConnectLinkSelected(gpointer, gint, guintptr);
import "C"

// GType values.
var (
	GTypeHypertext = coreglib.Type(C.atk_hypertext_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHypertext, F: marshalHypertext},
	})
}

// HypertextOverrider contains methods that are overridable.
type HypertextOverrider interface {
	// Link gets the link in this hypertext document at index link_index.
	//
	// The function takes the following parameters:
	//
	//    - linkIndex: integer specifying the desired link.
	//
	// The function returns the following values:
	//
	//    - hyperlink: link in this hypertext document at index link_index.
	//
	Link(linkIndex int) *Hyperlink
	// LinkIndex gets the index into the array of hyperlinks that is associated
	// with the character specified by char_index.
	//
	// The function takes the following parameters:
	//
	//    - charIndex: character index.
	//
	// The function returns the following values:
	//
	//    - gint: index into the array of hyperlinks in hypertext, or -1 if there
	//      is no hyperlink associated with this character.
	//
	LinkIndex(charIndex int) int
	// NLinks gets the number of links within this hypertext document.
	//
	// The function returns the following values:
	//
	//    - gint: number of links within this hypertext document.
	//
	NLinks() int
	// The function takes the following parameters:
	//
	LinkSelected(linkIndex int)
}

// Hypertext: interface used for objects which implement linking between
// multiple resource or content locations, or multiple 'markers' within a single
// document. A Hypertext instance is associated with one or more Hyperlinks,
// which are associated with particular offsets within the Hypertext's included
// content. While this interface is derived from Text, there is no requirement
// that Hypertext instances have textual content; they may implement Image as
// well, and Hyperlinks need not have non-zero text offsets.
//
// Hypertext wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Hypertext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Hypertext)(nil)
)

// Hypertexter describes Hypertext's interface methods.
type Hypertexter interface {
	coreglib.Objector

	// Link gets the link in this hypertext document at index link_index.
	Link(linkIndex int) *Hyperlink
	// LinkIndex gets the index into the array of hyperlinks that is associated
	// with the character specified by char_index.
	LinkIndex(charIndex int) int
	// NLinks gets the number of links within this hypertext document.
	NLinks() int

	// Link-selected: "link-selected" signal is emitted by an AtkHyperText
	// object when one of the hyperlinks associated with the object is selected.
	ConnectLinkSelected(func(arg1 int)) coreglib.SignalHandle
}

var _ Hypertexter = (*Hypertext)(nil)

func ifaceInitHypertexter(gifacePtr, data C.gpointer) {
	iface := (*C.AtkHypertextIface)(unsafe.Pointer(gifacePtr))
	iface.get_link = (*[0]byte)(C._gotk4_atk1_HypertextIface_get_link)
	iface.get_link_index = (*[0]byte)(C._gotk4_atk1_HypertextIface_get_link_index)
	iface.get_n_links = (*[0]byte)(C._gotk4_atk1_HypertextIface_get_n_links)
	iface.link_selected = (*[0]byte)(C._gotk4_atk1_HypertextIface_link_selected)
}

//export _gotk4_atk1_HypertextIface_get_link
func _gotk4_atk1_HypertextIface_get_link(arg0 *C.AtkHypertext, arg1 C.gint) (cret *C.AtkHyperlink) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(HypertextOverrider)

	var _linkIndex int // out

	_linkIndex = int(arg1)

	hyperlink := iface.Link(_linkIndex)

	cret = (*C.AtkHyperlink)(unsafe.Pointer(coreglib.InternObject(hyperlink).Native()))

	return cret
}

//export _gotk4_atk1_HypertextIface_get_link_index
func _gotk4_atk1_HypertextIface_get_link_index(arg0 *C.AtkHypertext, arg1 C.gint) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(HypertextOverrider)

	var _charIndex int // out

	_charIndex = int(arg1)

	gint := iface.LinkIndex(_charIndex)

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_HypertextIface_get_n_links
func _gotk4_atk1_HypertextIface_get_n_links(arg0 *C.AtkHypertext) (cret C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(HypertextOverrider)

	gint := iface.NLinks()

	cret = C.gint(gint)

	return cret
}

//export _gotk4_atk1_HypertextIface_link_selected
func _gotk4_atk1_HypertextIface_link_selected(arg0 *C.AtkHypertext, arg1 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(HypertextOverrider)

	var _linkIndex int // out

	_linkIndex = int(arg1)

	iface.LinkSelected(_linkIndex)
}

func wrapHypertext(obj *coreglib.Object) *Hypertext {
	return &Hypertext{
		Object: obj,
	}
}

func marshalHypertext(p uintptr) (interface{}, error) {
	return wrapHypertext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
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

// ConnectLinkSelected: "link-selected" signal is emitted by an AtkHyperText
// object when one of the hyperlinks associated with the object is selected.
func (hypertext *Hypertext) ConnectLinkSelected(f func(arg1 int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(hypertext, "link-selected", false, unsafe.Pointer(C._gotk4_atk1_Hypertext_ConnectLinkSelected), f)
}

// Link gets the link in this hypertext document at index link_index.
//
// The function takes the following parameters:
//
//    - linkIndex: integer specifying the desired link.
//
// The function returns the following values:
//
//    - hyperlink: link in this hypertext document at index link_index.
//
func (hypertext *Hypertext) Link(linkIndex int) *Hyperlink {
	var _arg0 *C.AtkHypertext // out
	var _arg1 C.gint          // out
	var _cret *C.AtkHyperlink // in

	_arg0 = (*C.AtkHypertext)(unsafe.Pointer(coreglib.InternObject(hypertext).Native()))
	_arg1 = C.gint(linkIndex)

	_cret = C.atk_hypertext_get_link(_arg0, _arg1)
	runtime.KeepAlive(hypertext)
	runtime.KeepAlive(linkIndex)

	var _hyperlink *Hyperlink // out

	_hyperlink = wrapHyperlink(coreglib.Take(unsafe.Pointer(_cret)))

	return _hyperlink
}

// LinkIndex gets the index into the array of hyperlinks that is associated with
// the character specified by char_index.
//
// The function takes the following parameters:
//
//    - charIndex: character index.
//
// The function returns the following values:
//
//    - gint: index into the array of hyperlinks in hypertext, or -1 if there is
//      no hyperlink associated with this character.
//
func (hypertext *Hypertext) LinkIndex(charIndex int) int {
	var _arg0 *C.AtkHypertext // out
	var _arg1 C.gint          // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkHypertext)(unsafe.Pointer(coreglib.InternObject(hypertext).Native()))
	_arg1 = C.gint(charIndex)

	_cret = C.atk_hypertext_get_link_index(_arg0, _arg1)
	runtime.KeepAlive(hypertext)
	runtime.KeepAlive(charIndex)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NLinks gets the number of links within this hypertext document.
//
// The function returns the following values:
//
//    - gint: number of links within this hypertext document.
//
func (hypertext *Hypertext) NLinks() int {
	var _arg0 *C.AtkHypertext // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkHypertext)(unsafe.Pointer(coreglib.InternObject(hypertext).Native()))

	_cret = C.atk_hypertext_get_n_links(_arg0)
	runtime.KeepAlive(hypertext)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

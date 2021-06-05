// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_popover_get_type()), F: marshalPopover},
	})
}

// Popover: gtkPopover is a bubble-like context window, primarily meant to
// provide context-dependent information or options. Popovers are attached to a
// widget, set with gtk_widget_set_parent(). By default they will point to the
// whole widget area, although this behavior can be changed through
// gtk_popover_set_pointing_to().
//
// The position of a popover relative to the widget it is attached to can also
// be changed through gtk_popover_set_position().
//
// By default, Popover performs a grab, in order to ensure input events get
// redirected to it while it is shown, and also so the popover is dismissed in
// the expected situations (clicks outside the popover, or the Escape key being
// pressed). If no such modal behavior is desired on a popover,
// gtk_popover_set_autohide() may be called on it to tweak its behavior.
//
//
// GtkPopover as menu replacement
//
// GtkPopover is often used to replace menus. The best was to do this is to use
// the PopoverMenu subclass which supports being populated from a Model with
// gtk_popover_menu_new_from_model().
//
//    <section>
//      <attribute name="display-hint">horizontal-buttons</attribute>
//      <item>
//        <attribute name="label">Cut</attribute>
//        <attribute name="action">app.cut</attribute>
//        <attribute name="verb-icon">edit-cut-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Copy</attribute>
//        <attribute name="action">app.copy</attribute>
//        <attribute name="verb-icon">edit-copy-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Paste</attribute>
//        <attribute name="action">app.paste</attribute>
//        <attribute name="verb-icon">edit-paste-symbolic</attribute>
//      </item>
//    </section>
//
// CSS nodes
//
//    popover[.menu]
//    ├── arrow
//    ╰── contents.background
//        ╰── <child>
//
// The contents child node always gets the .background style class and the
// popover itself gets the .menu style class if the popover is menu-like (i.e.
// PopoverMenu).
//
// Particular uses of GtkPopover, such as touch selection popups or magnifiers
// in Entry or TextView get style classes like .touch-selection or .magnifier to
// differentiate from plain popovers.
//
// When styling a popover directly, the popover node should usually not have any
// background.
//
// Note that, in order to accomplish appropriate arrow visuals, Popover uses
// custom drawing for the arrow node. This makes it possible for the arrow to
// change its shape dynamically, but it also limits the possibilities of styling
// it using CSS. In particular, the arrow gets drawn over the content node's
// border so they look like one shape, which means that the border-width of the
// content node and the arrow node should be the same. The arrow also does not
// support any border shape other than solid, no border-radius, only one border
// width (border-bottom-width is used) and no box-shadow.
type Popover interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Native
	ShortcutManager

	// Autohide returns whether the popover is modal.
	//
	// See gtk_popover_set_autohide() for the implications of this.
	Autohide() bool
	// CascadePopdown returns whether the popover will close after a modal child
	// is closed.
	CascadePopdown() bool
	// Child gets the child widget of @popover.
	Child() Widget
	// HasArrow gets whether this popover is showing an arrow pointing at the
	// widget that it is relative to.
	HasArrow() bool
	// MnemonicsVisible gets the value of the Popover:mnemonics-visible
	// property.
	MnemonicsVisible() bool
	// Offset gets the offset previous set with gtk_popover_set_offset().
	Offset() (xOffset int, yOffset int)
	// PointingTo: if a rectangle to point to has been set, this function will
	// return true and fill in @rect with such rectangle, otherwise it will
	// return false and fill in @rect with the attached widget coordinates.
	PointingTo() (rect gdk.Rectangle, ok bool)
	// Position returns the preferred position of @popover.
	Position() PositionType
	// Popdown pops @popover down.This is different than a gtk_widget_hide()
	// call in that it shows the popover with a transition. If you want to hide
	// the popover without a transition, use gtk_widget_hide().
	Popdown()
	// Popup pops @popover up. This is different than a gtk_widget_show() call
	// in that it shows the popover with a transition. If you want to show the
	// popover without a transition, use gtk_widget_show().
	Popup()
	// Present presents the popover to the user.
	Present()
	// SetAutohide sets whether @popover is modal.
	//
	// A modal popover will grab the keyboard focus on it when being displayed.
	// Clicking outside the popover area or pressing Esc will dismiss the
	// popover.
	//
	// Called this function on an already showing popup with a new autohide
	// value different from the current one, will cause the popup to be hidden.
	SetAutohide(autohide bool)
	// SetCascadePopdown: if @cascade_popdown is UE, the popover will be closed
	// when a child modal popover is closed. If LSE, @popover will stay visible.
	SetCascadePopdown(cascadePopdown bool)
	// SetChild sets the child widget of @popover.
	SetChild(child Widget)
	// SetDefaultWidget: the default widget is the widget that’s activated when
	// the user presses Enter in a dialog (for example). This function sets or
	// unsets the default widget for a Popover.
	SetDefaultWidget(widget Widget)
	// SetHasArrow sets whether this popover should draw an arrow pointing at
	// the widget it is relative to.
	SetHasArrow(hasArrow bool)
	// SetMnemonicsVisible sets the Popover:mnemonics-visible property.
	SetMnemonicsVisible(mnemonicsVisible bool)
	// SetOffset sets the offset to use when calculating the position of the
	// popover.
	//
	// These values are used when preparing the PopupLayout for positioning the
	// popover.
	SetOffset(xOffset int, yOffset int)
	// SetPointingTo sets the rectangle that @popover will point to, in the
	// coordinate space of the @popover parent.
	SetPointingTo(rect *gdk.Rectangle)
	// SetPosition sets the preferred position for @popover to appear. If the
	// @popover is currently visible, it will be immediately updated.
	//
	// This preference will be respected where possible, although on lack of
	// space (eg. if close to the window edges), the Popover may choose to
	// appear on the opposite side
	SetPosition(position PositionType)
}

// popover implements the Popover interface.
type popover struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Native
	ShortcutManager
}

var _ Popover = (*popover)(nil)

// WrapPopover wraps a GObject to the right type. It is
// primarily used internally.
func WrapPopover(obj *externglib.Object) Popover {
	return Popover{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Native:           WrapNative(obj),
		ShortcutManager:  WrapShortcutManager(obj),
	}
}

func marshalPopover(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPopover(obj), nil
}

// NewPopover constructs a class Popover.
func NewPopover() Popover {
	var cret C.GtkPopover
	var goret1 Popover

	cret = C.gtk_popover_new()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Popover)

	return goret1
}

// Autohide returns whether the popover is modal.
//
// See gtk_popover_set_autohide() for the implications of this.
func (p popover) Autohide() bool {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_popover_get_autohide(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// CascadePopdown returns whether the popover will close after a modal child
// is closed.
func (p popover) CascadePopdown() bool {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_popover_get_cascade_popdown(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Child gets the child widget of @popover.
func (p popover) Child() Widget {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret *C.GtkWidget
	var goret1 Widget

	cret = C.gtk_popover_get_child(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret1
}

// HasArrow gets whether this popover is showing an arrow pointing at the
// widget that it is relative to.
func (p popover) HasArrow() bool {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_popover_get_has_arrow(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// MnemonicsVisible gets the value of the Popover:mnemonics-visible
// property.
func (p popover) MnemonicsVisible() bool {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_popover_get_mnemonics_visible(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Offset gets the offset previous set with gtk_popover_set_offset().
func (p popover) Offset() (xOffset int, yOffset int) {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var arg1 *C.int
	var ret1 int
	var arg2 *C.int
	var ret2 int

	C.gtk_popover_get_offset(arg0, &arg1, &arg2)

	ret1 = *C.int(arg1)
	ret2 = *C.int(arg2)

	return ret1, ret2
}

// PointingTo: if a rectangle to point to has been set, this function will
// return true and fill in @rect with such rectangle, otherwise it will
// return false and fill in @rect with the attached widget coordinates.
func (p popover) PointingTo() (rect gdk.Rectangle, ok bool) {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var arg1 *C.GdkRectangle
	var ret1 *gdk.Rectangle
	var cret C.gboolean
	var goret2 bool

	cret = C.gtk_popover_get_pointing_to(arg0, &arg1)

	ret1 = gdk.WrapRectangle(unsafe.Pointer(arg1))
	goret2 = C.bool(cret) != C.false

	return ret1, goret2
}

// Position returns the preferred position of @popover.
func (p popover) Position() PositionType {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret C.GtkPositionType
	var goret1 PositionType

	cret = C.gtk_popover_get_position(arg0)

	goret1 = PositionType(cret)

	return goret1
}

// Popdown pops @popover down.This is different than a gtk_widget_hide()
// call in that it shows the popover with a transition. If you want to hide
// the popover without a transition, use gtk_widget_hide().
func (p popover) Popdown() {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popdown(arg0)
}

// Popup pops @popover up. This is different than a gtk_widget_show() call
// in that it shows the popover with a transition. If you want to show the
// popover without a transition, use gtk_widget_show().
func (p popover) Popup() {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popup(arg0)
}

// Present presents the popover to the user.
func (p popover) Present() {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_present(arg0)
}

// SetAutohide sets whether @popover is modal.
//
// A modal popover will grab the keyboard focus on it when being displayed.
// Clicking outside the popover area or pressing Esc will dismiss the
// popover.
//
// Called this function on an already showing popup with a new autohide
// value different from the current one, will cause the popup to be hidden.
func (p popover) SetAutohide(autohide bool) {
	var arg0 *C.GtkPopover
	var arg1 C.gboolean

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if autohide {
		arg1 = C.gboolean(1)
	}

	C.gtk_popover_set_autohide(arg0, autohide)
}

// SetCascadePopdown: if @cascade_popdown is UE, the popover will be closed
// when a child modal popover is closed. If LSE, @popover will stay visible.
func (p popover) SetCascadePopdown(cascadePopdown bool) {
	var arg0 *C.GtkPopover
	var arg1 C.gboolean

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if cascadePopdown {
		arg1 = C.gboolean(1)
	}

	C.gtk_popover_set_cascade_popdown(arg0, cascadePopdown)
}

// SetChild sets the child widget of @popover.
func (p popover) SetChild(child Widget) {
	var arg0 *C.GtkPopover
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_popover_set_child(arg0, child)
}

// SetDefaultWidget: the default widget is the widget that’s activated when
// the user presses Enter in a dialog (for example). This function sets or
// unsets the default widget for a Popover.
func (p popover) SetDefaultWidget(widget Widget) {
	var arg0 *C.GtkPopover
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_popover_set_default_widget(arg0, widget)
}

// SetHasArrow sets whether this popover should draw an arrow pointing at
// the widget it is relative to.
func (p popover) SetHasArrow(hasArrow bool) {
	var arg0 *C.GtkPopover
	var arg1 C.gboolean

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if hasArrow {
		arg1 = C.gboolean(1)
	}

	C.gtk_popover_set_has_arrow(arg0, hasArrow)
}

// SetMnemonicsVisible sets the Popover:mnemonics-visible property.
func (p popover) SetMnemonicsVisible(mnemonicsVisible bool) {
	var arg0 *C.GtkPopover
	var arg1 C.gboolean

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if mnemonicsVisible {
		arg1 = C.gboolean(1)
	}

	C.gtk_popover_set_mnemonics_visible(arg0, mnemonicsVisible)
}

// SetOffset sets the offset to use when calculating the position of the
// popover.
//
// These values are used when preparing the PopupLayout for positioning the
// popover.
func (p popover) SetOffset(xOffset int, yOffset int) {
	var arg0 *C.GtkPopover
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = C.int(xOffset)
	arg2 = C.int(yOffset)

	C.gtk_popover_set_offset(arg0, xOffset, yOffset)
}

// SetPointingTo sets the rectangle that @popover will point to, in the
// coordinate space of the @popover parent.
func (p popover) SetPointingTo(rect *gdk.Rectangle) {
	var arg0 *C.GtkPopover
	var arg1 *C.GdkRectangle

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect.Native()))

	C.gtk_popover_set_pointing_to(arg0, rect)
}

// SetPosition sets the preferred position for @popover to appear. If the
// @popover is currently visible, it will be immediately updated.
//
// This preference will be respected where possible, although on lack of
// space (eg. if close to the window edges), the Popover may choose to
// appear on the opposite side
func (p popover) SetPosition(position PositionType) {
	var arg0 *C.GtkPopover
	var arg1 C.GtkPositionType

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = (C.GtkPositionType)(position)

	C.gtk_popover_set_position(arg0, position)
}
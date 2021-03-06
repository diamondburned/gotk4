// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_CellRendererCombo_ConnectChanged(gpointer, gchar*, GtkTreeIter*, guintptr);
import "C"

// GType values.
var (
	GTypeCellRendererCombo = coreglib.Type(C.gtk_cell_renderer_combo_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellRendererCombo, F: marshalCellRendererCombo},
	})
}

// CellRendererCombo renders a combobox in a cell
//
// CellRendererCombo renders text in a cell like CellRendererText from which it
// is derived. But while CellRendererText offers a simple entry to edit the
// text, CellRendererCombo offers a ComboBox widget to edit the text. The values
// to display in the combo box are taken from the tree model specified in the
// CellRendererCombo:model property.
//
// The combo cell renderer takes care of adding a text cell renderer to the
// combo box and sets it to display the column specified by its
// CellRendererCombo:text-column property. Further properties of the combo box
// can be set in a handler for the CellRenderer::editing-started signal.
type CellRendererCombo struct {
	_ [0]func() // equal guard
	CellRendererText
}

var (
	_ CellRendererer = (*CellRendererCombo)(nil)
)

func wrapCellRendererCombo(obj *coreglib.Object) *CellRendererCombo {
	return &CellRendererCombo{
		CellRendererText: CellRendererText{
			CellRenderer: CellRenderer{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalCellRendererCombo(p uintptr) (interface{}, error) {
	return wrapCellRendererCombo(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_CellRendererCombo_ConnectChanged
func _gotk4_gtk4_CellRendererCombo_ConnectChanged(arg0 C.gpointer, arg1 *C.gchar, arg2 *C.GtkTreeIter, arg3 C.guintptr) {
	var f func(pathString string, newIter *TreeIter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(pathString string, newIter *TreeIter))
	}

	var _pathString string // out
	var _newIter *TreeIter // out

	_pathString = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_newIter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	f(_pathString, _newIter)
}

// ConnectChanged: this signal is emitted each time after the user selected an
// item in the combo box, either by using the mouse or the arrow keys. Contrary
// to GtkComboBox, GtkCellRendererCombo::changed is not emitted for changes made
// to a selected item in the entry. The argument new_iter corresponds to the
// newly selected item in the combo box and it is relative to the GtkTreeModel
// set via the model property on GtkCellRendererCombo.
//
// Note that as soon as you change the model displayed in the tree view, the
// tree view will immediately cease the editing operating. This means that you
// most probably want to refrain from changing the model until the combo cell
// renderer emits the edited or editing_canceled signal.
func (v *CellRendererCombo) ConnectChanged(f func(pathString string, newIter *TreeIter)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "changed", false, unsafe.Pointer(C._gotk4_gtk4_CellRendererCombo_ConnectChanged), f)
}

// NewCellRendererCombo creates a new CellRendererCombo. Adjust how text is
// drawn using object properties. Object properties can be set globally (with
// g_object_set()). Also, with TreeViewColumn, you can bind a property to a
// value in a TreeModel. For example, you can bind the ???text??? property on the
// cell renderer to a string value in the model, thus rendering a different
// string in each row of the TreeView.
//
// The function returns the following values:
//
//    - cellRendererCombo: new cell renderer.
//
func NewCellRendererCombo() *CellRendererCombo {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_combo_new()

	var _cellRendererCombo *CellRendererCombo // out

	_cellRendererCombo = wrapCellRendererCombo(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererCombo
}

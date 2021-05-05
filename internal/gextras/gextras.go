// Package gextras contains supplemental types to gotk3.
package gextras

import (
	"github.com/gotk3/gotk3/glib"
)

// Objector is an interface that describes partially the glib.Object type.
type Objector interface {
	Connect(string, interface{}) glib.SignalHandle
	ConnectAfter(string, interface{}) glib.SignalHandle
	Emit(string, ...interface{}) (interface{}, error)
	HandlerBlock(glib.SignalHandle)
	HandlerDisconnect(glib.SignalHandle)
	HandlerUnblock(glib.SignalHandle)
	IsA(glib.Type) bool
	Native() uintptr
	GetProperty(string) (interface{}, error)
	SetProperty(string, interface{}) error
	StopEmission(string)
	TypeFromInstance() glib.Type
}

var _ Objector = (*glib.Object)(nil)

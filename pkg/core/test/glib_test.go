package test

import (
	"testing"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func TestObjectData(t *testing.T) {
	gtk.Init()

	label := gtk.NewLabel("label")

	label.SetObjectData("foo", "bar")

	if label.ObjectData("foo") != "bar" {
		t.Fatal("returned data did not match expected data")
	}

	label.StealObjectData("foo")
}

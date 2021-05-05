package girgen

import "github.com/diamondburned/gotk4/gir"

// Generator is the current generator state.
type Generator struct {
	Repos gir.Repositories
}

func NewGenerator(repos gir.Repositories) *Generator {

}

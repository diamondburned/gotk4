package main

import (
	"log"
	"os"

	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

func main() {
	if os.Getenv("GOTK4_RUNTIME_LINK") == "1" {
		// TODO: remove this when RuntimeLinkMode is working.
		log.Println("warning: GOTK4_RUNTIME_LINK is set to 1")
		girgen.DefaultLinkMode = types.RuntimeLinkMode
	}

	genmain.Run(gendata.Main)
}

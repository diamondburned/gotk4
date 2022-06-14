package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/pkgconfig"
)

func main() {
	flag.Parse()

	name := flag.Arg(0)
	if name == "" {
		log.Fatalln("missing name; usage: gir_namespaces <pkg-config name>")
	}

	files, err := pkgconfig.FindGIRFiles(name)
	if err != nil {
		log.Fatalln("failed to get gir files for", name)
	}

	for _, file := range files {
		repo, err := gir.ParseRepository(file)
		if err != nil {
			log.Fatalln("failed to parse repository file", file)
		}

		fmt.Println(file, "v"+repo.Version)
		for _, namespace := range repo.Namespaces {
			fmt.Println(" ", namespace.Name, "v"+namespace.Version)
		}
	}
}

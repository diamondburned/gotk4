package main

import (
	"flag"
	"log"
	"os"

	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/genutil"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
)

const module = "github.com/diamondburned/gotk4/pkg"

var (
	output  string
	verbose bool
	listPkg bool
)

func init() {
	flag.StringVar(&output, "o", "", "output directory to mkdir in")
	flag.BoolVar(&verbose, "v", verbose, "log verbosely (debug mode)")
	flag.BoolVar(&listPkg, "l", listPkg, "only list packages and exit")
	flag.Parse()

	if !listPkg && output == "" {
		log.Fatalln("Missing -o output directory.")
	}

	if !verbose {
		verbose = os.Getenv("GIR_VERBOSE") == "1"
	}
}

func main() {
	repos := genutil.MustLoadPackages(gendata.Packages)
	genutil.PrintAddedPkgs(repos)

	if listPkg {
		return
	}

	gen := girgen.NewGenerator(repos, genutil.ModulePath(module, gendata.ImportOverrides))
	gen.Logger = log.New(os.Stderr, "girgen: ", log.Lmsgprefix)
	gen.ApplyPreprocessors(gendata.Preprocessors)
	gen.AddFilters(gendata.Filters)
	gen.AddProcessConverters(gendata.ConversionProcessors)

	if verbose {
		gen.LogLevel = logger.Debug
	}

	if err := genutil.CleanDirectory(output, gendata.PkgExceptions); err != nil {
		log.Fatalln("failed to clean output directory:", err)
	}

	if errors := genutil.GenerateAll(gen, output, gendata.GenerateExceptions); len(errors) > 0 {
		for _, err := range errors {
			log.Println("generation error:", err)
		}
		os.Exit(1)
	}

	if err := genutil.AppendGoFiles(output, gendata.Appends); err != nil {
		log.Fatalln("failed to append files post-generation:", err)
	}

	finalFiles := [][]string{gendata.PkgExceptions, gendata.PkgGenerated}
	if err := genutil.EnsureDirectory(output, finalFiles...); err != nil {
		log.Fatalln("error verifying generation:", err)
	}
}

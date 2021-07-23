package generators

import (
	"strconv"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var constantTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	const {{ .Name }} = {{ .Value }}
`)

type constantData struct {
	*gir.Constant
	Value string
}

func GenerateConstant(gen FileGeneratorWriter, constant *gir.Constant) bool {
	goType := types.GIRBuiltinGo(constant.Type.Name)
	if goType == "" {
		gen.Logln(logger.Debug, "unknown constant type", constant.Type.Name)
		return false
	}

	if !types.Filter(gen, constant.Name, constant.CType) {
		gen.Logln(logger.Debug, "filtered constant", constant.CType)
		return false
	}

	value := constant.Value
	if goType == "string" {
		value = strconv.Quote(value)
	}

	writer := FileWriterFromType(gen, constant)
	writer.Pen().WriteTmpl(constantTmpl, &constantData{
		Constant: constant,
		Value:    value,
	})

	return true
}

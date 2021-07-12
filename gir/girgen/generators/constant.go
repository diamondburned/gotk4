package generators

import (
	"strconv"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
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
		return false
	}

	if !types.Filter(gen, constant.Name, constant.CType) {
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

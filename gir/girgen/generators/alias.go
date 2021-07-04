package generators

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

var aliasTmpl = gotmpl.NewGoTemplate(`
	{{ $name := (PascalToGo .Name) }}
	{{ GoDoc . 0 }}
	type {{ $name }} = {{ .GoType }}
`)

type aliasData struct {
	*gir.Alias
	GoType string
}

// GenerateAlias generates an alias declaration into the given file generator.
// If the generation fails or is ignored, then false is returned.
func GenerateAlias(gen FileGeneratorWriter, alias *gir.Alias) bool {
	if !alias.IsIntrospectable() || types.Filter(gen, alias.Name, alias.CType) {
		return false
	}

	resolved := types.Resolve(gen, alias.Type)
	if resolved == nil {
		return false
	}

	goType := resolved.PublicType(resolved.NeedsNamespace(gen.Namespace()))
	if goType == "" {
		// Use the C type directly if we can't find the Go equivalent.
		goType = "C." + alias.Type.CType
	}

	writer := FileWriterFromType(gen, alias)
	writer.Header().ImportPubl(resolved)
	writer.Pen().WriteTmpl(aliasTmpl, aliasData{
		Alias:  alias,
		GoType: goType,
	})

	return true
}

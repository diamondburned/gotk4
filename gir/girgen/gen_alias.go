package girgen

import "github.com/diamondburned/gotk4/gir"

var aliasTmpl = newGoTemplate(`
	{{ $name := (PascalToGo .Name) }}
	{{ GoDoc .Doc 0 $name }}
	type {{ $name }} {{ .GoType }}
`)

type aliasData struct {
	gir.Alias
	GoType string
}

func (ng *NamespaceGenerator) generateAliases() {
	for _, alias := range ng.current.Namespace.Aliases {
		resolved := ng.resolveType(alias.Type)
		if resolved == nil {
			continue
		}
		if resolved.GoType == "" {
			// TODO: fix this.
			resolved.GoType = "struct{}"
		}

		ng.pen.BlockTmpl(aliasTmpl, aliasData{
			Alias:  alias,
			GoType: resolved.GoType,
		})
	}
}

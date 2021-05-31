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
		goType, ok := ng.ResolveToGoType(alias.Type, true)
		if !ok {
			continue
		}

		if ng.mustIgnore(alias.Name, alias.CType) {
			continue
		}

		if goType == "" {
			// TODO: fix this.
			goType = "struct{}"
		}

		ng.pen.BlockTmpl(aliasTmpl, aliasData{
			Alias:  alias,
			GoType: goType,
		})
	}
}

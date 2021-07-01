package generators

import "github.com/diamondburned/gotk4/gir"

var aliasTmpl = newGoTemplate(`
	{{ $name := (PascalToGo .Name) }}
	{{ GoDoc .Doc 0 $name }}
	type {{ $name }} = {{ .GoType }}
`)

type aliasData struct {
	gir.Alias
	GoType string
}

func (ng *NamespaceGenerator) generateAliases() {
	for _, alias := range ng.current.Namespace.Aliases {
		if !alias.IsIntrospectable() {
			continue
		}
		if ng.mustIgnore(&alias.Name, &alias.CType) {
			continue
		}

		resolved := ng.ResolveType(alias.Type)
		if resolved == nil {
			continue
		}

		needsNamespace := resolved.NeedsNamespace(ng.Namespace())

		goType := resolved.PublicType(needsNamespace)
		if goType == "" {
			ng.Logln(LogSkip, "alias", alias.Name, "is opaque type")
			continue
		}

		ng.pen.WriteTmpl(aliasTmpl, aliasData{
			Alias:  alias,
			GoType: goType,
		})
	}
}

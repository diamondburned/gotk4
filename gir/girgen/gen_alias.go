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
		if !alias.IsIntrospectable() {
			continue
		}
		if ng.mustIgnore(&alias.Name, &alias.CType) {
			continue
		}

		fg := ng.FileFromSource(alias.DocElements)

		goType, ok := GoType(fg, alias.Type, true)
		if !ok {
			continue
		}

		if goType == "" {
			fg.Logln(LogSkip, "alias", alias.Name, "is opaque type")
			continue
		}

		fg.pen.WriteTmpl(aliasTmpl, aliasData{
			Alias:  alias,
			GoType: goType,
		})
	}
}

package girgen

// var classTmpl = newGoTemplate(`
// 	{{ $type := (PascalToGo .Name) }}

// 	{{ GoDoc .Doc 0 $type }}
// 	type {{ $type }} struct {

// 	}
// `)

// type classGenerator struct {
// 	gir.Class
// 	ParentGo string

// 	ng *NamespaceGenerator
// }

// func (rg classGenerator) NeedsNative() bool {
// 	for _, field := range rg.Fields {
// 		if field.Private {
// 			return true
// 		}
// 	}

// 	return len(rg.Fields) == 0
// }

// func (rg *classGenerator) Use(class gir.Class) bool {
// 	// TODO: resolve parent type
// 	// TODO: recursively resolve parent type until end of chain
// }

// func (rg classGenerator) Field(field gir.Field) string {
// 	if field.Private {
// 		return ""
// 	}

// 	typ, ok := rg.ng.ResolveAnyType(field.AnyType)
// 	if !ok {
// 		return ""
// 	}

// 	name := SnakeToGo(true, field.Name)

// 	return GoDoc(field.Doc, 1, name) + "\n" + name + " " + typ
// }

// func (ng *NamespaceGenerator) generateClasss() {
// 	_ = gir.Field{}

// 	for _, class := range ng.current.Namespace.Classes {
// 		if ignoreClass(class) {
// 			continue
// 		}

// 		ng.pen.BlockTmpl(classTmpl, classGenerator{
// 			Class: class,
// 			ng:    ng,
// 		})
// 	}
// }

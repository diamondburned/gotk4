package girgen

// Go to C type conversions.

// func (ng *NamespaceGenerator) GoCConverter(value, target string, any gir.AnyType) string {

// }

// func (ng *NamespaceGenerator) gocArrayConverter(value, target string, array gir.Array) string {

// }

// func (ng *NamespaceGenerator) gocTypeConverter(value, target string, typ gir.Type) string {

// }

// func (ng *NamespaceGenerator) _gocTypeConverter(value, target string, typ gir.Type, create bool) string {
// 	if prim, ok := girPrimitiveGo[typ.Name]; ok {
// 		switch prim {
// 		case "string":
// 			p := pen.NewPiece()
// 			p.Linef(directCallOrCreate(value, target, "C.CString", create))
// 			p.Linef("defer C.free(unsafe.Pointer(%s))", value)
// 			return p.String()
// 		case "bool":
// 			ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
// 			return directCallOrCreate(value, target, "gextras.Cbool", create)
// 		default:
// 			return directCallOrCreate(value, target, "C."+typ.CType, create)
// 		}
// 	}

// 	switch typ.Name {
// 	case "gpointer":

// 	}
// }

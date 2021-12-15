// Package gocode provides utility functions to manipulate existing Go code.
package gocode

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"strconv"
	"strings"
)

// CoalesceTail coalesces certain parameters with the same type to be shorter.
func CoalesceTail(tail string) string {
	expr, err := parser.ParseExpr("func" + tail)
	if err != nil {
		log.Panicf("CoalesceTail: invalid tail %q: %v", tail, err)
	}

	fn := expr.(*ast.FuncType)
	if len(fn.Params.List) == 0 {
		return tail // no parameters
	}

	coalesceList := make([]*ast.Field, 0, len(fn.Params.List))
	last := fn.Params.List[0]

	for _, param := range fn.Params.List[1:] {
		if exprEq(param.Type, last.Type) {
			last.Names = append(last.Names, param.Names...)
			continue
		}
		coalesceList = append(coalesceList, last)
		last = param
	}
	coalesceList = append(coalesceList, last)

	fn.Params.List = coalesceList

	return strings.TrimPrefix(printNode(expr), "func")
}

func printNode(node any) string {
	var out strings.Builder
	if err := printer.Fprint(&out, token.NewFileSet(), node); err != nil {
		log.Panicf("CoalesceTail: cannot fprint source: %v", err)
	}
	return out.String()
}

// TODO: figure out something not as dumb.
func exprEq(expr1, expr2 ast.Expr) bool {
	return exprString(expr1) == exprString(expr2)
}

func exprString(v interface{}) string {
	switch v := v.(type) {
	case *ast.StarExpr:
		return "*" + exprString(v.X)
	case *ast.Ident:
		return v.String()
	case *ast.SelectorExpr:
		return exprString(v.Sel) + "." + exprString(v.X)
	case *ast.ArrayType:
		return "[]" + exprString(v.Elt)
	default:
		// Slow path.
		var out strings.Builder
		printer.Fprint(&out, token.NewFileSet(), v)
		return out.String()
	}
}

func splitParam(param string) (name, typ string) {
	param = strings.TrimSpace(param)

	parts := strings.SplitN(param, " ", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}

	panic("splitParam: invalid non-two-part param " + strconv.Quote(param))
}

// FormatReturn formats the return snippet from the given list of types.
func FormatReturn(parts []string) string {
	if len(parts) == 0 {
		return ""
	}

	types := make([]string, len(parts))

	for i, part := range parts {
		types[i] = extractTypeFromPair(part)
	}

	for i := range parts {
		for j := range parts {
			if i == j {
				continue
			}

			if types[i] == types[j] {
				goto dupeType
			}
		}
	}

	// No duplicate types, so use types only.
	parts = types

dupeType:
	if len(parts) == 1 && !strings.Contains(parts[0], " ") {
		return parts[0]
	}

	return "(" + strings.Join(parts, ", ") + ")"
}

// extractTypeFromPair returns the second word (which is the type) from the
// name-type pair.
func extractTypeFromPair(namePair string) string {
	return namePair[strings.IndexByte(namePair, ' ')+1:]
}

// ExtractDefer extracts only snippets of code that will be deferred. The order
// of execution is flipped to reflect the LIFO behavior. The function will only
// extract top-level defer statements.
func ExtractDefer(x string) string {
	expr, err := parser.ParseExpr("func() {" + x + "}")
	if err != nil {
		log.Println("invalid expression\n" + x)
		log.Panicln("ParseExpr:", err)
	}

	fn := expr.(*ast.FuncLit)
	newLit := make([]ast.Stmt, 0, len(fn.Body.List))

	for _, stmt := range fn.Body.List {
		defers, ok := stmt.(*ast.DeferStmt)
		if !ok {
			continue
		}

		newLit = append(newLit, &ast.ExprStmt{X: defers.Call})
	}

	// Reverse.
	for left, right := 0, len(newLit)-1; left < right; left, right = left+1, right-1 {
		newLit[left], newLit[right] = newLit[right], newLit[left]
	}

	fn.Body.List = newLit

	// ast.Print(token.NewFileSet(), expr)

	code := printNode(fn.Body.List)
	return code
}

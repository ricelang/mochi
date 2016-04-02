package compiler

import (
	"fmt"
	"github.com/mochi-lang/mochi/parser"
	"go/ast"
	"go/token"
)

func isNsDecl(node parser.Node) bool {
	if node.Type() != parser.NodeCall {
		return false
	}

	call := node.(*parser.CallNode)
	if call.Callee.(*parser.IdentNode).Ident != "ns" {
		return false
	}

	if len(call.Args) < 1 {
		return false
	}

	return true
}

func getNamespace(node *parser.CallNode) *ast.Ident {
	return &ast.Ident{
		Name: node.Args[0].(*parser.IdentNode).Ident,
	}
}

func getImports(node *parser.CallNode) []ast.Spec {
	imports := node.Args[1:]

	specs := make([]ast.Spec, len(imports))
	for i, a := range imports {
		fmt.Printf("N= %v\n", a)

		specs[i] = &ast.ImportSpec{
			Path: &ast.BasicLit{Value: a.(*parser.StringNode).Value},
			Name: &ast.Ident{Name: a.(*parser.StringNode).Value},
		}
	}

	return specs
}

func importsToDecl(specs []ast.Spec) ast.Decl {
	s := ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: specs,
	}
	s.Lparen = 1
	return &s
}

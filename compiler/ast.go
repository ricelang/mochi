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
			//@TODO use Name to implement name import
			Name: nil,
		}
	}

	return specs
}

func importsToDecl(specs []ast.Spec) ast.Decl {
	s := ast.GenDecl{
		Doc:   genComment(),
		Tok:   token.IMPORT,
		Specs: specs,
	}
	// https://godoc.org/go/ast#GenDecl
	// A valid Lparen position (Lparen.Line > 0) indicates a parenthesized declaration.
	s.Lparen = 1
	return &s
}

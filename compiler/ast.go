package compiler

import (
	"fmt"
	"github.com/mochi-lang/mochi/compiler/helper"
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
		Name: "main",
		// @TODO When we support build a whole dir, we will use below code
		//Name: node.Args[0].(*parser.IdentNode).Ident,
	}
}

func getImports(node *parser.CallNode) []ast.Spec {
	imports := node.Args[1:]
	imports = append(imports, &parser.StringNode{NodeType: parser.NodeString, Value: "\"github.com/mochi-lang/mochi/mochi/core\""})

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

// importsToDecl turns Lisp import into Go AST
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

// topLevelNodeToDecl converts an AST node to Go ast.Decl
func topLevelNodeToDecl(node *parser.CallNode) ast.Decl {
	if node.Callee.(*parser.IdentNode).Ident != "defn" {
		panic("Top level node has to be defn")
	}

	switch node.Callee.(*parser.IdentNode).Ident {
	case "defn":
		decl := ast.FuncDecl{
			Name: &ast.Ident{
				NamePos: 2,
				Name:    node.Args[0].(*parser.IdentNode).Ident,
			},
		}
		fmt.Printf("SHOW ARGS: %+v\n", node.Args[1:])

		// @TODO Fix this to support signature properly
		params := &ast.FieldList{
			Opening: 1,
			Closing: 3,
		}
		params.List = make([]*ast.Field, 1)
		params.List[0] = &ast.Field{
			Names: []*ast.Ident{
				&ast.Ident{
					Name: "lol",
				},
			},
			Type: &ast.Ident{
				Name: "interface{}",
			},
		}

		decl.Type = &ast.FuncType{
			Func:    1,
			Params:  params,
			Results: nil,
		}

		decl.Body = nodeFnBody(node.Args[1:])

		return &decl
	default:
		// The rest is normal function call probably
		//return nodeFnCall(node)
		fmt.Println("got nil %+v", node)
		return nil
	}
}

func nodeFnBody(nodes []parser.Node) *ast.BlockStmt {
	stmt := ast.BlockStmt{
		List: make([]ast.Stmt, len(nodes)),
	}

	for i, node := range nodes {
		switch node.Type() {
		case parser.NodeCall:
			fmt.Printf("FnBody %+v %s", node, node.Type())

			exprstmt := &ast.ExprStmt{}
			exprstmt.X = nodeFnCall(node.(*parser.CallNode))
			stmt.List[i] = exprstmt
		default:
			panic(fmt.Sprintf("Doesn't support node: %v here", node))
		}
	}

	return &stmt
}

// Convert a Lisp node fn call to AST
func nodeFnCall(node *parser.CallNode) ast.Expr {
	fmt.Printf("nodeFnCall node=%+v ident=%v args=%v", node, node.Callee.(*parser.IdentNode).Ident, node.Args)
	stmt := &ast.CallExpr{}
	switch node.Callee.(*parser.IdentNode).Ident {
	case "let":
		stmt := &ast.AssignStmt {
			Lhs: []ast.Expr{
				&ast.Ident {
					Name: node.Args[0].(*parser.IdentNode).Ident,
				},
			}
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.Ident {
					Name: node.Args[0].(*parser.IdentNode).Ident,
				},
			}

		}
		return stmt
	case "if":
		stmt.Fun = &ast.Ident{
			Name: "core.If",
		}
		stmt.Args = make([]ast.Expr, len(node.Args))
		switch len(node.Args) {
		case 1:
			panic("If requires >=2 arguments")
		}
	default:
		stmt = &ast.CallExpr{
			Fun: &ast.Ident{
				Name: helper.LispFnToGoName(node.Callee.(*parser.IdentNode).Ident),
			},
			Args: make([]ast.Expr, len(node.Args)),
		}
	}

	for i, a := range node.Args {
		stmt.Args[i] = nodeToStmt(a)
	}

	return stmt
}

func nodeToStmt(node parser.Node) ast.Expr {
	switch node.Type() {
	case parser.NodeCall:
		return nodeFnCall(node.(*parser.CallNode))
	case parser.NodeString:
		return &ast.BasicLit{
			Kind:  token.STRING,
			Value: node.(*parser.StringNode).Value,
		}
	case parser.NodeIdent:
		return &ast.BasicLit{
			Kind:  token.STRING,
			Value: node.(*parser.IdentNode).Ident,
		}
	}
	panic(fmt.Sprintf("Unknow NodeType %v: %v", node.Type(), node))
}

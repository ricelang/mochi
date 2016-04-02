package compiler

import (
	"fmt"
	"github.com/mochi-lang/mochi/parser"
	"go/ast"
	//"go/token"
)

func Compile(tree []parser.Node) *ast.File {
	for i, v := range tree {
		fmt.Printf("%d v= %v\n", i, v)
		if v.Type() == parser.NodeCall {
			for _, a := range v.(*parser.CallNode).Args {
				fmt.Printf("func: %v args: %v", v.(*parser.CallNode).Callee, a)
			}
		}
	}

	fmt.Println("Will compile")
	f := &ast.File{Name: ast.NewIdent("main")}
	decls := make([]ast.Decl, 0, len(tree))

	if len(tree) < 1 {
		return f
	}

	// you can only have (ns ...) as the first form
	if isNsDecl(tree[0]) {
		name := getNamespace(tree[0].(*parser.CallNode))
		imports := getImports(tree[0].(*parser.CallNode))

		f.Name = name
		//f.Comments = genComment()

		//decls = append(decls, genComment())

		if imports != nil {
			decls = append(decls, importsToDecl(imports))
			//f.Imports = imports
		}

		tree = tree[1:]
	}

	//decls = append(decls, generateDecls(tree)...)

	f.Decls = decls
	return f
}

package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mochi-lang/mochi/generator"
	"github.com/mochi-lang/mochi/parser"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func Run() {
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		line, _, _ := r.ReadLine()
		p := parser.ParseFromString("<REPL>", string(line)+"\n")
		fmt.Println(p)

		// a := generator.GenerateAST(p)
		a := generator.EvalExprs(p)
		fset := token.NewFileSet()
		ast.Print(fset, a)

		var buf bytes.Buffer
		printer.Fprint(&buf, fset, a)
		fmt.Printf("%s\n", buf.String())
	}
}

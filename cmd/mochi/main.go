package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mochi-lang/mochi/generator"
	"github.com/mochi-lang/mochi/parser"
	"go/ast"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
)

// Mochi version
//	-ldflags "-X main.Version=`git describe --always --tags`"
var VERSION string

const (
	Lang = "Mochi"
)

func args(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	p := parser.ParseFromString(filename, string(b)+"\n")

	a := generator.GenerateAST(p)

	fset := token.NewFileSet()

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, a)
	fmt.Printf("%s\n", buf.String())
}

func main() {
	fmt.Printf("\033[48;5;95;38;5;214m%s. Rev %s\033[0m\n", Lang, VERSION)

	if len(os.Args) > 1 {
		args(os.Args[1])
		return
	}

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

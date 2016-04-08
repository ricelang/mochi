package interpreter

import (
	//"bytes"
	//"fmt"
	"github.com/mochi-lang/mochi/compiler"
	//"github.com/mochi-lang/mochi/generator"
	"github.com/mochi-lang/mochi/parser"
	//"go/printer"
	//"go/token"
	"io/ioutil"
)

func Run(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	p := parser.ParseFromString(filename, string(b)+"\n")

	v := compiler.Compile(p)
	compiler.PrintOut(v)

	//a := generator.GenerateAST(p)

	//fset := token.NewFileSet()

	//var buf bytes.Buffer
	//printer.Fprint(&buf, fset, a)
	//fmt.Printf("%s\n", buf.String())
}

package compiler

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
)

func PrintOut(v *ast.File) {
	vset := token.NewFileSet()
	var vbuf bytes.Buffer
	printer.Fprint(&vbuf, vset, v)
	fmt.Printf("===============Vinh- compile\n%s\n end vinh=================", vbuf.String())
}

package main

import (
	"fmt"
	"github.com/mochi-lang/mochi/interpreter"
	"github.com/mochi-lang/mochi/repl"
	"os"
)

// Mochi version
//	-ldflags "-X main.Version=`git describe --always --tags`"
var VERSION string

const (
	Lang = "Mochi"
)

func main() {
	fmt.Printf("\033[48;5;95;38;5;214m%s. Rev %s\033[0m\n", Lang, VERSION)

	if len(os.Args) > 1 {
		interpreter.Run(os.Args[1])
		return
	}

	repl.Run()
}

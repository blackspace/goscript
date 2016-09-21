//go:generate -command yacc go tool yacc
//go:generate yacc -o parser/parser.go -v parser/parser.y.out  parser/parser.go.y

package main

import (
	"os"
	"bufio"
	"io"
	"goscript/vm"
)

func main() {
	stdin:=bufio.NewReader(os.Stdin)

	source:=""

	for {
		l,_,err:=stdin.ReadLine()

		if err==io.EOF {
			break
		}

		source+=string(l)
	}

	vm:=vm.NewVM()

	vm.Run(source)
}


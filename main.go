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
	var r * bufio.Reader
	if len(os.Args)==1 {
		r=bufio.NewReader(os.Stdin)
	} else {
		f_name:=os.Args[1]

		f,_:=os.Open(f_name)

		r=bufio.NewReader(f)
	}

	source:=""

	for {
		buf:=make([]byte,1024)


		n,err:=r.Read(buf)

		if err==io.EOF {
			break
		}

		source+=string(buf[:n])
	}

	vm:=vm.NewVM()

	vm.Run(source)
}


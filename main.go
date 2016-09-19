//go:generate -command yacc go tool yacc
//go:generate yacc -o parser/parser.go -v parser/parser.y.out  parser/parser.go.y

package main

import (
	"goscript/vm"
)

func main() {
	var vm = vm.NewVM()

	vm.Run("for a=1;a<10;a++ { a }" )
}


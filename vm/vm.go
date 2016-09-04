package vm

import (
	"goscript/parser"
	"reflect"
)

type VM struct {

}

func NewVM() *VM {
	return &VM{}
}

func (v * VM)Execute(src string) (reflect.Value,error) {
	expr,_ :=parser.ParseString(src)

	return expr.Eval(),nil
}


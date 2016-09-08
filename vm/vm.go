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

func (v * VM)Execute(src string) (result reflect.Value,err error) {
	expres,_ :=parser.ParseString(src)

	for _,e:= range expres {
		result=e.Eval()
	}

	return
}


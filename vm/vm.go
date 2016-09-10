package vm

import (
	"goscript/parser"
	"reflect"
	"goscript/vm/runtime"
)

type VM struct {
	Runtime	*runtime.Runtime
}

func NewVM() *VM {
	return &VM{Runtime:runtime.NewRuntime()}
}

func (v * VM)Run(src string) (result reflect.Value,err error) {
	exprs,err :=parser.ParseString(src)

	for _,e:= range exprs {
		result=e.Eval(v.Runtime)
	}

	return
}

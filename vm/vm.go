package vm

import (
	"goscript/parser"
	"reflect"
	"goscript/vm/runtime"
	"goscript/ast"
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
		result=v.Eval(e)
	}

	return
}


func (v *VM)Eval(expr ast.Expr) reflect.Value {
	return expr.Eval(v.Runtime)
}

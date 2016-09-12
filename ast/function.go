package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type FuncExpr struct {
}

func (b * FuncExpr)Eval(r *runtime.Runtime) (v reflect.Value,status int){
	return  v,0
}



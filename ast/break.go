package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type BreakExpr struct  {
}

func (n * BreakExpr)Eval(r *runtime.Runtime) (reflect.Value,int) {
	return reflect.Value{},BREAK
}



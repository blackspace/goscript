package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type BreakExpr struct  {
}

func (n * BreakExpr)Eval(r *runtime.Runtime) reflect.Value {
	return reflect.Value{}
}



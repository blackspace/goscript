package ast

import (
	"reflect"
	"goscript/vm/runtime"
)



type Number struct  {
	Int int64
}

func (n * Number)Eval(r *runtime.Runtime) reflect.Value {
	return reflect.ValueOf(n.Int)
}

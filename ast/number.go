package ast

import (
	"reflect"
	"goscript/runtime"
)



type Number struct  {
	Int int64
}

func (n * Number)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int) {
	return reflect.ValueOf(n.Int),0
}

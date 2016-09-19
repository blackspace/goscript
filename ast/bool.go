package ast

import (
	"reflect"
	"goscript/runtime"
)

type Bool struct  {
      Bool bool
}

func (n * Bool)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int) {
	return reflect.ValueOf(n.Bool),0
}


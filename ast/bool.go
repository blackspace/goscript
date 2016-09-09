package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type Bool struct  {
      Bool bool
}

func (n * Bool)Eval(r *runtime.Runtime) reflect.Value {
	return reflect.ValueOf(n.Bool)
}


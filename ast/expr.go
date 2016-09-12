package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type Expr interface {
      	Eval(r *runtime.Runtime) (reflect.Value,int)
}





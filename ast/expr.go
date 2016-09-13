package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

const (
	OK=iota
	BREAK
)

type Expr interface {
      	Eval(r *runtime.Runtime) (reflect.Value,int)
}





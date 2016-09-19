package ast

import (
	"reflect"
	"goscript/runtime"
)

const (
	OK=iota
	BREAK
	RETURN
)

type Expr interface {
      	Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int)
}





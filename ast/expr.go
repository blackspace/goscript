package ast

import (
	"goscript/runtime"
)

const (
	OK=iota
	BREAK
	RETURN
)

type Expr interface {
      	Eval(r *runtime.Runtime,args ...interface{}) (interface{},int)
}





package ast

import (
	"goscript/runtime"
)

type BreakExpr struct  {
}

func (n * BreakExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {
	return nil,BREAK
}



package ast

import (
	"goscript/runtime"
)



type Number struct  {
	Int int64
}

func (n * Number)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {
	return n.Int,0
}

package ast

import (
	"goscript/runtime"
)



type Int struct  {
	V int64
}

func (i * Int)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {
	return i.V,OK
}

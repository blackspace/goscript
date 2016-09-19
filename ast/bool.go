package ast

import (
	"goscript/runtime"
)

type Bool struct  {
      Bool bool
}

func (n * Bool)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {
	return n.Bool,0
}


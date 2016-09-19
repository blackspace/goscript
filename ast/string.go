package ast

import (
	"goscript/runtime"
)

type String struct {
	S string
}

func (s * String)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {
	return s.S,0
}


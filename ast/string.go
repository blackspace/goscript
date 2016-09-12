package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type String struct {
	S string
}

func (s * String)Eval(r *runtime.Runtime) (reflect.Value,int) {
	return reflect.ValueOf(s.S),0
}


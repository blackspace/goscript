package ast

import (
	"reflect"
	"goscript/runtime"
)

type String struct {
	S string
}

func (s * String)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int) {
	return reflect.ValueOf(s.S),0
}


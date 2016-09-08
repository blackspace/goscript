package ast

import "reflect"

type String struct {
	S string
}

func (s * String)Eval() reflect.Value {
	return reflect.ValueOf(s.S)
}


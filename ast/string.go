package ast

import "reflect"

type String string

func (s * String)Eval() reflect.Value {
	return reflect.ValueOf(*s)
}


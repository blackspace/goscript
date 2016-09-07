package ast

import "reflect"


type Number struct  {
	Int int64
}

func (n * Number)Eval() reflect.Value {
	return reflect.ValueOf(n.Int)
}

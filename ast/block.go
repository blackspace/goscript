package ast

import (
	"reflect"
	"goscript/vm/runtime"
)



type BlockExpr struct {
	Exprs []Expr
}


func (b * BlockExpr)Eval(r *runtime.Runtime) (v reflect.Value,status int) {

	s:=r.BeginScope()

	for _,e:=range b.Exprs {
		tv,status:=e.Eval(r)

		if status==BREAK {
			if true {
				return v,status
			} else {
				return reflect.ValueOf(nil),status
			}
		}

		v=tv
	}

	r.EndScope(s)

	return  v,OK
}


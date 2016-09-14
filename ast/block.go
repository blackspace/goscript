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

		switch status {
		case BREAK:
			if status==BREAK {
				if !v.IsValid() {
					return reflect.ValueOf(nil),status
				} else {
					return v,status
				}
			}
		case RETURN:
			return tv,OK
		}



		v=tv
	}

	r.EndScope(s)

	return  v,OK
}


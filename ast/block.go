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
		if _,ok:=e.(*BreakExpr);ok {
			return
		}

		v,_=e.Eval(r)


	}

	r.EndScope(s)

	return  v,0
}


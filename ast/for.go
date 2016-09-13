package ast

import (
	"goscript/vm/runtime"
	"reflect"
)

type ForExpr struct {
	Expr0 Expr
	Expr1 Expr
	Expr2 Expr
	Expr3 Expr
}


func (e *ForExpr)Eval(r *runtime.Runtime) (v reflect.Value,status int) {
	f1:=e.Expr0
	f2:=e.Expr1
	f3:=e.Expr2

	s:=r.BeginScope()

	if f1!=nil {
		f1.Eval(r)
	}

	for  {
		if f2!=nil {
			if v,_:=f2.Eval(r);!v.Bool() {
				break
			}
		}

		v,status=e.Expr3.Eval(r)

		if status==BREAK {
			return
		}

		if f3!=nil {
			f3.Eval(r)
		}

	}

	r.EndScope(s)

	return v,0

}


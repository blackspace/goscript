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


func (e *ForExpr)Eval(r *runtime.Runtime) (v reflect.Value) {
	f1:=e.Expr0
	f2:=e.Expr1
	f3:=e.Expr2

	s:=r.BeginScope()

	if f1!=nil {
		f1.Eval(r)
	}

	for  {
		if f2!=nil {
			if ! f2.Eval(r).Bool() {
				break
			}
		}


		v=e.Expr3.Eval(r)


		if f3!=nil {
			f3.Eval(r)
		}

	}

	r.EndScope(s)

	return

}


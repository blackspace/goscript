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


	if f1!=nil {
		f1.Eval(r)
	}

	for  {
		if f2!=nil {
			if ! f2.Eval(r).Bool() {
				break
			}
		}

		if e.Expr3!=nil {
			v=e.Expr3.Eval(r)
		}

		if f3!=nil {
			f3.Eval(r)
		}

	}

	return

}


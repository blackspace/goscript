package ast

import  (
	"reflect"
	"goscript/vm/runtime"
)

type IFExpr struct {
	Expr0 Expr
	Expr1 Expr
	Expr2 Expr
}


func (e * IFExpr)Eval(r *runtime.Runtime) (v reflect.Value,status int){

	v0,_ :=e.Expr0.Eval(r)

	if v0.Bool() {
		s:=r.BeginScope()
		v,status=e.Expr1.Eval(r)
		r.EndScope(s)
	} else {
		if e.Expr2==nil {
			return
		}

		s:=r.BeginScope()
		v,status=e.Expr2.Eval(r)
		r.EndScope(s)
	}


	return v,status
}



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


func (e * IFExpr)Eval(r *runtime.Runtime) (v reflect.Value){
	v0 :=e.Expr0.Eval(r)

	if v0.Bool() {
		v=e.Expr1.Eval(r)
	} else {
		v=e.Expr2.Eval(r)
	}

	return
}



package ast

import (
	"goscript/vm/runtime"
	"reflect"
)

type SubfixDoubleAddExpr struct {
	Expr0  Expr
}


func (e *SubfixDoubleAddExpr)Eval(r *runtime.Runtime) (v reflect.Value,status int) {
	name :=e.Expr0.(*VarExpr).Name

	if v,ok:=r.GetVarible(name);ok{
		n:=v.Int()

		n++
		r.SetVarible(name,reflect.ValueOf(n))
	}
	v,_=r.GetVarible(name);
	return v,0
}

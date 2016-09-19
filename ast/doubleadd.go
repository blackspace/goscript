package ast

import (
	"goscript/runtime"
	"reflect"
)

type SubfixDoubleAddExpr struct {
	Expr0  Expr
}


func (e *SubfixDoubleAddExpr)Eval(r *runtime.Runtime,args ...interface{}) (v reflect.Value,status int) {
	name :=e.Expr0.(*Variable).Name

	if v,ok:=r.GetVarible(name);ok{
		n:=v.Int()

		n++
		r.SetVarible(name,reflect.ValueOf(n))
	}
	v,_=r.GetVarible(name);
	return v,0
}

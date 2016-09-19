package ast

import (
	"goscript/runtime"
)

type SubfixDoubleAddExpr struct {
	Expr0  Expr
}


func (e *SubfixDoubleAddExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {
	name :=e.Expr0.(*Variable).Name

	if v,ok:=r.GetVarible(name);ok{
		n:=v.(int64)
		n++
		r.SetVarible(name,n)
	}
	v,_:=r.GetVarible(name);

	return v,OK
}

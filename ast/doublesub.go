package ast

import (
	"goscript/runtime"
)

type SubfixDoubleSubExpr struct {
	Expr0  Expr
}


func (e *SubfixDoubleSubExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {

	name :=e.Expr0.(*Variable).Name

	if v:=r.GetVarible(name);v!=nil{
		n:=v.(int64)
		n--
		r.SetVarible(name,n)
	}

	v:=r.GetVarible(name)

	return v,OK
}

package ast

import (
	"reflect"
	"goscript/runtime"
)

type AssignExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (a * AssignExpr)Eval(r *runtime.Runtime) (reflect.Value,int){
	v2,_ :=a.Expr2.Eval(r)

	v1,_:=a.Expr1.(*Variable)
	n:=v1.Name

	r.SetVarible(n,v2)

	return  v2,0
}


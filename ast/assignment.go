package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type AssignExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (a * AssignExpr)Eval(r *runtime.Runtime) reflect.Value{
	v :=a.Expr2.Eval(r)

	n:=a.Expr1.(*VarExpr).Name

	r.SetVarible(n,v)

	return  v
}


package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type AssignExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (a * AssignExpr)Eval() reflect.Value{
	v :=a.Expr2.Eval()

	n:=a.Expr1.(*VarExpr).Name

	runtime.AddSymbols( n,v)

	return  v
}


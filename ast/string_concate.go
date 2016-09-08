package ast

import "reflect"

type StringConcateExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * StringConcateExpr)Eval() reflect.Value{
	v1 :=s.Expr1.Eval()
	v2 :=s.Expr2.Eval()
	return  reflect.ValueOf(v1.String()+v2.String())
}


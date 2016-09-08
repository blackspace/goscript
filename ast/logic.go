package ast

import "reflect"

type ORExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * ORExpr)Eval() reflect.Value{
	v1 :=e.Expr1.Eval()
	v2 :=e.Expr2.Eval()
	return  reflect.ValueOf(v1.Bool()||v2.Bool())
}


type ANDExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * ANDExpr)Eval() reflect.Value{
	v1 :=e.Expr1.Eval()
	v2 :=e.Expr2.Eval()
	return  reflect.ValueOf(v1.Bool()&&v2.Bool())
}


package ast

import "reflect"

type AddExpr struct {
	Expr1 Expr
	Expr2 Expr
}

func (a  * AddExpr)Eval() reflect.Value{
	v1 :=a.Expr1.Eval()
	v2 :=a.Expr2.Eval()
	return reflect.ValueOf(v1.Int()+v2.Int())
}

type SubExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * SubExpr)Eval() reflect.Value{
	v1 :=s.Expr1.Eval()
	v2 :=s.Expr2.Eval()
	return  reflect.ValueOf(v1.Int()-v2.Int())
}

type MultiExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * MultiExpr)Eval() reflect.Value{
	v1 :=s.Expr1.Eval()
	v2 :=s.Expr2.Eval()
	return  reflect.ValueOf(v1.Int()*v2.Int())
}


type DivExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * DivExpr)Eval() reflect.Value{
	v1 :=s.Expr1.Eval()
	v2 :=s.Expr2.Eval()
	return  reflect.ValueOf(v1.Int()/v2.Int())
}

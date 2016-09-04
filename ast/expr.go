package ast

import (
	"reflect"
)

type Expr interface {
      	Eval() reflect.Value
}

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

type Number struct {
	Num int
}

func (n * Number)Eval() reflect.Value {
	return reflect.ValueOf(n.Num)
}
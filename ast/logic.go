package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type ORExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * ORExpr)Eval(r *runtime.Runtime) reflect.Value{
	v1 :=e.Expr1.Eval(r)
	v2 :=e.Expr2.Eval(r)
	return  reflect.ValueOf(v1.Bool()||v2.Bool())
}


type ANDExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * ANDExpr)Eval(r *runtime.Runtime) reflect.Value{
	v1 :=e.Expr1.Eval(r)
	v2 :=e.Expr2.Eval(r)
	return  reflect.ValueOf(v1.Bool()&&v2.Bool())
}


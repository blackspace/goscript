package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type LessThanExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * LessThanExpr)Eval(r *runtime.Runtime) reflect.Value{
	v1:=e.Expr1.Eval(r).Int()
	v2:=e.Expr2.Eval(r).Int()

	return  reflect.ValueOf(v1<v2)

}

type LessEqualExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * LessEqualExpr)Eval(r *runtime.Runtime) reflect.Value{
	v1:=e.Expr1.Eval(r).Int()
	v2:=e.Expr2.Eval(r).Int()

	return  reflect.ValueOf(v1<=v2)

}

type GreaterThanExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * GreaterThanExpr)Eval(r *runtime.Runtime) reflect.Value{
	v1:=e.Expr1.Eval(r).Int()
	v2:=e.Expr2.Eval(r).Int()

	return  reflect.ValueOf(v1>v2)
}


type GreaterEqualExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * GreaterEqualExpr)Eval(r *runtime.Runtime) reflect.Value{
	v1:=e.Expr1.Eval(r).Int()
	v2:=e.Expr2.Eval(r).Int()

	return  reflect.ValueOf(v1>=v2)

}

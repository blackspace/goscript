package ast

import (
	"reflect"
	"goscript/runtime"
)

type ORExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * ORExpr)Eval(r *runtime.Runtime) (reflect.Value,int){
	v1,_ :=e.Expr1.Eval(r)
	v2,_ :=e.Expr2.Eval(r)
	return  reflect.ValueOf(v1.Bool()||v2.Bool()),0
}


type ANDExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * ANDExpr)Eval(r *runtime.Runtime) (reflect.Value,int){
	v1,_ :=e.Expr1.Eval(r)
	v2,_ :=e.Expr2.Eval(r)
	return  reflect.ValueOf(v1.Bool()&&v2.Bool()),0
}


package ast

import (
	"reflect"
	"goscript/runtime"
)

type LessThanExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * LessThanExpr)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int){
	v1,_:=e.Expr1.Eval(r)
	v2,_:=e.Expr2.Eval(r)

	return  reflect.ValueOf(v1.Int()<v2.Int()),0

}

type LessEqualExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * LessEqualExpr)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int){
	v1,_:=e.Expr1.Eval(r)
	v2,_:=e.Expr2.Eval(r)

	return  reflect.ValueOf(v1.Int()<=v2.Int()),0

}

type GreaterThanExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * GreaterThanExpr)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int){
	v1,_:=e.Expr1.Eval(r)
	v2,_:=e.Expr2.Eval(r)

	return  reflect.ValueOf(v1.Int()>v2.Int()),0
}


type GreaterEqualExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * GreaterEqualExpr)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int){
	v1,_:=e.Expr1.Eval(r)
	v2,_:=e.Expr2.Eval(r)

	return  reflect.ValueOf(v1.Int()>=v2.Int()),0

}

package ast

import (
	"reflect"
	"strconv"
	"goscript/runtime"
)

type AddExpr struct {
	Expr1 Expr
	Expr2 Expr
}

func (a  * AddExpr)Eval(r *runtime.Runtime,args ...interface{}) (v reflect.Value,status int){
	v1,_ :=a.Expr1.Eval(r)
	v2,_ :=a.Expr2.Eval(r)

	t1:=v1.Type().Kind()
	t2:=v2.Type().Kind()


	switch t1 {
	case reflect.String:
		switch t2 {
		case reflect.String:
			return reflect.ValueOf(v1.String()+v2.String()),0
		case reflect.Int64:
			s2:=strconv.FormatInt(v2.Int(),10)
			return reflect.ValueOf(v1.String()+s2),0
		}
	case reflect.Int64:
		switch t2 {
		case reflect.String:
			s1:=strconv.FormatInt(v1.Int(),10)
			return reflect.ValueOf(s1+v2.String()),0
		case reflect.Int64:
			return reflect.ValueOf(v1.Int()+v2.Int()),0
		}
	}

	return
}

type SubExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * SubExpr)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int){
	v1,_ :=s.Expr1.Eval(r)
	v2,_ :=s.Expr2.Eval(r)
	return  reflect.ValueOf(v1.Int()-v2.Int()),0
}

type MultiExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * MultiExpr)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int){
	v1,_ :=s.Expr1.Eval(r)
	v2,_ :=s.Expr2.Eval(r)
	return  reflect.ValueOf(v1.Int()*v2.Int()),0
}


type DivExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * DivExpr)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int){
	v1,_ :=s.Expr1.Eval(r)
	v2,_ :=s.Expr2.Eval(r)
	return  reflect.ValueOf(v1.Int()/v2.Int()),0
}

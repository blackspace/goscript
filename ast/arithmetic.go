package ast

import (
	"reflect"
	"log"
	"strconv"
	"goscript/vm/runtime"
)

type AddExpr struct {
	Expr1 Expr
	Expr2 Expr
}

func (a  * AddExpr)Eval(r *runtime.Runtime) (v reflect.Value){
	v1 :=a.Expr1.Eval(r)
	v2 :=a.Expr2.Eval(r)

	t1:=v1.Type().Kind()
	t2:=v2.Type().Kind()

	log.Println(t1,t2)



	switch t1 {
	case reflect.String:
		switch t2 {
		case reflect.String:
			return reflect.ValueOf(v1.String()+v2.String())
		case reflect.Int64:
			s2:=strconv.FormatInt(v2.Int(),10)
			return reflect.ValueOf(v1.String()+s2)
		}
	case reflect.Int64:
		switch t2 {
		case reflect.String:
			s1:=strconv.FormatInt(v1.Int(),10)
			return reflect.ValueOf(s1+v2.String())
		case reflect.Int64:
			return reflect.ValueOf(v1.Int()+v2.Int())
		}
	}

	return
}

type SubExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * SubExpr)Eval(r *runtime.Runtime) reflect.Value{
	v1 :=s.Expr1.Eval(r)
	v2 :=s.Expr2.Eval(r)
	return  reflect.ValueOf(v1.Int()-v2.Int())
}

type MultiExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * MultiExpr)Eval(r *runtime.Runtime) reflect.Value{
	v1 :=s.Expr1.Eval(r)
	v2 :=s.Expr2.Eval(r)
	return  reflect.ValueOf(v1.Int()*v2.Int())
}


type DivExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * DivExpr)Eval(r *runtime.Runtime) reflect.Value{
	v1 :=s.Expr1.Eval(r)
	v2 :=s.Expr2.Eval(r)
	return  reflect.ValueOf(v1.Int()/v2.Int())
}

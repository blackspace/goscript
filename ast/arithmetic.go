package ast

import (
	"strconv"
	"goscript/runtime"
)

type AddExpr struct {
	Expr1 Expr
	Expr2 Expr
}

func (a  * AddExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int){
	v1,_ :=a.Expr1.Eval(r)
	v2,_ :=a.Expr2.Eval(r)


	switch v1:=v1.(type) {
	case string:
		switch v2:=v2.(type) {
		case string:
			return v1+v2,0
		case int64:
			s2:=strconv.FormatInt(v2,10)
			return v1+s2,0
		}
	case int64:
		switch v2:=v2.(type) {
		case string:
			s1:=strconv.FormatInt(v1,10)
			return s1+v2,0
		case int64:
			return v1+v2,0
		}
	}

	return
}

type SubExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * SubExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int){
	v1,_ :=s.Expr1.Eval(r)
	v2,_ :=s.Expr2.Eval(r)
	return  v1.(int64)-v2.(int64),0
}

type MultiExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * MultiExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int){
	v1,_ :=s.Expr1.Eval(r)
	v2,_ :=s.Expr2.Eval(r)
	return  v1.(int64)*v2.(int64),0
}


type DivExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (s * DivExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int){
	v1,_ :=s.Expr1.Eval(r)
	v2,_ :=s.Expr2.Eval(r)
	return  v1.(int64)/v2.(int64),0
}

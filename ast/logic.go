package ast

import (
	"goscript/runtime"
)

type ORExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * ORExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int){
	v1,_ :=e.Expr1.Eval(r)
	v2,_ :=e.Expr2.Eval(r)
	return  v1.(bool)||v2.(bool),0
}


type ANDExpr struct {
	Expr1 Expr
	Expr2 Expr
}


func (e * ANDExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int){
	v1,_ :=e.Expr1.Eval(r)
	v2,_ :=e.Expr2.Eval(r)
	return  v1.(bool)&&v2.(bool),0
}


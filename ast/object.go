package ast

import (
	"goscript/runtime"
	"goscript/buildin"
)

type KeyValueExpr struct {
	Key string
	Value Expr
}

func (oe * KeyValueExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {
	return
}


type ObjectExpr struct {
	Exprs []Expr
}


func (oe * ObjectExpr)Eval(r *runtime.Runtime,args ...interface{}) (result interface{},status int) {

	o:=buildin.EmptyObject.NewObject()

	for _,e:=range oe.Exprs {


		k:=e.(*KeyValueExpr).Key
		v,_:=e.(*KeyValueExpr).Value.Eval(r)

		o.SetMember(k,v)
	}

	return o,OK
}



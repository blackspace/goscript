package ast

import (
	"reflect"
	"goscript/vm/runtime"
)



type BlockExpr struct {
	Exprs []Expr
}


func (b * BlockExpr)Eval(r *runtime.Runtime) (v reflect.Value){

	for _,e:=range b.Exprs {
		v=e.Eval(r)
	}

	return  v
}


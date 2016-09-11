package ast

import (
	"reflect"
	"goscript/vm/runtime"
)



type BlockExpr struct {
	Exprs []Expr
}


func (b * BlockExpr)Eval(r *runtime.Runtime) (v reflect.Value){


	s:=r.BeginScope()

	for _,e:=range b.Exprs {
		v=e.Eval(r)
	}

	r.EndScope(s)

	return  v
}


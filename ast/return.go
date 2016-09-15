package ast

import (
	"reflect"
	"goscript/runtime"
)

type ReturnExpr struct  {
	ValueExpr Expr
}

func (n * ReturnExpr)Eval(r *runtime.Runtime) (reflect.Value,int) {
	v,_:=n.ValueExpr.Eval(r)


	return v,RETURN
}



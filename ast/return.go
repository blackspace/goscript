package ast

import (
	"goscript/runtime"
)

type ReturnExpr struct  {
	ValueExpr Expr
}

func (n * ReturnExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {
	v,_:=n.ValueExpr.Eval(r)


	return v,RETURN
}



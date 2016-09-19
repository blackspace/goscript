package ast

import (
	"goscript/runtime"
)



type BlockExpr struct {
	Exprs []Expr
}


func (b * BlockExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {

	s:=r.BeginScope()

	for _,e:=range b.Exprs {
		tv,status:=e.Eval(r)

		switch status {
		case BREAK:
			if status==BREAK {
				return v,status
			}
		case RETURN:
			return tv,OK
		}



		v=tv
	}

	r.EndScope(s)

	return  v,OK
}


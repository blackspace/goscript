package ast

import (
	"goscript/runtime"
)

type ForExpr struct {
	InitExpr Expr
	Expr1    Expr
	StepExpr Expr
	BodyExpr Expr
}


func (e *ForExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {
	f2:=e.Expr1

	s:=r.BeginScope()

	if e.InitExpr!=nil {
		e.InitExpr.Eval(r)
	}

	for  {
		if f2!=nil {
			if lv,_:=f2.Eval(r);!lv.(bool) {
				break
			}
		}

		v,status=e.BodyExpr.Eval(r)

		if status==BREAK {
			return nil,OK
		}

		if e.StepExpr!=nil {
			e.StepExpr.Eval(r)
		}

	}

	r.EndScope(s)

	return v,0

}


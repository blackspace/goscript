package ast

import  (
	"goscript/runtime"
)

type IFExpr struct {
	Expr0 Expr
	Expr1 Expr
	Expr2 Expr
}


func (e * IFExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int){

	v0,_ :=e.Expr0.Eval(r)

	if v0.(bool) {
		s:=r.BeginScope()
		v,status=e.Expr1.Eval(r)
		r.EndScope(s)
	} else if e.Expr2!=nil {
		s:=r.BeginScope()
		v,status=e.Expr2.Eval(r)
		r.EndScope(s)
	}


	return v,status
}



package ast

import "reflect"

type IFExpr struct {
	Expr0 Expr
	Expr1 []Expr
	Expr2 []Expr
}


func (e * IFExpr)Eval() (v reflect.Value){
	v0 :=e.Expr0.Eval()

	if v0.Bool() {
		for _,e:=range e.Expr1 {
			v=e.Eval()
		}
	} else {
		for _,e:=range e.Expr2 {
			v =e.Eval()
		}
	}

	return
}



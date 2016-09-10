package ast

import (
	"goscript/vm/runtime"
	"reflect"
)

type ForExpr struct {
	Expr0 Expr
	Expr1 Expr
	Expr2 Expr
	Expr3 []Expr
}


func (e *ForExpr)Eval(r *runtime.Runtime) (v reflect.Value) {
	return
}
package ast

import (
	"goscript/vm/runtime"
	"reflect"
	"log"
)

type SubfixDoubleAddExpr struct {
	Expr0  Expr
}


func (e *SubfixDoubleAddExpr)Eval(r *runtime.Runtime) (v reflect.Value) {

	name :=e.Expr0.(*VarExpr).Name

	log.Println(name)

	if v,ok:=r.GetSymbolValue(name);ok{
		n:=v.Int()
		n++
		r.AddSymbols(name,reflect.ValueOf(n))
	}
	v,_=r.GetSymbolValue(name);
	return
}

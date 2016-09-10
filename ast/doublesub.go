package ast

import (
	"goscript/vm/runtime"
	"reflect"
	"log"
)

type SubfixDoubleSubExpr struct {
	Expr0  Expr
}


func (e *SubfixDoubleSubExpr)Eval(r *runtime.Runtime) (v reflect.Value) {

	name :=e.Expr0.(*VarExpr).Name

	log.Println(name)

	if v,ok:=r.Symbols.Get(name);ok{
		n:=v.Int()
		n--
		r.Symbols.Set(name,reflect.ValueOf(n))
	}
	v,_=r.Symbols.Get(name);
	return
}

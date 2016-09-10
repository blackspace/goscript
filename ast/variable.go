package ast

import (
	"reflect"
	"goscript/vm/runtime"
	"errors"
)

type VarExpr struct {
	Name string
}

func (s * VarExpr)Eval(r *runtime.Runtime) reflect.Value {
	if v,ok:=r.Symbols.Get(s.Name);ok {
		return v
	} else {
		panic(errors.New("The "+s.Name+" varible is not existing"))
	}

}



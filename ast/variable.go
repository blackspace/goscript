package ast

import (
	"reflect"
	"goscript/vm/runtime"
	"errors"
)

type VarExpr struct {
	Name string
}

func (s * VarExpr)Eval(r *runtime.Runtime) (reflect.Value,int) {
	if v,ok:=r.GetVarible(s.Name);ok {
		return v,OK
	} else {
		panic(errors.New("The "+s.Name+" varible is not existing"))
	}

}



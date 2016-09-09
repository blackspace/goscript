package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type VarExpr struct {
	Name string
}

func (s * VarExpr)Eval(r *runtime.Runtime) reflect.Value {
	if v,ok:=r.Symbols[s.Name];ok {
		return v
	} else {
		r.Symbols[s.Name]=reflect.Value{}
		return runtime.Symbols[s.Name]
	}

}



package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type VarExpr struct {
	Name string
}

func (s * VarExpr)Eval() reflect.Value {
	if v,ok:=runtime.Symbols[s.Name];ok {

		return v
	} else {
		runtime.Symbols[s.Name]=reflect.Value{}
		return runtime.Symbols[s.Name]
	}

}



package ast

import (
	"goscript/runtime"
	"errors"
)

type Variable struct {
	Name string
}

func (s *Variable)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {
	if v:=r.GetVarible(s.Name);v!=nil {
		return v,OK
	} else {
		panic(errors.New("The "+s.Name+" varible is not existing"))
	}

}



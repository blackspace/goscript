package ast

import (
	"reflect"
	"goscript/runtime"
	"errors"
)

type Variable struct {
	Name string
}

func (s *Variable)Eval(r *runtime.Runtime,args ...interface{}) (reflect.Value,int) {
	if v,ok:=r.GetVarible(s.Name);ok {
		return v,OK
	} else {
		panic(errors.New("The "+s.Name+" varible is not existing"))
	}

}



package vm

import (
	"goscript/parser"
	"reflect"
	"goscript/runtime"
)

type VM struct {
	Runtime	*runtime.Runtime
}

func NewVM() (v *VM) {
	v=&VM{Runtime:runtime.NewRuntime()}
	v.Runtime.BeginScope()
	return
}

func (v * VM)Run(src string) (result reflect.Value,err error) {
	if exprs,err :=parser.ParseString(src);err==nil {
		for _,e:= range exprs {
			result,_=e.Eval(v.Runtime)
		}
	} else {
		return result,err
	}

	return

}

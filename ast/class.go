package ast

import  (
	"reflect"
	"goscript/vm/runtime"
)

type Class struct  {
	Bool bool
}

func (c * Class)Eval(r *runtime.Runtime) (v reflect.Value,status int) {
	return
}


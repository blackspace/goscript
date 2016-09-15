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

type Method struct  {
	Bool bool
}

func (m * Method)Eval(r *runtime.Runtime) (v reflect.Value,status int) {
	return
}

type Attribute struct  {
	Bool bool
}

func (m * Attribute)Eval(r *runtime.Runtime) (v reflect.Value,status int) {
	return
}

type ClassAttribute struct  {
	Bool bool
}

func (m * ClassAttribute)Eval(r *runtime.Runtime) (v reflect.Value,status int) {
	return
}


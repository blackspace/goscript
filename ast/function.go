package ast

import (
	"reflect"
	"goscript/vm/runtime"
)

type Params []Variable

type FuncDefineExpr struct {
	Name string
	Params []string
	Body  Expr
}

func (f * FuncDefineExpr)Eval(r *runtime.Runtime) (v reflect.Value,status int){
	return
}

type FuncCallExpr struct {
	Name string
	Params []reflect.Value
}

func (f * FuncCallExpr)Eval(r *runtime.Runtime) (v reflect.Value,status int){
	return
}



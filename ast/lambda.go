package ast

import (
	"goscript/runtime"
)

type LambdaDefineExpr struct {
	Params []string
	Body  []Expr
}


func (f * LambdaDefineExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int){
	return MakeFunction(r,f.Params,f.Body),OK
}

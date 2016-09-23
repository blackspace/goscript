package ast

import (
	"goscript/runtime"
)

type FuncDefineExpr struct {
	Name string
	Params []string
	Body  []Expr
}


func MakeFunc(r *runtime.Runtime,params []string,body []Expr) runtime.Function {
	return func(r *runtime.Runtime,params []string,body []Expr) runtime.Function {
		return func(in []interface{}) (result interface{},status int) {
			s:=r.BeginScope()

			if params!=nil {
				for i,p:=range params{
					if i<len(in) { // The  count of passing value to the called function less the params count
						r.SetVarible(p,in[i])
					}
				}
			}


			for _,e:=range body {
				result,status=e.Eval(r)
			}



			r.EndScope(s)

			return result,status
		}
	}(r,params,body)
}


func (f * FuncDefineExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int){
	r.SetFunction(f.Name,MakeFunc(r,f.Params,f.Body))

	return nil,OK
}




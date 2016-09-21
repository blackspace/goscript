package ast

import (
	"goscript/runtime"
	"errors"
)

type FuncDefineExpr struct {
	Name string
	Params []string
	Body  Expr
}


func (f * FuncDefineExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int){
	r.SetFunction(f.Name,func(f * FuncDefineExpr) runtime.Function {
		return func(in []interface{}) (interface{},int) {
			s:=r.BeginScope()

			if f.Params!=nil {
				for i,p:=range f.Params{
					if i<len(in) { // The  count of passing value to the called function less the params count
						r.SetVarible(p,in[i])
					}
				}
			}

			lv,lstatus:=f.Body.Eval(r)

			r.EndScope(s)

			return lv,lstatus
		}
	}(f))

	return nil,OK
}

type FuncCallExpr struct {
	Name string
	Params []Expr
}

func (f * FuncCallExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int){
	F:=r.GetFunction(f.Name)


	if F==nil {
		panic(errors.New("Can't find the "+f.Name+" function"))
	}

	if f.Params==nil {

		in:=make([]interface{},0,10)

		v,status=F(in)

		return
	} else {

		vs:=make([]interface{},0,10)

		for _,p:=range f.Params {
			v,_:=p.Eval(r)
			vs=append(vs,v)
		}

		v,status=F(vs)

		return

	}

	return
}



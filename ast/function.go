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



type _Function func(in []reflect.Value) (reflect.Value,int)

func (f * FuncDefineExpr)Eval(r *runtime.Runtime) (v reflect.Value,status int){
	r.SetFunction(f.Name,func(f * FuncDefineExpr) _Function {
		return func(in []reflect.Value) (reflect.Value,int) {
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

	return
}

type FuncCallExpr struct {
	Name string
	Params []Expr
}

func (f * FuncCallExpr)Eval(r *runtime.Runtime) (v reflect.Value,status int){
	F:=r.GetFunction(f.Name).(_Function)


	if f.Params==nil {

		in:=make([]reflect.Value,0,10)

		v,status=F(in)

		return
	} else {

		vs:=make([]reflect.Value,0,10)

		for _,p:=range f.Params {
			v,_:=p.Eval(r)
			vs=append(vs,v)
		}

		v,status=F(vs)

		return

	}

	return
}



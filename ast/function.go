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
	r.SetFunction(f.Name,func(f * FuncDefineExpr) interface{} {
		return func(in []reflect.Value) (interface{},int) {
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
	F:=reflect.ValueOf(r.GetFunction(f.Name))



	if f.Params==nil {

		vs:=make([]reflect.Value,0,10)

		in :=[]reflect.Value{reflect.ValueOf(vs)}

		result:=F.Call(in)

		return result[0].Interface().(reflect.Value),int(result[1].Int())
	} else {

		vs:=make([]reflect.Value,0,10)

		for _,p:=range f.Params {
			v,_:=p.Eval(r)
			vs=append(vs,v)
		}

		in :=[]reflect.Value{reflect.ValueOf(vs)}

		result:=F.Call(in)

		return result[0].Interface().(reflect.Value),int(result[1].Int())

	}

	return
}



package ast

import (
	"goscript/runtime"
	"goscript/buildin"
	"errors"
)

type ArrayExpr struct {
	Exprs []Expr
}


func (ae * ArrayExpr)Eval(r *runtime.Runtime,args ...interface{}) (result interface{},status int) {
	o:=buildin.Array.NewObject()

	for _,e:=range ae.Exprs {
		v,_:=e.Eval(r)

		buildin.Array_AddElement(o,v)
	}

	return o,OK
}

type ArrayGetExpr struct {
	Path []string
	Index Expr
}


func (age * ArrayGetExpr)Eval(r *runtime.Runtime,args ...interface{}) (result interface{},status int) {
	ns:=r.PathToNodes(age.Path)
	
	if len(ns)==len(age.Path) {
		name:=age.Path[len(age.Path)-1]


		var o interface{}

		if len(ns)==1 {
			o=r.GetVarible(name)
		} else {
			o=ns[len(ns)-1].GetMember(name)
		}

		v:=o.(*runtime.Object).GetMember("_")

		i,_:=age.Index.Eval(r)

		array:=v.(buildin.LArray)

		result=array[i.(int64)]


		return result,OK

	}

	if len(ns)==len(age.Path)-1 {

	}

	if len(ns)<len(age.Path)-1 {
		panic(errors.New("This path  is illegal"))
	}

	return 
}


type ArraySetExpr struct {
	Path []string
	Index Expr
	Value interface{}
}


func (age * ArraySetExpr)Eval(r *runtime.Runtime,args ...interface{}) (result interface{},status int) {



	return 
}


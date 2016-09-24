package ast

import (
	"goscript/runtime"
	"goscript/buildin"
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
	Path      []string
	IndexExpr Expr
}


func (age * ArrayGetExpr)Eval(r *runtime.Runtime,args ...interface{}) (result interface{},status int) {
	path:=runtime.NewPath(r,age.Path)
	ns:=path.GetNodes()
	
	if len(ns)==len(age.Path) {
		o:=ns[len(ns)-1]

		v:=o.(*runtime.Object).GetMember("_")

		i,_:=age.IndexExpr.Eval(r)

		array:=v.(buildin.LArray)

		result=array[i.(int64)]


		return result,OK

	}
	return 
}


type ArraySetExpr struct {
	Path      []string
	IndexExpr Expr
	ValueExpr Expr
}


func (ase * ArraySetExpr)Eval(r *runtime.Runtime,args ...interface{}) (result interface{},status int) {
	path:=runtime.NewPath(r,ase.Path)
	ns:=path.GetNodes()

	if len(ns)==len(ase.Path) {

		o:=ns[len(ns)-1]

		v:=o.(*runtime.Object).GetMember("_")

		i,_:= ase.IndexExpr.Eval(r)
		value,_:= ase.ValueExpr.Eval(r)

		array:=v.(buildin.LArray)


		array[i.(int64)]=value


		return o,OK

	}


	return 
}


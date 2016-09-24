package ast

import (
	"goscript/runtime"
	"errors"
)

type SetExpr struct {
	Path []string
	Expr2 Expr
}


func (se * SetExpr)Eval(r *runtime.Runtime,args ...interface{}) (result interface{},status int){
	path:=runtime.NewPath(r,se.Path)
	ns:=path.GetNodes()

	if len(ns)==len(se.Path) {
		panic(errors.New("Nodes can't be assigned"))
	}

	if len(ns)==len(se.Path)-1 {
		name:=se.Path[len(se.Path)-1]


		if LDE,ok:=se.Expr2.(*LambdaDefineExpr);ok {
			if len(ns)==0 {
				F,_:= LDE.Eval(r)
				r.SetFunction(name, F.(runtime.Function))
			} else {
				n:=ns[len(ns)-1]
				switch co:=n.(type) {
				case *runtime.Class:
					co.SetClassMember(name,MakeClassMethod(r, LDE.Params, LDE.Body))
				case *runtime.Object:
					co.SetMember(name,MakeObjectMethod(r, LDE.Params, LDE.Body))
				}

			}


		} else {
			value,_:=se.Expr2.Eval(r)

			if len(ns)==0 {
				r.SetVarible(name,value)
			} else {

				ns[len(ns)-1].SetMember(name,value)
			}
		}



	}
	return
}


type GetExpr struct {
	Path []string
}

func (ge * GetExpr)Eval(r *runtime.Runtime,args ...interface{}) (result interface{},status int) {
	path:=runtime.NewPath(r,ge.Path)
	ns:=path.GetNodes()

	if len(ns)==len(ge.Path) {
		return ns[len(ns)-1],OK
	}

	if len(ns)==len(ge.Path)-1 {
		name:=ge.Path[len(ge.Path)-1]
		if len(ns)==0 {
			return r.GetVarible(name),OK
		} else 	{
			return ns[len(ns)-1].GetMember(name),OK
		}

	}


	return
}
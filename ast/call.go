package ast

import (
	"goscript/runtime"
	"errors"
)

type CalledExpr struct {
	Path []string
	Params     []Expr
}

func (ce * CalledExpr)Eval(r *runtime.Runtime,args ...interface{}) (result interface{},status int) {
	ns:=r.PathToNodes(ce.Path)

	if len(ns)==len(ce.Path) {
		panic(errors.New("Nodes can't be assigned"))
	}

	if len(ns)==len(ce.Path)-1 {
		name:=ce.Path[len(ce.Path)-1]

		vs:=make([]interface{},0,len(ce.Params))

		for _,p:=range ce.Params {
			lv,_:=p.Eval(r)

			vs=append(vs,lv)
		}

		var m interface{}
		if len(ns)==0 {
			m=r.GetFunction(name)
		} else {
			m=ns[len(ns)-1].GetMember(name)
		}



		switch F:=m.(type) {
		case runtime.Function:
			result,status=F(vs)
		case runtime.ClassMethod:
			c:=ns[len(ns)-1].(interface{}).(*runtime.Class)
			result=F(c,vs)
		case runtime.ObjectMethod:
			o:=ns[len(ns)-1].(interface{}).(*runtime.Object)
			result=F(o,vs)
		}

	}

	if len(ns)<len(ce.Path)-1 {
		panic(errors.New("This path  is illegal"))
	}

	return result,status

}



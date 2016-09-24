package ast

import (
	"goscript/runtime"
	"errors"
)

type SubfixDoubleSubExpr struct {
	Path []string
}


func (e *SubfixDoubleSubExpr)Eval(r *runtime.Runtime,args ...interface{}) (result interface{},status int) {
	path:=runtime.NewPath(r,e.Path)
	ns:=path.GetNodes()

	if len(ns)==len(e.Path) {
		panic(errors.New("Nodes can't be assigned"))
	}

	if len(ns)==len(e.Path)-1 {
		if len(ns)==0 {
			mn:=e.Path[len(e.Path)-1]
			v:=r.GetMember(mn)

			temp:=v.(int64)
			result=temp-1
			r.SetMember(mn,result)


		} else {
			n:=ns[len(ns)-1]
			mn:=e.Path[len(e.Path)-1]

			v:=n.GetMember(mn)

			temp:=v.(int64)
			result=temp-1

			n.SetMember(mn,result)

		}


	}

	return result,OK
}


package ast

import  (
	"goscript/runtime"
)

type Class struct  {
	Name string
	Exprs []Expr
}


func (c * Class)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {
	class := runtime.NewClass()

	r.SetClass(c.Name,class)

	s:=r.BeginScope()

	for _,e:=range c.Exprs {
		e.Eval(r,class)
	}

	r.EndScope(s)

	return
}



func MakeObjectMethod(r *runtime.Runtime,params []string,body []Expr) runtime.ObjectMethod {
	return func(o *runtime.Object,in []interface{}) interface{} {
		s:=r.BeginScope()

		s.Set("this",o)

		var result interface{}
		var status int

		for _, e := range body {
			result,status= e.Eval(r,o)

			switch status {
			case RETURN:
				return result
			}
		}

		r.EndScope(s)

		return result

	}
}

func MakeClassMethod(r *runtime.Runtime,params []string,body []Expr) runtime.ClassMethod {
	return func(c *runtime.Class,in []interface{}) interface{} {

		s:=r.BeginScope()

		s.Set("this",c)

		var result interface{}
		var status int

		for _, e := range body {
			result, status = e.Eval(r,c)

			switch status {
			case RETURN:
				return result
			}
		}

		r.EndScope(s)

		return result

	}
}

type MethodDefineExpr struct  {
	MethodName string
	Params []string
	Body       []Expr
}

func (m *MethodDefineExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {
	class:=args[0].(*runtime.Class)
	class.SetObjectMembers(m.MethodName, MakeObjectMethod(r,m.Params,m.Body))
	return
}







package ast

import  (
	"goscript/runtime"
	"errors"
)

type Class struct  {
	Name string
	Exprs []Expr
}


func (c * Class)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {
	class := runtime.NewClass()

	r.SetClass(c.Name,class)

	s:=r.BeginScope()

	s.Set("this",class)

	for _,e:=range c.Exprs {
		e.Eval(r,class)
	}

	r.EndScope(s)

	return
}


type MethodDefineExpr struct  {
	ObjectName string
	MethodName string
	Params     []string
	Body       []Expr
}

func (m *MethodDefineExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {
	if m.ObjectName=="" {
		class:=args[0].(*runtime.Class)

		class.SetObjectMembers(m.MethodName, runtime.ObjectMethod(func(o *runtime.Object,in []interface{}) interface{} {

			s:=r.BeginScope()

			s.Set("this",o)

			var result interface{}

			for _, e := range m.Body {
				result, status := e.Eval(r,o)

				switch status {
				case RETURN:
					return result
				}
			}

			r.EndScope(s)

			return result

		}))
	} else {
		if pc:=r.GetClass(m.ObjectName);pc!=nil {
			pc.SetClassMembers(m.MethodName, runtime.ClassMethod(func(c *runtime.Class,in []interface{}) interface{} {

				s:=r.BeginScope()

				s.Set("this",c)

				var result interface{}

				for _, e := range m.Body {
					result, status := e.Eval(r,c)

					switch status {
					case RETURN:
						return result
					}
				}

				r.EndScope(s)

				return result

			}))
		}


		if po:=r.GetVarible(m.ObjectName);po!=nil {
			switch v:=po.(type) {
			case *runtime.Object:
				v.SetAttribute(m.MethodName, runtime.ObjectMethod(func(o *runtime.Object,in []interface{}) interface{} {

					s:=r.BeginScope()

					s.Set("this",o)

					var result interface{}

					for _, e := range m.Body {
						result, status := e.Eval(r,o)

						switch status {
						case RETURN:
							return result
						}
					}

					r.EndScope(s)

					return result

				}))
			case *runtime.Class:
				v.SetClassMembers(m.MethodName, runtime.ClassMethod(func(c *runtime.Class,in []interface{}) interface{} {

					s:=r.BeginScope()

					s.Set("this",c)

					var result interface{}

					for _, e := range m.Body {
						result, status = e.Eval(r,c)

						switch status {
						case RETURN:
							return result
						}
					}

					r.EndScope(s)

					return result

				}))
			}

		}

	}




	return
}


type AttributeSetExpr struct {
	ObjectName string
	AttributeName string
	ValueExpr Expr
}

func (ase * AttributeSetExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {

	v,_=ase.ValueExpr.Eval(r)

	if po:=r.GetVarible(ase.ObjectName);po!=nil {
		switch o:=po.(type) {
		case (*runtime.Object):
			o.SetAttribute(ase.AttributeName,v)
		case (*runtime.Class):
			o.SetClassMembers(ase.AttributeName,v)
		}

	}

	return
}


type AttributeExpr struct {
	ObjectName string
	AttributeName string
}

func (ae * AttributeExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {
	if pc:=r.GetClass(ae.ObjectName);pc!=nil {
		a := pc.GetClassMembers(ae.AttributeName)

		if a!=nil {
			return a,OK
		} else {
			panic(errors.New("The "+ae.ObjectName+"class hasn't "+"the "+ae.AttributeName+" attribute"))
		}

	}


	if po:=r.GetVarible(ae.ObjectName);po!=nil {
		a := po.(*runtime.Object).GetObjectAttribute(ae.AttributeName)

		if a!=nil {
			return a,OK
		} else {
			panic(errors.New("The "+ae.ObjectName+"object hasn't "+"the "+ae.AttributeName+" attribute"))
		}
	}

	return
}


type MethodCalledExpr struct {
	ObjectName string
	MethodName string
	Params     []Expr
}

func (m * MethodCalledExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {
	if v:=r.GetVarible(m.ObjectName);v!=nil {
		o:= (v).(*runtime.Object)

		m:=o.GetObjectAttribute(m.MethodName)

		if m!=nil {
			F:=m.(runtime.ObjectMethod)
			result:=F(o,nil)
			return result,OK
		}
	}


	if c:=r.FindClass(m.ObjectName);c!=nil {
		m:=c.GetClassMembers(m.MethodName)

		if m!=nil {
			F:=m.(runtime.ClassMethod)
			result:=F(c,nil)
			return result,OK
		}
	}


	panic(errors.New("Can't find the "+m.MethodName+" method of "+" "+m.ObjectName))
}





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
	Params []string
	Body       []Expr
}


func _MakeObjectMethod(r *runtime.Runtime,params []string,body []Expr) runtime.ObjectMethod {
	return func(o *runtime.Object,in []interface{}) interface{} {
		s:=r.BeginScope()

		s.Set("this",o)

		var result interface{}

		for _, e := range body {
			result, status := e.Eval(r,o)

			switch status {
			case RETURN:
				return result
			}
		}

		r.EndScope(s)

		return result

	}
}

func _MakeClassMethod(r *runtime.Runtime,params []string,body []Expr) runtime.ClassMethod {
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

func (m *MethodDefineExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {
	if m.ObjectName=="" {
		class:=args[0].(*runtime.Class)

		class.SetObjectMembers(m.MethodName, _MakeObjectMethod(r,m.Params,m.Body))
	} else {
		if pc:=r.GetClass(m.ObjectName);pc!=nil {
			pc.SetClassMembers(m.MethodName, _MakeClassMethod(r,m.Params,m.Body))
		}


		if po:=r.GetVarible(m.ObjectName);po!=nil {
			switch v:=po.(type) {
			case *runtime.Object:
				v.SetAttribute(m.MethodName, _MakeObjectMethod(r,m.Params,m.Body))
			case *runtime.Class:
				v.SetClassMembers(m.MethodName, _MakeClassMethod(r,m.Params,m.Body))
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

func (ase * AttributeSetExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {
	pc:=r.GetClass(ase.ObjectName)
	po:=r.GetVarible(ase.ObjectName)

	if pc!=nil && po!=nil {
		panic(errors.New("The name has both a variable and a class"))
	}

	v,status:=ase.ValueExpr.Eval(r)

	if pc!=nil {
		pc.SetClassMembers(ase.AttributeName,v)
	}


	if po!=nil {
		switch lo:=po.(type) {
		case *runtime.Object:
			lo.SetAttribute(ase.AttributeName,v)
		case *runtime.Class: //these lines is for "this" variable in a class define,
			             // when it is ,it is a variable not a class
			lo.SetClassMembers(ase.AttributeName,v)
		}

	}

	return v,status
}


type AttributeExpr struct {
	ObjectName string
	AttributeName string
}

func (ae * AttributeExpr)Eval(r *runtime.Runtime,args ...interface{}) (v interface{},status int) {
	pc:=r.GetClass(ae.ObjectName)
	po:=r.GetVarible(ae.ObjectName)

	if pc!=nil && po!=nil {
		panic(errors.New("The name has both a variable and a class"))
	}

	if pc!=nil {
		a := pc.GetClassMembers(ae.AttributeName)

		if a!=nil {
			return a,OK
		} else {
			panic(errors.New("The "+ae.ObjectName+"class hasn't "+"the "+ae.AttributeName+" attribute"))
		}

	}


	if po!=nil {
		a := po.(*runtime.Object).GetObjectAttribute(ae.AttributeName)

		if a!=nil {
			return a,OK
		} else {
			panic(errors.New("The "+ae.ObjectName+" object hasn't "+"the "+ae.AttributeName+" attribute"))
		}
	}

	return
}


type MethodCalledExpr struct {
	ObjectName string
	MethodName string
	Params     []Expr
}

func (m * MethodCalledExpr)Eval(r *runtime.Runtime,args ...interface{}) (interface{},int) {
	v:=r.GetVarible(m.ObjectName)
	c:=r.FindClass(m.ObjectName)

	if (v!=nil&&c!=nil) {
		panic(errors.New("The name has both a variable and a class"))
	}


	if v!=nil {
		o:= v.(*runtime.Object)

		m:=o.GetObjectAttribute(m.MethodName)

		if m!=nil {
			F:=m.(runtime.ObjectMethod)
			result:=F(o,nil)
			return result,OK
		}
	}


	if c!=nil {
		m:=c.GetClassMembers(m.MethodName)

		if m!=nil {
			F:=m.(runtime.ClassMethod)
			result:=F(c,nil)
			return result,OK
		}
	}


	panic(errors.New("Can't find the "+m.MethodName+" method of "+" "+m.ObjectName))
}





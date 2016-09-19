package ast

import  (
	"reflect"
	"goscript/runtime"
	"errors"
)

type Class struct  {
	Name string
	Exprs []Expr
}


func (c * Class)Eval(r *runtime.Runtime,args ...interface{}) (v reflect.Value,status int) {
	class := runtime.NewClass()

	r.SetClass(c.Name,class)


	for _,e:=range c.Exprs {
		e.Eval(r,class)
	}

	return
}


type ClassAttribute struct  {
	Name string
	Value interface{}
}

func (m * ClassAttribute)Eval(r *runtime.Runtime,args ...interface{}) (v reflect.Value,status int) {
	return
}

type ObjectMethod struct  {
	Name string
	Params []string
	Body  []Expr
}

func (m * ObjectMethod)Eval(r *runtime.Runtime,args ...interface{}) (v reflect.Value,status int) {
	class:=args[0].(*runtime.Class)

	class.SetObjectMembers(m.Name, runtime.ObjectMethod(func(o *runtime.Object,in []interface{}) interface{} {
			return &runtime.Object{}
		}))

	return
}


type ObjectAttribute struct  {
	Name string
	Value interface{}
}

func (m * ObjectAttribute)Eval(r *runtime.Runtime,args ...interface{}) (v reflect.Value,status int) {
	return
}

type MethodCalledExpr struct {
	Name       string
	MethodName string
	Params     []Expr
}

func (m * MethodCalledExpr)Eval(r *runtime.Runtime,args ...interface{}) (v reflect.Value,status int) {
	if v,ok:=r.GetVarible(m.Name);ok {
		o:=v.Interface().(*runtime.Object)

		m:=o.GetAttribute(m.MethodName)

		if m!=nil {
			F:=m.(runtime.ObjectMethod)
			F(o,nil)
			return reflect.ValueOf(1),OK
		}
	}


	if c:=r.FindClass(m.Name);c!=nil {
		m:=c.GetClassMembers(m.MethodName)

		if m!=nil {
			F:=m.(runtime.ClassMethod)
			result:=F(c,nil)
			return reflect.ValueOf(result),OK
		}
	}


	panic(errors.New("Can't find the "+m.MethodName+" method of "+" "+m.Name))
}





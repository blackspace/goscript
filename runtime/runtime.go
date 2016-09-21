package runtime

import (
	. "goscript/mycontainer"
)

type Runtime struct {
	scopes List
	functions map[string]Function
	classes map[string]*Class
}

func NewRuntime() (r *Runtime) {
	r=&Runtime{scopes:List{EqualFun:func(e1 interface{},e2 interface{}) bool {
		return e1.(*Scope)==e2.(*Scope)
	}},

	functions:make(map[string]Function),
	classes:make(map[string]*Class)}

	r.SetFunction("print",Print)

	return r

}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (r *Runtime)BeginScope() (s *Scope){
	ls := NewScope()
	r.scopes.Add(ls)
	return ls
}


func (r *Runtime)EndScope(s *Scope) {
	r.scopes.Remove(s)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (r *Runtime)GetVarible(n string) interface{} {
	e:=r.scopes.FindByLambda(func (e interface{}) bool {
		if _,ok:=e.(*Scope).Get(n);ok {
			return true
		} else {
			return false
		}

	})

	if e==nil {
		return nil
	}

	result,_:=e.(*Scope).Get(n)

	return	result
}

func (r *Runtime)SetVarible(n string,v interface{}) {

	e:=r.scopes.FindByLambda(func (e interface{}) bool {
		if _,ok:=e.(*Scope).Get(n);ok {
			return true
		}
		return false
	})



	//If it didn't find the varible to set in the scope list,must set a new varible in the last scope
	if e==nil {
		r.scopes.Last().(*Scope).Set(n, v)
	} else {
		s:=e.(*Scope)
		s.Set(n,v)
	}


	return
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


func (r*Runtime)SetFunction(n string,f Function){
		r.functions[n]=f

}


func (r*Runtime)GetFunction(n string) Function  {
	f:=r.functions[n]
	return f
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (r *Runtime)SetClass(n string,c *Class) {
	r.classes[n]=c
}

func (r *Runtime)GetClass(n string) *Class {
	if c,ok:=r.classes[n];ok {
		return c
	} else {
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (r *Runtime)FindObject(n string) *Object {
	return nil
}

func (r *Runtime)FindClass(n string) *Class {
	if c,ok:=r.classes[n];ok{
		return c
	} else {
		return nil
	}
}


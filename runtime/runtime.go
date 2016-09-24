package runtime

import (
	. "goscript/mycontainer"
	"errors"
)

type Runtime struct {
	scopes List
	functions map[string]Function
	classes map[string]*Class
	modules List
}

func NewRuntime() (r *Runtime) {
	r=&Runtime{scopes:List{EqualFun:func(e1 interface{},e2 interface{}) bool {
		return e1.(*Scope)==e2.(*Scope)
	}},

	functions:make(map[string]Function),
	classes:make(map[string]*Class),
	modules:List{EqualFun:func(e1 interface{},e2 interface{}) bool {
		return e1.(*Module)==e2.(*Module)
	}}}
	return r

}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Scope
func (r *Runtime)BeginScope() (s *Scope){
	ls := NewScope()
	r.scopes.Add(ls)
	return ls
}


func (r *Runtime)EndScope(s *Scope) {
	r.scopes.Remove(s)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Variable
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
//Function
func (r*Runtime)SetFunction(n string,f Function){
		r.functions[n]=f

}


func (r*Runtime)GetFunction(n string) Function  {
	f:=r.functions[n]
	return f
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Class
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
//Module
func (r *Runtime)LoadModule(m *Module) {
		r.modules.Add(m)
}


func (r *Runtime)GetModule(n string) *Module {
	e:=r.modules.FindByLambda(func (e interface{}) bool {
		if v,ok:=e.(*Module);ok {
			if v.Name==n {
				return true
			} else {
				return false
			}
		} else {
			return false
		}

	})

	if e==nil {
		return nil
	}

	result,_:=e.(*Module)

	return	result
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (r *Runtime)GetChildNode(name string) Node {
	v:=r.GetVarible(name)
	m:=r.GetModule(name)
	c:=r.GetClass(name)

	if v==nil&&m==nil&&c==nil {
		return nil
	}

	i:=0

	if v!=nil {i++}
	if m!=nil {i++}
	if c!=nil {i++}

	if i>1 {
		panic(errors.New("There are too many thing for the name:"+name))
	}

	if v!=nil {
		if lv,ok:=v.(Node);ok {
			return lv
		} else {
			return nil
		}

	}
	if m!=nil {return m}
	if c!=nil {return c}

	return nil
}

func (r *Runtime)GetMember(name string) interface{}  {
	v:=r.GetVarible(name)

	if v!=nil {
		return v
	}

	return nil
}

func (r *Runtime)SetMember(name string,value interface{})  {
	r.SetVarible(name,value)
}

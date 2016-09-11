package runtime

import (
	"reflect"
	. "goscript/mycontainer"
)

type Runtime struct {
	scopes List
}

func NewRuntime() (r *Runtime) {
	return &Runtime{scopes:List{EqualFun:func(e1 interface{},e2 interface{}) bool {
		return e1.(*Scope)==e2.(*Scope)
	}}}
}


func (r *Runtime)BeginScope() (s *Scope){
	ls := NewScope()
	r.scopes.Add(ls)
	return ls
}


func (r *Runtime)EndScope(s *Scope) {
	r.scopes.Remove(s)
}

func (r *Runtime)GetVarible(n string) (v reflect.Value,ok bool) {
	e:=r.scopes.FindByLambda(func (e interface{}) bool {
		if _,ok=e.(*Scope).Get(n);ok {
			return true
		}
		return false
	})

	s:=e.(*Scope)

	v,ok=s.Get(n)

	return
}

func (r *Runtime)SetVarible(n string,v reflect.Value) {

	e:=r.scopes.FindByLambda(func (e interface{}) bool {
		if _,ok:=e.(*Scope).Get(n);ok {
			return true
		}
		return false
	})



	//If it didn't find the varible to set in the scope list,must set a new varible in the last scope
	if e==nil {
		r.scopes.LastNode().Element.(*Scope).Set(n, v)
	} else {
		s:=e.(*Scope)
		s.Set(n,v)
	}


	return
}
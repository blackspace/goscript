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
	for sn:=r.scopes.LastNode();;sn=sn.Pre {

		if v,ok=sn.Element.(*Scope).Get(n);ok {
			return v,ok
		}


		if sn==r.scopes.Root  {
			break
		}
	}

	return
}

func (r *Runtime)SetVarible(n string,v reflect.Value) {
	for sn:=r.scopes.LastNode();;sn=sn.Pre {

		if _,ok:=sn.Element.(*Scope).Get(n);ok {
			sn.Element.(*Scope).Set(n,v)
			return
		}

		if sn==r.scopes.Root  {
			break
		}
	}

	r.scopes.LastNode().Element.(*Scope).Set(n, v)
	return
}
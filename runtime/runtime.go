package runtime

import (
	"reflect"
	. "goscript/mycontainer"
	"errors"
)

type Runtime struct {
	scopes List
	Functions
}

func NewRuntime() (r *Runtime) {
	return &Runtime{scopes:List{EqualFun:func(e1 interface{},e2 interface{}) bool {
		return e1.(*Scope)==e2.(*Scope)
	}},
	Functions: MakeFunctions()}
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
func (r *Runtime)GetVarible(n string) (v reflect.Value,ok bool) {
	e:=r.scopes.FindByLambda(func (e interface{}) bool {
		if _,ok=e.(*Scope).Get(n);ok {
			return true
		} else {
			return false
		}

	})

	if e==nil {
		panic(errors.New("Can't get the varible:"+n))
	}

	return	e.(*Scope).Get(n)
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
		r.scopes.Last().(*Scope).Set(n, v)
	} else {
		s:=e.(*Scope)
		s.Set(n,v)
	}


	return
}


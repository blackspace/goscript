package runtime

import (
	"reflect"
)

type Runtime struct {
	Scopes []*Scope
}


func NewRuntime() (r *Runtime) {
	return &Runtime{Scopes:make([]*Scope,0,1000)}
}



func (r *Runtime)BeginScope() (s *Scope){
	s=NewScope()
	r.Scopes=append(r.Scopes,s)
	return
}


func (r *Runtime)EndScope(s *Scope) {
	for i,ls:=range r.Scopes {
		if ls==s {
			r.Scopes[i]=nil
		}
	}
}

func (r *Runtime)GetVarible(n string) (v reflect.Value,ok bool) {
	for i:=len(r.Scopes)-1;i>=0;i-- {
		s:=r.Scopes[i]
		if s==nil {
			continue
		}

		if v,ok=s.Get(n);ok {
			return v,ok
		}
	}

	return
}

func (r *Runtime)SetVarible(n string,v reflect.Value) {
	for i:=len(r.Scopes)-1;i>=0;i-- {
		s:=r.Scopes[i]
		if s==nil {
			continue
		}

		if _,ok:=s.Get(n);ok {
			s.Set(n,v)
			return
		}
	}

	for i:=len(r.Scopes)-1;i>=0;i-- {
		s := r.Scopes[i]
		if s == nil {
			continue
		}

		s.Set(n, v)
		return
	}



}
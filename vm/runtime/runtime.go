package runtime

import (
	"reflect"
)

type Runtime struct {
	ScopeListHead *ScopeNode
}

type ScopeNode struct {
	pre *ScopeNode
	next *ScopeNode
	*Scope
}

func NewScopeNode() *ScopeNode {
	return &ScopeNode{Scope:NewScope()}
}


func NewRuntime() (r *Runtime) {
	return &Runtime{}
}

func (r *Runtime)LastScopeNode() (sn *ScopeNode) {
	sn=r.ScopeListHead
	for {
		if sn==nil  {
			return sn
		} else if sn.next==nil {
			return sn
		}
		sn=sn.next
	}

	return
}

func (r *Runtime)FindScopeNode(s *Scope) (sn *ScopeNode) {
	for sn=r.ScopeListHead;sn.next!=nil;sn=sn.next {
		if sn.Scope==s {
			return
		}
	}
	return
}



func (r *Runtime)BeginScope() (s *Scope){
	sn:=NewScopeNode()

	if r.ScopeListHead==nil {
		r.ScopeListHead=sn
	} else {
		sn.pre=r.LastScopeNode()
		r.LastScopeNode().next=sn
	}

	return sn.Scope
}


func (r *Runtime)EndScope(s *Scope) {
	n:=r.FindScopeNode(s)

	if n.pre!=nil {
		n.pre.next = n.next
	}

	if n.next!=nil {
		n.next.pre=n.pre
	}
}

func (r *Runtime)GetVarible(n string) (v reflect.Value,ok bool) {
	for sn:=r.LastScopeNode();;sn=sn.pre {

		if v,ok=sn.Get(n);ok {
			return v,ok
		}

		if sn.pre==nil  {
			break
		}
	}

	return
}

func (r *Runtime)SetVarible(n string,v reflect.Value) {
	for sn:=r.LastScopeNode();;sn=sn.pre {

		if _,ok:=sn.Get(n);ok {
			sn.Set(n,v)
			return
		}

		if sn.pre==nil  {
			break
		}
	}

	r.LastScopeNode().Set(n, v)
	return
}
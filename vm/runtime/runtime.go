package runtime

import (
	"reflect"
)

type Runtime struct {
	scopeListHead *_ScopeNode
}

type _ScopeNode struct {
	pre *_ScopeNode
	next *_ScopeNode
	*Scope
}

func _NewScopeNode() *_ScopeNode {
	return &_ScopeNode{Scope:NewScope()}
}


func NewRuntime() (r *Runtime) {
	return &Runtime{}
}

func (r *Runtime)_LastScopeNode() (sn *_ScopeNode) {
	sn=r.scopeListHead
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

func (r *Runtime)_FindScopeNode(s *Scope) (sn *_ScopeNode) {
	for sn=r.scopeListHead;sn.next!=nil;sn=sn.next {
		if sn.Scope==s {
			return
		}
	}
	return
}



func (r *Runtime)BeginScope() (s *Scope){
	sn:= _NewScopeNode()

	if r.scopeListHead ==nil {
		r.scopeListHead =sn
	} else {
		sn.pre=r._LastScopeNode()
		r._LastScopeNode().next=sn
	}

	return sn.Scope
}


func (r *Runtime)EndScope(s *Scope) {
	n:=r._FindScopeNode(s)

	if n.pre!=nil {
		n.pre.next = n.next
	}

	if n.next!=nil {
		n.next.pre=n.pre
	}
}

func (r *Runtime)GetVarible(n string) (v reflect.Value,ok bool) {
	for sn:=r._LastScopeNode();;sn=sn.pre {

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
	for sn:=r._LastScopeNode();;sn=sn.pre {

		if _,ok:=sn.Get(n);ok {
			sn.Set(n,v)
			return
		}

		if sn.pre==nil  {
			break
		}
	}

	r._LastScopeNode().Set(n, v)
	return
}
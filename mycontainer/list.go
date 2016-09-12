package mycontainer


type _Node struct {
	pre     *_Node
	next    *_Node
	element interface{}
}

func _NewNode() *_Node {
	return &_Node{}
}


type List struct {
	Root *_Node
	EqualFun func (interface{},interface{}) bool
}


func NewList() *List {
	return &List{}
}

func (l *List)_FindNode(e interface{}) (n *_Node) {
	for ln:=l.Root;;ln=ln.next {

		if l.EqualFun(ln.element,e) {
			return ln
		}

		if ln.next ==nil {
			break
		}

	}
	return
}

func (l * List)_LastNode() (n *_Node) {
	n=l.Root
	for {
		if n==nil  {
			return n
		} else if n.next ==nil {
			return n
		}
		n=n.next
	}

	return
}


func (l * List)_RemoveNode(n *_Node) {
	if n.pre !=nil {
		n.pre.next = n.next
	} else {
		l.Root=n.next
	}

	if n.next !=nil {
		n.next.pre =n.pre
	}
}


func (l *List)Len() int {
	i:=0
	n:=l.Root
	if n==nil {
		return 0
	}
	for {
		i++
		if n.next ==nil {
			return i
		}
		n=n.next

	}
}

func (l *List)FindByLambda(P func (e interface{}) bool) (e interface{}) {

	if l.Len()==0 {
		return
	}

	for n:=l._LastNode();;n=n.pre {
		if P(n.element) {
			e=n.element
			return e
		}

		if n.pre==nil {
			return nil
		}
	}
	return nil
}



func (l * List)Last() (e interface{}) {
	n:=l.Root
	for {
		if n==nil  {
			return n.element
		} else if n.next ==nil {
			return n.element
		}
		n=n.next
	}

	return n.element
}



func (l *List)Add(e interface{}) {
	n := _NewNode()
	n.element =e
	if l.Root ==nil {
		l.Root = n
	} else {
		n.pre =l._LastNode()
		l._LastNode().next =n
	}
}

func (l *List)Remove(e interface{}) {
	n:=l._FindNode(e)

	if n==nil {
		return
	}

	l._RemoveNode(n)

}

func (l *List)Push(e interface{}) {
	l.Add(e)
}


func (l *List)Pop()  (e interface{}) {
	n:=l._LastNode()

	e=n.element

	l._RemoveNode(n)

	return

}


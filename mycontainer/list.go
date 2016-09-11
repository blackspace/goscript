package mycontainer


type List struct {
	Root *_Node
	EqualFun func (interface{},interface{}) bool
}

type _Node struct {
	pre     *_Node
	next    *_Node
	Element interface{}
}

func NewNode() *_Node {
	return &_Node{}
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

func (l *List)FindByLambda(lambda func (e interface{}) bool) (e interface{}) {

	if l.Len()==0 {
		return
	}

	for n:=l.LastNode();;n=n.pre {
		if lambda(n.Element) {
			e=n.Element
			return e
		}

		if n.pre==nil {
			return nil
		}
	}
	return nil
}

func (l * List)LastNode() (n *_Node) {
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

func (l *List)_FindNode(e interface{}) (n *_Node) {
	for ln:=l.Root;;ln=ln.next {

		if l.EqualFun(ln.Element,e) {
			return ln
		}

		if ln.next ==nil {
			break
		}

	}
	return
}

func (l *List)Add(e interface{}) {
	n :=NewNode()
	n.Element =e
	if l.Root ==nil {
		l.Root = n
	} else {
		n.pre =l.LastNode()
		l.LastNode().next =n
	}
}

func (l *List)Remove(e interface{}) {
	n:=l._FindNode(e)

	if n==nil {
		return
	}

	if n.pre !=nil {
		n.pre.next = n.next
	} else {
		l.Root=n.next
	}

	if n.next !=nil {
		n.next.pre =n.pre
	}
}

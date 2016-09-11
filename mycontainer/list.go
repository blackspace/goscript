package mycontainer

import (
	"log"
)

type List struct {
	Root *Node
	EqualFun func (interface{},interface{}) bool
}

type Node struct {
	Pre *Node
	Next *Node
	Element interface{}
}

func NewNode() *Node {
	return &Node{}
}

func (l *List)Len() int {
	i:=0
	n:=l.Root
	if n==nil {
		return 0
	}
	for {
		i++
		if n.Next==nil {
			return i
		}
		n=n.Next

	}
}

func (l *List)ToString() {
	for n:=l.Root;;n=n.Next {
		log.Println("n:",n)

		if n.Next==nil {
			break
		}
	}
}

func (l * List)LastNode() (n *Node) {
	n=l.Root
	for {
		if n==nil  {
			return n
		} else if n.Next==nil {
			return n
		}
		n=n.Next
	}

	return
}

func (l *List)FindNode(e interface{}) (n *Node) {
	for ln:=l.Root;;ln=ln.Next {

		l.ToString()

		if l.EqualFun(ln.Element,e) {
			return ln
		}

		if ln.Next==nil {
			break
		}

	}
	return
}

func (l *List)Add(e interface{}) {
	n :=NewNode()
	n.Element=e
	if l.Root ==nil {
		l.Root = n
	} else {
		n.Pre=l.LastNode()
		l.LastNode().Next=n
	}
}

func (l *List)Remove(e interface{}) {
	n:=l.FindNode(e)

	if n==nil {

	}

	if n.Pre!=nil {
		n.Pre.Next = n.Next
	}

	if n.Next!=nil {
		n.Next.Pre=n.Pre
	}
}

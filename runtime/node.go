package runtime

import "errors"

type Node interface {
	GetChildNode(string) Node
	GetMember(string) interface{}
	SetMember(string,interface{})
}


type Path struct {
	runtime *Runtime
	path []string
	nodes []Node
}


func NewPath(r *Runtime,path []string) *Path{
	p:=&Path{r,path,make([]Node,0,256)}
	p.nodes=p._ToNodes()

	return p
}

func (p *Path)GetNodes() []Node {
	return p.nodes
}

func (p *Path)_ToNodes() []Node{
	var n Node=p.runtime

	var ns=make([]Node,0,len(p.path))

	for _,p :=range p.path {
		n=n.GetChildNode(p)

		if n==nil {
			break
		}

		ns=append(ns,n)
	}

	if len(ns)<len(p.path)-1 {
		panic(errors.New("This path  is illegal"))
	}

	return ns
}
package runtime

type Node interface {
	GetChildNode(string) Node
	GetMember(string) interface{}
	SetMember(string,interface{})
}


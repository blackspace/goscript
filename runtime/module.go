package runtime

type Module struct {
	Name string
	Members	map[string]interface{}
}

func NewModule(n string) * Module{
	return &Module{Name:n,Members:make(map[string]interface{})}
}

func (m *Module)AddMember(name string,member interface{}) {
	m.Members[name]=member
}

func (m *Module)GetChildNode(name string) Node {
       member:=m.Members[name]

	if n,ok:=member.(Node);ok {
		return n
	} else {
		return nil
	}

	return nil
}

func (m *Module)GetMember(name string) interface{} {
	return m.Members[name]
}

func (m *Module)SetMember(name string,value interface{}) {
	m.Members[name]=value
}
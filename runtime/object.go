package runtime

type ObjectMethod func(*Object,...interface{}) interface{}

type Object struct {
	class   *Class
	Members map[string]interface{}
}

func (o *Object)SetMember(n string,v interface{}) {
	o.Members[n]=v
}

func (o *Object)GetMember(n string) interface{} {
	if a,ok:=o.Members[n];ok {
		return a
	} else {
		return nil
	}
}

func (o *Object)GetChildNode(name string) Node {
	member:=o.GetMember(name)


	if n,ok:=member.(Node);ok {
		return n
	} else {
		return nil
	}
}

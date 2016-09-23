package runtime


type ClassMethod func(*Class, []interface {}) interface{}

type Class struct {
	ObjectMembers         map[string]interface{}
	ClassMembers          map[string]interface{}
}

func NewClass() *Class {
	class := &Class{
		ObjectMembers:make(map[string]interface{}),
		ClassMembers:make(map[string]interface{}),
	}

	class.SetClassMember("new",ClassMethod(func(c *Class,in []interface{}) interface{} {
		o:=class.NewObject()

		if F:=o.GetMember("init");F!=nil {
			initFun:=F.(ObjectMethod)

			initFun(o,nil)
		}

		return o
	}))

	return class
}

func (c *Class)SetObjectMembers(n string,v interface{}) {
	c.ObjectMembers[n]=v
}

func (c *Class)GetObjectMembers(n string) interface {} {
	if m,ok:=c.ObjectMembers[n];ok {
		return m
	} else {
		return nil
	}
}

func (c *Class)SetClassMember(n string,v interface{}) {
	c.ClassMembers[n]=v
}

func (c *Class)GetClassMember(n string) interface {} {
	if m,ok:=c.ClassMembers[n];ok {
		return m
	} else {
		return nil
	}
}


func (c *Class)GetChildNode(name string) Node {
	member:=c.GetClassMember(name)

	if n,ok:=member.(Node);ok {
		return n
	} else {
		return nil
	}
	return nil
}

func (c *Class)GetMember(name string) interface{} {
	return c.GetClassMember(name)
}

func (c *Class)SetMember(name string,value interface{}) {
	c.SetClassMember(name,value)
}

func (c * Class)NewObject() *Object {
	o:=Object{class:c,
		Members:make(map[string]interface{}),
	}

	for n,v:=range c.ObjectMembers {
		o.SetMember(n,v)
	}

	return &o
}





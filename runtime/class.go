package runtime

type ObjectMethod func(*Object,[]interface{}) interface{}
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

	class.SetClassMembers("new",ClassMethod(func(c *Class,in []interface{}) interface{} {
		return class.NewObject()
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

func (c *Class)SetClassMembers(n string,v interface{}) {
	c.ClassMembers[n]=v
}

func (c *Class)GetClassMembers(n string) interface {} {
	if m,ok:=c.ClassMembers[n];ok {
		return m
	} else {
		return nil
	}
}

func (c * Class)NewObject() *Object {
	o:=Object{class:c,
		Attributes:make(map[string]interface{}),
	}

	for n,v:=range c.ObjectMembers {
		o.SetAttribute(n,v)
	}

	return &o
}

type Object struct {
	class        *Class
	Attributes   map[string]interface{}
}

func (o *Object)SetAttribute(n string,v interface{}) {
	o.Attributes[n]=v
}

func (o *Object)GetAttribute(n string) interface{} {
	if a,ok:=o.Attributes[n];ok {
		return a
	} else {
		return nil
	}
}




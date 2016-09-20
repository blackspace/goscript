package runtime


var Cint *Class = NewClass()

func Init() {
	Cint.SetObjectMembers("value",0)
	Cint.SetObjectMembers("+",func(o *Object,in []Object) interface{} {
		return o.GetObjectAttribute("value").(int)+in[0].GetObjectAttribute("value").(int)
	})
	Cint.SetObjectMembers("-",func(o *Object,in []Object) interface{} {
		return o.GetObjectAttribute("value").(int)-in[0].GetObjectAttribute("value").(int)
	})
	Cint.SetObjectMembers("*",func(o *Object,in []Object) interface{} {
		return o.GetObjectAttribute("value").(int)*in[0].GetObjectAttribute("value").(int)
	})
	Cint.SetObjectMembers("/",func(o *Object,in []Object) interface{} {
		return o.GetObjectAttribute("value").(int)/in[0].GetObjectAttribute("value").(int)
	})
}


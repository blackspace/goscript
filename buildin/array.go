package buildin

import (
	"goscript/runtime"
)

var Array * runtime.Class

type LArray []interface{}

func Array_AddElement(o *runtime.Object,vs ...interface{}) interface{} {
	old:=o.GetMember("_").(LArray)


	for _,e:=range vs {
		old=append(old,e)
	}

	o.SetMember("_",old)

	return  o
}

func init() {
	Array=runtime.NewClass()

	Array.SetObjectMembers("_",make(LArray,0,1024))

	Array.SetObjectMembers("Add",runtime.ObjectMethod(Array_AddElement))
}

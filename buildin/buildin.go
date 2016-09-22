package buildin

import (
	"fmt"
	"goscript/runtime"
)

var Console * runtime.Object

func init() {
	c:=runtime.NewClass()

	c.SetObjectMembers("Print",runtime.ObjectMethod(func(r *runtime.Object,args []interface{}) interface{} {
		fmt.Print(args...)
		return nil
	}))


	c.SetObjectMembers("Println",runtime.ObjectMethod(func(r *runtime.Object,args []interface{}) interface{} {
		fmt.Println(args...)
		return nil
	}))

	Console =c.NewObject()
}


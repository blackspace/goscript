package runtime

import (
	"fmt"
)

var console * Object

func init() {
	c:=NewClass()

	c.SetObjectMembers("Print",ObjectMethod(func(r *Object,args []interface{}) interface{} {
		fmt.Print(args...)
		return nil
	}))


	c.SetObjectMembers("Println",ObjectMethod(func(r *Object,args []interface{}) interface{} {
		fmt.Println(args...)
		return nil
	}))

	console =c.NewObject()
}

func Print(in []interface{}) (interface{},int) {
	fmt.Println(in...)
	return nil,0
}

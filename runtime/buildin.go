package runtime

import (
	"fmt"
)

var Console * Object

func init() {
	c:=NewClass()

	c.SetObjectMembers("Print",ObjectMethod(func(r *Object,args []interface{}) interface{} {
		fmt.Print(args...)
		return nil
	}))

	Console=c.NewObject()
}




func Print(in []interface{}) (interface{},int) {
	fmt.Println(in...)
	return nil,0
}

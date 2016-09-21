package runtime

import "fmt"

func Print(in []interface{}) (interface{},int) {
	fmt.Println(in...)
	return nil,0
}

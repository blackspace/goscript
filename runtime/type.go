package runtime

import "reflect"

type Function func(in []reflect.Value) (reflect.Value,int)

type Class struct {
	Attributes map[string]interface{}
	Methods map[string]interface{}
}
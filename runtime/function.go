package runtime

import (
	"reflect"
)

type Function func(in []reflect.Value) (reflect.Value,int)


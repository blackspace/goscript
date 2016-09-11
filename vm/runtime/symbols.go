package runtime

import (
	"reflect"
)

type Symbols map[string]reflect.Value


func NewSymbols() Symbols {
	return make(map[string]reflect.Value)
}

func (s Symbols)Set(n string,v reflect.Value) {
	s[n]=v
}

func (s Symbols)Get(n string) (v reflect.Value,ok bool) {
	v,ok=s[n]

	return
}

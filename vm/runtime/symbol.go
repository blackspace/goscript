package runtime

import "reflect"

type Runtime struct {
	Symbols map[string]reflect.Value
}


func NewRuntime() (r *Runtime) {
	return &Runtime{Symbols:make(map[string]reflect.Value)}
}


func (r *Runtime)AddSymbols(n string,v reflect.Value) {
	r.Symbols[n]=v
}


var Symbols = make(map[string]reflect.Value)

func AddSymbols(n string,v reflect.Value) {
	Symbols[n]=v
}

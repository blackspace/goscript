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


func (r*Runtime)GetSymbolValue(n string) (v reflect.Value,ok bool) {
	v,ok=r.Symbols[n]
	return
}



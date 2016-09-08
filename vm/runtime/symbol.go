package runtime

import "reflect"

var Symbols = make(map[string]reflect.Value)

func AddSymbols(n string,v reflect.Value) {
	Symbols[n]=v
}

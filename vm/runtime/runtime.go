package runtime

type Runtime struct {
	Symbols Symbols
}


func NewRuntime() (r *Runtime) {
	return &Runtime{Symbols:NewSymbols()}
}





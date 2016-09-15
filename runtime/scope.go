package runtime

import "reflect"

type Scope struct {
	Symbols Symbols
}


func NewScope() *Scope {
	return &Scope{Symbols:NewSymbols()}
}

func (s *Scope)Set(n string,v reflect.Value){
	s.Symbols.Set(n,v)
}

func (s *Scope)Get(n string) (v reflect.Value,ok bool) {
	v,ok = s.Symbols.Get(n)

	return
}
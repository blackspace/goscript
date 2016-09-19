package runtime

type Symbols map[string]interface{}


func NewSymbols() Symbols {
	return make(map[string]interface{})
}

func (s Symbols)Set(n string,v interface{}) {
	s[n]=v
}

func (s Symbols)Get(n string) (v interface{},ok bool) {
	v,ok=s[n]

	return
}


type Scope struct {
	Symbols Symbols
}


func NewScope() *Scope {
	return &Scope{Symbols:NewSymbols()}
}

func (s *Scope)Set(n string,v interface{}){
	s.Symbols.Set(n,v)
}

func (s *Scope)Get(n string) (v interface{},ok bool) {
	v,ok = s.Symbols.Get(n)

	return
}
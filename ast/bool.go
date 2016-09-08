package ast

import "reflect"


type Bool struct  {
      Bool bool
}

func (n * Bool)Eval() reflect.Value {
	return reflect.ValueOf(n.Bool)
}


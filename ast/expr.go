package ast

import (
	"reflect"
)

type Expr interface {
      	Eval() reflect.Value
}





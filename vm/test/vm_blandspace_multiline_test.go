package test

import (
	"testing"
	. "goscript/vm"
)
func TestBlankspace(t *testing.T) {
	var vm = NewVM()

	if val,err :=vm.Run("3   +   3"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=6 {
		t.Fail()

	}

}


func TestMultiExprExpr(t *testing.T) {
	var vm = NewVM()

	s:=`
	a=1


	a=a+3

5

a+2
`
	if val,err :=vm.Run(s); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=6 {
		t.Fail()

	}

}

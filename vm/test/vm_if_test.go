package test

import (
	"testing"
	. "goscript/vm"
)

func TestIf(t *testing.T) {
	vm:=NewVM()

	if v,err:=vm.Execute(`if true { 1 }`);err==nil {
		if v.Int()!=1 {
			t.Fail()
		}
	}
}



func TestIfElse(t *testing.T) {
	vm:=NewVM()

	if v,err:=vm.Execute(`if true&&false { 1 } else { 2 }`);err==nil {
		if v.Int()!=2 {
			t.Fail()
		}
	}

	if v,err:=vm.Execute(`if false&&false { 1 } else { 2 3 }`);err==nil {
		if v.Int()!=3 {
			t.Fail()
		}
	}
}

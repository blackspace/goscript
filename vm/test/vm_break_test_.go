package test

import (
	"testing"
	"goscript/vm"
)

func TestBreakBlock(t *testing.T) {
	vm:=vm.NewVM()

	if v,err:=vm.Run("{ 1 2 3  break 4 5 }");err==nil {
		if v.Int()!=3 {
			t.Fail()
		}
	}
}

func TestBreakFor(t *testing.T) {
	vm:=vm.NewVM()

	if v,err:=vm.Run("for a=1;;a++ { if a>10 { break } }");err==nil {
		if v.Int()!=3 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}



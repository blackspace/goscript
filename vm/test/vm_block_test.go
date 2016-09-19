package test

import (
	"testing"
	"goscript/vm"
)

func TestBlock(t *testing.T) {
	vm:=vm.NewVM()

	if v,err:=vm.Run("{ 1 2 3  4 5 }");err==nil {
		if v.(int64)!=5 {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	if v,err:=vm.Run("{  { 1 2 3  4 5 } 1 }");err==nil {
		if v.(int64)!=1 {
			t.Fail()
		}
	} else {
		t.Fail()
	}

}




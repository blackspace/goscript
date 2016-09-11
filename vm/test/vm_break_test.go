package test

import (
	"testing"
	"goscript/vm"
)

func TestBreak(t *testing.T) {
	vm:=vm.NewVM()

	if v,err:=vm.Run("{ 1 2 3  break 4 5 }");err==nil {
		if v.Int()!=5 {
			t.Fail()
		}
	}
}



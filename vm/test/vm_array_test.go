package test

import (
	"testing"
	"goscript/vm"
)

func TestArray(t *testing.T) {
	vm:=vm.NewVM().Init()

	vm.Run("a=[]")

	vm.Run("b=[1,1+2,3+4]")

	v,_:=vm.Run("b[2]")

	if v.(int64)!=7 {
		t.Fail()
	}
}

package test

import (
	"testing"
	"goscript/vm"
)

func TestSimpleFunction(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run("func a() {  a=1 a++ a++  }")

	v,_:=vm.Run("a()")


	if v.Int()!=3 {
		t.Fail()
	}
}


func TestHasParamsFunction(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run("func a(a,b,c) {  b  }")

	v,_:=vm.Run("a(3,4)")


	if v.Int()!=4 {
		t.Fail()
	}
}

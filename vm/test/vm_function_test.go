package test

import (
	"testing"
	"goscript/vm"
)

func TestFunctionSimple(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run("func a() {  a=1 a++ a++  }")

	v,_:=vm.Run("a()")


	if v.Int()!=3 {
		t.Fail()
	}
}


func TestFunctionHasParams(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run("func a(a,b,c) {  b  }")

	v,_:=vm.Run("a(3,4)")


	if v.Int()!=4 {
		t.Fail()
	}
}

func TestFunctionBreak(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run("func a(a,b,c) {  a b break c  }")

	v,_:=vm.Run("a(3,4,5)")

	if v.Int()!=4 {
		t.Fail()
	}
}

func TestFunctionReturn(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run("func a(a,b,c) {  a b return c  }")

	v,_:=vm.Run("a(3,4,5)")

	if v.Int()!=5 {
		t.Fail()
	}
}

func TestFunctionEmbedCall(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run("func a(a,b,c) {  a b return c  }")


	vm.Run("func b() { a(3,4,5)+1 }")

	v,_:=vm.Run("b()")

	if v.Int()!=6 {
		t.Fail()
	}
}




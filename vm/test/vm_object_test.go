package test

import (
	"testing"
	"goscript/vm"
)

func TestObject(t *testing.T) {
	vm:=vm.NewVM().Init()

	vm.Run("a={a:10}")

	vm.Run("a.f=lambda() { this.a}")

	v,_:=vm.Run("a.f()")


	if v.(int64)!=10 {
		t.Fail()
	}


	vm.Run("b={a:1,b:{a:1,b:2},c:[1,3]}")

	v,_=vm.Run("b.b.b")


	if v.(int64)!=2 {
		t.Fail()
	}

	v,_=vm.Run("b.c[1]")


	if v.(int64)!=3 {
		t.Fail()
	}

	vm.Run("b.b.f=lambda() {this.a}")

	v,_=vm.Run("b.b.f()")

	if v.(int64)!=1 {
		t.Fail()
	}
}


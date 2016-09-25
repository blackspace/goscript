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


	vm.Run("b={a:1,b:2,c:3}")

	v,_=vm.Run("b.b")


	if v.(int64)!=2 {
		t.Fail()
	}
}

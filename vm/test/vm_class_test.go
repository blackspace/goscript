package test

import (
	"testing"
	"goscript/vm"
)

func TestClassDefineMethod(t *testing.T) {
	vm:=vm.NewVM()

	v,_:=vm.Run(`class A {
		a=1

		a++

		def hello() { 1 }

	}

	a=A.new()

	a.hello()`)

	if v.Int()!=1 {
		t.Fail()
	}
}

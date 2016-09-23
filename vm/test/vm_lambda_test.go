package test

import (
	"testing"
	"goscript/vm"
)

func TestLambda(t *testing.T) {
	vm:=vm.NewVM().Init()

	v,_:=vm.Run(`

	a=lambda() { 100}

	a()
	`)

	if v.(int64)!=100 {
		t.Fail()
	}

	vm.Run(`
	class A{}

	A.a=lambda() {1000}
	`)



	v,_=vm.Run(`A.a()`)

	if v.(int64)!=1000 {
		t.Fail()
	}


	vm.Run(`
		o=A.new()

		o.a=lambda() {
			10000
		}
	`)

	v,_=vm.Run(`o.a()`)

	if v.(int64)!=10000 {
		t.Fail()
	}

}

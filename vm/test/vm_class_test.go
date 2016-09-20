package test

import (
	"testing"
	"goscript/vm"
)

func TestClassObjectMember(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run(`class A {
		def hello() {
			this.b=2
		}
	}`)

	vm.Run("a=A.new()")


	vm.Run("a.hello()")


	v,_:=vm.Run("a.b+1")

	if v.(int64)!=3 {
		t.Fail()
	}
}

func TestClassAttributeMember(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run(`class A {

		this.a=5

	}`)

	v,_:=vm.Run("A.a")

	if v.(int64)!=5 {
		t.Fail()
	}
}

func TestClassMethodMember(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run(`class A {

		def this.hello() {
			6
		}

	}`)

	v,_:=vm.Run("A.hello()")

	if v.(int64)!=6 {
		t.Fail()
	}
}
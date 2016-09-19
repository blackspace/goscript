package test

import (
	"testing"
	"goscript/vm"
)

func TestClassDefineMethod(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run(`class A {
		def hello() {
			this.b=2
		}
	}`)

	vm.Run("a=A.new()")


	vm.Run("a.hello()")


	v,_:=vm.Run("a.b")

	if v.(int64)!=2 {
		t.Fail()
	}
}

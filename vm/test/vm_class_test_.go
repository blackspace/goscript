package test

import (
	"testing"
	"goscript/vm"
)

func TestClassDefineMethod(t *testing.T) {
	vm:=vm.NewVM()

	vm.Run(`class A {
		def hello() {
			this.b=1
		}
	}`)

	vm.Run("a=A.new()")


	v,_:=vm.Run("a.b")

	if v.(int)!=2 {
		t.Fail()
	}
}

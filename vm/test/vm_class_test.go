package test

import (
	"testing"
	"goscript/vm"
)

func TestClassObjectMember(t *testing.T) {
	vm:=vm.NewVM().Init()

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
	vm:=vm.NewVM().Init()

	vm.Run(`class A {
		A.a=5

		A.b=3
	}`)

	vm.Run("A.c=2")

	v,_:=vm.Run("A.a+A.b+A.c")

	if v.(int64)!=10 {
		t.Fail()
	}
}

func TestClassMethodMember(t *testing.T) {
	vm:=vm.NewVM().Init()

	vm.Run(`class A {
		def A.hello() {
			6
		}
	}`)

	v,_:=vm.Run("A.hello()")

	if v.(int64)!=6 {
		t.Fail()
	}
}

func TestClassReturn(t *testing.T) {
	vm:=vm.NewVM().Init()

	vm.Run(`class A {
		def A.hello() {
			return 6
		}

		def hi() {
			return 7
		}
	}`)

	if v,_:=vm.Run("A.hello()"); v.(int64)!=6 {
		t.Fail()
	}

	vm.Run("a=A.new()")


	if v,_:=vm.Run("a.hi()");v.(int64)!=7 {
		t.Fail()
	}
}

func TestClassInitWithoutParams(t *testing.T) {
	vm:=vm.NewVM().Init()

	vm.Run(`class A {
		def init() {
			this.a=6
			this.b=3
		}
	}`)

	vm.Run("a=A.new()")

	if v,_:=vm.Run("a.a");v.(int64)!=6 {
		t.Fail()
	}

	if v,_:=vm.Run("a.b");v.(int64)!=3 {
		t.Fail()
	}
}
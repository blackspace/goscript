package test

import (
	"testing"
	"goscript/vm"
)

func TestBreakBlock(t *testing.T) {
	vm:=vm.NewVM()

	if v,err:=vm.Run("{ 1 2 3  break 4 5 }");err==nil {
		if v.Int()!=3 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestBreakIF(t *testing.T) {
	vm:=vm.NewVM()

	if v,err:=vm.Run("if true { 1 break }");err==nil {
		if v.IsValid() {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}



func TestBreakIF1(t *testing.T) {
	vm:=vm.NewVM()

	if v,err:=vm.Run("if true { 3 break }");err==nil {
		if v.IsValid() {
			if v.Int()!=3 {
				t.Fail()
			}

		}
	} else {
		t.Fail()
	}
}

func TestBreakFor(t *testing.T) {
	vm:=vm.NewVM()

	if v,err:=vm.Run("for a=1;;a++ { 1 break  }");err==nil {
		if v.IsValid() {
			if v.Int()!=1 {
				t.Fail()
			}

		}
	} else {
		t.Fail()
	}
}




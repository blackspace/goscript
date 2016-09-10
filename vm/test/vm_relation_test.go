package test

import (
	"testing"
	"goscript/vm"
)

func TestRelation(t *testing.T) {
	vm:=vm.NewVM()

	if v,err:=vm.Run("1<2");err==nil {
		if v.Bool()!=true {
			t.Fail()
		}
	}

	if v,err:=vm.Run("2<=2");err==nil {
		if v.Bool()!=true {
			t.Fail()
		}
	}

	if v,err:=vm.Run("3>2");err==nil {
		if v.Bool()!=true {
			t.Fail()
		}
	}

	if v,err:=vm.Run("2<=2");err==nil {
		if v.Bool()!=true {
			t.Fail()
		}
	}
}

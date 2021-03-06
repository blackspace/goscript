package test

import (
	"testing"
	"goscript/vm"
)

func TestRelation(t *testing.T) {
	vm:=vm.NewVM().Init()

	if v,err:=vm.Run("1<2");err==nil {
		if v.(bool)!=true {
			t.Fail()
		}
	}else {
		t.Fatal(err)
	}

	if v,err:=vm.Run("2<=2");err==nil {
		if v.(bool)!=true {
			t.Fail()
		}
	}else {
		t.Fatal(err)
	}

	if v,err:=vm.Run("3>2");err==nil {
		if v.(bool)!=true {
			t.Fail()
		}
	}else {
		t.Fatal(err)
	}

	if v,err:=vm.Run("2<=2");err==nil {
		if v.(bool)!=true {
			t.Fail()
		}
	}else {
		t.Fatal(err)
	}
}

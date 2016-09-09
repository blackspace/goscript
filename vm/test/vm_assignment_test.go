package test

import (
	"testing"
	. "goscript/vm"
)

func TestAsignExpr(t *testing.T) {
	var vm = NewVM()

	if val,err :=vm.Run("a=1"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=1 {
		t.Fail()

	}

	if val,err :=vm.Run("a=1+1*1+1/1"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=3 {
		t.Fail()

	}

	if val,err :=vm.Run("b=3"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=3 {
		t.Fail()

	}

	if val,err :=vm.Run("b+3"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=6 {
		t.Fail()

	}

}


package vm

import "testing"

func TestAsignExpr(t *testing.T) {
	var vm = NewVM()

	if val,err :=vm.Execute("a=1"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=1 {
		t.Fail()

	}

	if val,err :=vm.Execute("a=1+1*1+1/1"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=3 {
		t.Fail()

	}

	if val,err :=vm.Execute("b=3"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=3 {
		t.Fail()

	}

	if val,err :=vm.Execute("b+3"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=6 {
		t.Fail()

	}

}


package vm

import (
	"testing"
)

func TestExecuteSigleDigitNumber(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("1")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=1 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}


func TestExecuteMultiDigitNumber(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("123")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=123 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}
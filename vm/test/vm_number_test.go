package test

import (
	"testing"
	. "goscript/vm"
)

func TestExecuteSigleDigitNumber(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Run("1")

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

	val,err :=vm.Run("123")

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
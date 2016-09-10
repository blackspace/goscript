package test

import (
	"testing"
	. "goscript/vm"
)

func TestDoubleSubExpress(t *testing.T) {
	var vm = NewVM()

	val, err := vm.Run("a=1" )

	if err != nil {
		t.Error(err)
	}

	val, err = vm.Run("a--" )

	if val.IsValid() {
		if val.Int() != 0 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}


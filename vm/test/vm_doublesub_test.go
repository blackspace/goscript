package test

import (
	"testing"
	. "goscript/vm"
)

func TestDoubleSubExpress(t *testing.T) {
	var vm = NewVM()

	val, err := vm.Run("a=2" )

	if err != nil {
		t.Error(err)
	}

	val, err = vm.Run("a--" )

	val, err = vm.Run("a--" )

	if val!=nil {
		if val.(int64) != 0 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}


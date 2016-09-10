package test

import (
	"testing"
	. "goscript/vm"
)

func TestFor(t *testing.T) {
	var vm = NewVM()

	val, err := vm.Run("for a=1;a<10;a++ { a }" )

	if err != nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int() != 9 {
			t.Fail()
		}
	} else {
		t.Fail()
	}

}



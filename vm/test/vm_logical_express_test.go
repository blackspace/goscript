package test

import (
	"testing"
	. "goscript/vm"
)

func TestBool(t *testing.T) {
	var vm = NewVM().Init()

	val,err :=vm.Run("true")

	if err!=nil {
		t.Error(err)
	}

	if val!=nil {
		if val.(bool)!=true {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Run("false")

	if err!=nil {
		t.Error(err)
	}

	if val!=nil {
		if val.(bool)!=false {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestBoolExpress(t *testing.T) {
	var vm = NewVM().Init()

	val,err :=vm.Run("true&&false")

	if err!=nil {
		t.Error(err)
	}

	if val!=nil {
		if val.(bool)!=false {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Run("false||true")

	if err!=nil {
		t.Error(err)
	}

	if val!=nil {
		if val.(bool)!=true {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Run("true&&false||true")

	if err!=nil {
		t.Error(err)
	}

	if val!=nil {
		if val.(bool)!=true {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

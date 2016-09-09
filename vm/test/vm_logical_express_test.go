package test

import (
	"testing"
	. "goscript/vm"
)

func TestBool(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("true")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Bool()!=true {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Execute("false")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Bool()!=false {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestBoolExpress(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("true&&false")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Bool()!=false {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Execute("false||true")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Bool()!=true {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Execute("true&&false||true")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Bool()!=true {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

package test

import (
	"testing"
	. "goscript/vm"
)

func TestString(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute(`"123"`)

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.String()!="123" {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestStringConcate(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute(`"12a3"+"123"`)

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.String()!="12a3123" {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}


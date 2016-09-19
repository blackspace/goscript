package test

import (
	"testing"
	. "goscript/vm"
)

func TestString(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Run(`"123"`)

	if err!=nil {
		t.Error(err)
	}

	if val!=nil {
		if val.(string)!="123" {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestStringConcate(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Run(`"12a3"+"123"`)

	if err!=nil {
		t.Error(err)
	}

	if val!=nil {
		if val.(string)!="12a3123" {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}


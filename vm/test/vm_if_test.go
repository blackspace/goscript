package test

import (
	"testing"
	. "goscript/vm"
)

func TestIf(t *testing.T) {
	vm:=NewVM()

	if v,err:=vm.Run(`if true { 1 }`);err==nil {
		if v.Int()!=1 {
			t.Fail()
		}
	}
}



func TestIfElse(t *testing.T) {
	vm:=NewVM()

	if v,err:=vm.Run(`if true&&true { a=1 a++ a++ } `);err==nil {
		if v.Int()!=3 {
			t.Fail()
		}
	}else {
		t.Fatal(err)
	}

	if v,err:=vm.Run(`if false&&false { 1 } else { 2 3 }`);err==nil {
		if v.Int()!=3 {
			t.Fail()
		}
	}else {
		t.Fatal(err)
	}
}

package test

import (
	"testing"
	"goscript/vm"
)

func TestScope(t *testing.T) {
	vm :=vm.NewVM()

	vm.Run("a=1")

	if v,err:=vm.Run("a++");err==nil {
		if v.Int()!=2 {
			t.Fail()
		}
	}

	v,err:=vm.Run(`
	a=1

	{
		a=2
		a++
	}


	`)


	if err==nil {
		if v.Int()!=3 {
			t.Fail()
		}
	}



	v,err=vm.Run(`
	a=1

	{
		b=2
		b++
	}


	`)


	if err==nil {
		if v.Int()!=3 {
			t.Fail()
		}
	}

}

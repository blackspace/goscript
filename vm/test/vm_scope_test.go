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
	} else {
		t.Fatal(err)
	}


}

func TestScope1(t *testing.T) {
	vm :=vm.NewVM()


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
	} else {
		t.Fatal(err)
	}


}

func TestScope2(t *testing.T) {
	vm :=vm.NewVM()


	s:=`

	a=1


	{

	b=2


	b++

	}

	`

	v,err:=vm.Run(s)


	if err==nil {
		if v.Int()!=3 {
			t.Fail()
		}
	} else {
		t.Fatal(err)
	}

}

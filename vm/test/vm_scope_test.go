package test

import (
	"testing"
	"goscript/vm"
)

func TestScope(t *testing.T) {
	vm :=vm.NewVM().Init()

	vm.Run("a=1")

	if v,err:=vm.Run("a++");err==nil {
		if v.(int64)!=2 {
			t.Fail()
		}
	} else {
		t.Fatal(err)
	}


}

func TestScope1(t *testing.T) {
	vm :=vm.NewVM().Init()


	v,err:=vm.Run(`


	a=1


	{

	a=2

	a++

	}

	`)


	if err==nil {
		if v.(int64)!=3 {
			t.Fail()
		}
	} else {
		t.Fatal(err)
	}


}

func TestScope2(t *testing.T) {
	vm :=vm.NewVM().Init()


	s:=`

	a=1


	{

	b=2


	b++

	}

	`

	v,err:=vm.Run(s)


	if err==nil {
		if v.(int64)!=3 {
			t.Fail()
		}
	} else {
		t.Fatal(err)
	}

}

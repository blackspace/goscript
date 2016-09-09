package test

import (
	"testing"
	. "goscript/vm"
)

func TestExecuteAddExpress(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Run("1+1")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=2 {
			t.Fail()
		}
	} else {
		t.Fail()
	}


	val,err =vm.Run("1+1+1")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=3 {
			t.Fail()
		}
	} else {
		t.Fail()
	}


	val,err =vm.Run("1+123+456")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=580 {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Run("1+123+456+0")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=580 {
			t.Fail()
		}
	} else {
		t.Fail()
	}

}

func TestExecuteSubExpress(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Run("1-1")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=0 {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Run("123-23-1")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=99 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestExecuteMultiExpress(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Run("1*1")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=1 {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Run("1000*10")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=10000 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestExecuteDivExpress(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Run("1/1")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=1 {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Run("1000/10")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=100 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}


func TestExecuteArithmetic(t *testing.T) {
	var vm = NewVM()

	val,err :=vm.Run("1+1/1+1*1")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=3 {
			t.Fail()
		}
	} else {
		t.Fail()
	}

	val,err =vm.Run("10+1000/10-10+1*10+1")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=111 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}


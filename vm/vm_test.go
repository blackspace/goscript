package vm

import (
	"testing"
)

func TestExecuteSigleDigitNumber(t * testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("1")

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
}


func TestExecuteMultiDigitNumber(t * testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("123")

	if err!=nil {
		t.Error(err)
	}

	if val.IsValid() {
		if val.Int()!=123 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestExecuteAddExpress(t * testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("1+1")

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


	val,err =vm.Execute("1+1+1")

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


	val,err =vm.Execute("1+123+456")

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

	val,err =vm.Execute("1+123+456+0")

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

func TestExecuteSubExpress(t * testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("1-1")

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

	val,err =vm.Execute("123-23-1")

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

func TestExecuteMultiExpress(t * testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("1*1")

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

	val,err =vm.Execute("1000*10")

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

func TestExecuteDivExpress(t * testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("1/1")

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

	val,err =vm.Execute("1000/10")

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


func TestExecuteArithmetic(t * testing.T) {
	var vm = NewVM()

	val,err :=vm.Execute("1+1/1+1*1")

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

	val,err =vm.Execute("10+1000/10-10+1*10+1")

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

func TestAsignExpr(t * testing.T) {
	var vm = NewVM()

	if val,err :=vm.Execute("a=1"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=1 {
		t.Fail()

	}

	if val,err :=vm.Execute("a=1+1*1+1/1"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=3 {
		t.Fail()

	}

	if val,err :=vm.Execute("b=3"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=3 {
		t.Fail()

	}

	if val,err :=vm.Execute("b+3"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=6 {
		t.Fail()

	}

}


func TestBlankspace(t * testing.T) {
	var vm = NewVM()

	if val,err :=vm.Execute("3   +   3"); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=6 {
		t.Fail()

	}

}


func TestMultiExprExpr(t * testing.T) {
	var vm = NewVM()

	s:=`
a=1


a+3

5
`
	if val,err :=vm.Execute(s); err!=nil {
		t.Error(err)
	} else if val.IsValid() &&val.Int()!=5 {
		t.Fail()

	}

}


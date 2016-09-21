package test

import (
	"testing"
	"goscript/vm"
)

func TestCommentLine(t *testing.T) {
	vm:=vm.NewVM()

	v,_:=vm.Run(`
	############ddddd
	#sssssss
	#ffffff

	234 #ddddd`)

	if v.(int64)!=234 {
		t.Fail()
	}
}

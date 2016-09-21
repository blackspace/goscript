package test

import (
	"testing"
	"goscript/vm"
)

func TestCommentLine(t *testing.T) {
	vm:=vm.NewVM()

	v,_:=vm.Run(`#! /bin/goscript
	############ddddd
	#sssssss
	#ffffff

	//ddddddddd

	234 #ddddd`)

	if v.(int64)!=234 {
		t.Fail()
	}
}

package test

import (
	"testing"
	"goscript/vm"
)

func TestCommentLine(t *testing.T) {
	vm:=vm.NewVM().Init()

	v,_:=vm.Run(`#! /bin/goscript
	############ddddd
	#sssssss
	#ffffff

	223 //ddddddddd

	/*dddd*/

	234 #ddddd`)

	if v.(int64)!=234 {
		t.Fail()
	}
}

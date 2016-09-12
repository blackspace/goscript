package test

import  (
	"testing"
	. "goscript/mycontainer"
)

func TestPush(t *testing.T) {
	l:=NewList()

	l.Push(1)
	l.Push(2)
	l.Push(3)

	if l.Pop().(int)!=3 {
		t.Fail()
	}

	if l.Pop().(int)!=2 {
		t.Fail()
	}

	if l.Pop().(int)!=1 {
		t.Fail()
	}

}


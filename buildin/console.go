package buildin

import (
	"fmt"
	"goscript/runtime"
)

var Console * runtime.Module

func init() {
	m:=runtime.NewModule("console")

	m.AddMember("Print",runtime.Function(func(args []interface{}) (interface{},int) {
		fmt.Print(args...)
		return nil,0
	}))

	m.AddMember("Println",runtime.Function(func(args []interface{}) (interface{},int) {
		fmt.Println(args...)
		return nil,0
	}))

	Console=m
}


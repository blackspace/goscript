package test

import "goscript/vm"

func ExamplePrint() {
	vm:=vm.NewVM().Init()

	vm.Run(`console.Print("hello goscript")`)
	vm.Run(`console.Println("hello goscript")`)

	//Output:hello goscripthello goscript
}

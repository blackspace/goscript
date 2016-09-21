package test

import "goscript/vm"

func ExamplePrint() {
	vm:=vm.NewVM()

	vm.Run(`print("hello goscript")`)
	vm.Run(`console.Print("hello goscript")`)

	//Output:hello goscript
	//hello goscript
}

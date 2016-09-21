package test

import "goscript/vm"

func ExamplePrint() {
	vm:=vm.NewVM()

	vm.Run(`print("hello goscript")`)

	//Output:hello goscript
}

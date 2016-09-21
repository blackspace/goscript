package test

import "goscript/vm"

func ExamplePrint() {
	vm:=vm.NewVM()

	vm.Run(`Print("hello goscript")`)

	//Output:hello goscript
}

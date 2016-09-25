package buildin

import "goscript/runtime"


var EmptyObject * runtime.Class

func init() {
	EmptyObject=runtime.NewClass()
}



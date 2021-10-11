package screeps

import "syscall/js"

type console struct {
	ref js.Value
}

func (c console) Log(args ...interface{}) {
	c.ref.Call("log", args...)
}

var Console = console{
	ref: js.Global().Get("console"),
}

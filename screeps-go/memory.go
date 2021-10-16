package screeps

import (
	"syscall/js"
)

func Memory() js.Value {
	return js.Global().Get("Memory")
}

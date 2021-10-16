package screeps

import "syscall/js"

var object = js.Global().Get("Object")

var isMMO = !js.Global().Get("Game").Get("cpu").Get("generatePixel").IsUndefined()

func IsMMO() bool {
	return isMMO
}

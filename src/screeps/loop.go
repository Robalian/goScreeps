package screeps

import (
	"syscall/js"
)

//export preMain
func preMain() {
	Game = game{
		ref: js.Global().Get("Game"),
	}
	loadSegments()
}

//export postMain
func postMain() {
	saveSegments()
}

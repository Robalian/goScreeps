package screeps

import (
	"syscall/js"
)

//export preMain
func PreMain() {
	Game = game{
		ref: js.Global().Get("Game"),
	}
	loadSegments()
}

//export postMain
func PostMain() {
	saveSegments()
}

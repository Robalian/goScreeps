package main

import (
	. "screepsgo/screeps-go"
	"syscall/js"
)

func main() {
	channel := make(chan bool)

	runLoop := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		channel <- true
		return nil
	})
	js.Global().Set("runLoop", runLoop)

	for {
		<-channel
		PreMain()

		runCreeps()

		//
		spawnCreeps()

		//
		runTowers()

		//
		pathfinderExample()

		//
		setGrafanaStats()

		PostMain()
	}
}

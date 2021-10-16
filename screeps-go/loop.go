package screeps

import "syscall/js"

var channel chan bool

func InitScreeps() {
	channel = make(chan bool)

	runLoop := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		channel <- true
		return nil
	})
	js.Global().Set("runLoop", runLoop)

	js.Global().Set("goRoomCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		goRoomCallback()
		return nil
	}))

	js.Global().Set("goOrderFilter", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		goOrderFilter()
		return nil
	}))

	js.Global().Set("goRouteCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		goRouteCallback()
		return nil
	}))

	js.Global().Set("goCostCallback", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		goCostCallback()
		return nil
	}))

	Console.Log("Initialized Screeps Go")
}

func PreMain() {
	<-channel
	updateGame()
	updateRawMemory()
	loadSegments()
}

func PostMain() {
	saveSegments()
}

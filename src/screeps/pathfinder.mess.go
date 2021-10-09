package screeps

import "syscall/js"

var roomCallbackArgumentBuffer [20]byte

//export getRoomCallbackArgumentBuffer
func getRoomCallbackArgumentBuffer() *byte {
	return &roomCallbackArgumentBuffer[0]
}

var currentRoomCallback *RoomCallback

//export goRoomCallback
func goRoomCallback() {
	roomName := js.Global().Get("jsRoomCallbackArgument").String()
	var result js.Value
	if currentRoomCallback == nil {
		result = js.ValueOf(NewCostMatrix().ref)
	} else {
		roomCallbackResult := (*currentRoomCallback)(roomName)
		if roomCallbackResult == nil {
			result = js.ValueOf(false)
		} else {
			result = roomCallbackResult.ref
		}
	}

	js.Global().Set("goRoomCallbackResult", result)
}

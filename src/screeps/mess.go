package screeps

import "syscall/js"

var currentRoomCallback *RoomCallback

//export goRoomCallback
func goRoomCallback() {
	roomName := js.Global().Get("roomCallbackArgument").String()
	var result js.Value

	roomCallbackResult := (*currentRoomCallback)(roomName)
	if roomCallbackResult == nil {
		result = js.ValueOf(false)
	} else {
		result = roomCallbackResult.ref
	}

	js.Global().Set("roomCallbackResult", result)
}

var currentOrderFilter *OrderFilter

//export goOrderFilter
func goOrderFilter() bool {
	jsOrder := js.Global().Get("orderFilterArgument")
	order := unpackOrder(jsOrder)
	return (*currentOrderFilter)(order)
}

var currentRouteCallback *RouteCallback

//export goRouteCallback
func goRouteCallback() float64 {
	roomName := js.Global().Get("routeCallbackArgument1").String()
	fromRoomName := js.Global().Get("routeCallbackArgument2").String()
	return (*currentRouteCallback)(roomName, fromRoomName)
}

var currentCostCallback *CostCallback

//export goCostCallback
func goCostCallback() {
	roomName := js.Global().Get("costCallbackArgument1").String()
	costMatrix := CostMatrix{
		ref: js.Global().Get("costCallbackArgument2"),
	}
	(*currentCostCallback)(roomName, &costMatrix)
}

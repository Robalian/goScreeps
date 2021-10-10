package screeps

import "syscall/js"

type RouteCallback func(roomName string, fromRoomName string) float64
type FindRouteOpts struct {
	routeCallback RouteCallback
}

type FindRouteResult []struct {
	exit FindExitConstant
	room string
}

type Map struct {
	ref js.Value
}

func (m Map) DescribeExits(roomName string) map[DirectionConstant]string {
	jsResult := m.ref.Call("describeExits", roomName)
	result := map[DirectionConstant]string{}

	topRoom := jsResult.Get("1")
	if !topRoom.IsUndefined() {
		result[TOP] = topRoom.String()
	}
	rightRoom := jsResult.Get("3")
	if !rightRoom.IsUndefined() {
		result[RIGHT] = rightRoom.String()
	}
	bottomRoom := jsResult.Get("5")
	if !bottomRoom.IsUndefined() {
		result[BOTTOM] = bottomRoom.String()
	}
	leftRoom := jsResult.Get("7")
	if !leftRoom.IsUndefined() {
		result[LEFT] = leftRoom.String()
	}
	return result
}

func (m Map) FindExit(fromRoom string, toRoom string, opts *FindRouteOpts) (*FindExitConstant, ErrorCode) {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		currentRouteCallback = &opts.routeCallback
		jsOpts = js.ValueOf(map[string]interface{}{
			"routeCallback": js.Global().Get("jsRouteCallback"),
		})
	}
	callResult := m.ref.Call("findExit", fromRoom, toRoom, jsOpts).Int()
	if callResult < 0 {
		return nil, ErrorCode(callResult)
	} else {
		result := FindExitConstant(callResult)
		return &result, OK
	}
}

func (m Map) FindRoute(fromRoom string, toRoom string, opts *FindRouteOpts) (*FindRouteResult, ErrorCode) {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		currentRouteCallback = &opts.routeCallback
		jsOpts = js.ValueOf(map[string]interface{}{
			"routeCallback": js.Global().Get("jsRouteCallback"),
		})
	}

	callResult := m.ref.Call("findExit", fromRoom, toRoom, jsOpts)
	if callResult.Type() == js.TypeNumber {
		return nil, ErrorCode(callResult.Int())
	} else {
		routeLength := callResult.Length()
		result := make(FindRouteResult, routeLength)
		for i := 0; i < routeLength; i++ {
			routeStep := callResult.Index(i)
			result[i] = struct {
				exit FindExitConstant
				room string
			}{
				exit: FindExitConstant(routeStep.Get("exit").Int()),
				room: routeStep.Get("room").String(),
			}
		}
		return &result, OK
	}
}

func (m Map) GetRoomLinearDistance(roomName1 string, roomName2 string, continuous bool) int {
	return m.ref.Call("getRoomLinearDistance", roomName1, roomName2, continuous).Int()
}

func (m Map) GetRoomTerrain(roomName string) Terrain {
	jsTerrain := m.ref.Call("getRoomTerrain", roomName)
	return Terrain{ref: jsTerrain}
}

func (m Map) GetWorldSize() int {
	return m.ref.Call("getWorldSize").Int()
}

type RoomStatus struct {
	status    string
	timestamp int
}

func (m Map) GetRoomStatus(roomName string) RoomStatus {
	jsRoomStatus := m.ref.Call("getRoomStatus", roomName)
	return RoomStatus{
		status:    jsRoomStatus.Get("status").String(),
		timestamp: jsRoomStatus.Get("timestamp").Int(),
	}
}

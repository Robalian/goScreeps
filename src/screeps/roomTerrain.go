package screeps

import "syscall/js"

func MakeRoomTerrain(roomName string) RoomTerrain {
	result := js.Global().Get("Room").Get("Terrain").New(roomName)
	return RoomTerrain{
		ref: result,
	}
}

type RoomTerrain struct {
	ref js.Value
}

func (terrain RoomTerrain) Get(x int, y int) int {
	return terrain.ref.Call("get", x, y).Int()
}

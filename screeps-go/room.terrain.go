package screeps

import "syscall/js"

func MakeTerrain(roomName string) Terrain {
	result := js.Global().Get("Room").Get("Terrain").New(roomName)
	return Terrain{
		ref: result,
	}
}

type Terrain struct {
	ref js.Value
}

func (terrain Terrain) Get(x int, y int) int {
	return terrain.ref.Call("get", x, y).Int()
}

//TODO
// func (terrain Terrain) GetRawBuffer(destinationArray *[]uint8) ErrorCode {
//
// }

package screeps

import "syscall/js"

type Room struct {
	ref js.Value
}
type roomConstructor struct {
	ref js.Value
}

var RoomConstructor = roomConstructor{
	ref: js.Global().Get("Room"),
}

func (room Room) Controller() *StructureController {
	var controller = room.ref.Get("controller")
	if controller.IsUndefined() {
		return nil
	} else {
		var result = new(StructureController)
		result.ref = controller
		return result
	}
}

func (room Room) EnergyAvailable() int {
	return room.ref.Get("energyAvailable").Int()
}

func (room Room) EnergyCapacityAvailable() int {
	return room.ref.Get("energyCapacityAvailable").Int()
}

// TODO
//func (room Room) Memory() ??? {
//}

func (room Room) Name() string {
	return room.ref.Get("name").String()
}

func (room Room) Storage() *StructureStorage {
	var storage = room.ref.Get("storage")
	if storage.IsUndefined() {
		return nil
	} else {
		var result = new(StructureStorage)
		result.ref = storage
		return result
	}
}

func (room Room) Terminal() *StructureTerminal {
	var terminal = room.ref.Get("storage")
	if terminal.IsUndefined() {
		return nil
	} else {
		var result = new(StructureTerminal)
		result.ref = terminal
		return result
	}
}

func (room Room) Visual() RoomVisual {
	var visual = room.ref.Get("visual")
	return RoomVisual{
		ref: visual,
	}
}

func (room roomConstructor) SerializePath(path FindPathResult) string {
	var packedPath = packFindPathResult(path)
	return room.ref.Call("serializePath", packedPath).String()
}

func (room roomConstructor) DeserializePath(path string) FindPathResult {
	var deserializedPath = room.ref.Call("deserializePath", path)
	return unpackFindPathResult(deserializedPath)
}

func (room Room) CreateConstructionSite(pos RoomPosition, structureType StructureConstant, name *string) ErrorCode {
	var jsName js.Value
	if name == nil {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(*name)
	}
	var result = room.ref.Call("createConstructionSite", pos.ref, string(structureType), jsName).Int()
	return ErrorCode(result)
}

func (room Room) CreateConstructionSite_XY(x int, y int, structureType StructureConstant, name *string) ErrorCode {
	var jsName js.Value
	if name == nil {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(*name)
	}
	var result = room.ref.Call("createConstructionSite", x, y, string(structureType), jsName).Int()
	return ErrorCode(result)
}

func (room Room) CreateFlag(pos RoomPosition, name *string, color *ColorConstant, secondaryColor *ColorConstant) ErrorCode {
	var jsName js.Value
	if name == nil {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(*name)
	}

	var jsColor js.Value
	if color == nil {
		jsColor = js.Undefined()
	} else {
		jsColor = js.ValueOf(int(*color))
	}

	var jsSecondaryColor js.Value
	if secondaryColor == nil {
		jsSecondaryColor = js.Undefined()
	} else {
		jsSecondaryColor = js.ValueOf(int(*secondaryColor))
	}

	var result = room.ref.Call("createFlag", pos.ref, jsName, jsColor, jsSecondaryColor).Int()
	return ErrorCode(result)
}

func (room Room) CreateFlag_XY(x int, y int, name *string, color *ColorConstant, secondaryColor *ColorConstant) ErrorCode {
	var jsName js.Value
	if name == nil {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(*name)
	}

	var jsColor js.Value
	if color == nil {
		jsColor = js.Undefined()
	} else {
		jsColor = js.ValueOf(int(*color))
	}

	var jsSecondaryColor js.Value
	if secondaryColor == nil {
		jsSecondaryColor = js.Undefined()
	} else {
		jsSecondaryColor = js.ValueOf(int(*secondaryColor))
	}

	var result = room.ref.Call("createFlag", x, y, jsName, jsColor, jsSecondaryColor).Int()
	return ErrorCode(result)
}

func (room Room) Find(findType FindRoomObjectConstant) []RoomObject {
	var foundPositions = room.ref.Call("find", int(findType))
	var length = foundPositions.Length()
	var result = make([]RoomObject, length)
	for i := 0; i < length; i++ {
		result[i] = RoomObject{
			ref: foundPositions.Index(i),
		}
	}
	return result
}

func (room Room) Find_Positions(findType FindExitConstant) []RoomPosition {
	var foundPositions = room.ref.Call("find", int(findType))
	var length = foundPositions.Length()
	var result = make([]RoomPosition, length)
	for i := 0; i < length; i++ {
		result[i] = makeRoomPosition(foundPositions.Index(i))
	}
	return result
}

func (room Room) FindExitTo(otherRoom Room) int {
	return room.ref.Call("findExitTo", otherRoom.ref).Int()
}

func (room Room) FindPath(fromPos RoomPosition, toPos RoomPosition, opts *FindPathOpts) FindPathResult {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(packFindPathOpts(*opts))
	}
	var path = room.ref.Call("findPath", fromPos.ref, toPos.ref, jsOpts)
	return unpackFindPathResult(path)
}

func (room Room) GetEventLog() string {
	return room.ref.Call("getEventLog", true).String()
}

func (room Room) GetPositionAt(x int, y int) RoomPosition {
	var pos = room.ref.Call("getPositionAt", x, y)
	return makeRoomPosition(pos)
}

func (room Room) GetTerrain() RoomTerrain {
	var terrain = room.ref.Call("getTerrain")
	return RoomTerrain{
		ref: terrain,
	}
}

func (room Room) LookAt(pos RoomPosition) LookAtResult {
	var lookAtResult = room.ref.Call("lookAt", pos.ref)
	return unpackLookAtResult(lookAtResult)
}

func (room Room) LookAt_XY(x int, y int) LookAtResult {
	var lookAtResult = room.ref.Call("lookAt", x, y)
	return unpackLookAtResult(lookAtResult)
}

func (room Room) LookAtArea(top int, left int, bottom int, right int) LookAtAreaResult {
	var lookAtAreaResult = room.ref.Call("lookAtArea", top, left, bottom, right, true)

	var length = lookAtAreaResult.Length()
	var result = LookAtAreaResult{}
	for i := 0; i < length; i++ {
		var v = lookAtAreaResult.Index(i)
		var t = v.Get("type").String()
		if t == "terrain" {
			continue
		}

		result = append(result,
			struct {
				X      int
				Y      int
				Type   string
				Object RoomObject
			}{
				X:    v.Get("x").Int(),
				Y:    v.Get("y").Int(),
				Type: t,
				Object: RoomObject{
					ref: v.Get(t),
				},
			})
	}

	return result
}

func (room Room) LookForAt(lookType LookConstant, pos RoomPosition) []RoomObject {
	var lookForAtResult = room.ref.Call("lookForAt", string(lookType), pos.ref)
	return unpackLookForAtResult(lookForAtResult)
}

func (room Room) LookForAt_XY(lookType LookConstant, x int, y int) []RoomObject {
	var lookForAtResult = room.ref.Call("lookForAt", string(lookType), x, y)
	return unpackLookForAtResult(lookForAtResult)
}

func (room Room) LookForAtArea(lookType LookConstant, top int, left int, bottom int, right int) LookForAtAreaResult {
	var lookForAtAreaResult = room.ref.Call("lookForAtArea", string(lookType), top, left, bottom, right, true)

	var length = lookForAtAreaResult.Length()
	var result = LookForAtAreaResult{}
	for i := 0; i < length; i++ {
		var v = lookForAtAreaResult.Index(i)

		result = append(result,
			struct {
				X      int
				Y      int
				Object RoomObject
			}{
				X: v.Get("x").Int(),
				Y: v.Get("y").Int(),
				Object: RoomObject{
					ref: v.Get(string(lookType)),
				},
			})
	}

	return result
}

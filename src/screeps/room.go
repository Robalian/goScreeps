package screeps

import "syscall/js"

type LookAtResult []struct {
	Type   string
	Object RoomObject
}

func unpackLookAtResult(lookAtResult js.Value) LookAtResult {
	length := lookAtResult.Length()
	result := LookAtResult{}
	for i := 0; i < length; i++ {
		v := lookAtResult.Index(i)
		objectType := v.Get("type").String()

		// we skip terrain, for type consistency (it's not a RoomObject) + it's better accessed through Game.map anyway
		if objectType == "terrain" {
			continue
		}

		result = append(result, struct {
			Type   string
			Object RoomObject
		}{
			Type: objectType,
			Object: RoomObject{
				ref: v.Get(objectType),
			},
		})
	}
	return result
}

type LookAtAreaResult []struct {
	X      int
	Y      int
	Type   string
	Object RoomObject
}

type LookForAtAreaResult []struct {
	X      int
	Y      int
	Object RoomObject
}

type FindPathResult []struct {
	x         int
	y         int
	dx        int
	dy        int
	direction DirectionConstant
}

func packFindPathResult(path FindPathResult) []interface{} {
	pathLength := len(path)
	result := make([]interface{}, pathLength)
	for i := 0; i < pathLength; i++ {
		step := map[string]interface{}{}
		step["x"] = path[i].x
		step["y"] = path[i].y
		step["dx"] = path[i].dx
		step["dy"] = path[i].dy
		step["direction"] = int(path[i].direction)
		result[i] = step
	}
	return result
}

func unpackFindPathResult(path js.Value) FindPathResult {
	pathLength := path.Length()
	result := make(FindPathResult, pathLength)
	for i := 0; i < pathLength; i++ {
		step := path.Index(i)
		result[i] = struct {
			x         int
			y         int
			dx        int
			dy        int
			direction DirectionConstant
		}{
			x:         step.Get("x").Int(),
			y:         step.Get("y").Int(),
			dx:        step.Get("dx").Int(),
			dy:        step.Get("dy").Int(),
			direction: DirectionConstant(step.Get("direction").Int()),
		}
	}
	return result
}

type CostCallback func(string, *CostMatrix)

type FindPathOpts struct {
	IgnoreCreeps                 *bool
	IgnoreDestructibleStructures *bool
	IgnoreRoads                  *bool
	CostCallback                 *CostCallback
	MaxOps                       *uint
	HeuristicWeight              *float64
	Serialize                    *bool
	MaxRooms                     *uint
	Range                        *uint
	PlainCost                    *uint
	SwampCost                    *uint
}

func packFindPathOpts(opts FindPathOpts) map[string]interface{} {
	result := map[string]interface{}{}

	if opts.HeuristicWeight != nil {
		result["heuristicWeight"] = *opts.HeuristicWeight
	}
	if opts.IgnoreCreeps != nil {
		result["ignoreCreeps"] = *opts.IgnoreCreeps
	}
	if opts.IgnoreDestructibleStructures != nil {
		result["ignoreDestructibleStructures"] = *opts.IgnoreDestructibleStructures
	}
	if opts.IgnoreRoads != nil {
		result["ignoreRoads"] = *opts.IgnoreRoads
	}
	if opts.CostCallback != nil {
		currentCostCallback = opts.CostCallback
		result["costCallback"] = js.Global().Get("jsCostCallback")
	}
	if opts.MaxOps != nil {
		result["maxOps"] = *opts.MaxOps
	}
	if opts.MaxRooms != nil {
		result["maxRooms"] = *opts.MaxRooms
	}
	if opts.PlainCost != nil {
		result["plainCost"] = *opts.PlainCost
	}
	if opts.Range != nil {
		result["range"] = *opts.Range
	}
	if opts.Serialize != nil {
		result["serialize"] = *opts.Serialize
	}
	if opts.SwampCost != nil {
		result["swampCost"] = *opts.SwampCost
	}

	return result
}

type FindResult struct {
	AsPosition   *[]RoomPosition
	AsRoomObject *[]RoomObject
}

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
	jsController := room.ref.Get("controller")
	if jsController.IsUndefined() {
		return nil
	} else {
		result := new(StructureController)
		result.ref = jsController
		return result
	}
}

func (room Room) EnergyAvailable() int {
	return room.ref.Get("energyAvailable").Int()
}

func (room Room) EnergyCapacityAvailable() int {
	return room.ref.Get("energyCapacityAvailable").Int()
}

func (room Room) Memory() js.Value {
	return room.ref.Get("memory")
}

func (room Room) Name() string {
	return room.ref.Get("name").String()
}

func (room Room) Storage() *StructureStorage {
	jsStorage := room.ref.Get("storage")
	if jsStorage.IsUndefined() {
		return nil
	} else {
		result := new(StructureStorage)
		result.ref = jsStorage
		return result
	}
}

func (room Room) Terminal() *StructureTerminal {
	jsTerminal := room.ref.Get("storage")
	if jsTerminal.IsUndefined() {
		return nil
	} else {
		result := new(StructureTerminal)
		result.ref = jsTerminal
		return result
	}
}

func (room Room) Visual() RoomVisual {
	jsVisual := room.ref.Get("visual")
	return RoomVisual{
		ref: jsVisual,
	}
}

func (room roomConstructor) SerializePath(path FindPathResult) string {
	packedPath := packFindPathResult(path)
	return room.ref.Call("serializePath", packedPath).String()
}

func (room roomConstructor) DeserializePath(path string) FindPathResult {
	deserializedPath := room.ref.Call("deserializePath", path)
	return unpackFindPathResult(deserializedPath)
}

func (room Room) CreateConstructionSite(pos RoomPosition, structureType StructureConstant, name *string) ErrorCode {
	var jsName js.Value
	if name == nil {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(*name)
	}
	result := room.ref.Call("createConstructionSite", pos.ref, string(structureType), jsName).Int()
	return ErrorCode(result)
}

func (room Room) CreateConstructionSite_XY(x int, y int, structureType StructureConstant, name *string) ErrorCode {
	var jsName js.Value
	if name == nil {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(*name)
	}
	result := room.ref.Call("createConstructionSite", x, y, string(structureType), jsName).Int()
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

	result := room.ref.Call("createFlag", pos.ref, jsName, jsColor, jsSecondaryColor).Int()
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

	result := room.ref.Call("createFlag", x, y, jsName, jsColor, jsSecondaryColor).Int()
	return ErrorCode(result)
}

func (room Room) Find(findType FindRoomObjectConstant) []RoomObject {
	foundPositions := room.ref.Call("find", int(findType))
	foundPositionsCount := foundPositions.Length()
	result := make([]RoomObject, foundPositionsCount)
	for i := 0; i < foundPositionsCount; i++ {
		result[i] = RoomObject{
			ref: foundPositions.Index(i),
		}
	}
	return result
}

func (room Room) Find_Positions(findType FindExitConstant) []RoomPosition {
	foundPositions := room.ref.Call("find", int(findType))
	foundPositionsCount := foundPositions.Length()
	result := make([]RoomPosition, foundPositionsCount)
	for i := 0; i < foundPositionsCount; i++ {
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
	path := room.ref.Call("findPath", fromPos.ref, toPos.ref, jsOpts)
	return unpackFindPathResult(path)
}

func (room Room) GetEventLog() string {
	return room.ref.Call("getEventLog", true).String()
}

func (room Room) GetPositionAt(x int, y int) RoomPosition {
	jsPos := room.ref.Call("getPositionAt", x, y)
	return makeRoomPosition(jsPos)
}

func (room Room) GetTerrain() Terrain {
	jsTerrain := room.ref.Call("getTerrain")
	return Terrain{
		ref: jsTerrain,
	}
}

func (room Room) LookAt(pos RoomPosition) LookAtResult {
	lookAtResult := room.ref.Call("lookAt", pos.ref)
	return unpackLookAtResult(lookAtResult)
}

func (room Room) LookAt_XY(x int, y int) LookAtResult {
	lookAtResult := room.ref.Call("lookAt", x, y)
	return unpackLookAtResult(lookAtResult)
}

func (room Room) LookAtArea(top int, left int, bottom int, right int) LookAtAreaResult {
	lookAtAreaResult := room.ref.Call("lookAtArea", top, left, bottom, right, true)

	length := lookAtAreaResult.Length()
	result := LookAtAreaResult{}
	for i := 0; i < length; i++ {
		v := lookAtAreaResult.Index(i)
		objectType := v.Get("type").String()

		// we skip terrain, for type consistency (it's not a RoomObject) + it's better accessed through Game.map anyway
		if objectType == "terrain" {
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
				Type: objectType,
				Object: RoomObject{
					ref: v.Get(objectType),
				},
			})
	}

	return result
}

func (room Room) LookForAt(lookType LookConstant, x int, y int) []RoomObject {
	jsResult := room.ref.Call("lookForAt", string(lookType), x, y)

	length := jsResult.Length()
	result := make([]RoomObject, length)
	for i := 0; i < length; i++ {
		result[i] = RoomObject{
			ref: jsResult.Index(i),
		}
	}
	return result
}

func (room Room) LookForAtArea(lookType LookConstant, top int, left int, bottom int, right int) LookForAtAreaResult {
	lookForAtAreaResult := room.ref.Call("lookForAtArea", string(lookType), top, left, bottom, right, true)

	length := lookForAtAreaResult.Length()
	result := LookForAtAreaResult{}
	for i := 0; i < length; i++ {
		v := lookForAtAreaResult.Index(i)

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

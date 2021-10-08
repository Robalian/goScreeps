package screeps

import "syscall/js"

func packFindClosestByPathOpts(opts FindClosestByPathOpts) map[string]interface{} {
	result := packFindPathOpts(opts.FindPathOpts)
	if opts.Algorithm != nil {
		result["algorithm"] = string(*opts.Algorithm)
	}
	return result
}

func makeRoomPosition(pos js.Value) RoomPosition {
	return RoomPosition{
		ref:      pos,
		X:        pos.Get("x").Int(),
		Y:        pos.Get("y").Int(),
		RoomName: pos.Get("roomName").String(),
	}
}

type RoomPosition struct {
	ref      js.Value
	X        int
	Y        int
	RoomName string
}

func (pos RoomPosition) CreateConstructionSite(structureType StructureConstant, name *string) ErrorCode {
	var jsName js.Value
	if name == nil {
		jsName = js.Undefined()
	} else {
		jsName = js.ValueOf(*name)
	}
	result := pos.ref.Call("createConstructionSite", string(structureType), jsName).Int()
	return ErrorCode(result)
}

func (pos RoomPosition) CreateFlag(name *string, color *ColorConstant, secondaryColor *ColorConstant) ErrorCode {
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

	result := pos.ref.Call("createFlag", jsName, jsColor, jsSecondaryColor).Int()
	return ErrorCode(result)
}

func (pos RoomPosition) FindClosestByPath(findType FindRoomObjectConstant, opts *FindClosestByPathOpts) *RoomObject {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		packedOpts := packFindClosestByPathOpts(*opts)
		jsOpts = js.ValueOf(packedOpts)
	}

	findResult := pos.ref.Call("findClosestByPath", int(findType), jsOpts)
	if findResult.IsNull() {
		return nil
	} else {
		return &RoomObject{
			ref: findResult,
		}
	}
}

func (pos RoomPosition) FindClosestByPath_Exit(findType FindExitConstant, opts *FindClosestByPathOpts) *RoomPosition {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(*opts)
	}

	findResult := pos.ref.Call("findClosestByPath", int(findType), jsOpts)
	if findResult.IsNull() {
		return nil
	} else {
		result := makeRoomPosition(findResult)
		return &result
	}
}

func (pos RoomPosition) FindClosestByPath_Objects(objects []RoomObject, opts *FindClosestByPathOpts) *RoomObject {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(*opts)
	}

	jsObjects := make([]interface{}, len(objects))
	for i, object := range objects {
		jsObjects[i] = object.ref
	}

	findResult := pos.ref.Call("findClosestByPath", jsObjects, jsOpts)
	if findResult.IsNull() {
		return nil
	} else {
		return &RoomObject{
			ref: findResult,
		}
	}
}

func (pos RoomPosition) FindClosestByPath_Positions(positions []RoomPosition, opts *FindClosestByPathOpts) *RoomPosition {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(*opts)
	}

	jsPositions := make([]interface{}, len(positions))
	for i, v := range positions {
		jsPositions[i] = v.ref
	}

	findResult := pos.ref.Call("findClosestByPath", jsPositions, jsOpts)
	if findResult.IsNull() {
		return nil
	} else {
		result := makeRoomPosition(findResult)
		return &result
	}
}

func (pos RoomPosition) FindClosestByRange(findType FindRoomObjectConstant) *RoomObject {
	findResult := pos.ref.Call("findClosestByRange", int(findType))
	if findResult.IsNull() {
		return nil
	} else {
		return &RoomObject{
			ref: findResult,
		}
	}
}

func (pos RoomPosition) FindClosestByRange_Exit(findType FindExitConstant) *RoomPosition {
	findResult := pos.ref.Call("findClosestByRange", int(findType))
	if findResult.IsNull() {
		return nil
	} else {
		result := makeRoomPosition(findResult)
		return &result
	}
}

func (pos RoomPosition) FindClosestByRange_Objects(objects []RoomObject) *RoomObject {
	jsObjects := make([]interface{}, len(objects))
	for i, v := range objects {
		jsObjects[i] = v.ref
	}

	findResult := pos.ref.Call("findClosestByRange", jsObjects)
	if findResult.IsNull() {
		return nil
	} else {
		return &RoomObject{
			ref: findResult,
		}
	}
}

func (pos RoomPosition) FindClosestByRange_Positions(positions []RoomPosition) *RoomPosition {
	jsPositions := make([]interface{}, len(positions))
	for i, v := range positions {
		jsPositions[i] = v.ref
	}

	findResult := pos.ref.Call("findClosestByRange", jsPositions)
	if findResult.IsNull() {
		return nil
	} else {
		result := makeRoomPosition(findResult)
		return &result
	}
}

func (pos RoomPosition) FindInRange(findType FindRoomObjectConstant, distance int) []RoomObject {
	findResult := pos.ref.Call("findInRange", int(findType), distance)
	length := findResult.Length()
	result := make([]RoomObject, length)
	for i := 0; i < length; i++ {
		result[i] = RoomObject{ref: findResult.Index(i)}
	}
	return result
}

func (pos RoomPosition) FindInRange_Exit(findType FindExitConstant, distance int) []RoomPosition {
	findResult := pos.ref.Call("findInRange", int(findType), distance)
	length := findResult.Length()
	result := make([]RoomPosition, length)
	for i := 0; i < length; i++ {
		result[i] = makeRoomPosition(findResult.Index(i))
	}
	return result
}

func (pos RoomPosition) FindInRange_Objects(objects []RoomObject, distance int) []RoomObject {
	jsObjects := make([]interface{}, len(objects))
	for i, v := range objects {
		jsObjects[i] = v.ref
	}

	findResult := pos.ref.Call("findInRange", jsObjects, distance)
	length := findResult.Length()
	result := make([]RoomObject, length)
	for i := 0; i < length; i++ {
		result[i] = RoomObject{ref: findResult.Index(i)}
	}
	return result
}

func (pos RoomPosition) FindInRange_Positions(positions []RoomPosition, distance int) []RoomPosition {
	jsPositions := make([]interface{}, len(positions))
	for i, v := range positions {
		jsPositions[i] = v.ref
	}

	findResult := pos.ref.Call("findInRange", jsPositions, distance)
	length := findResult.Length()
	result := make([]RoomPosition, length)
	for i := 0; i < length; i++ {
		result[i] = makeRoomPosition(findResult.Index(i))
	}
	return result
}

func (pos RoomPosition) FindPathTo(toPos RoomPosition, opts *FindPathOpts) FindPathResult {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(packFindPathOpts(*opts))
	}
	path := pos.ref.Call("findPathTo", toPos.ref, jsOpts)
	return unpackFindPathResult(path)
}

func (pos RoomPosition) FindPathTo_XY(x int, y int, opts *FindPathOpts) FindPathResult {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(packFindPathOpts(*opts))
	}
	path := pos.ref.Call("findPathTo", x, y, jsOpts)
	return unpackFindPathResult(path)
}

func (pos RoomPosition) GetDirectionTo(target RoomPosition) DirectionConstant {
	result := pos.ref.Call("getDirectionTo", target.ref).Int()
	return DirectionConstant(result)
}

func (pos RoomPosition) GetDirectionTo_XY(x int, y int) DirectionConstant {
	result := pos.ref.Call("getDirectionTo", x, y).Int()
	return DirectionConstant(result)
}

func (pos RoomPosition) GetRangeTo(target RoomPosition) int {
	return pos.ref.Call("getRangeTo", target.ref).Int()
}

func (pos RoomPosition) GetRangeTo_XY(x int, y int) int {
	return pos.ref.Call("getRangeTo", x, y).Int()
}

func (pos RoomPosition) InRangeTo(otherPos RoomPosition, distance int) bool {
	return pos.ref.Call("inRangeTo", otherPos.ref, distance).Bool()
}

func (pos RoomPosition) InRangeTo_XY(x int, y int, distance int) bool {
	return pos.ref.Call("inRangeTo", x, y, distance).Bool()
}

func (pos RoomPosition) IsEqualTo(otherPos RoomPosition, distance int) bool {
	return pos.ref.Call("isEqualTo", otherPos.ref, distance).Bool()
}

func (pos RoomPosition) IsEqualTo_XY(x int, y int, distance int) bool {
	return pos.ref.Call("isEqualTo", x, y, distance).Bool()
}

func (pos RoomPosition) IsNearTo(otherPos RoomPosition, distance int) bool {
	return pos.ref.Call("isNearTo", otherPos.ref, distance).Bool()
}

func (pos RoomPosition) IsNearTo_XY(x int, y int, distance int) bool {
	return pos.ref.Call("isNearTo", x, y, distance).Bool()
}

func (pos RoomPosition) Look() LookAtResult {
	lookAtResult := pos.ref.Call("look")
	return unpackLookAtResult(lookAtResult)
}

func (pos RoomPosition) LookFor() []RoomObject {
	lookAtResult := pos.ref.Call("lookFor")
	return unpackLookForAtResult(lookAtResult)
}

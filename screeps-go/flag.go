package screeps

import "syscall/js"

type Flag struct {
	RoomObject
}

func (flag Flag) Color() ColorConstant {
	result := flag.ref.Get("color").Int()
	return ColorConstant(result)
}

func (flag Flag) Memory() js.Value {
	return flag.ref.Get("memory")
}

func (flag Flag) Name() string {
	return flag.ref.Get("name").String()
}

func (flag Flag) SecondaryColor() ColorConstant {
	result := flag.ref.Get("secondaryColor").Int()
	return ColorConstant(result)
}

func (flag Flag) Remove() ErrorCode {
	return ErrorCode(flag.ref.Call("remove").Int())
}

func (flag Flag) SetColor(color ColorConstant, secondaryColor *ColorConstant) ErrorCode {
	if secondaryColor == nil {
		return ErrorCode(flag.ref.Call("setColor", int(color)).Int())
	} else {
		return ErrorCode(flag.ref.Call("setColor", int(color), int(*secondaryColor)).Int())
	}
}

func (flag Flag) SetPosition_XY(x int, y int) ErrorCode {
	return ErrorCode(flag.ref.Call("setPosition", x, y).Int())
}

func (flag Flag) SetPosition_Pos(target RoomPosition) ErrorCode {
	return ErrorCode(flag.ref.Call("setPosition", target.ref).Int())
}

func (flag Flag) SetPosition_HasPos(target HasRoomPosition) ErrorCode {
	return ErrorCode(flag.ref.Call("setPosition", target.Pos().ref).Int())
}

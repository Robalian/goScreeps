package screeps

import "syscall/js"

type RoomObject struct {
	ref js.Value
}

func (roomObject RoomObject) getRef() js.Value {
	return roomObject.ref
}

func (roomObject RoomObject) Effects() []Effect {
	var effects = roomObject.ref.Get("effects")

	var length = effects.Get("length").Int()
	var result = make([]Effect, length)
	for i := 0; i < length; i++ {
		var v = effects.Index(i)
		result[i] = Effect{
			Effect:         EffectId(v.Get("effect").Int()),
			Level:          nil,
			TicksRemaining: v.Get("ticksRemaining").Int(),
		}

		// level
		var level = v.Get("level")
		if !level.IsUndefined() {
			result[i].Level = new(int)
			*result[i].Level = level.Int()
		}
	}

	return result
}

func (roomObject RoomObject) Pos() RoomPosition {
	return makeRoomPosition(roomObject.ref.Get("pos"))
}
func (roomObject RoomObject) Room() Room {
	var roomRef = roomObject.ref.Get("room")
	return Room{
		ref: roomRef,
	}
}

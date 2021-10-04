package screeps

import "syscall/js"

type RoomObject struct {
	ref js.Value
}

func (roomObject RoomObject) getRef() js.Value {
	return roomObject.ref
}

func (roomObject RoomObject) Effects() []Effect {
	jsEffects := roomObject.ref.Get("effects")

	effectCount := jsEffects.Get("length").Int()
	result := make([]Effect, effectCount)
	for i := 0; i < effectCount; i++ {
		effect := jsEffects.Index(i)
		result[i] = Effect{
			Effect:         EffectId(effect.Get("effect").Int()),
			Level:          nil,
			TicksRemaining: effect.Get("ticksRemaining").Int(),
		}

		// level
		level := effect.Get("level")
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
	jsRoom := roomObject.ref.Get("room")
	return Room{
		ref: jsRoom,
	}
}

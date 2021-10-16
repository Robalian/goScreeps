package screeps

type Tombstone struct {
	RoomObject
}

func (tombstone Tombstone) Creep() AnyCreep {
	jsResult := tombstone.ref.Get("creep")
	if jsResult.InstanceOf(creepConstructor) {
		return Creep{
			RoomObject{
				ref: jsResult,
			},
		}
	} else {
		return PowerCreep{
			RoomObject{
				ref: jsResult,
			},
		}
	}
}

func (tombstone Tombstone) DeathTime() int {
	return tombstone.ref.Get("deathTime").Int()
}

func (tombstone Tombstone) Id() string {
	return tombstone.ref.Get("id").String()
}

func (tombstone Tombstone) Store() Store {
	jsStore := tombstone.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (tombstone Tombstone) TicksToDecay() int {
	return tombstone.ref.Get("tickToDecay").Int()
}

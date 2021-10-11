package screeps

type Ruin struct {
	RoomObject
}

func (ruin Ruin) DestroyTime() int {
	return ruin.ref.Get("destroyTime").Int()
}

func (ruin Ruin) Id() string {
	return ruin.ref.Get("id").String()
}

func (ruin Ruin) Store() Store {
	jsStore := ruin.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (ruin Ruin) Structure() Structure {
	jsStructure := ruin.ref.Get("structure")
	return Structure{
		RoomObject{
			ref: jsStructure,
		},
	}
}

func (ruin Ruin) TicksToDecay() int {
	return ruin.ref.Get("tickToDecay").Int()
}

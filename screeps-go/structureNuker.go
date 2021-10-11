package screeps

type StructureNuker struct {
	OwnedStructure
}

func (nuker StructureNuker) Cooldown() int {
	return nuker.ref.Get("cooldown").Int()
}

func (nuker StructureNuker) Store() Store {
	jsStore := nuker.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (nuker StructureNuker) LaunchNuke(pos RoomPosition) ErrorCode {
	result := nuker.ref.Call("launchNuke", pos.ref).Int()
	return ErrorCode(result)
}

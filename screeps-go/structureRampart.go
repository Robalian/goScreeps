package screeps

type StructureRampart struct {
	OwnedStructure
}

func (rampart StructureRampart) IsPublic() bool {
	return rampart.ref.Get("isPublic").Bool()
}

func (rampart StructureRampart) TicksToDecay() int {
	return rampart.ref.Get("tickToDecay").Int()
}

func (rampart StructureRampart) SetPublic(isPublic bool) ErrorCode {
	result := rampart.ref.Call("setPublic", isPublic).Int()
	return ErrorCode(result)
}

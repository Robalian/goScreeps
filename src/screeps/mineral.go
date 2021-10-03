package screeps

type Mineral struct {
	RoomObject
}

func (mineral Mineral) isHarvestable() bool {
	return true
}

func (mineral Mineral) Id() string {
	return mineral.ref.Get("id").String()
}

package screeps

type Mineral struct {
	RoomObject
}

func (mineral Mineral) aAmHarvestable() {}

func (mineral Mineral) Id() string {
	return mineral.ref.Get("id").String()
}

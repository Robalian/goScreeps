package screeps

type Mineral struct {
	RoomObject
}

func (mineral Mineral) iAmHarvestable() {}

func (mineral Mineral) Density() int {
	return mineral.ref.Get("density").Int()
}

func (mineral Mineral) MineralAmount() int {
	return mineral.ref.Get("mineralAmount").Int()
}

func (mineral Mineral) MineralType() ResourceConstant {
	return ResourceConstant(mineral.ref.Get("mineralType").String())
}

func (mineral Mineral) Id() string {
	return mineral.ref.Get("id").String()
}

func (mineral Mineral) TicksToRegeneration() *int {
	jsResult := mineral.ref.Get("ticksToRegeneration")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.Int()
		return &result
	}
}

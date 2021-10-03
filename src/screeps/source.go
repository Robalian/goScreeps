package screeps

type Source struct {
	RoomObject
}

func (source Source) isHarvestable() bool {
	return true
}

func (source Source) Energy() int {
	return source.ref.Get("energy").Int()
}
func (source Source) EnergyCapacity() int {
	return source.ref.Get("energyCapacity").Int()
}
func (source Source) Id() string {
	return source.ref.Get("id").String()
}
func (source Source) TicksToRegeneration() int {
	return source.ref.Get("ticksToRegeneration").Int()
}

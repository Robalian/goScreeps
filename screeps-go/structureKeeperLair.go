package screeps

type StructureKeeperLair struct {
	OwnedStructure
}

func (keeperLair StructureKeeperLair) TicksToSpawn() int {
	return keeperLair.ref.Get("ticksToSpawn").Int()
}

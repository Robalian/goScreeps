package screeps

type StructureRoad struct {
	Structure
}

func (road StructureRoad) TicksToDecay() int {
	return road.ref.Get("tickToDecay").Int()
}

package screeps

type StructureInvaderCore struct {
	OwnedStructure
}

func (invaderCore StructureInvaderCore) Level() int {
	return invaderCore.ref.Get("level").Int()
}

func (invaderCore StructureInvaderCore) TicksToDeploy() *int {
	jsResult := invaderCore.ref.Get("ticksToDeploy")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.Int()
		return &result
	}
}

func (invaderCore StructureInvaderCore) Spawning() *StructureSpawnSpawning {
	jsSpawning := invaderCore.ref.Get("spawning")
	if jsSpawning.IsNull() {
		return nil
	} else {
		result := new(StructureSpawnSpawning)
		result.ref = jsSpawning
		result.Directions = nil
		result.Name = jsSpawning.Get("name").String()
		result.NeedTime = jsSpawning.Get("needTime").Int()
		result.RemainingTime = jsSpawning.Get("remainingTime").Int()
		result.Spawn = nil // TODO - not sure what comes here, documentation doesn't mention a thing

		// jsDirections
		jsDirections := jsSpawning.Get("directions")
		if !jsDirections.IsUndefined() {
			directionsCount := jsDirections.Length()
			result.Directions = make([]DirectionConstant, directionsCount)
			for i := 0; i < directionsCount; i++ {
				result.Directions[i] = DirectionConstant(jsDirections.Index(i).Int())
			}
		}
		return result
	}
}

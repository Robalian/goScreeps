package screeps

type StructurePowerSpawn struct {
	OwnedStructure
}

func (powerSpawn StructurePowerSpawn) iAmRenewingPowerCreeps() {}

func (powerSpawn StructurePowerSpawn) Store() Store {
	jsStore := powerSpawn.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (powerSpawn StructurePowerSpawn) ProcessPower() ErrorCode {
	result := powerSpawn.ref.Call("processPower").Int()
	return ErrorCode(result)
}

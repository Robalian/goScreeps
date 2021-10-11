package screeps

type StructureTower struct {
	OwnedStructure
}

func (tower StructureTower) Store() Store {
	jsStore := tower.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (tower StructureTower) Attack(target Attackable) ErrorCode {
	result := tower.ref.Call("attack", target.getRef()).Int()
	return ErrorCode(result)
}

func (tower StructureTower) Heal(target AnyCreep) ErrorCode {
	result := tower.ref.Call("heal", target.getRef()).Int()
	return ErrorCode(result)
}

func (tower StructureTower) Repair(target Structure) ErrorCode {
	result := tower.ref.Call("repair", target.getRef()).Int()
	return ErrorCode(result)
}

package screeps

type StructureTower struct {
	OwnedStructure
}

func (tower StructureTower) Store() Store {
	var store = tower.ref.Get("store")
	return Store{
		ref: store,
	}
}

func (tower StructureTower) Attack(target Attackable) ErrorCode {
	var result = tower.ref.Call("attack", target.getRef()).Int()
	return ErrorCode(result)
}

func (tower StructureTower) Heal(target AnyCreep) ErrorCode {
	var result = tower.ref.Call("heal", target.getRef()).Int()
	return ErrorCode(result)
}

func (tower StructureTower) Repair(target Structure) ErrorCode {
	var result = tower.ref.Call("repair", target.getRef()).Int()
	return ErrorCode(result)
}

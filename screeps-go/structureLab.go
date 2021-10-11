package screeps

import "syscall/js"

type StructureLab struct {
	OwnedStructure
}

func (lab StructureLab) Cooldown() int {
	return lab.ref.Get("cooldown").Int()
}

func (lab StructureLab) MineralType() ResourceConstant {
	return ResourceConstant(lab.ref.Get("mineralType").String())
}

func (lab StructureLab) Store() Store {
	jsStore := lab.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (lab StructureLab) BoostCreep(creep Creep, bodyPartsCount *int) ErrorCode {
	var jsBodyPartsCount js.Value
	if bodyPartsCount == nil {
		jsBodyPartsCount = js.Undefined()
	} else {
		jsBodyPartsCount = js.ValueOf(*bodyPartsCount)
	}

	result := lab.ref.Call("boostCreep", creep.ref, jsBodyPartsCount).Int()
	return ErrorCode(result)
}

func (lab StructureLab) ReverseReaction(lab1 StructureLab, lab2 StructureLab) ErrorCode {
	result := lab.ref.Call("reverseReaction", lab1.ref, lab2.ref).Int()
	return ErrorCode(result)
}

func (lab StructureLab) RunReaction(lab1 StructureLab, lab2 StructureLab) ErrorCode {
	result := lab.ref.Call("runReaction", lab1.ref, lab2.ref).Int()
	return ErrorCode(result)
}

func (lab StructureLab) UnboostCreep(creep Creep) ErrorCode {
	result := lab.ref.Call("unboostCreep", creep.ref).Int()
	return ErrorCode(result)
}

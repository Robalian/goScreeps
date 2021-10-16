package screeps

import "syscall/js"

type StructureLink struct {
	OwnedStructure
}

func (link StructureLink) Cooldown() int {
	return link.ref.Get("cooldown").Int()
}

func (link StructureLink) Store() Store {
	jsStore := link.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (link StructureLink) TransferEnergy(target StructureLink, amount *int) ErrorCode {
	var jsAmount js.Value
	if amount == nil {
		jsAmount = js.Undefined()
	} else {
		jsAmount = js.ValueOf(amount)
	}

	result := link.ref.Call("transferEnergy", target.ref, jsAmount).Int()
	return ErrorCode(result)
}

package screeps

type Structure struct {
	RoomObject
}

func (structure Structure) Hits() *int {
	jsHits := structure.ref.Get("hits")
	if jsHits.IsUndefined() {
		return nil
	} else {
		result := jsHits.Int()
		return &result
	}
}

func (structure Structure) HitsMax() *int {
	jsHits := structure.ref.Get("hitsMax")
	if jsHits.IsUndefined() {
		return nil
	} else {
		result := jsHits.Int()
		return &result
	}
}

func (structure Structure) Id() string {
	return structure.ref.Get("id").String()
}

func (structure Structure) StructureType() StructureConstant {
	result := structure.ref.Get("structureType").String()
	return StructureConstant(result)
}

func (structure Structure) Destroy() ErrorCode {
	result := structure.ref.Call("destroy").Int()
	return ErrorCode(result)
}

func (structure Structure) IsActive() bool {
	return structure.ref.Call("isActive").Bool()
}

func (structure Structure) NotifyWhenAttacked(enabled bool) ErrorCode {
	result := structure.ref.Call("notifyWhenAttacked", enabled).Int()
	return ErrorCode(result)
}

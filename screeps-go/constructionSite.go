package screeps

type ConstructionSite struct {
	RoomObject
}

func (cs ConstructionSite) Id() string {
	return cs.ref.Get("id").String()
}
func (cs ConstructionSite) My() bool {
	return cs.ref.Get("my").Bool()
}
func (cs ConstructionSite) Owner() struct{ username string } {
	return struct{ username string }{
		username: cs.ref.Get("owner").Get("username").String(),
	}
}
func (cs ConstructionSite) Progress() int {
	return cs.ref.Get("progress").Int()
}
func (cs ConstructionSite) ProgressTotal() int {
	return cs.ref.Get("progressTotal").Int()
}
func (cs ConstructionSite) StructureType() StructureConstant {
	return StructureConstant(cs.ref.Get("structureType").String())
}
func (cs ConstructionSite) Remove() ErrorCode {
	return ErrorCode(cs.ref.Call("remove").Int())
}
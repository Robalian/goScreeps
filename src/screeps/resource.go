package screeps

type Resource struct {
	RoomObject
}

func (resource Resource) Amount() int {
	return resource.ref.Get("amount").Int()
}

func (resource Resource) Id() string {
	return resource.ref.Get("id").String()
}

func (resource Resource) ResourceType() ResourceConstant {
	var result = resource.ref.Get("resourceType").String()
	return ResourceConstant(result)
}

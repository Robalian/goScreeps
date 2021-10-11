package screeps

type StructureFactory struct {
	OwnedStructure
}

func (factory StructureFactory) Cooldown() int {
	return factory.ref.Get("cooldown").Int()
}

func (factory StructureFactory) Level() int {
	return factory.ref.Get("level").Int()
}

func (factory StructureFactory) Store() Store {
	jsStore := factory.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (factory StructureFactory) Produce(resourceType StructureConstant) ErrorCode {
	result := factory.ref.Call("produce", string(resourceType)).Int()
	return ErrorCode(result)
}

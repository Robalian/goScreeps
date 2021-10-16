package screeps

type StructureContainer struct {
	Structure
}

func (container StructureContainer) Store() Store {
	jsStore := container.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (container StructureContainer) TicksToDecay() int {
	return container.ref.Get("tickToDecay").Int()
}

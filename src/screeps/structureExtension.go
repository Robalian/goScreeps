package screeps

type StructureExtension struct {
	OwnedStructure
}

func (extension StructureExtension) Store() Store {
	var store = extension.ref.Get("store")
	return Store{
		ref: store,
	}
}

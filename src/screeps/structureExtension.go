package screeps

type StructureExtension struct {
	OwnedStructure
}

func (extension StructureExtension) Store() Store {
	jsStore := extension.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

package screeps

type StructureStorage struct {
	OwnedStructure
}

func (storage StructureStorage) Store() Store {
	jsStore := storage.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

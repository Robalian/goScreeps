package screeps

type StructureStorage struct {
	OwnedStructure
}

func (storage StructureStorage) Store() Store {
	var store = storage.ref.Get("store")
	return Store{
		ref: store,
	}
}

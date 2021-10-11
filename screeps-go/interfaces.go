package screeps

import "syscall/js"

type Attackable interface {
	getRef() js.Value

	Hits() int
}
type Harvestable interface {
	getRef() js.Value
	iAmHarvestable()

	Pos() RoomPosition
	Room() Room
	Id() string
}

type StructureRenewingPowerCreeps interface {
	getRef() js.Value
	iAmRenewingPowerCreeps()
}

type Transferable interface {
	getRef() js.Value
}

type Withdrawable interface {
	getRef() js.Value
}

type AnyCreep interface {
	getRef() js.Value
	iAmAnyCreep()

	HitsMax() int
	Id() string
}

type HasRoomPosition interface {
	Pos() RoomPosition
}

type SpawnOrExtension interface {
	getRef() js.Value
}

type StoreStructure struct {
	Structure
}

func (storeStructure StoreStructure) Store() Store {
	jsStore := storeStructure.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

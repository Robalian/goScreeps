package screeps

import "syscall/js"

type Attackable interface {
	getRef() js.Value

	Hits() int
}
type Harvestable interface {
	getRef() js.Value
	isHarvestable() bool

	Pos() RoomPosition
	Room() Room
	Id() string
}
type Transferable interface {
	getRef() js.Value
}

type Withdrawable interface {
	getRef() js.Value
}

type AnyCreep interface {
	getRef() js.Value

	Hits()
	HitsMax()
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
	var store = storeStructure.ref.Get("store")
	return Store{
		ref: store,
	}
}

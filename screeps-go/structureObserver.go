package screeps

type StructureObserver struct {
	OwnedStructure
}

func (observer StructureObserver) ObserveRoom(roomName string) ErrorCode {
	result := observer.ref.Call("observeRoom", roomName).Int()
	return ErrorCode(result)
}

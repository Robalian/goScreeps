package screeps

type StructurePortal struct {
	Structure
}

type PortalDestination struct {
	Interroom  *RoomPosition
	Intershard *struct {
		Shard string
		Room  string
	}
}

func (portal StructurePortal) Destination() PortalDestination {
	jsResult := portal.ref.Get("destination")
	isIntershard := !jsResult.InstanceOf(roomPositionConstructor)
	if isIntershard {
		return PortalDestination{
			Interroom: nil,
			Intershard: &struct {
				Shard string
				Room  string
			}{
				Shard: jsResult.Get("shard").String(),
				Room:  jsResult.Get("room").String(),
			},
		}
	} else {
		pos := makeRoomPosition(jsResult)
		return PortalDestination{
			Interroom:  &pos,
			Intershard: nil,
		}
	}
}

func (portal StructurePortal) TicksToDecay() *int {
	jsResult := portal.ref.Get("tickToDecay")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.Int()
		return &result
	}
}

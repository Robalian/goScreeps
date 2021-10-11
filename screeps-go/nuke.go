package screeps

type Nuke struct {
	RoomObject
}

func (nuke Nuke) Id() string {
	return nuke.ref.Get("id").String()
}

func (nuke Nuke) LaunchRoomName() string {
	return nuke.ref.Get("launchRoomName").String()
}

func (nuke Nuke) TimeToLand() int {
	return nuke.ref.Get("timeToLand").Int()
}

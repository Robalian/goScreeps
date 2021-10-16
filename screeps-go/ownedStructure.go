package screeps

type OwnedStructure struct {
	Structure
}

func (ownedStructure OwnedStructure) My() bool {
	return ownedStructure.ref.Get("my").Bool()
}
func (ownedStructure OwnedStructure) Owner() struct{ username string } {
	return struct{ username string }{
		username: ownedStructure.ref.Get("owner").Get("username").String(),
	}
}

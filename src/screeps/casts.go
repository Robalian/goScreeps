package screeps

func (roomObject RoomObject) AsConstructionSite() ConstructionSite {
	return ConstructionSite{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsCreep() Creep {
	return Creep{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsDeposit() Deposit {
	return Deposit{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsFlag() Flag {
	return Flag{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsMineral() Mineral {
	return Mineral{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsNuke() Nuke {
	return Nuke{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsPowerCreep() PowerCreep {
	return PowerCreep{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsResource() Resource {
	return Resource{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsRuin() Ruin {
	return Ruin{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsSource() Source {
	return Source{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsStructure() Structure {
	return Structure{
		RoomObject: roomObject,
	}
}

func (roomObject RoomObject) AsStoreStructure() StoreStructure {
	return StoreStructure{
		Structure{
			RoomObject: roomObject,
		},
	}
}

func (roomObject RoomObject) AsOwnedStructure() OwnedStructure {
	return OwnedStructure{
		Structure{
			RoomObject: roomObject,
		},
	}
}

func (roomObject RoomObject) AsStructureContainer() StructureContainer {
	return StructureContainer{
		Structure{
			RoomObject: roomObject,
		},
	}
}

func (roomObject RoomObject) AsStructureController() StructureController {
	return StructureController{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureExtension() StructureExtension {
	return StructureExtension{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureExtractor() StructureExtractor {
	return StructureExtractor{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureFactory() StructureFactory {
	return StructureFactory{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureInvaderCore() StructureInvaderCore {
	return StructureInvaderCore{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureKeeperLair() StructureKeeperLair {
	return StructureKeeperLair{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureLab() StructureLab {
	return StructureLab{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureLink() StructureLink {
	return StructureLink{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureNuker() StructureNuker {
	return StructureNuker{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureObserver() StructureObserver {
	return StructureObserver{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructurePortal() StructurePortal {
	return StructurePortal{
		Structure{
			RoomObject: roomObject,
		},
	}
}

func (roomObject RoomObject) AsStructurePowerBank() StructurePowerBank {
	return StructurePowerBank{
		Structure{
			RoomObject: roomObject,
		},
	}
}

func (roomObject RoomObject) AsStructurePowerSpawn() StructurePowerSpawn {
	return StructurePowerSpawn{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureRampart() StructureRampart {
	return StructureRampart{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureRoad() StructureRoad {
	return StructureRoad{
		Structure{
			RoomObject: roomObject,
		},
	}
}

func (roomObject RoomObject) AsStructureSpawn() StructureSpawn {
	return StructureSpawn{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureStorage() StructureStorage {
	return StructureStorage{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureTerminal() StructureTerminal {
	return StructureTerminal{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureTower() StructureTower {
	return StructureTower{
		OwnedStructure{
			Structure{
				RoomObject: roomObject,
			},
		},
	}
}

func (roomObject RoomObject) AsStructureWall() StructureWall {
	return StructureWall{
		Structure{
			RoomObject: roomObject,
		},
	}
}

func (roomObject RoomObject) AsTombstone() Tombstone {
	return Tombstone{
		RoomObject: roomObject,
	}
}

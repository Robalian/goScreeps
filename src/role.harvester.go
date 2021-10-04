package main

import . "screepsgo/src/screeps"

func roleHarvester(creep Creep) {
	if *creep.Store().GetFreeCapacity(nil) > 0 {
		source := creep.Pos().FindClosestByPath(FIND_SOURCES_ACTIVE, nil)
		if source != nil {
			if creep.Harvest(Source{RoomObject: *source}) == ERR_NOT_IN_RANGE { // TODO
				var color = "#ffaa00" // TODO
				creep.MoveTo(source.Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
			}
		}
	} else {
		targets := creep.Room().Find(FIND_STRUCTURES)
		validTargets := []Structure{}
		for _, v := range targets {
			structure := Structure{RoomObject: v} // TODO
			if structure.StructureType() == STRUCTURE_EXTENSION ||
				structure.StructureType() == STRUCTURE_SPAWN ||
				structure.StructureType() == STRUCTURE_TOWER {
				storeStructure := StoreStructure{Structure: structure} // TODO
				if *storeStructure.Store().GetFreeCapacity(&RESOURCE_ENERGY) > 0 {
					validTargets = append(validTargets, structure)
				}
			}
		}

		if len(validTargets) > 0 {
			var color = "#ffffff" // TODO
			if creep.Transfer(validTargets[0], RESOURCE_ENERGY, nil) == ERR_NOT_IN_RANGE {
				creep.MoveTo(validTargets[0].Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
			}
		} else {
			creep.Say("???", false)
		}
	}
}

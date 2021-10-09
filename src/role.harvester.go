package main

import . "screepsgo/src/screeps"

func roleHarvester(creep Creep) {
	if *creep.Store().GetFreeCapacity(nil) > 0 {
		source := creep.Pos().FindClosestByPath(FIND_SOURCES_ACTIVE, nil)
		if source != nil {
			if creep.Harvest(source.AsSource()) == ERR_NOT_IN_RANGE {
				color := "#ffaa00" // TODO
				creep.MoveTo(source.Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
			}
		}
	} else {
		targets := creep.Room().Find(FIND_STRUCTURES)
		validTargets := []Structure{}
		for _, v := range targets {
			structure := v.AsStructure()
			if structure.StructureType() == STRUCTURE_EXTENSION ||
				structure.StructureType() == STRUCTURE_SPAWN ||
				structure.StructureType() == STRUCTURE_TOWER {
				storeStructure := structure.AsStoreStructure()
				if *storeStructure.Store().GetFreeCapacity(&RESOURCE_ENERGY) > 0 {
					validTargets = append(validTargets, structure)
				}
			}
		}

		if len(validTargets) > 0 {
			color := "#ffffff" // TODO
			if creep.Transfer(validTargets[0], RESOURCE_ENERGY, nil) == ERR_NOT_IN_RANGE {
				creep.MoveTo(validTargets[0].Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
			}
		} else {
			creep.Say("???", false)
		}
	}
}

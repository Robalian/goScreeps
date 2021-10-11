package main

import . "screepsgo/screeps-go"

func roleUpgrader(creep Creep) {
	if creep.Memory().Get("upgrading").Truthy() && *creep.Store().GetUsedCapacity(&RESOURCE_ENERGY) == 0 {
		creep.Memory().Set("upgrading", false)
		creep.Say("🔄 harvest", false)
	}
	if !creep.Memory().Get("upgrading").Truthy() && *creep.Store().GetFreeCapacity(nil) == 0 {
		creep.Memory().Set("upgrading", true)
		creep.Say("⚡ upgrade", false)
	}

	if creep.Memory().Get("upgrading").Truthy() {
		color := "#ffffff"
		if creep.UpgradeController(*creep.Room().Controller()) == ERR_NOT_IN_RANGE {
			creep.MoveTo(creep.Room().Controller().Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
		}
	} else {
		color := "#ffaa00"
		source := creep.Pos().FindClosestByPath(FIND_SOURCES_ACTIVE, nil)
		if source != nil {
			if creep.Harvest(source.AsSource()) == ERR_NOT_IN_RANGE {
				creep.MoveTo(source.Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
			}
		}
	}
}

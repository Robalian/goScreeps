package main

import . "screepsgo/src/screeps"

func roleBuilder(creep Creep) {
	if creep.Memory().Get("building").Truthy() && *creep.Store().GetUsedCapacity(&RESOURCE_ENERGY) == 0 {
		creep.Memory().Set("building", false)
		creep.Say("🔄 harvest", false)
	}
	if !creep.Memory().Get("building").Truthy() && *creep.Store().GetFreeCapacity(nil) == 0 {
		creep.Memory().Set("building", true)
		creep.Say("🚧 build", false)
	}

	if creep.Memory().Get("building").Truthy() {
		closestConstructionSite := creep.Pos().FindClosestByPath(FIND_CONSTRUCTION_SITES, nil)
		if closestConstructionSite != nil {
			constructionSite := ConstructionSite{RoomObject: *closestConstructionSite} // TODO
			if creep.Build(constructionSite) == ERR_NOT_IN_RANGE {
				var color = "#ffffff" // TODO
				creep.MoveTo(constructionSite.Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
			}
		}
	} else {
		source := creep.Pos().FindClosestByPath(FIND_SOURCES_ACTIVE, nil)
		if source != nil {
			if creep.Harvest(Source{RoomObject: *source}) == ERR_NOT_IN_RANGE { // TODO
				var color = "#ffaa00" // TODO
				creep.MoveTo(source.Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
			}
		}
	}
}

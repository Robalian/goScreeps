package main

import . "screepsgo/screeps-go"

func roleHarvester(creep Creep) {
	if *creep.Store().GetFreeCapacity(nil) > 0 {
		source := creep.Pos().FindClosestByPath(FIND_SOURCES_ACTIVE, nil)
		if source != nil {
			if creep.Harvest(source.AsSource()) == ERR_NOT_IN_RANGE {
				color := "#ffaa00"
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
			color := "#ffffff"
			if creep.Transfer(validTargets[0], RESOURCE_ENERGY, nil) == ERR_NOT_IN_RANGE {
				creep.MoveTo(validTargets[0].Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
			}
		} else {
			creep.Say("???", false)
		}
	}
}

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
			constructionSite := closestConstructionSite.AsConstructionSite()
			if creep.Build(constructionSite) == ERR_NOT_IN_RANGE {
				color := "#ffffff"
				creep.MoveTo(constructionSite.Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
			}
		}
	} else {
		source := creep.Pos().FindClosestByPath(FIND_SOURCES_ACTIVE, nil)
		if source != nil {
			if creep.Harvest(source.AsSource()) == ERR_NOT_IN_RANGE {
				color := "#ffaa00"
				creep.MoveTo(source.Pos(), &MoveToOpts{VisualizePathStyle: &PolyStyle{Stroke: &color}})
			}
		}
	}
}

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

func runCreeps() {
	for _, creep := range Game.Creeps() {
		creepRole := creep.Memory().Get("role")
		if creepRole.IsUndefined() {
			continue
		}

		if creep.Memory().Get("role").String() == "harvester" {
			roleHarvester(creep)
		}
		if creep.Memory().Get("role").String() == "upgrader" {
			roleUpgrader(creep)
		}
		if creep.Memory().Get("role").String() == "builder" {
			roleBuilder(creep)
		}
	}
}

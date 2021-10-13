package main

import . "screepsgo/screeps-go"

func runTowers() {

	room := Game.Spawns()["Spawn1"].Room()
	myStructures := room.Find(FIND_MY_STRUCTURES)
	var tower *StructureTower
	for _, v := range myStructures {
		structure := v.AsStructure()
		if structure.StructureType() == STRUCTURE_TOWER {
			asTower := structure.AsStructureTower() // TODO
			tower = &asTower
			break
		}
	}
	if tower != nil {
		structures := tower.Room().Find(FIND_STRUCTURES)
		damagedStructures := []RoomObject{} // TODO
		for _, v := range structures {
			structure := v.AsStructure()
			if structure.Hits() != nil && *structure.Hits() < *structure.HitsMax() {
				damagedStructures = append(damagedStructures, v)
			}
		}

		closestDamagedStructure := tower.Pos().FindClosestByRange_RoomObjects(damagedStructures)
		if closestDamagedStructure != nil {
			tower.Repair(closestDamagedStructure.AsStructure())
		}

		closestHostile := tower.Pos().FindClosestByRange_ObjectConstant(FIND_HOSTILE_CREEPS)
		if closestHostile != nil {
			tower.Attack(Creep{RoomObject: *closestHostile})
		}
	}
}

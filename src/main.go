package main

import (
	. "screepsgo/src/screeps"
	"strconv"
	"syscall/js"
)

//export Loop
func Loop() {
	spawn := Game.Spawns()["Spawn1"]
	room := spawn.Room()

	myStructures := room.Find(FIND_MY_STRUCTURES)
	var tower *StructureTower
	for _, v := range myStructures {
		structure := Structure{RoomObject: v} // TODO
		if structure.StructureType() == STRUCTURE_TOWER {
			tower = &StructureTower{OwnedStructure: OwnedStructure{Structure: structure}} // TODO
			break
		}
	}
	if tower != nil {
		structures := tower.Room().Find(FIND_STRUCTURES)
		damagedStructures := []RoomObject{} // TODO
		for _, v := range structures {
			structure := Structure{RoomObject: v} // TODO
			if structure.Hits() != nil && *structure.Hits() < *structure.HitsMax() {
				damagedStructures = append(damagedStructures, v)
			}
		}

		closestDamagedStructure := tower.Pos().FindClosestByRange_Objects(damagedStructures)
		if closestDamagedStructure != nil {
			tower.Repair(Structure{RoomObject: *closestDamagedStructure}) // TODO
		}

		closestHostile := tower.Pos().FindClosestByRange(FIND_HOSTILE_CREEPS)
		if closestHostile != nil {
			tower.Attack(Creep{RoomObject: *closestHostile})
		}
	}

	harvestersCount := 0
	upgradersCount := 0
	buildersCount := 0
	for _, creep := range Game.Creeps() {
		creepRole := creep.Memory().Get("role")
		if creepRole.IsUndefined() {
			continue
		}

		if creep.Memory().Get("role").String() == "harvester" {
			roleHarvester(creep)
			harvestersCount++
		}
		if creep.Memory().Get("role").String() == "upgrader" {
			roleUpgrader(creep)
			upgradersCount++
		}
		if creep.Memory().Get("role").String() == "builder" {
			roleBuilder(creep)
			buildersCount++
		}
	}

	if spawn.Spawning() == nil {
		if harvestersCount < 2 {
			creepBody := []BodyPartConstant{MOVE, MOVE, WORK, CARRY, CARRY}
			creepName := "harvester" + strconv.Itoa(Game.Time())
			creepMemory := js.ValueOf(map[string]interface{}{
				"role": "harvester",
			})
			spawn.SpawnCreep(creepBody, creepName, &SpawnOpts{Memory: &creepMemory})
		} else if upgradersCount < 3 {
			creepBody := []BodyPartConstant{MOVE, MOVE, WORK, CARRY, CARRY}
			creepName := "upgrader" + strconv.Itoa(Game.Time())
			creepMemory := js.ValueOf(map[string]interface{}{
				"role": "upgrader",
			})
			spawn.SpawnCreep(creepBody, creepName, &SpawnOpts{Memory: &creepMemory})
		} else if buildersCount < 3 {
			creepBody := []BodyPartConstant{MOVE, MOVE, WORK, CARRY, CARRY}
			creepName := "builder" + strconv.Itoa(Game.Time())
			creepMemory := js.ValueOf(map[string]interface{}{
				"role": "builder",
			})
			spawn.SpawnCreep(creepBody, creepName, &SpawnOpts{Memory: &creepMemory})
		}
	}
}

func main() {
}

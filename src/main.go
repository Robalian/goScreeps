package main

import (
	. "screepsgo/src/screeps"
	"strconv"
	"syscall/js"
)

//export Loop
func Loop() {
	var spawn = Game.Spawns()["Spawn1"]
	var room = spawn.Room()

	var myStructures = room.Find(FIND_MY_STRUCTURES)
	var tower *StructureTower = nil
	for _, v := range myStructures {
		var structure = Structure{RoomObject: v} // TODO
		if structure.StructureType() == STRUCTURE_TOWER {
			tower = &StructureTower{OwnedStructure: OwnedStructure{Structure: structure}} // TODO
			break
		}
	}
	if tower != nil {
		var structures = tower.Room().Find(FIND_STRUCTURES)
		var damagedStructures = []RoomObject{} // TODO
		for _, v := range structures {
			var structure = Structure{RoomObject: v} // TODO
			if structure.Hits() != nil && *structure.Hits() < *structure.HitsMax() {
				damagedStructures = append(damagedStructures, v)
			}
		}

		var closestDamagedStructure = tower.Pos().FindClosestByRange_Objects(damagedStructures)
		if closestDamagedStructure != nil {
			tower.Repair(Structure{RoomObject: *closestDamagedStructure}) // TODO
		}

		var closestHostile = tower.Pos().FindClosestByRange(FIND_HOSTILE_CREEPS)
		if closestHostile != nil {
			tower.Attack(Creep{RoomObject: *closestHostile})
		}
	}

	var harvesters = 0
	var upgraders = 0
	var builders = 0
	for _, creep := range Game.Creeps() {
		var role = creep.Memory().Get("role")
		if role.IsUndefined() {
			continue
		}

		if creep.Memory().Get("role").String() == "harvester" {
			roleHarvester(creep)
			harvesters++
		}
		if creep.Memory().Get("role").String() == "upgrader" {
			roleUpgrader(creep)
			upgraders++
		}
		if creep.Memory().Get("role").String() == "builder" {
			roleBuilder(creep)
			builders++
		}
	}

	if spawn.Spawning() == nil {
		if harvesters < 2 {
			var creepBody = []BodyPartConstant{MOVE, MOVE, WORK, CARRY, CARRY}
			var creepName = "harvester" + strconv.Itoa(Game.Time())
			var creepMemory = js.ValueOf(map[string]interface{}{
				"role": "harvester",
			})
			spawn.SpawnCreep(creepBody, creepName, &SpawnOpts{Memory: &creepMemory})
		} else if upgraders < 3 {
			var creepBody = []BodyPartConstant{MOVE, MOVE, WORK, CARRY, CARRY}
			var creepName = "upgrader" + strconv.Itoa(Game.Time())
			var creepMemory = js.ValueOf(map[string]interface{}{
				"role": "upgrader",
			})
			spawn.SpawnCreep(creepBody, creepName, &SpawnOpts{Memory: &creepMemory})
		} else if builders < 3 {
			var creepBody = []BodyPartConstant{MOVE, MOVE, WORK, CARRY, CARRY}
			var creepName = "builder" + strconv.Itoa(Game.Time())
			var creepMemory = js.ValueOf(map[string]interface{}{
				"role": "builder",
			})
			spawn.SpawnCreep(creepBody, creepName, &SpawnOpts{Memory: &creepMemory})
		}
	}
}

func main() {
}

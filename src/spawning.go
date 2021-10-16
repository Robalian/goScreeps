package main

import (
	. "screepsgo/screeps-go"
	"strconv"
	"syscall/js"
)

func spawnCreeps() {
	harvestersCount := 0
	upgradersCount := 0
	buildersCount := 0

	// count creeps by role
	for _, creep := range Game.Creeps() {
		creepRole := creep.Memory().Get("role")
		if creepRole.IsUndefined() {
			continue
		}

		if creep.Memory().Get("role").String() == "harvester" {
			harvestersCount++
		}
		if creep.Memory().Get("role").String() == "upgrader" {
			upgradersCount++
		}
		if creep.Memory().Get("role").String() == "builder" {
			buildersCount++
		}
	}

	// spawn
	spawn := Game.Spawns()["Spawn1"]
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

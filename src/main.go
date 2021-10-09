package main

import (
	. "screepsgo/src/screeps"
	"strconv"
	"syscall/js"
)

func main() {
	PreMain()

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

	//
	flag1, ok1 := Game.Flags()["Flag1"]
	flag2, ok2 := Game.Flags()["Flag2"]
	if ok1 && ok2 {
		goals := []PathFinderGoal{
			{
				Pos:   flag2.Pos(),
				Range: 0,
			},
		}
		var roomCb RoomCallback = func(roomName string) *CostMatrix {
			result := NewCostMatrix()

			flags := Game.Flags()
			for flagName, flag := range flags {
				if flag.Color() == COLOR_RED {
					result.Set(flag.Pos().X, flag.Pos().Y, 254)
					Console.Log(roomName, flagName, flag.Pos().X, flag.Pos().Y, result.Get(flag.Pos().X, flag.Pos().Y))
				}
			}

			return &result
		}
		opts := PathFinderOpts{
			RoomCallback: &roomCb,
		}
		path := PathFinder.Search(flag1.Pos(), goals, &opts)
		if path.Incomplete {
			Console.Log("Flag1-Flag2 Path incomplete")
		} else {
			//Console.Log("Drawing path", "length =", len(path.Path))
			visual := spawn.Room().Visual()
			visual.Poly(path.Path, nil)
		}
	}
	PostMain()
}

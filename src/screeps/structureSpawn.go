package screeps

import "syscall/js"

type SpawnOpts struct {
	Memory           *js.Value
	EnergyStructures *[]SpawnOrExtension
	DryRun           *bool
	Directions       *[]DirectionConstant
}

type StructureSpawnSpawning struct {
	ref           js.Value
	Directions    []DirectionConstant
	Name          string
	NeedTime      int
	RemainingTime int
	Spawn         *StructureSpawn
}

func (spawning StructureSpawnSpawning) Cancel() ErrorCode {
	var result = spawning.ref.Call("cancel").Int()
	return ErrorCode(result)
}

func (spawning StructureSpawnSpawning) SetDirections(directions []DirectionConstant) ErrorCode {
	var packedDirections = make([]interface{}, len(directions))
	for i, v := range directions {
		packedDirections[i] = int(v)
	}
	var result = spawning.ref.Call("setDirections", packedDirections).Int()
	return ErrorCode(result)
}

type StructureSpawn struct {
	OwnedStructure
}

// TODO
//func (spawn StructureSpawn) Memory() ??? {
//}
func (spawn StructureSpawn) Name() string {
	return spawn.ref.Get("name").String()
}

func (spawn StructureSpawn) Spawning() *StructureSpawnSpawning {
	var spawning = spawn.ref.Get("spawning")
	if spawning.IsNull() {
		return nil
	} else {
		var result = new(StructureSpawnSpawning)
		result.ref = spawning
		result.Directions = nil
		result.Name = spawning.Get("name").String()
		result.NeedTime = spawning.Get("needTime").Int()
		result.RemainingTime = spawning.Get("remainingTime").Int()
		result.Spawn = &spawn

		// directions
		var directions = spawning.Get("directions")
		if !directions.IsUndefined() {
			var length = directions.Length()
			result.Directions = make([]DirectionConstant, length)
			for i := 0; i < length; i++ {
				result.Directions[i] = DirectionConstant(directions.Index(i).Int())
			}
		}
		return result
	}
}

func (spawn StructureSpawn) Store() Store {
	var store = spawn.ref.Get("store")
	return Store{
		ref: store,
	}
}

func (spawn StructureSpawn) SpawnCreep(body []BodyPartConstant, name string, opts *SpawnOpts) ErrorCode {
	var convertedBody = make([]interface{}, len(body))
	for i, v := range body {
		convertedBody[i] = string(v)
	}

	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		var optsMap = map[string]interface{}{}
		if opts.Memory != nil {
			optsMap["memory"] = *opts.Memory
		}
		if opts.EnergyStructures != nil {
			var energyStructures = make([]interface{}, len(*opts.EnergyStructures))
			for i := 0; i < len(energyStructures); i++ {
				energyStructures[i] = (*opts.EnergyStructures)[i].getRef()
			}
			optsMap["energyStructures"] = energyStructures
		}
		if opts.DryRun != nil {
			optsMap["dryRun"] = *opts.DryRun
		}
		if opts.Directions != nil {
			var directions = make([]interface{}, len(*opts.Directions))
			for i := 0; i < len(directions); i++ {
				directions[i] = int((*opts.Directions)[i])
			}
			optsMap["directions"] = directions
		}
		jsOpts = js.ValueOf(optsMap)
	}
	var result = spawn.ref.Call("spawnCreep", convertedBody, name, jsOpts).Int()
	return ErrorCode(result)
}

func (spawn StructureSpawn) RecycleCreep(target Creep) ErrorCode {
	var result = spawn.ref.Call("recycleCreep", target.ref).Int()
	return ErrorCode(result)
}

func (spawn StructureSpawn) RenewCreep(target Creep) ErrorCode {
	var result = spawn.ref.Call("renewCreep", target.ref).Int()
	return ErrorCode(result)
}

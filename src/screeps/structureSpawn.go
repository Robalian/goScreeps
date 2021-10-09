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
	result := spawning.ref.Call("cancel").Int()
	return ErrorCode(result)
}

func (spawning StructureSpawnSpawning) SetDirections(directions []DirectionConstant) ErrorCode {
	packedDirections := make([]interface{}, len(directions))
	for i, v := range directions {
		packedDirections[i] = int(v)
	}
	result := spawning.ref.Call("setDirections", packedDirections).Int()
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
	jsSpawning := spawn.ref.Get("spawning")
	if jsSpawning.IsNull() {
		return nil
	} else {
		result := new(StructureSpawnSpawning)
		result.ref = jsSpawning
		result.Directions = nil
		result.Name = jsSpawning.Get("name").String()
		result.NeedTime = jsSpawning.Get("needTime").Int()
		result.RemainingTime = jsSpawning.Get("remainingTime").Int()
		result.Spawn = &spawn

		// jsDirections
		jsDirections := jsSpawning.Get("directions")
		if !jsDirections.IsUndefined() {
			directionsCount := jsDirections.Length()
			result.Directions = make([]DirectionConstant, directionsCount)
			for i := 0; i < directionsCount; i++ {
				result.Directions[i] = DirectionConstant(jsDirections.Index(i).Int())
			}
		}
		return result
	}
}

func (spawn StructureSpawn) Store() Store {
	jsStore := spawn.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}

func (spawn StructureSpawn) SpawnCreep(body []BodyPartConstant, name string, opts *SpawnOpts) ErrorCode {
	convertedBody := make([]interface{}, len(body))
	for i, bodypartConstant := range body {
		convertedBody[i] = string(bodypartConstant)
	}

	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		tmpOpts := map[string]interface{}{}
		if opts.Memory != nil {
			tmpOpts["memory"] = *opts.Memory
		}
		if opts.EnergyStructures != nil {
			energyStructures := make([]interface{}, len(*opts.EnergyStructures))
			for i := 0; i < len(energyStructures); i++ {
				energyStructures[i] = (*opts.EnergyStructures)[i].getRef()
			}
			tmpOpts["energyStructures"] = energyStructures
		}
		if opts.DryRun != nil {
			tmpOpts["dryRun"] = *opts.DryRun
		}
		if opts.Directions != nil {
			directions := make([]interface{}, len(*opts.Directions))
			for i := 0; i < len(directions); i++ {
				directions[i] = int((*opts.Directions)[i])
			}
			tmpOpts["directions"] = directions
		}
		jsOpts = js.ValueOf(tmpOpts)
	}
	result := spawn.ref.Call("spawnCreep", convertedBody, name, jsOpts).Int()
	return ErrorCode(result)
}

func (spawn StructureSpawn) RecycleCreep(target Creep) ErrorCode {
	result := spawn.ref.Call("recycleCreep", target.ref).Int()
	return ErrorCode(result)
}

func (spawn StructureSpawn) RenewCreep(target Creep) ErrorCode {
	result := spawn.ref.Call("renewCreep", target.ref).Int()
	return ErrorCode(result)
}

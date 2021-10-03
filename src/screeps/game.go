package screeps

import (
	"syscall/js"
)

type Cpu struct {
	Limit        int
	TickLimit    int
	Bucket       int
	ShardLimits  map[string]int
	Unlocked     bool
	UnlockedTime *int
}

type GlobalControlLevel struct {
	Level         int
	Progress      int
	ProgressTotal int
}
type GlobalPowerLevel struct {
	Level         int
	Progress      int
	ProgressTotal int
}

type game struct {
	ref js.Value
}

func (g game) ConstructionSites() map[string]ConstructionSite {
	var constructionSites = g.ref.Get("constructionSites")
	var result = map[string]ConstructionSite{}

	var entries = object.Call("entries", constructionSites)
	var length = entries.Get("length").Int()
	for i := 0; i < length; i++ {
		var v = entries.Index(i)
		result[v.Index(0).String()] = ConstructionSite{RoomObject{
			ref: v.Index(1),
		}}
	}

	return result
}

func (g game) Cpu() Cpu {
	var cpu = g.ref.Get("cpu")

	var result = Cpu{
		Limit:        cpu.Get("limit").Int(),
		TickLimit:    cpu.Get("tickLimit").Int(),
		Bucket:       cpu.Get("bucket").Int(),
		ShardLimits:  map[string]int{},
		Unlocked:     cpu.Get("unlocked").Bool(),
		UnlockedTime: nil,
	}

	// shard limits
	var shardLimitsEntries = object.Call("entries", cpu.Get("shardLimits"))
	var length = shardLimitsEntries.Get("length").Int()
	for i := 0; i < length; i++ {
		var v = shardLimitsEntries.Index(i)
		result.ShardLimits[v.Index(0).String()] = v.Index(1).Int()
	}

	// unlocked time
	var unlockedTime = cpu.Get("unlockedTime")
	if !unlockedTime.IsUndefined() {
		result.UnlockedTime = new(int)
		*result.UnlockedTime = unlockedTime.Int()
	}

	//
	return result
}

func (g game) Creeps() map[string]Creep {
	var creeps = g.ref.Get("creeps")
	var result = map[string]Creep{}

	var entries = object.Call("entries", creeps)
	var length = entries.Get("length").Int()
	for i := 0; i < length; i++ {
		var v = entries.Index(i)
		result[v.Index(0).String()] = Creep{RoomObject{
			ref: v.Index(1),
		}}
	}

	return result
}

func (g game) Flags() map[string]Flag {
	var creeps = g.ref.Get("flags")
	var result = map[string]Flag{}

	var entries = object.Call("entries", creeps)
	var length = entries.Get("length").Int()
	for i := 0; i < length; i++ {
		var v = entries.Index(i)
		result[v.Index(0).String()] = Flag{RoomObject{
			ref: v.Index(1),
		}}
	}

	return result
}

func (g game) Gcl() GlobalControlLevel {
	var gcl = g.ref.Get("gcl")
	return GlobalControlLevel{
		Level:         gcl.Get("level").Int(),
		Progress:      gcl.Get("progress").Int(),
		ProgressTotal: gcl.Get("progressTotal").Int(),
	}
}

func (g game) Gpl() GlobalPowerLevel {
	var gpl = g.ref.Get("gpl")
	return GlobalPowerLevel{
		Level:         gpl.Get("level").Int(),
		Progress:      gpl.Get("progress").Int(),
		ProgressTotal: gpl.Get("progressTotal").Int(),
	}
}

func (g game) GetObjectById(id string) RoomObject {
	var objectById = g.ref.Call("getObjectById", id)
	return RoomObject{
		ref: objectById,
	}
}

func (g game) Spawns() map[string]StructureSpawn {
	var spawns = g.ref.Get("spawns")
	var result = map[string]StructureSpawn{}

	var entries = object.Call("entries", spawns)
	var length = entries.Get("length").Int()
	for i := 0; i < length; i++ {
		var v = entries.Index(i)
		result[v.Index(0).String()] = StructureSpawn{OwnedStructure{Structure{RoomObject{
			ref: v.Index(1),
		}}}}
	}

	return result
}

func (g game) Structures() map[string]Structure {
	var structures = g.ref.Get("structures")
	var result = map[string]Structure{}

	var entries = object.Call("entries", structures)
	var length = entries.Get("length").Int()
	for i := 0; i < length; i++ {
		var v = entries.Index(i)
		result[v.Index(0).String()] = Structure{
			RoomObject{
				ref: v.Index(1),
			},
		}
	}

	return result
}

func (g game) Time() int {
	return g.ref.Get("time").Int()
}

var Game game

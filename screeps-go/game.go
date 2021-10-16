package screeps

import (
	"syscall/js"
)

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
	jsConstructionSites := g.ref.Get("constructionSites")
	result := map[string]ConstructionSite{}

	entries := object.Call("entries", jsConstructionSites)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		result[entry.Index(0).String()] = ConstructionSite{RoomObject{
			ref: entry.Index(1),
		}}
	}

	return result
}

func (g game) Cpu() Cpu {
	jsCpu := g.ref.Get("cpu")

	result := Cpu{
		ref:          jsCpu,
		Limit:        jsCpu.Get("limit").Int(),
		TickLimit:    jsCpu.Get("tickLimit").Int(),
		Bucket:       jsCpu.Get("bucket").Int(),
		ShardLimits:  nil,
		Unlocked:     nil,
		UnlockedTime: nil,
	}

	// shard limits
	jsShardLimits := jsCpu.Get("shardLimits")
	if !jsShardLimits.IsUndefined() {
		shardLimits := map[string]int{}
		shardLimitsEntries := object.Call("entries", jsCpu.Get("shardLimits"))
		shardLimitsLength := shardLimitsEntries.Get("length").Int()
		for i := 0; i < shardLimitsLength; i++ {
			entry := shardLimitsEntries.Index(i)
			key := entry.Index(0).String()
			value := entry.Index(1).Int()
			shardLimits[key] = value
		}
		result.ShardLimits = &shardLimits
	}

	// unlocked
	jsUnlocked := jsCpu.Get("unlocked")
	if !jsUnlocked.IsUndefined() {
		unlocked := jsUnlocked.Bool()
		result.Unlocked = &unlocked
	}

	// unlocked time
	jsUnlockedTime := jsCpu.Get("unlockedTime")
	if !jsUnlockedTime.IsUndefined() {
		unlockedTime := jsUnlockedTime.Int()
		result.UnlockedTime = &unlockedTime
	}

	//
	return result
}

func (g game) Creeps() map[string]Creep {
	jsCreeps := g.ref.Get("creeps")
	result := map[string]Creep{}

	entries := object.Call("entries", jsCreeps)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = Creep{RoomObject{
			ref: value,
		}}
	}

	return result
}

func (g game) Flags() map[string]Flag {
	jsFlags := g.ref.Get("flags")
	result := map[string]Flag{}

	entries := object.Call("entries", jsFlags)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = Flag{RoomObject{
			ref: value,
		}}
	}

	return result
}

func (g game) Gcl() GlobalControlLevel {
	jsGcl := g.ref.Get("gcl")
	return GlobalControlLevel{
		Level:         jsGcl.Get("level").Int(),
		Progress:      jsGcl.Get("progress").Int(),
		ProgressTotal: jsGcl.Get("progressTotal").Int(),
	}
}

func (g game) Gpl() GlobalPowerLevel {
	jsGpl := g.ref.Get("gpl")
	return GlobalPowerLevel{
		Level:         jsGpl.Get("level").Int(),
		Progress:      jsGpl.Get("progress").Int(),
		ProgressTotal: jsGpl.Get("progressTotal").Int(),
	}
}

func (g game) Map() Map {
	return Map{ref: g.ref.Get("map")}
}

func (g game) Market() Market {
	return Market{ref: g.ref.Get("market")}
}

func (g game) PowerCreeps() map[string]PowerCreep {
	jsPowerCreeps := g.ref.Get("powerCreeps")
	result := map[string]PowerCreep{}

	entries := object.Call("entries", jsPowerCreeps)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = PowerCreep{RoomObject{
			ref: value,
		}}
	}

	return result
}

func (g game) Resources() map[string]int {
	jsResources := g.ref.Get("resources")
	result := map[string]int{}

	entries := object.Call("entries", jsResources)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1).Int()
		result[key] = value
	}

	return result
}

func (g game) Rooms() map[string]Room {
	jsPowerCreeps := g.ref.Get("rooms")
	result := map[string]Room{}

	entries := object.Call("entries", jsPowerCreeps)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = Room{
			ref: value,
		}
	}

	return result
}

type Shard struct {
	Name string
	Type string
	Ptr  bool
}

func (g game) Shard() Shard {
	jsShard := g.ref.Get("shard")
	return Shard{
		Name: jsShard.Get("name").String(),
		Type: jsShard.Get("type").String(),
		Ptr:  jsShard.Get("ptr").Bool(),
	}
}

func (g game) Spawns() map[string]StructureSpawn {
	jsSpawns := g.ref.Get("spawns")
	result := map[string]StructureSpawn{}

	entries := object.Call("entries", jsSpawns)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = StructureSpawn{OwnedStructure{Structure{RoomObject{
			ref: value,
		}}}}
	}

	return result
}

func (g game) Structures() map[string]Structure {
	jsStructures := g.ref.Get("structures")
	result := map[string]Structure{}

	entries := object.Call("entries", jsStructures)
	length := entries.Get("length").Int()
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)
		result[key] = Structure{
			RoomObject{
				ref: value,
			},
		}
	}

	return result
}

func (g game) Time() int {
	return g.ref.Get("time").Int()
}

func (g game) GetObjectById(id string) RoomObject {
	jsRoomObject := g.ref.Call("getObjectById", id)
	return RoomObject{
		ref: jsRoomObject,
	}
}

func (g game) Notify(message string, groupInterval *int) {
	var jsGroupInterval js.Value
	if groupInterval == nil {
		jsGroupInterval = js.Undefined()
	} else {
		jsGroupInterval = js.ValueOf(*groupInterval)
	}
	g.ref.Call("notify", message, jsGroupInterval)
}

var Game = game{
	ref: js.Global().Get("Game"),
}

func updateGame() {
	Game = game{
		ref: js.Global().Get("Game"),
	}
}

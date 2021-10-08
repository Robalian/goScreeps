package screeps

import (
	"strconv"
	"syscall/js"
)

type powerCreepConstructor struct {
	ref js.Value
}

var PowerCreepConstructor = powerCreepConstructor{
	ref: js.Global().Get("PowerCreep"),
}

func (powerCreep powerCreepConstructor) Create(name string, className PowerClassConstant) ErrorCode {
	result := powerCreep.ref.Call("create", name, string(className)).Int()
	return ErrorCode(result)
}

type PowerCreep struct {
	RoomObject
}

func (powerCreep PowerCreep) ClassName() PowerClassConstant {
	return PowerClassConstant(powerCreep.ref.Get("className").String())
}

func (powerCreep PowerCreep) DeleteTime() *int {
	jsResult := powerCreep.ref.Get("deleteTime")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.Int()
		return &result
	}
}

func (powerCreep PowerCreep) Hits() *int {
	jsResult := powerCreep.ref.Get("hits")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.Int()
		return &result
	}
}

func (powerCreep PowerCreep) HitsMax() int {
	return powerCreep.ref.Get("hitsMax").Int()
}

func (powerCreep PowerCreep) Id() string {
	return powerCreep.ref.Get("id").String()
}

func (powerCreep PowerCreep) Level() int {
	return powerCreep.ref.Get("level").Int()
}

func (powerCreep PowerCreep) Memory() js.Value {
	return powerCreep.ref.Get("memory")
}

func (powerCreep PowerCreep) My() bool {
	return powerCreep.ref.Get("my").Bool()
}

func (powerCreep PowerCreep) Name() string {
	return powerCreep.ref.Get("name").String()
}

func (powerCreep PowerCreep) Owner() struct{ username string } {
	return struct{ username string }{
		username: powerCreep.ref.Get("owner").Get("username").String(),
	}
}

func (powerCreep PowerCreep) Store() Store {
	return Store{ref: powerCreep.ref.Get("store")}
}

type PowerCreepPower struct {
	level    int
	cooldown *int
}

func (powerCreep PowerCreep) Powers() map[PowerConstant]PowerCreepPower {
	jsPowers := powerCreep.ref.Get("powers")
	entries := object.Call("entries", jsPowers)
	length := entries.Length()
	result := map[PowerConstant]PowerCreepPower{}
	for i := 0; i < length; i++ {
		entry := entries.Index(i)
		key := entry.Index(0).String()
		value := entry.Index(1)

		power := PowerCreepPower{
			level:    value.Get("level").Int(),
			cooldown: nil,
		}
		jsCooldown := value.Get("cooldown")
		if !jsCooldown.IsUndefined() {
			cooldown := jsCooldown.Int()
			power.cooldown = &cooldown
		}

		intKey, _ := strconv.Atoi(key)
		result[PowerConstant(intKey)] = power
	}
	return result
}

func (powerCreep PowerCreep) Saying() *string {
	jsResult := powerCreep.ref.Get("saying")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.String()
		return &result
	}
}

func (powerCreep PowerCreep) Shard() *string {
	jsResult := powerCreep.ref.Get("shard")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.String()
		return &result
	}
}

func (powerCreep PowerCreep) SpawnCooldownTime() *int {
	jsResult := powerCreep.ref.Get("spawnCooldownTime")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.Int()
		return &result
	}
}

func (powerCreep PowerCreep) TicksToLive() *int {
	jsResult := powerCreep.ref.Get("ticksToLive")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.Int()
		return &result
	}
}

func (powerCreep PowerCreep) CancelOrder(methodName string) ErrorCode {
	result := powerCreep.ref.Call("cancelOrder", methodName).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Delete(cancel *bool) ErrorCode {
	var jsCancel js.Value
	if cancel == nil {
		jsCancel = js.Undefined()
	} else {
		jsCancel = js.ValueOf(*cancel)
	}
	result := powerCreep.ref.Call("delete", jsCancel).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Drop(resourceType ResourceConstant, amount *int) ErrorCode {
	var jsAmount js.Value
	if amount == nil {
		jsAmount = js.Undefined()
	} else {
		jsAmount = js.ValueOf(*amount)
	}
	result := powerCreep.ref.Call("drop", string(resourceType), jsAmount).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) EnableRoom(controller StructureController) ErrorCode {
	result := powerCreep.ref.Call("enableRoom", controller.ref).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Move(direction DirectionConstant) ErrorCode {
	result := powerCreep.ref.Call("move", int(direction)).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) MoveByPath_String(path string) ErrorCode {
	result := powerCreep.ref.Call("moveByPath", path).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) MoveByPath_Array(path FindPathResult) ErrorCode {
	packedPath := packFindPathResult(path)
	result := powerCreep.ref.Call("moveByPath", packedPath).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) MoveTo(target RoomPosition, opts *MoveToOpts) ErrorCode {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(packMoveToOpts(*opts))
	}

	result := powerCreep.ref.Call("moveTo", target.ref, jsOpts).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) MoveTo_XY(x int, y int, opts *MoveToOpts) ErrorCode {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(packMoveToOpts(*opts))
	}

	result := powerCreep.ref.Call("moveTo", x, y, jsOpts).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) NotifyWhenAttacked(enabled bool) ErrorCode {
	result := powerCreep.ref.Call("notifyWhenAttacked", enabled).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Pickup(resource Resource) ErrorCode {
	result := powerCreep.ref.Call("pickup", resource.ref).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Rename(name string) ErrorCode {
	result := powerCreep.ref.Call("rename", name).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Renew(target StructureRenewingPowerCreeps) ErrorCode {
	result := powerCreep.ref.Call("renew", target.getRef()).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Say(message string, public bool) ErrorCode {
	result := powerCreep.ref.Call("say", message, public).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Spawn(powerSpawn StructurePowerSpawn) ErrorCode {
	result := powerCreep.ref.Call("spawn", powerSpawn.ref).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Suicide() ErrorCode {
	result := powerCreep.ref.Call("suicide").Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Transfer(target Transferable, resourceType ResourceConstant, amount *int) ErrorCode {
	var jsAmount js.Value
	if amount == nil {
		jsAmount = js.Undefined()
	} else {
		jsAmount = js.ValueOf(*amount)
	}

	result := powerCreep.ref.Call("transfer", target.getRef(), string(resourceType), jsAmount).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Upgrade(power PowerConstant) ErrorCode {
	result := powerCreep.ref.Call("upgrade", int(power)).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) UsePower(power PowerConstant, target RoomObject) ErrorCode {
	result := powerCreep.ref.Call("usePower", int(power), target.ref).Int()
	return ErrorCode(result)
}

func (powerCreep PowerCreep) Withdraw(target Withdrawable, resourceType ResourceConstant, amount *int) ErrorCode {
	var jsAmount js.Value
	if amount == nil {
		jsAmount = js.Undefined()
	} else {
		jsAmount = js.ValueOf(*amount)
	}

	result := powerCreep.ref.Call("withdraw", target.getRef(), string(resourceType), jsAmount).Int()
	return ErrorCode(result)
}

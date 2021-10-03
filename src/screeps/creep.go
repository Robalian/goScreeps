package screeps

import "syscall/js"

type Creep struct {
	RoomObject
}

func (creep Creep) Body() []BodyPart {
	var body = creep.ref.Get("body")
	var length = body.Get("length").Int()
	var result = make([]BodyPart, length)
	for i := 0; i < length; i++ {
		var v = body.Index(i)
		result[i] = BodyPart{
			Boost: nil,
			Type:  BodyPartConstant(v.Get("type").String()),
			Hits:  v.Get("hits").Int(),
		}

		// boost
		var boost = v.Get("boost")
		if !boost.IsUndefined() {
			result[i].Boost = new(ResourceConstant)
			*result[i].Boost = ResourceConstant(boost.String())
		}
	}

	return result
}
func (creep Creep) Fatigue() int {
	return creep.ref.Get("fatigue").Int()
}
func (creep Creep) Hits() int {
	return creep.ref.Get("hits").Int()
}
func (creep Creep) HitsMax() int {
	return creep.ref.Get("hitsMax").Int()
}
func (creep Creep) Id() string {
	return creep.ref.Get("id").String()
}

func (creep Creep) Memory() js.Value {
	return creep.ref.Get("memory")
}

func (creep Creep) My() bool {
	return creep.ref.Get("my").Bool()
}
func (creep Creep) Name() string {
	return creep.ref.Get("name").String()
}
func (creep Creep) Owner() struct{ username string } {
	return struct{ username string }{
		username: creep.ref.Get("owner").Get("username").String(),
	}
}
func (creep Creep) Saying() string {
	return creep.ref.Get("saying").String()
}
func (creep Creep) Spawning() bool {
	return creep.ref.Get("spawning").Bool()
}
func (creep Creep) Store() Store {
	var store = creep.ref.Get("store")
	return Store{
		ref: store,
	}
}
func (creep Creep) TicksToLive() int {
	return creep.ref.Get("ticksToLive").Int()
}

func (creep Creep) Attack(target Attackable) ErrorCode {
	var result = creep.ref.Call("attack", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) AttackController(target StructureController) ErrorCode {
	var result = creep.ref.Call("attackController", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) Build(target ConstructionSite) ErrorCode {
	var result = creep.ref.Call("build", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) CancelOrder(methodName string) ErrorCode {
	var result = creep.ref.Call("cancelOrder", methodName).Int()
	return ErrorCode(result)
}
func (creep Creep) ClaimController(target StructureController) ErrorCode {
	var result = creep.ref.Call("claimController", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) Dismantle(target Structure) ErrorCode {
	var result = creep.ref.Call("dismantle", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) Drop(resourceType ResourceConstant) ErrorCode {
	var result = creep.ref.Call("drop", string(resourceType)).Int()
	return ErrorCode(result)
}
func (creep Creep) GenerateSafeMode(controller StructureController) ErrorCode {
	var result = creep.ref.Call("generateSafeMode", controller.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) GetActiveBodyparts(bodypartType BodyPartConstant) ErrorCode {
	var result = creep.ref.Call("getActiveBodyparts", string(bodypartType)).Int()
	return ErrorCode(result)
}
func (creep Creep) Harvest(target Harvestable) ErrorCode {
	var result = creep.ref.Call("harvest", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) Heal(target AnyCreep) ErrorCode {
	var result = creep.ref.Call("heal", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) Move(direction DirectionConstant) ErrorCode {
	var result = creep.ref.Call("move", int(direction)).Int()
	return ErrorCode(result)
}
func (creep Creep) Move_Creep(otherCreep Creep) ErrorCode {
	var result = creep.ref.Call("move", otherCreep.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) MoveByPath_String(path string) ErrorCode {
	var result = creep.ref.Call("moveByPath", path).Int()
	return ErrorCode(result)
}

func (creep Creep) MoveByPath_Array(path FindPathResult) ErrorCode {
	var packedPath = packFindPathResult(path)
	var result = creep.ref.Call("moveByPath", packedPath).Int()
	return ErrorCode(result)
}

func (creep Creep) MoveTo(target RoomPosition, opts *MoveToOpts) ErrorCode {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(packMoveToOpts(*opts))
	}

	var result = creep.ref.Call("moveTo", target.ref, jsOpts).Int()
	return ErrorCode(result)
}

func (creep Creep) MoveTo_XY(x int, y int, opts *MoveToOpts) ErrorCode {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(packMoveToOpts(*opts))
	}

	var result = creep.ref.Call("moveTo", x, y, jsOpts).Int()
	return ErrorCode(result)
}
func (creep Creep) NotifyWhenAttacked(enabled bool) ErrorCode {
	var result = creep.ref.Call("notifyWhenAttacked", enabled).Int()
	return ErrorCode(result)
}
func (creep Creep) Pickup(resource Resource) ErrorCode {
	var result = creep.ref.Call("pickup", resource.ref).Int()
	return ErrorCode(result)
}
func (creep Creep) Pull(target Creep) ErrorCode {
	var result = creep.ref.Call("pull", target.ref).Int()
	return ErrorCode(result)
}
func (creep Creep) RangedAttack(target Attackable) ErrorCode {
	var result = creep.ref.Call("rangedAttack", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) RangedHeal(target AnyCreep) ErrorCode {
	var result = creep.ref.Call("rangedHeal", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) RangedMassAttack() ErrorCode {
	var result = creep.ref.Call("rangedMassAttack").Int()
	return ErrorCode(result)
}
func (creep Creep) Repair(target Structure) ErrorCode {
	var result = creep.ref.Call("repair", target.ref).Int()
	return ErrorCode(result)
}
func (creep Creep) ReserveController(target StructureController) ErrorCode {
	var result = creep.ref.Call("reserveController", target.ref).Int()
	return ErrorCode(result)
}
func (creep Creep) Say(message string, public bool) ErrorCode {
	var result = creep.ref.Call("say", message, public).Int()
	return ErrorCode(result)
}
func (creep Creep) SignController(target StructureController, text string) ErrorCode {
	var result = creep.ref.Call("signController", target.ref, text).Int()
	return ErrorCode(result)
}
func (creep Creep) Suicide() ErrorCode {
	var result = creep.ref.Call("suicide").Int()
	return ErrorCode(result)
}
func (creep Creep) Transfer(target Transferable, resourceType ResourceConstant, amount *int) ErrorCode {
	var jsAmount js.Value
	if amount == nil {
		jsAmount = js.Undefined()
	} else {
		jsAmount = js.ValueOf(*amount)
	}

	var result = creep.ref.Call("transfer", target.getRef(), string(resourceType), jsAmount).Int()
	return ErrorCode(result)
}
func (creep Creep) UpgradeController(target StructureController) ErrorCode {
	var result = creep.ref.Call("upgradeController", target.ref).Int()
	return ErrorCode(result)
}
func (creep Creep) Withdraw(target Withdrawable, resourceType ResourceConstant, amount *int) ErrorCode {
	var jsAmount js.Value
	if amount == nil {
		jsAmount = js.Undefined()
	} else {
		jsAmount = js.ValueOf(*amount)
	}

	var result = creep.ref.Call("withdraw", target.getRef(), string(resourceType), jsAmount).Int()
	return ErrorCode(result)
}

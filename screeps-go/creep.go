package screeps

import "syscall/js"

type BodyPart struct {
	Boost *ResourceConstant
	Type  BodyPartConstant
	Hits  int
}

type MoveToOpts struct {
	FindPathOpts

	ReusePath          *uint
	SerializeMemory    *bool
	NoPathFinding      *bool
	VisualizePathStyle *PolyStyle
}

func packMoveToOpts(opts MoveToOpts) map[string]interface{} {
	result := packFindPathOpts(opts.FindPathOpts)

	if opts.NoPathFinding != nil {
		result["noPathFinding"] = *opts.NoPathFinding
	}
	if opts.ReusePath != nil {
		result["reusePath"] = *opts.ReusePath
	}
	if opts.SerializeMemory != nil {
		result["serializeMemory"] = *opts.SerializeMemory
	}
	if opts.VisualizePathStyle != nil {
		result["visualizePathStyle"] = packPolyStyle(*opts.VisualizePathStyle)
	}

	return result
}

var creepConstructor = js.Global().Get("Creep")

type Creep struct {
	RoomObject
}

func (creep Creep) iAmAnyCreep() {}

func (creep Creep) Body() []BodyPart {
	body := creep.ref.Get("body")
	bodyLength := body.Get("length").Int()
	result := make([]BodyPart, bodyLength)
	for i := 0; i < bodyLength; i++ {
		jsBodypart := body.Index(i)
		result[i] = BodyPart{
			Boost: nil,
			Type:  BodyPartConstant(jsBodypart.Get("type").String()),
			Hits:  jsBodypart.Get("hits").Int(),
		}

		// jsBoost
		jsBoost := jsBodypart.Get("boost")
		if !jsBoost.IsUndefined() {
			result[i].Boost = new(ResourceConstant)
			*result[i].Boost = ResourceConstant(jsBoost.String())
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
	jsStore := creep.ref.Get("store")
	return Store{
		ref: jsStore,
	}
}
func (creep Creep) TicksToLive() int {
	return creep.ref.Get("ticksToLive").Int()
}

func (creep Creep) Attack(target Attackable) ErrorCode {
	result := creep.ref.Call("attack", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) AttackController(target StructureController) ErrorCode {
	result := creep.ref.Call("attackController", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) Build(target ConstructionSite) ErrorCode {
	result := creep.ref.Call("build", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) CancelOrder(methodName string) ErrorCode {
	result := creep.ref.Call("cancelOrder", methodName).Int()
	return ErrorCode(result)
}
func (creep Creep) ClaimController(target StructureController) ErrorCode {
	result := creep.ref.Call("claimController", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) Dismantle(target Structure) ErrorCode {
	result := creep.ref.Call("dismantle", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) Drop(resourceType ResourceConstant, amount *int) ErrorCode {
	var jsAmount js.Value
	if amount == nil {
		jsAmount = js.Undefined()
	} else {
		jsAmount = js.ValueOf(*amount)
	}
	result := creep.ref.Call("drop", string(resourceType), jsAmount).Int()
	return ErrorCode(result)
}
func (creep Creep) GenerateSafeMode(controller StructureController) ErrorCode {
	result := creep.ref.Call("generateSafeMode", controller.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) GetActiveBodyparts(bodypartType BodyPartConstant) ErrorCode {
	result := creep.ref.Call("getActiveBodyparts", string(bodypartType)).Int()
	return ErrorCode(result)
}
func (creep Creep) Harvest(target Harvestable) ErrorCode {
	result := creep.ref.Call("harvest", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) Heal(target AnyCreep) ErrorCode {
	result := creep.ref.Call("heal", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) Move(direction DirectionConstant) ErrorCode {
	result := creep.ref.Call("move", int(direction)).Int()
	return ErrorCode(result)
}
func (creep Creep) Move_ToCreep(otherCreep Creep) ErrorCode {
	result := creep.ref.Call("move", otherCreep.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) MoveByPath_String(path string) ErrorCode {
	result := creep.ref.Call("moveByPath", path).Int()
	return ErrorCode(result)
}

func (creep Creep) MoveByPath_Array(path FindPathResult) ErrorCode {
	packedPath := packFindPathResult(path)
	result := creep.ref.Call("moveByPath", packedPath).Int()
	return ErrorCode(result)
}

func (creep Creep) MoveTo(target RoomPosition, opts *MoveToOpts) ErrorCode {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(packMoveToOpts(*opts))
	}

	result := creep.ref.Call("moveTo", target.ref, jsOpts).Int()
	return ErrorCode(result)
}

func (creep Creep) MoveTo_XY(x int, y int, opts *MoveToOpts) ErrorCode {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		jsOpts = js.ValueOf(packMoveToOpts(*opts))
	}

	result := creep.ref.Call("moveTo", x, y, jsOpts).Int()
	return ErrorCode(result)
}
func (creep Creep) NotifyWhenAttacked(enabled bool) ErrorCode {
	result := creep.ref.Call("notifyWhenAttacked", enabled).Int()
	return ErrorCode(result)
}
func (creep Creep) Pickup(resource Resource) ErrorCode {
	result := creep.ref.Call("pickup", resource.ref).Int()
	return ErrorCode(result)
}
func (creep Creep) Pull(target Creep) ErrorCode {
	result := creep.ref.Call("pull", target.ref).Int()
	return ErrorCode(result)
}
func (creep Creep) RangedAttack(target Attackable) ErrorCode {
	result := creep.ref.Call("rangedAttack", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) RangedHeal(target AnyCreep) ErrorCode {
	result := creep.ref.Call("rangedHeal", target.getRef()).Int()
	return ErrorCode(result)
}
func (creep Creep) RangedMassAttack() ErrorCode {
	result := creep.ref.Call("rangedMassAttack").Int()
	return ErrorCode(result)
}
func (creep Creep) Repair(target Structure) ErrorCode {
	result := creep.ref.Call("repair", target.ref).Int()
	return ErrorCode(result)
}
func (creep Creep) ReserveController(target StructureController) ErrorCode {
	result := creep.ref.Call("reserveController", target.ref).Int()
	return ErrorCode(result)
}
func (creep Creep) Say(message string, public bool) ErrorCode {
	result := creep.ref.Call("say", message, public).Int()
	return ErrorCode(result)
}
func (creep Creep) SignController(target StructureController, text string) ErrorCode {
	result := creep.ref.Call("signController", target.ref, text).Int()
	return ErrorCode(result)
}
func (creep Creep) Suicide() ErrorCode {
	result := creep.ref.Call("suicide").Int()
	return ErrorCode(result)
}
func (creep Creep) Transfer(target Transferable, resourceType ResourceConstant, amount *int) ErrorCode {
	var jsAmount js.Value
	if amount == nil {
		jsAmount = js.Undefined()
	} else {
		jsAmount = js.ValueOf(*amount)
	}

	result := creep.ref.Call("transfer", target.getRef(), string(resourceType), jsAmount).Int()
	return ErrorCode(result)
}
func (creep Creep) UpgradeController(target StructureController) ErrorCode {
	result := creep.ref.Call("upgradeController", target.ref).Int()
	return ErrorCode(result)
}
func (creep Creep) Withdraw(target Withdrawable, resourceType ResourceConstant, amount *int) ErrorCode {
	var jsAmount js.Value
	if amount == nil {
		jsAmount = js.Undefined()
	} else {
		jsAmount = js.ValueOf(*amount)
	}

	result := creep.ref.Call("withdraw", target.getRef(), string(resourceType), jsAmount).Int()
	return ErrorCode(result)
}

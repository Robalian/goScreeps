package screeps

type Deposit struct {
	RoomObject
}

func (deposit Deposit) iAmHarvestable() {}

func (deposit Deposit) Cooldown() int {
	return deposit.ref.Get("cooldown").Int()
}

func (deposit Deposit) DepositType() ResourceConstant {
	return ResourceConstant(deposit.ref.Get("depositType").String())
}

func (deposit Deposit) Id() string {
	return deposit.ref.Get("id").String()
}

func (deposit Deposit) LastCooldown() int {
	return deposit.ref.Get("lastCooldown").Int()
}

func (deposit Deposit) TicksToDecay() int {
	return deposit.ref.Get("tickToDecay").Int()
}

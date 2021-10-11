package screeps

type StructurePowerBank struct {
	Structure
}

func (powerBank StructurePowerBank) iAmRenewingPowerCreeps() {}

func (powerBank StructurePowerBank) Power() int {
	return powerBank.ref.Get("power").Int()
}

func (powerBank StructurePowerBank) TicksToDecay() int {
	return powerBank.ref.Get("ticksToDecay").Int()
}

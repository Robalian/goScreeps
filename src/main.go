package main

import (
	. "screepsgo/screeps-go"
	"syscall/js"
)

func main() {
	channel := make(chan bool)

	runLoop := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		channel <- true
		return nil
	})
	js.Global().Set("runLoop", runLoop)

	for {
		<-channel
		PreMain()

		for _, creep := range Game.Creeps() {
			creepRole := creep.Memory().Get("role")
			if creepRole.IsUndefined() {
				continue
			}

			if creep.Memory().Get("role").String() == "harvester" {
				roleHarvester(creep)
			}
			if creep.Memory().Get("role").String() == "upgrader" {
				roleUpgrader(creep)
			}
			if creep.Memory().Get("role").String() == "builder" {
				roleBuilder(creep)
			}
		}

		//
		spawnCreeps()

		//
		runTowers()

		//
		pathfinderExample()

		//
		setGrafanaStats()

		PostMain()
	}
}

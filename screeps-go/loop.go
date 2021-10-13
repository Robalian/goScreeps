package screeps

func PreMain() {
	updateGame()
	updateRawMemory()
	loadSegments()
}

func PostMain() {
	saveSegments()
}

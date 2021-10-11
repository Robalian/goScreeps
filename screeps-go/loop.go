package screeps

//export preMain
func PreMain() {
	updateGame()
	loadSegments()
}

//export postMain
func PostMain() {
	saveSegments()
}

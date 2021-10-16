package main

import (
	"runtime"
	. "screepsgo/screeps-go"
)

// Stack returns a formatted stack trace of the goroutine that calls it.
// It calls runtime.Stack with a large enough buffer to capture the entire trace.
func Stack() []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			Console.Log(r)
			Console.Log(string(Stack()))
		}
	}()

	InitScreeps()

	for {
		PreMain()

		runCreeps()

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

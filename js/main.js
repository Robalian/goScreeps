let go = undefined
function loadWasm() {
	require('wasm_exec');
	const bytecode = require('screepsgo');
	const wasmModule = new WebAssembly.Module(bytecode);
	
	let localGo = new Go();
	let wasmInstance = new WebAssembly.Instance(wasmModule, localGo.importObject)
    localGo.init(wasmInstance);
	//localGo.defaultRun(wasmInstance)
	go = localGo;
	global.go = go
}

function jsRoomCallback(roomName) {
	global.jsRoomCallbackArgument = roomName
	let result = go._inst.exports.goRoomCallback(roomName)
	
	console.log(global.goRoomCallbackResult)
	
	return global.goRoomCallbackResult
}
global.jsRoomCallback = jsRoomCallback

let globalResetTick = Game.time;
module.exports.loop = () => {
	if (!go && Game.cpu.bucket === 10000)
		loadWasm();
	if (!go)
		return;
		
	go.run()

	//	wasmInstance.exports.preMain();
	//	go.run()
	//	wasmInstance.exports.Loop();
	//	wasmInstance.exports.postMain();

	RawMemory.segments[99] = JSON.stringify({
		globalAge: Game.time - globalResetTick,
		profiling: {
			bucket: Game.cpu.bucket,
			limit: Game.cpu.limit,
			heapStatistics: Game.cpu.getHeapStatistics(),
			memory: RawMemory.get().length
		}
	});
}

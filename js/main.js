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
	global.roomCallbackArgument = roomName
	go._inst.exports.goRoomCallback()
	return global.roomCallbackResult
}
global.jsRoomCallback = jsRoomCallback

function jsOrderFilter(order) {
	global.orderFilterArgument = order
	return go._inst.exports.goOrderFilter()
}
global.jsOrderFilter = jsOrderFilter

function jsRouteCallback(roomName, fromRoomName) {
	global.routeCallbackArgument1 = roomName
	global.routeCallbackArgument2 = fromRoomName
	return go._inst.exports.goRouteCallback()
}
global.jsRouteCallback = jsRouteCallback

function jsCostCallback(roomName, costMatrix) {
	global.costCallbackArgument1 = roomName
	global.costCallbackArgument2 = costMatrix
	go._inst.exports.goCostCallback()
}
global.jsCostCallback = jsCostCallback

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

require('wasm_exec');
const bytecode = require('screepsgo');
const wasmModule = new WebAssembly.Module(bytecode);
let wasmInstance;
let go = new Go();

wasmInstance = new WebAssembly.Instance(wasmModule, go.importObject)
go.run(wasmInstance);
global.wasmInstance = wasmInstance;



let globalResetTick = Game.time;
module.exports.loop = () => {
	wasmInstance.exports.preMain();
	wasmInstance.exports.Loop();
	wasmInstance.exports.postMain();

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

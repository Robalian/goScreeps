var utf8 = require('utf8')

class TextEncoder {
	constructor(format) {
	}
	encode(text) {
		let encodedString = utf8.encode(text);
		let buffer = new Int8Array(encodedString.length);
		for (let i=0; i<encodedString.length; ++i)
			buffer[i] = encodedString.charCodeAt(i);
		return buffer;
	}
}
global.TextEncoder = TextEncoder;

class TextDecoder {
	constructor(format) {
	}
	
	decode(dataView) {
		let memoryView = new Int8Array(dataView.buffer, dataView.byteOffset, dataView.byteLength)
		let text = String.fromCharCode.apply(null, memoryView)
		return utf8.decode(text);
	}
}
global.TextDecoder = TextDecoder;

const crypto = {
	getRandomValues: function(typedArray) {
		for (let i=0; i<typedArray.length; ++i) {
			typedArray[i] = _.random(0, 255, false);
		}
		return typedArray;
	}
};
global.crypto = crypto;

const Console = {
	log: console.log,
	error: console.log,
	warn: console.log
};
global.Console = Console;

const process = {
	hrtime() {
		return [Date.now(), 0];
	}
}
global.process = process;

const fs = {
	constants: { O_WRONLY: -1, O_RDWR: -1, O_CREAT: -1, O_TRUNC: -1, O_APPEND: -1, O_EXCL: -1 }, // unused
	writeSync(fd, buf) {
		outputBuf += decoder.decode(buf);
		const nl = outputBuf.lastIndexOf("\n");
		if (nl != -1) {
			console.log(outputBuf.substr(0, nl));
			outputBuf = outputBuf.substr(nl + 1);
		}
		return buf.length;
	},
	write(fd, buf, offset, length, position, callback) {
		if (offset !== 0 || length !== buf.length || position !== null) {
			callback(enosys());
			return;
		}
		const n = this.writeSync(fd, buf);
		callback(null, n);
	},
	chmod(path, mode, callback) { callback(enosys()); },
	chown(path, uid, gid, callback) { callback(enosys()); },
	close(fd, callback) { callback(enosys()); },
	fchmod(fd, mode, callback) { callback(enosys()); },
	fchown(fd, uid, gid, callback) { callback(enosys()); },
	fstat(fd, callback) { callback(enosys()); },
	fsync(fd, callback) { callback(null); },
	ftruncate(fd, length, callback) { callback(enosys()); },
	lchown(path, uid, gid, callback) { callback(enosys()); },
	link(path, link, callback) { callback(enosys()); },
	lstat(path, callback) { callback(enosys()); },
	mkdir(path, perm, callback) { callback(enosys()); },
	open(path, flags, mode, callback) { callback(enosys()); },
	read(fd, buffer, offset, length, position, callback) { callback(enosys()); },
	readdir(path, callback) { callback(enosys()); },
	readlink(path, callback) { callback(enosys()); },
	rename(from, to, callback) { callback(enosys()); },
	rmdir(path, callback) { callback(enosys()); },
	stat(path, callback) { callback(enosys()); },
	symlink(path, link, callback) { callback(enosys()); },
	truncate(path, length, callback) { callback(enosys()); },
	unlink(path, callback) { callback(enosys()); },
	utimes(path, atime, mtime, callback) { callback(enosys()); },
};
global.fs = fs;

require('wasm_exec')

//---------------------------------------------------------------------------------

let go = undefined
function loadWasm() {
	const bytecode = require('screepsgo');
	const wasmModule = new WebAssembly.Module(bytecode);

	let localGo = new Go();
	let wasmInstance = new WebAssembly.Instance(wasmModule, localGo.importObject)
	localGo.run(wasmInstance)
	//localGo.defaultRun(wasmInstance)
	go = localGo;
	global.go = go
}

function jsRoomCallback(roomName) {
	global.roomCallbackArgument = roomName
	global.goRoomCallback()
	return global.roomCallbackResult
}
global.jsRoomCallback = jsRoomCallback

function jsOrderFilter(order) {
	global.orderFilterArgument = order
	return global.goOrderFilter()
}
global.jsOrderFilter = jsOrderFilter

function jsRouteCallback(roomName, fromRoomName) {
	global.routeCallbackArgument1 = roomName
	global.routeCallbackArgument2 = fromRoomName
	return global.goRouteCallback()
}
global.jsRouteCallback = jsRouteCallback

function jsCostCallback(roomName, costMatrix) {
	global.costCallbackArgument1 = roomName
	global.costCallbackArgument2 = costMatrix
	global.goCostCallback()
}
global.jsCostCallback = jsCostCallback

module.exports.loop = () => {
	if (!go && Game.cpu.bucket === 10000)
		loadWasm();
	if (!go)
		return;
	
	runLoop();
}

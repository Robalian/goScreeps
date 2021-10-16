/*! https://mths.be/utf8js v3.0.0 by @mathias */
var utf8 = {}
;(function(root) {
	var stringFromCharCode = String.fromCharCode;

	// Taken from https://mths.be/punycode
	function ucs2decode(string) {
		var output = [];
		var counter = 0;
		var length = string.length;
		var value;
		var extra;
		while (counter < length) {
			value = string.charCodeAt(counter++);
			if (value >= 0xD800 && value <= 0xDBFF && counter < length) {
				// high surrogate, and there is a next character
				extra = string.charCodeAt(counter++);
				if ((extra & 0xFC00) == 0xDC00) { // low surrogate
					output.push(((value & 0x3FF) << 10) + (extra & 0x3FF) + 0x10000);
				} else {
					// unmatched surrogate; only append this code unit, in case the next
					// code unit is the high surrogate of a surrogate pair
					output.push(value);
					counter--;
				}
			} else {
				output.push(value);
			}
		}
		return output;
	}

	// Taken from https://mths.be/punycode
	function ucs2encode(array) {
		var length = array.length;
		var index = -1;
		var value;
		var output = '';
		while (++index < length) {
			value = array[index];
			if (value > 0xFFFF) {
				value -= 0x10000;
				output += stringFromCharCode(value >>> 10 & 0x3FF | 0xD800);
				value = 0xDC00 | value & 0x3FF;
			}
			output += stringFromCharCode(value);
		}
		return output;
	}

	function checkScalarValue(codePoint) {
		if (codePoint >= 0xD800 && codePoint <= 0xDFFF) {
			throw Error(
				'Lone surrogate U+' + codePoint.toString(16).toUpperCase() +
				' is not a scalar value'
			);
		}
	}

	function createByte(codePoint, shift) {
		return stringFromCharCode(((codePoint >> shift) & 0x3F) | 0x80);
	}

	function encodeCodePoint(codePoint) {
		if ((codePoint & 0xFFFFFF80) == 0) { // 1-byte sequence
			return stringFromCharCode(codePoint);
		}
		var symbol = '';
		if ((codePoint & 0xFFFFF800) == 0) { // 2-byte sequence
			symbol = stringFromCharCode(((codePoint >> 6) & 0x1F) | 0xC0);
		}
		else if ((codePoint & 0xFFFF0000) == 0) { // 3-byte sequence
			checkScalarValue(codePoint);
			symbol = stringFromCharCode(((codePoint >> 12) & 0x0F) | 0xE0);
			symbol += createByte(codePoint, 6);
		}
		else if ((codePoint & 0xFFE00000) == 0) { // 4-byte sequence
			symbol = stringFromCharCode(((codePoint >> 18) & 0x07) | 0xF0);
			symbol += createByte(codePoint, 12);
			symbol += createByte(codePoint, 6);
		}
		symbol += stringFromCharCode((codePoint & 0x3F) | 0x80);
		return symbol;
	}

	function utf8encode(string) {
		var codePoints = ucs2decode(string);
		var length = codePoints.length;
		var index = -1;
		var codePoint;
		var byteString = '';
		while (++index < length) {
			codePoint = codePoints[index];
			byteString += encodeCodePoint(codePoint);
		}
		return byteString;
	}

	function readContinuationByte() {
		if (byteIndex >= byteCount) {
			throw Error('Invalid byte index');
		}

		var continuationByte = byteArray[byteIndex] & 0xFF;
		byteIndex++;

		if ((continuationByte & 0xC0) == 0x80) {
			return continuationByte & 0x3F;
		}

		// If we end up here, it’s not a continuation byte
		throw Error('Invalid continuation byte');
	}

	function decodeSymbol() {
		var byte1;
		var byte2;
		var byte3;
		var byte4;
		var codePoint;

		if (byteIndex > byteCount) {
			throw Error('Invalid byte index');
		}

		if (byteIndex == byteCount) {
			return false;
		}

		// Read first byte
		byte1 = byteArray[byteIndex] & 0xFF;
		byteIndex++;

		// 1-byte sequence (no continuation bytes)
		if ((byte1 & 0x80) == 0) {
			return byte1;
		}

		// 2-byte sequence
		if ((byte1 & 0xE0) == 0xC0) {
			byte2 = readContinuationByte();
			codePoint = ((byte1 & 0x1F) << 6) | byte2;
			if (codePoint >= 0x80) {
				return codePoint;
			} else {
				throw Error('Invalid continuation byte');
			}
		}

		// 3-byte sequence (may include unpaired surrogates)
		if ((byte1 & 0xF0) == 0xE0) {
			byte2 = readContinuationByte();
			byte3 = readContinuationByte();
			codePoint = ((byte1 & 0x0F) << 12) | (byte2 << 6) | byte3;
			if (codePoint >= 0x0800) {
				checkScalarValue(codePoint);
				return codePoint;
			} else {
				throw Error('Invalid continuation byte');
			}
		}

		// 4-byte sequence
		if ((byte1 & 0xF8) == 0xF0) {
			byte2 = readContinuationByte();
			byte3 = readContinuationByte();
			byte4 = readContinuationByte();
			codePoint = ((byte1 & 0x07) << 0x12) | (byte2 << 0x0C) |
				(byte3 << 0x06) | byte4;
			if (codePoint >= 0x010000 && codePoint <= 0x10FFFF) {
				return codePoint;
			}
		}

		throw Error('Invalid UTF-8 detected');
	}

	var byteArray;
	var byteCount;
	var byteIndex;
	function utf8decode(byteString) {
		byteArray = ucs2decode(byteString);
		byteCount = byteArray.length;
		byteIndex = 0;
		var codePoints = [];
		var tmp;
		while ((tmp = decodeSymbol()) !== false) {
			codePoints.push(tmp);
		}
		return ucs2encode(codePoints);
	}
	root.version = '3.0.0';
	root.encode = utf8encode;
	root.decode = utf8decode;

}(utf8));

//---------------------------------------------------------------------------------

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file has been modified for use by the TinyGo compiler.

const enosys = () => {
	const err = new Error("not implemented");
	err.code = "ENOSYS";
	return err;
};
if (!global.fs) {
	let outputBuf = "";
	global.fs = {
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
}

global.Go = class {
	constructor() {
		this.argv = ["js"];
		this.env = {};
		this.exit = (code) => {
			if (code !== 0) {
				console.log("exit code:", code);
			}
		};
		this._pendingEvent = null;
		this.encoder = {
			encode: function(text) {
				let encodedString = utf8.encode(text);
				let buffer = new Int8Array(encodedString.length);
				for (let i=0; i<encodedString.length; ++i)
					buffer[i] = encodedString.charCodeAt(i);
				return buffer;
			}
		};
		this.decoder = {
			decode: function(dataView) {
				let memoryView = new Int8Array(dataView.buffer, dataView.byteOffset, dataView.byteLength)
				let text = String.fromCharCode.apply(null, memoryView)
				return utf8.decode(text);
			}
		};
		this.getRandomValues = function(typedArray) {
			for (let i=0; i<typedArray.length; ++i) {
				typedArray[i] = _.random(0, 255, false);
			}
			return typedArray;
		};
		this._scheduledTimeouts = new Map();
		this._nextCallbackTimeoutID = 1;

		const setInt64 = (addr, v) => {
			this.mem.setUint32(addr + 0, v, true);
			this.mem.setUint32(addr + 4, Math.floor(v / 4294967296), true);
		}

		const getInt64 = (addr) => {
			const low = this.mem.getUint32(addr + 0, true);
			const high = this.mem.getInt32(addr + 4, true);
			return low + high * 4294967296;
		}

		const loadValue = (addr) => {
			const f = this.mem.getFloat64(addr, true);
			if (f === 0) {
				return undefined;
			}
			if (!isNaN(f)) {
				return f;
			}

			const id = this.mem.getUint32(addr, true);
			return this._values[id];
		}

		const storeValue = (addr, v) => {
			const nanHead = 0x7FF80000;

			if (typeof v === "number" && v !== 0) {
				if (isNaN(v)) {
					this.mem.setUint32(addr + 4, nanHead, true);
					this.mem.setUint32(addr, 0, true);
					return;
				}
				this.mem.setFloat64(addr, v, true);
				return;
			}

			if (v === undefined) {
				this.mem.setFloat64(addr, 0, true);
				return;
			}

			let id = this._ids.get(v);
			if (id === undefined) {
				id = this._idPool.pop();
				if (id === undefined) {
					id = this._values.length;
				}
				this._values[id] = v;
				this._goRefCounts[id] = 0;
				this._ids.set(v, id);
			}
			this._goRefCounts[id]++;
			let typeFlag = 0;
			switch (typeof v) {
				case "object":
					if (v !== null) {
						typeFlag = 1;
					}
					break;
				case "string":
					typeFlag = 2;
					break;
				case "symbol":
					typeFlag = 3;
					break;
				case "function":
					typeFlag = 4;
					break;
			}
			this.mem.setUint32(addr + 4, nanHead | typeFlag, true);
			this.mem.setUint32(addr, id, true);
		}

		const loadSlice = (addr) => {
			const array = getInt64(addr + 0);
			const len = getInt64(addr + 8);
			return new Uint8Array(this._inst.exports.mem.buffer, array, len);
		}

		const loadSliceOfValues = (addr) => {
			const array = getInt64(addr + 0);
			const len = getInt64(addr + 8);
			const a = new Array(len);
			for (let i = 0; i < len; i++) {
				a[i] = loadValue(array + i * 8);
			}
			return a;
		}

		const loadString = (addr) => {
			const saddr = getInt64(addr + 0);
			const len = getInt64(addr + 8);
			return this.decoder.decode(new DataView(this._inst.exports.mem.buffer, saddr, len));
		}

		const timeOrigin = Date.now();
		this.importObject = {
			go: {
				// Go's SP does not change as long as no Go code is running. Some operations (e.g. calls, getters and setters)
				// may synchronously trigger a Go event handler. This makes Go code get executed in the middle of the imported
				// function. A goroutine can switch to a new stack if the current stack is too small (see morestack function).
				// This changes the SP, thus we have to update the SP used by the imported function.

				// func wasmExit(code int32)
				"runtime.wasmExit": (sp) => {
					sp >>>= 0;
					const code = this.mem.getInt32(sp + 8, true);
					this.exited = true;
					delete this._inst;
					delete this._values;
					delete this._goRefCounts;
					delete this._ids;
					delete this._idPool;
					this.exit(code);
				},

				// func wasmWrite(fd uintptr, p unsafe.Pointer, n int32)
				"runtime.wasmWrite": (sp) => {
					//sp >>>= 0;
					//const fd = getInt64(sp + 8);
					//const p = getInt64(sp + 16);
					//const n = this.mem.getInt32(sp + 24, true);
					//fs.writeSync(fd, new Uint8Array(this._inst.exports.mem.buffer, p, n));
				},

				// func resetMemoryDataView()
				"runtime.resetMemoryDataView": (sp) => {
					sp >>>= 0;
					this.mem = new DataView(this._inst.exports.mem.buffer);
				},

				// func nanotime1() int64
				"runtime.nanotime1": (sp) => {
					sp >>>= 0;
					setInt64(sp + 8, Date.now() * 1000000);
				},

				// func walltime() (sec int64, nsec int32)
				"runtime.walltime": (sp) => {
					sp >>>= 0;
					const msec = (new Date).getTime();
					setInt64(sp + 8, msec / 1000);
					this.mem.setInt32(sp + 16, (msec % 1000) * 1000000, true);
				},

				// func scheduleTimeoutEvent(delay int64) int32
				"runtime.scheduleTimeoutEvent": (sp) => {
					// screeps doesn't have "setTimeout". Hopefully this function won't be needed.

					//sp >>>= 0;
					//const id = this._nextCallbackTimeoutID;
					//this._nextCallbackTimeoutID++;
					//this._scheduledTimeouts.set(id, setTimeout(
					//	() => {
					//		this._resume();
					//		while (this._scheduledTimeouts.has(id)) {
					//			// for some reason Go failed to register the timeout event, log and try again
					//			// (temporary workaround for https://github.com/golang/go/issues/28975)
					//			console.warn("scheduleTimeoutEvent: missed timeout event");
					//			this._resume();
					//		}
					//	},
					//	getInt64(sp + 8) + 1, // setTimeout has been seen to fire up to 1 millisecond early
					//));
					//this.mem.setInt32(sp + 16, id, true);
				},

				// func clearTimeoutEvent(id int32)
				"runtime.clearTimeoutEvent": (sp) => {
					// screeps doesn't have "setTimeout". Hopefully this function won't be needed.

					//sp >>>= 0;
					//const id = this.mem.getInt32(sp + 8, true);
					//clearTimeout(this._scheduledTimeouts.get(id));
					//this._scheduledTimeouts.delete(id);
				},

				// func getRandomData(r []byte)
				"runtime.getRandomData": (sp) => {
					sp >>>= 0;
					this.getRandomValues(loadSlice(sp + 8));
				},

				// func finalizeRef(v ref)
				"syscall/js.finalizeRef": (sp) => {
					sp >>>= 0;
					const id = this.mem.getUint32(sp + 8, true);
					this._goRefCounts[id]--;
					if (this._goRefCounts[id] === 0) {
						const v = this._values[id];
						this._values[id] = null;
						this._ids.delete(v);
						this._idPool.push(id);
					}
				},

				// func stringVal(value string) ref
				"syscall/js.stringVal": (sp) => {
					sp >>>= 0;
					storeValue(sp + 24, loadString(sp + 8));
				},

				// func valueGet(v ref, p string) ref
				"syscall/js.valueGet": (sp) => {
					sp >>>= 0;
					const result = Reflect.get(loadValue(sp + 8), loadString(sp + 16));
					sp = this._inst.exports.getsp() >>> 0; // see comment above
					storeValue(sp + 32, result);
				},

				// func valueSet(v ref, p string, x ref)
				"syscall/js.valueSet": (sp) => {
					sp >>>= 0;
					Reflect.set(loadValue(sp + 8), loadString(sp + 16), loadValue(sp + 32));
				},

				// func valueDelete(v ref, p string)
				"syscall/js.valueDelete": (sp) => {
					sp >>>= 0;
					Reflect.deleteProperty(loadValue(sp + 8), loadString(sp + 16));
				},

				// func valueIndex(v ref, i int) ref
				"syscall/js.valueIndex": (sp) => {
					sp >>>= 0;
					storeValue(sp + 24, Reflect.get(loadValue(sp + 8), getInt64(sp + 16)));
				},

				// valueSetIndex(v ref, i int, x ref)
				"syscall/js.valueSetIndex": (sp) => {
					sp >>>= 0;
					Reflect.set(loadValue(sp + 8), getInt64(sp + 16), loadValue(sp + 24));
				},

				// func valueCall(v ref, m string, args []ref) (ref, bool)
				"syscall/js.valueCall": (sp) => {
					sp >>>= 0;
					try {
						const v = loadValue(sp + 8);
						const m = Reflect.get(v, loadString(sp + 16));
						const args = loadSliceOfValues(sp + 32);
						const result = Reflect.apply(m, v, args);
						sp = this._inst.exports.getsp() >>> 0; // see comment above
						storeValue(sp + 56, result);
						this.mem.setUint8(sp + 64, 1);
					} catch (err) {
						sp = this._inst.exports.getsp() >>> 0; // see comment above
						storeValue(sp + 56, err);
						this.mem.setUint8(sp + 64, 0);
					}
				},

				// func valueInvoke(v ref, args []ref) (ref, bool)
				"syscall/js.valueInvoke": (sp) => {
					sp >>>= 0;
					try {
						const v = loadValue(sp + 8);
						const args = loadSliceOfValues(sp + 16);
						const result = Reflect.apply(v, undefined, args);
						sp = this._inst.exports.getsp() >>> 0; // see comment above
						storeValue(sp + 40, result);
						this.mem.setUint8(sp + 48, 1);
					} catch (err) {
						sp = this._inst.exports.getsp() >>> 0; // see comment above
						storeValue(sp + 40, err);
						this.mem.setUint8(sp + 48, 0);
					}
				},

				// func valueNew(v ref, args []ref) (ref, bool)
				"syscall/js.valueNew": (sp) => {
					sp >>>= 0;
					try {
						const v = loadValue(sp + 8);
						const args = loadSliceOfValues(sp + 16);
						const result = Reflect.construct(v, args);
						sp = this._inst.exports.getsp() >>> 0; // see comment above
						storeValue(sp + 40, result);
						this.mem.setUint8(sp + 48, 1);
					} catch (err) {
						sp = this._inst.exports.getsp() >>> 0; // see comment above
						storeValue(sp + 40, err);
						this.mem.setUint8(sp + 48, 0);
					}
				},

				// func valueLength(v ref) int
				"syscall/js.valueLength": (sp) => {
					sp >>>= 0;
					setInt64(sp + 16, parseInt(loadValue(sp + 8).length));
				},

				// valuePrepareString(v ref) (ref, int)
				"syscall/js.valuePrepareString": (sp) => {
					sp >>>= 0;
					const str = this.encoder.encode(String(loadValue(sp + 8)));
					storeValue(sp + 16, str);
					setInt64(sp + 24, str.length);
				},

				// valueLoadString(v ref, b []byte)
				"syscall/js.valueLoadString": (sp) => {
					sp >>>= 0;
					const str = loadValue(sp + 8);
					loadSlice(sp + 16).set(str);
				},

				// func valueInstanceOf(v ref, t ref) bool
				"syscall/js.valueInstanceOf": (sp) => {
					sp >>>= 0;
					this.mem.setUint8(sp + 24, (loadValue(sp + 8) instanceof loadValue(sp + 16)) ? 1 : 0);
				},

				// func copyBytesToGo(dst []byte, src ref) (int, bool)
				"syscall/js.copyBytesToGo": (sp) => {
					sp >>>= 0;
					const dst = loadSlice(sp + 8);
					const src = loadValue(sp + 32);
					if (!(src instanceof Uint8Array || src instanceof Uint8ClampedArray)) {
						this.mem.setUint8(sp + 48, 0);
						return;
					}
					const toCopy = src.subarray(0, dst.length);
					dst.set(toCopy);
					setInt64(sp + 40, toCopy.length);
					this.mem.setUint8(sp + 48, 1);
				},

				// func copyBytesToJS(dst ref, src []byte) (int, bool)
				"syscall/js.copyBytesToJS": (sp) => {
					sp >>>= 0;
					const dst = loadValue(sp + 8);
					const src = loadSlice(sp + 16);
					if (!(dst instanceof Uint8Array || dst instanceof Uint8ClampedArray)) {
						this.mem.setUint8(sp + 48, 0);
						return;
					}
					const toCopy = src.subarray(0, dst.length);
					dst.set(toCopy);
					setInt64(sp + 40, toCopy.length);
					this.mem.setUint8(sp + 48, 1);
				},

				"debug": (value) => {
					console.log(value);
				},
			}
		};
	}

	run(instance) {
		this._inst = instance;
		this.mem = new DataView(this._inst.exports.mem.buffer);
		this._values = [ // JS values that Go currently has references to, indexed by reference id
			NaN,
			0,
			null,
			true,
			false,
			global,
			this,
		];
		this._goRefCounts = new Array(this._values.length).fill(Infinity); // number of references that Go has to a JS value, indexed by reference id
		this._ids = new Map([ // mapping from JS values to reference ids
			[0, 1],
			[null, 2],
			[true, 3],
			[false, 4],
			[global, 5],
			[this, 6],
		]);
		this._idPool = [];      // unused ids that have been garbage collected
		this.exited = false;    // whether the Go program has exited

		// Pass command line arguments and environment variables to WebAssembly by writing them to the linear memory.
		let offset = 4096;

		const strPtr = (str) => {
			const ptr = offset;
			const bytes = this.encoder.encode(str + "\0");
			new Uint8Array(this.mem.buffer, offset, bytes.length).set(bytes);
			offset += bytes.length;
			if (offset % 8 !== 0) {
				offset += 8 - (offset % 8);
			}
			return ptr;
		};

		const argc = this.argv.length;
		const argvPtrs = [];
		this.argv.forEach((arg) => {
			argvPtrs.push(strPtr(arg));
		});
		argvPtrs.push(0);
		const keys = Object.keys(this.env).sort();
		keys.forEach((key) => {
			argvPtrs.push(strPtr(`${key}=${this.env[key]}`));
		});
		argvPtrs.push(0);

		const argv = offset;
		argvPtrs.forEach((ptr) => {
			this.mem.setUint32(offset, ptr, true);
			this.mem.setUint32(offset + 4, 0, true);
			offset += 8;
		});

		this._inst.exports.run(argc, argv);
	}
	
	_resume() {
		this._inst.exports.resume();
	}

	_makeFuncWrapper(id) {
		const go = this;
		return function () {
			const event = { id: id, this: this, args: arguments };
			go._pendingEvent = event;
			go._resume();
			return event.result;
		};
	}
}


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

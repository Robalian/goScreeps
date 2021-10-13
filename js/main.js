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

global.Go = class {
	constructor() {
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
		this._callbackTimeouts = new Map();
		this._nextCallbackTimeoutID = 1;

		const mem = () => {
			// The buffer may change when requesting more memory.
			return new DataView(this._inst.exports.memory.buffer);
		}

		const setInt64 = (addr, v) => {
			mem().setUint32(addr + 0, v, true);
			mem().setUint32(addr + 4, Math.floor(v / 4294967296), true);
		}

		const getInt64 = (addr) => {
			const low = mem().getUint32(addr + 0, true);
			const high = mem().getInt32(addr + 4, true);
			return low + high * 4294967296;
		}

		const loadValue = (addr) => {
			const f = mem().getFloat64(addr, true);
			if (f === 0) {
				return undefined;
			}
			if (!isNaN(f)) {
				return f;
			}

			const id = mem().getUint32(addr, true);
			return this._values[id];
		}

		const storeValue = (addr, v) => {
			const nanHead = 0x7FF80000;

			if (typeof v === "number") {
				if (isNaN(v)) {
					mem().setUint32(addr + 4, nanHead, true);
					mem().setUint32(addr, 0, true);
					return;
				}
				if (v === 0) {
					mem().setUint32(addr + 4, nanHead, true);
					mem().setUint32(addr, 1, true);
					return;
				}
				mem().setFloat64(addr, v, true);
				return;
			}

			switch (v) {
				case undefined:
					mem().setFloat64(addr, 0, true);
					return;
				case null:
					mem().setUint32(addr + 4, nanHead, true);
					mem().setUint32(addr, 2, true);
					return;
				case true:
					mem().setUint32(addr + 4, nanHead, true);
					mem().setUint32(addr, 3, true);
					return;
				case false:
					mem().setUint32(addr + 4, nanHead, true);
					mem().setUint32(addr, 4, true);
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
			let typeFlag = 1;
			switch (typeof v) {
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
			mem().setUint32(addr + 4, nanHead | typeFlag, true);
			mem().setUint32(addr, id, true);
		}

		const loadSlice = (array, len, cap) => {
			return new Uint8Array(this._inst.exports.memory.buffer, array, len);
		}

		const loadSliceOfValues = (array, len, cap) => {
			const a = new Array(len);
			for (let i = 0; i < len; i++) {
				a[i] = loadValue(array + i * 8);
			}
			return a;
		}

		const loadString = (ptr, len) => {
			return this.decoder.decode(new DataView(this._inst.exports.memory.buffer, ptr, len));
		}

		const timeOrigin = Date.now();
		this.importObject = {
			wasi_snapshot_preview1: {
				// https://github.com/WebAssembly/WASI/blob/main/phases/snapshot/docs.md#fd_write
				fd_write: function(fd, iovs_ptr, iovs_len, nwritten_ptr) {
					return 0;
				},
				proc_exit: (code) => {
					throw 'trying to exit with code ' + code;
				},
				random_get: (bufPtr, bufLen) => {
					crypto.getRandomValues(loadSlice(bufPtr, bufLen));
					return 0;
				},
			},
			env: {
				// func ticks() float64
				"runtime.ticks": () => {
					return Date.now() - timeOrigin
				},

				// func sleepTicks(timeout float64)
				"runtime.sleepTicks": (timeout) => {
					// Do not sleep, only reactivate scheduler after the given timeout.
					setTimeout(this._inst.exports.go_scheduler, timeout);
				},

				// func finalizeRef(v ref)
				"syscall/js.finalizeRef": (v_addr) => {
					// Note: TinyGo does not support finalizers so this is only called
					// for one specific case, by js.go:jsString.
					const id = mem().getUint32(v_addr, true);
					this._goRefCounts[id]--;
					if (this._goRefCounts[id] === 0) {
						const v = this._values[id];
						this._values[id] = null;
						this._ids.delete(v);
						this._idPool.push(id);
					}
				},

				// func stringVal(value string) ref
				"syscall/js.stringVal": (ret_ptr, value_ptr, value_len) => {
					const s = loadString(value_ptr, value_len);
					storeValue(ret_ptr, s);
				},

				// func valueGet(v ref, p string) ref
				"syscall/js.valueGet": (retval, v_addr, p_ptr, p_len) => {
					let prop = loadString(p_ptr, p_len);
					let value = loadValue(v_addr);
					let result = Reflect.get(value, prop);
					storeValue(retval, result);
				},

				// func valueSet(v ref, p string, x ref)
				"syscall/js.valueSet": (v_addr, p_ptr, p_len, x_addr) => {
					const v = loadValue(v_addr);
					const p = loadString(p_ptr, p_len);
					const x = loadValue(x_addr);
					Reflect.set(v, p, x);
				},

				// func valueDelete(v ref, p string)
				"syscall/js.valueDelete": (v_addr, p_ptr, p_len) => {
					const v = loadValue(v_addr);
					const p = loadString(p_ptr, p_len);
					Reflect.deleteProperty(v, p);
				},

				// func valueIndex(v ref, i int) ref
				"syscall/js.valueIndex": (ret_addr, v_addr, i) => {
					storeValue(ret_addr, Reflect.get(loadValue(v_addr), i));
				},

				// valueSetIndex(v ref, i int, x ref)
				"syscall/js.valueSetIndex": (v_addr, i, x_addr) => {
					Reflect.set(loadValue(v_addr), i, loadValue(x_addr));
				},

				// func valueCall(v ref, m string, args []ref) (ref, bool)
				"syscall/js.valueCall": (ret_addr, v_addr, m_ptr, m_len, args_ptr, args_len, args_cap) => {
					const v = loadValue(v_addr);
					const name = loadString(m_ptr, m_len);
					const args = loadSliceOfValues(args_ptr, args_len, args_cap);
					try {
						const m = Reflect.get(v, name);
						storeValue(ret_addr, Reflect.apply(m, v, args));
						mem().setUint8(ret_addr + 8, 1);
					} catch (err) {
						storeValue(ret_addr, err);
						mem().setUint8(ret_addr + 8, 0);
					}
				},

				// func valueInvoke(v ref, args []ref) (ref, bool)
				"syscall/js.valueInvoke": (ret_addr, v_addr, args_ptr, args_len, args_cap) => {
					try {
						const v = loadValue(v_addr);
						const args = loadSliceOfValues(args_ptr, args_len, args_cap);
						storeValue(ret_addr, Reflect.apply(v, undefined, args));
						mem().setUint8(ret_addr + 8, 1);
					} catch (err) {
						storeValue(ret_addr, err);
						mem().setUint8(ret_addr + 8, 0);
					}
				},

				// func valueNew(v ref, args []ref) (ref, bool)
				"syscall/js.valueNew": (ret_addr, v_addr, args_ptr, args_len, args_cap) => {
					const v = loadValue(v_addr);
					const args = loadSliceOfValues(args_ptr, args_len, args_cap);
					try {
						storeValue(ret_addr, Reflect.construct(v, args));
						mem().setUint8(ret_addr + 8, 1);
					} catch (err) {
						storeValue(ret_addr, err);
						mem().setUint8(ret_addr+ 8, 0);
					}
				},

				// func valueLength(v ref) int
				"syscall/js.valueLength": (v_addr) => {
					return loadValue(v_addr).length;
				},

				// valuePrepareString(v ref) (ref, int)
				"syscall/js.valuePrepareString": (ret_addr, v_addr) => {
					const s = String(loadValue(v_addr));
					const str = this.encoder.encode(s);
					storeValue(ret_addr, str);
					setInt64(ret_addr + 8, str.length);
				},

				// valueLoadString(v ref, b []byte)
				"syscall/js.valueLoadString": (v_addr, slice_ptr, slice_len, slice_cap) => {
					const str = loadValue(v_addr);
					loadSlice(slice_ptr, slice_len, slice_cap).set(str);
				},

				// func valueInstanceOf(v ref, t ref) bool
				"syscall/js.valueInstanceOf": (v_addr, t_addr) => {
					return loadValue(v_addr) instanceof loadValue(t_addr);
				},

				// func copyBytesToGo(dst []byte, src ref) (int, bool)
				"syscall/js.copyBytesToGo": (ret_addr, dest_addr, dest_len, dest_cap, source_addr) => {
					let num_bytes_copied_addr = ret_addr;
					let returned_status_addr = ret_addr + 4; // Address of returned boolean status variable

					const dst = loadSlice(dest_addr, dest_len);
					const src = loadValue(source_addr);
					if (!(src instanceof Uint8Array)) {
						mem().setUint8(returned_status_addr, 0); // Return "not ok" status
						return;
					}
					const toCopy = src.subarray(0, dst.length);
					dst.set(toCopy);
					setInt64(num_bytes_copied_addr, toCopy.length);
					mem().setUint8(returned_status_addr, 1); // Return "ok" status
				},

				// copyBytesToJS(dst ref, src []byte) (int, bool)
				// Originally copied from upstream Go project, then modified:
				//   https://github.com/golang/go/blob/3f995c3f3b43033013013e6c7ccc93a9b1411ca9/misc/wasm/wasm_exec.js#L404-L416
				"syscall/js.copyBytesToJS": (ret_addr, dest_addr, source_addr, source_len, source_cap) => {
					let num_bytes_copied_addr = ret_addr;
					let returned_status_addr = ret_addr + 4; // Address of returned boolean status variable

					const dst = loadValue(dest_addr);
					const src = loadSlice(source_addr, source_len);
					if (!(dst instanceof Uint8Array)) {
						mem().setUint8(returned_status_addr, 0); // Return "not ok" status
						return;
					}
					const toCopy = src.subarray(0, dst.length);
					dst.set(toCopy);
					setInt64(num_bytes_copied_addr, toCopy.length);
					mem().setUint8(returned_status_addr, 1); // Return "ok" status
				},
			}
		};
	}

	run(instance) {
		this._inst = instance;
		this._values = [ // JS values that Go currently has references to, indexed by reference id
			NaN,
			0,
			null,
			true,
			false,
			global,
			this,
		];
		this._goRefCounts = []; // number of references that Go has to a JS value, indexed by reference id
		this._ids = new Map();  // mapping from JS values to reference ids
		this._idPool = [];      // unused ids that have been garbage collected
		this.exited = false;    // whether the Go program has exited
		this.started = false

		this._inst.exports._start();
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

	async defaultRun(instance) {
		this._inst = instance;
		this._values = [ // JS values that Go currently has references to, indexed by reference id
			NaN,
			0,
			null,
			true,
			false,
			global,
			this,
		];
		this._goRefCounts = []; // number of references that Go has to a JS value, indexed by reference id
		this._ids = new Map();  // mapping from JS values to reference ids
		this._idPool = [];      // unused ids that have been garbage collected
		this.exited = false;    // whether the Go program has exited

		const mem = new DataView(this._inst.exports.memory.buffer)

		while (true) {
			const callbackPromise = new Promise((resolve) => {
				this._resolveCallbackPromise = () => {
					if (this.exited) {
						throw new Error("bad callback: Go program has already exited");
					}
					setTimeout(resolve, 0); // make sure it is asynchronous
				};
			});
			this._inst.exports._start();
			if (this.exited) {
				break;
			}
			await callbackPromise;
		}
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
	
	runLoop();
}

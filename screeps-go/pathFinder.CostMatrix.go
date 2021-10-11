package screeps

import "syscall/js"

var costMatrixConstructor = js.Global().Get("PathFinder").Get("CostMatrix")

func NewCostMatrix() CostMatrix {
	return CostMatrix{
		ref: costMatrixConstructor.New(),
	}
}

type CostMatrix struct {
	ref js.Value
}

func (cm CostMatrix) Set(x int, y int, cost uint8) {
	cm.ref.Call("set", x, y, cost)
}

func (cm CostMatrix) Get(x int, y int) uint8 {
	return uint8(cm.ref.Call("get", x, y).Int())
}

func (cm CostMatrix) Clone() CostMatrix {
	return CostMatrix{
		ref: cm.ref.Call("clone"),
	}
}

func (cm CostMatrix) Serialize() []uint32 {
	jsResult := cm.ref.Call("serialize")
	length := jsResult.Length()
	result := make([]uint32, length)
	for i := 0; i < length; i++ {
		result[i] = uint32(jsResult.Index(i).Int())
	}
	return result
}

func CostMatrixDeserialize(val []uint32) CostMatrix {
	jsVal := []interface{}{}
	length := len(val)
	for i := 0; i < length; i++ {
		jsVal[i] = val[i]
	}

	jsResult := costMatrixConstructor.Call("deserialize", jsVal)
	return CostMatrix{
		ref: jsResult,
	}
}

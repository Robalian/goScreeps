package screeps

import "syscall/js"

type interShardMemory struct {
	ref js.Value
}

var InterShardMemory *interShardMemory = func() *interShardMemory {
	if IsMMO() {
		return &interShardMemory{
			ref: js.Global().Get("InterShardMemory"),
		}
	} else {
		return nil
	}
}()

func (ism interShardMemory) GetLocal() string {
	return ism.ref.Call("getLocal").String()
}

func (ism interShardMemory) SetLocal(value string) {
	ism.ref.Call("setLocal", value)
}

func (ism interShardMemory) GetRemote(shard string) *string {
	jsResult := ism.ref.Call("getRemote", shard)
	if jsResult.IsNull() {
		return nil
	} else {
		result := jsResult.String()
		return &result
	}
}

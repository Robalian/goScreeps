package screeps

import "syscall/js"

type Store struct {
	ref js.Value
}

func (store Store) GetCapacity(resource *ResourceConstant) *int {
	var callResult js.Value
	if resource == nil {
		callResult = store.ref.Call("getCapacity")
	} else {
		callResult = store.ref.Call("getCapacity", string(*resource))
	}

	if callResult.IsNull() {
		return nil
	} else {
		result := new(int)
		*result = callResult.Int()
		return result
	}
}

func (store Store) GetFreeCapacity(resource *ResourceConstant) *int {
	var callResult js.Value
	if resource == nil {
		callResult = store.ref.Call("getFreeCapacity")
	} else {
		callResult = store.ref.Call("getFreeCapacity", string(*resource))
	}

	if callResult.IsNull() {
		return nil
	} else {
		result := new(int)
		*result = callResult.Int()
		return result
	}
}

func (store Store) GetUsedCapacity(resource *ResourceConstant) *int {
	var callResult js.Value
	if resource == nil {
		callResult = store.ref.Call("getUsedCapacity")
	} else {
		callResult = store.ref.Call("getUsedCapacity", string(*resource))
	}

	if callResult.IsNull() {
		return nil
	} else {
		result := new(int)
		*result = callResult.Int()
		return result
	}
}

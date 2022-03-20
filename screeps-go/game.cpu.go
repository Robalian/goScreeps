package screeps

import "syscall/js"

type Cpu struct {
	ref          js.Value
	Limit        int
	TickLimit    int
	Bucket       int
	ShardLimits  *map[string]int
	Unlocked     *bool // doesn't exist on private servers
	UnlockedTime *int
}

type HeapStatistics struct {
	Total_heap_size            int
	Total_heap_size_executable int
	Total_physical_size        int
	Total_available_size       int
	Used_heap_size             int
	Heap_size_limit            int
	Malloced_memory            int
	Peak_malloced_memory       int
	Does_zap_garbage           int
	Externally_allocated_size  int
}

func (cpu Cpu) GetHeapStatistics() HeapStatistics {
	jsHeapStatistics := cpu.ref.Call("getHeapStatistics")
	return HeapStatistics{
		Total_heap_size:            jsHeapStatistics.Get("total_heap_size").Int(),
		Total_heap_size_executable: jsHeapStatistics.Get("total_heap_size_executable").Int(),
		Total_physical_size:        jsHeapStatistics.Get("total_physical_size").Int(),
		Total_available_size:       jsHeapStatistics.Get("total_available_size").Int(),
		Used_heap_size:             jsHeapStatistics.Get("used_heap_size").Int(),
		Heap_size_limit:            jsHeapStatistics.Get("heap_size_limit").Int(),
		Malloced_memory:            jsHeapStatistics.Get("malloced_memory").Int(),
		Peak_malloced_memory:       jsHeapStatistics.Get("peak_malloced_memory").Int(),
		Externally_allocated_size:  jsHeapStatistics.Get("externally_allocated_size").Int(),
	}
}

func (cpu Cpu) GetUsed() float64 {
	return cpu.ref.Call("getUsed").Float()
}

func (cpu Cpu) Halt() {
	cpu.ref.Call("halt")
}

func (cpu Cpu) SetShardLimits(limits map[string]int) ErrorCode {
	jsLimits := map[string]interface{}{}
	for k, v := range limits {
		jsLimits[k] = v
	}

	result := cpu.ref.Call("setShardLimits", jsLimits).Int()
	return ErrorCode(result)
}

func (cpu Cpu) Unlock() ErrorCode {
	result := cpu.ref.Call("unlock").Int()
	return ErrorCode(result)
}

func (cpu Cpu) GeneratePixel() ErrorCode {
	result := cpu.ref.Call("generatePixel").Int()
	return ErrorCode(result)
}

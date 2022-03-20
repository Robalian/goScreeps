package src

import (
	"encoding/json"
	. "screepsgo/screeps-go"
	"strings"
)

var globalResetTick = Game.Time()

func removeWhiteSpace(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != '\t' && s[i] != ' ' && s[i] != '\r' && s[i] != '\n' {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

func setGrafanaStats() {
	//*
	bytes, _ := json.Marshal(map[string]interface{}{
		"globalAge": Game.Time() - globalResetTick,
		"profiling": map[string]interface{}{
			"bucket":         Game.Cpu().Bucket,
			"limit":          Game.Cpu().Limit,
			"tick":           Game.Cpu().GetUsed(),
			"heapStatistics": Game.Cpu().GetHeapStatistics(),
			"memory":         len(RawMemory.Get()),
		},
	})
	(*RawMemory.Segments())[99] = string(bytes)

	/*/

	heapStatistics := Game.Cpu().GetHeapStatistics()
	(*RawMemory.Segments())[99] = removeWhiteSpace(fmt.Sprintf(`{
		"globalAge": %d,
		"profiling": {
			"bucket": %d,
			"limit": %d,
			"heapStatistics": {
				"total_heap_size": %d,
				"total_heap_size_executable": %d,
				"total_physical_size": %d,
				"total_available_size": %d,
				"used_heap_size": %d,
				"heap_size_limit": %d,
				"malloced_memory": %d,
				"peak_malloced_memory": %d,
				"externally_allocated_size": %d
			},
			"memory": %d
		}
	}`,
		Game.Time()-globalResetTick,
		Game.Cpu().Bucket,
		Game.Cpu().Limit,
		heapStatistics.Total_heap_size,
		heapStatistics.Total_heap_size_executable,
		heapStatistics.Total_physical_size,
		heapStatistics.Total_available_size,
		heapStatistics.Used_heap_size,
		heapStatistics.Heap_size_limit,
		heapStatistics.Malloced_memory,
		heapStatistics.Peak_malloced_memory,
		heapStatistics.Externally_allocated_size,
		len(RawMemory.Get()),
	))
	//*/
}

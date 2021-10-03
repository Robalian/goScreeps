package screeps

import (
	"strconv"
	"syscall/js"
)

type ForeignSegment struct {
	Username string
	Id       int
	Data     string
}

type rawMemory struct {
	ref      js.Value
	segments map[int]string
}

func (rm rawMemory) Segments() *map[int]string {
	return &rm.segments
}

func (rm rawMemory) ForeignSegment() *ForeignSegment {
	var jsForeignSegment = rm.ref.Get("ForeignSegment")
	if jsForeignSegment.IsNull() {
		return nil
	} else {
		var result = new(ForeignSegment)
		result.Username = jsForeignSegment.Get("username").String()
		result.Id = jsForeignSegment.Get("id").Int()
		result.Data = jsForeignSegment.Get("data").String()
		return result
	}
}

func (rm rawMemory) Get() string {
	return rm.ref.Call("get").String()
}

func (rm rawMemory) Set(value string) {
	rm.ref.Call("set", value)
}

func (rm rawMemory) SetActiveSegments(ids []int) {
	rm.ref.Call("setActiveSegments", ids)
}

func (rm rawMemory) SetActiveForeignSegment(username string, id *int) {
	var jsId js.Value
	if id == nil {
		jsId = js.Null()
	} else {
		jsId = js.ValueOf(*id)
	}
	rm.ref.Call("setActiveForeignSegment", username, jsId)
}

func (rm rawMemory) SetDefaultPublicSegment(id int) {
	rm.ref.Call("setDefaultPublicSegment", id)
}

func (rm rawMemory) SetPublicSegments(ids []int) {
	rm.ref.Call("setPublicSegments", ids)
}

var RawMemory = rawMemory{
	ref:      js.Global().Get("RawMemory"),
	segments: map[int]string{},
}

func loadSegments() {
	for k := range RawMemory.segments {
		delete(RawMemory.segments, k)
	}

	var jsSegments = RawMemory.ref.Get("segments")
	var entries = object.Call("entries", jsSegments)
	var length = entries.Get("length").Int()
	for i := 0; i < length; i++ {
		var v = entries.Index(i)
		var k, _ = strconv.Atoi(v.Index(0).String())
		RawMemory.segments[k] = v.Index(1).String()
		jsSegments.Set(v.Index(0).String(), js.Undefined())
	}
}

func saveSegments() {
	var jsSegments = RawMemory.ref.Get("segments")
	for k := range RawMemory.segments {
		jsSegments.Set(strconv.Itoa(k), RawMemory.segments[k])
	}
}

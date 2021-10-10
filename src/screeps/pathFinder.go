package screeps

import (
	"syscall/js"
)

type PathFinderGoal struct {
	Pos   RoomPosition
	Range int
}
type RoomCallback func(roomName string) *CostMatrix
type PathFinderOpts struct {
	RoomCallback    *RoomCallback
	PlainCost       *uint8
	SwampCost       *uint8
	Flee            *bool
	MaxOps          *uint
	MaxRooms        *uint // default is 16, maximum is 64
	MaxCost         *uint
	HeuristicWeight *float64
}
type PathFinderSearchResult struct {
	Path       []RoomPosition
	Ops        int
	Cost       int
	Incomplete bool
}

type pathFinder struct {
	ref js.Value
}

var PathFinder = pathFinder{
	ref: js.Global().Get("PathFinder"),
}

func (pf pathFinder) Search(origin RoomPosition, goal []PathFinderGoal, opts *PathFinderOpts) PathFinderSearchResult {
	var jsOpts js.Value
	if opts == nil {
		jsOpts = js.Undefined()
	} else {
		tmpOpts := map[string]interface{}{}

		if opts.RoomCallback != nil {
			currentRoomCallback = opts.RoomCallback
			tmpOpts["roomCallback"] = js.Global().Get("jsRoomCallback")
		}

		if opts.PlainCost != nil {
			tmpOpts["plainCost"] = *opts.PlainCost
		}

		if opts.SwampCost != nil {
			tmpOpts["swampCost"] = *opts.SwampCost
		}

		if opts.Flee != nil {
			tmpOpts["flee"] = *opts.Flee
		}

		if opts.MaxOps != nil {
			tmpOpts["maxOps"] = *opts.MaxOps
		}

		if opts.MaxRooms != nil {
			tmpOpts["maxRooms"] = *opts.MaxRooms
		}

		if opts.MaxCost != nil {
			tmpOpts["maxCost"] = *opts.MaxCost
		}

		if opts.HeuristicWeight != nil {
			tmpOpts["heuristicWeight"] = *opts.HeuristicWeight
		}

		jsOpts = js.ValueOf(tmpOpts)
	}

	goalLength := len(goal)
	jsGoal := make([]interface{}, goalLength)
	for i := 0; i < goalLength; i++ {
		jsGoal[i] = map[string]interface{}{
			"pos":   goal[i].Pos.ref,
			"range": goal[i].Range,
		}
	}

	jsResult := pf.ref.Call("search", origin.ref, jsGoal, jsOpts)
	jsPath := jsResult.Get("path")
	pathLength := jsPath.Length()
	path := make([]RoomPosition, pathLength)
	for i := 0; i < pathLength; i++ {
		path[i] = makeRoomPosition(jsPath.Index(i))
	}
	result := PathFinderSearchResult{
		Path:       path,
		Ops:        jsResult.Get("ops").Int(),
		Cost:       jsResult.Get("cost").Int(),
		Incomplete: jsResult.Get("incomplete").Bool(),
	}

	return result
}

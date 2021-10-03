package screeps

import "syscall/js"

var object = js.Global().Get("Object")

func packLineStyle(style LineStyle) map[string]interface{} {
	var result = map[string]interface{}{}
	if style.Color != nil {
		result["color"] = *style.Color
	}
	if style.Width != nil {
		result["width"] = *style.Width
	}
	if style.Opacity != nil {
		result["opacity"] = *style.Opacity
	}
	if style.LineStyle != nil {
		result["lineStyle"] = *style.LineStyle
	}
	return result
}

func packCircleStyle(style CircleStyle) map[string]interface{} {
	var result = map[string]interface{}{}
	if style.Radius != nil {
		result["radius"] = *style.Radius
	}
	if style.Fill != nil {
		result["fill"] = *style.Fill
	}
	if style.Opacity != nil {
		result["opacity"] = *style.Opacity
	}
	if style.Stroke != nil {
		result["stroke"] = *style.Stroke
	}
	if style.StrokeWidth != nil {
		result["strokeWidth"] = *style.StrokeWidth
	}
	if style.LineStyle != nil {
		result["lineStyle"] = *style.LineStyle
	}
	return result
}

func packRectStyle(style RectStyle) map[string]interface{} {
	var result = map[string]interface{}{}
	if style.Fill != nil {
		result["fill"] = *style.Fill
	}
	if style.Opacity != nil {
		result["opacity"] = *style.Opacity
	}
	if style.Stroke != nil {
		result["stroke"] = *style.Stroke
	}
	if style.StrokeWidth != nil {
		result["strokeWidth"] = *style.StrokeWidth
	}
	if style.LineStyle != nil {
		result["lineStyle"] = *style.LineStyle
	}
	return result
}

func packPolyStyle(style PolyStyle) map[string]interface{} {
	var result = map[string]interface{}{}
	if style.Fill != nil {
		result["fill"] = *style.Fill
	}
	if style.LineStyle != nil {
		result["lineStyle"] = *style.LineStyle
	}
	if style.Opacity != nil {
		result["opacity"] = *style.Opacity
	}
	if style.Stroke != nil {
		result["stroke"] = *style.Stroke
	}
	if style.StrokeWidth != nil {
		result["strokeWidth"] = *style.StrokeWidth
	}
	return result
}
func packFindPathOpts(opts FindPathOpts) map[string]interface{} {
	var result = map[string]interface{}{}

	if opts.HeuristicWeight != nil {
		result["heuristicWeight"] = *opts.HeuristicWeight
	}
	if opts.IgnoreCreeps != nil {
		result["ignoreCreeps"] = *opts.IgnoreCreeps
	}
	if opts.IgnoreDestructibleStructures != nil {
		result["ignoreDestructibleStructures"] = *opts.IgnoreDestructibleStructures
	}
	if opts.IgnoreRoads != nil {
		result["ignoreRoads"] = *opts.IgnoreRoads
	}
	if opts.MaxOps != nil {
		result["maxOps"] = *opts.MaxOps
	}
	if opts.MaxRooms != nil {
		result["maxRooms"] = *opts.MaxRooms
	}
	if opts.PlainCost != nil {
		result["plainCost"] = *opts.PlainCost
	}
	if opts.Range != nil {
		result["range"] = *opts.Range
	}
	if opts.Serialize != nil {
		result["serialize"] = *opts.Serialize
	}
	if opts.SwampCost != nil {
		result["swampCost"] = *opts.SwampCost
	}

	return result
}

func packMoveToOpts(opts MoveToOpts) map[string]interface{} {
	var result = packFindPathOpts(opts.FindPathOpts)

	if opts.NoPathFinding != nil {
		result["noPathFinding"] = *opts.NoPathFinding
	}
	if opts.ReusePath != nil {
		result["reusePath"] = *opts.ReusePath
	}
	if opts.SerializeMemory != nil {
		result["serializeMemory"] = *opts.SerializeMemory
	}
	if opts.VisualizePathStyle != nil {
		result["visualizePathStyle"] = packPolyStyle(*opts.VisualizePathStyle)
	}

	return result
}

func packFindPathResult(path FindPathResult) []interface{} {
	var pathLength = len(path)
	var result = make([]interface{}, pathLength)
	for i := 0; i < pathLength; i++ {
		var step = map[string]interface{}{}
		step["x"] = path[i].x
		step["y"] = path[i].y
		step["dx"] = path[i].dx
		step["dy"] = path[i].dy
		step["direction"] = int(path[i].direction)
		result[i] = step
	}
	return result
}

func unpackFindPathResult(path js.Value) FindPathResult {
	var length = path.Length()
	var result = make(FindPathResult, length)
	for i := 0; i < length; i++ {
		var v = path.Index(i)
		result[i] = struct {
			x         int
			y         int
			dx        int
			dy        int
			direction DirectionConstant
		}{
			x:         v.Get("x").Int(),
			y:         v.Get("y").Int(),
			dx:        v.Get("dx").Int(),
			dy:        v.Get("dy").Int(),
			direction: DirectionConstant(v.Get("direction").Int()),
		}
	}
	return result
}

func unpackLookAtResult(lookAtResult js.Value) LookAtResult {
	var length = lookAtResult.Length()
	var result = LookAtResult{}
	for i := 0; i < length; i++ {
		var v = lookAtResult.Index(i)
		var t = v.Get("type").String()
		if t == "terrain" {
			continue
		}

		result = append(result, struct {
			Type   string
			Object RoomObject
		}{
			Type: t,
			Object: RoomObject{
				ref: v.Get(t),
			},
		})
	}
	return result
}

func unpackLookForAtResult(lookForAtResult js.Value) []RoomObject {
	var length = lookForAtResult.Length()
	var result = make([]RoomObject, length)
	for i := 0; i < length; i++ {
		result[i] = RoomObject{
			ref: lookForAtResult.Index(i),
		}
	}
	return result
}

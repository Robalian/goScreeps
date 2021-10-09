package screeps

import "syscall/js"

var object = js.Global().Get("Object")

var isMMO = !js.Global().Get("Game").Get("cpu").Get("generatePixel").IsUndefined()

func IsMMO() bool {
	return isMMO
}

func packLineStyle(style LineStyle) map[string]interface{} {
	result := map[string]interface{}{}
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
	result := map[string]interface{}{}
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
	result := map[string]interface{}{}
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
	result := map[string]interface{}{}
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

func packTextStyle(style TextStyle) map[string]interface{} {
	result := map[string]interface{}{}
	if style.Color != nil {
		result["color"] = *style.Color
	}
	if style.Font != nil {
		result["font"] = *style.Font
	}
	if style.Stroke != nil {
		result["stroke"] = *style.Stroke
	}
	if style.StrokeWidth != nil {
		result["strokeWidth"] = *style.StrokeWidth
	}
	if style.BackgroundColor != nil {
		result["backgroundColor"] = *style.BackgroundColor
	}
	if style.BackgroundPadding != nil {
		result["backgroundPadding"] = *style.BackgroundPadding
	}
	if style.Align != nil {
		result["align"] = string(*style.Align)
	}
	if style.Opacity != nil {
		result["opacity"] = *style.Opacity
	}
	return result
}

func packFindPathOpts(opts FindPathOpts) map[string]interface{} {
	result := map[string]interface{}{}

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
	result := packFindPathOpts(opts.FindPathOpts)

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
	pathLength := len(path)
	result := make([]interface{}, pathLength)
	for i := 0; i < pathLength; i++ {
		step := map[string]interface{}{}
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
	pathLength := path.Length()
	result := make(FindPathResult, pathLength)
	for i := 0; i < pathLength; i++ {
		step := path.Index(i)
		result[i] = struct {
			x         int
			y         int
			dx        int
			dy        int
			direction DirectionConstant
		}{
			x:         step.Get("x").Int(),
			y:         step.Get("y").Int(),
			dx:        step.Get("dx").Int(),
			dy:        step.Get("dy").Int(),
			direction: DirectionConstant(step.Get("direction").Int()),
		}
	}
	return result
}

func unpackLookAtResult(lookAtResult js.Value) LookAtResult {
	length := lookAtResult.Length()
	result := LookAtResult{}
	for i := 0; i < length; i++ {
		v := lookAtResult.Index(i)
		objectType := v.Get("type").String()

		// we skip terrain, for type consistency (it's not a RoomObject) + it's better accessed through Game.map anyway
		if objectType == "terrain" {
			continue
		}

		result = append(result, struct {
			Type   string
			Object RoomObject
		}{
			Type: objectType,
			Object: RoomObject{
				ref: v.Get(objectType),
			},
		})
	}
	return result
}

func unpackLookForAtResult(lookForAtResult js.Value) []RoomObject {
	length := lookForAtResult.Length()
	result := make([]RoomObject, length)
	for i := 0; i < length; i++ {
		result[i] = RoomObject{
			ref: lookForAtResult.Index(i),
		}
	}
	return result
}

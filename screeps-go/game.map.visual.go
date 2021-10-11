package screeps

import "syscall/js"

type MapVisual struct {
	ref js.Value
}

func (visual MapVisual) Line(pos1 RoomPosition, pos2 RoomPosition, style *LineStyle) MapVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packLineStyle(*style))
	}
	visual.ref.Call("line", pos1.ref, pos2.ref, jsStyle)
	return visual
}

func (visual MapVisual) Circle(pos RoomPosition, style *CircleStyle) MapVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packCircleStyle(*style))
	}
	visual.ref.Call("circle", pos.ref, jsStyle)
	return visual
}

func (visual MapVisual) Rect(topLeftPos RoomPosition, width int, height int, style *RectStyle) MapVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packRectStyle(*style))
	}
	visual.ref.Call("rect", topLeftPos.ref, width, height, jsStyle)
	return visual
}

func (visual MapVisual) Poly(points []RoomPosition, style *PolyStyle) MapVisual {
	pointsCount := len(points)
	jsPoints := make([]interface{}, pointsCount)
	for i := 0; i < pointsCount; i++ {
		jsPoints[i] = points[i].ref
	}
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packPolyStyle(*style))
	}
	visual.ref.Call("poly", jsPoints, jsStyle)
	return visual
}

func (visual MapVisual) Text(text string, pos RoomPosition, style *TextStyle) MapVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packTextStyle(*style))
	}
	visual.ref.Call("text", text, pos.ref, jsStyle)
	return visual
}

func (visual MapVisual) Clear() MapVisual {
	visual.ref.Call("clear")
	return visual
}

func (visual MapVisual) GetSize() int {
	return visual.ref.Call("getSize").Int()
}

func (visual MapVisual) Export() *string {
	jsResult := visual.ref.Call("export")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.String()
		return &result
	}
}

func (visual MapVisual) Import(val string) MapVisual {
	jsResult := visual.ref.Call("import", val)
	return MapVisual{ref: jsResult}
}

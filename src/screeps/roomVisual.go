package screeps

import "syscall/js"

type LineStyle struct {
	Width     *float32
	Color     *string
	Opacity   *float32
	LineStyle *string
}

type CircleStyle struct {
	Radius      *float32
	Fill        *string
	Opacity     *float32
	Stroke      *string
	StrokeWidth *float32
	LineStyle   *string
}

type RectStyle struct {
	Fill        *string
	Opacity     *float32
	Stroke      *string
	StrokeWidth *float32
	LineStyle   *string
}

type PolyStyle struct {
	Fill        *string
	Opacity     *float32
	Stroke      *string
	StrokeWidth *float32
	LineStyle   *string
}

type PolyPoint [2]int

func MakeRoomVisual(roomName *string) RoomVisual {
	var jsRoomName js.Value
	if roomName == nil {
		jsRoomName = js.Undefined()
	} else {
		jsRoomName = js.ValueOf(*roomName)
	}

	var result = js.Global().Get("RoomVisual").New(jsRoomName)
	return RoomVisual{
		ref: result,
	}
}

type RoomVisual struct {
	ref js.Value
}

func (visual RoomVisual) RoomName() string {
	return visual.ref.Get("roomName").String()
}

func (visual RoomVisual) Line(pos1 RoomPosition, pos2 RoomPosition, style *LineStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packLineStyle(*style))
	}
	visual.ref.Call("line", pos1.ref, pos2.ref, jsStyle)
	return visual
}

func (visual RoomVisual) Line_XY(x1 int, y1 int, x2 int, y2 int, style *LineStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packLineStyle(*style))
	}
	visual.ref.Call("line", x1, y1, x2, y2, jsStyle)
	return visual
}

func (visual RoomVisual) Circle(pos RoomPosition, style *CircleStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packCircleStyle(*style))
	}
	visual.ref.Call("circle", pos.ref, jsStyle)
	return visual
}

func (visual RoomVisual) Circle_XY(x int, y int, style *CircleStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packCircleStyle(*style))
	}
	visual.ref.Call("circle", x, y, jsStyle)
	return visual
}

func (visual RoomVisual) Rect(topLeftPos RoomPosition, width int, height int, style *RectStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packRectStyle(*style))
	}
	visual.ref.Call("rect", topLeftPos.ref, width, height, jsStyle)
	return visual
}

func (visual RoomVisual) Rect_XY(x int, y int, width int, height int, style *RectStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packRectStyle(*style))
	}
	visual.ref.Call("rect", x, y, width, height, jsStyle)
	return visual
}

func (visual RoomVisual) Poly_XY(points []PolyPoint, style *PolyStyle) RoomVisual {
	var pointLength = len(points)
	var jsPoints = make([]interface{}, pointLength)
	for i := 0; i < pointLength; i++ {
		jsPoints[i] = []interface{}{
			points[i][0],
			points[i][1],
		}
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

func (visual RoomVisual) Poly(points []RoomPosition, style *PolyStyle) RoomVisual {
	var pointLength = len(points)
	var jsPoints = make([]interface{}, pointLength)
	for i := 0; i < pointLength; i++ {
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

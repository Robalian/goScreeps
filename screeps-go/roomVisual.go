package screeps

import "syscall/js"

type LineStyle struct {
	Width     *float64
	Color     *string
	Opacity   *float64
	LineStyle *string
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

type CircleStyle struct {
	Radius      *float64
	Fill        *string
	Opacity     *float64
	Stroke      *string
	StrokeWidth *float64
	LineStyle   *string
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

type RectStyle struct {
	Fill        *string
	Opacity     *float64
	Stroke      *string
	StrokeWidth *float64
	LineStyle   *string
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

type PolyStyle struct {
	Fill        *string
	Opacity     *float64
	Stroke      *string
	StrokeWidth *float64
	LineStyle   *string
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

type TextStyleAlign string

const (
	AlignLeft   TextStyleAlign = "left"
	AlignCenter TextStyleAlign = "center"
	AlignRight  TextStyleAlign = "right"
)

type PolyPoint [2]int

type TextStyle struct {
	Color             *string
	Font              *float64 // TODO - number/string
	Stroke            *string
	StrokeWidth       *float64
	BackgroundColor   *string
	BackgroundPadding *float64
	Align             *TextStyleAlign
	Opacity           *float64
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

func NewRoomVisual(roomName *string) RoomVisual {
	var jsRoomName js.Value
	if roomName == nil {
		jsRoomName = js.Undefined()
	} else {
		jsRoomName = js.ValueOf(*roomName)
	}

	result := js.Global().Get("RoomVisual").New(jsRoomName)
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

func (visual RoomVisual) Line_XY(x1 float64, y1 float64, x2 float64, y2 float64, style *LineStyle) RoomVisual {
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

func (visual RoomVisual) Circle_XY(x float64, y float64, style *CircleStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packCircleStyle(*style))
	}
	visual.ref.Call("circle", x, y, jsStyle)
	return visual
}

func (visual RoomVisual) Rect(topLeftPos RoomPosition, width float64, height float64, style *RectStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packRectStyle(*style))
	}
	visual.ref.Call("rect", topLeftPos.ref, width, height, jsStyle)
	return visual
}

func (visual RoomVisual) Rect_XY(x float64, y float64, width float64, height float64, style *RectStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packRectStyle(*style))
	}
	visual.ref.Call("rect", x, y, width, height, jsStyle)
	return visual
}

func (visual RoomVisual) Poly(points []RoomPosition, style *PolyStyle) RoomVisual {
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

func (visual RoomVisual) Poly_XY(points []PolyPoint, style *PolyStyle) RoomVisual {
	pointsCount := len(points)
	jsPoints := make([]interface{}, pointsCount)
	for i := 0; i < pointsCount; i++ {
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

func (visual RoomVisual) Text(text string, pos RoomPosition, style *TextStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packTextStyle(*style))
	}
	visual.ref.Call("text", text, pos.ref, jsStyle)
	return visual
}

func (visual RoomVisual) Text_XY(text string, x float64, y float64, style *TextStyle) RoomVisual {
	var jsStyle js.Value
	if style == nil {
		jsStyle = js.Undefined()
	} else {
		jsStyle = js.ValueOf(packTextStyle(*style))
	}
	visual.ref.Call("text", text, x, y, jsStyle)
	return visual
}

func (visual RoomVisual) Clear() RoomVisual {
	visual.ref.Call("clear")
	return visual
}

func (visual RoomVisual) GetSize() int {
	return visual.ref.Call("getSize").Int()
}

func (visual RoomVisual) Export() *string {
	jsResult := visual.ref.Call("export")
	if jsResult.IsUndefined() {
		return nil
	} else {
		result := jsResult.String()
		return &result
	}
}

func (visual RoomVisual) Import(val string) RoomVisual {
	jsResult := visual.ref.Call("import", val)
	return RoomVisual{ref: jsResult}
}

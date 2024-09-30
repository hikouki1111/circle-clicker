package utility

import (
	"fmt"
	"math"
	"syscall/js"
)

var (
	currentCanvas js.Value
	currentMode   string
	defaultFont   = "Comfortaa"
	systemFont    = "sans-serif"
	currentCtx    js.Value
)

func BeginRender(canvas js.Value, mode string) {
	currentCanvas = canvas
	currentMode = mode
	beginPath()
	w := currentCanvas.Get("width").Float()
	h := currentCanvas.Get("height").Float()
	currentCtx.Call("clearRect", 0, 0, w, h)
	closePath()
}

func EndRender() {
	currentCanvas = js.Null()
	currentMode = ""
}

func DrawImage(x, y, w, h float32, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	img := js.Global().Get("Image").New()
	img.Set("src", "assets/bg.jpg")
	currentCtx.Call("drawImage", img, x, y, w, h)
	closePath()
}

func DrawBackground() {
	if !isReady() {
		return
	}

	beginPath()
	SetFilter("blur(25px)")
	DrawImage(0, 0, float32(currentCanvas.Get("width").Float()), float32(currentCanvas.Get("height").Float()))
	closePath()
}

func DrawRect(x, y, w, h, lineWidth float32, strokeStyle string, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	SetStrokeStyle(strokeStyle)
	SetLineWidth(lineWidth)
	StrokeRect(x, y, w, h)
	closePath()
}

func DrawFilledRect(x, y, w, h float32, fillStyle string, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	SetFillStyle(fillStyle)
	FillRect(x, y, w, h)
	closePath()
}

func DrawCircle(x, y, r, lineWidth float32, strokeStyle string, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	Arc(x, y, r, 0, math.Pi*2)
	SetLineWidth(lineWidth)
	SetStrokeStyle(strokeStyle)
	Stroke()
	closePath()
}

func DrawFilledCircle(x, y, r float32, fillStyle string, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	Arc(x, y, r, 0, math.Pi*2)
	SetFillStyle(fillStyle)
	Fill()
	closePath()
}

func DrawRoundedRect(x, y, w, h, r, lineWidth float32, strokeStyle string, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	MoveTo(x+r, y)
	LineTo(x+w-r, y)
	QuadraticCurveTo(x+w, y, x+w, y+r)
	LineTo(x+w, y+h-r)
	QuadraticCurveTo(x+w, y+h, x+w-r, y+h)
	LineTo(x+r, y+h)
	QuadraticCurveTo(x, y+h, x, y+h-r)
	LineTo(x, y+r)
	QuadraticCurveTo(x, y, x+r, y)
	SetStrokeStyle(strokeStyle)
	SetLineWidth(lineWidth)
	Stroke()
	closePath()
}

func DrawFilledRoundedRect(x, y, w, h, r float32, fillStyle string, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	MoveTo(x+r, y)
	LineTo(x+w-r, y)
	QuadraticCurveTo(x+w, y, x+w, y+r)
	LineTo(x+w, y+h-r)
	QuadraticCurveTo(x+w, y+h, x+w-r, y+h)
	LineTo(x+r, y+h)
	QuadraticCurveTo(x, y+h, x, y+h-r)
	LineTo(x, y+r)
	QuadraticCurveTo(x, y, x+r, y)
	SetFillStyle(fillStyle)
	Fill()
	closePath()
}

func DrawFilledText(text string, x, y, size float32, fillStyle string, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	SetFont(defaultFont, size)
	SetFillStyle(fillStyle)
	currentCtx.Call("fillText", text, x, y)
	closePath()
}

func DrawText(text string, x, y, lineWidth, size float32, strokeStyle string, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	SetFont(defaultFont, size)
	SetStrokeStyle(strokeStyle)
	SetLineWidth(lineWidth)
	currentCtx.Call("strokeText", text, x, y)
	closePath()
}

func DrawCenteredFilledText(text string, rx, ry, rw, rh, size float32, fillStyle string, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	fontW, fontH := GetFontSize(text, size)
	centerX, centerY := GetRectCenter(rx, ry, rw, rh, fontW, fontH)
	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	SetFont(defaultFont, size)
	SetFillStyle(fillStyle)
	currentCtx.Call("fillText", text, centerX-fontW/2, centerY+fontH)
	closePath()
}

func DrawCenteredText(text string, rx, ry, rw, rh, lineWidth, size float32, strokeStyle string, styleFunc ...func(ctx js.Value)) {
	if !isReady() {
		return
	}

	fontW, fontH := GetFontSize(text, size)
	centerX, centerY := GetRectCenter(rx, ry, rw, rh, fontW, fontH)
	beginPath()
	for _, f := range styleFunc {
		f(currentCtx)
	}
	SetFont(defaultFont, size)
	SetStrokeStyle(strokeStyle)
	SetLineWidth(lineWidth)
	currentCtx.Call("strokeText", text, centerX-fontW/2, centerY+fontH)
	closePath()
}

func GetCenter(w, h float32) (x, y float32) {
	if !isReady() {
		return
	}

	centerX := float32(currentCanvas.Get("width").Float() / 2)
	centerY := float32(currentCanvas.Get("height").Float() / 2)

	return centerX - w/2, centerY - h/2
}

func GetRectCenter(rx, ry, rw, rh, w, h float32) (x, y float32) {
	centerX := rx + rw/2
	centerY := ry + rh/2

	return centerX - w/2, centerY - h/2
}

func GetFontSize(text string, size float32) (w, h float32) {
	if !isReady() {
		return
	}

	beginPath()
	SetFont(defaultFont, size)
	width := currentCtx.Call("measureText", text).Get("width").Float()
	closePath()

	return float32(width / 2), size
}

func isReady() bool {
	return !currentCanvas.IsNull() && currentMode != ""
}

func beginPath() js.Value {
	if !isReady() {
		return js.Null()
	}

	ctx := currentCanvas.Call("getContext", currentMode)
	ctx.Call("beginPath")
	currentCtx = ctx

	return currentCtx
}

func closePath() {
	if !isReadyToPath() {
		return
	}

	SetFilter("none")
	SetFont(systemFont, 10)
	SetShadow(0, "#000000")
	currentCtx.Call("closePath")
	currentCtx = js.Null()
}

func MoveTo(x, y float32) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Call("moveTo", x, y)
}

func LineTo(x, y float32) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Call("lineTo", x, y)
}

func Arc(x, y, r, startAngle, endAngle float32, counterClockWise ...bool) {
	if !isReadyToPath() {
		return
	}

	ccw := false
	if len(counterClockWise) > 0 {
		ccw = counterClockWise[0]
	}

	currentCtx.Call("arc", x, y, r, startAngle, endAngle, ccw)
}

func QuadraticCurveTo(cp1x, cp1y, x, y float32) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Call("quadraticCurveTo", cp1x, cp1y, x, y)
}

func SetFillStyle(fillStyle string) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Set("fillStyle", fillStyle)
}

func Fill() {
	if !isReadyToPath() {
		return
	}

	currentCtx.Call("fill")
}

func SetStrokeStyle(strokeStyle string) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Set("strokeStyle", strokeStyle)
}

func SetLineWidth(lineWidth float32) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Set("lineWidth", lineWidth)
}

func SetFont(font string, size float32) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Set("font", fmt.Sprintf("%fpx %s", size, font))
}

func Stroke() {
	if !isReadyToPath() {
		return
	}

	currentCtx.Call("stroke")
}

func FillRect(x, y, w, h float32) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Call("fillRect", x, y, w, h)
}

func StrokeRect(x, y, w, h float32) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Call("strokeRect", x, y, w, h)
}

func SetFilter(filter string) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Set("filter", filter)
}

func SetShadow(blur float32, color string) {
	if !isReadyToPath() {
		return
	}

	currentCtx.Set("shadowColor", color)
	currentCtx.Set("shadowBlur", blur)
}

func isReadyToPath() bool {
	return isReady() && !currentCtx.IsNull()
}

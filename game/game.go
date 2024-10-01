package game

import (
	"circle-clicker/game/item"
	"circle-clicker/game/screen"
	"syscall/js"
)

var (
	global   = js.Global()
	document = global.Get("document")
	canvas   js.Value
)

func Start() {
	canvas = document.Call("getElementById", "main-canvas")

	resizeCanvas()
	global.Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resizeCanvas()
		return nil
	}))

	screen.TitleScreen().SetScreen(global, canvas, document)
	var tickFunc js.Func
	tickFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		tick()
		global.Call("requestAnimationFrame", tickFunc)
		return nil
	})
	global.Call("requestAnimationFrame", tickFunc)
	screen.AddEvents(global, canvas, document)

	item.MultiplierItem().Register()
	item.ClickerItem().Register()
}

func resizeCanvas() {
	width := global.Get("innerWidth").Float()
	height := global.Get("innerHeight").Float()
	canvas.Set("width", width)
	canvas.Set("height", height)
}

func tick() {
	drawScreenTick()
}

func drawScreenTick() {
	screen.CurrentScreen.Render(global, canvas, document)
}

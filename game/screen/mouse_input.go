package screen

import (
	"syscall/js"
)

func AddEvents(global, canvas, document js.Value) {
	document.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		CurrentScreen.OnClick(event.Get("button").Int())
		return nil
	}))

	document.Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		MouseX = float32(event.Get("clientX").Float())
		MouseY = float32(event.Get("clientY").Float())
		return nil
	}))
}

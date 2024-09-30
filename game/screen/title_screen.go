package screen

import (
	"circle-clicker/game/utility"
	"syscall/js"
)

func TitleScreen() *Screen {
	screen := Screen{
		Render:  TitleRender,
		OnClick: TitleOnClick,
		OnInit:  TitleOnInit,
	}

	return &screen
}

func TitleOnInit(global, canvas, document js.Value) {
}

func TitleOnClick(button int) {
	for _, b := range Buttons {
		if button == 0 && b.IsHovered() {
			b.Func()
		}
	}
}

func TitleRender(global, canvas, document js.Value) {
	utility.BeginRender(canvas, "2d")
	buttonW, buttonH := float32(250), float32(50)
	titleSize := float32(70)
	margin := float32(20)
	x, y := utility.GetCenter(buttonW, buttonH)
	AddButton(Button{
		Func: func() {
			GameScreen().SetScreen(global, canvas, document)
		},
		X:      x,
		Y:      y + (titleSize + margin),
		Width:  buttonW,
		Height: buttonH,
		ID:     "Start",
	})
	shadowFunc := func(ctx js.Value) {
		utility.SetShadow(30, "#000000")
	}
	utility.DrawBackground()
	utility.DrawFilledRoundedRect(x, y+(titleSize+margin), buttonW, buttonH, 24, "#ffffff", shadowFunc)
	utility.DrawCenteredFilledText("Click to continue", x, y+(titleSize+margin), buttonW, buttonH, 24, "#000000", shadowFunc)
	utility.DrawCenteredFilledText("Circle Clicker", 0, 0, float32(canvas.Get("width").Float()), float32(canvas.Get("height").Float()), titleSize, "#ffffff", shadowFunc)
	utility.EndRender()
}

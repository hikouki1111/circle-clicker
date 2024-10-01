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
	document.Set("title", "Circle Clicker")
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
	yOffset := titleSize + margin
	button := AddButton(Button{
		Func: func() {
			GameScreen().SetScreen(global, canvas, document)
		},
		X:      x,
		Y:      y + yOffset,
		Width:  buttonW,
		Height: buttonH,
		ID:     "Start",
	})
	shadowFunc := func(ctx js.Value) {
		utility.SetShadow(30, "#000000")
	}
	utility.DrawBackground()
	utility.DrawCenteredFilledText("Circle Clicker", 0, 0, float32(canvas.Get("width").Float()), float32(canvas.Get("height").Float()), titleSize, "#ffffff", shadowFunc)
	utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 24, "#ffffff", shadowFunc)
	utility.DrawCenteredFilledText("Click to continue", button.X, button.Y, button.Width, button.Height, 24, "#000000", shadowFunc)
	yOffset += buttonH + margin

	button = AddButton(Button{
		Func: func() {
			SettingsScreen().SetScreen(global, canvas, document)
		},
		X:      x,
		Y:      y + yOffset,
		Width:  buttonW,
		Height: buttonH,
		ID:     "Settings",
	})
	utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 24, "#ffffff", shadowFunc)
	utility.DrawCenteredFilledText("Settings", button.X, button.Y, button.Width, button.Height, 24, "#000000", shadowFunc)
	yOffset += buttonH + margin

	xOffset := margin
	button = AddButton(Button{
		Func: func() {
			global.Get("navigator").Get("clipboard").Call("writeText", document.Get("location").Get("href").String())
		},
		X:      xOffset,
		Y:      float32(canvas.Get("height").Float()) - (50 + margin),
		Width:  50,
		Height: 50,
		ID:     "Share",
	})
	utility.DrawImage(button.X, button.Y, button.Width, button.Height, "assets/share.svg")
	xOffset += button.Width + margin

	button = AddButton(Button{
		Func: func() {
			global.Get("window").Call("open", "https://github.com/hikouki1111/circle-clicker", "_blank")
		},
		X:      xOffset,
		Y:      float32(canvas.Get("height").Float()) - (50 + margin),
		Width:  50,
		Height: 50,
		ID:     "Github",
	})
	utility.DrawImage(button.X, button.Y, button.Width, button.Height, "assets/github-mark-white.svg")

	utility.EndRender()
}

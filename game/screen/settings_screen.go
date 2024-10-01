package screen

import (
	"circle-clicker/game/utility"
	"syscall/js"
)

var WaveAnimation = true

func SettingsScreen() *Screen {
	return &Screen{
		Render:  SettingsRender,
		OnInit:  SettingsOnInit,
		OnClick: SettingsOnClick,
	}
}

func SettingsOnInit(global, canvas, document js.Value) {

}

func SettingsOnClick(button int) {
	for _, b := range Buttons {
		if button == 0 && b.IsHovered() {
			b.Func()
		}
	}
}

func SettingsRender(global, canvas, document js.Value) {
	utility.BeginRender(canvas, "2d")
	margin := float32(20)
	shadowFunc := func(ctx js.Value) {
		utility.SetShadow(30, "#000000")
	}
	utility.DrawBackground()
	buttonW, buttonH := float32(50), float32(50)
	yOffset := margin
	button := AddButton(Button{
		Func: func() {
			WaveAnimation = !WaveAnimation
		},
		X:      margin,
		Y:      yOffset,
		Width:  buttonW,
		Height: buttonH,
		ID:     "Wave Animation",
	})
	utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 5, "#ffffff", shadowFunc)
	utility.DrawFilledText(button.ID, button.X+button.Width+margin, button.Y+button.Height/1.5, 24, "#ffffff")
	if WaveAnimation {
		utility.DrawCenteredFilledText("O", button.X, button.Y, button.Width, button.Height, 24, "#000000")
	}
	yOffset += button.Width + margin

	buttonW, buttonH = float32(250), float32(50)
	button = AddButton(Button{
		Func: func() {
			TitleScreen().SetScreen(global, canvas, document)
		},
		X:      margin,
		Y:      (float32(canvas.Get("height").Float()) - buttonH) - margin,
		Width:  buttonW,
		Height: buttonH,
		ID:     "Back",
	})
	utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 24, "#ffffff", shadowFunc)
	utility.DrawCenteredFilledText(button.ID, button.X, button.Y, button.Width, button.Height, 24, "#000000")

	utility.EndRender()
}

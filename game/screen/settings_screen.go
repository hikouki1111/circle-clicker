package screen

import (
	"circle-clicker/game/utility"
	"fmt"
	"syscall/js"
)

var (
	WaveAnimation    = true
	CountUPAnimation = true
)

func SettingsScreen() *Screen {
	return &Screen{
		Render:  SettingsRender,
		OnInit:  SettingsOnInit,
		OnClick: SettingsOnClick,
	}
}

func SettingsOnInit(global, canvas, document js.Value) {
	cookies := utility.ParseCookie(document)
	if cookies != nil {
		b, err := utility.ParseBool(cookies["waveanimation"])
		if err == nil {
			WaveAnimation = b
		}

		b, err = utility.ParseBool(cookies["countupanimation"])
		if err == nil {
			CountUPAnimation = b
		}
	}
}

func SettingsOnClick(button int) {
	for _, b := range Buttons {
		if button == 0 && b.IsHovered() {
			b.Func()
		}
	}
}

func SettingsRender(global, canvas, document js.Value) {
	storeSettingsCookie(document)

	utility.BeginRender(canvas, "2d")
	margin := float32(20)
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
	utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 5, "#ffffff", utility.GetSF(button.IsHovered()))
	utility.DrawFilledText(button.ID, button.X+button.Width+margin, button.Y+button.Height/1.5, 24, "#ffffff")
	if WaveAnimation {
		utility.DrawCenteredFilledText("O", button.X, button.Y, button.Width, button.Height, 24, "#000000")
	}
	yOffset += button.Width + margin

	button = AddButton(Button{
		Func: func() {
			CountUPAnimation = !CountUPAnimation
		},
		X:      margin,
		Y:      yOffset,
		Width:  buttonW,
		Height: buttonH,
		ID:     "Count UP Animation",
	})
	utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 5, "#ffffff", utility.GetSF(button.IsHovered()))
	utility.DrawFilledText(button.ID, button.X+button.Width+margin, button.Y+button.Height/1.5, 24, "#ffffff")
	if CountUPAnimation {
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
	utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 24, "#ffffff", utility.GetSF(button.IsHovered()))
	utility.DrawCenteredFilledText(button.ID, button.X, button.Y, button.Width, button.Height, 24, "#000000")

	utility.EndRender()
}

func storeSettingsCookie(document js.Value) {
	document.Set("cookie", fmt.Sprintf("waveanimation=%t;", WaveAnimation))
	document.Set("cookie", fmt.Sprintf("countupanimation=%t;", CountUPAnimation))
}

package screen

import (
	"circle-clicker/game/item"
	"circle-clicker/game/utility"
	"fmt"
	"strconv"
	"syscall/js"
)

func GameScreen() *Screen {
	screen := Screen{
		Render:  GameRender,
		OnClick: GameOnClick,
		OnInit:  GameOnInit,
	}

	return &screen
}

var (
	lastCircles int
	waveAnims   []utility.Animation
	cntUPAnims  []struct {
		utility.Animation
		int
	}
)

func GameOnInit(global, canvas, document js.Value) {
	waveAnims = []utility.Animation{}
	cntUPAnims = []struct {
		utility.Animation
		int
	}{}

	cookies := utility.ParseCookie(document)
	if cookies != nil {
		i, err := strconv.Atoi(cookies["multiplier"])
		if err == nil {
			item.Multiplier = i
		}

		i, err = strconv.Atoi(cookies["circles"])
		if err == nil {
			item.Circles = i
		}

		i, err = strconv.Atoi(cookies["clickers"])
		if err == nil {
			item.Clickers = i
		}

		i, err = strconv.Atoi(cookies["totalcircles"])
		if err == nil {
			item.TotalCircles = i
		}

		b, err := utility.ParseBool(cookies["waveanimation"])
		if err == nil {
			WaveAnimation = b
		}

		b, err = utility.ParseBool(cookies["countupanimation"])
		if err == nil {
			CountUPAnimation = b
		}
	}
	lastCircles = item.Circles
}

func GameOnClick(button int) {
	for _, b := range Buttons {
		if button == 0 && b.IsHovered() {
			b.Func()
		}
	}
}

func GameRender(global, canvas, document js.Value) {
	document.Set("title", fmt.Sprintf("%d - Circle Clicker", item.Circles))
	storeGameCookie(document)
	margin := float32(20)

	for _, i := range item.Items {
		i.OnUpdate()
	}

	utility.BeginRender(canvas, "2d")
	utility.DrawBackground()
	circleRadius := float32(150.0)
	if lastCircles < item.Circles {
		waveAnims = append(waveAnims,
			*utility.NewAnimation(
				circleRadius,
				circleRadius,
				float32(canvas.Get("width").Float()),
				float32(canvas.Get("width").Float()),
				5000,
				utility.LinerMode,
			),
		)

		cntUPAnims = append(cntUPAnims,
			struct {
				utility.Animation
				int
			}{
				*utility.NewAnimation(
					MouseX,
					MouseY,
					MouseX,
					MouseY-250,
					1000,
					utility.LinerMode,
				),
				item.Circles - lastCircles,
			},
		)
	}

	circleX, circleY := utility.GetCenter(circleRadius*2, circleRadius*2)
	button := AddButton(Button{
		Func: func() {
			item.Circles += item.Multiplier
			item.TotalCircles += item.Multiplier
		},
		X:      circleX,
		Y:      circleY,
		Width:  circleRadius * 2,
		Height: circleRadius * 2,
		ID:     "Circle",
	})

	circleX, circleY = circleX+circleRadius, circleY+circleRadius
	text := strconv.Itoa(item.Circles)
	utility.DrawFilledCircle(circleX, circleY, circleRadius, "#ffffff", utility.GetSF(button.IsHovered()))
	updateAnims()
	utility.DrawCenteredFilledText(text, 0, 0, float32(canvas.Get("width").Float()), float32(canvas.Get("height").Float())/2, 48, "#ffffff", utility.GetSF(false))

	detailSize := float32(24)
	dYOffset := detailSize + margin
	utility.DrawFilledText(fmt.Sprintf("Multiplier %d", item.Multiplier), margin, dYOffset, detailSize, "#ffffff", utility.GetSF(false))
	dYOffset += detailSize + margin
	utility.DrawFilledText(fmt.Sprintf("Clicker %d", item.Clickers), margin, dYOffset, detailSize, "#ffffff", utility.GetSF(false))
	dYOffset += detailSize + margin

	buttonW, buttonH := float32(250), float32(50)
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

	iYOffset := margin
	for _, i := range item.Items {
		button = AddButton(Button{
			Func: func() {
				i.OnBuy()
			},
			X:      (float32(canvas.Get("width").Float()) - buttonW) - margin,
			Y:      iYOffset,
			Width:  buttonW,
			Height: buttonH,
			ID:     i.Name,
		})
		utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 24, "#ffffff", utility.GetSF(button.IsHovered()))
		utility.DrawCenteredFilledText(fmt.Sprintf("O%d %s", i.Cost, i.Name), button.X, button.Y, button.Width, button.Height, 18, "#000000")
		button = AddButton(Button{
			Func: func() {
				for {
					suc := i.OnBuy()
					i.OnUpdate()
					if !suc {
						break
					}
				}
			},
			X:      (float32(canvas.Get("width").Float()) - (buttonW + margin + 50)) - margin,
			Y:      iYOffset,
			Width:  50,
			Height: 50,
			ID:     i.Name + "All",
		})
		utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 5, "#ffffff", utility.GetSF(button.IsHovered()))
		utility.DrawCenteredFilledText("All", button.X, button.Y, button.Width, button.Height, 18, "#000000")
		iYOffset += buttonH + margin
	}

	text = fmt.Sprintf("Total %d", item.TotalCircles)
	w, h := utility.GetFontSize(text, detailSize)
	utility.DrawFilledText(text, (float32(canvas.Get("width").Float())-w*2)-margin, float32(canvas.Get("height").Float())-(h/2+margin), detailSize, "#ffffff", utility.GetSF(false))

	utility.EndRender()
	lastCircles = item.Circles
}

func storeGameCookie(document js.Value) {
	document.Set("cookie", fmt.Sprintf("multiplier=%d;", item.Multiplier))
	document.Set("cookie", fmt.Sprintf("circles=%d;", item.Circles))
	document.Set("cookie", fmt.Sprintf("clickers=%d;", item.Clickers))
	document.Set("cookie", fmt.Sprintf("totalcircles=%d;", item.TotalCircles))
	document.Set("cookie", fmt.Sprintf("waveanimation=%t;", WaveAnimation))
	document.Set("cookie", fmt.Sprintf("countupanimation=%t;", CountUPAnimation))
}

func updateAnims() {
	circleX, circleY := utility.GetCenter(150*2, 150*2)
	circleX, circleY = circleX+150, circleY+150

	for i, a := range waveAnims {
		if WaveAnimation {
			utility.DrawCircle(circleX, circleY, a.X, 1, "#ffffff")
		}

		if len(waveAnims) > i {
			if a.IsFinished() {
				if len(waveAnims) > 1 {
					waveAnims = append(waveAnims[:i], waveAnims[i+1:]...)
				} else {
					waveAnims = []utility.Animation{}
				}
			} else {
				waveAnims[i].Update()
			}
		}
	}

	for i, m := range cntUPAnims {
		if CountUPAnimation {
			utility.DrawFilledText(fmt.Sprintf("+%d", m.int), m.X, m.Y, 48, "#000000")
		}

		if len(cntUPAnims) > i {
			if m.IsFinished() {
				if len(cntUPAnims) > 1 {
					cntUPAnims = append(cntUPAnims[:i], cntUPAnims[i+1:]...)
				} else {
					cntUPAnims = []struct {
						utility.Animation
						int
					}{}
				}
			} else {
				cntUPAnims[i].Update()
			}
		}
	}
}

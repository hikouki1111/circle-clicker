package screen

import (
	"circle-clicker/game/item"
	"circle-clicker/game/utility"
	"fmt"
	"strconv"
	"strings"
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
	cntUPAnims  []utility.Animation
)

func GameOnInit(global, canvas, document js.Value) {
	lastCircles = item.Circles
	waveAnims = []utility.Animation{}
	cntUPAnims = []utility.Animation{}

	cookies := parseCookie(document)
	if cookies != nil {
		i, err := strconv.Atoi(cookies["multiplier"])
		if err == nil {
			item.Multiplier = i
		} else {
			fmt.Println(err)
		}

		i, err = strconv.Atoi(cookies["circles"])
		if err == nil {
			item.Circles = i
		} else {
			fmt.Println(err)
		}

		i, err = strconv.Atoi(cookies["clickers"])
		if err == nil {
			item.Clickers = i
		} else {
			fmt.Println(err)
		}

		i, err = strconv.Atoi(cookies["totalcircles"])
		if err == nil {
			item.TotalCircles = i
		} else {
			fmt.Println(err)
		}
	}
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
	storeCookie(document)
	margin := float32(20)

	for _, i := range item.Items {
		i.OnUpdate()
	}

	utility.BeginRender(canvas, "2d")
	utility.DrawBackground()
	circleRadius := float32(150.0)
	circleX, circleY := utility.GetCenter(circleRadius*2, circleRadius*2)
	button := AddButton(Button{
		Func: func() {
			item.Circles += item.Multiplier
			item.TotalCircles += item.Multiplier
			cntUPAnims = append(cntUPAnims,
				*utility.NewAnimation(
					MouseX,
					MouseY,
					MouseX,
					MouseY-250,
					1000,
					utility.LinerMode,
				),
			)
		},
		X:      circleX,
		Y:      circleY,
		Width:  circleRadius * 2,
		Height: circleRadius * 2,
		ID:     "Circle",
	})

	circleX, circleY = circleX+circleRadius, circleY+circleRadius
	text := strconv.Itoa(item.Circles)
	shadowFunc := func(ctx js.Value) {
		utility.SetShadow(30, "#000000")
	}
	if button != nil && button.IsHovered() {
		shadowFunc = func(ctx js.Value) {
			utility.SetShadow(30, "#ffffff")
		}
	}

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
	}

	utility.DrawFilledCircle(circleX, circleY, circleRadius, "#ffffff", shadowFunc)
	updateAnims(circleX, circleY)
	shadowFunc = func(ctx js.Value) {
		utility.SetShadow(30, "#000000")
	}
	utility.DrawCenteredFilledText(text, 0, 0, float32(canvas.Get("width").Float()), float32(canvas.Get("height").Float())/2, 48, "#ffffff", shadowFunc)

	detailSize := float32(24)
	dYOffset := detailSize + margin
	utility.DrawFilledText(fmt.Sprintf("Multiplier %d", item.Multiplier), margin, dYOffset, detailSize, "#ffffff", shadowFunc)
	dYOffset += detailSize + margin
	utility.DrawFilledText(fmt.Sprintf("Clicker %d", item.Clickers), margin, dYOffset, detailSize, "#ffffff", shadowFunc)
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
	utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 24, "#ffffff", shadowFunc)
	utility.DrawCenteredFilledText(button.ID, button.X, button.Y, button.Width, button.Height, 24, "#000000")

	iYOffset := margin
	for _, i := range item.Items {
		AddButton(Button{
			Func: func() {
				if item.Circles >= i.Cost {
					item.Circles -= i.Cost
					i.OnBuy()
				}
			},
			X:      (float32(canvas.Get("width").Float()) - buttonW) - margin,
			Y:      iYOffset,
			Width:  buttonW,
			Height: buttonH,
			ID:     i.Name,
		})
		button = GetButton(i.Name)
		utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 24, "#ffffff", shadowFunc)
		utility.DrawCenteredFilledText(fmt.Sprintf("O%d %s", i.Cost, i.Name), button.X, button.Y, button.Width, button.Height, 18, "#000000")
		iYOffset += button.Height + margin
	}

	utility.EndRender()
	lastCircles = item.Circles
}

func storeCookie(document js.Value) {
	document.Set("cookie", fmt.Sprintf("multiplier=%d;", item.Multiplier))
	document.Set("cookie", fmt.Sprintf("circles=%d;", item.Circles))
	document.Set("cookie", fmt.Sprintf("clickers=%d;", item.Clickers))
	document.Set("cookie", fmt.Sprintf("totalcircles=%d;", item.TotalCircles))
}

func parseCookie(document js.Value) map[string]string {
	cookieStr := document.Get("cookie").String()
	cookies := map[string]string{}

	if cookieStr == "" {
		return nil
	}

	cookieArray := strings.Split(cookieStr, "; ")
	for _, cookie := range cookieArray {
		pair := strings.SplitN(cookie, "=", 2)
		if len(pair) == 2 {
			cookies[pair[0]] = pair[1]
		}
	}

	return cookies
}

func updateAnims(circleX, circleY float32) {
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

	for i, a := range cntUPAnims {
		if CountUPAnimation {
			utility.DrawFilledText(fmt.Sprintf("+%d", item.Multiplier), a.X, a.Y, 48, "#000000")
		}

		if len(cntUPAnims) > i {
			if a.IsFinished() {
				if len(cntUPAnims) > 1 {
					cntUPAnims = append(cntUPAnims[:i], cntUPAnims[i+1:]...)
				} else {
					cntUPAnims = []utility.Animation{}
				}
			} else {
				cntUPAnims[i].Update()
			}
		}
	}
}

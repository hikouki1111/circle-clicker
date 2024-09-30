package screen

import (
	"circle-clicker/game/utility"
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

var (
	Circles    int
	Multiplier = 1
)

func GameScreen() *Screen {
	screen := Screen{
		Render:  GameRender,
		OnClick: GameOnClick,
		OnInit:  GameOnInit,
	}

	return &screen
}

func GameOnInit(global, canvas, document js.Value) {
	cookies := parseCookie(document)
	if cookies == nil {
		storeCookie(document)
	} else {
		i, err := strconv.Atoi(cookies["multiplier"])
		if err != nil {
			Multiplier = i
		}

		i, err = strconv.Atoi(cookies["circles"])
		if err != nil {
			Circles = i
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
	utility.BeginRender(canvas, "2d")
	circleRadius := float32(150.0)
	circleX, circleY := utility.GetCenter(circleRadius*2, circleRadius*2)

	AddButton(Button{
		Func: func() {
			Circles += Multiplier
			document.Set("title", fmt.Sprintf("%d - Circle Clicker", Circles))
			storeCookie(document)
		},
		X:      circleX,
		Y:      circleY,
		Width:  circleRadius * 2,
		Height: circleRadius * 2,
		ID:     "Circle",
	})

	circleX, circleY = circleX+circleRadius, circleY+circleRadius
	text := strconv.Itoa(Circles)
	utility.DrawBackground()
	button := GetButton("Circle")
	shadowFunc := func(ctx js.Value) {
		utility.SetShadow(30, "#000000")
	}
	if button != nil && button.IsHovered() {
		shadowFunc = func(ctx js.Value) {
			utility.SetShadow(30, "#ffffff")
		}
	}
	utility.DrawFilledCircle(circleX, circleY, circleRadius, "#ffffff", shadowFunc)
	shadowFunc = func(ctx js.Value) {
		utility.SetShadow(30, "#000000")
	}
	detailSize := float32(24)
	utility.DrawFilledText(fmt.Sprintf("Multiplier %d", Multiplier), 0, detailSize, detailSize, "#ffffff", shadowFunc)
	utility.DrawCenteredFilledText(text, 0, 0, float32(canvas.Get("width").Float()), float32(canvas.Get("height").Float())/2, 48, "#ffffff", shadowFunc)
	utility.EndRender()
}

func storeCookie(document js.Value) {
	document.Set("cookie", fmt.Sprintf("multiplier=%d;", Multiplier))
	document.Set("cookie", fmt.Sprintf("circles=%d;", Circles))
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

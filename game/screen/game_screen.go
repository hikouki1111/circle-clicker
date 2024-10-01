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

func GameOnInit(global, canvas, document js.Value) {
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
	for _, i := range item.Items {
		i.OnUpdate()
	}

	utility.BeginRender(canvas, "2d")
	circleRadius := float32(150.0)
	circleX, circleY := utility.GetCenter(circleRadius*2, circleRadius*2)

	AddButton(Button{
		Func: func() {
			item.Circles += item.Multiplier
			document.Set("title", fmt.Sprintf("%d - Circle Clicker", item.Circles))
			storeCookie(document)
		},
		X:      circleX,
		Y:      circleY,
		Width:  circleRadius * 2,
		Height: circleRadius * 2,
		ID:     "Circle",
	})

	circleX, circleY = circleX+circleRadius, circleY+circleRadius
	text := strconv.Itoa(item.Circles)
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
	utility.DrawCenteredFilledText(text, 0, 0, float32(canvas.Get("width").Float()), float32(canvas.Get("height").Float())/2, 48, "#ffffff", shadowFunc)

	detailSize := float32(24)
	dYOffset := detailSize
	utility.DrawFilledText(fmt.Sprintf("Multiplier %d", item.Multiplier), 0, dYOffset, detailSize, "#ffffff", shadowFunc)
	dYOffset += detailSize
	utility.DrawFilledText(fmt.Sprintf("Clicker %d", item.Clickers), 0, dYOffset, detailSize, "#ffffff", shadowFunc)
	dYOffset += detailSize

	iYOffset := float32(0)
	for _, i := range item.Items {
		AddButton(Button{
			Func: func() {
				if item.Circles >= i.Cost {
					item.Circles -= i.Cost
					i.OnBuy()
					document.Set("title", fmt.Sprintf("%d - Circle Clicker", item.Circles))
					storeCookie(document)
				}
			},
			X:      float32(canvas.Get("width").Float()) - 200,
			Y:      iYOffset,
			Width:  200,
			Height: 50,
			ID:     i.Name,
		})
		button = GetButton(i.Name)
		utility.DrawFilledRoundedRect(button.X, button.Y, button.Width, button.Height, 24, "#ffffff", shadowFunc)
		utility.DrawCenteredFilledText(fmt.Sprintf("●%d %s", i.Cost, i.Name), button.X, button.Y, button.Width, button.Height, 18, "#000000")
		iYOffset += button.Height + 10
	}

	utility.EndRender()
}

func storeCookie(document js.Value) {
	document.Set("cookie", fmt.Sprintf("multiplier=%d;", item.Multiplier))
	document.Set("cookie", fmt.Sprintf("circles=%d;", item.Circles))
	document.Set("cookie", fmt.Sprintf("clickers=%d;", item.Clickers))
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

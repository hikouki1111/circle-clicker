package screen

import (
	"circle-clicker/game/utility"
	"fmt"
	"strconv"
	"syscall/js"
)

var (
	Circles int
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
			Circles++
			document.Set("title", fmt.Sprintf("%d - Circle Clicker", Circles))
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
	utility.DrawCenteredFilledText(text, 0, 0, float32(canvas.Get("width").Float()), float32(canvas.Get("height").Float())/2, 48, "#ffffff", shadowFunc)
	utility.EndRender()
}

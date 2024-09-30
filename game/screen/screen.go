package screen

import (
	"syscall/js"
)

var (
	MouseX        float32
	MouseY        float32
	CurrentScreen Screen
	Buttons       []Button
)

type Screen struct {
	Render  func(global, canvas, document js.Value)
	OnInit  func(global, canvas, document js.Value)
	OnClick func(pressed int)
}

func AddButton(button Button) {
	for i := range Buttons {
		if Buttons[i].ID == button.ID {
			Buttons[i] = button
			return
		}
	}

	Buttons = append(Buttons, button)
}

func RemoveButton(id string) {
	for i := range Buttons {
		if Buttons[i].ID == id {
			Buttons = append(Buttons[:i], Buttons[i+1:]...)
			return
		}
	}
}

func GetButton(id string) *Button {
	for i := range Buttons {
		if Buttons[i].ID == id {
			return &Buttons[i]
		}
	}

	return nil
}

func (s *Screen) SetScreen(global, canvas, document js.Value) {
	Buttons = []Button{}
	CurrentScreen = *s
	CurrentScreen.OnInit(global, canvas, document)
}

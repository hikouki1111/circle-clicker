package screen

type Button struct {
	Func   func()
	X      float32
	Y      float32
	Width  float32
	Height float32
	ID     string
}

func (b *Button) IsHovered() bool {
	return MouseX >= b.X && MouseX <= b.X+b.Width && MouseY >= b.Y && MouseY <= b.Y+b.Height
}

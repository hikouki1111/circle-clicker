package utility

import "time"

type AnimateMode int

const (
	LinerMode AnimateMode = iota
	EaseInMode
	EaseOutMode
	EaseInOutMode
)

type Animation struct {
	X         float32
	Y         float32
	StartX    float32
	StartY    float32
	GoalX     float32
	GoalY     float32
	Duration  float32
	StartTime time.Time
	Mode      AnimateMode
}

func NewAnimation(x, y, goalX, goalY, duration float32, mode AnimateMode) *Animation {
	return &Animation{
		X:         x,
		Y:         y,
		StartX:    x,
		StartY:    y,
		GoalX:     goalX,
		GoalY:     goalY,
		Duration:  duration,
		StartTime: time.Now(),
		Mode:      mode,
	}
}

func (a *Animation) getProgress() float32 {
	elapsed := float32(time.Since(a.StartTime).Seconds()) * 1000
	progress := elapsed / a.Duration

	if progress > 1 {
		progress = 1
	}
	return progress
}

func (a *Animation) Update() {
	var progress float32

	switch a.Mode {
	case LinerMode:
		progress = a.getProgress()
	case EaseInMode:
		progress = EaseIn(a.getProgress())
	case EaseOutMode:
		progress = EaseOut(a.getProgress())
	case EaseInOutMode:
		progress = EaseInOut(a.getProgress())
	}

	a.X = Lerp(a.StartX, a.GoalX, progress)
	a.Y = Lerp(a.StartY, a.GoalY, progress)
}

package utility

func Lerp(start, end, t float32) float32 {
	return start + (end-start)*t
}

func EaseIn(t float32) float32 {
	return t * t
}

func EaseOut(t float32) float32 {
	return 1 - EaseIn(1-t)
}

func EaseInOut(t float32) float32 {
	return t * t * t * (t*(t*6-15) + 10)
}

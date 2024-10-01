package item

import "circle-clicker/game/utility"

var (
	Clickers  = 0
	stopwatch = utility.NewStopwatch()
)

func ClickerItem() *Item {
	return &Item{
		"Clicker",
		100,
		func() {
			Clickers++
		},
		func() {
			if stopwatch.IsFinished(1000, true) {
				Circles += Clickers * Multiplier
			}
		},
	}
}

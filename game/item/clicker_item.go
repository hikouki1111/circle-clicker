package item

import "circle-clicker/game/utility"

var (
	Clickers  = 0
	stopwatch = utility.NewStopwatch()
)

func ClickerItem() *Item {
	return &Item{
		"Clicker",
		300,
		300,
		func(i *Item) {
			Clickers++
		},
		func(i *Item) {
			i.Cost = (Clickers + 1) * i.InitCost
			if stopwatch.IsFinished(1000, true) {
				Circles += Clickers * Multiplier
				TotalCircles += Clickers * Multiplier
			}
		},
	}
}

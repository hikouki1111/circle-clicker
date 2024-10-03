package item

import "circle-clicker/game/utility"

var (
	Clickers  = 0
	stopwatch = utility.NewStopwatch()
)

func ClickerItem() *Item {
	i := &Item{
		Name:     "Clicker",
		InitCost: 50,
		Cost:     50,
	}

	i.OnBuy = func() bool {
		if Circles >= i.Cost {
			Circles -= i.Cost
			Clickers++

			return true
		}

		return false
	}

	i.OnUpdate = func() {
		if Clickers > 0 {
			i.Cost = Clickers * (i.InitCost * (Clickers + 1))
		}

		if stopwatch.IsFinished(1000, true) {
			Circles += Clickers * Multiplier
			TotalCircles += Clickers * Multiplier
		}
	}

	return i
}

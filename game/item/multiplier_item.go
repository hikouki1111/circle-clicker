package item

var Multiplier = 1

func MultiplierItem() *Item {
	i := &Item{
		Name:     "Multiplier",
		InitCost: 100,
		Cost:     100,
	}

	i.OnBuy = func() bool {
		if Circles >= i.Cost {
			Circles -= i.Cost
			Multiplier++

			return true
		}

		return false
	}

	i.OnUpdate = func() {
		if Multiplier > 1 {
			i.Cost = Multiplier * (i.InitCost * Multiplier)
		}
	}

	return i
}

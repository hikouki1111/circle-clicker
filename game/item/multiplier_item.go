package item

var Multiplier = 1

func MultiplierItem() *Item {
	i := &Item{
		Name:     "Multiplier",
		InitCost: 400,
		Cost:     400,
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
		i.Cost = Multiplier * i.InitCost
	}

	return i
}

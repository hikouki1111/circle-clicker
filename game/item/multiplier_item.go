package item

var Multiplier = 1

func MultiplierItem() *Item {
	return &Item{
		Name:     "Multiplier",
		InitCost: 400,
		Cost:     400,
		OnBuy: func(i *Item) {
			Multiplier++
		},
		OnUpdate: func(i *Item) {
			i.Cost = Multiplier * i.InitCost
		},
	}
}

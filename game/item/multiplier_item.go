package item

var Multiplier = 1

func MultiplierItem() *Item {
	return &Item{
		Name: "Multiplier",
		Cost: 500,
		OnBuy: func() {
			Multiplier++
		},
		OnUpdate: func() {},
	}
}

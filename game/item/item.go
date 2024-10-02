package item

var (
	Items        []Item
	Circles      int
	TotalCircles int
)

type Item struct {
	Name     string
	InitCost int
	Cost     int
	OnBuy    func(i *Item)
	OnUpdate func(i *Item)
}

func (i *Item) Register() {
	Items = append(Items, *i)
}

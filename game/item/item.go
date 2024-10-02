package item

var (
	Items        []*Item
	Circles      int
	TotalCircles int
)

type Item struct {
	Name     string
	InitCost int
	Cost     int
	OnBuy    func() bool
	OnUpdate func()
}

func (i *Item) Register() {
	Items = append(Items, i)
}

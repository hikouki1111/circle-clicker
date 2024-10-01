package item

var (
	Items   []Item
	Circles int
)

type Item struct {
	Name     string
	Cost     int
	OnBuy    func()
	OnUpdate func()
}

func (i *Item) Register() {
	Items = append(Items, *i)
}

package product

type IKeyboard interface {
	SetLogo(logo string)
	SetPrice(price float64)
	GetLogo() string
	GetPrice() float64
}

type Keyboard struct {
	logo  string
	price float64
}

func (k *Keyboard) SetLogo(logo string) {
	k.logo = logo
}

func (k *Keyboard) SetPrice(price float64) {
	k.price = price
}

func (k *Keyboard) GetLogo() string {
	return k.logo
}

func (k *Keyboard) GetPrice() float64 {
	return k.price
}

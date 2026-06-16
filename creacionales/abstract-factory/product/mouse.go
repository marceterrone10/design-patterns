package product

type IMouse interface {
	SetLogo(logo string)
	SetPrice(price float64)
	GetLogo() string
	GetPrice() float64
}

type Mouse struct {
	logo  string
	price float64
}

func (m *Mouse) SetLogo(logo string) {
	m.logo = logo
}

func (m *Mouse) SetPrice(price float64) {
	m.price = price
}

func (m *Mouse) GetLogo() string {
	return m.logo
}

func (m *Mouse) GetPrice() float64 {
	return m.price
}

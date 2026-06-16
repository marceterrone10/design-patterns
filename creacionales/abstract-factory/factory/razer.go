package factory

import "design-patterns/creacionales/abstract-factory/product"

type RazerMouse struct {
	product.Mouse
}

type RazerKeyboard struct {
	product.Keyboard
}

type Razer struct{}

func (r *Razer) CreateMouse() product.IMouse {
	m := &RazerMouse{}
	m.SetLogo("razer")
	m.SetPrice(100)
	return m
}

func (r *Razer) CreateKeyboard() product.IKeyboard {
	k := &RazerKeyboard{}
	k.SetLogo("razer")
	k.SetPrice(200)
	return k
}

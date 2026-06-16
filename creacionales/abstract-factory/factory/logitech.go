package factory

import "design-patterns/creacionales/abstract-factory/product"

type LogitechMouse struct {
	product.Mouse
}

type LogitechKeyboard struct {
	product.Keyboard
}

type Logitech struct{}

func (l *Logitech) CreateMouse() product.IMouse {
	m := &LogitechMouse{}
	m.SetLogo("logitech")
	m.SetPrice(100)
	return m
}

func (l *Logitech) CreateKeyboard() product.IKeyboard {
	k := &LogitechKeyboard{}
	k.SetLogo("logitech")
	k.SetPrice(200)
	return k
}

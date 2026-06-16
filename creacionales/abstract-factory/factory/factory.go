package factory

import (
	"fmt"

	"design-patterns/creacionales/abstract-factory/product"
)

type IPerifericosFactory interface {
	CreateMouse() product.IMouse
	CreateKeyboard() product.IKeyboard
}

func GetPerifericosFactory(brand string) (IPerifericosFactory, error) {
	switch brand {
	case "logitech":
		return &Logitech{}, nil
	case "razer":
		return &Razer{}, nil
	default:
		return nil, fmt.Errorf("brand %s not found", brand)
	}
}

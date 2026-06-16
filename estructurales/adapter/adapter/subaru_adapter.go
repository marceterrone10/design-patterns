package adapter

import (
	"strings"

	"design-patterns/estructurales/adapter/adaptee"
)

type SubaruAdapter struct {
	subaru *adaptee.SubaruImpreza
}

func NewSubaruAdapter(subaru *adaptee.SubaruImpreza) *SubaruAdapter {
	return &SubaruAdapter{subaru: subaru}
}

func (a *SubaruAdapter) Avanzar() string {
	steps := []string{
		a.subaru.PresionarPedalFreno(),
		a.subaru.InsertarLlave(),
		a.subaru.GirarLlave(),
		a.subaru.SoltarFreno(),
		"Subaru Impreza: avanzando",
	}
	return strings.Join(steps, "\n")
}

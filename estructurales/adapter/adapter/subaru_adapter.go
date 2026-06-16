package adapter

import (
	"strings"

	"design-patterns/estructurales/adapter/adaptee"
	"design-patterns/estructurales/adapter/target"
)

type SubaruAdapter struct {
	target.Vehicle
	subaru *adaptee.SubaruImpreza
}

func NewSubaruAdapter(subaru *adaptee.SubaruImpreza) *SubaruAdapter {
	return &SubaruAdapter{
		Vehicle: target.Vehicle{Nombre: "Subaru Impreza"},
		subaru:  subaru,
	}
}

func (a *SubaruAdapter) Avanzar() string {
	steps := []string{
		a.subaru.PresionarPedalFreno(),
		a.subaru.InsertarLlave(),
		a.subaru.GirarLlave(),
		a.subaru.SoltarFreno(),
	}
	result := strings.Join(steps, "\n")
	target.PrintAvance(result, a.Nombre)
	return result
}

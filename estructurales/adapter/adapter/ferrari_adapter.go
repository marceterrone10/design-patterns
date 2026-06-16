package adapter

import (
	"strings"

	"design-patterns/estructurales/adapter/adaptee"
	"design-patterns/estructurales/adapter/target"
)

type FerrariAdapter struct {
	target.Vehicle
	ferrari *adaptee.FerrariModerna
}

func NewFerrariAdapter(ferrari *adaptee.FerrariModerna) *FerrariAdapter {
	return &FerrariAdapter{
		Vehicle: target.Vehicle{Nombre: "Ferrari moderna"},
		ferrari: ferrari,
	}
}

func (a *FerrariAdapter) Avanzar() string {
	steps := []string{
		a.ferrari.PresionarFreno(),
		a.ferrari.PonerEnNeutral(),
		a.ferrari.ApretarBotonStart(),
	}
	result := strings.Join(steps, "\n")
	target.PrintAvance(result, a.Nombre)
	return result
}

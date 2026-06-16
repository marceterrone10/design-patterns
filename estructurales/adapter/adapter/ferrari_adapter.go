package adapter

import (
	"strings"

	"design-patterns/estructurales/adapter/adaptee"
)

type FerrariAdapter struct {
	ferrari *adaptee.FerrariModerna
}

func NewFerrariAdapter(ferrari *adaptee.FerrariModerna) *FerrariAdapter {
	return &FerrariAdapter{ferrari: ferrari}
}

func (a *FerrariAdapter) Avanzar() string {
	steps := []string{
		a.ferrari.PresionarFreno(),
		a.ferrari.PonerEnNeutral(),
		a.ferrari.ApretarBotonStart(),
		"Ferrari moderna: avanzando",
	}
	return strings.Join(steps, "\n")
}
